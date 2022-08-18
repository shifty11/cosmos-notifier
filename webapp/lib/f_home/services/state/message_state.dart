import 'package:freezed_annotation/freezed_annotation.dart';

part 'message_state.freezed.dart';

@freezed
class MessageState with _$MessageState {
  const MessageState._();

  const factory MessageState.initial() = Initial;

  factory MessageState.received({String? info, String? error}) = Received;
}
