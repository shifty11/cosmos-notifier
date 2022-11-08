import 'package:cosmos_notifier/api/protobuf/dart/admin_service.pbgrpc.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_admin/widget/services/admin_service.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final adminProvider = StateNotifierProvider<AdminStateNotifier, void>((ref) {
  return AdminStateNotifier(ref.read(messageProvider.notifier), adminService);
});

class AdminStateNotifier extends StateNotifier<void> {
  final MessageNotifier _messageNotifier;
  final AdminService _adminService;

  AdminStateNotifier(this._messageNotifier, this._adminService) : super(null);

  Future<void> broadcastMessage(String message, BroadcastMessageRequest_MessageType type) async {
    try {
      final stream = _adminService.broadcastMessage(BroadcastMessageRequest(message: message, type: type));
      stream.listen(
        (resp) {
          switch (resp.status) {
            case BroadcastMessageResponse_Status.SENDING:
              _messageNotifier.sendMsg(info: resp.response);
              break;
            case BroadcastMessageResponse_Status.SENT:
              _messageNotifier.sendMsg(info: resp.response);
              break;
            case BroadcastMessageResponse_Status.FAILED:
              _messageNotifier.sendMsg(error: "Error broadcasting message to $type chats");
              break;
          }
        },
        onError: (e) => _messageNotifier.sendMsg(error: e.toString()),
        onDone: () {},
        cancelOnError: true,
      );
    } catch (e) {
      _messageNotifier.sendMsg(error: e.toString());
    }
  }
}
