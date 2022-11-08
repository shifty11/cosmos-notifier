///
//  Generated code. Do not modify.
//  source: auth_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'auth_service.pb.dart' as $1;
export 'auth_service.pb.dart';

class AuthServiceClient extends $grpc.Client {
  static final _$telegramLogin =
      $grpc.ClientMethod<$1.TelegramLoginRequest, $1.LoginResponse>(
          '/daodao_notifier_grpc.AuthService/TelegramLogin',
          ($1.TelegramLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.LoginResponse.fromBuffer(value));
  static final _$discordLogin =
      $grpc.ClientMethod<$1.DiscordLoginRequest, $1.LoginResponse>(
          '/daodao_notifier_grpc.AuthService/DiscordLogin',
          ($1.DiscordLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.LoginResponse.fromBuffer(value));
  static final _$refreshAccessToken = $grpc.ClientMethod<
          $1.RefreshAccessTokenRequest, $1.RefreshAccessTokenResponse>(
      '/daodao_notifier_grpc.AuthService/RefreshAccessToken',
      ($1.RefreshAccessTokenRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.RefreshAccessTokenResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.LoginResponse> telegramLogin(
      $1.TelegramLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$telegramLogin, request, options: options);
  }

  $grpc.ResponseFuture<$1.LoginResponse> discordLogin(
      $1.DiscordLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$discordLogin, request, options: options);
  }

  $grpc.ResponseFuture<$1.RefreshAccessTokenResponse> refreshAccessToken(
      $1.RefreshAccessTokenRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$refreshAccessToken, request, options: options);
  }
}

abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'daodao_notifier_grpc.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.TelegramLoginRequest, $1.LoginResponse>(
        'TelegramLogin',
        telegramLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.TelegramLoginRequest.fromBuffer(value),
        ($1.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.DiscordLoginRequest, $1.LoginResponse>(
        'DiscordLogin',
        discordLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.DiscordLoginRequest.fromBuffer(value),
        ($1.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.RefreshAccessTokenRequest,
            $1.RefreshAccessTokenResponse>(
        'RefreshAccessToken',
        refreshAccessToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.RefreshAccessTokenRequest.fromBuffer(value),
        ($1.RefreshAccessTokenResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.LoginResponse> telegramLogin_Pre($grpc.ServiceCall call,
      $async.Future<$1.TelegramLoginRequest> request) async {
    return telegramLogin(call, await request);
  }

  $async.Future<$1.LoginResponse> discordLogin_Pre($grpc.ServiceCall call,
      $async.Future<$1.DiscordLoginRequest> request) async {
    return discordLogin(call, await request);
  }

  $async.Future<$1.RefreshAccessTokenResponse> refreshAccessToken_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.RefreshAccessTokenRequest> request) async {
    return refreshAccessToken(call, await request);
  }

  $async.Future<$1.LoginResponse> telegramLogin(
      $grpc.ServiceCall call, $1.TelegramLoginRequest request);
  $async.Future<$1.LoginResponse> discordLogin(
      $grpc.ServiceCall call, $1.DiscordLoginRequest request);
  $async.Future<$1.RefreshAccessTokenResponse> refreshAccessToken(
      $grpc.ServiceCall call, $1.RefreshAccessTokenRequest request);
}
