// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'tracker_row.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$TrackerRow {
  Int64 get id => throw _privateConstructorUsedError;
  String get address => throw _privateConstructorUsedError;
  pb.Duration get notificationInterval => throw _privateConstructorUsedError;
  TrackerChatRoom? get chatRoom => throw _privateConstructorUsedError;
  DateTime? get updatedAt => throw _privateConstructorUsedError;
  String get validatorMoniker => throw _privateConstructorUsedError;
  bool get isAddressValid => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $TrackerRowCopyWith<TrackerRow> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TrackerRowCopyWith<$Res> {
  factory $TrackerRowCopyWith(
          TrackerRow value, $Res Function(TrackerRow) then) =
      _$TrackerRowCopyWithImpl<$Res, TrackerRow>;
  @useResult
  $Res call(
      {Int64 id,
      String address,
      pb.Duration notificationInterval,
      TrackerChatRoom? chatRoom,
      DateTime? updatedAt,
      String validatorMoniker,
      bool isAddressValid});
}

/// @nodoc
class _$TrackerRowCopyWithImpl<$Res, $Val extends TrackerRow>
    implements $TrackerRowCopyWith<$Res> {
  _$TrackerRowCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? address = null,
    Object? notificationInterval = null,
    Object? chatRoom = freezed,
    Object? updatedAt = freezed,
    Object? validatorMoniker = null,
    Object? isAddressValid = null,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as Int64,
      address: null == address
          ? _value.address
          : address // ignore: cast_nullable_to_non_nullable
              as String,
      notificationInterval: null == notificationInterval
          ? _value.notificationInterval
          : notificationInterval // ignore: cast_nullable_to_non_nullable
              as pb.Duration,
      chatRoom: freezed == chatRoom
          ? _value.chatRoom
          : chatRoom // ignore: cast_nullable_to_non_nullable
              as TrackerChatRoom?,
      updatedAt: freezed == updatedAt
          ? _value.updatedAt
          : updatedAt // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      validatorMoniker: null == validatorMoniker
          ? _value.validatorMoniker
          : validatorMoniker // ignore: cast_nullable_to_non_nullable
              as String,
      isAddressValid: null == isAddressValid
          ? _value.isAddressValid
          : isAddressValid // ignore: cast_nullable_to_non_nullable
              as bool,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_TrackerRowCopyWith<$Res>
    implements $TrackerRowCopyWith<$Res> {
  factory _$$_TrackerRowCopyWith(
          _$_TrackerRow value, $Res Function(_$_TrackerRow) then) =
      __$$_TrackerRowCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {Int64 id,
      String address,
      pb.Duration notificationInterval,
      TrackerChatRoom? chatRoom,
      DateTime? updatedAt,
      String validatorMoniker,
      bool isAddressValid});
}

/// @nodoc
class __$$_TrackerRowCopyWithImpl<$Res>
    extends _$TrackerRowCopyWithImpl<$Res, _$_TrackerRow>
    implements _$$_TrackerRowCopyWith<$Res> {
  __$$_TrackerRowCopyWithImpl(
      _$_TrackerRow _value, $Res Function(_$_TrackerRow) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? address = null,
    Object? notificationInterval = null,
    Object? chatRoom = freezed,
    Object? updatedAt = freezed,
    Object? validatorMoniker = null,
    Object? isAddressValid = null,
  }) {
    return _then(_$_TrackerRow(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as Int64,
      address: null == address
          ? _value.address
          : address // ignore: cast_nullable_to_non_nullable
              as String,
      notificationInterval: null == notificationInterval
          ? _value.notificationInterval
          : notificationInterval // ignore: cast_nullable_to_non_nullable
              as pb.Duration,
      chatRoom: freezed == chatRoom
          ? _value.chatRoom
          : chatRoom // ignore: cast_nullable_to_non_nullable
              as TrackerChatRoom?,
      updatedAt: freezed == updatedAt
          ? _value.updatedAt
          : updatedAt // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      validatorMoniker: null == validatorMoniker
          ? _value.validatorMoniker
          : validatorMoniker // ignore: cast_nullable_to_non_nullable
              as String,
      isAddressValid: null == isAddressValid
          ? _value.isAddressValid
          : isAddressValid // ignore: cast_nullable_to_non_nullable
              as bool,
    ));
  }
}

/// @nodoc

class _$_TrackerRow extends _TrackerRow {
  const _$_TrackerRow(
      {required this.id,
      required this.address,
      required this.notificationInterval,
      required this.chatRoom,
      required this.updatedAt,
      required this.validatorMoniker,
      this.isAddressValid = true})
      : super._();

  @override
  final Int64 id;
  @override
  final String address;
  @override
  final pb.Duration notificationInterval;
  @override
  final TrackerChatRoom? chatRoom;
  @override
  final DateTime? updatedAt;
  @override
  final String validatorMoniker;
  @override
  @JsonKey()
  final bool isAddressValid;

  @override
  String toString() {
    return 'TrackerRow(id: $id, address: $address, notificationInterval: $notificationInterval, chatRoom: $chatRoom, updatedAt: $updatedAt, validatorMoniker: $validatorMoniker, isAddressValid: $isAddressValid)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_TrackerRow &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.address, address) || other.address == address) &&
            (identical(other.notificationInterval, notificationInterval) ||
                other.notificationInterval == notificationInterval) &&
            (identical(other.chatRoom, chatRoom) ||
                other.chatRoom == chatRoom) &&
            (identical(other.updatedAt, updatedAt) ||
                other.updatedAt == updatedAt) &&
            (identical(other.validatorMoniker, validatorMoniker) ||
                other.validatorMoniker == validatorMoniker) &&
            (identical(other.isAddressValid, isAddressValid) ||
                other.isAddressValid == isAddressValid));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType,
      id,
      address,
      notificationInterval,
      chatRoom,
      updatedAt,
      validatorMoniker,
      isAddressValid);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_TrackerRowCopyWith<_$_TrackerRow> get copyWith =>
      __$$_TrackerRowCopyWithImpl<_$_TrackerRow>(this, _$identity);
}

abstract class _TrackerRow extends TrackerRow {
  const factory _TrackerRow(
      {required final Int64 id,
      required final String address,
      required final pb.Duration notificationInterval,
      required final TrackerChatRoom? chatRoom,
      required final DateTime? updatedAt,
      required final String validatorMoniker,
      final bool isAddressValid}) = _$_TrackerRow;
  const _TrackerRow._() : super._();

  @override
  Int64 get id;
  @override
  String get address;
  @override
  pb.Duration get notificationInterval;
  @override
  TrackerChatRoom? get chatRoom;
  @override
  DateTime? get updatedAt;
  @override
  String get validatorMoniker;
  @override
  bool get isAddressValid;
  @override
  @JsonKey(ignore: true)
  _$$_TrackerRowCopyWith<_$_TrackerRow> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
mixin _$TrackerSortState {
  bool get isAscending => throw _privateConstructorUsedError;
  TrackerSortType get sortType => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $TrackerSortStateCopyWith<TrackerSortState> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TrackerSortStateCopyWith<$Res> {
  factory $TrackerSortStateCopyWith(
          TrackerSortState value, $Res Function(TrackerSortState) then) =
      _$TrackerSortStateCopyWithImpl<$Res, TrackerSortState>;
  @useResult
  $Res call({bool isAscending, TrackerSortType sortType});
}

/// @nodoc
class _$TrackerSortStateCopyWithImpl<$Res, $Val extends TrackerSortState>
    implements $TrackerSortStateCopyWith<$Res> {
  _$TrackerSortStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isAscending = null,
    Object? sortType = null,
  }) {
    return _then(_value.copyWith(
      isAscending: null == isAscending
          ? _value.isAscending
          : isAscending // ignore: cast_nullable_to_non_nullable
              as bool,
      sortType: null == sortType
          ? _value.sortType
          : sortType // ignore: cast_nullable_to_non_nullable
              as TrackerSortType,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_TrackerSortStateCopyWith<$Res>
    implements $TrackerSortStateCopyWith<$Res> {
  factory _$$_TrackerSortStateCopyWith(
          _$_TrackerSortState value, $Res Function(_$_TrackerSortState) then) =
      __$$_TrackerSortStateCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({bool isAscending, TrackerSortType sortType});
}

/// @nodoc
class __$$_TrackerSortStateCopyWithImpl<$Res>
    extends _$TrackerSortStateCopyWithImpl<$Res, _$_TrackerSortState>
    implements _$$_TrackerSortStateCopyWith<$Res> {
  __$$_TrackerSortStateCopyWithImpl(
      _$_TrackerSortState _value, $Res Function(_$_TrackerSortState) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isAscending = null,
    Object? sortType = null,
  }) {
    return _then(_$_TrackerSortState(
      isAscending: null == isAscending
          ? _value.isAscending
          : isAscending // ignore: cast_nullable_to_non_nullable
              as bool,
      sortType: null == sortType
          ? _value.sortType
          : sortType // ignore: cast_nullable_to_non_nullable
              as TrackerSortType,
    ));
  }
}

/// @nodoc

class _$_TrackerSortState extends _TrackerSortState {
  const _$_TrackerSortState({required this.isAscending, required this.sortType})
      : super._();

  @override
  final bool isAscending;
  @override
  final TrackerSortType sortType;

  @override
  String toString() {
    return 'TrackerSortState(isAscending: $isAscending, sortType: $sortType)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_TrackerSortState &&
            (identical(other.isAscending, isAscending) ||
                other.isAscending == isAscending) &&
            (identical(other.sortType, sortType) ||
                other.sortType == sortType));
  }

  @override
  int get hashCode => Object.hash(runtimeType, isAscending, sortType);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_TrackerSortStateCopyWith<_$_TrackerSortState> get copyWith =>
      __$$_TrackerSortStateCopyWithImpl<_$_TrackerSortState>(this, _$identity);
}

abstract class _TrackerSortState extends TrackerSortState {
  const factory _TrackerSortState(
      {required final bool isAscending,
      required final TrackerSortType sortType}) = _$_TrackerSortState;
  const _TrackerSortState._() : super._();

  @override
  bool get isAscending;
  @override
  TrackerSortType get sortType;
  @override
  @JsonKey(ignore: true)
  _$$_TrackerSortStateCopyWith<_$_TrackerSortState> get copyWith =>
      throw _privateConstructorUsedError;
}
