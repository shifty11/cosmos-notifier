// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'auth_state.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$CannySSO {
  bool get isCannySSO => throw _privateConstructorUsedError;
  String get ssoToken => throw _privateConstructorUsedError;
  String get redirectUrl => throw _privateConstructorUsedError;
  String get companyId => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $CannySSOCopyWith<CannySSO> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $CannySSOCopyWith<$Res> {
  factory $CannySSOCopyWith(CannySSO value, $Res Function(CannySSO) then) =
      _$CannySSOCopyWithImpl<$Res, CannySSO>;
  @useResult
  $Res call(
      {bool isCannySSO, String ssoToken, String redirectUrl, String companyId});
}

/// @nodoc
class _$CannySSOCopyWithImpl<$Res, $Val extends CannySSO>
    implements $CannySSOCopyWith<$Res> {
  _$CannySSOCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isCannySSO = null,
    Object? ssoToken = null,
    Object? redirectUrl = null,
    Object? companyId = null,
  }) {
    return _then(_value.copyWith(
      isCannySSO: null == isCannySSO
          ? _value.isCannySSO
          : isCannySSO // ignore: cast_nullable_to_non_nullable
              as bool,
      ssoToken: null == ssoToken
          ? _value.ssoToken
          : ssoToken // ignore: cast_nullable_to_non_nullable
              as String,
      redirectUrl: null == redirectUrl
          ? _value.redirectUrl
          : redirectUrl // ignore: cast_nullable_to_non_nullable
              as String,
      companyId: null == companyId
          ? _value.companyId
          : companyId // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_CannySSOCopyWith<$Res> implements $CannySSOCopyWith<$Res> {
  factory _$$_CannySSOCopyWith(
          _$_CannySSO value, $Res Function(_$_CannySSO) then) =
      __$$_CannySSOCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {bool isCannySSO, String ssoToken, String redirectUrl, String companyId});
}

/// @nodoc
class __$$_CannySSOCopyWithImpl<$Res>
    extends _$CannySSOCopyWithImpl<$Res, _$_CannySSO>
    implements _$$_CannySSOCopyWith<$Res> {
  __$$_CannySSOCopyWithImpl(
      _$_CannySSO _value, $Res Function(_$_CannySSO) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isCannySSO = null,
    Object? ssoToken = null,
    Object? redirectUrl = null,
    Object? companyId = null,
  }) {
    return _then(_$_CannySSO(
      null == isCannySSO
          ? _value.isCannySSO
          : isCannySSO // ignore: cast_nullable_to_non_nullable
              as bool,
      null == ssoToken
          ? _value.ssoToken
          : ssoToken // ignore: cast_nullable_to_non_nullable
              as String,
      null == redirectUrl
          ? _value.redirectUrl
          : redirectUrl // ignore: cast_nullable_to_non_nullable
              as String,
      null == companyId
          ? _value.companyId
          : companyId // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc

class _$_CannySSO with DiagnosticableTreeMixin implements _CannySSO {
  const _$_CannySSO(
      this.isCannySSO, this.ssoToken, this.redirectUrl, this.companyId);

  @override
  final bool isCannySSO;
  @override
  final String ssoToken;
  @override
  final String redirectUrl;
  @override
  final String companyId;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'CannySSO(isCannySSO: $isCannySSO, ssoToken: $ssoToken, redirectUrl: $redirectUrl, companyId: $companyId)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'CannySSO'))
      ..add(DiagnosticsProperty('isCannySSO', isCannySSO))
      ..add(DiagnosticsProperty('ssoToken', ssoToken))
      ..add(DiagnosticsProperty('redirectUrl', redirectUrl))
      ..add(DiagnosticsProperty('companyId', companyId));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_CannySSO &&
            (identical(other.isCannySSO, isCannySSO) ||
                other.isCannySSO == isCannySSO) &&
            (identical(other.ssoToken, ssoToken) ||
                other.ssoToken == ssoToken) &&
            (identical(other.redirectUrl, redirectUrl) ||
                other.redirectUrl == redirectUrl) &&
            (identical(other.companyId, companyId) ||
                other.companyId == companyId));
  }

  @override
  int get hashCode =>
      Object.hash(runtimeType, isCannySSO, ssoToken, redirectUrl, companyId);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_CannySSOCopyWith<_$_CannySSO> get copyWith =>
      __$$_CannySSOCopyWithImpl<_$_CannySSO>(this, _$identity);
}

abstract class _CannySSO implements CannySSO {
  const factory _CannySSO(final bool isCannySSO, final String ssoToken,
      final String redirectUrl, final String companyId) = _$_CannySSO;

  @override
  bool get isCannySSO;
  @override
  String get ssoToken;
  @override
  String get redirectUrl;
  @override
  String get companyId;
  @override
  @JsonKey(ignore: true)
  _$$_CannySSOCopyWith<_$_CannySSO> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
mixin _$AuthState {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() initial,
    required TResult Function() loading,
    required TResult Function(bool redirect, CannySSO cannySSO) authenticated,
    required TResult Function(CannySSO cannySSO) unauthenticated,
    required TResult Function(Exception error) error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? initial,
    TResult? Function()? loading,
    TResult? Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult? Function(CannySSO cannySSO)? unauthenticated,
    TResult? Function(Exception error)? error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? initial,
    TResult Function()? loading,
    TResult Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult Function(CannySSO cannySSO)? unauthenticated,
    TResult Function(Exception error)? error,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Loading value) loading,
    required TResult Function(Authenticated value) authenticated,
    required TResult Function(Unauthenticated value) unauthenticated,
    required TResult Function(Error value) error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Loading value)? loading,
    TResult? Function(Authenticated value)? authenticated,
    TResult? Function(Unauthenticated value)? unauthenticated,
    TResult? Function(Error value)? error,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Loading value)? loading,
    TResult Function(Authenticated value)? authenticated,
    TResult Function(Unauthenticated value)? unauthenticated,
    TResult Function(Error value)? error,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $AuthStateCopyWith<$Res> {
  factory $AuthStateCopyWith(AuthState value, $Res Function(AuthState) then) =
      _$AuthStateCopyWithImpl<$Res, AuthState>;
}

/// @nodoc
class _$AuthStateCopyWithImpl<$Res, $Val extends AuthState>
    implements $AuthStateCopyWith<$Res> {
  _$AuthStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;
}

/// @nodoc
abstract class _$$InitialCopyWith<$Res> {
  factory _$$InitialCopyWith(_$Initial value, $Res Function(_$Initial) then) =
      __$$InitialCopyWithImpl<$Res>;
}

/// @nodoc
class __$$InitialCopyWithImpl<$Res>
    extends _$AuthStateCopyWithImpl<$Res, _$Initial>
    implements _$$InitialCopyWith<$Res> {
  __$$InitialCopyWithImpl(_$Initial _value, $Res Function(_$Initial) _then)
      : super(_value, _then);
}

/// @nodoc

class _$Initial extends Initial with DiagnosticableTreeMixin {
  const _$Initial() : super._();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AuthState.initial()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'AuthState.initial'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$Initial);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() initial,
    required TResult Function() loading,
    required TResult Function(bool redirect, CannySSO cannySSO) authenticated,
    required TResult Function(CannySSO cannySSO) unauthenticated,
    required TResult Function(Exception error) error,
  }) {
    return initial();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? initial,
    TResult? Function()? loading,
    TResult? Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult? Function(CannySSO cannySSO)? unauthenticated,
    TResult? Function(Exception error)? error,
  }) {
    return initial?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? initial,
    TResult Function()? loading,
    TResult Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult Function(CannySSO cannySSO)? unauthenticated,
    TResult Function(Exception error)? error,
    required TResult orElse(),
  }) {
    if (initial != null) {
      return initial();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Loading value) loading,
    required TResult Function(Authenticated value) authenticated,
    required TResult Function(Unauthenticated value) unauthenticated,
    required TResult Function(Error value) error,
  }) {
    return initial(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Loading value)? loading,
    TResult? Function(Authenticated value)? authenticated,
    TResult? Function(Unauthenticated value)? unauthenticated,
    TResult? Function(Error value)? error,
  }) {
    return initial?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Loading value)? loading,
    TResult Function(Authenticated value)? authenticated,
    TResult Function(Unauthenticated value)? unauthenticated,
    TResult Function(Error value)? error,
    required TResult orElse(),
  }) {
    if (initial != null) {
      return initial(this);
    }
    return orElse();
  }
}

abstract class Initial extends AuthState {
  const factory Initial() = _$Initial;
  const Initial._() : super._();
}

/// @nodoc
abstract class _$$LoadingCopyWith<$Res> {
  factory _$$LoadingCopyWith(_$Loading value, $Res Function(_$Loading) then) =
      __$$LoadingCopyWithImpl<$Res>;
}

/// @nodoc
class __$$LoadingCopyWithImpl<$Res>
    extends _$AuthStateCopyWithImpl<$Res, _$Loading>
    implements _$$LoadingCopyWith<$Res> {
  __$$LoadingCopyWithImpl(_$Loading _value, $Res Function(_$Loading) _then)
      : super(_value, _then);
}

/// @nodoc

class _$Loading extends Loading with DiagnosticableTreeMixin {
  const _$Loading() : super._();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AuthState.loading()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'AuthState.loading'));
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
    required TResult Function() initial,
    required TResult Function() loading,
    required TResult Function(bool redirect, CannySSO cannySSO) authenticated,
    required TResult Function(CannySSO cannySSO) unauthenticated,
    required TResult Function(Exception error) error,
  }) {
    return loading();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? initial,
    TResult? Function()? loading,
    TResult? Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult? Function(CannySSO cannySSO)? unauthenticated,
    TResult? Function(Exception error)? error,
  }) {
    return loading?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? initial,
    TResult Function()? loading,
    TResult Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult Function(CannySSO cannySSO)? unauthenticated,
    TResult Function(Exception error)? error,
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
    required TResult Function(Initial value) initial,
    required TResult Function(Loading value) loading,
    required TResult Function(Authenticated value) authenticated,
    required TResult Function(Unauthenticated value) unauthenticated,
    required TResult Function(Error value) error,
  }) {
    return loading(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Loading value)? loading,
    TResult? Function(Authenticated value)? authenticated,
    TResult? Function(Unauthenticated value)? unauthenticated,
    TResult? Function(Error value)? error,
  }) {
    return loading?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Loading value)? loading,
    TResult Function(Authenticated value)? authenticated,
    TResult Function(Unauthenticated value)? unauthenticated,
    TResult Function(Error value)? error,
    required TResult orElse(),
  }) {
    if (loading != null) {
      return loading(this);
    }
    return orElse();
  }
}

abstract class Loading extends AuthState {
  const factory Loading() = _$Loading;
  const Loading._() : super._();
}

/// @nodoc
abstract class _$$AuthenticatedCopyWith<$Res> {
  factory _$$AuthenticatedCopyWith(
          _$Authenticated value, $Res Function(_$Authenticated) then) =
      __$$AuthenticatedCopyWithImpl<$Res>;
  @useResult
  $Res call({bool redirect, CannySSO cannySSO});

  $CannySSOCopyWith<$Res> get cannySSO;
}

/// @nodoc
class __$$AuthenticatedCopyWithImpl<$Res>
    extends _$AuthStateCopyWithImpl<$Res, _$Authenticated>
    implements _$$AuthenticatedCopyWith<$Res> {
  __$$AuthenticatedCopyWithImpl(
      _$Authenticated _value, $Res Function(_$Authenticated) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? redirect = null,
    Object? cannySSO = null,
  }) {
    return _then(_$Authenticated(
      null == redirect
          ? _value.redirect
          : redirect // ignore: cast_nullable_to_non_nullable
              as bool,
      null == cannySSO
          ? _value.cannySSO
          : cannySSO // ignore: cast_nullable_to_non_nullable
              as CannySSO,
    ));
  }

  @override
  @pragma('vm:prefer-inline')
  $CannySSOCopyWith<$Res> get cannySSO {
    return $CannySSOCopyWith<$Res>(_value.cannySSO, (value) {
      return _then(_value.copyWith(cannySSO: value));
    });
  }
}

/// @nodoc

class _$Authenticated extends Authenticated with DiagnosticableTreeMixin {
  const _$Authenticated(this.redirect, this.cannySSO) : super._();

  @override
  final bool redirect;
  @override
  final CannySSO cannySSO;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AuthState.authenticated(redirect: $redirect, cannySSO: $cannySSO)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'AuthState.authenticated'))
      ..add(DiagnosticsProperty('redirect', redirect))
      ..add(DiagnosticsProperty('cannySSO', cannySSO));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$Authenticated &&
            (identical(other.redirect, redirect) ||
                other.redirect == redirect) &&
            (identical(other.cannySSO, cannySSO) ||
                other.cannySSO == cannySSO));
  }

  @override
  int get hashCode => Object.hash(runtimeType, redirect, cannySSO);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$AuthenticatedCopyWith<_$Authenticated> get copyWith =>
      __$$AuthenticatedCopyWithImpl<_$Authenticated>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() initial,
    required TResult Function() loading,
    required TResult Function(bool redirect, CannySSO cannySSO) authenticated,
    required TResult Function(CannySSO cannySSO) unauthenticated,
    required TResult Function(Exception error) error,
  }) {
    return authenticated(redirect, cannySSO);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? initial,
    TResult? Function()? loading,
    TResult? Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult? Function(CannySSO cannySSO)? unauthenticated,
    TResult? Function(Exception error)? error,
  }) {
    return authenticated?.call(redirect, cannySSO);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? initial,
    TResult Function()? loading,
    TResult Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult Function(CannySSO cannySSO)? unauthenticated,
    TResult Function(Exception error)? error,
    required TResult orElse(),
  }) {
    if (authenticated != null) {
      return authenticated(redirect, cannySSO);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Loading value) loading,
    required TResult Function(Authenticated value) authenticated,
    required TResult Function(Unauthenticated value) unauthenticated,
    required TResult Function(Error value) error,
  }) {
    return authenticated(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Loading value)? loading,
    TResult? Function(Authenticated value)? authenticated,
    TResult? Function(Unauthenticated value)? unauthenticated,
    TResult? Function(Error value)? error,
  }) {
    return authenticated?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Loading value)? loading,
    TResult Function(Authenticated value)? authenticated,
    TResult Function(Unauthenticated value)? unauthenticated,
    TResult Function(Error value)? error,
    required TResult orElse(),
  }) {
    if (authenticated != null) {
      return authenticated(this);
    }
    return orElse();
  }
}

abstract class Authenticated extends AuthState {
  const factory Authenticated(final bool redirect, final CannySSO cannySSO) =
      _$Authenticated;
  const Authenticated._() : super._();

  bool get redirect;
  CannySSO get cannySSO;
  @JsonKey(ignore: true)
  _$$AuthenticatedCopyWith<_$Authenticated> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class _$$UnauthenticatedCopyWith<$Res> {
  factory _$$UnauthenticatedCopyWith(
          _$Unauthenticated value, $Res Function(_$Unauthenticated) then) =
      __$$UnauthenticatedCopyWithImpl<$Res>;
  @useResult
  $Res call({CannySSO cannySSO});

  $CannySSOCopyWith<$Res> get cannySSO;
}

/// @nodoc
class __$$UnauthenticatedCopyWithImpl<$Res>
    extends _$AuthStateCopyWithImpl<$Res, _$Unauthenticated>
    implements _$$UnauthenticatedCopyWith<$Res> {
  __$$UnauthenticatedCopyWithImpl(
      _$Unauthenticated _value, $Res Function(_$Unauthenticated) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? cannySSO = null,
  }) {
    return _then(_$Unauthenticated(
      null == cannySSO
          ? _value.cannySSO
          : cannySSO // ignore: cast_nullable_to_non_nullable
              as CannySSO,
    ));
  }

  @override
  @pragma('vm:prefer-inline')
  $CannySSOCopyWith<$Res> get cannySSO {
    return $CannySSOCopyWith<$Res>(_value.cannySSO, (value) {
      return _then(_value.copyWith(cannySSO: value));
    });
  }
}

/// @nodoc

class _$Unauthenticated extends Unauthenticated with DiagnosticableTreeMixin {
  const _$Unauthenticated(this.cannySSO) : super._();

  @override
  final CannySSO cannySSO;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AuthState.unauthenticated(cannySSO: $cannySSO)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'AuthState.unauthenticated'))
      ..add(DiagnosticsProperty('cannySSO', cannySSO));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$Unauthenticated &&
            (identical(other.cannySSO, cannySSO) ||
                other.cannySSO == cannySSO));
  }

  @override
  int get hashCode => Object.hash(runtimeType, cannySSO);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$UnauthenticatedCopyWith<_$Unauthenticated> get copyWith =>
      __$$UnauthenticatedCopyWithImpl<_$Unauthenticated>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() initial,
    required TResult Function() loading,
    required TResult Function(bool redirect, CannySSO cannySSO) authenticated,
    required TResult Function(CannySSO cannySSO) unauthenticated,
    required TResult Function(Exception error) error,
  }) {
    return unauthenticated(cannySSO);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? initial,
    TResult? Function()? loading,
    TResult? Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult? Function(CannySSO cannySSO)? unauthenticated,
    TResult? Function(Exception error)? error,
  }) {
    return unauthenticated?.call(cannySSO);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? initial,
    TResult Function()? loading,
    TResult Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult Function(CannySSO cannySSO)? unauthenticated,
    TResult Function(Exception error)? error,
    required TResult orElse(),
  }) {
    if (unauthenticated != null) {
      return unauthenticated(cannySSO);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Loading value) loading,
    required TResult Function(Authenticated value) authenticated,
    required TResult Function(Unauthenticated value) unauthenticated,
    required TResult Function(Error value) error,
  }) {
    return unauthenticated(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Loading value)? loading,
    TResult? Function(Authenticated value)? authenticated,
    TResult? Function(Unauthenticated value)? unauthenticated,
    TResult? Function(Error value)? error,
  }) {
    return unauthenticated?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Loading value)? loading,
    TResult Function(Authenticated value)? authenticated,
    TResult Function(Unauthenticated value)? unauthenticated,
    TResult Function(Error value)? error,
    required TResult orElse(),
  }) {
    if (unauthenticated != null) {
      return unauthenticated(this);
    }
    return orElse();
  }
}

abstract class Unauthenticated extends AuthState {
  const factory Unauthenticated(final CannySSO cannySSO) = _$Unauthenticated;
  const Unauthenticated._() : super._();

  CannySSO get cannySSO;
  @JsonKey(ignore: true)
  _$$UnauthenticatedCopyWith<_$Unauthenticated> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class _$$ErrorCopyWith<$Res> {
  factory _$$ErrorCopyWith(_$Error value, $Res Function(_$Error) then) =
      __$$ErrorCopyWithImpl<$Res>;
  @useResult
  $Res call({Exception error});
}

/// @nodoc
class __$$ErrorCopyWithImpl<$Res> extends _$AuthStateCopyWithImpl<$Res, _$Error>
    implements _$$ErrorCopyWith<$Res> {
  __$$ErrorCopyWithImpl(_$Error _value, $Res Function(_$Error) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? error = null,
  }) {
    return _then(_$Error(
      null == error
          ? _value.error
          : error // ignore: cast_nullable_to_non_nullable
              as Exception,
    ));
  }
}

/// @nodoc

class _$Error extends Error with DiagnosticableTreeMixin {
  const _$Error(this.error) : super._();

  @override
  final Exception error;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AuthState.error(error: $error)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'AuthState.error'))
      ..add(DiagnosticsProperty('error', error));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$Error &&
            (identical(other.error, error) || other.error == error));
  }

  @override
  int get hashCode => Object.hash(runtimeType, error);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$ErrorCopyWith<_$Error> get copyWith =>
      __$$ErrorCopyWithImpl<_$Error>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() initial,
    required TResult Function() loading,
    required TResult Function(bool redirect, CannySSO cannySSO) authenticated,
    required TResult Function(CannySSO cannySSO) unauthenticated,
    required TResult Function(Exception error) error,
  }) {
    return error(this.error);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? initial,
    TResult? Function()? loading,
    TResult? Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult? Function(CannySSO cannySSO)? unauthenticated,
    TResult? Function(Exception error)? error,
  }) {
    return error?.call(this.error);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? initial,
    TResult Function()? loading,
    TResult Function(bool redirect, CannySSO cannySSO)? authenticated,
    TResult Function(CannySSO cannySSO)? unauthenticated,
    TResult Function(Exception error)? error,
    required TResult orElse(),
  }) {
    if (error != null) {
      return error(this.error);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(Initial value) initial,
    required TResult Function(Loading value) loading,
    required TResult Function(Authenticated value) authenticated,
    required TResult Function(Unauthenticated value) unauthenticated,
    required TResult Function(Error value) error,
  }) {
    return error(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(Initial value)? initial,
    TResult? Function(Loading value)? loading,
    TResult? Function(Authenticated value)? authenticated,
    TResult? Function(Unauthenticated value)? unauthenticated,
    TResult? Function(Error value)? error,
  }) {
    return error?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(Initial value)? initial,
    TResult Function(Loading value)? loading,
    TResult Function(Authenticated value)? authenticated,
    TResult Function(Unauthenticated value)? unauthenticated,
    TResult Function(Error value)? error,
    required TResult orElse(),
  }) {
    if (error != null) {
      return error(this);
    }
    return orElse();
  }
}

abstract class Error extends AuthState {
  const factory Error(final Exception error) = _$Error;
  const Error._() : super._();

  Exception get error;
  @JsonKey(ignore: true)
  _$$ErrorCopyWith<_$Error> get copyWith => throw _privateConstructorUsedError;
}
