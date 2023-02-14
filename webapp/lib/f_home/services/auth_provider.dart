import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:cosmos_notifier/f_home/services/canny_query_param_provider.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/state/auth_state.dart';

import 'auth_service.dart';

final authProvider = Provider<AuthService>((ref) => authService);

final authStateProvider = StateNotifierProvider<AuthNotifier, AuthState>(
  (ref) => AuthNotifier(ref.watch(authProvider), ref.watch(cannySSOProvider)),
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
  final CannySSO _cannySSO;

  AuthNotifier(this._authService, this._cannySSO) : super(const AuthState.initial()) {
    if (_authService.hasLoginData) {
      login();
      _authService.addListener(() {
        if (!_authService.isAuthenticated) {
          state = AuthState.error(AuthExpiredError());
        }
      });
    } else {
      state = AuthState.unauthenticated(_cannySSO);
      _authService.backgroundInit().then((isAuthenticated) async {
        if (isAuthenticated) {
          var cannySSOResult = await _getCannySSO();
          state = AuthState.authenticated(false, cannySSOResult);
        }
      });
    }
  }

  Future<CannySSO> _getCannySSO() async {
    if (_cannySSO.isCannySSO) {
      var response = await _authService.cannySSO(Empty());
      return _cannySSO.copyWith(ssoToken: response.accessToken);
    }
    return const CannySSO(false, "", "", "");
  }

  Future<void> login() async {
    try {
      state = const AuthState.loading();
      await _authService.init();
      var cannySSOResult = await _getCannySSO();
      state = AuthState.authenticated(true, cannySSOResult);
    } catch (e) {
      if (cDebugMode) {
        print("AuthNotifier: error -> $e");
      }
      state = AuthState.error(e as Exception);
    }
  }
}
