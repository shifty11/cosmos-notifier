///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class ChatRoom_Type extends $pb.ProtobufEnum {
  static const ChatRoom_Type TYPE_UNSPECIFIED = ChatRoom_Type._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TYPE_UNSPECIFIED');
  static const ChatRoom_Type TELEGRAM = ChatRoom_Type._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TELEGRAM');
  static const ChatRoom_Type DISCORD = ChatRoom_Type._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DISCORD');

  static const $core.List<ChatRoom_Type> values = <ChatRoom_Type> [
    TYPE_UNSPECIFIED,
    TELEGRAM,
    DISCORD,
  ];

  static final $core.Map<$core.int, ChatRoom_Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ChatRoom_Type? valueOf($core.int value) => _byValue[value];

  const ChatRoom_Type._($core.int v, $core.String n) : super(v, n);
}

class AddDaoResponse_Status extends $pb.ProtobufEnum {
  static const AddDaoResponse_Status STATUS_UNSPECIFIED = AddDaoResponse_Status._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_UNSPECIFIED');
  static const AddDaoResponse_Status ADDED = AddDaoResponse_Status._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ADDED');
  static const AddDaoResponse_Status ALREADY_ADDED = AddDaoResponse_Status._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ALREADY_ADDED');
  static const AddDaoResponse_Status IS_ADDING = AddDaoResponse_Status._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'IS_ADDING');
  static const AddDaoResponse_Status FAILED = AddDaoResponse_Status._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'FAILED');

  static const $core.List<AddDaoResponse_Status> values = <AddDaoResponse_Status> [
    STATUS_UNSPECIFIED,
    ADDED,
    ALREADY_ADDED,
    IS_ADDING,
    FAILED,
  ];

  static final $core.Map<$core.int, AddDaoResponse_Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static AddDaoResponse_Status? valueOf($core.int value) => _byValue[value];

  const AddDaoResponse_Status._($core.int v, $core.String n) : super(v, n);
}

