///
//  Generated code. Do not modify.
//  source: admin_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use broadcastMessageRequestDescriptor instead')
const BroadcastMessageRequest$json = const {
  '1': 'BroadcastMessageRequest',
  '2': const [
    const {'1': 'message', '3': 1, '4': 1, '5': 9, '10': 'message'},
    const {'1': 'type', '3': 2, '4': 1, '5': 14, '6': '.cosmos_notifier_grpc.BroadcastMessageRequest.MessageType', '10': 'type'},
  ],
  '4': const [BroadcastMessageRequest_MessageType$json],
};

@$core.Deprecated('Use broadcastMessageRequestDescriptor instead')
const BroadcastMessageRequest_MessageType$json = const {
  '1': 'MessageType',
  '2': const [
    const {'1': 'MESSAGE_TYPE_UNSPECIFIED', '2': 0},
    const {'1': 'TELEGRAM_TEST', '2': 1},
    const {'1': 'DISCORD_TEST', '2': 2},
    const {'1': 'TELEGRAM', '2': 3},
    const {'1': 'DISCORD', '2': 4},
  ],
};

/// Descriptor for `BroadcastMessageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List broadcastMessageRequestDescriptor = $convert.base64Decode('ChdCcm9hZGNhc3RNZXNzYWdlUmVxdWVzdBIYCgdtZXNzYWdlGAEgASgJUgdtZXNzYWdlEk0KBHR5cGUYAiABKA4yOS5jb3Ntb3Nfbm90aWZpZXJfZ3JwYy5Ccm9hZGNhc3RNZXNzYWdlUmVxdWVzdC5NZXNzYWdlVHlwZVIEdHlwZSJrCgtNZXNzYWdlVHlwZRIcChhNRVNTQUdFX1RZUEVfVU5TUEVDSUZJRUQQABIRCg1URUxFR1JBTV9URVNUEAESEAoMRElTQ09SRF9URVNUEAISDAoIVEVMRUdSQU0QAxILCgdESVNDT1JEEAQ=');
@$core.Deprecated('Use broadcastMessageResponseDescriptor instead')
const BroadcastMessageResponse$json = const {
  '1': 'BroadcastMessageResponse',
  '2': const [
    const {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.cosmos_notifier_grpc.BroadcastMessageResponse.Status', '10': 'status'},
    const {'1': 'response', '3': 2, '4': 1, '5': 9, '10': 'response'},
  ],
  '4': const [BroadcastMessageResponse_Status$json],
};

@$core.Deprecated('Use broadcastMessageResponseDescriptor instead')
const BroadcastMessageResponse_Status$json = const {
  '1': 'Status',
  '2': const [
    const {'1': 'STATUS_UNSPECIFIED', '2': 0},
    const {'1': 'SENDING', '2': 1},
    const {'1': 'SENT', '2': 2},
    const {'1': 'FAILED', '2': 3},
  ],
};

/// Descriptor for `BroadcastMessageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List broadcastMessageResponseDescriptor = $convert.base64Decode('ChhCcm9hZGNhc3RNZXNzYWdlUmVzcG9uc2USTQoGc3RhdHVzGAEgASgOMjUuY29zbW9zX25vdGlmaWVyX2dycGMuQnJvYWRjYXN0TWVzc2FnZVJlc3BvbnNlLlN0YXR1c1IGc3RhdHVzEhoKCHJlc3BvbnNlGAIgASgJUghyZXNwb25zZSJDCgZTdGF0dXMSFgoSU1RBVFVTX1VOU1BFQ0lGSUVEEAASCwoHU0VORElORxABEggKBFNFTlQQAhIKCgZGQUlMRUQQAw==');
@$core.Deprecated('Use getStatsResponseDescriptor instead')
const GetStatsResponse$json = const {
  '1': 'GetStatsResponse',
  '2': const [
    const {'1': 'chains', '3': 1, '4': 1, '5': 5, '10': 'chains'},
    const {'1': 'contracts', '3': 2, '4': 1, '5': 5, '10': 'contracts'},
    const {'1': 'users', '3': 3, '4': 1, '5': 5, '10': 'users'},
    const {'1': 'telegram_users', '3': 4, '4': 1, '5': 5, '10': 'telegramUsers'},
    const {'1': 'discord_users', '3': 5, '4': 1, '5': 5, '10': 'discordUsers'},
    const {'1': 'telegram_chats', '3': 6, '4': 1, '5': 5, '10': 'telegramChats'},
    const {'1': 'discord_channels', '3': 7, '4': 1, '5': 5, '10': 'discordChannels'},
    const {'1': 'subscriptions', '3': 8, '4': 1, '5': 5, '10': 'subscriptions'},
    const {'1': 'telegram_subscriptions', '3': 9, '4': 1, '5': 5, '10': 'telegramSubscriptions'},
    const {'1': 'discord_subscriptions', '3': 10, '4': 1, '5': 5, '10': 'discordSubscriptions'},
  ],
};

/// Descriptor for `GetStatsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getStatsResponseDescriptor = $convert.base64Decode('ChBHZXRTdGF0c1Jlc3BvbnNlEhYKBmNoYWlucxgBIAEoBVIGY2hhaW5zEhwKCWNvbnRyYWN0cxgCIAEoBVIJY29udHJhY3RzEhQKBXVzZXJzGAMgASgFUgV1c2VycxIlCg50ZWxlZ3JhbV91c2VycxgEIAEoBVINdGVsZWdyYW1Vc2VycxIjCg1kaXNjb3JkX3VzZXJzGAUgASgFUgxkaXNjb3JkVXNlcnMSJQoOdGVsZWdyYW1fY2hhdHMYBiABKAVSDXRlbGVncmFtQ2hhdHMSKQoQZGlzY29yZF9jaGFubmVscxgHIAEoBVIPZGlzY29yZENoYW5uZWxzEiQKDXN1YnNjcmlwdGlvbnMYCCABKAVSDXN1YnNjcmlwdGlvbnMSNQoWdGVsZWdyYW1fc3Vic2NyaXB0aW9ucxgJIAEoBVIVdGVsZWdyYW1TdWJzY3JpcHRpb25zEjMKFWRpc2NvcmRfc3Vic2NyaXB0aW9ucxgKIAEoBVIUZGlzY29yZFN1YnNjcmlwdGlvbnM=');
