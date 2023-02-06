///
//  Generated code. Do not modify.
//  source: auth_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use telegramLoginRequestDescriptor instead')
const TelegramLoginRequest$json = const {
  '1': 'TelegramLoginRequest',
  '2': const [
    const {'1': 'userId', '3': 1, '4': 1, '5': 3, '10': 'userId'},
    const {'1': 'dataStr', '3': 2, '4': 1, '5': 9, '10': 'dataStr'},
    const {'1': 'username', '3': 4, '4': 1, '5': 9, '10': 'username'},
    const {'1': 'authDate', '3': 6, '4': 1, '5': 3, '10': 'authDate'},
    const {'1': 'hash', '3': 3, '4': 1, '5': 9, '10': 'hash'},
  ],
};

/// Descriptor for `TelegramLoginRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List telegramLoginRequestDescriptor = $convert.base64Decode('ChRUZWxlZ3JhbUxvZ2luUmVxdWVzdBIWCgZ1c2VySWQYASABKANSBnVzZXJJZBIYCgdkYXRhU3RyGAIgASgJUgdkYXRhU3RyEhoKCHVzZXJuYW1lGAQgASgJUgh1c2VybmFtZRIaCghhdXRoRGF0ZRgGIAEoA1IIYXV0aERhdGUSEgoEaGFzaBgDIAEoCVIEaGFzaA==');
@$core.Deprecated('Use discordLoginRequestDescriptor instead')
const DiscordLoginRequest$json = const {
  '1': 'DiscordLoginRequest',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 9, '10': 'code'},
  ],
};

/// Descriptor for `DiscordLoginRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List discordLoginRequestDescriptor = $convert.base64Decode('ChNEaXNjb3JkTG9naW5SZXF1ZXN0EhIKBGNvZGUYASABKAlSBGNvZGU=');
@$core.Deprecated('Use loginResponseDescriptor instead')
const LoginResponse$json = const {
  '1': 'LoginResponse',
  '2': const [
    const {'1': 'access_token', '3': 1, '4': 1, '5': 9, '10': 'accessToken'},
    const {'1': 'refresh_token', '3': 2, '4': 1, '5': 9, '10': 'refreshToken'},
  ],
};

/// Descriptor for `LoginResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginResponseDescriptor = $convert.base64Decode('Cg1Mb2dpblJlc3BvbnNlEiEKDGFjY2Vzc190b2tlbhgBIAEoCVILYWNjZXNzVG9rZW4SIwoNcmVmcmVzaF90b2tlbhgCIAEoCVIMcmVmcmVzaFRva2Vu');
@$core.Deprecated('Use refreshAccessTokenRequestDescriptor instead')
const RefreshAccessTokenRequest$json = const {
  '1': 'RefreshAccessTokenRequest',
  '2': const [
    const {'1': 'refresh_token', '3': 1, '4': 1, '5': 9, '10': 'refreshToken'},
  ],
};

/// Descriptor for `RefreshAccessTokenRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List refreshAccessTokenRequestDescriptor = $convert.base64Decode('ChlSZWZyZXNoQWNjZXNzVG9rZW5SZXF1ZXN0EiMKDXJlZnJlc2hfdG9rZW4YASABKAlSDHJlZnJlc2hUb2tlbg==');
@$core.Deprecated('Use refreshAccessTokenResponseDescriptor instead')
const RefreshAccessTokenResponse$json = const {
  '1': 'RefreshAccessTokenResponse',
  '2': const [
    const {'1': 'access_token', '3': 1, '4': 1, '5': 9, '10': 'accessToken'},
  ],
};

/// Descriptor for `RefreshAccessTokenResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List refreshAccessTokenResponseDescriptor = $convert.base64Decode('ChpSZWZyZXNoQWNjZXNzVG9rZW5SZXNwb25zZRIhCgxhY2Nlc3NfdG9rZW4YASABKAlSC2FjY2Vzc1Rva2Vu');
