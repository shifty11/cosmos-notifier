///
//  Generated code. Do not modify.
//  source: auth_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'auth_service.pb.dart' as $2;
import 'google/protobuf/empty.pb.dart' as $1;
export 'auth_service.pb.dart';

class AuthServiceClient extends $grpc.Client {
  static final _$telegramLogin =
      $grpc.ClientMethod<$2.TelegramLoginRequest, $2.LoginResponse>(
          '/cosmos_notifier_grpc.AuthService/TelegramLogin',
          ($2.TelegramLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.LoginResponse.fromBuffer(value));
  static final _$discordLogin =
      $grpc.ClientMethod<$2.DiscordLoginRequest, $2.LoginResponse>(
          '/cosmos_notifier_grpc.AuthService/DiscordLogin',
          ($2.DiscordLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.LoginResponse.fromBuffer(value));
  static final _$refreshAccessToken = $grpc.ClientMethod<
          $2.RefreshAccessTokenRequest, $2.RefreshAccessTokenResponse>(
      '/cosmos_notifier_grpc.AuthService/RefreshAccessToken',
      ($2.RefreshAccessTokenRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $2.RefreshAccessTokenResponse.fromBuffer(value));
  static final _$cannySSO = $grpc.ClientMethod<$1.Empty, $2.CannySSOResponse>(
      '/cosmos_notifier_grpc.AuthService/CannySSO',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.CannySSOResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$2.LoginResponse> telegramLogin(
      $2.TelegramLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$telegramLogin, request, options: options);
  }

  $grpc.ResponseFuture<$2.LoginResponse> discordLogin(
      $2.DiscordLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$discordLogin, request, options: options);
  }

  $grpc.ResponseFuture<$2.RefreshAccessTokenResponse> refreshAccessToken(
      $2.RefreshAccessTokenRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$refreshAccessToken, request, options: options);
  }

  $grpc.ResponseFuture<$2.CannySSOResponse> cannySSO($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$cannySSO, request, options: options);
  }
}

abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'cosmos_notifier_grpc.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.TelegramLoginRequest, $2.LoginResponse>(
        'TelegramLogin',
        telegramLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $2.TelegramLoginRequest.fromBuffer(value),
        ($2.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.DiscordLoginRequest, $2.LoginResponse>(
        'DiscordLogin',
        discordLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $2.DiscordLoginRequest.fromBuffer(value),
        ($2.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.RefreshAccessTokenRequest,
            $2.RefreshAccessTokenResponse>(
        'RefreshAccessToken',
        refreshAccessToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $2.RefreshAccessTokenRequest.fromBuffer(value),
        ($2.RefreshAccessTokenResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.CannySSOResponse>(
        'CannySSO',
        cannySSO_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.CannySSOResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.LoginResponse> telegramLogin_Pre($grpc.ServiceCall call,
      $async.Future<$2.TelegramLoginRequest> request) async {
    return telegramLogin(call, await request);
  }

  $async.Future<$2.LoginResponse> discordLogin_Pre($grpc.ServiceCall call,
      $async.Future<$2.DiscordLoginRequest> request) async {
    return discordLogin(call, await request);
  }

  $async.Future<$2.RefreshAccessTokenResponse> refreshAccessToken_Pre(
      $grpc.ServiceCall call,
      $async.Future<$2.RefreshAccessTokenRequest> request) async {
    return refreshAccessToken(call, await request);
  }

  $async.Future<$2.CannySSOResponse> cannySSO_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return cannySSO(call, await request);
  }

  $async.Future<$2.LoginResponse> telegramLogin(
      $grpc.ServiceCall call, $2.TelegramLoginRequest request);
  $async.Future<$2.LoginResponse> discordLogin(
      $grpc.ServiceCall call, $2.DiscordLoginRequest request);
  $async.Future<$2.RefreshAccessTokenResponse> refreshAccessToken(
      $grpc.ServiceCall call, $2.RefreshAccessTokenRequest request);
  $async.Future<$2.CannySSOResponse> cannySSO(
      $grpc.ServiceCall call, $1.Empty request);
}
