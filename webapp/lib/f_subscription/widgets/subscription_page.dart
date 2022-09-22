import 'dart:html';

import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_slidable/flutter_slidable.dart';
import 'package:responsive_framework/responsive_framework.dart';
import 'package:riverpod_messages/riverpod_messages.dart';
import 'package:webapp/api/protobuf/dart/subscription_service.pb.dart';
import 'package:webapp/config.dart';
import 'package:webapp/f_home/services/message_provider.dart';
import 'package:webapp/f_home/widgets/bottom_navigation_bar_widget.dart';
import 'package:webapp/f_subscription/services/subscription_provider.dart';
import 'package:webapp/style.dart';

import '../../f_home/services/chat_id_provider.dart';

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
                },
                child: const Text('Delete'),
              );
            }),
          ],
        );
      },
    );
  }

  Widget buildSliderForAdmins({required BuildContext context, required Subscription subscription, required WidgetRef ref, required Widget child}) {
    if (!jwtManager.isAdmin) {
      return child;
    }
    return Slidable(
      startActionPane: ActionPane(
        motion: const ScrollMotion(),
        children: [
          SlidableAction(
            onPressed: (_) {
              showConfirmDeleteDaoDialog(context, subscription);
            },
            backgroundColor: Styles.dangerBgColor,
            foregroundColor: Styles.dangerTextColor,
            icon: Icons.delete,
            label: 'Delete',
          ),
        ],
      ),
      child: child,
    );
  }

  Widget subscriptionList() {
    return Expanded(
      child: Consumer(
        builder: (BuildContext context, WidgetRef ref, Widget? child) {
          final state = ref.watch(chatroomListStateProvider);
          return state.when(
            loading: () => const Center(child: CircularProgressIndicator()),
            data: (chatRooms) {
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
                  return buildSliderForAdmins(
                    context: context,
                    subscription: subscription,
                    ref: ref,
                    child: Container(
                      decoration: BoxDecoration(
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
                          ref.read(chatroomListStateProvider.notifier).toggleSubscription(chatRoomId, subData.subscription.id);
                        },
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            const SizedBox(width: sidePadding),
                            CircleAvatar(
                              backgroundColor: Colors.transparent,
                              child: subscription.thumbnailUrl.isNotEmpty ? ClipOval(
                                child: Image.asset(
                                  subscription.thumbnailUrl,
                                  width: 40,
                                  height: 40,
                                ),
                              ) : Container(
                                decoration: BoxDecoration(
                                  border: Border.all(
                                    width: 5,
                                    color: Colors.black,
                                  ),
                                  shape: BoxShape.circle,
                                ),
                              )
                            ),
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
                                  Text(
                                    "${subscription.contractAddress.substring(0, 10)}...${subscription.contractAddress.substring(subscription.contractAddress.length - 5)}",
                                    style: TextStyle(fontSize: 12, color: Theme.of(context).inputDecorationTheme.hintStyle!.color),
                                    overflow: TextOverflow.ellipsis,
                                  ),
                                ],
                              ),
                            ),
                            subscription.isSubscribed
                                ? Padding(
                                    padding: const EdgeInsets.only(right: sidePadding),
                                    child: Icon(Icons.check_circle_rounded, color: Theme.of(context).primaryColor, size: 24),
                                  )
                                : const SizedBox(width: sidePadding),
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
      ),
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

  Widget chatDropdownWidget(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final state = ref.watch(chatroomListStateProvider);
      return state.when(
        loading: () => Container(),
        data: (chatRooms) {
          if (chatRooms.isEmpty) {
            return Container();
          }
          if (chatRooms.length == 1) {
            return Text(chatRooms.first.name);
          }
          return DropdownButton<ChatRoom>(
            value: ref.watch(selectedChatRoomProvider),
            icon: const Padding(
              padding: EdgeInsets.only(left: 4.0),
              child: Icon(Icons.chat, size: 20),
            ),
            onChanged: (ChatRoom? newValue) {
              ref.watch(selectedChatRoomProvider.notifier).state = newValue;
              ref.read(chatIdProvider.notifier).state = newValue?.id ?? ref.read(chatIdProvider.notifier).state;
              final uri = Uri.parse(Uri.base.origin).replace(queryParameters: {'chat-id': newValue?.id.toString() ?? ""});
              window.history.pushState(null, "", uri.toString());
            },
            items: chatRooms.map<DropdownMenuItem<ChatRoom>>((ChatRoom chatRoom) {
              return DropdownMenuItem<ChatRoom>(
                value: chatRoom,
                child: Text(chatRoom.name),
              );
            }).toList(),
          );
        },
      );
    });
  }

  Widget settingsRow(BuildContext context) {
    return Row(
      children: [
        chatDropdownWidget(context),
      ],
    );
  }

  void showAddDaoDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (_) {
        var addressController = TextEditingController();
        return AlertDialog(
          title: const Text('Add DAO'),
          content: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Text("Enter the Juno address of the DAO you want to add"),
                const SizedBox(height: 15),
                TextFormField(
                  controller: addressController,
                  decoration: const InputDecoration(
                    labelText: 'Contract Address',
                    hintText: 'juno1z3zqgz7t0hcu2fx4wusuyjq0gc2m33la8l64saunfz7vmqwa2d5sz6jnep',
                  ),
                ),
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
                onPressed: () {
                  var address = addressController.text;
                  ref.read(chatroomListStateProvider.notifier).addDao(address);
                  Navigator.pop(context);
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
    return Center(
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(
          shape: const RoundedRectangleBorder(
            borderRadius: BorderRadius.all(Radius.circular(5)),
          ),
          minimumSize: ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? const Size(double.infinity, 40) : const Size(200, 40),
        ),
        onPressed: () {
          // ref.read(chatroomListStateProvider.notifier).addChatRoom();
          showAddDaoDialog(context);
        },
        child: const Text("Add DAO"),
      ),
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
              padding: const EdgeInsets.all(40),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text("Subscriptions", style: Theme.of(context).textTheme.headline2),
                  const SizedBox(height: 10),
                  chatDropdownWidget(context),
                  const SizedBox(height: 10),
                  const Text("Select the DAO's that you want to follow. You will receive notifications about new governance proposals."),
                  const SizedBox(height: 20),
                  searchWidget(context),
                  const SizedBox(height: 20),
                  subscriptionList(),
                  const SizedBox(height: 20),
                  addDaoButton(context),
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
