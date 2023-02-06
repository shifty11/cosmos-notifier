///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use subscriptionStatsDescriptor instead')
const SubscriptionStats$json = const {
  '1': 'SubscriptionStats',
  '2': const [
    const {'1': 'total', '3': 1, '4': 1, '5': 5, '10': 'total'},
    const {'1': 'telegram', '3': 2, '4': 1, '5': 5, '10': 'telegram'},
    const {'1': 'discord', '3': 3, '4': 1, '5': 5, '10': 'discord'},
  ],
};

/// Descriptor for `SubscriptionStats`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List subscriptionStatsDescriptor = $convert.base64Decode('ChFTdWJzY3JpcHRpb25TdGF0cxIUCgV0b3RhbBgBIAEoBVIFdG90YWwSGgoIdGVsZWdyYW0YAiABKAVSCHRlbGVncmFtEhgKB2Rpc2NvcmQYAyABKAVSB2Rpc2NvcmQ=');
@$core.Deprecated('Use subscriptionDescriptor instead')
const Subscription$json = const {
  '1': 'Subscription',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'is_subscribed', '3': 3, '4': 1, '5': 8, '10': 'isSubscribed'},
    const {'1': 'is_enabled', '3': 4, '4': 1, '5': 8, '10': 'isEnabled'},
    const {'1': 'thumbnail_url', '3': 5, '4': 1, '5': 9, '10': 'thumbnailUrl'},
    const {'1': 'contract_address', '3': 6, '4': 1, '5': 9, '10': 'contractAddress'},
    const {'1': 'stats', '3': 7, '4': 1, '5': 11, '6': '.cosmos_notifier_grpc.SubscriptionStats', '10': 'stats'},
  ],
};

/// Descriptor for `Subscription`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List subscriptionDescriptor = $convert.base64Decode('CgxTdWJzY3JpcHRpb24SDgoCaWQYASABKANSAmlkEhIKBG5hbWUYAiABKAlSBG5hbWUSIwoNaXNfc3Vic2NyaWJlZBgDIAEoCFIMaXNTdWJzY3JpYmVkEh0KCmlzX2VuYWJsZWQYBCABKAhSCWlzRW5hYmxlZBIjCg10aHVtYm5haWxfdXJsGAUgASgJUgx0aHVtYm5haWxVcmwSKQoQY29udHJhY3RfYWRkcmVzcxgGIAEoCVIPY29udHJhY3RBZGRyZXNzEj0KBXN0YXRzGAcgASgLMicuY29zbW9zX25vdGlmaWVyX2dycGMuU3Vic2NyaXB0aW9uU3RhdHNSBXN0YXRz');
@$core.Deprecated('Use chatRoomDescriptor instead')
const ChatRoom$json = const {
  '1': 'ChatRoom',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'TYPE', '3': 3, '4': 1, '5': 14, '6': '.cosmos_notifier_grpc.ChatRoom.Type', '10': 'TYPE'},
    const {'1': 'subscriptions', '3': 4, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.Subscription', '10': 'subscriptions'},
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
final $typed_data.Uint8List chatRoomDescriptor = $convert.base64Decode('CghDaGF0Um9vbRIOCgJpZBgBIAEoA1ICaWQSEgoEbmFtZRgCIAEoCVIEbmFtZRI3CgRUWVBFGAMgASgOMiMuY29zbW9zX25vdGlmaWVyX2dycGMuQ2hhdFJvb20uVHlwZVIEVFlQRRJICg1zdWJzY3JpcHRpb25zGAQgAygLMiIuY29zbW9zX25vdGlmaWVyX2dycGMuU3Vic2NyaXB0aW9uUg1zdWJzY3JpcHRpb25zIiEKBFR5cGUSDAoIVEVMRUdSQU0QABILCgdESVNDT1JEEAE=');
@$core.Deprecated('Use getSubscriptionsResponseDescriptor instead')
const GetSubscriptionsResponse$json = const {
  '1': 'GetSubscriptionsResponse',
  '2': const [
    const {'1': 'chain_chat_rooms', '3': 1, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.ChatRoom', '10': 'chainChatRooms'},
    const {'1': 'contract_chat_rooms', '3': 2, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.ChatRoom', '10': 'contractChatRooms'},
  ],
};

/// Descriptor for `GetSubscriptionsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getSubscriptionsResponseDescriptor = $convert.base64Decode('ChhHZXRTdWJzY3JpcHRpb25zUmVzcG9uc2USSAoQY2hhaW5fY2hhdF9yb29tcxgBIAMoCzIeLmNvc21vc19ub3RpZmllcl9ncnBjLkNoYXRSb29tUg5jaGFpbkNoYXRSb29tcxJOChNjb250cmFjdF9jaGF0X3Jvb21zGAIgAygLMh4uY29zbW9zX25vdGlmaWVyX2dycGMuQ2hhdFJvb21SEWNvbnRyYWN0Q2hhdFJvb21z');
@$core.Deprecated('Use toggleChainSubscriptionRequestDescriptor instead')
const ToggleChainSubscriptionRequest$json = const {
  '1': 'ToggleChainSubscriptionRequest',
  '2': const [
    const {'1': 'chatRoomId', '3': 1, '4': 1, '5': 3, '10': 'chatRoomId'},
    const {'1': 'chainId', '3': 2, '4': 1, '5': 3, '10': 'chainId'},
  ],
};

/// Descriptor for `ToggleChainSubscriptionRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toggleChainSubscriptionRequestDescriptor = $convert.base64Decode('Ch5Ub2dnbGVDaGFpblN1YnNjcmlwdGlvblJlcXVlc3QSHgoKY2hhdFJvb21JZBgBIAEoA1IKY2hhdFJvb21JZBIYCgdjaGFpbklkGAIgASgDUgdjaGFpbklk');
@$core.Deprecated('Use toggleContractSubscriptionRequestDescriptor instead')
const ToggleContractSubscriptionRequest$json = const {
  '1': 'ToggleContractSubscriptionRequest',
  '2': const [
    const {'1': 'chatRoomId', '3': 1, '4': 1, '5': 3, '10': 'chatRoomId'},
    const {'1': 'contractId', '3': 2, '4': 1, '5': 3, '10': 'contractId'},
  ],
};

/// Descriptor for `ToggleContractSubscriptionRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toggleContractSubscriptionRequestDescriptor = $convert.base64Decode('CiFUb2dnbGVDb250cmFjdFN1YnNjcmlwdGlvblJlcXVlc3QSHgoKY2hhdFJvb21JZBgBIAEoA1IKY2hhdFJvb21JZBIeCgpjb250cmFjdElkGAIgASgDUgpjb250cmFjdElk');
@$core.Deprecated('Use toggleSubscriptionResponseDescriptor instead')
const ToggleSubscriptionResponse$json = const {
  '1': 'ToggleSubscriptionResponse',
  '2': const [
    const {'1': 'isSubscribed', '3': 1, '4': 1, '5': 8, '10': 'isSubscribed'},
  ],
};

/// Descriptor for `ToggleSubscriptionResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toggleSubscriptionResponseDescriptor = $convert.base64Decode('ChpUb2dnbGVTdWJzY3JpcHRpb25SZXNwb25zZRIiCgxpc1N1YnNjcmliZWQYASABKAhSDGlzU3Vic2NyaWJlZA==');
@$core.Deprecated('Use addDaoRequestDescriptor instead')
const AddDaoRequest$json = const {
  '1': 'AddDaoRequest',
  '2': const [
    const {'1': 'contractAddress', '3': 1, '4': 1, '5': 9, '10': 'contractAddress'},
    const {'1': 'customQuery', '3': 2, '4': 1, '5': 9, '10': 'customQuery'},
  ],
};

/// Descriptor for `AddDaoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addDaoRequestDescriptor = $convert.base64Decode('Cg1BZGREYW9SZXF1ZXN0EigKD2NvbnRyYWN0QWRkcmVzcxgBIAEoCVIPY29udHJhY3RBZGRyZXNzEiAKC2N1c3RvbVF1ZXJ5GAIgASgJUgtjdXN0b21RdWVyeQ==');
@$core.Deprecated('Use addDaoResponseDescriptor instead')
const AddDaoResponse$json = const {
  '1': 'AddDaoResponse',
  '2': const [
    const {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.cosmos_notifier_grpc.AddDaoResponse.Status', '10': 'status'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'contractAddress', '3': 3, '4': 1, '5': 9, '10': 'contractAddress'},
  ],
  '4': const [AddDaoResponse_Status$json],
};

@$core.Deprecated('Use addDaoResponseDescriptor instead')
const AddDaoResponse_Status$json = const {
  '1': 'Status',
  '2': const [
    const {'1': 'ADDED', '2': 0},
    const {'1': 'ALREADY_ADDED', '2': 1},
    const {'1': 'IS_ADDING', '2': 2},
    const {'1': 'FAILED', '2': 3},
  ],
};

/// Descriptor for `AddDaoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addDaoResponseDescriptor = $convert.base64Decode('Cg5BZGREYW9SZXNwb25zZRJDCgZzdGF0dXMYASABKA4yKy5jb3Ntb3Nfbm90aWZpZXJfZ3JwYy5BZGREYW9SZXNwb25zZS5TdGF0dXNSBnN0YXR1cxISCgRuYW1lGAIgASgJUgRuYW1lEigKD2NvbnRyYWN0QWRkcmVzcxgDIAEoCVIPY29udHJhY3RBZGRyZXNzIkEKBlN0YXR1cxIJCgVBRERFRBAAEhEKDUFMUkVBRFlfQURERUQQARINCglJU19BRERJTkcQAhIKCgZGQUlMRUQQAw==');
@$core.Deprecated('Use deleteDaoRequestDescriptor instead')
const DeleteDaoRequest$json = const {
  '1': 'DeleteDaoRequest',
  '2': const [
    const {'1': 'contractId', '3': 1, '4': 1, '5': 3, '10': 'contractId'},
  ],
};

/// Descriptor for `DeleteDaoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteDaoRequestDescriptor = $convert.base64Decode('ChBEZWxldGVEYW9SZXF1ZXN0Eh4KCmNvbnRyYWN0SWQYASABKANSCmNvbnRyYWN0SWQ=');
@$core.Deprecated('Use enableChainRequestDescriptor instead')
const EnableChainRequest$json = const {
  '1': 'EnableChainRequest',
  '2': const [
    const {'1': 'chainId', '3': 1, '4': 1, '5': 3, '10': 'chainId'},
    const {'1': 'isEnabled', '3': 2, '4': 1, '5': 8, '10': 'isEnabled'},
  ],
};

/// Descriptor for `EnableChainRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List enableChainRequestDescriptor = $convert.base64Decode('ChJFbmFibGVDaGFpblJlcXVlc3QSGAoHY2hhaW5JZBgBIAEoA1IHY2hhaW5JZBIcCglpc0VuYWJsZWQYAiABKAhSCWlzRW5hYmxlZA==');
