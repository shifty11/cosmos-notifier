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
  static final _$toggleChainSubscription = $grpc.ClientMethod<
          $2.ToggleChainSubscriptionRequest, $2.ToggleSubscriptionResponse>(
      '/daodao_notifier_grpc.SubscriptionService/ToggleChainSubscription',
      ($2.ToggleChainSubscriptionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $2.ToggleSubscriptionResponse.fromBuffer(value));
  static final _$toggleContractSubscription = $grpc.ClientMethod<
          $2.ToggleContractSubscriptionRequest, $2.ToggleSubscriptionResponse>(
      '/daodao_notifier_grpc.SubscriptionService/ToggleContractSubscription',
      ($2.ToggleContractSubscriptionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $2.ToggleSubscriptionResponse.fromBuffer(value));
  static final _$addDao =
      $grpc.ClientMethod<$2.AddDaoRequest, $2.AddDaoResponse>(
          '/daodao_notifier_grpc.SubscriptionService/AddDao',
          ($2.AddDaoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.AddDaoResponse.fromBuffer(value));
  static final _$deleteDao = $grpc.ClientMethod<$2.DeleteDaoRequest, $1.Empty>(
      '/daodao_notifier_grpc.SubscriptionService/DeleteDao',
      ($2.DeleteDaoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$enableChain =
      $grpc.ClientMethod<$2.EnableChainRequest, $1.Empty>(
          '/daodao_notifier_grpc.SubscriptionService/EnableChain',
          ($2.EnableChainRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  SubscriptionServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$2.GetSubscriptionsResponse> getSubscriptions(
      $1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getSubscriptions, request, options: options);
  }

  $grpc.ResponseFuture<$2.ToggleSubscriptionResponse> toggleChainSubscription(
      $2.ToggleChainSubscriptionRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleChainSubscription, request,
        options: options);
  }

  $grpc.ResponseFuture<$2.ToggleSubscriptionResponse>
      toggleContractSubscription($2.ToggleContractSubscriptionRequest request,
          {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleContractSubscription, request,
        options: options);
  }

  $grpc.ResponseStream<$2.AddDaoResponse> addDao($2.AddDaoRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$addDao, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$1.Empty> deleteDao($2.DeleteDaoRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteDao, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> enableChain($2.EnableChainRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$enableChain, request, options: options);
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
    $addMethod($grpc.ServiceMethod<$2.ToggleChainSubscriptionRequest,
            $2.ToggleSubscriptionResponse>(
        'ToggleChainSubscription',
        toggleChainSubscription_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $2.ToggleChainSubscriptionRequest.fromBuffer(value),
        ($2.ToggleSubscriptionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.ToggleContractSubscriptionRequest,
            $2.ToggleSubscriptionResponse>(
        'ToggleContractSubscription',
        toggleContractSubscription_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $2.ToggleContractSubscriptionRequest.fromBuffer(value),
        ($2.ToggleSubscriptionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.AddDaoRequest, $2.AddDaoResponse>(
        'AddDao',
        addDao_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $2.AddDaoRequest.fromBuffer(value),
        ($2.AddDaoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.DeleteDaoRequest, $1.Empty>(
        'DeleteDao',
        deleteDao_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.DeleteDaoRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.EnableChainRequest, $1.Empty>(
        'EnableChain',
        enableChain_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $2.EnableChainRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$2.GetSubscriptionsResponse> getSubscriptions_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getSubscriptions(call, await request);
  }

  $async.Future<$2.ToggleSubscriptionResponse> toggleChainSubscription_Pre(
      $grpc.ServiceCall call,
      $async.Future<$2.ToggleChainSubscriptionRequest> request) async {
    return toggleChainSubscription(call, await request);
  }

  $async.Future<$2.ToggleSubscriptionResponse> toggleContractSubscription_Pre(
      $grpc.ServiceCall call,
      $async.Future<$2.ToggleContractSubscriptionRequest> request) async {
    return toggleContractSubscription(call, await request);
  }

  $async.Stream<$2.AddDaoResponse> addDao_Pre(
      $grpc.ServiceCall call, $async.Future<$2.AddDaoRequest> request) async* {
    yield* addDao(call, await request);
  }

  $async.Future<$1.Empty> deleteDao_Pre($grpc.ServiceCall call,
      $async.Future<$2.DeleteDaoRequest> request) async {
    return deleteDao(call, await request);
  }

  $async.Future<$1.Empty> enableChain_Pre($grpc.ServiceCall call,
      $async.Future<$2.EnableChainRequest> request) async {
    return enableChain(call, await request);
  }

  $async.Future<$2.GetSubscriptionsResponse> getSubscriptions(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$2.ToggleSubscriptionResponse> toggleChainSubscription(
      $grpc.ServiceCall call, $2.ToggleChainSubscriptionRequest request);
  $async.Future<$2.ToggleSubscriptionResponse> toggleContractSubscription(
      $grpc.ServiceCall call, $2.ToggleContractSubscriptionRequest request);
  $async.Stream<$2.AddDaoResponse> addDao(
      $grpc.ServiceCall call, $2.AddDaoRequest request);
  $async.Future<$1.Empty> deleteDao(
      $grpc.ServiceCall call, $2.DeleteDaoRequest request);
  $async.Future<$1.Empty> enableChain(
      $grpc.ServiceCall call, $2.EnableChainRequest request);
}
