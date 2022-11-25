///
//  Generated code. Do not modify.
//  source: admin_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'admin_service.pbenum.dart';

export 'admin_service.pbenum.dart';

class BroadcastMessageRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BroadcastMessageRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message')
    ..e<BroadcastMessageRequest_MessageType>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: BroadcastMessageRequest_MessageType.TELEGRAM_TEST, valueOf: BroadcastMessageRequest_MessageType.valueOf, enumValues: BroadcastMessageRequest_MessageType.values)
    ..hasRequiredFields = false
  ;

  BroadcastMessageRequest._() : super();
  factory BroadcastMessageRequest({
    $core.String? message,
    BroadcastMessageRequest_MessageType? type,
  }) {
    final _result = create();
    if (message != null) {
      _result.message = message;
    }
    if (type != null) {
      _result.type = type;
    }
    return _result;
  }
  factory BroadcastMessageRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BroadcastMessageRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BroadcastMessageRequest clone() => BroadcastMessageRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BroadcastMessageRequest copyWith(void Function(BroadcastMessageRequest) updates) => super.copyWith((message) => updates(message as BroadcastMessageRequest)) as BroadcastMessageRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BroadcastMessageRequest create() => BroadcastMessageRequest._();
  BroadcastMessageRequest createEmptyInstance() => create();
  static $pb.PbList<BroadcastMessageRequest> createRepeated() => $pb.PbList<BroadcastMessageRequest>();
  @$core.pragma('dart2js:noInline')
  static BroadcastMessageRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BroadcastMessageRequest>(create);
  static BroadcastMessageRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get message => $_getSZ(0);
  @$pb.TagNumber(1)
  set message($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMessage() => $_has(0);
  @$pb.TagNumber(1)
  void clearMessage() => clearField(1);

  @$pb.TagNumber(2)
  BroadcastMessageRequest_MessageType get type => $_getN(1);
  @$pb.TagNumber(2)
  set type(BroadcastMessageRequest_MessageType v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasType() => $_has(1);
  @$pb.TagNumber(2)
  void clearType() => clearField(2);
}

class BroadcastMessageResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BroadcastMessageResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..e<BroadcastMessageResponse_Status>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: BroadcastMessageResponse_Status.SENDING, valueOf: BroadcastMessageResponse_Status.valueOf, enumValues: BroadcastMessageResponse_Status.values)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'response')
    ..hasRequiredFields = false
  ;

  BroadcastMessageResponse._() : super();
  factory BroadcastMessageResponse({
    BroadcastMessageResponse_Status? status,
    $core.String? response,
  }) {
    final _result = create();
    if (status != null) {
      _result.status = status;
    }
    if (response != null) {
      _result.response = response;
    }
    return _result;
  }
  factory BroadcastMessageResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BroadcastMessageResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BroadcastMessageResponse clone() => BroadcastMessageResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BroadcastMessageResponse copyWith(void Function(BroadcastMessageResponse) updates) => super.copyWith((message) => updates(message as BroadcastMessageResponse)) as BroadcastMessageResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BroadcastMessageResponse create() => BroadcastMessageResponse._();
  BroadcastMessageResponse createEmptyInstance() => create();
  static $pb.PbList<BroadcastMessageResponse> createRepeated() => $pb.PbList<BroadcastMessageResponse>();
  @$core.pragma('dart2js:noInline')
  static BroadcastMessageResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BroadcastMessageResponse>(create);
  static BroadcastMessageResponse? _defaultInstance;

  @$pb.TagNumber(1)
  BroadcastMessageResponse_Status get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(BroadcastMessageResponse_Status v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get response => $_getSZ(1);
  @$pb.TagNumber(2)
  set response($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasResponse() => $_has(1);
  @$pb.TagNumber(2)
  void clearResponse() => clearField(2);
}

