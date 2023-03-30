///
//  Generated code. Do not modify.
//  source: tracker_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'pbcommon.pb.dart' as $6;
import 'google/protobuf/duration.pb.dart' as $7;
import 'google/protobuf/timestamp.pb.dart' as $8;

enum TrackerChatRoom_Type {
  discord, 
  telegram, 
  notSet
}

class TrackerChatRoom extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, TrackerChatRoom_Type> _TrackerChatRoom_TypeByTag = {
    2 : TrackerChatRoom_Type.discord,
    3 : TrackerChatRoom_Type.telegram,
    0 : TrackerChatRoom_Type.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TrackerChatRoom', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..oo(0, [2, 3])
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOM<$6.DiscordType>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'discord', subBuilder: $6.DiscordType.create)
    ..aOM<$6.TelegramType>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'telegram', subBuilder: $6.TelegramType.create)
    ..hasRequiredFields = false
  ;

  TrackerChatRoom._() : super();
  factory TrackerChatRoom({
    $core.String? name,
    $6.DiscordType? discord,
    $6.TelegramType? telegram,
  }) {
    final _result = create();
    if (name != null) {
      _result.name = name;
    }
    if (discord != null) {
      _result.discord = discord;
    }
    if (telegram != null) {
      _result.telegram = telegram;
    }
    return _result;
  }
  factory TrackerChatRoom.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TrackerChatRoom.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TrackerChatRoom clone() => TrackerChatRoom()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TrackerChatRoom copyWith(void Function(TrackerChatRoom) updates) => super.copyWith((message) => updates(message as TrackerChatRoom)) as TrackerChatRoom; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TrackerChatRoom create() => TrackerChatRoom._();
  TrackerChatRoom createEmptyInstance() => create();
  static $pb.PbList<TrackerChatRoom> createRepeated() => $pb.PbList<TrackerChatRoom>();
  @$core.pragma('dart2js:noInline')
  static TrackerChatRoom getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TrackerChatRoom>(create);
  static TrackerChatRoom? _defaultInstance;

  TrackerChatRoom_Type whichType() => _TrackerChatRoom_TypeByTag[$_whichOneof(0)]!;
  void clearType() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $6.DiscordType get discord => $_getN(1);
  @$pb.TagNumber(2)
  set discord($6.DiscordType v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasDiscord() => $_has(1);
  @$pb.TagNumber(2)
  void clearDiscord() => clearField(2);
  @$pb.TagNumber(2)
  $6.DiscordType ensureDiscord() => $_ensure(1);

  @$pb.TagNumber(3)
  $6.TelegramType get telegram => $_getN(2);
  @$pb.TagNumber(3)
  set telegram($6.TelegramType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTelegram() => $_has(2);
  @$pb.TagNumber(3)
  void clearTelegram() => clearField(3);
  @$pb.TagNumber(3)
  $6.TelegramType ensureTelegram() => $_ensure(2);
}

class Tracker extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Tracker', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'address')
    ..aOM<$7.Duration>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationInterval', protoName: 'notificationInterval', subBuilder: $7.Duration.create)
    ..aOM<TrackerChatRoom>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRoom', protoName: 'chatRoom', subBuilder: TrackerChatRoom.create)
    ..aOM<$8.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'updatedAt', protoName: 'updatedAt', subBuilder: $8.Timestamp.create)
    ..hasRequiredFields = false
  ;

  Tracker._() : super();
  factory Tracker({
    $fixnum.Int64? id,
    $core.String? address,
    $7.Duration? notificationInterval,
    TrackerChatRoom? chatRoom,
    $8.Timestamp? updatedAt,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (address != null) {
      _result.address = address;
    }
    if (notificationInterval != null) {
      _result.notificationInterval = notificationInterval;
    }
    if (chatRoom != null) {
      _result.chatRoom = chatRoom;
    }
    if (updatedAt != null) {
      _result.updatedAt = updatedAt;
    }
    return _result;
  }
  factory Tracker.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Tracker.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Tracker clone() => Tracker()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Tracker copyWith(void Function(Tracker) updates) => super.copyWith((message) => updates(message as Tracker)) as Tracker; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Tracker create() => Tracker._();
  Tracker createEmptyInstance() => create();
  static $pb.PbList<Tracker> createRepeated() => $pb.PbList<Tracker>();
  @$core.pragma('dart2js:noInline')
  static Tracker getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Tracker>(create);
  static Tracker? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get address => $_getSZ(1);
  @$pb.TagNumber(2)
  set address($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAddress() => $_has(1);
  @$pb.TagNumber(2)
  void clearAddress() => clearField(2);

  @$pb.TagNumber(3)
  $7.Duration get notificationInterval => $_getN(2);
  @$pb.TagNumber(3)
  set notificationInterval($7.Duration v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasNotificationInterval() => $_has(2);
  @$pb.TagNumber(3)
  void clearNotificationInterval() => clearField(3);
  @$pb.TagNumber(3)
  $7.Duration ensureNotificationInterval() => $_ensure(2);

  @$pb.TagNumber(4)
  TrackerChatRoom get chatRoom => $_getN(3);
  @$pb.TagNumber(4)
  set chatRoom(TrackerChatRoom v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasChatRoom() => $_has(3);
  @$pb.TagNumber(4)
  void clearChatRoom() => clearField(4);
  @$pb.TagNumber(4)
  TrackerChatRoom ensureChatRoom() => $_ensure(3);

  @$pb.TagNumber(5)
  $8.Timestamp get updatedAt => $_getN(4);
  @$pb.TagNumber(5)
  set updatedAt($8.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasUpdatedAt() => $_has(4);
  @$pb.TagNumber(5)
  void clearUpdatedAt() => clearField(5);
  @$pb.TagNumber(5)
  $8.Timestamp ensureUpdatedAt() => $_ensure(4);
}

class Validator extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Validator', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'address')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chainName')
    ..hasRequiredFields = false
  ;

  Validator._() : super();
  factory Validator({
    $fixnum.Int64? id,
    $core.String? address,
    $core.String? chainName,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (address != null) {
      _result.address = address;
    }
    if (chainName != null) {
      _result.chainName = chainName;
    }
    return _result;
  }
  factory Validator.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Validator.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Validator clone() => Validator()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Validator copyWith(void Function(Validator) updates) => super.copyWith((message) => updates(message as Validator)) as Validator; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Validator create() => Validator._();
  Validator createEmptyInstance() => create();
  static $pb.PbList<Validator> createRepeated() => $pb.PbList<Validator>();
  @$core.pragma('dart2js:noInline')
  static Validator getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Validator>(create);
  static Validator? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get address => $_getSZ(1);
  @$pb.TagNumber(2)
  set address($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAddress() => $_has(1);
  @$pb.TagNumber(2)
  void clearAddress() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get chainName => $_getSZ(2);
  @$pb.TagNumber(3)
  set chainName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasChainName() => $_has(2);
  @$pb.TagNumber(3)
  void clearChainName() => clearField(3);
}

class ValidatorBundle extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ValidatorBundle', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'moniker')
    ..pc<Validator>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'validators', $pb.PbFieldType.PM, subBuilder: Validator.create)
    ..aOB(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isTracked')
    ..hasRequiredFields = false
  ;

  ValidatorBundle._() : super();
  factory ValidatorBundle({
    $core.String? moniker,
    $core.Iterable<Validator>? validators,
    $core.bool? isTracked,
  }) {
    final _result = create();
    if (moniker != null) {
      _result.moniker = moniker;
    }
    if (validators != null) {
      _result.validators.addAll(validators);
    }
    if (isTracked != null) {
      _result.isTracked = isTracked;
    }
    return _result;
  }
  factory ValidatorBundle.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ValidatorBundle.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ValidatorBundle clone() => ValidatorBundle()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ValidatorBundle copyWith(void Function(ValidatorBundle) updates) => super.copyWith((message) => updates(message as ValidatorBundle)) as ValidatorBundle; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ValidatorBundle create() => ValidatorBundle._();
  ValidatorBundle createEmptyInstance() => create();
  static $pb.PbList<ValidatorBundle> createRepeated() => $pb.PbList<ValidatorBundle>();
  @$core.pragma('dart2js:noInline')
  static ValidatorBundle getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ValidatorBundle>(create);
  static ValidatorBundle? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get moniker => $_getSZ(0);
  @$pb.TagNumber(1)
  set moniker($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMoniker() => $_has(0);
  @$pb.TagNumber(1)
  void clearMoniker() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<Validator> get validators => $_getList(1);

  @$pb.TagNumber(3)
  $core.bool get isTracked => $_getBF(2);
  @$pb.TagNumber(3)
  set isTracked($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasIsTracked() => $_has(2);
  @$pb.TagNumber(3)
  void clearIsTracked() => clearField(3);
}

class GetTrackersResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetTrackersResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..pc<Tracker>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'trackers', $pb.PbFieldType.PM, subBuilder: Tracker.create)
    ..pc<TrackerChatRoom>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRooms', $pb.PbFieldType.PM, protoName: 'chatRooms', subBuilder: TrackerChatRoom.create)
    ..pc<ValidatorBundle>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'validatorBundles', $pb.PbFieldType.PM, subBuilder: ValidatorBundle.create)
    ..hasRequiredFields = false
  ;

  GetTrackersResponse._() : super();
  factory GetTrackersResponse({
    $core.Iterable<Tracker>? trackers,
    $core.Iterable<TrackerChatRoom>? chatRooms,
    $core.Iterable<ValidatorBundle>? validatorBundles,
  }) {
    final _result = create();
    if (trackers != null) {
      _result.trackers.addAll(trackers);
    }
    if (chatRooms != null) {
      _result.chatRooms.addAll(chatRooms);
    }
    if (validatorBundles != null) {
      _result.validatorBundles.addAll(validatorBundles);
    }
    return _result;
  }
  factory GetTrackersResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetTrackersResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetTrackersResponse clone() => GetTrackersResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetTrackersResponse copyWith(void Function(GetTrackersResponse) updates) => super.copyWith((message) => updates(message as GetTrackersResponse)) as GetTrackersResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetTrackersResponse create() => GetTrackersResponse._();
  GetTrackersResponse createEmptyInstance() => create();
  static $pb.PbList<GetTrackersResponse> createRepeated() => $pb.PbList<GetTrackersResponse>();
  @$core.pragma('dart2js:noInline')
  static GetTrackersResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTrackersResponse>(create);
  static GetTrackersResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Tracker> get trackers => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<TrackerChatRoom> get chatRooms => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<ValidatorBundle> get validatorBundles => $_getList(2);
}

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
    ..aOM<$7.Duration>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationInterval', protoName: 'notificationInterval', subBuilder: $7.Duration.create)
    ..aOM<TrackerChatRoom>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRoom', protoName: 'chatRoom', subBuilder: TrackerChatRoom.create)
    ..hasRequiredFields = false
  ;

  AddTrackerRequest._() : super();
  factory AddTrackerRequest({
    $core.String? address,
    $7.Duration? notificationInterval,
    TrackerChatRoom? chatRoom,
  }) {
    final _result = create();
    if (address != null) {
      _result.address = address;
    }
    if (notificationInterval != null) {
      _result.notificationInterval = notificationInterval;
    }
    if (chatRoom != null) {
      _result.chatRoom = chatRoom;
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
  $7.Duration get notificationInterval => $_getN(1);
  @$pb.TagNumber(2)
  set notificationInterval($7.Duration v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotificationInterval() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotificationInterval() => clearField(2);
  @$pb.TagNumber(2)
  $7.Duration ensureNotificationInterval() => $_ensure(1);

  @$pb.TagNumber(3)
  TrackerChatRoom get chatRoom => $_getN(2);
  @$pb.TagNumber(3)
  set chatRoom(TrackerChatRoom v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasChatRoom() => $_has(2);
  @$pb.TagNumber(3)
  void clearChatRoom() => clearField(3);
  @$pb.TagNumber(3)
  TrackerChatRoom ensureChatRoom() => $_ensure(2);
}

class UpdateTrackerRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateTrackerRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'trackerId', protoName: 'trackerId')
    ..aOM<$7.Duration>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationInterval', protoName: 'notificationInterval', subBuilder: $7.Duration.create)
    ..aOM<TrackerChatRoom>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRoom', protoName: 'chatRoom', subBuilder: TrackerChatRoom.create)
    ..hasRequiredFields = false
  ;

  UpdateTrackerRequest._() : super();
  factory UpdateTrackerRequest({
    $fixnum.Int64? trackerId,
    $7.Duration? notificationInterval,
    TrackerChatRoom? chatRoom,
  }) {
    final _result = create();
    if (trackerId != null) {
      _result.trackerId = trackerId;
    }
    if (notificationInterval != null) {
      _result.notificationInterval = notificationInterval;
    }
    if (chatRoom != null) {
      _result.chatRoom = chatRoom;
    }
    return _result;
  }
  factory UpdateTrackerRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateTrackerRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateTrackerRequest clone() => UpdateTrackerRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateTrackerRequest copyWith(void Function(UpdateTrackerRequest) updates) => super.copyWith((message) => updates(message as UpdateTrackerRequest)) as UpdateTrackerRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateTrackerRequest create() => UpdateTrackerRequest._();
  UpdateTrackerRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateTrackerRequest> createRepeated() => $pb.PbList<UpdateTrackerRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateTrackerRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateTrackerRequest>(create);
  static UpdateTrackerRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get trackerId => $_getI64(0);
  @$pb.TagNumber(1)
  set trackerId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTrackerId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTrackerId() => clearField(1);

  @$pb.TagNumber(2)
  $7.Duration get notificationInterval => $_getN(1);
  @$pb.TagNumber(2)
  set notificationInterval($7.Duration v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotificationInterval() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotificationInterval() => clearField(2);
  @$pb.TagNumber(2)
  $7.Duration ensureNotificationInterval() => $_ensure(1);

  @$pb.TagNumber(3)
  TrackerChatRoom get chatRoom => $_getN(2);
  @$pb.TagNumber(3)
  set chatRoom(TrackerChatRoom v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasChatRoom() => $_has(2);
  @$pb.TagNumber(3)
  void clearChatRoom() => clearField(3);
  @$pb.TagNumber(3)
  TrackerChatRoom ensureChatRoom() => $_ensure(2);
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

class TrackValidatorsRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TrackValidatorsRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..pPS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'monikers')
    ..aOM<$7.Duration>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationInterval', protoName: 'notificationInterval', subBuilder: $7.Duration.create)
    ..aOM<TrackerChatRoom>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRoom', protoName: 'chatRoom', subBuilder: TrackerChatRoom.create)
    ..hasRequiredFields = false
  ;

  TrackValidatorsRequest._() : super();
  factory TrackValidatorsRequest({
    $core.Iterable<$core.String>? monikers,
    $7.Duration? notificationInterval,
    TrackerChatRoom? chatRoom,
  }) {
    final _result = create();
    if (monikers != null) {
      _result.monikers.addAll(monikers);
    }
    if (notificationInterval != null) {
      _result.notificationInterval = notificationInterval;
    }
    if (chatRoom != null) {
      _result.chatRoom = chatRoom;
    }
    return _result;
  }
  factory TrackValidatorsRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TrackValidatorsRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TrackValidatorsRequest clone() => TrackValidatorsRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TrackValidatorsRequest copyWith(void Function(TrackValidatorsRequest) updates) => super.copyWith((message) => updates(message as TrackValidatorsRequest)) as TrackValidatorsRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TrackValidatorsRequest create() => TrackValidatorsRequest._();
  TrackValidatorsRequest createEmptyInstance() => create();
  static $pb.PbList<TrackValidatorsRequest> createRepeated() => $pb.PbList<TrackValidatorsRequest>();
  @$core.pragma('dart2js:noInline')
  static TrackValidatorsRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TrackValidatorsRequest>(create);
  static TrackValidatorsRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get monikers => $_getList(0);

  @$pb.TagNumber(2)
  $7.Duration get notificationInterval => $_getN(1);
  @$pb.TagNumber(2)
  set notificationInterval($7.Duration v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotificationInterval() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotificationInterval() => clearField(2);
  @$pb.TagNumber(2)
  $7.Duration ensureNotificationInterval() => $_ensure(1);

  @$pb.TagNumber(3)
  TrackerChatRoom get chatRoom => $_getN(2);
  @$pb.TagNumber(3)
  set chatRoom(TrackerChatRoom v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasChatRoom() => $_has(2);
  @$pb.TagNumber(3)
  void clearChatRoom() => clearField(3);
  @$pb.TagNumber(3)
  TrackerChatRoom ensureChatRoom() => $_ensure(2);
}

class TrackValidatorsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TrackValidatorsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..pc<Tracker>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'addedTrackers', $pb.PbFieldType.PM, protoName: 'addedTrackers', subBuilder: Tracker.create)
    ..p<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'deletedTrackerIds', $pb.PbFieldType.K6, protoName: 'deletedTrackerIds')
    ..hasRequiredFields = false
  ;

  TrackValidatorsResponse._() : super();
  factory TrackValidatorsResponse({
    $core.Iterable<Tracker>? addedTrackers,
    $core.Iterable<$fixnum.Int64>? deletedTrackerIds,
  }) {
    final _result = create();
    if (addedTrackers != null) {
      _result.addedTrackers.addAll(addedTrackers);
    }
    if (deletedTrackerIds != null) {
      _result.deletedTrackerIds.addAll(deletedTrackerIds);
    }
    return _result;
  }
  factory TrackValidatorsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TrackValidatorsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TrackValidatorsResponse clone() => TrackValidatorsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TrackValidatorsResponse copyWith(void Function(TrackValidatorsResponse) updates) => super.copyWith((message) => updates(message as TrackValidatorsResponse)) as TrackValidatorsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TrackValidatorsResponse create() => TrackValidatorsResponse._();
  TrackValidatorsResponse createEmptyInstance() => create();
  static $pb.PbList<TrackValidatorsResponse> createRepeated() => $pb.PbList<TrackValidatorsResponse>();
  @$core.pragma('dart2js:noInline')
  static TrackValidatorsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TrackValidatorsResponse>(create);
  static TrackValidatorsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Tracker> get addedTrackers => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$fixnum.Int64> get deletedTrackerIds => $_getList(1);
}

