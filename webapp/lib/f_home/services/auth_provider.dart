import 'package:webapp/config.dart';
import 'package:webapp/f_home/services/state/auth_state.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'auth_service.dart';

final authProvider = Provider<AuthService>((ref) => authService);

final authStateProvider = StateNotifierProvider<AuthNotifier, AuthState>(
  (ref) => AuthNotifier(ref.watch(authProvider)),
);

final authStateValueProvider = Provider<ValueNotifier<AuthState>>((ref) {
  final notifier = ValueNotifier<AuthState>(const AuthState.loading());
  ref.listen(authStateProvider, (_, next) {
    notifier.value = next as AuthState;
    notifier.notifyListeners();
  });
  return notifier;
});

class AuthNotifier extends StateNotifier<AuthState> {
  final AuthService _authService;

  AuthNotifier(this._authService) : super(const AuthState.loading()) {
    login();
    _authService.addListener(() {
      if (!_authService.isAuthenticated) {
        state = const AuthState.expired();
      }
    });
  }

  Future<void> login() async {
    try {
      state = const AuthState.loading();
      await _authService.init();
      state = const AuthState.authorized();
    } catch (e) {
      if (cDebugMode) {
        print("AuthNotifier: error -> $e");
      }
      if (e is AuthExpiredError) {
        state = const AuthState.expired();
      } else if (e is AuthUserNotFoundError) {
        state = const AuthState.userNotFound();
      } else {
        state = const AuthState.error();
      }
    }
  }
}
