///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'subscription_service.pbenum.dart';

export 'subscription_service.pbenum.dart';

class SubscriptionStats extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SubscriptionStats', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'total', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'telegram', $pb.PbFieldType.O3)
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'discord', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  SubscriptionStats._() : super();
  factory SubscriptionStats({
    $core.int? total,
    $core.int? telegram,
    $core.int? discord,
  }) {
    final _result = create();
    if (total != null) {
      _result.total = total;
    }
    if (telegram != null) {
      _result.telegram = telegram;
    }
    if (discord != null) {
      _result.discord = discord;
    }
    return _result;
  }
  factory SubscriptionStats.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SubscriptionStats.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SubscriptionStats clone() => SubscriptionStats()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SubscriptionStats copyWith(void Function(SubscriptionStats) updates) => super.copyWith((message) => updates(message as SubscriptionStats)) as SubscriptionStats; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SubscriptionStats create() => SubscriptionStats._();
  SubscriptionStats createEmptyInstance() => create();
  static $pb.PbList<SubscriptionStats> createRepeated() => $pb.PbList<SubscriptionStats>();
  @$core.pragma('dart2js:noInline')
  static SubscriptionStats getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SubscriptionStats>(create);
  static SubscriptionStats? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get total => $_getIZ(0);
  @$pb.TagNumber(1)
  set total($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTotal() => $_has(0);
  @$pb.TagNumber(1)
  void clearTotal() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get telegram => $_getIZ(1);
  @$pb.TagNumber(2)
  set telegram($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTelegram() => $_has(1);
  @$pb.TagNumber(2)
  void clearTelegram() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get discord => $_getIZ(2);
  @$pb.TagNumber(3)
  set discord($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDiscord() => $_has(2);
  @$pb.TagNumber(3)
  void clearDiscord() => clearField(3);
}

class Subscription extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Subscription', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOB(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isSubscribed')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'thumbnailUrl')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractAddress')
    ..aOM<SubscriptionStats>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'stats', subBuilder: SubscriptionStats.create)
    ..hasRequiredFields = false
  ;

  Subscription._() : super();
  factory Subscription({
    $fixnum.Int64? id,
    $core.String? name,
    $core.bool? isSubscribed,
    $core.String? thumbnailUrl,
    $core.String? contractAddress,
    SubscriptionStats? stats,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (name != null) {
      _result.name = name;
    }
    if (isSubscribed != null) {
      _result.isSubscribed = isSubscribed;
    }
    if (thumbnailUrl != null) {
      _result.thumbnailUrl = thumbnailUrl;
    }
    if (contractAddress != null) {
      _result.contractAddress = contractAddress;
    }
    if (stats != null) {
      _result.stats = stats;
    }
    return _result;
  }
  factory Subscription.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Subscription.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Subscription clone() => Subscription()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Subscription copyWith(void Function(Subscription) updates) => super.copyWith((message) => updates(message as Subscription)) as Subscription; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Subscription create() => Subscription._();
  Subscription createEmptyInstance() => create();
  static $pb.PbList<Subscription> createRepeated() => $pb.PbList<Subscription>();
  @$core.pragma('dart2js:noInline')
  static Subscription getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Subscription>(create);
  static Subscription? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get isSubscribed => $_getBF(2);
  @$pb.TagNumber(3)
  set isSubscribed($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasIsSubscribed() => $_has(2);
  @$pb.TagNumber(3)
  void clearIsSubscribed() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get thumbnailUrl => $_getSZ(3);
  @$pb.TagNumber(4)
  set thumbnailUrl($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasThumbnailUrl() => $_has(3);
  @$pb.TagNumber(4)
  void clearThumbnailUrl() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get contractAddress => $_getSZ(4);
  @$pb.TagNumber(5)
  set contractAddress($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasContractAddress() => $_has(4);
  @$pb.TagNumber(5)
  void clearContractAddress() => clearField(5);

  @$pb.TagNumber(6)
  SubscriptionStats get stats => $_getN(5);
  @$pb.TagNumber(6)
  set stats(SubscriptionStats v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasStats() => $_has(5);
  @$pb.TagNumber(6)
  void clearStats() => clearField(6);
  @$pb.TagNumber(6)
  SubscriptionStats ensureStats() => $_ensure(5);
}

class ChatRoom extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ChatRoom', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..e<ChatRoom_Type>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TYPE', $pb.PbFieldType.OE, protoName: 'TYPE', defaultOrMaker: ChatRoom_Type.TELEGRAM, valueOf: ChatRoom_Type.valueOf, enumValues: ChatRoom_Type.values)
    ..pc<Subscription>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subscriptions', $pb.PbFieldType.PM, subBuilder: Subscription.create)
    ..hasRequiredFields = false
  ;

  ChatRoom._() : super();
  factory ChatRoom({
    $fixnum.Int64? id,
    $core.String? name,
    ChatRoom_Type? tYPE,
    $core.Iterable<Subscription>? subscriptions,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (name != null) {
      _result.name = name;
    }
    if (tYPE != null) {
      _result.tYPE = tYPE;
    }
    if (subscriptions != null) {
      _result.subscriptions.addAll(subscriptions);
    }
    return _result;
  }
  factory ChatRoom.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ChatRoom.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ChatRoom clone() => ChatRoom()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ChatRoom copyWith(void Function(ChatRoom) updates) => super.copyWith((message) => updates(message as ChatRoom)) as ChatRoom; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ChatRoom create() => ChatRoom._();
  ChatRoom createEmptyInstance() => create();
  static $pb.PbList<ChatRoom> createRepeated() => $pb.PbList<ChatRoom>();
  @$core.pragma('dart2js:noInline')
  static ChatRoom getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ChatRoom>(create);
  static ChatRoom? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  ChatRoom_Type get tYPE => $_getN(2);
  @$pb.TagNumber(3)
  set tYPE(ChatRoom_Type v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTYPE() => $_has(2);
  @$pb.TagNumber(3)
  void clearTYPE() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<Subscription> get subscriptions => $_getList(3);
}

class GetSubscriptionsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetSubscriptionsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..pc<ChatRoom>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRooms', $pb.PbFieldType.PM, subBuilder: ChatRoom.create)
    ..hasRequiredFields = false
  ;

  GetSubscriptionsResponse._() : super();
  factory GetSubscriptionsResponse({
    $core.Iterable<ChatRoom>? chatRooms,
  }) {
    final _result = create();
    if (chatRooms != null) {
      _result.chatRooms.addAll(chatRooms);
    }
    return _result;
  }
  factory GetSubscriptionsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetSubscriptionsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetSubscriptionsResponse clone() => GetSubscriptionsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetSubscriptionsResponse copyWith(void Function(GetSubscriptionsResponse) updates) => super.copyWith((message) => updates(message as GetSubscriptionsResponse)) as GetSubscriptionsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetSubscriptionsResponse create() => GetSubscriptionsResponse._();
  GetSubscriptionsResponse createEmptyInstance() => create();
  static $pb.PbList<GetSubscriptionsResponse> createRepeated() => $pb.PbList<GetSubscriptionsResponse>();
  @$core.pragma('dart2js:noInline')
  static GetSubscriptionsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetSubscriptionsResponse>(create);
  static GetSubscriptionsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<ChatRoom> get chatRooms => $_getList(0);
}

class ToggleSubscriptionRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ToggleSubscriptionRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRoomId', protoName: 'chatRoomId')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractId', protoName: 'contractId')
    ..hasRequiredFields = false
  ;

  ToggleSubscriptionRequest._() : super();
  factory ToggleSubscriptionRequest({
    $fixnum.Int64? chatRoomId,
    $fixnum.Int64? contractId,
  }) {
    final _result = create();
    if (chatRoomId != null) {
      _result.chatRoomId = chatRoomId;
    }
    if (contractId != null) {
      _result.contractId = contractId;
    }
    return _result;
  }
  factory ToggleSubscriptionRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ToggleSubscriptionRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ToggleSubscriptionRequest clone() => ToggleSubscriptionRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ToggleSubscriptionRequest copyWith(void Function(ToggleSubscriptionRequest) updates) => super.copyWith((message) => updates(message as ToggleSubscriptionRequest)) as ToggleSubscriptionRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ToggleSubscriptionRequest create() => ToggleSubscriptionRequest._();
  ToggleSubscriptionRequest createEmptyInstance() => create();
  static $pb.PbList<ToggleSubscriptionRequest> createRepeated() => $pb.PbList<ToggleSubscriptionRequest>();
  @$core.pragma('dart2js:noInline')
  static ToggleSubscriptionRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ToggleSubscriptionRequest>(create);
  static ToggleSubscriptionRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get chatRoomId => $_getI64(0);
  @$pb.TagNumber(1)
  set chatRoomId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChatRoomId() => $_has(0);
  @$pb.TagNumber(1)
  void clearChatRoomId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get contractId => $_getI64(1);
  @$pb.TagNumber(2)
  set contractId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasContractId() => $_has(1);
  @$pb.TagNumber(2)
  void clearContractId() => clearField(2);
}

class ToggleSubscriptionResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ToggleSubscriptionResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..aOB(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isSubscribed', protoName: 'isSubscribed')
    ..hasRequiredFields = false
  ;

  ToggleSubscriptionResponse._() : super();
  factory ToggleSubscriptionResponse({
    $core.bool? isSubscribed,
  }) {
    final _result = create();
    if (isSubscribed != null) {
      _result.isSubscribed = isSubscribed;
    }
    return _result;
  }
  factory ToggleSubscriptionResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ToggleSubscriptionResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ToggleSubscriptionResponse clone() => ToggleSubscriptionResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ToggleSubscriptionResponse copyWith(void Function(ToggleSubscriptionResponse) updates) => super.copyWith((message) => updates(message as ToggleSubscriptionResponse)) as ToggleSubscriptionResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ToggleSubscriptionResponse create() => ToggleSubscriptionResponse._();
  ToggleSubscriptionResponse createEmptyInstance() => create();
  static $pb.PbList<ToggleSubscriptionResponse> createRepeated() => $pb.PbList<ToggleSubscriptionResponse>();
  @$core.pragma('dart2js:noInline')
  static ToggleSubscriptionResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ToggleSubscriptionResponse>(create);
  static ToggleSubscriptionResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isSubscribed => $_getBF(0);
  @$pb.TagNumber(1)
  set isSubscribed($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsSubscribed() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsSubscribed() => clearField(1);
}

class AddDaoRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddDaoRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractAddress', protoName: 'contractAddress')
    ..hasRequiredFields = false
  ;

  AddDaoRequest._() : super();
  factory AddDaoRequest({
    $core.String? contractAddress,
  }) {
    final _result = create();
    if (contractAddress != null) {
      _result.contractAddress = contractAddress;
    }
    return _result;
  }
  factory AddDaoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddDaoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddDaoRequest clone() => AddDaoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddDaoRequest copyWith(void Function(AddDaoRequest) updates) => super.copyWith((message) => updates(message as AddDaoRequest)) as AddDaoRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AddDaoRequest create() => AddDaoRequest._();
  AddDaoRequest createEmptyInstance() => create();
  static $pb.PbList<AddDaoRequest> createRepeated() => $pb.PbList<AddDaoRequest>();
  @$core.pragma('dart2js:noInline')
  static AddDaoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddDaoRequest>(create);
  static AddDaoRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get contractAddress => $_getSZ(0);
  @$pb.TagNumber(1)
  set contractAddress($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasContractAddress() => $_has(0);
  @$pb.TagNumber(1)
  void clearContractAddress() => clearField(1);
}

class AddDaoResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddDaoResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..e<AddDaoResponse_Status>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: AddDaoResponse_Status.ADDED, valueOf: AddDaoResponse_Status.valueOf, enumValues: AddDaoResponse_Status.values)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractAddress', protoName: 'contractAddress')
    ..hasRequiredFields = false
  ;

  AddDaoResponse._() : super();
  factory AddDaoResponse({
    AddDaoResponse_Status? status,
    $core.String? name,
    $core.String? contractAddress,
  }) {
    final _result = create();
    if (status != null) {
      _result.status = status;
    }
    if (name != null) {
      _result.name = name;
    }
    if (contractAddress != null) {
      _result.contractAddress = contractAddress;
    }
    return _result;
  }
  factory AddDaoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddDaoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddDaoResponse clone() => AddDaoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddDaoResponse copyWith(void Function(AddDaoResponse) updates) => super.copyWith((message) => updates(message as AddDaoResponse)) as AddDaoResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AddDaoResponse create() => AddDaoResponse._();
  AddDaoResponse createEmptyInstance() => create();
  static $pb.PbList<AddDaoResponse> createRepeated() => $pb.PbList<AddDaoResponse>();
  @$core.pragma('dart2js:noInline')
  static AddDaoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddDaoResponse>(create);
  static AddDaoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  AddDaoResponse_Status get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(AddDaoResponse_Status v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get contractAddress => $_getSZ(2);
  @$pb.TagNumber(3)
  set contractAddress($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasContractAddress() => $_has(2);
  @$pb.TagNumber(3)
  void clearContractAddress() => clearField(3);
}

class DeleteDaoRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteDaoRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'daodao_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractId', protoName: 'contractId')
    ..hasRequiredFields = false
  ;

  DeleteDaoRequest._() : super();
  factory DeleteDaoRequest({
    $fixnum.Int64? contractId,
  }) {
    final _result = create();
    if (contractId != null) {
      _result.contractId = contractId;
    }
    return _result;
  }
  factory DeleteDaoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteDaoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteDaoRequest clone() => DeleteDaoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteDaoRequest copyWith(void Function(DeleteDaoRequest) updates) => super.copyWith((message) => updates(message as DeleteDaoRequest)) as DeleteDaoRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteDaoRequest create() => DeleteDaoRequest._();
  DeleteDaoRequest createEmptyInstance() => create();
  static $pb.PbList<DeleteDaoRequest> createRepeated() => $pb.PbList<DeleteDaoRequest>();
  @$core.pragma('dart2js:noInline')
  static DeleteDaoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteDaoRequest>(create);
  static DeleteDaoRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get contractId => $_getI64(0);
  @$pb.TagNumber(1)
  set contractId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasContractId() => $_has(0);
  @$pb.TagNumber(1)
  void clearContractId() => clearField(1);
}

