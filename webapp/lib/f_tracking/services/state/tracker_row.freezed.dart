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
  Int64 get chatId => throw _privateConstructorUsedError;
  DateTime? get updatedAt => throw _privateConstructorUsedError;
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
      Int64 chatId,
      DateTime? updatedAt,
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
    Object? chatId = null,
    Object? updatedAt = freezed,
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
      chatId: null == chatId
          ? _value.chatId
          : chatId // ignore: cast_nullable_to_non_nullable
              as Int64,
      updatedAt: freezed == updatedAt
          ? _value.updatedAt
          : updatedAt // ignore: cast_nullable_to_non_nullable
              as DateTime?,
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
      Int64 chatId,
      DateTime? updatedAt,
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
    Object? chatId = null,
    Object? updatedAt = freezed,
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
      chatId: null == chatId
          ? _value.chatId
          : chatId // ignore: cast_nullable_to_non_nullable
              as Int64,
      updatedAt: freezed == updatedAt
          ? _value.updatedAt
          : updatedAt // ignore: cast_nullable_to_non_nullable
              as DateTime?,
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
      required this.chatId,
      required this.updatedAt,
      this.isAddressValid = true})
      : super._();

  @override
  final Int64 id;
  @override
  final String address;
  @override
  final pb.Duration notificationInterval;
  @override
  final Int64 chatId;
  @override
  final DateTime? updatedAt;
  @override
  @JsonKey()
  final bool isAddressValid;

  @override
  String toString() {
    return 'TrackerRow(id: $id, address: $address, notificationInterval: $notificationInterval, chatId: $chatId, updatedAt: $updatedAt, isAddressValid: $isAddressValid)';
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
            (identical(other.chatId, chatId) || other.chatId == chatId) &&
            (identical(other.updatedAt, updatedAt) ||
                other.updatedAt == updatedAt) &&
            (identical(other.isAddressValid, isAddressValid) ||
                other.isAddressValid == isAddressValid));
  }

  @override
  int get hashCode => Object.hash(runtimeType, id, address,
      notificationInterval, chatId, updatedAt, isAddressValid);

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
      required final Int64 chatId,
      required final DateTime? updatedAt,
      final bool isAddressValid}) = _$_TrackerRow;
  const _TrackerRow._() : super._();

  @override
  Int64 get id;
  @override
  String get address;
  @override
  pb.Duration get notificationInterval;
  @override
  Int64 get chatId;
  @override
  DateTime? get updatedAt;
  @override
  bool get isAddressValid;
  @override
  @JsonKey(ignore: true)
  _$$_TrackerRowCopyWith<_$_TrackerRow> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
mixin _$TrackerList {
  List<TrackerRow> get trackers => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $TrackerListCopyWith<TrackerList> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TrackerListCopyWith<$Res> {
  factory $TrackerListCopyWith(
          TrackerList value, $Res Function(TrackerList) then) =
      _$TrackerListCopyWithImpl<$Res, TrackerList>;
  @useResult
  $Res call({List<TrackerRow> trackers});
}

/// @nodoc
class _$TrackerListCopyWithImpl<$Res, $Val extends TrackerList>
    implements $TrackerListCopyWith<$Res> {
  _$TrackerListCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? trackers = null,
  }) {
    return _then(_value.copyWith(
      trackers: null == trackers
          ? _value.trackers
          : trackers // ignore: cast_nullable_to_non_nullable
              as List<TrackerRow>,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_TrackerListCopyWith<$Res>
    implements $TrackerListCopyWith<$Res> {
  factory _$$_TrackerListCopyWith(
          _$_TrackerList value, $Res Function(_$_TrackerList) then) =
      __$$_TrackerListCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({List<TrackerRow> trackers});
}

/// @nodoc
class __$$_TrackerListCopyWithImpl<$Res>
    extends _$TrackerListCopyWithImpl<$Res, _$_TrackerList>
    implements _$$_TrackerListCopyWith<$Res> {
  __$$_TrackerListCopyWithImpl(
      _$_TrackerList _value, $Res Function(_$_TrackerList) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? trackers = null,
  }) {
    return _then(_$_TrackerList(
      trackers: null == trackers
          ? _value._trackers
          : trackers // ignore: cast_nullable_to_non_nullable
              as List<TrackerRow>,
    ));
  }
}

/// @nodoc

class _$_TrackerList implements _TrackerList {
  const _$_TrackerList({required final List<TrackerRow> trackers})
      : _trackers = trackers;

  final List<TrackerRow> _trackers;
  @override
  List<TrackerRow> get trackers {
    if (_trackers is EqualUnmodifiableListView) return _trackers;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_trackers);
  }

  @override
  String toString() {
    return 'TrackerList(trackers: $trackers)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_TrackerList &&
            const DeepCollectionEquality().equals(other._trackers, _trackers));
  }

  @override
  int get hashCode =>
      Object.hash(runtimeType, const DeepCollectionEquality().hash(_trackers));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_TrackerListCopyWith<_$_TrackerList> get copyWith =>
      __$$_TrackerListCopyWithImpl<_$_TrackerList>(this, _$identity);
}

abstract class _TrackerList implements TrackerList {
  const factory _TrackerList({required final List<TrackerRow> trackers}) =
      _$_TrackerList;

  @override
  List<TrackerRow> get trackers;
  @override
  @JsonKey(ignore: true)
  _$$_TrackerListCopyWith<_$_TrackerList> get copyWith =>
      throw _privateConstructorUsedError;
}
