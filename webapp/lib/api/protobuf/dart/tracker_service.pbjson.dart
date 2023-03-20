///
//  Generated code. Do not modify.
//  source: tracker_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use isAddressValidRequestDescriptor instead')
const IsAddressValidRequest$json = const {
  '1': 'IsAddressValidRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
  ],
};

/// Descriptor for `IsAddressValidRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List isAddressValidRequestDescriptor = $convert.base64Decode('ChVJc0FkZHJlc3NWYWxpZFJlcXVlc3QSGAoHYWRkcmVzcxgBIAEoCVIHYWRkcmVzcw==');
@$core.Deprecated('Use isAddressValidResponseDescriptor instead')
const IsAddressValidResponse$json = const {
  '1': 'IsAddressValidResponse',
  '2': const [
    const {'1': 'isValid', '3': 1, '4': 1, '5': 8, '10': 'isValid'},
  ],
};

/// Descriptor for `IsAddressValidResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List isAddressValidResponseDescriptor = $convert.base64Decode('ChZJc0FkZHJlc3NWYWxpZFJlc3BvbnNlEhgKB2lzVmFsaWQYASABKAhSB2lzVmFsaWQ=');
@$core.Deprecated('Use addTrackerRequestDescriptor instead')
const AddTrackerRequest$json = const {
  '1': 'AddTrackerRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'notificationInterval', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'notificationInterval'},
    const {'1': 'discordChannelId', '3': 3, '4': 1, '5': 3, '10': 'discordChannelId'},
    const {'1': 'telegramChatId', '3': 4, '4': 1, '5': 3, '10': 'telegramChatId'},
  ],
};

/// Descriptor for `AddTrackerRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addTrackerRequestDescriptor = $convert.base64Decode('ChFBZGRUcmFja2VyUmVxdWVzdBIYCgdhZGRyZXNzGAEgASgJUgdhZGRyZXNzEk0KFG5vdGlmaWNhdGlvbkludGVydmFsGAIgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uUhRub3RpZmljYXRpb25JbnRlcnZhbBIqChBkaXNjb3JkQ2hhbm5lbElkGAMgASgDUhBkaXNjb3JkQ2hhbm5lbElkEiYKDnRlbGVncmFtQ2hhdElkGAQgASgDUg50ZWxlZ3JhbUNoYXRJZA==');
@$core.Deprecated('Use addTrackerResponseDescriptor instead')
const AddTrackerResponse$json = const {
  '1': 'AddTrackerResponse',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'notificationInterval', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'notificationInterval'},
    const {'1': 'discordChannelId', '3': 3, '4': 1, '5': 3, '10': 'discordChannelId'},
    const {'1': 'telegramChatId', '3': 4, '4': 1, '5': 3, '10': 'telegramChatId'},
    const {'1': 'trackerId', '3': 5, '4': 1, '5': 3, '10': 'trackerId'},
  ],
};

/// Descriptor for `AddTrackerResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addTrackerResponseDescriptor = $convert.base64Decode('ChJBZGRUcmFja2VyUmVzcG9uc2USGAoHYWRkcmVzcxgBIAEoCVIHYWRkcmVzcxJNChRub3RpZmljYXRpb25JbnRlcnZhbBgCIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvblIUbm90aWZpY2F0aW9uSW50ZXJ2YWwSKgoQZGlzY29yZENoYW5uZWxJZBgDIAEoA1IQZGlzY29yZENoYW5uZWxJZBImCg50ZWxlZ3JhbUNoYXRJZBgEIAEoA1IOdGVsZWdyYW1DaGF0SWQSHAoJdHJhY2tlcklkGAUgASgDUgl0cmFja2VySWQ=');
@$core.Deprecated('Use deleteTrackerRequestDescriptor instead')
const DeleteTrackerRequest$json = const {
  '1': 'DeleteTrackerRequest',
  '2': const [
    const {'1': 'trackerId', '3': 1, '4': 1, '5': 3, '10': 'trackerId'},
  ],
};

/// Descriptor for `DeleteTrackerRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTrackerRequestDescriptor = $convert.base64Decode('ChREZWxldGVUcmFja2VyUmVxdWVzdBIcCgl0cmFja2VySWQYASABKANSCXRyYWNrZXJJZA==');
