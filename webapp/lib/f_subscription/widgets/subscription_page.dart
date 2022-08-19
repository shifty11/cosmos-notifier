import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:responsive_framework/responsive_framework.dart';
import 'package:riverpod_messages/riverpod_messages.dart';
import 'package:tuple/tuple.dart';
import 'package:webapp/api/protobuf/dart/subscription_service.pb.dart';
import 'package:webapp/config.dart';
import 'package:webapp/f_home/services/message_provider.dart';
import 'package:webapp/f_home/widgets/bottom_navigation_bar_widget.dart';
import 'package:webapp/f_subscription/services/subscription_provider.dart';
import 'package:webapp/f_subscription/services/type/subscription_data_type.dart';
import 'package:webapp/style.dart';

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

  Widget subscriptionsLoaded(BuildContext context, ChatroomData chatRoom) {
    return GridView.builder(
      shrinkWrap: true,
      gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: getCrossAxisCount(context),
        crossAxisSpacing: 10,
        mainAxisSpacing: 10,
        mainAxisExtent: 60,
      ),
      itemCount: chatRoom.filtered.length,
      itemBuilder: (BuildContext context, int index) {
        return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
          final data = Tuple2(chatRoom.chatRoomId, chatRoom.getUnfilteredIndex(index));
          final state = ref.watch(subscriptionStateProvider(data));
          const double sidePadding = 12;
          return state.when(
            loaded: (subscription) => Container(
              decoration: BoxDecoration(
                  border: Border.all(
                    width: Styles.selectCardBorderWidth,
                    color: Theme.of(context).inputDecorationTheme.enabledBorder!.borderSide.color,
                  ),
                  borderRadius: const BorderRadius.all(Radius.circular(5))),
              child: InkWell(
                onTap: () {
                  FocusScopeNode currentFocus = FocusScope.of(context);
                  if (!currentFocus.hasPrimaryFocus) {
                    currentFocus.focusedChild?.unfocus();
                  }
                  ref.read(subscriptionStateProvider(data).notifier).toggleSubscription();
                },
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    const SizedBox(width: sidePadding),
                    CircleAvatar(
                      child: ClipOval(
                        child: Image.asset(
                          subscription.thumbnailUrl,
                          width: 40,
                          height: 40,
                        ),
                      ),
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
                            style: TextStyle(fontSize: 12, color: Theme.of(context).inputDecorationTheme.iconColor),
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
        });
      },
    );
  }

  Widget subscriptionList() {
    return Expanded(
      child: Consumer(
        builder: (BuildContext context, WidgetRef ref, Widget? child) {
          final state = ref.watch(chatroomListStateProvider);
          return state.when(
            loading: () => const Center(child: CircularProgressIndicator()),
            data: (chatRooms) => subscriptionsLoaded(context, ref.watch(searchedSubsProvider)),
            error: (err, stackTrace) => Container(),
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
            value: ref.watch(chatRoomProvider),
            icon: const Padding(
              padding: EdgeInsets.only(left: 4.0),
              child: Icon(Icons.chat, size: 20),
            ),
            onChanged: (ChatRoom? newValue) {
              ref.watch(chatRoomProvider.notifier).state = newValue;
            },
            items: chatRooms.map<DropdownMenuItem<ChatRoom>>((ChatRoom chatRoom) {
              return DropdownMenuItem<ChatRoom>(
                value: chatRoom,
                child: Text(chatRoom.name),
              );
            }).toList(),
          );
        },
        error: (err, stackTrace) => Container(),
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
                  const Text(
                      "Select the DAO's that you want to follow. You will receive notifications about new governance proposals."),
                  const SizedBox(height: 20),
                  searchWidget(context),
                  const SizedBox(height: 20),
                  subscriptionList(),
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
