@JS()
library script.js;

import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:grpc/grpc.dart';
import 'package:grpc/grpc_connection_interface.dart';
import 'package:js/js.dart';
import 'package:cosmos_notifier/api/protobuf/dart/auth_service.pbgrpc.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/jwt_manager.dart';
import 'package:cosmos_notifier/f_home/services/type/login_data.dart';

@JS()
external String getTelegramInitData();

class AuthExpiredError implements Exception {
  String wdExpMsg() => 'AuthExpiredError';
}

class AuthUserNotFoundError implements Exception {
  String wdExpMsg() => 'AuthUserNotFoundError';
}

class AuthService extends AuthServiceClient with ChangeNotifier {
  static AuthService? _singleton;

  final Duration refreshBeforeExpDuration;
  final JwtManager jwtManager;

  factory AuthService(
          ClientChannelBase channel, Iterable<ClientInterceptor> interceptors, JwtManager jwtManager, refreshBeforeExpDuration) =>
      _singleton ??= AuthService._internal(channel, interceptors, jwtManager, refreshBeforeExpDuration);

  AuthService._internal(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors, this.jwtManager, this.refreshBeforeExpDuration)
      : super(channel, interceptors: interceptors);

  bool get hasLoginData => _isTelegramLogin() || _isDiscordLogin();

  // Authenticate using query params for Telegram/Discord login
  // It they are not available try to authenticate with previous saved JWT refresh token
  // If authentication not successful an error will be thrown
  init() async {
    if (canRefreshAccessToken) {
      try {
        await _login();
      } catch (e) {
        await _refreshAccessToken();
        if (!isAuthenticated) {
          rethrow;
        }
      }
    } else {
      await _login();
    }
    _scheduleRefreshAccessToken();
  }

  // Try to authenticate without login.
  // This works just in case there is already a JWT in local storage.
  Future<bool> backgroundInit() async {
    try {
      if (isAuthenticated) {
        _scheduleRefreshAccessToken();
      } else if (canRefreshAccessToken) {
        await _refreshAccessToken();
        if (isAuthenticated) {
          _scheduleRefreshAccessToken();
        }
      }
      // ignore: empty_catches
    } catch (e) {}
    return isAuthenticated;
  }

  bool get isAuthenticated {
    return jwtManager.isAccessTokenValid;
  }

  bool get canRefreshAccessToken {
    return jwtManager.isRefreshTokenValid;
  }

  TelegramLoginData _getTelegramLoginData() {
    final loginData = TelegramLoginData(Uri.base.queryParameters.entries.map((e) => "${e.key}=${e.value}").join("\n"));
    if (loginData.isValid) {
      return loginData;
    }
    final data = getTelegramInitData();
    if (data.isNotEmpty) {
      final loginData = TelegramLoginData(Uri.decodeComponent(data));
      if (loginData.isValid) {
        return loginData;
      }
    } else {
      throw Exception("Login information are not available");
    }
    throw Exception("Login information are not available");
  }

  DiscordLoginData _getDiscordLoginData() {
    final data = Uri.base.queryParameters.entries.map((e) => "${e.key}=${e.value}").join("\n");
    final loginData = DiscordLoginData(data);
    if (loginData.isValid) {
      return loginData;
    }
    throw Exception("Discord login information are not available");
  }

  bool _isDiscordLogin() {
    return Uri.base.queryParameters.containsKey("code");
  }

  bool _isTelegramLogin() {
    return Uri.base.queryParameters.containsKey("hash");
  }

  _login() async {
    if (cDebugMode) {
      print("AuthService: login");
    }
    try {
      if (_isDiscordLogin()) {
        final loginData = _getDiscordLoginData();
        var data = DiscordLoginRequest(code: loginData.code);
        var response = await discordLogin(data);
        jwtManager.accessToken = response.accessToken;
        jwtManager.refreshToken = response.refreshToken;
      } else if (_isTelegramLogin()) {
        final loginData = _getTelegramLoginData();
        var data = TelegramLoginRequest(
          userId: loginData.id,
          dataStr: loginData.data,
          username: loginData.username,
          authDate: loginData.authDate,
          hash: loginData.hash,
        );
        var response = await telegramLogin(data);
        jwtManager.accessToken = response.accessToken;
        jwtManager.refreshToken = response.refreshToken;
      } else {
        throw Exception("Login information are not available");
      }
    } on GrpcError catch (e) {
      if (e.code == StatusCode.notFound && e.message == "user not found") {
        throw AuthUserNotFoundError();
      } else if ((e).code == StatusCode.unauthenticated && e.message == "login expired") {
        throw AuthExpiredError();
      }
      rethrow;
    }
  }

  _logout() {
    if (cDebugMode) {
      print("AuthService: logout");
    }
    jwtManager.accessToken = "";
    jwtManager.refreshToken = "";
    notifyListeners();
  }

  Future<bool> _refreshAccessToken() async {
    if (cDebugMode) {
      print("AuthService: Refresh access token");
    }
    try {
      var response = await refreshAccessToken(RefreshAccessTokenRequest(refreshToken: jwtManager.refreshToken));
      jwtManager.accessToken = response.accessToken;
      return true;
    } catch (e) {
      if (cDebugMode) {
        print("AuthService: Error while refreshing access token: $e");
      }
    }
    return false;
  }

  _scheduleRefreshAccessToken() {
    try {
      final int exp = jwtManager.accessTokenDecoded["exp"] ?? 0;
      var expDateTime = DateTime.fromMillisecondsSinceEpoch(exp * 1000);
      var sleep = expDateTime.subtract(refreshBeforeExpDuration).difference(DateTime.now()).inSeconds;
      if (sleep <= 0) {
        sleep = 0;
      }
      if (cDebugMode) {
        print("AuthService: sleep for $sleep seconds");
      }
      Timer(Duration(seconds: sleep), () async {
        bool hasRefreshed = false;
        if (canRefreshAccessToken) {
          hasRefreshed = await _refreshAccessToken();
          _scheduleRefreshAccessToken();
        }
        if (!hasRefreshed) {
          _logout();
        }
      });
    } on FormatException catch (e) {
      if (cDebugMode) {
        print("AuthService: Error while executing _scheduleRefreshAccessToken: $e");
      }
      _logout();
      return;
    }
  }
}
