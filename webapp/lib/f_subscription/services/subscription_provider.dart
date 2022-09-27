import 'package:dao_dao_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:dao_dao_notifier/api/protobuf/dart/subscription_service.pb.dart';
import 'package:dao_dao_notifier/config.dart';
import 'package:dao_dao_notifier/f_home/services/chat_id_provider.dart';
import 'package:dao_dao_notifier/f_home/services/message_provider.dart';
import 'package:dao_dao_notifier/f_subscription/services/state/subscription_list_state.dart';
import 'package:dao_dao_notifier/f_subscription/services/subscription_service.dart';
import 'package:dao_dao_notifier/f_subscription/services/type/subscription_data_type.dart';
import 'package:fixnum/fixnum.dart' as fixnum;
import 'package:fixnum/fixnum.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:protobuf/protobuf.dart';

final subscriptionProvider = Provider<SubscriptionService>((ref) => subsService);

final chatroomListStateProvider = StateNotifierProvider<SubscriptionListNotifier, SubscriptionListState>((ref) {
  return SubscriptionListNotifier(ref);
});

class SubscriptionListNotifier extends StateNotifier<SubscriptionListState> {
  final StateNotifierProviderRef _ref;
  List<ChatRoom> _chatRooms = [];

  late final SubscriptionService _subsService;

  SubscriptionListNotifier(this._ref) : super(SubscriptionListState.loading()) {
    _subsService = _ref.read(subscriptionProvider);
    _loadSubscriptions();
  }

  Future<void> _loadSubscriptions() async {
    try {
      final response = await _subsService.getSubscriptions(Empty());
      _chatRooms = response.chatRooms;
      state = SubscriptionListState.data(chatRooms: _chatRooms);
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }

  _updateSubscription(Int64 chatRoomId, Subscription subscription) {
    final chatRoom = _chatRooms.firstWhere((element) => element.id == chatRoomId);
    final index = chatRoom.subscriptions.indexWhere((element) => element.id == subscription.id);
    chatRoom.subscriptions[index] = subscription;

    state = SubscriptionListState.data(chatRooms: _chatRooms);
  }

  _removeSubscriptions(Int64 id) {
    for (var chatRoom in _chatRooms) {
      chatRoom.subscriptions.removeWhere((element) => element.id == id);
    }

    state = SubscriptionListState.data(chatRooms: _chatRooms);
  }

  Future<void> toggleSubscription(chatRoomId, Int64 subscriptionId) async {
    try {
      final response = await _subsService.toggleSubscription(ToggleSubscriptionRequest(chatRoomId: chatRoomId, contractId: subscriptionId));
      final sub = _chatRooms.firstWhere((c) => c.id == chatRoomId).subscriptions.firstWhere(((s) => s.id == subscriptionId)).deepCopy();
      sub.isSubscribed = response.isSubscribed;
      _updateSubscription(chatRoomId, sub);
      if (jwtManager.isAdmin) {
        _loadSubscriptions(); // to update statistics
      }
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }

  Future<void> deleteDao(Int64 contractId, String name) async {
    try {
      await _subsService.deleteDao(DeleteDaoRequest(contractId: contractId));
      _removeSubscriptions(contractId);
      _ref.read(messageProvider.notifier).sendMsg(info: "DAO $name deleted");
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }

  Future<void> addDao(String contractAddress) async {
    final msgProvider = _ref.read(messageProvider.notifier);
    try {
      final stream = _subsService.addDao(AddDaoRequest(contractAddress: contractAddress));
      stream.listen(
        (resp) {
          switch (resp.status) {
            case AddDaoResponse_Status.ADDED:
              msgProvider.sendMsg(info: "DAO ${resp.name} added 🚀");
              _loadSubscriptions();
              break;
            case AddDaoResponse_Status.ALREADY_ADDED:
              msgProvider.sendMsg(info: "DAO has already been added");
              break;
            case AddDaoResponse_Status.IS_ADDING:
              msgProvider.sendMsg(info: "DAO is being added... this can take a minute 😴");
              break;
            case AddDaoResponse_Status.FAILED:
              msgProvider.sendMsg(error: "Could not add DAO! 😭\nProbably because the contract address is invalid");
              break;
          }
        },
        onError: (e) => msgProvider.sendMsg(error: e.toString()),
        onDone: () {},
        cancelOnError: true,
      );
    } catch (e) {
      msgProvider.sendMsg(error: e.toString());
    }
  }
}

final searchSubsProvider = StateProvider((ref) => "");

final selectedChatRoomProvider = StateProvider<ChatRoom?>((ref) {
  final chatRooms = ref.watch(chatroomListStateProvider);
  return chatRooms.maybeWhen(
    data: (chatRooms) {
      final chatId = ref.read(chatIdProvider) ?? Int64(0);
      return chatRooms.firstWhere((element) => element.id == chatId, orElse: () => chatRooms.first);
    },
    orElse: () => null,
  );
});

final searchedSubsProvider = Provider<ChatroomData>((ref) {
  final search = ref.watch(searchSubsProvider);
  final chatRoom = ref.watch(selectedChatRoomProvider);
  final subs = ref.watch(chatroomListStateProvider);
  return subs.whenOrNull(data: (chatRooms) {
        for (var cr in chatRooms) {
          if (cr.id == chatRoom?.id || chatRoom == null) {
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
                    .where((e) =>
                        e.value.name.toLowerCase().contains(search.toLowerCase()) ||
                        e.value.contractAddress.toLowerCase().contains(search.toLowerCase()))
                    .map((e) => SubscriptionData(e.value, e.key))
                    .toList());
          }
        }
        return ChatroomData(fixnum.Int64(), "", [], []);
      }) ??
      ChatroomData(fixnum.Int64(), "", [], []);
});
