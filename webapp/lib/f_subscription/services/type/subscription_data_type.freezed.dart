// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'subscription_data_type.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$SubscriptionData {
  Subscription get subscription => throw _privateConstructorUsedError;
  int get index => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $SubscriptionDataCopyWith<SubscriptionData> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $SubscriptionDataCopyWith<$Res> {
  factory $SubscriptionDataCopyWith(
          SubscriptionData value, $Res Function(SubscriptionData) then) =
      _$SubscriptionDataCopyWithImpl<$Res, SubscriptionData>;
  @useResult
  $Res call({Subscription subscription, int index});
}

/// @nodoc
class _$SubscriptionDataCopyWithImpl<$Res, $Val extends SubscriptionData>
    implements $SubscriptionDataCopyWith<$Res> {
  _$SubscriptionDataCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? subscription = null,
    Object? index = null,
  }) {
    return _then(_value.copyWith(
      subscription: null == subscription
          ? _value.subscription
          : subscription // ignore: cast_nullable_to_non_nullable
              as Subscription,
      index: null == index
          ? _value.index
          : index // ignore: cast_nullable_to_non_nullable
              as int,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_SubscriptionDataCopyWith<$Res>
    implements $SubscriptionDataCopyWith<$Res> {
  factory _$$_SubscriptionDataCopyWith(
          _$_SubscriptionData value, $Res Function(_$_SubscriptionData) then) =
      __$$_SubscriptionDataCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({Subscription subscription, int index});
}

/// @nodoc
class __$$_SubscriptionDataCopyWithImpl<$Res>
    extends _$SubscriptionDataCopyWithImpl<$Res, _$_SubscriptionData>
    implements _$$_SubscriptionDataCopyWith<$Res> {
  __$$_SubscriptionDataCopyWithImpl(
      _$_SubscriptionData _value, $Res Function(_$_SubscriptionData) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? subscription = null,
    Object? index = null,
  }) {
    return _then(_$_SubscriptionData(
      null == subscription
          ? _value.subscription
          : subscription // ignore: cast_nullable_to_non_nullable
              as Subscription,
      null == index
          ? _value.index
          : index // ignore: cast_nullable_to_non_nullable
              as int,
    ));
  }
}

/// @nodoc

class _$_SubscriptionData implements _SubscriptionData {
  const _$_SubscriptionData(this.subscription, this.index);

  @override
  final Subscription subscription;
  @override
  final int index;

  @override
  String toString() {
    return 'SubscriptionData(subscription: $subscription, index: $index)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_SubscriptionData &&
            (identical(other.subscription, subscription) ||
                other.subscription == subscription) &&
            (identical(other.index, index) || other.index == index));
  }

  @override
  int get hashCode => Object.hash(runtimeType, subscription, index);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_SubscriptionDataCopyWith<_$_SubscriptionData> get copyWith =>
      __$$_SubscriptionDataCopyWithImpl<_$_SubscriptionData>(this, _$identity);
}

abstract class _SubscriptionData implements SubscriptionData {
  const factory _SubscriptionData(
      final Subscription subscription, final int index) = _$_SubscriptionData;

  @override
  Subscription get subscription;
  @override
  int get index;
  @override
  @JsonKey(ignore: true)
  _$$_SubscriptionDataCopyWith<_$_SubscriptionData> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
mixin _$ChatroomData {
  fixnum.Int64 get chatRoomId => throw _privateConstructorUsedError;
  String get username => throw _privateConstructorUsedError;
  List<Subscription> get subscriptions => throw _privateConstructorUsedError;
  List<SubscriptionData> get filtered => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $ChatroomDataCopyWith<ChatroomData> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $ChatroomDataCopyWith<$Res> {
  factory $ChatroomDataCopyWith(
          ChatroomData value, $Res Function(ChatroomData) then) =
      _$ChatroomDataCopyWithImpl<$Res, ChatroomData>;
  @useResult
  $Res call(
      {fixnum.Int64 chatRoomId,
      String username,
      List<Subscription> subscriptions,
      List<SubscriptionData> filtered});
}

/// @nodoc
class _$ChatroomDataCopyWithImpl<$Res, $Val extends ChatroomData>
    implements $ChatroomDataCopyWith<$Res> {
  _$ChatroomDataCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? chatRoomId = null,
    Object? username = null,
    Object? subscriptions = null,
    Object? filtered = null,
  }) {
    return _then(_value.copyWith(
      chatRoomId: null == chatRoomId
          ? _value.chatRoomId
          : chatRoomId // ignore: cast_nullable_to_non_nullable
              as fixnum.Int64,
      username: null == username
          ? _value.username
          : username // ignore: cast_nullable_to_non_nullable
              as String,
      subscriptions: null == subscriptions
          ? _value.subscriptions
          : subscriptions // ignore: cast_nullable_to_non_nullable
              as List<Subscription>,
      filtered: null == filtered
          ? _value.filtered
          : filtered // ignore: cast_nullable_to_non_nullable
              as List<SubscriptionData>,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_ChatroomDataCopyWith<$Res>
    implements $ChatroomDataCopyWith<$Res> {
  factory _$$_ChatroomDataCopyWith(
          _$_ChatroomData value, $Res Function(_$_ChatroomData) then) =
      __$$_ChatroomDataCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {fixnum.Int64 chatRoomId,
      String username,
      List<Subscription> subscriptions,
      List<SubscriptionData> filtered});
}

/// @nodoc
class __$$_ChatroomDataCopyWithImpl<$Res>
    extends _$ChatroomDataCopyWithImpl<$Res, _$_ChatroomData>
    implements _$$_ChatroomDataCopyWith<$Res> {
  __$$_ChatroomDataCopyWithImpl(
      _$_ChatroomData _value, $Res Function(_$_ChatroomData) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? chatRoomId = null,
    Object? username = null,
    Object? subscriptions = null,
    Object? filtered = null,
  }) {
    return _then(_$_ChatroomData(
      null == chatRoomId
          ? _value.chatRoomId
          : chatRoomId // ignore: cast_nullable_to_non_nullable
              as fixnum.Int64,
      null == username
          ? _value.username
          : username // ignore: cast_nullable_to_non_nullable
              as String,
      null == subscriptions
          ? _value._subscriptions
          : subscriptions // ignore: cast_nullable_to_non_nullable
              as List<Subscription>,
      null == filtered
          ? _value._filtered
          : filtered // ignore: cast_nullable_to_non_nullable
              as List<SubscriptionData>,
    ));
  }
}

/// @nodoc

class _$_ChatroomData extends _ChatroomData {
  const _$_ChatroomData(
      this.chatRoomId,
      this.username,
      final List<Subscription> subscriptions,
      final List<SubscriptionData> filtered)
      : _subscriptions = subscriptions,
        _filtered = filtered,
        super._();

  @override
  final fixnum.Int64 chatRoomId;
  @override
  final String username;
  final List<Subscription> _subscriptions;
  @override
  List<Subscription> get subscriptions {
    if (_subscriptions is EqualUnmodifiableListView) return _subscriptions;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_subscriptions);
  }

  final List<SubscriptionData> _filtered;
  @override
  List<SubscriptionData> get filtered {
    if (_filtered is EqualUnmodifiableListView) return _filtered;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_filtered);
  }

  @override
  String toString() {
    return 'ChatroomData(chatRoomId: $chatRoomId, username: $username, subscriptions: $subscriptions, filtered: $filtered)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_ChatroomData &&
            (identical(other.chatRoomId, chatRoomId) ||
                other.chatRoomId == chatRoomId) &&
            (identical(other.username, username) ||
                other.username == username) &&
            const DeepCollectionEquality()
                .equals(other._subscriptions, _subscriptions) &&
            const DeepCollectionEquality().equals(other._filtered, _filtered));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType,
      chatRoomId,
      username,
      const DeepCollectionEquality().hash(_subscriptions),
      const DeepCollectionEquality().hash(_filtered));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_ChatroomDataCopyWith<_$_ChatroomData> get copyWith =>
      __$$_ChatroomDataCopyWithImpl<_$_ChatroomData>(this, _$identity);
}

abstract class _ChatroomData extends ChatroomData {
  const factory _ChatroomData(
      final fixnum.Int64 chatRoomId,
      final String username,
      final List<Subscription> subscriptions,
      final List<SubscriptionData> filtered) = _$_ChatroomData;
  const _ChatroomData._() : super._();

  @override
  fixnum.Int64 get chatRoomId;
  @override
  String get username;
  @override
  List<Subscription> get subscriptions;
  @override
  List<SubscriptionData> get filtered;
  @override
  @JsonKey(ignore: true)
  _$$_ChatroomDataCopyWith<_$_ChatroomData> get copyWith =>
      throw _privateConstructorUsedError;
}
