///
//  Generated code. Do not modify.
//  source: dev_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'dev_service.pbenum.dart';

export 'dev_service.pbenum.dart';

class DevLoginRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DevLoginRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userId')
    ..e<DevLoginRequest_UserType>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: DevLoginRequest_UserType.TELEGRAM, valueOf: DevLoginRequest_UserType.valueOf, enumValues: DevLoginRequest_UserType.values)
    ..e<DevLoginRequest_Role>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'role', $pb.PbFieldType.OE, defaultOrMaker: DevLoginRequest_Role.ADMIN, valueOf: DevLoginRequest_Role.valueOf, enumValues: DevLoginRequest_Role.values)
    ..hasRequiredFields = false
  ;

  DevLoginRequest._() : super();
  factory DevLoginRequest({
    $fixnum.Int64? userId,
    DevLoginRequest_UserType? type,
    DevLoginRequest_Role? role,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    if (type != null) {
      _result.type = type;
    }
    if (role != null) {
      _result.role = role;
    }
    return _result;
  }
  factory DevLoginRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DevLoginRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DevLoginRequest clone() => DevLoginRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DevLoginRequest copyWith(void Function(DevLoginRequest) updates) => super.copyWith((message) => updates(message as DevLoginRequest)) as DevLoginRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DevLoginRequest create() => DevLoginRequest._();
  DevLoginRequest createEmptyInstance() => create();
  static $pb.PbList<DevLoginRequest> createRepeated() => $pb.PbList<DevLoginRequest>();
  @$core.pragma('dart2js:noInline')
  static DevLoginRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DevLoginRequest>(create);
  static DevLoginRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get userId => $_getI64(0);
  @$pb.TagNumber(1)
  set userId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);

  @$pb.TagNumber(2)
  DevLoginRequest_UserType get type => $_getN(1);
  @$pb.TagNumber(2)
  set type(DevLoginRequest_UserType v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasType() => $_has(1);
  @$pb.TagNumber(2)
  void clearType() => clearField(2);

  @$pb.TagNumber(3)
  DevLoginRequest_Role get role => $_getN(2);
  @$pb.TagNumber(3)
  set role(DevLoginRequest_Role v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasRole() => $_has(2);
  @$pb.TagNumber(3)
  void clearRole() => clearField(3);
}

