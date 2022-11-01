import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:dao_dao_notifier/config.dart';
import 'package:dao_dao_notifier/f_home/services/auth_provider.dart';
import 'package:dao_dao_notifier/f_home/services/auth_service.dart';
import 'package:dao_dao_notifier/f_home/widgets/subwidgets/footer_widget.dart';
import 'package:dao_dao_notifier/f_home/widgets/subwidgets/logo_widget.dart';
import 'package:dao_dao_notifier/style.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key, errorCode}) : super(key: key);

  Widget botButtons(BuildContext context) {
    const buttonWith = 170.0;
    const spaceBetween = 20.0;
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        ElevatedButton.icon(
          onPressed: () async => await launchUrl(tgBotUrl),
          icon: const Icon(Icons.telegram),
          label: const Text("Telegram"),
          style: ElevatedButton.styleFrom(
            minimumSize: const Size(buttonWith, 50),
            primary: Styles.telegramColor,
            onPrimary: Colors.white,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(32.0),
            ),
          ),
        ),
        const SizedBox(width: spaceBetween, height: spaceBetween),
        ElevatedButton.icon(
          onPressed: () async => await launchUrl(discordBotUrl),
          icon: const Icon(Icons.discord),
          label: const Text("Discord"),
          style: ElevatedButton.styleFrom(
            minimumSize: const Size(buttonWith, 50),
            primary: Styles.discordColor,
            onPrimary: Colors.white,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(32.0),
            ),
          ),
        ),
      ],
    );
  }

  Widget subscriptionButton() {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final auth = ref.watch(authProvider);
      if (auth.isAuthenticated || auth.canRefreshAccessToken) {
        return Column(
          children: [
            OutlinedButton.icon(
              onPressed: () => GoRouter.of(context).go(rSubscriptions.path),
              icon: const Icon(Icons.notifications),
              label: const Text("Manage Subscriptions"),
              style: OutlinedButton.styleFrom(
                minimumSize: const Size(380, 50),
              ),
            ),
          ],
        );
      }
      return Container();
    });
  }

  Widget errorMsg() {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final state = ref.watch(authStateValueProvider);
      return state.value.whenOrNull(
            error: (err) {
              var text = "Unknown error.\nPlease login again.";
              if (err is AuthExpiredError) {
                text = "Your session has expired.\nPlease login again.";
              }
              if (err is AuthUserNotFoundError) {
                text = "User was not found.\nUse the Telegram or Discord bot to register.";
              }
              return Container(
                  padding: const EdgeInsets.all(20),
                  decoration: BoxDecoration(
                    color: Styles.dangerBgColor,
                    borderRadius: BorderRadius.circular(10),
                    boxShadow: [
                      BoxShadow(
                        color: Colors.black.withOpacity(0.2),
                        spreadRadius: 2,
                        blurRadius: 5,
                        offset: const Offset(0, 3), // changes position of shadow
                      ),
                    ],
                  ),
                  child: Text(
                      text,
                      textAlign: TextAlign.center,
                      style: Theme.of(context).textTheme.headline5?.copyWith(color: Styles.dangerTextColor)));
            },
          ) ??
          Container();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: Styles.sidePadding),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Spacer(flex: 2),
            MediaQuery.of(context).size.height >= 400 ? const LogoWidget() : Container(),
            const Spacer(flex: 2),
            Flexible(
              flex: 0,
              child: Text("Get notified about governance proposals of Cosmos chains and DAO's",
                  textAlign: TextAlign.center, style: Theme.of(context).textTheme.headline5),
            ),
            const SizedBox(height: 10),
            Flexible(
                flex: 0,
                child: Text("Now available on Telegram and Discord",
                    textAlign: TextAlign.center, style: Theme.of(context).textTheme.headline6)),
            const Spacer(flex: 1),
            botButtons(context),
            const Spacer(flex: 1),
            subscriptionButton(),
            errorMsg(),
            const Spacer(flex: 4),
            const Flexible(flex: 0, child: FooterWidget()),
          ],
        ),
      ),
    );
  }
}
