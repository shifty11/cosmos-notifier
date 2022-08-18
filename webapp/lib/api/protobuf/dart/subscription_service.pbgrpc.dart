///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'google/protobuf/empty.pb.dart' as $1;
import 'subscription_service.pb.dart' as $2;
export 'subscription_service.pb.dart';

class SubscriptionServiceClient extends $grpc.Client {
  static final _$getSubscriptions =
      $grpc.ClientMethod<$1.Empty, $2.GetSubscriptionsResponse>(
          '/daodao_notifier_grpc.SubscriptionService/GetSubscriptions',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $2.GetSubscriptionsResponse.fromBuffer(value));
  static final _$toggleSubscription = $grpc.ClientMethod<
          $2.ToggleSubscriptionRequest, $2.ToggleSubscriptionResponse>(
      '/daodao_notifier_grpc.SubscriptionService/ToggleSubscription',
      ($2.ToggleSubscriptionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $2.ToggleSubscriptionResponse.fromBuffer(value));

  SubscriptionServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$2.GetSubscriptionsResponse> getSubscriptions(
      $1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getSubscriptions, request, options: options);
  }

  $grpc.ResponseFuture<$2.ToggleSubscriptionResponse> toggleSubscription(
      $2.ToggleSubscriptionRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleSubscription, request, options: options);
  }
}

abstract class SubscriptionServiceBase extends $grpc.Service {
  $core.String get $name => 'daodao_notifier_grpc.SubscriptionService';

  SubscriptionServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.GetSubscriptionsResponse>(
        'GetSubscriptions',
        getSubscriptions_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.GetSubscriptionsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.ToggleSubscriptionRequest,
            $2.ToggleSubscriptionResponse>(
        'ToggleSubscription',
        toggleSubscription_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $2.ToggleSubscriptionRequest.fromBuffer(value),
        ($2.ToggleSubscriptionResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.GetSubscriptionsResponse> getSubscriptions_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getSubscriptions(call, await request);
  }

  $async.Future<$2.ToggleSubscriptionResponse> toggleSubscription_Pre(
      $grpc.ServiceCall call,
      $async.Future<$2.ToggleSubscriptionRequest> request) async {
    return toggleSubscription(call, await request);
  }

  $async.Future<$2.GetSubscriptionsResponse> getSubscriptions(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$2.ToggleSubscriptionResponse> toggleSubscription(
      $grpc.ServiceCall call, $2.ToggleSubscriptionRequest request);
}
