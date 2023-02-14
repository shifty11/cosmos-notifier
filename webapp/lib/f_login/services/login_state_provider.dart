import 'package:cosmos_notifier/f_home/services/auth_provider.dart';
import 'package:cosmos_notifier/f_home/services/state/auth_state.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'state/login_state.dart';


final loginStateProvider = StateNotifierProvider<LoginNotifier, LoginState>(
  (ref) => LoginNotifier(ref),
);

class LoginNotifier extends StateNotifier<LoginState> {
  final StateNotifierProviderRef<LoginNotifier, LoginState> _ref;

  LoginNotifier(this._ref) : super(const LoginState.loading()) {
    _ref.watch(authStateProvider).maybeWhen(authenticated: (_, cannySSO) {
      if (cannySSO.ssoToken.isEmpty) {
        state = const LoginState.loading();
      } else {
        state = LoginState.authenticated(cannySSO);
      }
    }, unauthenticated: (cannySSO) {
      state = LoginState.unauthenticated(cannySSO);
    }, orElse: () {
      state = const LoginState.unauthenticated(CannySSO(false, "", "", ""));
    });
  }
}
