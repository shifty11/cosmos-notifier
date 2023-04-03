import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:cosmos_notifier/api/protobuf/dart/subscription_service.pb.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/chat_id_provider.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_subscription/services/state/subscription_list_state.dart';
import 'package:cosmos_notifier/f_subscription/services/subscription_service.dart';
import 'package:cosmos_notifier/f_subscription/services/type/subscription_data_type.dart';
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
  List<ChatRoom> _chainChatRooms = [];
  List<ChatRoom> _contractChatRooms = [];

  late final SubscriptionService _subsService;

  SubscriptionListNotifier(this._ref) : super(SubscriptionListState.loading()) {
    _subsService = _ref.watch(subscriptionProvider);
    _loadSubscriptions();
  }

  Future<void> _loadSubscriptions() async {
    try {
      final response = await _subsService.listSubscriptions(Empty());
      _chainChatRooms = response.chainChatRooms;
      _contractChatRooms = response.contractChatRooms;
      state = SubscriptionListState.data(chainChatRooms: _chainChatRooms, contractChatRooms: _contractChatRooms);
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }

  _updateSubscription(Int64 chatRoomId, Int64 subscriptionId, ToggleSubscriptionResponse response) {
    ChatRoom chatRoom;
    if (_ref.read(isChainsSelectedProvider)) {
      chatRoom = _chainChatRooms.firstWhere((element) => element.id == chatRoomId);
    } else {
      chatRoom = _contractChatRooms.firstWhere((element) => element.id == chatRoomId);
    }
    final sub = chatRoom.subscriptions.firstWhere(((s) => s.id == subscriptionId)).deepCopy();
    sub.isSubscribed = response.isSubscribed;
    final index = chatRoom.subscriptions.indexWhere((element) => element.id == sub.id);
    chatRoom.subscriptions[index] = sub;

    state = SubscriptionListState.data(chainChatRooms: _chainChatRooms, contractChatRooms: _contractChatRooms);
  }

  _removeSubscription(Int64 id) {
    if (_ref.read(isChainsSelectedProvider)) {
      for (var chatRoom in _chainChatRooms) {
        chatRoom.subscriptions.removeWhere((element) => element.id == id);
      }
    } else {
      for (var chatRoom in _contractChatRooms) {
        chatRoom.subscriptions.removeWhere((element) => element.id == id);
      }
    }

    state = SubscriptionListState.data(chainChatRooms: _chainChatRooms, contractChatRooms: _contractChatRooms);
  }

  _setSubscriptionEnabled(Int64 id, bool isEnabled) {
    for (var chatRoom in _chainChatRooms) {
      chatRoom.subscriptions.firstWhere((element) => element.id == id).isEnabled = isEnabled;
    }
    state = SubscriptionListState.data(chainChatRooms: _chainChatRooms, contractChatRooms: _contractChatRooms);
  }

  Future<void> toggleSubscription(Int64 chatRoomId, Subscription subscription) async {
    try {
      ToggleSubscriptionResponse response;
      if (_ref.read(isChainsSelectedProvider)) {
        response =
            await _subsService.toggleChainSubscription(ToggleChainSubscriptionRequest(chatRoomId: chatRoomId, chainId: subscription.id));
      } else {
        response = await _subsService
            .toggleContractSubscription(ToggleContractSubscriptionRequest(chatRoomId: chatRoomId, contractId: subscription.id));
      }
      _updateSubscription(chatRoomId, subscription.id, response);
      if (response.isSubscribed) {
        _ref.read(messageProvider.notifier).sendMsg(info: 'Subscribed to ${subscription.name}');
      }
      if (jwtManager.isAdmin) {
        _loadSubscriptions(); // to update statistics
      }
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }

  Future<void> enableChain(Int64 chainId, String name, bool isEnabled) async {
    try {
      await _subsService.enableChain(EnableChainRequest(chainId: chainId, isEnabled: isEnabled));
      _setSubscriptionEnabled(chainId, isEnabled);
      _ref.read(messageProvider.notifier).sendMsg(info: "Chain $name ${isEnabled ? 'enabled' : 'disabled'}");
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }

  Future<void> deleteDao(Int64 contractId, String name) async {
    try {
      await _subsService.deleteDao(DeleteDaoRequest(contractId: contractId));
      _removeSubscription(contractId);
      _ref.read(messageProvider.notifier).sendMsg(info: "DAO $name deleted");
    } catch (e) {
      _ref.read(messageProvider.notifier).sendMsg(error: e.toString());
    }
  }

  Future<void> addDao(String contractAddress, String proposalQuery) async {
    final msgProvider = _ref.read(messageProvider.notifier);
    try {
      final stream = _subsService.addDao(AddDaoRequest(contractAddress: contractAddress, customQuery: proposalQuery));
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

final isChainsSelectedProvider = StateProvider((ref) => true);

final selectedChatRoomProvider = StateProvider<ChatRoom?>((ref) {
  final chatRooms = ref.watch(chatroomListStateProvider);
  final isChain = ref.watch(isChainsSelectedProvider);
  return chatRooms.maybeWhen(
    data: (chainChatRooms, contractChatRooms) {
      final chatId = ref.read(chatIdProvider) ?? Int64(0);
      if (isChain) {
        return chainChatRooms.firstWhere((element) => element.id == chatId, orElse: () => chainChatRooms.first);
      } else {
        return contractChatRooms.firstWhere((element) => element.id == chatId, orElse: () => contractChatRooms.first);
      }
    },
    orElse: () => null,
  );
});

final searchedSubsProvider = Provider<ChatroomData>((ref) {
  final search = ref.watch(searchSubsProvider);
  final chatRoom = ref.watch(selectedChatRoomProvider);
  final subs = ref.watch(chatroomListStateProvider);
  final isChain = ref.watch(isChainsSelectedProvider);
  return subs.whenOrNull(data: (chainChatRooms, contractChatRooms) {
        for (var cr in isChain ? chainChatRooms : contractChatRooms) {
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
