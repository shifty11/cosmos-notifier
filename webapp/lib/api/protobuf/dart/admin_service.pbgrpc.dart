///
//  Generated code. Do not modify.
//  source: admin_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'admin_service.pb.dart' as $0;
export 'admin_service.pb.dart';

class AdminServiceClient extends $grpc.Client {
  static final _$broadcastMessage = $grpc.ClientMethod<
          $0.BroadcastMessageRequest, $0.BroadcastMessageResponse>(
      '/daodao_notifier_grpc.AdminService/BroadcastMessage',
      ($0.BroadcastMessageRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.BroadcastMessageResponse.fromBuffer(value));

  AdminServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseStream<$0.BroadcastMessageResponse> broadcastMessage(
      $0.BroadcastMessageRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$broadcastMessage, $async.Stream.fromIterable([request]),
        options: options);
  }
}

abstract class AdminServiceBase extends $grpc.Service {
  $core.String get $name => 'daodao_notifier_grpc.AdminService';

  AdminServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.BroadcastMessageRequest,
            $0.BroadcastMessageResponse>(
        'BroadcastMessage',
        broadcastMessage_Pre,
        false,
        true,
        ($core.List<$core.int> value) =>
            $0.BroadcastMessageRequest.fromBuffer(value),
        ($0.BroadcastMessageResponse value) => value.writeToBuffer()));
  }

  $async.Stream<$0.BroadcastMessageResponse> broadcastMessage_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.BroadcastMessageRequest> request) async* {
    yield* broadcastMessage(call, await request);
  }

  $async.Stream<$0.BroadcastMessageResponse> broadcastMessage(
      $grpc.ServiceCall call, $0.BroadcastMessageRequest request);
}
