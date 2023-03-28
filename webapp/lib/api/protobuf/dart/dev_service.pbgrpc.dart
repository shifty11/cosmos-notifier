///
//  Generated code. Do not modify.
//  source: dev_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'dev_service.pb.dart' as $3;
import 'auth_service.pb.dart' as $2;
export 'dev_service.pb.dart';

class DevServiceClient extends $grpc.Client {
  static final _$login =
      $grpc.ClientMethod<$3.DevLoginRequest, $2.LoginResponse>(
          '/cosmos_notifier_grpc.DevService/Login',
          ($3.DevLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.LoginResponse.fromBuffer(value));

  DevServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$2.LoginResponse> login($3.DevLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$login, request, options: options);
  }
}

abstract class DevServiceBase extends $grpc.Service {
  $core.String get $name => 'cosmos_notifier_grpc.DevService';

  DevServiceBase() {
    $addMethod($grpc.ServiceMethod<$3.DevLoginRequest, $2.LoginResponse>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.DevLoginRequest.fromBuffer(value),
        ($2.LoginResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.LoginResponse> login_Pre(
      $grpc.ServiceCall call, $async.Future<$3.DevLoginRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$2.LoginResponse> login(
      $grpc.ServiceCall call, $3.DevLoginRequest request);
}
