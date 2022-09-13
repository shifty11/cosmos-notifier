///
//  Generated code. Do not modify.
//  source: auth_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'auth_service.pb.dart' as $0;
export 'auth_service.pb.dart';

class AuthServiceClient extends $grpc.Client {
  static final _$telegramLogin =
      $grpc.ClientMethod<$0.TelegramLoginRequest, $0.LoginResponse>(
          '/daodao_notifier_grpc.AuthService/TelegramLogin',
          ($0.TelegramLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.LoginResponse.fromBuffer(value));
  static final _$discordLogin =
      $grpc.ClientMethod<$0.DiscordLoginRequest, $0.LoginResponse>(
          '/daodao_notifier_grpc.AuthService/DiscordLogin',
          ($0.DiscordLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.LoginResponse.fromBuffer(value));
  static final _$refreshAccessToken = $grpc.ClientMethod<
          $0.RefreshAccessTokenRequest, $0.RefreshAccessTokenResponse>(
      '/daodao_notifier_grpc.AuthService/RefreshAccessToken',
      ($0.RefreshAccessTokenRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.RefreshAccessTokenResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.LoginResponse> telegramLogin(
      $0.TelegramLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$telegramLogin, request, options: options);
  }

  $grpc.ResponseFuture<$0.LoginResponse> discordLogin(
      $0.DiscordLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$discordLogin, request, options: options);
  }

  $grpc.ResponseFuture<$0.RefreshAccessTokenResponse> refreshAccessToken(
      $0.RefreshAccessTokenRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$refreshAccessToken, request, options: options);
  }
}

abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'daodao_notifier_grpc.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.TelegramLoginRequest, $0.LoginResponse>(
        'TelegramLogin',
        telegramLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.TelegramLoginRequest.fromBuffer(value),
        ($0.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DiscordLoginRequest, $0.LoginResponse>(
        'DiscordLogin',
        discordLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.DiscordLoginRequest.fromBuffer(value),
        ($0.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RefreshAccessTokenRequest,
            $0.RefreshAccessTokenResponse>(
        'RefreshAccessToken',
        refreshAccessToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.RefreshAccessTokenRequest.fromBuffer(value),
        ($0.RefreshAccessTokenResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.LoginResponse> telegramLogin_Pre($grpc.ServiceCall call,
      $async.Future<$0.TelegramLoginRequest> request) async {
    return telegramLogin(call, await request);
  }

  $async.Future<$0.LoginResponse> discordLogin_Pre($grpc.ServiceCall call,
      $async.Future<$0.DiscordLoginRequest> request) async {
    return discordLogin(call, await request);
  }

  $async.Future<$0.RefreshAccessTokenResponse> refreshAccessToken_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.RefreshAccessTokenRequest> request) async {
    return refreshAccessToken(call, await request);
  }

  $async.Future<$0.LoginResponse> telegramLogin(
      $grpc.ServiceCall call, $0.TelegramLoginRequest request);
  $async.Future<$0.LoginResponse> discordLogin(
      $grpc.ServiceCall call, $0.DiscordLoginRequest request);
  $async.Future<$0.RefreshAccessTokenResponse> refreshAccessToken(
      $grpc.ServiceCall call, $0.RefreshAccessTokenRequest request);
}
