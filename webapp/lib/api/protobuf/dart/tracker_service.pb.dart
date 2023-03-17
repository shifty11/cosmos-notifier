///
//  Generated code. Do not modify.
//  source: tracker_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class IsAddressValidRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IsAddressValidRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'address')
    ..hasRequiredFields = false
  ;

  IsAddressValidRequest._() : super();
  factory IsAddressValidRequest({
    $core.String? address,
  }) {
    final _result = create();
    if (address != null) {
      _result.address = address;
    }
    return _result;
  }
  factory IsAddressValidRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IsAddressValidRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IsAddressValidRequest clone() => IsAddressValidRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IsAddressValidRequest copyWith(void Function(IsAddressValidRequest) updates) => super.copyWith((message) => updates(message as IsAddressValidRequest)) as IsAddressValidRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IsAddressValidRequest create() => IsAddressValidRequest._();
  IsAddressValidRequest createEmptyInstance() => create();
  static $pb.PbList<IsAddressValidRequest> createRepeated() => $pb.PbList<IsAddressValidRequest>();
  @$core.pragma('dart2js:noInline')
  static IsAddressValidRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IsAddressValidRequest>(create);
  static IsAddressValidRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get address => $_getSZ(0);
  @$pb.TagNumber(1)
  set address($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAddress() => $_has(0);
  @$pb.TagNumber(1)
  void clearAddress() => clearField(1);
}

class IsAddressValidResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IsAddressValidResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOB(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isValid', protoName: 'isValid')
    ..hasRequiredFields = false
  ;

  IsAddressValidResponse._() : super();
  factory IsAddressValidResponse({
    $core.bool? isValid,
  }) {
    final _result = create();
    if (isValid != null) {
      _result.isValid = isValid;
    }
    return _result;
  }
  factory IsAddressValidResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IsAddressValidResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IsAddressValidResponse clone() => IsAddressValidResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IsAddressValidResponse copyWith(void Function(IsAddressValidResponse) updates) => super.copyWith((message) => updates(message as IsAddressValidResponse)) as IsAddressValidResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IsAddressValidResponse create() => IsAddressValidResponse._();
  IsAddressValidResponse createEmptyInstance() => create();
  static $pb.PbList<IsAddressValidResponse> createRepeated() => $pb.PbList<IsAddressValidResponse>();
  @$core.pragma('dart2js:noInline')
  static IsAddressValidResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IsAddressValidResponse>(create);
  static IsAddressValidResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isValid => $_getBF(0);
  @$pb.TagNumber(1)
  set isValid($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsValid() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsValid() => clearField(1);
}

class AddTrackerRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddTrackerRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'address')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationInterval', protoName: 'notificationInterval')
    ..aInt64(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'discordChannelId', protoName: 'discordChannelId')
    ..aInt64(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'telegramChatId', protoName: 'telegramChatId')
    ..hasRequiredFields = false
  ;

  AddTrackerRequest._() : super();
  factory AddTrackerRequest({
    $core.String? address,
    $fixnum.Int64? notificationInterval,
    $fixnum.Int64? discordChannelId,
    $fixnum.Int64? telegramChatId,
  }) {
    final _result = create();
    if (address != null) {
      _result.address = address;
    }
    if (notificationInterval != null) {
      _result.notificationInterval = notificationInterval;
    }
    if (discordChannelId != null) {
      _result.discordChannelId = discordChannelId;
    }
    if (telegramChatId != null) {
      _result.telegramChatId = telegramChatId;
    }
    return _result;
  }
  factory AddTrackerRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddTrackerRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddTrackerRequest clone() => AddTrackerRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddTrackerRequest copyWith(void Function(AddTrackerRequest) updates) => super.copyWith((message) => updates(message as AddTrackerRequest)) as AddTrackerRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AddTrackerRequest create() => AddTrackerRequest._();
  AddTrackerRequest createEmptyInstance() => create();
  static $pb.PbList<AddTrackerRequest> createRepeated() => $pb.PbList<AddTrackerRequest>();
  @$core.pragma('dart2js:noInline')
  static AddTrackerRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddTrackerRequest>(create);
  static AddTrackerRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get address => $_getSZ(0);
  @$pb.TagNumber(1)
  set address($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAddress() => $_has(0);
  @$pb.TagNumber(1)
  void clearAddress() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get notificationInterval => $_getI64(1);
  @$pb.TagNumber(2)
  set notificationInterval($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotificationInterval() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotificationInterval() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get discordChannelId => $_getI64(2);
  @$pb.TagNumber(3)
  set discordChannelId($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDiscordChannelId() => $_has(2);
  @$pb.TagNumber(3)
  void clearDiscordChannelId() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get telegramChatId => $_getI64(3);
  @$pb.TagNumber(4)
  set telegramChatId($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTelegramChatId() => $_has(3);
  @$pb.TagNumber(4)
  void clearTelegramChatId() => clearField(4);
}

class AddTrackerResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddTrackerResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'address')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationInterval', protoName: 'notificationInterval')
    ..aInt64(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'discordChannelId', protoName: 'discordChannelId')
    ..aInt64(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'telegramChatId', protoName: 'telegramChatId')
    ..a<$core.int>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'trackerId', $pb.PbFieldType.O3, protoName: 'trackerId')
    ..hasRequiredFields = false
  ;

  AddTrackerResponse._() : super();
  factory AddTrackerResponse({
    $core.String? address,
    $fixnum.Int64? notificationInterval,
    $fixnum.Int64? discordChannelId,
    $fixnum.Int64? telegramChatId,
    $core.int? trackerId,
  }) {
    final _result = create();
    if (address != null) {
      _result.address = address;
    }
    if (notificationInterval != null) {
      _result.notificationInterval = notificationInterval;
    }
    if (discordChannelId != null) {
      _result.discordChannelId = discordChannelId;
    }
    if (telegramChatId != null) {
      _result.telegramChatId = telegramChatId;
    }
    if (trackerId != null) {
      _result.trackerId = trackerId;
    }
    return _result;
  }
  factory AddTrackerResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddTrackerResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddTrackerResponse clone() => AddTrackerResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddTrackerResponse copyWith(void Function(AddTrackerResponse) updates) => super.copyWith((message) => updates(message as AddTrackerResponse)) as AddTrackerResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AddTrackerResponse create() => AddTrackerResponse._();
  AddTrackerResponse createEmptyInstance() => create();
  static $pb.PbList<AddTrackerResponse> createRepeated() => $pb.PbList<AddTrackerResponse>();
  @$core.pragma('dart2js:noInline')
  static AddTrackerResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddTrackerResponse>(create);
  static AddTrackerResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get address => $_getSZ(0);
  @$pb.TagNumber(1)
  set address($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAddress() => $_has(0);
  @$pb.TagNumber(1)
  void clearAddress() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get notificationInterval => $_getI64(1);
  @$pb.TagNumber(2)
  set notificationInterval($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotificationInterval() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotificationInterval() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get discordChannelId => $_getI64(2);
  @$pb.TagNumber(3)
  set discordChannelId($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDiscordChannelId() => $_has(2);
  @$pb.TagNumber(3)
  void clearDiscordChannelId() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get telegramChatId => $_getI64(3);
  @$pb.TagNumber(4)
  set telegramChatId($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTelegramChatId() => $_has(3);
  @$pb.TagNumber(4)
  void clearTelegramChatId() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get trackerId => $_getIZ(4);
  @$pb.TagNumber(5)
  set trackerId($core.int v) { $_setSignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTrackerId() => $_has(4);
  @$pb.TagNumber(5)
  void clearTrackerId() => clearField(5);
}

class DeleteTrackerRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteTrackerRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'trackerId', protoName: 'trackerId')
    ..hasRequiredFields = false
  ;

  DeleteTrackerRequest._() : super();
  factory DeleteTrackerRequest({
    $fixnum.Int64? trackerId,
  }) {
    final _result = create();
    if (trackerId != null) {
      _result.trackerId = trackerId;
    }
    return _result;
  }
  factory DeleteTrackerRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteTrackerRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteTrackerRequest clone() => DeleteTrackerRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteTrackerRequest copyWith(void Function(DeleteTrackerRequest) updates) => super.copyWith((message) => updates(message as DeleteTrackerRequest)) as DeleteTrackerRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteTrackerRequest create() => DeleteTrackerRequest._();
  DeleteTrackerRequest createEmptyInstance() => create();
  static $pb.PbList<DeleteTrackerRequest> createRepeated() => $pb.PbList<DeleteTrackerRequest>();
  @$core.pragma('dart2js:noInline')
  static DeleteTrackerRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteTrackerRequest>(create);
  static DeleteTrackerRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get trackerId => $_getI64(0);
  @$pb.TagNumber(1)
  set trackerId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTrackerId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTrackerId() => clearField(1);
}

