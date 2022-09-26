import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:dao_dao_notifier/config.dart';
import 'package:dao_dao_notifier/f_home/services/state/auth_state.dart';

import 'auth_service.dart';

final authProvider = Provider<AuthService>((ref) => authService);

final authStateProvider = StateNotifierProvider<AuthNotifier, AuthState>(
  (ref) => AuthNotifier(ref.watch(authProvider)),
);

final authStateValueProvider = Provider<ValueNotifier<AuthState>>((ref) {
  final notifier = ValueNotifier<AuthState>(ref.read(authStateProvider.notifier).state);
  ref.listen(authStateProvider, (_, next) {
    notifier.value = next as AuthState;
    notifier.notifyListeners();
  });
  return notifier;
});

class AuthNotifier extends StateNotifier<AuthState> {
  final AuthService _authService;

  AuthNotifier(this._authService) : super(const AuthState.initial()) {
    if (_authService.hasLoginData) {
      login();
      _authService.addListener(() {
        if (!_authService.isAuthenticated) {
          state = AuthState.error(AuthExpiredError());
        }
      });
    } else {
      state = const AuthState.unauthenticated();
      _authService.backgroundInit().then((isAuthenticated) {
        if (isAuthenticated) {
          state = const AuthState.authenticated(false);
        }
      });
    }
  }

  Future<void> login() async {
    try {
      state = const AuthState.loading();
      await _authService.init();
      state = const AuthState.authenticated(true);
    } catch (e) {
      if (cDebugMode) {
        print("AuthNotifier: error -> $e");
      }
      state = AuthState.error(e as Exception);
    }
  }
}
