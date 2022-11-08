import 'package:cosmos_notifier/api/protobuf/dart/admin_service.pbgrpc.dart';
import 'package:grpc/grpc.dart';
import 'package:grpc/grpc_connection_interface.dart';

class AdminService extends AdminServiceClient {
  static AdminService? _singleton;

  factory AdminService(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors) =>
      _singleton ??= AdminService._internal(channel, interceptors);

  AdminService._internal(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors)
      : super(channel, interceptors: interceptors);
}
