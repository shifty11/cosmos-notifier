import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';

part 'auth_state.freezed.dart';

@freezed
class AuthState with _$AuthState {
  const AuthState._();

  const factory AuthState.loading() = Loading;
  const factory AuthState.authorized() = Authorized;
  const factory AuthState.expired() = Expired;
  const factory AuthState.error() = Error;
}