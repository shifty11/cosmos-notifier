///
//  Generated code. Do not modify.
//  source: admin_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'admin_service.pbenum.dart';

export 'admin_service.pbenum.dart';

class BroadcastMessageRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BroadcastMessageRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message')
    ..e<BroadcastMessageRequest_MessageType>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: BroadcastMessageRequest_MessageType.MESSAGE_TYPE_UNSPECIFIED, valueOf: BroadcastMessageRequest_MessageType.valueOf, enumValues: BroadcastMessageRequest_MessageType.values)
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
    ..e<BroadcastMessageResponse_Status>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: BroadcastMessageResponse_Status.STATUS_UNSPECIFIED, valueOf: BroadcastMessageResponse_Status.valueOf, enumValues: BroadcastMessageResponse_Status.values)
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

class GetStatsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetStatsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chains', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contracts', $pb.PbFieldType.O3)
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'users', $pb.PbFieldType.O3)
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'telegramUsers', $pb.PbFieldType.O3)
    ..a<$core.int>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'discordUsers', $pb.PbFieldType.O3)
    ..a<$core.int>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'telegramChats', $pb.PbFieldType.O3)
    ..a<$core.int>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'discordChannels', $pb.PbFieldType.O3)
    ..a<$core.int>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subscriptions', $pb.PbFieldType.O3)
    ..a<$core.int>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'telegramSubscriptions', $pb.PbFieldType.O3)
    ..a<$core.int>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'discordSubscriptions', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  GetStatsResponse._() : super();
  factory GetStatsResponse({
    $core.int? chains,
    $core.int? contracts,
    $core.int? users,
    $core.int? telegramUsers,
    $core.int? discordUsers,
    $core.int? telegramChats,
    $core.int? discordChannels,
    $core.int? subscriptions,
    $core.int? telegramSubscriptions,
    $core.int? discordSubscriptions,
  }) {
    final _result = create();
    if (chains != null) {
      _result.chains = chains;
    }
    if (contracts != null) {
      _result.contracts = contracts;
    }
    if (users != null) {
      _result.users = users;
    }
    if (telegramUsers != null) {
      _result.telegramUsers = telegramUsers;
    }
    if (discordUsers != null) {
      _result.discordUsers = discordUsers;
    }
    if (telegramChats != null) {
      _result.telegramChats = telegramChats;
    }
    if (discordChannels != null) {
      _result.discordChannels = discordChannels;
    }
    if (subscriptions != null) {
      _result.subscriptions = subscriptions;
    }
    if (telegramSubscriptions != null) {
      _result.telegramSubscriptions = telegramSubscriptions;
    }
    if (discordSubscriptions != null) {
      _result.discordSubscriptions = discordSubscriptions;
    }
    return _result;
  }
  factory GetStatsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetStatsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetStatsResponse clone() => GetStatsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetStatsResponse copyWith(void Function(GetStatsResponse) updates) => super.copyWith((message) => updates(message as GetStatsResponse)) as GetStatsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetStatsResponse create() => GetStatsResponse._();
  GetStatsResponse createEmptyInstance() => create();
  static $pb.PbList<GetStatsResponse> createRepeated() => $pb.PbList<GetStatsResponse>();
  @$core.pragma('dart2js:noInline')
  static GetStatsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetStatsResponse>(create);
  static GetStatsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get chains => $_getIZ(0);
  @$pb.TagNumber(1)
  set chains($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChains() => $_has(0);
  @$pb.TagNumber(1)
  void clearChains() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get contracts => $_getIZ(1);
  @$pb.TagNumber(2)
  set contracts($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasContracts() => $_has(1);
  @$pb.TagNumber(2)
  void clearContracts() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get users => $_getIZ(2);
  @$pb.TagNumber(3)
  set users($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasUsers() => $_has(2);
  @$pb.TagNumber(3)
  void clearUsers() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get telegramUsers => $_getIZ(3);
  @$pb.TagNumber(4)
  set telegramUsers($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTelegramUsers() => $_has(3);
  @$pb.TagNumber(4)
  void clearTelegramUsers() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get discordUsers => $_getIZ(4);
  @$pb.TagNumber(5)
  set discordUsers($core.int v) { $_setSignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasDiscordUsers() => $_has(4);
  @$pb.TagNumber(5)
  void clearDiscordUsers() => clearField(5);

  @$pb.TagNumber(6)
  $core.int get telegramChats => $_getIZ(5);
  @$pb.TagNumber(6)
  set telegramChats($core.int v) { $_setSignedInt32(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasTelegramChats() => $_has(5);
  @$pb.TagNumber(6)
  void clearTelegramChats() => clearField(6);

  @$pb.TagNumber(7)
  $core.int get discordChannels => $_getIZ(6);
  @$pb.TagNumber(7)
  set discordChannels($core.int v) { $_setSignedInt32(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasDiscordChannels() => $_has(6);
  @$pb.TagNumber(7)
  void clearDiscordChannels() => clearField(7);

  @$pb.TagNumber(8)
  $core.int get subscriptions => $_getIZ(7);
  @$pb.TagNumber(8)
  set subscriptions($core.int v) { $_setSignedInt32(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasSubscriptions() => $_has(7);
  @$pb.TagNumber(8)
  void clearSubscriptions() => clearField(8);

  @$pb.TagNumber(9)
  $core.int get telegramSubscriptions => $_getIZ(8);
  @$pb.TagNumber(9)
  set telegramSubscriptions($core.int v) { $_setSignedInt32(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTelegramSubscriptions() => $_has(8);
  @$pb.TagNumber(9)
  void clearTelegramSubscriptions() => clearField(9);

  @$pb.TagNumber(10)
  $core.int get discordSubscriptions => $_getIZ(9);
  @$pb.TagNumber(10)
  set discordSubscriptions($core.int v) { $_setSignedInt32(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasDiscordSubscriptions() => $_has(9);
  @$pb.TagNumber(10)
  void clearDiscordSubscriptions() => clearField(10);
}

