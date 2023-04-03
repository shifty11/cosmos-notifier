///
//  Generated code. Do not modify.
//  source: dev_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class DevLoginRequest_UserType extends $pb.ProtobufEnum {
  static const DevLoginRequest_UserType USER_TYPE_UNSPECIFIED = DevLoginRequest_UserType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'USER_TYPE_UNSPECIFIED');
  static const DevLoginRequest_UserType TELEGRAM = DevLoginRequest_UserType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TELEGRAM');
  static const DevLoginRequest_UserType DISCORD = DevLoginRequest_UserType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DISCORD');

  static const $core.List<DevLoginRequest_UserType> values = <DevLoginRequest_UserType> [
    USER_TYPE_UNSPECIFIED,
    TELEGRAM,
    DISCORD,
  ];

  static final $core.Map<$core.int, DevLoginRequest_UserType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static DevLoginRequest_UserType? valueOf($core.int value) => _byValue[value];

  const DevLoginRequest_UserType._($core.int v, $core.String n) : super(v, n);
}

class DevLoginRequest_Role extends $pb.ProtobufEnum {
  static const DevLoginRequest_Role ROLE_UNSPECIFIED = DevLoginRequest_Role._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ROLE_UNSPECIFIED');
  static const DevLoginRequest_Role ADMIN = DevLoginRequest_Role._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ADMIN');
  static const DevLoginRequest_Role USER = DevLoginRequest_Role._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'USER');

  static const $core.List<DevLoginRequest_Role> values = <DevLoginRequest_Role> [
    ROLE_UNSPECIFIED,
    ADMIN,
    USER,
  ];

  static final $core.Map<$core.int, DevLoginRequest_Role> _byValue = $pb.ProtobufEnum.initByValue(values);
  static DevLoginRequest_Role? valueOf($core.int value) => _byValue[value];

  const DevLoginRequest_Role._($core.int v, $core.String n) : super(v, n);
}

