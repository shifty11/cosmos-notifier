import 'package:webapp/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:webapp/api/protobuf/dart/subscription_service.pb.dart';
import 'package:webapp/config.dart';
import 'package:webapp/f_home/services/message_provider.dart';
import 'package:webapp/f_subscription/services/state/subscription_state.dart';
import 'package:webapp/f_subscription/services/subscription_service.dart';
import 'package:webapp/f_subscription/services/type/subscription_data_type.dart';
import 'package:fixnum/fixnum.dart' as fixnum;
import 'package:fixnum/fixnum.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tuple/tuple.dart';

final subscriptionProvider = Provider<SubscriptionService>((ref) => subsService);

final chatroomListStateProvider = FutureProvider<List<ChatRoom>>((ref) async {
  final subsService = ref.read(subscriptionProvider);
  final response = await subsService.getSubscriptions(Empty());

  final selectedChatRoom = ref.read(chatRoomProvider.notifier).state;
  if (selectedChatRoom != null) { // if chat room was selected before, select it again
    ref.read(chatRoomProvider.notifier).state = response.chatRooms.where((c) => c.id == selectedChatRoom.id).first;
  } else {
    if (response.chatRooms.length > 1) {
      // if query params contain `chat_id` put that one first
      final chatIdStr = Uri.base.queryParameters['chat_id'];
      var chatId = Int64();
      if (chatIdStr != null) {
        try {
          chatId = Int64.parseInt(chatIdStr);
        } on FormatException {
          // ignore exceptions since the query param could be anything
        }
      }
      response.chatRooms.sort(((a, b) {
        if (a.id == chatId) {
          return -1;
        }
        if (b.id == chatId) {
          return 1;
        }
        return a.name.compareTo(b.name);
      }));
    }
    ref.read(chatRoomProvider.notifier).state = response.chatRooms.first;
  }
  return response.chatRooms;
});

final subscriptionStateProvider = StateNotifierProvider.family<SubscriptionNotifier, SubscriptionState, Tuple2<fixnum.Int64, int>>(
  (ref, tuple) {
    final chatRoom = ref.watch(chatroomListStateProvider).value!.firstWhere((c) => c.id == tuple.item1);
    final subscription = chatRoom.subscriptions[tuple.item2];
    return SubscriptionNotifier(ref, subscription, chatRoom.id);
  },
);

class SubscriptionNotifier extends StateNotifier<SubscriptionState> {
  final Subscription _subscription;
  final fixnum.Int64 _chatRoomId;
  final StateNotifierProviderRef _ref;

  SubscriptionNotifier(this._ref, this._subscription, this._chatRoomId) : super(SubscriptionState.loaded(subscription: _subscription));

  Future<void> toggleSubscription() async {
    try {
      final subsService = _ref.read(subscriptionProvider);
      final response = await subsService.toggleSubscription(ToggleSubscriptionRequest(chatRoomId: _chatRoomId, contractId: _subscription.id));
      _subscription.isSubscribed = response.isSubscribed;
      state = SubscriptionState.loaded(subscription: _subscription);
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }
}

final searchSubsProvider = StateProvider((ref) => "");

final chatRoomProvider = StateProvider<ChatRoom?>((ref) => null);

final searchedSubsProvider = Provider<ChatroomData>((ref) {
  final search = ref.watch(searchSubsProvider);
  final chatRoom = ref.watch(chatRoomProvider);
  final subs = ref.watch(chatroomListStateProvider);
  return subs.whenOrNull(data: (chatRooms) {
        for (var cr in chatRooms) {
          if (cr == chatRoom || chatRoom == null) {
            if (search.isEmpty) {
              return ChatroomData(
                cr.id,
                cr.name,
                cr.subscriptions,
                cr.subscriptions.asMap().entries.map((e) => SubscriptionData(e.value, e.key)).toList(),
              );
            }
            return ChatroomData(
                cr.id,
                cr.name,
                cr.subscriptions,
                cr.subscriptions
                    .asMap()
                    .entries
                    .where((e) => e.value.name.toLowerCase().contains(search.toLowerCase()))
                    .map((e) => SubscriptionData(e.value, e.key))
                    .toList());
          }
        }
        return ChatroomData(fixnum.Int64(), "", [], []);
      }) ??
      ChatroomData(fixnum.Int64(), "", [], []);
});

final wantsPreVotePropsStateProvider = StateNotifierProvider<WantsPreVotePropsNotifier, bool>((ref) => WantsPreVotePropsNotifier(ref));

class WantsPreVotePropsNotifier extends StateNotifier<bool> {
  final StateNotifierProviderRef _ref;

  WantsPreVotePropsNotifier(this._ref) : super(false);

  Future<void> toggleWantsPreVoteProps() async {
    try {
      final subsService = _ref.read(subscriptionProvider);
      final chatRoom = _ref.read(chatRoomProvider);
      if (chatRoom == null) {
        return;
      }
      // final response = await subsService.updateSettings(UpdateSettingsRequest(chatRoomId: chatRoom.id, wantsDraftProposals: !chatRoom.wantsDraftProposals));
      // chatRoom.wantsDraftProposals = response.wantsDraftProposals;
      // state = response.wantsDraftProposals;
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }
}

final wantsPreVotePropsProvider = Provider<bool>((ref) {
  final chatRoom = ref.watch(chatRoomProvider);
  ref.watch(wantsPreVotePropsStateProvider);
  return false;
});
