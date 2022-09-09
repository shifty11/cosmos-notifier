import 'package:grpc/grpc_web.dart';
import 'package:webapp/env.dart';
import 'package:webapp/f_home/services/auth_interceptor.dart';
import 'package:webapp/f_home/services/auth_service.dart';
import 'package:webapp/f_home/services/jwt_manager.dart';
import 'package:webapp/f_subscription/services/subscription_service.dart';

const refreshBeforeExpDuration = Duration(seconds: 10 * 60);

final channel = GrpcWebClientChannel.xhr(Uri.parse(uri));

final jwtManager = JwtManager();
final authInterceptor = AuthInterceptor(jwtManager);
final authService = AuthService(channel, [authInterceptor], jwtManager, refreshBeforeExpDuration);
final subsService = SubscriptionService(channel, [authInterceptor]);

class RouteData {
  final String name;
  final String path;

  const RouteData(this.name, this.path);
}

const rRoot = RouteData("root", "/");
const rUnauthenticated = RouteData("unauthenticated", "/login");
const rSubscriptions = RouteData("subscriptions", "/subscriptions");

const bool cDebugMode = true;
