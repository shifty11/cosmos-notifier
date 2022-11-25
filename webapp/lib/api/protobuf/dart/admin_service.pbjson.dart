///
//  Generated code. Do not modify.
//  source: admin_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

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
    const {'1': 'TELEGRAM_TEST', '2': 0},
    const {'1': 'DISCORD_TEST', '2': 1},
    const {'1': 'TELEGRAM', '2': 2},
    const {'1': 'DISCORD', '2': 3},
  ],
};

/// Descriptor for `BroadcastMessageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List broadcastMessageRequestDescriptor = $convert.base64Decode('ChdCcm9hZGNhc3RNZXNzYWdlUmVxdWVzdBIYCgdtZXNzYWdlGAEgASgJUgdtZXNzYWdlEk0KBHR5cGUYAiABKA4yOS5jb3Ntb3Nfbm90aWZpZXJfZ3JwYy5Ccm9hZGNhc3RNZXNzYWdlUmVxdWVzdC5NZXNzYWdlVHlwZVIEdHlwZSJNCgtNZXNzYWdlVHlwZRIRCg1URUxFR1JBTV9URVNUEAASEAoMRElTQ09SRF9URVNUEAESDAoIVEVMRUdSQU0QAhILCgdESVNDT1JEEAM=');
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
    const {'1': 'SENDING', '2': 0},
    const {'1': 'SENT', '2': 1},
    const {'1': 'FAILED', '2': 3},
  ],
};

/// Descriptor for `BroadcastMessageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List broadcastMessageResponseDescriptor = $convert.base64Decode('ChhCcm9hZGNhc3RNZXNzYWdlUmVzcG9uc2USTQoGc3RhdHVzGAEgASgOMjUuY29zbW9zX25vdGlmaWVyX2dycGMuQnJvYWRjYXN0TWVzc2FnZVJlc3BvbnNlLlN0YXR1c1IGc3RhdHVzEhoKCHJlc3BvbnNlGAIgASgJUghyZXNwb25zZSIrCgZTdGF0dXMSCwoHU0VORElORxAAEggKBFNFTlQQARIKCgZGQUlMRUQQAw==');
