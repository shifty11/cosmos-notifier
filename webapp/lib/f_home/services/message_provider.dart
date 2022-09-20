import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:webapp/config.dart';
import 'package:webapp/f_home/services/jwt_manager.dart';
import 'package:webapp/f_home/services/state/message_state.dart';

final messageProvider = StateNotifierProvider<MessageNotifier, MessageState>((ref) => MessageNotifier(ref, jwtManager));

class MessageNotifier extends MessageNotifierBase {
  MessageNotifier(StateNotifierProviderRef<MessageNotifierBase, MessageState> ref, JwtManager jwtManager) : super(ref, jwtManager) {
    if (jwtManager.isAdmin) {}
  }
}

class MessageNotifierBase extends StateNotifier<MessageState> {
  final StateNotifierProviderRef<MessageNotifierBase, MessageState> ref;
  final JwtManager jwtManager;

  MessageNotifierBase(this.ref, this.jwtManager) : super(const MessageState.initial());

  sendMsg({String? info, String? error}) async {
    state = MessageState.received(info: info, error: error);
    await Future.delayed(const Duration(seconds: 1));
    state = const MessageState.initial();
  }
}
