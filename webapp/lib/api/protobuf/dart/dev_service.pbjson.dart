///
//  Generated code. Do not modify.
//  source: dev_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use devLoginRequestDescriptor instead')
const DevLoginRequest$json = const {
  '1': 'DevLoginRequest',
  '2': const [
    const {'1': 'user_id', '3': 1, '4': 1, '5': 3, '10': 'userId'},
    const {'1': 'user_type', '3': 2, '4': 1, '5': 14, '6': '.cosmos_notifier_grpc.DevLoginRequest.UserType', '10': 'userType'},
    const {'1': 'role', '3': 3, '4': 1, '5': 14, '6': '.cosmos_notifier_grpc.DevLoginRequest.Role', '10': 'role'},
  ],
  '4': const [DevLoginRequest_UserType$json, DevLoginRequest_Role$json],
};

@$core.Deprecated('Use devLoginRequestDescriptor instead')
const DevLoginRequest_UserType$json = const {
  '1': 'UserType',
  '2': const [
    const {'1': 'USER_TYPE_UNSPECIFIED', '2': 0},
    const {'1': 'TELEGRAM', '2': 1},
    const {'1': 'DISCORD', '2': 2},
  ],
};

@$core.Deprecated('Use devLoginRequestDescriptor instead')
const DevLoginRequest_Role$json = const {
  '1': 'Role',
  '2': const [
    const {'1': 'ROLE_UNSPECIFIED', '2': 0},
    const {'1': 'ADMIN', '2': 1},
    const {'1': 'USER', '2': 2},
  ],
};

/// Descriptor for `DevLoginRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List devLoginRequestDescriptor = $convert.base64Decode('Cg9EZXZMb2dpblJlcXVlc3QSFwoHdXNlcl9pZBgBIAEoA1IGdXNlcklkEksKCXVzZXJfdHlwZRgCIAEoDjIuLmNvc21vc19ub3RpZmllcl9ncnBjLkRldkxvZ2luUmVxdWVzdC5Vc2VyVHlwZVIIdXNlclR5cGUSPgoEcm9sZRgDIAEoDjIqLmNvc21vc19ub3RpZmllcl9ncnBjLkRldkxvZ2luUmVxdWVzdC5Sb2xlUgRyb2xlIkAKCFVzZXJUeXBlEhkKFVVTRVJfVFlQRV9VTlNQRUNJRklFRBAAEgwKCFRFTEVHUkFNEAESCwoHRElTQ09SRBACIjEKBFJvbGUSFAoQUk9MRV9VTlNQRUNJRklFRBAAEgkKBUFETUlOEAESCAoEVVNFUhAC');
