// ignore: avoid_web_libraries_in_flutter
import 'dart:html' show window;

import 'package:jwt_decode/jwt_decode.dart';

class JwtManager {
  static JwtManager? _singleton;

  String _accessToken = "";
  String _refreshToken = "";
  final storage = window.localStorage;

  factory JwtManager() => _singleton ??= JwtManager._internal();

  JwtManager._internal() {
    _accessToken = storage["accessToken"] ?? "";
    _refreshToken = storage["refreshToken"] ?? "";
  }

  set accessToken(String accessToken) {
    _accessToken = accessToken;
    if (accessToken.isEmpty) {
      storage.remove("accessToken");
    } else {
      storage["accessToken"] = accessToken;
    }
  }

  String get accessToken {
    return _accessToken;
  }

  bool get isAccessTokenValid {
    try {
      return _accessToken.isNotEmpty && !Jwt.isExpired(_accessToken);
    } on FormatException {
      return false;
    }
  }

  Map<String, dynamic> get accessTokenDecoded {
    return Jwt.parseJwt(_accessToken);
  }

  set refreshToken(String refreshToken) {
    _refreshToken = refreshToken;
    if (refreshToken.isEmpty) {
      storage.remove("refreshToken");
    } else {
      storage["refreshToken"] = refreshToken;
    }
  }

  String get refreshToken {
    return _refreshToken;
  }

  bool get isRefreshTokenValid {
    try {
      return _refreshToken.isNotEmpty && !Jwt.isExpired(_refreshToken);
    } on FormatException {
      return false;
    }
  }

  Map<String, dynamic> get refreshTokenDecoded {
    return Jwt.parseJwt(_refreshToken);
  }

  bool get isAdmin {
    return (accessTokenDecoded["role"] ?? "").toString().toLowerCase() == "admin";
  }

  bool get isTelegramUser {
    return (accessTokenDecoded["type"] ?? "").toString().toLowerCase() == "telegram";
  }

  bool get isDiscordUser {
    return (accessTokenDecoded["type"] ?? "").toString().toLowerCase() == "discord";
  }
}
