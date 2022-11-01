// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

part of 'subscription_list_state.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$SubscriptionListState {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() loading,
    required TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)
        data,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult Function()? loading,
    TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)?
        data,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? loading,
    TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)?
        data,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Loading value) loading,
    required TResult Function(Data value) data,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult Function(Loading value)? loading,
    TResult Function(Data value)? data,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Loading value)? loading,
    TResult Function(Data value)? data,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $SubscriptionListStateCopyWith<$Res> {
  factory $SubscriptionListStateCopyWith(SubscriptionListState value,
          $Res Function(SubscriptionListState) then) =
      _$SubscriptionListStateCopyWithImpl<$Res>;
}

/// @nodoc
class _$SubscriptionListStateCopyWithImpl<$Res>
    implements $SubscriptionListStateCopyWith<$Res> {
  _$SubscriptionListStateCopyWithImpl(this._value, this._then);

  final SubscriptionListState _value;
  // ignore: unused_field
  final $Res Function(SubscriptionListState) _then;
}

/// @nodoc
abstract class _$$LoadingCopyWith<$Res> {
  factory _$$LoadingCopyWith(_$Loading value, $Res Function(_$Loading) then) =
      __$$LoadingCopyWithImpl<$Res>;
}

/// @nodoc
class __$$LoadingCopyWithImpl<$Res>
    extends _$SubscriptionListStateCopyWithImpl<$Res>
    implements _$$LoadingCopyWith<$Res> {
  __$$LoadingCopyWithImpl(_$Loading _value, $Res Function(_$Loading) _then)
      : super(_value, (v) => _then(v as _$Loading));

  @override
  _$Loading get _value => super._value as _$Loading;
}

/// @nodoc

class _$Loading extends Loading {
  _$Loading() : super._();

  @override
  String toString() {
    return 'SubscriptionListState.loading()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$Loading);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() loading,
    required TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)
        data,
  }) {
    return loading();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult Function()? loading,
    TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)?
        data,
  }) {
    return loading?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? loading,
    TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)?
        data,
    required TResult orElse(),
  }) {
    if (loading != null) {
      return loading();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Loading value) loading,
    required TResult Function(Data value) data,
  }) {
    return loading(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult Function(Loading value)? loading,
    TResult Function(Data value)? data,
  }) {
    return loading?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Loading value)? loading,
    TResult Function(Data value)? data,
    required TResult orElse(),
  }) {
    if (loading != null) {
      return loading(this);
    }
    return orElse();
  }
}

abstract class Loading extends SubscriptionListState {
  factory Loading() = _$Loading;
  Loading._() : super._();
}

/// @nodoc
abstract class _$$DataCopyWith<$Res> {
  factory _$$DataCopyWith(_$Data value, $Res Function(_$Data) then) =
      __$$DataCopyWithImpl<$Res>;
  $Res call({List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms});
}

/// @nodoc
class __$$DataCopyWithImpl<$Res>
    extends _$SubscriptionListStateCopyWithImpl<$Res>
    implements _$$DataCopyWith<$Res> {
  __$$DataCopyWithImpl(_$Data _value, $Res Function(_$Data) _then)
      : super(_value, (v) => _then(v as _$Data));

  @override
  _$Data get _value => super._value as _$Data;

  @override
  $Res call({
    Object? chainChatRooms = freezed,
    Object? contractChatRooms = freezed,
  }) {
    return _then(_$Data(
      chainChatRooms: chainChatRooms == freezed
          ? _value._chainChatRooms
          : chainChatRooms // ignore: cast_nullable_to_non_nullable
              as List<ChatRoom>,
      contractChatRooms: contractChatRooms == freezed
          ? _value._contractChatRooms
          : contractChatRooms // ignore: cast_nullable_to_non_nullable
              as List<ChatRoom>,
    ));
  }
}

/// @nodoc

class _$Data extends Data {
  _$Data(
      {required final List<ChatRoom> chainChatRooms,
      required final List<ChatRoom> contractChatRooms})
      : _chainChatRooms = chainChatRooms,
        _contractChatRooms = contractChatRooms,
        super._();

  final List<ChatRoom> _chainChatRooms;
  @override
  List<ChatRoom> get chainChatRooms {
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_chainChatRooms);
  }

  final List<ChatRoom> _contractChatRooms;
  @override
  List<ChatRoom> get contractChatRooms {
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_contractChatRooms);
  }

  @override
  String toString() {
    return 'SubscriptionListState.data(chainChatRooms: $chainChatRooms, contractChatRooms: $contractChatRooms)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$Data &&
            const DeepCollectionEquality()
                .equals(other._chainChatRooms, _chainChatRooms) &&
            const DeepCollectionEquality()
                .equals(other._contractChatRooms, _contractChatRooms));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(_chainChatRooms),
      const DeepCollectionEquality().hash(_contractChatRooms));

  @JsonKey(ignore: true)
  @override
  _$$DataCopyWith<_$Data> get copyWith =>
      __$$DataCopyWithImpl<_$Data>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() loading,
    required TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)
        data,
  }) {
    return data(chainChatRooms, contractChatRooms);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult Function()? loading,
    TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)?
        data,
  }) {
    return data?.call(chainChatRooms, contractChatRooms);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? loading,
    TResult Function(
            List<ChatRoom> chainChatRooms, List<ChatRoom> contractChatRooms)?
        data,
    required TResult orElse(),
  }) {
    if (data != null) {
      return data(chainChatRooms, contractChatRooms);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Loading value) loading,
    required TResult Function(Data value) data,
  }) {
    return data(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult Function(Loading value)? loading,
    TResult Function(Data value)? data,
  }) {
    return data?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Loading value)? loading,
    TResult Function(Data value)? data,
    required TResult orElse(),
  }) {
    if (data != null) {
      return data(this);
    }
    return orElse();
  }
}

abstract class Data extends SubscriptionListState {
  factory Data(
      {required final List<ChatRoom> chainChatRooms,
      required final List<ChatRoom> contractChatRooms}) = _$Data;
  Data._() : super._();

  List<ChatRoom> get chainChatRooms;
  List<ChatRoom> get contractChatRooms;
  @JsonKey(ignore: true)
  _$$DataCopyWith<_$Data> get copyWith => throw _privateConstructorUsedError;
}
