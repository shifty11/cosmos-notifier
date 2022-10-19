///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use subscriptionDescriptor instead')
const Subscription$json = const {
  '1': 'Subscription',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'is_subscribed', '3': 3, '4': 1, '5': 8, '10': 'isSubscribed'},
  ],
};

/// Descriptor for `Subscription`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List subscriptionDescriptor = $convert.base64Decode('CgxTdWJzY3JpcHRpb24SDgoCaWQYASABKANSAmlkEhIKBG5hbWUYAiABKAlSBG5hbWUSIwoNaXNfc3Vic2NyaWJlZBgDIAEoCFIMaXNTdWJzY3JpYmVk');
@$core.Deprecated('Use chatRoomDescriptor instead')
const ChatRoom$json = const {
  '1': 'ChatRoom',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'TYPE', '3': 3, '4': 1, '5': 14, '6': '.daodao_notifier_grpc.ChatRoom.Type', '10': 'TYPE'},
    const {'1': 'subscriptions', '3': 4, '4': 3, '5': 11, '6': '.daodao_notifier_grpc.Subscription', '10': 'subscriptions'},
  ],
  '4': const [ChatRoom_Type$json],
};

@$core.Deprecated('Use chatRoomDescriptor instead')
const ChatRoom_Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'TELEGRAM', '2': 0},
    const {'1': 'DISCORD', '2': 1},
  ],
};

/// Descriptor for `ChatRoom`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List chatRoomDescriptor = $convert.base64Decode('CghDaGF0Um9vbRIOCgJpZBgBIAEoA1ICaWQSEgoEbmFtZRgCIAEoCVIEbmFtZRI3CgRUWVBFGAMgASgOMiMuZGFvZGFvX25vdGlmaWVyX2dycGMuQ2hhdFJvb20uVHlwZVIEVFlQRRJICg1zdWJzY3JpcHRpb25zGAQgAygLMiIuZGFvZGFvX25vdGlmaWVyX2dycGMuU3Vic2NyaXB0aW9uUg1zdWJzY3JpcHRpb25zIiEKBFR5cGUSDAoIVEVMRUdSQU0QABILCgdESVNDT1JEEAE=');
@$core.Deprecated('Use getSubscriptionsResponseDescriptor instead')
const GetSubscriptionsResponse$json = const {
  '1': 'GetSubscriptionsResponse',
  '2': const [
    const {'1': 'chat_rooms', '3': 1, '4': 3, '5': 11, '6': '.daodao_notifier_grpc.ChatRoom', '10': 'chatRooms'},
  ],
};

/// Descriptor for `GetSubscriptionsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getSubscriptionsResponseDescriptor = $convert.base64Decode('ChhHZXRTdWJzY3JpcHRpb25zUmVzcG9uc2USPQoKY2hhdF9yb29tcxgBIAMoCzIeLmRhb2Rhb19ub3RpZmllcl9ncnBjLkNoYXRSb29tUgljaGF0Um9vbXM=');
@$core.Deprecated('Use toggleSubscriptionRequestDescriptor instead')
const ToggleSubscriptionRequest$json = const {
  '1': 'ToggleSubscriptionRequest',
  '2': const [
    const {'1': 'chatRoomId', '3': 1, '4': 1, '5': 3, '10': 'chatRoomId'},
    const {'1': 'contractId', '3': 2, '4': 1, '5': 3, '10': 'contractId'},
  ],
};

/// Descriptor for `ToggleSubscriptionRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toggleSubscriptionRequestDescriptor = $convert.base64Decode('ChlUb2dnbGVTdWJzY3JpcHRpb25SZXF1ZXN0Eh4KCmNoYXRSb29tSWQYASABKANSCmNoYXRSb29tSWQSHgoKY29udHJhY3RJZBgCIAEoA1IKY29udHJhY3RJZA==');
@$core.Deprecated('Use toggleSubscriptionResponseDescriptor instead')
const ToggleSubscriptionResponse$json = const {
  '1': 'ToggleSubscriptionResponse',
  '2': const [
    const {'1': 'isSubscribed', '3': 1, '4': 1, '5': 8, '10': 'isSubscribed'},
  ],
};

/// Descriptor for `ToggleSubscriptionResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toggleSubscriptionResponseDescriptor = $convert.base64Decode('ChpUb2dnbGVTdWJzY3JpcHRpb25SZXNwb25zZRIiCgxpc1N1YnNjcmliZWQYASABKAhSDGlzU3Vic2NyaWJlZA==');
