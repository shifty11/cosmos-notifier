///
//  Generated code. Do not modify.
//  source: tracker_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'google/protobuf/empty.pb.dart' as $1;
import 'tracker_service.pb.dart' as $4;
export 'tracker_service.pb.dart';

class TrackerServiceClient extends $grpc.Client {
  static final _$getTrackers =
      $grpc.ClientMethod<$1.Empty, $4.GetTrackersResponse>(
          '/cosmos_notifier_grpc.TrackerService/GetTrackers',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.GetTrackersResponse.fromBuffer(value));
  static final _$isAddressValid =
      $grpc.ClientMethod<$4.IsAddressValidRequest, $4.IsAddressValidResponse>(
          '/cosmos_notifier_grpc.TrackerService/IsAddressValid',
          ($4.IsAddressValidRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.IsAddressValidResponse.fromBuffer(value));
  static final _$addTracker =
      $grpc.ClientMethod<$4.AddTrackerRequest, $4.Tracker>(
          '/cosmos_notifier_grpc.TrackerService/AddTracker',
          ($4.AddTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $4.Tracker.fromBuffer(value));
  static final _$updateTracker =
      $grpc.ClientMethod<$4.UpdateTrackerRequest, $4.Tracker>(
          '/cosmos_notifier_grpc.TrackerService/UpdateTracker',
          ($4.UpdateTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $4.Tracker.fromBuffer(value));
  static final _$deleteTracker =
      $grpc.ClientMethod<$4.DeleteTrackerRequest, $1.Empty>(
          '/cosmos_notifier_grpc.TrackerService/DeleteTracker',
          ($4.DeleteTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  TrackerServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$4.GetTrackersResponse> getTrackers($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getTrackers, request, options: options);
  }

  $grpc.ResponseFuture<$4.IsAddressValidResponse> isAddressValid(
      $4.IsAddressValidRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$isAddressValid, request, options: options);
  }

  $grpc.ResponseFuture<$4.Tracker> addTracker($4.AddTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$addTracker, request, options: options);
  }

  $grpc.ResponseFuture<$4.Tracker> updateTracker(
      $4.UpdateTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateTracker, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> deleteTracker($4.DeleteTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteTracker, request, options: options);
  }
}

abstract class TrackerServiceBase extends $grpc.Service {
  $core.String get $name => 'cosmos_notifier_grpc.TrackerService';

  TrackerServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.Empty, $4.GetTrackersResponse>(
        'GetTrackers',
        getTrackers_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($4.GetTrackersResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.IsAddressValidRequest,
            $4.IsAddressValidResponse>(
        'IsAddressValid',
        isAddressValid_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.IsAddressValidRequest.fromBuffer(value),
        ($4.IsAddressValidResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.AddTrackerRequest, $4.Tracker>(
        'AddTracker',
        addTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.AddTrackerRequest.fromBuffer(value),
        ($4.Tracker value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.UpdateTrackerRequest, $4.Tracker>(
        'UpdateTracker',
        updateTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.UpdateTrackerRequest.fromBuffer(value),
        ($4.Tracker value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.DeleteTrackerRequest, $1.Empty>(
        'DeleteTracker',
        deleteTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.DeleteTrackerRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$4.GetTrackersResponse> getTrackers_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getTrackers(call, await request);
  }

  $async.Future<$4.IsAddressValidResponse> isAddressValid_Pre(
      $grpc.ServiceCall call,
      $async.Future<$4.IsAddressValidRequest> request) async {
    return isAddressValid(call, await request);
  }

  $async.Future<$4.Tracker> addTracker_Pre($grpc.ServiceCall call,
      $async.Future<$4.AddTrackerRequest> request) async {
    return addTracker(call, await request);
  }

  $async.Future<$4.Tracker> updateTracker_Pre($grpc.ServiceCall call,
      $async.Future<$4.UpdateTrackerRequest> request) async {
    return updateTracker(call, await request);
  }

  $async.Future<$1.Empty> deleteTracker_Pre($grpc.ServiceCall call,
      $async.Future<$4.DeleteTrackerRequest> request) async {
    return deleteTracker(call, await request);
  }

  $async.Future<$4.GetTrackersResponse> getTrackers(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$4.IsAddressValidResponse> isAddressValid(
      $grpc.ServiceCall call, $4.IsAddressValidRequest request);
  $async.Future<$4.Tracker> addTracker(
      $grpc.ServiceCall call, $4.AddTrackerRequest request);
  $async.Future<$4.Tracker> updateTracker(
      $grpc.ServiceCall call, $4.UpdateTrackerRequest request);
  $async.Future<$1.Empty> deleteTracker(
      $grpc.ServiceCall call, $4.DeleteTrackerRequest request);
}
