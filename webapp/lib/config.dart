import 'package:cosmos_notifier/f_admin/widget/services/admin_service.dart';
import 'package:grpc/grpc_web.dart';
import 'package:cosmos_notifier/env.dart';
import 'package:cosmos_notifier/f_home/services/auth_interceptor.dart';
import 'package:cosmos_notifier/f_home/services/auth_service.dart';
import 'package:cosmos_notifier/f_home/services/jwt_manager.dart';
import 'package:cosmos_notifier/f_subscription/services/subscription_service.dart';

const refreshBeforeExpDuration = Duration(seconds: 10 * 60);

final channel = GrpcWebClientChannel.xhr(Uri.parse(uri));

final jwtManager = JwtManager();
final authInterceptor = AuthInterceptor(jwtManager);
final authService = AuthService(channel, [authInterceptor], jwtManager, refreshBeforeExpDuration);
final subsService = SubscriptionService(channel, [authInterceptor]);
final adminService = AdminService(channel, [authInterceptor]);

class RouteData {
  final String name;
  final String path;

  const RouteData(this.name, this.path);
}

const rRoot = RouteData("root", "/");
const rLoading = RouteData("loading", "/loading");
const rLogin = RouteData("login", "/login");
const rSubscriptions = RouteData("subscriptions", "/subscriptions");
const rAdmin = RouteData("admin", "/admin");

const bool cDebugMode = true;

final tgBotUrl = Uri.parse('https://t.me/DaoDaoNotifierBot');
final discordBotUrl = Uri.parse('https://discord.com/api/oauth2/authorize?client_id=953923165808107540&redirect_uri=https%3A%2F%2Fcosmos-notifier.decrypto.online&response_type=code&scope=identify');

final tgContactUrl = Uri.parse('https://t.me/rapha_decrypto');
final discordContactUrl = Uri.parse('https://discord.com/users/228978159440232453');
final githubUrl = Uri.parse('https://github.com/shifty11/dao-dao-notifier');
final twitterUrl = Uri.parse('https://twitter.com/Rapha90');