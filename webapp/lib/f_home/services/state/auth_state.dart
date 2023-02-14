import 'package:flutter/foundation.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'auth_state.freezed.dart';

@freezed
class CannySSO with _$CannySSO {
  const factory CannySSO(bool isCannySSO, String ssoToken, String redirectUrl, String companyId) = _CannySSO;
}

@freezed
class AuthState with _$AuthState {
  const AuthState._();

  const factory AuthState.initial() = Initial;

  const factory AuthState.loading() = Loading;

  const factory AuthState.authenticated(bool redirect, CannySSO cannySSO) = Authenticated;

  const factory AuthState.unauthenticated(CannySSO cannySSO) = Unauthenticated;

  const factory AuthState.error(Exception error) = Error;
}
