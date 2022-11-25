import 'dart:async';

import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/jwt_manager.dart';
import 'package:grpc/grpc.dart';

class AuthInterceptor extends ClientInterceptor {
  static AuthInterceptor? _singleton;
  JwtManager jwtManager;

  String? accessToken;

  String? refreshToken;

  factory AuthInterceptor(JwtManager jwtManager) => _singleton ??= AuthInterceptor._internal(jwtManager);

  AuthInterceptor._internal(this.jwtManager);

  FutureOr<void> _attachToken(Map<String, String> metadata, String uri) async {
    final token = jwtManager.accessToken;
    metadata['Authorization'] = "Bearer $token";
  }

  @override
  ResponseFuture<R> interceptUnary<Q, R>(ClientMethod<Q, R> method, Q request, CallOptions options, invoker) {
    if (cDebugMode) {
      print("interceptUnary --> ${method.path}");
    }

    if (isAuthNeeded(method.path)) {
      options = options.mergedWith(CallOptions(providers: [_attachToken]));
    }
    return super.interceptUnary(method, request, options, invoker);
  }

  @override
  ResponseStream<R> interceptStreaming<Q, R>(
      ClientMethod<Q, R> method, Stream<Q> requests, CallOptions options, ClientStreamingInvoker<Q, R> invoker) {
    if (cDebugMode) {
      print("interceptStreaming --> ${method.path}");
    }

    if (isAuthNeeded(method.path)) {
      options = options.mergedWith(CallOptions(providers: [_attachToken]));
    }

    return super.interceptStreaming(method, requests, options, invoker);
  }

  Map<String, bool> authMethod() {
    const path = "/cosmos_notifier_grpc.AuthService/";

    return {"${path}TelegramLogin": false, "${path}TokenLogin": false, "${path}RefreshAccessToken": false};
  }

  bool isAuthNeeded(String method) {
    var authInfo = authMethod()[method];
    if (authInfo == null) {
      return true;
    }
    return authInfo;
  }
}
