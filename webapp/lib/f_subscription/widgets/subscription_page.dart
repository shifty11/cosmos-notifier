import 'package:badges/badges.dart';
import 'package:cosmos_notifier/api/protobuf/dart/subscription_service.pb.dart';
import 'package:cosmos_notifier/common/header_widget.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/bottom_navigation_bar_widget.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/footer_widget.dart';
import 'package:cosmos_notifier/f_subscription/services/subscription_provider.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_slidable/flutter_slidable.dart';
import 'package:responsive_framework/responsive_framework.dart';
import 'package:riverpod_messages/riverpod_messages.dart';

class SubscriptionPage extends StatelessWidget {
  final double sideBarWith = 0;

  const SubscriptionPage({Key? key}) : super(key: key);

  int getCrossAxisCount(BuildContext context) {
    if (ResponsiveWrapper.of(context).isSmallerThan(TABLET)) {
      return 1;
    }
    if (ResponsiveWrapper.of(context).isSmallerThan(DESKTOP)) {
      return 3;
    }
    return 4;
  }

  void showConfirmEnableChainDialog(BuildContext context, Subscription subscription) {
    showDialog(
      context: context,
      builder: (_) {
        return AlertDialog(
          title: Text('${subscription.isEnabled ? 'Disable' : 'Enable'} chain?'),
          content: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text("Are you sure you want to ${subscription.isEnabled ? 'disable' : 'enable'} ${subscription.name}?"),
              ],
            ),
          ),
          actions: [
            TextButton(
              onPressed: () => Navigator.pop(context),
              child: const Text('Cancel'),
            ),
            Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
              return ElevatedButton(
                style: ElevatedButton.styleFrom(primary: Styles.dangerBgColor, onPrimary: Styles.dangerTextColor),
                onPressed: () {
                  ref.read(chatroomListStateProvider.notifier).enableChain(subscription.id, subscription.name, !subscription.isEnabled);
                  Navigator.pop(context);
                },
                child: Text(subscription.isEnabled ? 'Disable' : 'Enable'),
              );
            }),
          ],
        );
      },
    );
  }

  void showConfirmDeleteDaoDialog(BuildContext context, Subscription subscription) {
    showDialog(
      context: context,
      builder: (_) {
        return AlertDialog(
          title: const Text('Delete DAO?'),
          content: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text("Are you sure you want to delete ${subscription.name}?"),
              ],
            ),
          ),
          actions: [
            TextButton(
              onPressed: () => Navigator.pop(context),
              child: const Text('Cancel'),
            ),
            Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
              return ElevatedButton(
                style: ElevatedButton.styleFrom(primary: Styles.dangerBgColor, onPrimary: Styles.dangerTextColor),
                onPressed: () {
                  ref.read(chatroomListStateProvider.notifier).deleteDao(subscription.id, subscription.name);
                  Navigator.pop(context);
                },
                child: const Text('Delete'),
              );
            }),
          ],
        );
      },
    );
  }

  Widget _buildSliderForAdmins(
      {required BuildContext context, required Subscription subscription, required WidgetRef ref, required Widget child}) {
    if (!jwtManager.isAdmin) {
      return child;
    }
    return Slidable(
      startActionPane: ActionPane(
        motion: const ScrollMotion(),
        children: [
          SlidableAction(
            onPressed: (_) {
              ref.watch(isChainsSelectedProvider)
                  ? showConfirmEnableChainDialog(context, subscription)
                  : showConfirmDeleteDaoDialog(context, subscription);
            },
            backgroundColor: Styles.dangerBgColor,
            foregroundColor: Styles.dangerTextColor,
            icon: Icons.delete,
            label: ref.watch(isChainsSelectedProvider)
                ? subscription.isEnabled
                    ? 'Disable'
                    : 'Enable'
                : 'Delete',
          ),
        ],
      ),
      child: child,
    );
  }

  Widget _subscriptionSecondLine(BuildContext context, Subscription subscription) {
    var text = Text(
      subscription.contractAddress.isEmpty
          ? ""
          : "${subscription.contractAddress.substring(0, 10)}...${subscription.contractAddress.substring(subscription.contractAddress.length - 5)}",
      style: TextStyle(fontSize: 12, color: Theme.of(context).inputDecorationTheme.hintStyle!.color),
      overflow: TextOverflow.ellipsis,
    );
    if (!jwtManager.isAdmin) {
      return text;
    }
    return Row(
      children: [
        text,
        const Spacer(),
        Tooltip(
          message: "Subscribed Telegram chats",
          child: Badge(
            toAnimate: false,
            badgeColor: Styles.telegramColor.withOpacity(0.5),
            badgeContent: Text(
              "${subscription.stats.telegram}",
              style: const TextStyle(fontSize: 12, color: Colors.white),
            ),
          ),
        ),
        const SizedBox(width: 3),
        Tooltip(
          message: "Subscribed Discord channels",
          child: Badge(
            toAnimate: false,
            badgeColor: Styles.discordColor.withOpacity(0.5),
            badgeContent: Text(
              "${subscription.stats.discord}",
              style: const TextStyle(fontSize: 12, color: Colors.white),
            ),
          ),
        ),
        const SizedBox(width: 3),
        Tooltip(
          message: "Total subscriptions",
          child: Badge(
            toAnimate: false,
            badgeColor: Colors.black.withOpacity(0.5),
            badgeContent: Text(
              "${subscription.stats.total}",
              style: const TextStyle(fontSize: 12, color: Colors.white),
            ),
          ),
        ),
      ],
    );
  }

  Widget subscriptionList() {
    return Consumer(
      builder: (BuildContext context, WidgetRef ref, Widget? child) {
        final state = ref.watch(chatroomListStateProvider);
        return state.when(
          loading: () => const Center(child: CircularProgressIndicator()),
          data: (chainChatRooms, contractChatRooms) {
            return GridView.builder(
              shrinkWrap: true,
              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: getCrossAxisCount(context),
                crossAxisSpacing: 10,
                mainAxisSpacing: 10,
                mainAxisExtent: 60,
              ),
              itemCount: ref.watch(searchedSubsProvider).filtered.length,
              itemBuilder: (BuildContext context, int index) {
                final chatRoomId = ref.read(selectedChatRoomProvider)?.id ?? Int64(0);
                final subData = ref.watch(searchedSubsProvider).filtered[index];
                final subscription = subData.subscription;
                const double sidePadding = 12;
                const double checkMarkSize = 24;
                return _buildSliderForAdmins(
                  context: context,
                  subscription: subscription,
                  ref: ref,
                  child: Container(
                    decoration: BoxDecoration(
                        color: subscription.isEnabled ? null : Colors.grey,
                        border: Border.all(
                          width: Styles.selectCardBorderWidth,
                          color: Theme.of(context).inputDecorationTheme.enabledBorder!.borderSide.color,
                        ),
                        borderRadius: const BorderRadius.all(Radius.circular(5))),
                    child: InkWell(
                      hoverColor: Theme.of(context).primaryColor.intensifyBg(context, 0.2),
                      onTap: () {
                        FocusScopeNode currentFocus = FocusScope.of(context);
                        if (!currentFocus.hasPrimaryFocus) {
                          currentFocus.focusedChild?.unfocus();
                        }
                        ref.read(chatroomListStateProvider.notifier).toggleSubscription(chatRoomId, subData.subscription);
                      },
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          const SizedBox(width: sidePadding),
                          CircleAvatar(
                              backgroundColor: Colors.transparent,
                              child: subscription.thumbnailUrl.isEmpty
                                  ? Container(
                                      decoration: BoxDecoration(
                                        border: Border.all(
                                          width: 5,
                                          color: Colors.black,
                                        ),
                                        shape: BoxShape.circle,
                                      ),
                                    )
                                  : ClipOval(
                                      child: Image.asset(
                                        subscription.thumbnailUrl,
                                        width: 40,
                                        height: 40,
                                      ),
                                    )),
                          const SizedBox(width: 10),
                          Expanded(
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Text(
                                  subscription.name,
                                  style: const TextStyle(fontSize: 20),
                                  overflow: TextOverflow.ellipsis,
                                ),
                                _subscriptionSecondLine(context, subscription),
                              ],
                            ),
                          ),
                          subscription.isSubscribed
                              ? Padding(
                                  padding: const EdgeInsets.only(right: sidePadding),
                                  child: Icon(Icons.check_circle_rounded, color: Theme.of(context).primaryColor, size: checkMarkSize),
                                )
                              : const SizedBox(width: sidePadding + checkMarkSize),
                        ],
                      ),
                    ),
                  ),
                );
              },
            );
          },
        );
      },
    );
  }

  Widget searchWidget(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      return TextField(
        onChanged: (value) => ref.watch(searchSubsProvider.notifier).state = value,
        decoration: const InputDecoration(
          prefixIcon: Icon(Icons.search),
          border: OutlineInputBorder(),
          hintText: "Search",
        ),
      );
    });
  }

  Widget subscriptionTypeWidget(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final isChainsSelected = ref.watch(isChainsSelectedProvider);
      return Column(children: [
        ToggleButtons(
          borderRadius: BorderRadius.circular(10),
          constraints: ResponsiveWrapper.of(context).isSmallerThan(TABLET)
              ? BoxConstraints(maxWidth: MediaQuery.of(context).size.width / 2 - 42)
              : null,
          isSelected: [isChainsSelected, !isChainsSelected],
          onPressed: (int index) {
            ref.read(isChainsSelectedProvider.notifier).state = index != 1;
          },
          children: [
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.end,
                children: const [
                  Icon(Icons.link),
                  SizedBox(width: 5),
                  Text("Chains"),
                ],
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Row(
                children: const [
                  Text("DAO's"),
                  SizedBox(width: 5),
                  Icon(Icons.people),
                ],
              ),
            ),
          ],
        )
      ]);
    });
  }

  void showAddDaoDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (_) {
        var addressController = TextEditingController();
        var proposalQueryController = TextEditingController();
        var formKey = GlobalKey<FormState>();
        return AlertDialog(
          title: const Text('Add DAO'),
          content: Form(
            key: formKey,
            child: SingleChildScrollView(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text("Enter the address of the DAO you want to add"),
                  const SizedBox(height: 15),
                  TextFormField(
                    controller: addressController,
                    validator: (String? text) {
                      if (text == null) {
                        return null;
                      }
                      if (text.isEmpty) {
                        return "Address cannot be empty";
                      }
                      return null;
                    },
                    decoration: const InputDecoration(
                      labelText: 'Contract Address',
                      hintText: 'juno1z3zqgz7t0hcu2fx4wusuyjq0gc2m33la8l64saunfz7vmqwa2d5sz6jnep',
                    ),
                  ),
                  const SizedBox(height: 10),
                  TextFormField(
                    controller: proposalQueryController,
                    decoration: const InputDecoration(
                      labelText: 'Proposal query (optional)',
                      hintText: '{"proposals":{"query":{"everything":{}}}}',
                    ),
                  ),
                ],
              ),
            ),
          ),
          actions: [
            TextButton(
              onPressed: () => Navigator.pop(context),
              child: const Text('Cancel'),
            ),
            Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
              return ElevatedButton(
                onPressed: () {
                  if (formKey.currentState!.validate()) {
                    var address = addressController.text.toLowerCase();
                    var proposalQuery = proposalQueryController.text;
                    ref.read(chatroomListStateProvider.notifier).addDao(address, proposalQuery);
                    Navigator.pop(context);
                  }
                },
                child: const Text('Add'),
              );
            }),
          ],
        );
      },
    );
  }

  Widget addDaoButton(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      if (ref.watch(isChainsSelectedProvider) || !jwtManager.isAdmin) {
        return Container();
      }
      return Center(
        child: OutlinedButton(
          style: OutlinedButton.styleFrom(
            minimumSize: ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? const Size(double.infinity, 50) : const Size(200, 50),
          ),
          onPressed: () {
            showAddDaoDialog(context);
          },
          child: const Text("Add DAO"),
        ),
      );
    });
  }

  Widget title(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          children: [
            Text("Subscriptions", style: Theme.of(context).textTheme.headline2),
            Tooltip(
                triggerMode: TooltipTriggerMode.tap,
                showDuration: const Duration(seconds: 5),
                message: "Select the chains and DAO's that you want to follow. "
                    "You will receive notifications about new governance proposals.",
                child: Icon(Icons.info, size: 20, color: Theme.of(context).disabledColor)),
            const Spacer(),
            ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? Container() : subscriptionTypeWidget(context),
          ],
        ),
        ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? subscriptionTypeWidget(context) : Container(),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          MessageOverlayListener(
            provider: messageProvider,
            child: Container(
              width: MediaQuery.of(context).size.width,
              padding: const EdgeInsets.only(top: Styles.topPadding, left: Styles.sidePadding, right: Styles.sidePadding),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const HeaderWidget(),
                  const SizedBox(height: 10),
                  title(context),
                  const SizedBox(height: 20),
                  searchWidget(context),
                  const SizedBox(height: 20),
                  Expanded(child: subscriptionList()),
                  const SizedBox(height: 20),
                  addDaoButton(context),
                  const Flexible(flex: 0, child: FooterWidget()),
                ],
              ),
            ),
          ),
        ],
      ),
      bottomNavigationBar: BottomNavigationBarWidget(jwtManager: jwtManager),
    );
  }
}
