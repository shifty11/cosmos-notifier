// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'validator_bundle.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$FreezedValidator {
  Int64 get id => throw _privateConstructorUsedError;
  String get address => throw _privateConstructorUsedError;
  String get chainName => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $FreezedValidatorCopyWith<FreezedValidator> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $FreezedValidatorCopyWith<$Res> {
  factory $FreezedValidatorCopyWith(
          FreezedValidator value, $Res Function(FreezedValidator) then) =
      _$FreezedValidatorCopyWithImpl<$Res, FreezedValidator>;
  @useResult
  $Res call({Int64 id, String address, String chainName});
}

/// @nodoc
class _$FreezedValidatorCopyWithImpl<$Res, $Val extends FreezedValidator>
    implements $FreezedValidatorCopyWith<$Res> {
  _$FreezedValidatorCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? address = null,
    Object? chainName = null,
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
      chainName: null == chainName
          ? _value.chainName
          : chainName // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_ValidatorCopyWith<$Res>
    implements $FreezedValidatorCopyWith<$Res> {
  factory _$$_ValidatorCopyWith(
          _$_Validator value, $Res Function(_$_Validator) then) =
      __$$_ValidatorCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({Int64 id, String address, String chainName});
}

/// @nodoc
class __$$_ValidatorCopyWithImpl<$Res>
    extends _$FreezedValidatorCopyWithImpl<$Res, _$_Validator>
    implements _$$_ValidatorCopyWith<$Res> {
  __$$_ValidatorCopyWithImpl(
      _$_Validator _value, $Res Function(_$_Validator) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? address = null,
    Object? chainName = null,
  }) {
    return _then(_$_Validator(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as Int64,
      address: null == address
          ? _value.address
          : address // ignore: cast_nullable_to_non_nullable
              as String,
      chainName: null == chainName
          ? _value.chainName
          : chainName // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc

class _$_Validator extends _Validator {
  const _$_Validator(
      {required this.id, required this.address, required this.chainName})
      : super._();

  @override
  final Int64 id;
  @override
  final String address;
  @override
  final String chainName;

  @override
  String toString() {
    return 'FreezedValidator(id: $id, address: $address, chainName: $chainName)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_Validator &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.address, address) || other.address == address) &&
            (identical(other.chainName, chainName) ||
                other.chainName == chainName));
  }

  @override
  int get hashCode => Object.hash(runtimeType, id, address, chainName);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_ValidatorCopyWith<_$_Validator> get copyWith =>
      __$$_ValidatorCopyWithImpl<_$_Validator>(this, _$identity);
}

abstract class _Validator extends FreezedValidator {
  const factory _Validator(
      {required final Int64 id,
      required final String address,
      required final String chainName}) = _$_Validator;
  const _Validator._() : super._();

  @override
  Int64 get id;
  @override
  String get address;
  @override
  String get chainName;
  @override
  @JsonKey(ignore: true)
  _$$_ValidatorCopyWith<_$_Validator> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
mixin _$FreezedValidatorBundle {
  String get moniker => throw _privateConstructorUsedError;
  List<FreezedValidator> get validators => throw _privateConstructorUsedError;
  bool get isTracked => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $FreezedValidatorBundleCopyWith<FreezedValidatorBundle> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $FreezedValidatorBundleCopyWith<$Res> {
  factory $FreezedValidatorBundleCopyWith(FreezedValidatorBundle value,
          $Res Function(FreezedValidatorBundle) then) =
      _$FreezedValidatorBundleCopyWithImpl<$Res, FreezedValidatorBundle>;
  @useResult
  $Res call(
      {String moniker, List<FreezedValidator> validators, bool isTracked});
}

/// @nodoc
class _$FreezedValidatorBundleCopyWithImpl<$Res,
        $Val extends FreezedValidatorBundle>
    implements $FreezedValidatorBundleCopyWith<$Res> {
  _$FreezedValidatorBundleCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? moniker = null,
    Object? validators = null,
    Object? isTracked = null,
  }) {
    return _then(_value.copyWith(
      moniker: null == moniker
          ? _value.moniker
          : moniker // ignore: cast_nullable_to_non_nullable
              as String,
      validators: null == validators
          ? _value.validators
          : validators // ignore: cast_nullable_to_non_nullable
              as List<FreezedValidator>,
      isTracked: null == isTracked
          ? _value.isTracked
          : isTracked // ignore: cast_nullable_to_non_nullable
              as bool,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_ValidatorBundleCopyWith<$Res>
    implements $FreezedValidatorBundleCopyWith<$Res> {
  factory _$$_ValidatorBundleCopyWith(
          _$_ValidatorBundle value, $Res Function(_$_ValidatorBundle) then) =
      __$$_ValidatorBundleCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String moniker, List<FreezedValidator> validators, bool isTracked});
}

/// @nodoc
class __$$_ValidatorBundleCopyWithImpl<$Res>
    extends _$FreezedValidatorBundleCopyWithImpl<$Res, _$_ValidatorBundle>
    implements _$$_ValidatorBundleCopyWith<$Res> {
  __$$_ValidatorBundleCopyWithImpl(
      _$_ValidatorBundle _value, $Res Function(_$_ValidatorBundle) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? moniker = null,
    Object? validators = null,
    Object? isTracked = null,
  }) {
    return _then(_$_ValidatorBundle(
      moniker: null == moniker
          ? _value.moniker
          : moniker // ignore: cast_nullable_to_non_nullable
              as String,
      validators: null == validators
          ? _value._validators
          : validators // ignore: cast_nullable_to_non_nullable
              as List<FreezedValidator>,
      isTracked: null == isTracked
          ? _value.isTracked
          : isTracked // ignore: cast_nullable_to_non_nullable
              as bool,
    ));
  }
}

/// @nodoc

class _$_ValidatorBundle extends _ValidatorBundle {
  const _$_ValidatorBundle(
      {required this.moniker,
      required final List<FreezedValidator> validators,
      required this.isTracked})
      : _validators = validators,
        super._();

  @override
  final String moniker;
  final List<FreezedValidator> _validators;
  @override
  List<FreezedValidator> get validators {
    if (_validators is EqualUnmodifiableListView) return _validators;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_validators);
  }

  @override
  final bool isTracked;

  @override
  String toString() {
    return 'FreezedValidatorBundle(moniker: $moniker, validators: $validators, isTracked: $isTracked)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_ValidatorBundle &&
            (identical(other.moniker, moniker) || other.moniker == moniker) &&
            const DeepCollectionEquality()
                .equals(other._validators, _validators) &&
            (identical(other.isTracked, isTracked) ||
                other.isTracked == isTracked));
  }

  @override
  int get hashCode => Object.hash(runtimeType, moniker,
      const DeepCollectionEquality().hash(_validators), isTracked);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_ValidatorBundleCopyWith<_$_ValidatorBundle> get copyWith =>
      __$$_ValidatorBundleCopyWithImpl<_$_ValidatorBundle>(this, _$identity);
}

abstract class _ValidatorBundle extends FreezedValidatorBundle {
  const factory _ValidatorBundle(
      {required final String moniker,
      required final List<FreezedValidator> validators,
      required final bool isTracked}) = _$_ValidatorBundle;
  const _ValidatorBundle._() : super._();

  @override
  String get moniker;
  @override
  List<FreezedValidator> get validators;
  @override
  bool get isTracked;
  @override
  @JsonKey(ignore: true)
  _$$_ValidatorBundleCopyWith<_$_ValidatorBundle> get copyWith =>
      throw _privateConstructorUsedError;
}
