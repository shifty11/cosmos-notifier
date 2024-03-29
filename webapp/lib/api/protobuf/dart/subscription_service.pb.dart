///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'subscription_service.pbenum.dart';

export 'subscription_service.pbenum.dart';

class SubscriptionStats extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SubscriptionStats', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
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
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Subscription', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOB(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isSubscribed')
    ..aOB(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isEnabled')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'thumbnailUrl')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractAddress')
    ..aOM<SubscriptionStats>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'stats', subBuilder: SubscriptionStats.create)
    ..hasRequiredFields = false
  ;

  Subscription._() : super();
  factory Subscription({
    $fixnum.Int64? id,
    $core.String? name,
    $core.bool? isSubscribed,
    $core.bool? isEnabled,
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
    if (isEnabled != null) {
      _result.isEnabled = isEnabled;
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
  $core.bool get isEnabled => $_getBF(3);
  @$pb.TagNumber(4)
  set isEnabled($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasIsEnabled() => $_has(3);
  @$pb.TagNumber(4)
  void clearIsEnabled() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get thumbnailUrl => $_getSZ(4);
  @$pb.TagNumber(5)
  set thumbnailUrl($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasThumbnailUrl() => $_has(4);
  @$pb.TagNumber(5)
  void clearThumbnailUrl() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get contractAddress => $_getSZ(5);
  @$pb.TagNumber(6)
  set contractAddress($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasContractAddress() => $_has(5);
  @$pb.TagNumber(6)
  void clearContractAddress() => clearField(6);

  @$pb.TagNumber(7)
  SubscriptionStats get stats => $_getN(6);
  @$pb.TagNumber(7)
  set stats(SubscriptionStats v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasStats() => $_has(6);
  @$pb.TagNumber(7)
  void clearStats() => clearField(7);
  @$pb.TagNumber(7)
  SubscriptionStats ensureStats() => $_ensure(6);
}

class ChatRoom extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ChatRoom', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..e<ChatRoom_Type>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: ChatRoom_Type.TYPE_UNSPECIFIED, valueOf: ChatRoom_Type.valueOf, enumValues: ChatRoom_Type.values)
    ..pc<Subscription>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subscriptions', $pb.PbFieldType.PM, subBuilder: Subscription.create)
    ..hasRequiredFields = false
  ;

  ChatRoom._() : super();
  factory ChatRoom({
    $fixnum.Int64? id,
    $core.String? name,
    ChatRoom_Type? type,
    $core.Iterable<Subscription>? subscriptions,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (name != null) {
      _result.name = name;
    }
    if (type != null) {
      _result.type = type;
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
  ChatRoom_Type get type => $_getN(2);
  @$pb.TagNumber(3)
  set type(ChatRoom_Type v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasType() => $_has(2);
  @$pb.TagNumber(3)
  void clearType() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<Subscription> get subscriptions => $_getList(3);
}

class ListSubscriptionsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListSubscriptionsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..pc<ChatRoom>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chainChatRooms', $pb.PbFieldType.PM, subBuilder: ChatRoom.create)
    ..pc<ChatRoom>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractChatRooms', $pb.PbFieldType.PM, subBuilder: ChatRoom.create)
    ..hasRequiredFields = false
  ;

  ListSubscriptionsResponse._() : super();
  factory ListSubscriptionsResponse({
    $core.Iterable<ChatRoom>? chainChatRooms,
    $core.Iterable<ChatRoom>? contractChatRooms,
  }) {
    final _result = create();
    if (chainChatRooms != null) {
      _result.chainChatRooms.addAll(chainChatRooms);
    }
    if (contractChatRooms != null) {
      _result.contractChatRooms.addAll(contractChatRooms);
    }
    return _result;
  }
  factory ListSubscriptionsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListSubscriptionsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListSubscriptionsResponse clone() => ListSubscriptionsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListSubscriptionsResponse copyWith(void Function(ListSubscriptionsResponse) updates) => super.copyWith((message) => updates(message as ListSubscriptionsResponse)) as ListSubscriptionsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListSubscriptionsResponse create() => ListSubscriptionsResponse._();
  ListSubscriptionsResponse createEmptyInstance() => create();
  static $pb.PbList<ListSubscriptionsResponse> createRepeated() => $pb.PbList<ListSubscriptionsResponse>();
  @$core.pragma('dart2js:noInline')
  static ListSubscriptionsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListSubscriptionsResponse>(create);
  static ListSubscriptionsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<ChatRoom> get chainChatRooms => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<ChatRoom> get contractChatRooms => $_getList(1);
}

class ToggleChainSubscriptionRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ToggleChainSubscriptionRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRoomId', protoName: 'chatRoomId')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chainId', protoName: 'chainId')
    ..hasRequiredFields = false
  ;

  ToggleChainSubscriptionRequest._() : super();
  factory ToggleChainSubscriptionRequest({
    $fixnum.Int64? chatRoomId,
    $fixnum.Int64? chainId,
  }) {
    final _result = create();
    if (chatRoomId != null) {
      _result.chatRoomId = chatRoomId;
    }
    if (chainId != null) {
      _result.chainId = chainId;
    }
    return _result;
  }
  factory ToggleChainSubscriptionRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ToggleChainSubscriptionRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ToggleChainSubscriptionRequest clone() => ToggleChainSubscriptionRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ToggleChainSubscriptionRequest copyWith(void Function(ToggleChainSubscriptionRequest) updates) => super.copyWith((message) => updates(message as ToggleChainSubscriptionRequest)) as ToggleChainSubscriptionRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ToggleChainSubscriptionRequest create() => ToggleChainSubscriptionRequest._();
  ToggleChainSubscriptionRequest createEmptyInstance() => create();
  static $pb.PbList<ToggleChainSubscriptionRequest> createRepeated() => $pb.PbList<ToggleChainSubscriptionRequest>();
  @$core.pragma('dart2js:noInline')
  static ToggleChainSubscriptionRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ToggleChainSubscriptionRequest>(create);
  static ToggleChainSubscriptionRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get chatRoomId => $_getI64(0);
  @$pb.TagNumber(1)
  set chatRoomId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChatRoomId() => $_has(0);
  @$pb.TagNumber(1)
  void clearChatRoomId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get chainId => $_getI64(1);
  @$pb.TagNumber(2)
  set chainId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasChainId() => $_has(1);
  @$pb.TagNumber(2)
  void clearChainId() => clearField(2);
}

class ToggleContractSubscriptionRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ToggleContractSubscriptionRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatRoomId', protoName: 'chatRoomId')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractId', protoName: 'contractId')
    ..hasRequiredFields = false
  ;

  ToggleContractSubscriptionRequest._() : super();
  factory ToggleContractSubscriptionRequest({
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
  factory ToggleContractSubscriptionRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ToggleContractSubscriptionRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ToggleContractSubscriptionRequest clone() => ToggleContractSubscriptionRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ToggleContractSubscriptionRequest copyWith(void Function(ToggleContractSubscriptionRequest) updates) => super.copyWith((message) => updates(message as ToggleContractSubscriptionRequest)) as ToggleContractSubscriptionRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ToggleContractSubscriptionRequest create() => ToggleContractSubscriptionRequest._();
  ToggleContractSubscriptionRequest createEmptyInstance() => create();
  static $pb.PbList<ToggleContractSubscriptionRequest> createRepeated() => $pb.PbList<ToggleContractSubscriptionRequest>();
  @$core.pragma('dart2js:noInline')
  static ToggleContractSubscriptionRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ToggleContractSubscriptionRequest>(create);
  static ToggleContractSubscriptionRequest? _defaultInstance;

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
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ToggleSubscriptionResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
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
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddDaoRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contractAddress', protoName: 'contractAddress')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'customQuery', protoName: 'customQuery')
    ..hasRequiredFields = false
  ;

  AddDaoRequest._() : super();
  factory AddDaoRequest({
    $core.String? contractAddress,
    $core.String? customQuery,
  }) {
    final _result = create();
    if (contractAddress != null) {
      _result.contractAddress = contractAddress;
    }
    if (customQuery != null) {
      _result.customQuery = customQuery;
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

  @$pb.TagNumber(2)
  $core.String get customQuery => $_getSZ(1);
  @$pb.TagNumber(2)
  set customQuery($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCustomQuery() => $_has(1);
  @$pb.TagNumber(2)
  void clearCustomQuery() => clearField(2);
}

class AddDaoResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AddDaoResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..e<AddDaoResponse_Status>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: AddDaoResponse_Status.STATUS_UNSPECIFIED, valueOf: AddDaoResponse_Status.valueOf, enumValues: AddDaoResponse_Status.values)
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
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteDaoRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
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

class EnableChainRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'EnableChainRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmos_notifier_grpc'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chainId', protoName: 'chainId')
    ..aOB(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'isEnabled', protoName: 'isEnabled')
    ..hasRequiredFields = false
  ;

  EnableChainRequest._() : super();
  factory EnableChainRequest({
    $fixnum.Int64? chainId,
    $core.bool? isEnabled,
  }) {
    final _result = create();
    if (chainId != null) {
      _result.chainId = chainId;
    }
    if (isEnabled != null) {
      _result.isEnabled = isEnabled;
    }
    return _result;
  }
  factory EnableChainRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory EnableChainRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  EnableChainRequest clone() => EnableChainRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  EnableChainRequest copyWith(void Function(EnableChainRequest) updates) => super.copyWith((message) => updates(message as EnableChainRequest)) as EnableChainRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static EnableChainRequest create() => EnableChainRequest._();
  EnableChainRequest createEmptyInstance() => create();
  static $pb.PbList<EnableChainRequest> createRepeated() => $pb.PbList<EnableChainRequest>();
  @$core.pragma('dart2js:noInline')
  static EnableChainRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<EnableChainRequest>(create);
  static EnableChainRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get chainId => $_getI64(0);
  @$pb.TagNumber(1)
  set chainId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChainId() => $_has(0);
  @$pb.TagNumber(1)
  void clearChainId() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get isEnabled => $_getBF(1);
  @$pb.TagNumber(2)
  set isEnabled($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasIsEnabled() => $_has(1);
  @$pb.TagNumber(2)
  void clearIsEnabled() => clearField(2);
}

