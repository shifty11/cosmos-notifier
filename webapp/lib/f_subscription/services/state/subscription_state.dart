import 'package:webapp/api/protobuf/dart/subscription_service.pb.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'subscription_state.freezed.dart';

@freezed
class SubscriptionState with _$SubscriptionState {
  const SubscriptionState._();

  factory SubscriptionState.loaded({required Subscription subscription}) = Loaded;
}
