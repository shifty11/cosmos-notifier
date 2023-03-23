///
//  Generated code. Do not modify.
//  source: pbcommon.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use discordTypeDescriptor instead')
const DiscordType$json = const {
  '1': 'DiscordType',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
    const {'1': 'channel_id', '3': 2, '4': 1, '5': 3, '10': 'channelId'},
  ],
};

/// Descriptor for `DiscordType`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List discordTypeDescriptor = $convert.base64Decode('CgtEaXNjb3JkVHlwZRIOCgJpZBgBIAEoBVICaWQSHQoKY2hhbm5lbF9pZBgCIAEoA1IJY2hhbm5lbElk');
@$core.Deprecated('Use telegramTypeDescriptor instead')
const TelegramType$json = const {
  '1': 'TelegramType',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
    const {'1': 'chat_id', '3': 2, '4': 1, '5': 3, '10': 'chatId'},
  ],
};

/// Descriptor for `TelegramType`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List telegramTypeDescriptor = $convert.base64Decode('CgxUZWxlZ3JhbVR5cGUSDgoCaWQYASABKAVSAmlkEhcKB2NoYXRfaWQYAiABKANSBmNoYXRJZA==');
