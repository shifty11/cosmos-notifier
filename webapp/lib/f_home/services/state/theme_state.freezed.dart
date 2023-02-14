// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'theme_state.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$ThemeState {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function(ThemeData darkStyle, ThemeData lightStyle)
        initial,
    required TResult Function(ThemeData style) custom,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function(ThemeData darkStyle, ThemeData lightStyle)? initial,
    TResult? Function(ThemeData style)? custom,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function(ThemeData darkStyle, ThemeData lightStyle)? initial,
    TResult Function(ThemeData style)? custom,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Custom value) custom,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Custom value)? custom,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Custom value)? custom,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $ThemeStateCopyWith<$Res> {
  factory $ThemeStateCopyWith(
          ThemeState value, $Res Function(ThemeState) then) =
      _$ThemeStateCopyWithImpl<$Res, ThemeState>;
}

/// @nodoc
class _$ThemeStateCopyWithImpl<$Res, $Val extends ThemeState>
    implements $ThemeStateCopyWith<$Res> {
  _$ThemeStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;
}

/// @nodoc
abstract class _$$InitialCopyWith<$Res> {
  factory _$$InitialCopyWith(_$Initial value, $Res Function(_$Initial) then) =
      __$$InitialCopyWithImpl<$Res>;
  @useResult
  $Res call({ThemeData darkStyle, ThemeData lightStyle});
}

/// @nodoc
class __$$InitialCopyWithImpl<$Res>
    extends _$ThemeStateCopyWithImpl<$Res, _$Initial>
    implements _$$InitialCopyWith<$Res> {
  __$$InitialCopyWithImpl(_$Initial _value, $Res Function(_$Initial) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? darkStyle = null,
    Object? lightStyle = null,
  }) {
    return _then(_$Initial(
      darkStyle: null == darkStyle
          ? _value.darkStyle
          : darkStyle // ignore: cast_nullable_to_non_nullable
              as ThemeData,
      lightStyle: null == lightStyle
          ? _value.lightStyle
          : lightStyle // ignore: cast_nullable_to_non_nullable
              as ThemeData,
    ));
  }
}

/// @nodoc

class _$Initial extends Initial {
  const _$Initial({required this.darkStyle, required this.lightStyle})
      : super._();

  @override
  final ThemeData darkStyle;
  @override
  final ThemeData lightStyle;

  @override
  String toString() {
    return 'ThemeState.initial(darkStyle: $darkStyle, lightStyle: $lightStyle)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$Initial &&
            (identical(other.darkStyle, darkStyle) ||
                other.darkStyle == darkStyle) &&
            (identical(other.lightStyle, lightStyle) ||
                other.lightStyle == lightStyle));
  }

  @override
  int get hashCode => Object.hash(runtimeType, darkStyle, lightStyle);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$InitialCopyWith<_$Initial> get copyWith =>
      __$$InitialCopyWithImpl<_$Initial>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function(ThemeData darkStyle, ThemeData lightStyle)
        initial,
    required TResult Function(ThemeData style) custom,
  }) {
    return initial(darkStyle, lightStyle);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function(ThemeData darkStyle, ThemeData lightStyle)? initial,
    TResult? Function(ThemeData style)? custom,
  }) {
    return initial?.call(darkStyle, lightStyle);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function(ThemeData darkStyle, ThemeData lightStyle)? initial,
    TResult Function(ThemeData style)? custom,
    required TResult orElse(),
  }) {
    if (initial != null) {
      return initial(darkStyle, lightStyle);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Custom value) custom,
  }) {
    return initial(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Custom value)? custom,
  }) {
    return initial?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Custom value)? custom,
    required TResult orElse(),
  }) {
    if (initial != null) {
      return initial(this);
    }
    return orElse();
  }
}

abstract class Initial extends ThemeState {
  const factory Initial(
      {required final ThemeData darkStyle,
      required final ThemeData lightStyle}) = _$Initial;
  const Initial._() : super._();

  ThemeData get darkStyle;
  ThemeData get lightStyle;
  @JsonKey(ignore: true)
  _$$InitialCopyWith<_$Initial> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class _$$CustomCopyWith<$Res> {
  factory _$$CustomCopyWith(_$Custom value, $Res Function(_$Custom) then) =
      __$$CustomCopyWithImpl<$Res>;
  @useResult
  $Res call({ThemeData style});
}

/// @nodoc
class __$$CustomCopyWithImpl<$Res>
    extends _$ThemeStateCopyWithImpl<$Res, _$Custom>
    implements _$$CustomCopyWith<$Res> {
  __$$CustomCopyWithImpl(_$Custom _value, $Res Function(_$Custom) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? style = null,
  }) {
    return _then(_$Custom(
      style: null == style
          ? _value.style
          : style // ignore: cast_nullable_to_non_nullable
              as ThemeData,
    ));
  }
}

/// @nodoc

class _$Custom extends Custom {
  const _$Custom({required this.style}) : super._();

  @override
  final ThemeData style;

  @override
  String toString() {
    return 'ThemeState.custom(style: $style)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$Custom &&
            (identical(other.style, style) || other.style == style));
  }

  @override
  int get hashCode => Object.hash(runtimeType, style);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$CustomCopyWith<_$Custom> get copyWith =>
      __$$CustomCopyWithImpl<_$Custom>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function(ThemeData darkStyle, ThemeData lightStyle)
        initial,
    required TResult Function(ThemeData style) custom,
  }) {
    return custom(style);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function(ThemeData darkStyle, ThemeData lightStyle)? initial,
    TResult? Function(ThemeData style)? custom,
  }) {
    return custom?.call(style);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function(ThemeData darkStyle, ThemeData lightStyle)? initial,
    TResult Function(ThemeData style)? custom,
    required TResult orElse(),
  }) {
    if (custom != null) {
      return custom(style);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Custom value) custom,
  }) {
    return custom(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Custom value)? custom,
  }) {
    return custom?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Custom value)? custom,
    required TResult orElse(),
  }) {
    if (custom != null) {
      return custom(this);
    }
    return orElse();
  }
}

abstract class Custom extends ThemeState {
  const factory Custom({required final ThemeData style}) = _$Custom;
  const Custom._() : super._();

  ThemeData get style;
  @JsonKey(ignore: true)
  _$$CustomCopyWith<_$Custom> get copyWith =>
      throw _privateConstructorUsedError;
}
