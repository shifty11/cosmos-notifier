// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

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
      _$SubscriptionDataCopyWithImpl<$Res>;
  $Res call({Subscription subscription, int index});
}

/// @nodoc
class _$SubscriptionDataCopyWithImpl<$Res>
    implements $SubscriptionDataCopyWith<$Res> {
  _$SubscriptionDataCopyWithImpl(this._value, this._then);

  final SubscriptionData _value;
  // ignore: unused_field
  final $Res Function(SubscriptionData) _then;

  @override
  $Res call({
    Object? subscription = freezed,
    Object? index = freezed,
  }) {
    return _then(_value.copyWith(
      subscription: subscription == freezed
          ? _value.subscription
          : subscription // ignore: cast_nullable_to_non_nullable
              as Subscription,
      index: index == freezed
          ? _value.index
          : index // ignore: cast_nullable_to_non_nullable
              as int,
    ));
  }
}

/// @nodoc
abstract class _$$_SubscriptionDataCopyWith<$Res>
    implements $SubscriptionDataCopyWith<$Res> {
  factory _$$_SubscriptionDataCopyWith(
          _$_SubscriptionData value, $Res Function(_$_SubscriptionData) then) =
      __$$_SubscriptionDataCopyWithImpl<$Res>;
  @override
  $Res call({Subscription subscription, int index});
}

/// @nodoc
class __$$_SubscriptionDataCopyWithImpl<$Res>
    extends _$SubscriptionDataCopyWithImpl<$Res>
    implements _$$_SubscriptionDataCopyWith<$Res> {
  __$$_SubscriptionDataCopyWithImpl(
      _$_SubscriptionData _value, $Res Function(_$_SubscriptionData) _then)
      : super(_value, (v) => _then(v as _$_SubscriptionData));

  @override
  _$_SubscriptionData get _value => super._value as _$_SubscriptionData;

  @override
  $Res call({
    Object? subscription = freezed,
    Object? index = freezed,
  }) {
    return _then(_$_SubscriptionData(
      subscription == freezed
          ? _value.subscription
          : subscription // ignore: cast_nullable_to_non_nullable
              as Subscription,
      index == freezed
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
            const DeepCollectionEquality()
                .equals(other.subscription, subscription) &&
            const DeepCollectionEquality().equals(other.index, index));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(subscription),
      const DeepCollectionEquality().hash(index));

  @JsonKey(ignore: true)
  @override
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
      _$ChatroomDataCopyWithImpl<$Res>;
  $Res call(
      {fixnum.Int64 chatRoomId,
      String username,
      List<Subscription> subscriptions,
      List<SubscriptionData> filtered});
}

/// @nodoc
class _$ChatroomDataCopyWithImpl<$Res> implements $ChatroomDataCopyWith<$Res> {
  _$ChatroomDataCopyWithImpl(this._value, this._then);

  final ChatroomData _value;
  // ignore: unused_field
  final $Res Function(ChatroomData) _then;

  @override
  $Res call({
    Object? chatRoomId = freezed,
    Object? username = freezed,
    Object? subscriptions = freezed,
    Object? filtered = freezed,
  }) {
    return _then(_value.copyWith(
      chatRoomId: chatRoomId == freezed
          ? _value.chatRoomId
          : chatRoomId // ignore: cast_nullable_to_non_nullable
              as fixnum.Int64,
      username: username == freezed
          ? _value.username
          : username // ignore: cast_nullable_to_non_nullable
              as String,
      subscriptions: subscriptions == freezed
          ? _value.subscriptions
          : subscriptions // ignore: cast_nullable_to_non_nullable
              as List<Subscription>,
      filtered: filtered == freezed
          ? _value.filtered
          : filtered // ignore: cast_nullable_to_non_nullable
              as List<SubscriptionData>,
    ));
  }
}

/// @nodoc
abstract class _$$_ChatroomDataCopyWith<$Res>
    implements $ChatroomDataCopyWith<$Res> {
  factory _$$_ChatroomDataCopyWith(
          _$_ChatroomData value, $Res Function(_$_ChatroomData) then) =
      __$$_ChatroomDataCopyWithImpl<$Res>;
  @override
  $Res call(
      {fixnum.Int64 chatRoomId,
      String username,
      List<Subscription> subscriptions,
      List<SubscriptionData> filtered});
}

/// @nodoc
class __$$_ChatroomDataCopyWithImpl<$Res>
    extends _$ChatroomDataCopyWithImpl<$Res>
    implements _$$_ChatroomDataCopyWith<$Res> {
  __$$_ChatroomDataCopyWithImpl(
      _$_ChatroomData _value, $Res Function(_$_ChatroomData) _then)
      : super(_value, (v) => _then(v as _$_ChatroomData));

  @override
  _$_ChatroomData get _value => super._value as _$_ChatroomData;

  @override
  $Res call({
    Object? chatRoomId = freezed,
    Object? username = freezed,
    Object? subscriptions = freezed,
    Object? filtered = freezed,
  }) {
    return _then(_$_ChatroomData(
      chatRoomId == freezed
          ? _value.chatRoomId
          : chatRoomId // ignore: cast_nullable_to_non_nullable
              as fixnum.Int64,
      username == freezed
          ? _value.username
          : username // ignore: cast_nullable_to_non_nullable
              as String,
      subscriptions == freezed
          ? _value._subscriptions
          : subscriptions // ignore: cast_nullable_to_non_nullable
              as List<Subscription>,
      filtered == freezed
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
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_subscriptions);
  }

  final List<SubscriptionData> _filtered;
  @override
  List<SubscriptionData> get filtered {
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
            const DeepCollectionEquality()
                .equals(other.chatRoomId, chatRoomId) &&
            const DeepCollectionEquality().equals(other.username, username) &&
            const DeepCollectionEquality()
                .equals(other._subscriptions, _subscriptions) &&
            const DeepCollectionEquality().equals(other._filtered, _filtered));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(chatRoomId),
      const DeepCollectionEquality().hash(username),
      const DeepCollectionEquality().hash(_subscriptions),
      const DeepCollectionEquality().hash(_filtered));

  @JsonKey(ignore: true)
  @override
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
