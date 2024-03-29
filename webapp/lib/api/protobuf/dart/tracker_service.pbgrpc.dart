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
import 'tracker_service.pb.dart' as $5;
export 'tracker_service.pb.dart';

class TrackerServiceClient extends $grpc.Client {
  static final _$listTrackers =
      $grpc.ClientMethod<$1.Empty, $5.ListTrackersResponse>(
          '/cosmos_notifier_grpc.TrackerService/ListTrackers',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $5.ListTrackersResponse.fromBuffer(value));
  static final _$isAddressValid =
      $grpc.ClientMethod<$5.IsAddressValidRequest, $5.IsAddressValidResponse>(
          '/cosmos_notifier_grpc.TrackerService/IsAddressValid',
          ($5.IsAddressValidRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $5.IsAddressValidResponse.fromBuffer(value));
  static final _$createTracker =
      $grpc.ClientMethod<$5.CreateTrackerRequest, $5.Tracker>(
          '/cosmos_notifier_grpc.TrackerService/CreateTracker',
          ($5.CreateTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $5.Tracker.fromBuffer(value));
  static final _$updateTracker =
      $grpc.ClientMethod<$5.UpdateTrackerRequest, $5.Tracker>(
          '/cosmos_notifier_grpc.TrackerService/UpdateTracker',
          ($5.UpdateTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $5.Tracker.fromBuffer(value));
  static final _$deleteTracker =
      $grpc.ClientMethod<$5.DeleteTrackerRequest, $1.Empty>(
          '/cosmos_notifier_grpc.TrackerService/DeleteTracker',
          ($5.DeleteTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$trackValidators =
      $grpc.ClientMethod<$5.TrackValidatorsRequest, $5.TrackValidatorsResponse>(
          '/cosmos_notifier_grpc.TrackerService/TrackValidators',
          ($5.TrackValidatorsRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $5.TrackValidatorsResponse.fromBuffer(value));

  TrackerServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$5.ListTrackersResponse> listTrackers($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listTrackers, request, options: options);
  }

  $grpc.ResponseFuture<$5.IsAddressValidResponse> isAddressValid(
      $5.IsAddressValidRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$isAddressValid, request, options: options);
  }

  $grpc.ResponseFuture<$5.Tracker> createTracker(
      $5.CreateTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createTracker, request, options: options);
  }

  $grpc.ResponseFuture<$5.Tracker> updateTracker(
      $5.UpdateTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateTracker, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> deleteTracker($5.DeleteTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteTracker, request, options: options);
  }

  $grpc.ResponseFuture<$5.TrackValidatorsResponse> trackValidators(
      $5.TrackValidatorsRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$trackValidators, request, options: options);
  }
}

abstract class TrackerServiceBase extends $grpc.Service {
  $core.String get $name => 'cosmos_notifier_grpc.TrackerService';

  TrackerServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.Empty, $5.ListTrackersResponse>(
        'ListTrackers',
        listTrackers_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($5.ListTrackersResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.IsAddressValidRequest,
            $5.IsAddressValidResponse>(
        'IsAddressValid',
        isAddressValid_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.IsAddressValidRequest.fromBuffer(value),
        ($5.IsAddressValidResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.CreateTrackerRequest, $5.Tracker>(
        'CreateTracker',
        createTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.CreateTrackerRequest.fromBuffer(value),
        ($5.Tracker value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.UpdateTrackerRequest, $5.Tracker>(
        'UpdateTracker',
        updateTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.UpdateTrackerRequest.fromBuffer(value),
        ($5.Tracker value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.DeleteTrackerRequest, $1.Empty>(
        'DeleteTracker',
        deleteTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.DeleteTrackerRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.TrackValidatorsRequest,
            $5.TrackValidatorsResponse>(
        'TrackValidators',
        trackValidators_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.TrackValidatorsRequest.fromBuffer(value),
        ($5.TrackValidatorsResponse value) => value.writeToBuffer()));
  }

  $async.Future<$5.ListTrackersResponse> listTrackers_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return listTrackers(call, await request);
  }

  $async.Future<$5.IsAddressValidResponse> isAddressValid_Pre(
      $grpc.ServiceCall call,
      $async.Future<$5.IsAddressValidRequest> request) async {
    return isAddressValid(call, await request);
  }

  $async.Future<$5.Tracker> createTracker_Pre($grpc.ServiceCall call,
      $async.Future<$5.CreateTrackerRequest> request) async {
    return createTracker(call, await request);
  }

  $async.Future<$5.Tracker> updateTracker_Pre($grpc.ServiceCall call,
      $async.Future<$5.UpdateTrackerRequest> request) async {
    return updateTracker(call, await request);
  }

  $async.Future<$1.Empty> deleteTracker_Pre($grpc.ServiceCall call,
      $async.Future<$5.DeleteTrackerRequest> request) async {
    return deleteTracker(call, await request);
  }

  $async.Future<$5.TrackValidatorsResponse> trackValidators_Pre(
      $grpc.ServiceCall call,
      $async.Future<$5.TrackValidatorsRequest> request) async {
    return trackValidators(call, await request);
  }

  $async.Future<$5.ListTrackersResponse> listTrackers(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$5.IsAddressValidResponse> isAddressValid(
      $grpc.ServiceCall call, $5.IsAddressValidRequest request);
  $async.Future<$5.Tracker> createTracker(
      $grpc.ServiceCall call, $5.CreateTrackerRequest request);
  $async.Future<$5.Tracker> updateTracker(
      $grpc.ServiceCall call, $5.UpdateTrackerRequest request);
  $async.Future<$1.Empty> deleteTracker(
      $grpc.ServiceCall call, $5.DeleteTrackerRequest request);
  $async.Future<$5.TrackValidatorsResponse> trackValidators(
      $grpc.ServiceCall call, $5.TrackValidatorsRequest request);
}
