import 'package:cosmos_notifier/api/protobuf/dart/subscription_service.pb.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_subscription/services/subscription_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import '../../f_home/services/chat_id_provider.dart';

class HeaderWidget extends StatelessWidget {
  bool showChatDropdownWidget;

  HeaderWidget(this.showChatDropdownWidget, {Key? key}) : super(key: key);

  Widget chatDropdownWidget(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final state = ref.watch(chatroomListStateProvider);
      return state.when(
        loading: () => Container(),
        data: (chainChatRooms, contractChatRooms) {
          var chatRooms = ref.read(isChainsSelectedProvider) ? chainChatRooms : contractChatRooms;
          if (chatRooms.isEmpty) {
            return Container();
          }
          if (chatRooms.length == 1) {
            return Text(chatRooms.first.name);
          }
          return DropdownButtonHideUnderline(
            child: DropdownButton<ChatRoom>(
              value: ref.watch(selectedChatRoomProvider),
              icon: const Padding(
                padding: EdgeInsets.only(left: 4.0),
                child: Icon(Icons.person, size: 20),
              ),
              onChanged: (ChatRoom? newValue) {
                ref.watch(selectedChatRoomProvider.notifier).state = newValue;
                ref.read(chatIdProvider.notifier).state = newValue?.id ?? ref.read(chatIdProvider.notifier).state;
                context.pushNamed(rSubscriptions.name, queryParams: {'chat-id': newValue?.id.toString() ?? ""});
              },
              items: chatRooms.map<DropdownMenuItem<ChatRoom>>((ChatRoom chatRoom) {
                return DropdownMenuItem<ChatRoom>(
                  value: chatRoom,
                  child: Text(chatRoom.name),
                );
              }).toList(),
            ),
          );
        },
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      var location = GoRouter.of(context).location;
      return Column(
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              TextButton.icon(
                onPressed: () => context.pushNamed(rRoot.name),
                icon: const Icon(Icons.home),
                label: const Text("Home"),
                style: TextButton.styleFrom(
                  foregroundColor: location == rRoot.path ? Theme.of(context).primaryColor : Theme.of(context).disabledColor,
                ),
              ),
              TextButton.icon(
                onPressed: () => context.pushNamed(rSubscriptions.name),
                icon: const Icon(Icons.notifications),
                label: const Text("Subscriptions"),
                style: TextButton.styleFrom(
                  foregroundColor: location == rSubscriptions.path ? Theme.of(context).primaryColor : Theme.of(context).disabledColor,
                ),
              ),
              TextButton.icon(
                onPressed: () => context.pushNamed(rTracking.name),
                icon: const Icon(Icons.my_location),
                label: const Text("Tracking"),
                style: TextButton.styleFrom(
                  foregroundColor: location == rTracking.path ? Theme.of(context).primaryColor : Theme.of(context).disabledColor,
                ),
              ),
              jwtManager.isAdmin ? TextButton.icon(
                onPressed: () => context.pushNamed(rAdmin.name),
                icon: const Icon(Icons.settings),
                label: const Text("Admin"),
                style: TextButton.styleFrom(
                  foregroundColor: location == rAdmin.path ? Theme.of(context).primaryColor : Theme.of(context).disabledColor,
                ),
              ) : Container(),
              const Spacer(),
              showChatDropdownWidget ? chatDropdownWidget(context) : Container(),
            ],
          ),
          const Divider(),
        ],
      );
    });
  }
}
