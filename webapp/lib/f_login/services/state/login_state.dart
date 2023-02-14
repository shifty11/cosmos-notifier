import 'package:cosmos_notifier/f_home/services/state/auth_state.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'login_state.freezed.dart';

@freezed
class LoginState with _$LoginState {
  const LoginState._();

  const factory LoginState.loading() = Loading;

  const factory LoginState.authenticated(CannySSO cannySSO) = Authenticated;

  const factory LoginState.unauthenticated(CannySSO cannySSO) = Unauthenticated;

  const factory LoginState.error(Exception error) = Error;
}
