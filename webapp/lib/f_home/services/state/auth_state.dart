import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';

part 'auth_state.freezed.dart';

@freezed
class AuthState with _$AuthState {
  const AuthState._();

  const factory AuthState.initial() = Initial;
  const factory AuthState.loading() = Loading;
  const factory AuthState.authenticated(bool redirect) = Authenticated;
  const factory AuthState.unauthenticated() = Unauthenticated;
  const factory AuthState.error(Exception error) = Error;
}