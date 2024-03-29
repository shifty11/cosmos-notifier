import 'package:cosmos_notifier/env.dart';
import 'package:cosmos_notifier/f_admin/widget/services/admin_service.dart';
import 'package:cosmos_notifier/f_admin/widget/services/dev_service.dart';
import 'package:cosmos_notifier/f_home/services/auth_interceptor.dart';
import 'package:cosmos_notifier/f_home/services/auth_service.dart';
import 'package:cosmos_notifier/f_home/services/jwt_manager.dart';
import 'package:cosmos_notifier/f_subscription/services/subscription_service.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_service.dart';
import 'package:flutter/foundation.dart';
import 'package:grpc/grpc_web.dart';

const refreshBeforeExpDuration = Duration(seconds: 10 * 60);

final channel = GrpcWebClientChannel.xhr(Uri.parse(uri));

final jwtManager = JwtManager();
final authInterceptor = AuthInterceptor(jwtManager);
final devService = DevService(channel, [authInterceptor]);
final authService = AuthService(channel, [authInterceptor], jwtManager, refreshBeforeExpDuration, devService);
final subsService = SubscriptionService(channel, [authInterceptor]);
final adminService = AdminService(channel, [authInterceptor]);
final trackerService = TrackerService(channel, [authInterceptor]);

class MyRouteData {
  final String name;
  final String path;

  const MyRouteData(this.name, this.path);
}

const rHome = MyRouteData("home", "/");
const rLoading = MyRouteData("loading", "/loading");
const rLogin = MyRouteData("login", "/login");
const rSubscriptions = MyRouteData("subscriptions", "/subscriptions");
const rAdmin = MyRouteData("admin", "/admin");
const rTracking = MyRouteData("reminder", "/reminder");

const bool cDebugMode = true;

final tgBotUrl = Uri.parse('https://t.me/cosmos_gov_bot');
final discordBotUrl = Uri.parse(
    'https://discord.com/api/oauth2/authorize?client_id=953923165808107540&permissions=2048&redirect_uri=https%3A%2F%2Fcosmos-notifier.odincloud.xyz&response_type=code&scope=bot%20identify');
final discordOAuth2Url = kReleaseMode
    ? Uri.parse(
        'https://discord.com/api/oauth2/authorize?client_id=953923165808107540&redirect_uri=https%3A%2F%2Fcosmos-notifier.odincloud.xyz&response_type=code&scope=identify')
    : Uri.parse(
        'https://discord.com/api/oauth2/authorize?client_id=955835724714872942&redirect_uri=http%3A%2F%2Ftest.mydomain.com%3A40001&response_type=code&scope=identify');

final tgContactUrl = Uri.parse('https://t.me/rapha_decrypto');
final discordContactUrl = Uri.parse('https://discord.com/users/228978159440232453');
final githubUrl = Uri.parse('https://github.com/shifty11/cosmos-notifier');
final twitterUrl = Uri.parse('https://twitter.com/Rapha90');
