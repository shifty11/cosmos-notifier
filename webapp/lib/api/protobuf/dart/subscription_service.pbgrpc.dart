///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'google/protobuf/empty.pb.dart' as $2;
import 'subscription_service.pb.dart' as $3;
export 'subscription_service.pb.dart';

class SubscriptionServiceClient extends $grpc.Client {
  static final _$getSubscriptions =
      $grpc.ClientMethod<$2.Empty, $3.GetSubscriptionsResponse>(
          '/daodao_notifier_grpc.SubscriptionService/GetSubscriptions',
          ($2.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $3.GetSubscriptionsResponse.fromBuffer(value));
  static final _$toggleChainSubscription = $grpc.ClientMethod<
          $3.ToggleChainSubscriptionRequest, $3.ToggleSubscriptionResponse>(
      '/daodao_notifier_grpc.SubscriptionService/ToggleChainSubscription',
      ($3.ToggleChainSubscriptionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $3.ToggleSubscriptionResponse.fromBuffer(value));
  static final _$toggleContractSubscription = $grpc.ClientMethod<
          $3.ToggleContractSubscriptionRequest, $3.ToggleSubscriptionResponse>(
      '/daodao_notifier_grpc.SubscriptionService/ToggleContractSubscription',
      ($3.ToggleContractSubscriptionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $3.ToggleSubscriptionResponse.fromBuffer(value));
  static final _$addDao =
      $grpc.ClientMethod<$3.AddDaoRequest, $3.AddDaoResponse>(
          '/daodao_notifier_grpc.SubscriptionService/AddDao',
          ($3.AddDaoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $3.AddDaoResponse.fromBuffer(value));
  static final _$deleteDao = $grpc.ClientMethod<$3.DeleteDaoRequest, $2.Empty>(
      '/daodao_notifier_grpc.SubscriptionService/DeleteDao',
      ($3.DeleteDaoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Empty.fromBuffer(value));
  static final _$enableChain =
      $grpc.ClientMethod<$3.EnableChainRequest, $2.Empty>(
          '/daodao_notifier_grpc.SubscriptionService/EnableChain',
          ($3.EnableChainRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.Empty.fromBuffer(value));

  SubscriptionServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$3.GetSubscriptionsResponse> getSubscriptions(
      $2.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getSubscriptions, request, options: options);
  }

  $grpc.ResponseFuture<$3.ToggleSubscriptionResponse> toggleChainSubscription(
      $3.ToggleChainSubscriptionRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleChainSubscription, request,
        options: options);
  }

  $grpc.ResponseFuture<$3.ToggleSubscriptionResponse>
      toggleContractSubscription($3.ToggleContractSubscriptionRequest request,
          {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleContractSubscription, request,
        options: options);
  }

  $grpc.ResponseStream<$3.AddDaoResponse> addDao($3.AddDaoRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$addDao, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$2.Empty> deleteDao($3.DeleteDaoRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteDao, request, options: options);
  }

  $grpc.ResponseFuture<$2.Empty> enableChain($3.EnableChainRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$enableChain, request, options: options);
  }
}

abstract class SubscriptionServiceBase extends $grpc.Service {
  $core.String get $name => 'daodao_notifier_grpc.SubscriptionService';

  SubscriptionServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.Empty, $3.GetSubscriptionsResponse>(
        'GetSubscriptions',
        getSubscriptions_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Empty.fromBuffer(value),
        ($3.GetSubscriptionsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.ToggleChainSubscriptionRequest,
            $3.ToggleSubscriptionResponse>(
        'ToggleChainSubscription',
        toggleChainSubscription_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $3.ToggleChainSubscriptionRequest.fromBuffer(value),
        ($3.ToggleSubscriptionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.ToggleContractSubscriptionRequest,
            $3.ToggleSubscriptionResponse>(
        'ToggleContractSubscription',
        toggleContractSubscription_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $3.ToggleContractSubscriptionRequest.fromBuffer(value),
        ($3.ToggleSubscriptionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.AddDaoRequest, $3.AddDaoResponse>(
        'AddDao',
        addDao_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $3.AddDaoRequest.fromBuffer(value),
        ($3.AddDaoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.DeleteDaoRequest, $2.Empty>(
        'DeleteDao',
        deleteDao_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.DeleteDaoRequest.fromBuffer(value),
        ($2.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.EnableChainRequest, $2.Empty>(
        'EnableChain',
        enableChain_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $3.EnableChainRequest.fromBuffer(value),
        ($2.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$3.GetSubscriptionsResponse> getSubscriptions_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Empty> request) async {
    return getSubscriptions(call, await request);
  }

  $async.Future<$3.ToggleSubscriptionResponse> toggleChainSubscription_Pre(
      $grpc.ServiceCall call,
      $async.Future<$3.ToggleChainSubscriptionRequest> request) async {
    return toggleChainSubscription(call, await request);
  }

  $async.Future<$3.ToggleSubscriptionResponse> toggleContractSubscription_Pre(
      $grpc.ServiceCall call,
      $async.Future<$3.ToggleContractSubscriptionRequest> request) async {
    return toggleContractSubscription(call, await request);
  }

  $async.Stream<$3.AddDaoResponse> addDao_Pre(
      $grpc.ServiceCall call, $async.Future<$3.AddDaoRequest> request) async* {
    yield* addDao(call, await request);
  }

  $async.Future<$2.Empty> deleteDao_Pre($grpc.ServiceCall call,
      $async.Future<$3.DeleteDaoRequest> request) async {
    return deleteDao(call, await request);
  }

  $async.Future<$2.Empty> enableChain_Pre($grpc.ServiceCall call,
      $async.Future<$3.EnableChainRequest> request) async {
    return enableChain(call, await request);
  }

  $async.Future<$3.GetSubscriptionsResponse> getSubscriptions(
      $grpc.ServiceCall call, $2.Empty request);
  $async.Future<$3.ToggleSubscriptionResponse> toggleChainSubscription(
      $grpc.ServiceCall call, $3.ToggleChainSubscriptionRequest request);
  $async.Future<$3.ToggleSubscriptionResponse> toggleContractSubscription(
      $grpc.ServiceCall call, $3.ToggleContractSubscriptionRequest request);
  $async.Stream<$3.AddDaoResponse> addDao(
      $grpc.ServiceCall call, $3.AddDaoRequest request);
  $async.Future<$2.Empty> deleteDao(
      $grpc.ServiceCall call, $3.DeleteDaoRequest request);
  $async.Future<$2.Empty> enableChain(
      $grpc.ServiceCall call, $3.EnableChainRequest request);
}
