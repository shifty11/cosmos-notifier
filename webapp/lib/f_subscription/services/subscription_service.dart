import 'package:cosmos_notifier/api/protobuf/dart/subscription_service.pbgrpc.dart';
import 'package:grpc/grpc.dart';
import 'package:grpc/grpc_connection_interface.dart';

class SubscriptionService extends SubscriptionServiceClient {
  static SubscriptionService? _singleton;

  factory SubscriptionService(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors) =>
      _singleton ??= SubscriptionService._internal(channel, interceptors);

  SubscriptionService._internal(ClientChannelBase channel, Iterable<ClientInterceptor> interceptors)
      : super(channel, interceptors: interceptors);
}
