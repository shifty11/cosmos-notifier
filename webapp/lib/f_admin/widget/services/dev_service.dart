import 'package:cosmos_notifier/api/protobuf/dart/dev_service.pbgrpc.dart';
import 'package:grpc/grpc.dart';
import 'package:grpc/grpc_connection_interface.dart';

class DevService extends DevServiceClient {
  static DevService? _singleton;

  factory DevService(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors) =>
      _singleton ??= DevService._internal(channel, interceptors);

  DevService._internal(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors)
      : super(channel, interceptors: interceptors);
}
