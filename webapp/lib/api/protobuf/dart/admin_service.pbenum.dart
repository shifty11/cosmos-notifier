///
//  Generated code. Do not modify.
//  source: admin_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class BroadcastMessageRequest_MessageType extends $pb.ProtobufEnum {
  static const BroadcastMessageRequest_MessageType TELEGRAM_TEST = BroadcastMessageRequest_MessageType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TELEGRAM_TEST');
  static const BroadcastMessageRequest_MessageType DISCORD_TEST = BroadcastMessageRequest_MessageType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DISCORD_TEST');
  static const BroadcastMessageRequest_MessageType TELEGRAM = BroadcastMessageRequest_MessageType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TELEGRAM');
  static const BroadcastMessageRequest_MessageType DISCORD = BroadcastMessageRequest_MessageType._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DISCORD');

  static const $core.List<BroadcastMessageRequest_MessageType> values = <BroadcastMessageRequest_MessageType> [
    TELEGRAM_TEST,
    DISCORD_TEST,
    TELEGRAM,
    DISCORD,
  ];

  static final $core.Map<$core.int, BroadcastMessageRequest_MessageType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static BroadcastMessageRequest_MessageType? valueOf($core.int value) => _byValue[value];

  const BroadcastMessageRequest_MessageType._($core.int v, $core.String n) : super(v, n);
}

class BroadcastMessageResponse_Status extends $pb.ProtobufEnum {
  static const BroadcastMessageResponse_Status SENDING = BroadcastMessageResponse_Status._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SENDING');
  static const BroadcastMessageResponse_Status SENT = BroadcastMessageResponse_Status._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SENT');
  static const BroadcastMessageResponse_Status FAILED = BroadcastMessageResponse_Status._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'FAILED');

  static const $core.List<BroadcastMessageResponse_Status> values = <BroadcastMessageResponse_Status> [
    SENDING,
    SENT,
    FAILED,
  ];

  static final $core.Map<$core.int, BroadcastMessageResponse_Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static BroadcastMessageResponse_Status? valueOf($core.int value) => _byValue[value];

  const BroadcastMessageResponse_Status._($core.int v, $core.String n) : super(v, n);
}

