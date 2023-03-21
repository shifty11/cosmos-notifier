import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pbgrpc.dart';
import 'package:grpc/grpc.dart';
import 'package:grpc/grpc_connection_interface.dart';

class TrackerService extends TrackerServiceClient {
  static TrackerService? _singleton;

  factory TrackerService(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors) =>
      _singleton ??= TrackerService._internal(channel, interceptors);

  TrackerService._internal(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors)
      : super(channel, interceptors: interceptors);
}
