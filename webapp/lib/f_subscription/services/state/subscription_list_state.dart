import 'package:webapp/api/protobuf/dart/subscription_service.pb.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'subscription_list_state.freezed.dart';

@freezed
class SubscriptionListState with _$SubscriptionListState {
  const SubscriptionListState._();

  factory SubscriptionListState.loading() = Loading;
  factory SubscriptionListState.data({required List<ChatRoom> chatRooms}) = Data;
}