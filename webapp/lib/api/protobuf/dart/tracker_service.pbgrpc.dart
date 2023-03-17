///
//  Generated code. Do not modify.
//  source: tracker_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'tracker_service.pb.dart' as $4;
import 'google/protobuf/empty.pb.dart' as $1;
export 'tracker_service.pb.dart';

class TrackerServiceClient extends $grpc.Client {
  static final _$isAddressValid =
      $grpc.ClientMethod<$4.IsAddressValidRequest, $4.IsAddressValidResponse>(
          '/cosmos_notifier_grpc.TrackerService/IsAddressValid',
          ($4.IsAddressValidRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.IsAddressValidResponse.fromBuffer(value));
  static final _$addTracker =
      $grpc.ClientMethod<$4.AddTrackerRequest, $4.AddTrackerResponse>(
          '/cosmos_notifier_grpc.TrackerService/AddTracker',
          ($4.AddTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.AddTrackerResponse.fromBuffer(value));
  static final _$deleteTracker =
      $grpc.ClientMethod<$4.DeleteTrackerRequest, $1.Empty>(
          '/cosmos_notifier_grpc.TrackerService/DeleteTracker',
          ($4.DeleteTrackerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  TrackerServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$4.IsAddressValidResponse> isAddressValid(
      $4.IsAddressValidRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$isAddressValid, request, options: options);
  }

  $grpc.ResponseFuture<$4.AddTrackerResponse> addTracker(
      $4.AddTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$addTracker, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> deleteTracker($4.DeleteTrackerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteTracker, request, options: options);
  }
}

abstract class TrackerServiceBase extends $grpc.Service {
  $core.String get $name => 'cosmos_notifier_grpc.TrackerService';

  TrackerServiceBase() {
    $addMethod($grpc.ServiceMethod<$4.IsAddressValidRequest,
            $4.IsAddressValidResponse>(
        'IsAddressValid',
        isAddressValid_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.IsAddressValidRequest.fromBuffer(value),
        ($4.IsAddressValidResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.AddTrackerRequest, $4.AddTrackerResponse>(
        'AddTracker',
        addTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.AddTrackerRequest.fromBuffer(value),
        ($4.AddTrackerResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.DeleteTrackerRequest, $1.Empty>(
        'DeleteTracker',
        deleteTracker_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.DeleteTrackerRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$4.IsAddressValidResponse> isAddressValid_Pre(
      $grpc.ServiceCall call,
      $async.Future<$4.IsAddressValidRequest> request) async {
    return isAddressValid(call, await request);
  }

  $async.Future<$4.AddTrackerResponse> addTracker_Pre($grpc.ServiceCall call,
      $async.Future<$4.AddTrackerRequest> request) async {
    return addTracker(call, await request);
  }

  $async.Future<$1.Empty> deleteTracker_Pre($grpc.ServiceCall call,
      $async.Future<$4.DeleteTrackerRequest> request) async {
    return deleteTracker(call, await request);
  }

  $async.Future<$4.IsAddressValidResponse> isAddressValid(
      $grpc.ServiceCall call, $4.IsAddressValidRequest request);
  $async.Future<$4.AddTrackerResponse> addTracker(
      $grpc.ServiceCall call, $4.AddTrackerRequest request);
  $async.Future<$1.Empty> deleteTracker(
      $grpc.ServiceCall call, $4.DeleteTrackerRequest request);
}
