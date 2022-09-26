import 'package:dao_dao_notifier/api/protobuf/dart/subscription_service.pbgrpc.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:fixnum/fixnum.dart' as fixnum;

part 'subscription_data_type.freezed.dart';

@freezed
@immutable
class SubscriptionData with _$SubscriptionData {
  const factory SubscriptionData(Subscription subscription, int index) = _SubscriptionData;
}

@freezed
class ChatroomData with _$ChatroomData {
  const ChatroomData._();
  const factory ChatroomData(fixnum.Int64 chatRoomId, String username, List<Subscription> subscriptions, List<SubscriptionData> filtered) = _ChatroomData;

  int getUnfilteredIndex(int index) {
    return filtered[index].index;
  }
}
