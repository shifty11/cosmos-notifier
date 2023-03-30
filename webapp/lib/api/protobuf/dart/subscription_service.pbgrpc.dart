///
//  Generated code. Do not modify.
//  source: subscription_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'google/protobuf/empty.pb.dart' as $1;
import 'subscription_service.pb.dart' as $4;
export 'subscription_service.pb.dart';

class SubscriptionServiceClient extends $grpc.Client {
  static final _$getSubscriptions =
      $grpc.ClientMethod<$1.Empty, $4.GetSubscriptionsResponse>(
          '/cosmos_notifier_grpc.SubscriptionService/GetSubscriptions',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.GetSubscriptionsResponse.fromBuffer(value));
  static final _$toggleChainSubscription = $grpc.ClientMethod<
          $4.ToggleChainSubscriptionRequest, $4.ToggleSubscriptionResponse>(
      '/cosmos_notifier_grpc.SubscriptionService/ToggleChainSubscription',
      ($4.ToggleChainSubscriptionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $4.ToggleSubscriptionResponse.fromBuffer(value));
  static final _$toggleContractSubscription = $grpc.ClientMethod<
          $4.ToggleContractSubscriptionRequest, $4.ToggleSubscriptionResponse>(
      '/cosmos_notifier_grpc.SubscriptionService/ToggleContractSubscription',
      ($4.ToggleContractSubscriptionRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $4.ToggleSubscriptionResponse.fromBuffer(value));
  static final _$addDao =
      $grpc.ClientMethod<$4.AddDaoRequest, $4.AddDaoResponse>(
          '/cosmos_notifier_grpc.SubscriptionService/AddDao',
          ($4.AddDaoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $4.AddDaoResponse.fromBuffer(value));
  static final _$deleteDao = $grpc.ClientMethod<$4.DeleteDaoRequest, $1.Empty>(
      '/cosmos_notifier_grpc.SubscriptionService/DeleteDao',
      ($4.DeleteDaoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$enableChain =
      $grpc.ClientMethod<$4.EnableChainRequest, $1.Empty>(
          '/cosmos_notifier_grpc.SubscriptionService/EnableChain',
          ($4.EnableChainRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  SubscriptionServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$4.GetSubscriptionsResponse> getSubscriptions(
      $1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getSubscriptions, request, options: options);
  }

  $grpc.ResponseFuture<$4.ToggleSubscriptionResponse> toggleChainSubscription(
      $4.ToggleChainSubscriptionRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleChainSubscription, request,
        options: options);
  }

  $grpc.ResponseFuture<$4.ToggleSubscriptionResponse>
      toggleContractSubscription($4.ToggleContractSubscriptionRequest request,
          {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleContractSubscription, request,
        options: options);
  }

  $grpc.ResponseStream<$4.AddDaoResponse> addDao($4.AddDaoRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$addDao, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$1.Empty> deleteDao($4.DeleteDaoRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteDao, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> enableChain($4.EnableChainRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$enableChain, request, options: options);
  }
}

abstract class SubscriptionServiceBase extends $grpc.Service {
  $core.String get $name => 'cosmos_notifier_grpc.SubscriptionService';

  SubscriptionServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.Empty, $4.GetSubscriptionsResponse>(
        'GetSubscriptions',
        getSubscriptions_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($4.GetSubscriptionsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.ToggleChainSubscriptionRequest,
            $4.ToggleSubscriptionResponse>(
        'ToggleChainSubscription',
        toggleChainSubscription_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.ToggleChainSubscriptionRequest.fromBuffer(value),
        ($4.ToggleSubscriptionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.ToggleContractSubscriptionRequest,
            $4.ToggleSubscriptionResponse>(
        'ToggleContractSubscription',
        toggleContractSubscription_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.ToggleContractSubscriptionRequest.fromBuffer(value),
        ($4.ToggleSubscriptionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.AddDaoRequest, $4.AddDaoResponse>(
        'AddDao',
        addDao_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $4.AddDaoRequest.fromBuffer(value),
        ($4.AddDaoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.DeleteDaoRequest, $1.Empty>(
        'DeleteDao',
        deleteDao_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.DeleteDaoRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.EnableChainRequest, $1.Empty>(
        'EnableChain',
        enableChain_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.EnableChainRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$4.GetSubscriptionsResponse> getSubscriptions_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getSubscriptions(call, await request);
  }

  $async.Future<$4.ToggleSubscriptionResponse> toggleChainSubscription_Pre(
      $grpc.ServiceCall call,
      $async.Future<$4.ToggleChainSubscriptionRequest> request) async {
    return toggleChainSubscription(call, await request);
  }

  $async.Future<$4.ToggleSubscriptionResponse> toggleContractSubscription_Pre(
      $grpc.ServiceCall call,
      $async.Future<$4.ToggleContractSubscriptionRequest> request) async {
    return toggleContractSubscription(call, await request);
  }

  $async.Stream<$4.AddDaoResponse> addDao_Pre(
      $grpc.ServiceCall call, $async.Future<$4.AddDaoRequest> request) async* {
    yield* addDao(call, await request);
  }

  $async.Future<$1.Empty> deleteDao_Pre($grpc.ServiceCall call,
      $async.Future<$4.DeleteDaoRequest> request) async {
    return deleteDao(call, await request);
  }

  $async.Future<$1.Empty> enableChain_Pre($grpc.ServiceCall call,
      $async.Future<$4.EnableChainRequest> request) async {
    return enableChain(call, await request);
  }

  $async.Future<$4.GetSubscriptionsResponse> getSubscriptions(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$4.ToggleSubscriptionResponse> toggleChainSubscription(
      $grpc.ServiceCall call, $4.ToggleChainSubscriptionRequest request);
  $async.Future<$4.ToggleSubscriptionResponse> toggleContractSubscription(
      $grpc.ServiceCall call, $4.ToggleContractSubscriptionRequest request);
  $async.Stream<$4.AddDaoResponse> addDao(
      $grpc.ServiceCall call, $4.AddDaoRequest request);
  $async.Future<$1.Empty> deleteDao(
      $grpc.ServiceCall call, $4.DeleteDaoRequest request);
  $async.Future<$1.Empty> enableChain(
      $grpc.ServiceCall call, $4.EnableChainRequest request);
}
