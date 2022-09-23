import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:webapp/config.dart';
import 'package:webapp/f_home/services/auth_provider.dart';
import 'package:webapp/f_home/widgets/footer_widget.dart';
import 'package:webapp/style.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  Widget subscriptionButton() {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final auth = ref.watch(authProvider);
      if (auth.isAuthenticated || auth.canRefreshAccessToken) {
        return Column(
          children: [
            const SizedBox(
              height: 40,
            ),
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

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const SizedBox(height: 50),
          Stack(children: [
            CircleAvatar(
              radius: 180,
              backgroundColor: Styles.isDarkTheme(context) ? Colors.white : Colors.black,
              child: ClipOval(
                child: Image.asset(
                  "images/dove_square.png",
                  width: 340,
                  height: 340,
                ),
              ),
            ),
            Positioned(
                top: 70,
                child: SizedBox(width: 360, child: Center(child: Text("DAO DAO Notifier", style: Theme.of(context).textTheme.headline3)))),
          ]),
          const Spacer(flex: 2),
          Flexible(
            flex: 0,
            child: Text("Get notified about DAO DAO governance proposals",
                textAlign: TextAlign.center, style: Theme.of(context).textTheme.headline5),
          ),
          const SizedBox(height: 10),
          Flexible(
              flex: 0,
              child:
                  Text("Now available on Telegram and Discord", textAlign: TextAlign.center, style: Theme.of(context).textTheme.headline6)),
          const SizedBox(height: 40),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              ElevatedButton.icon(
                onPressed: () async => await launchUrl(tgBotUrl),
                icon: const Icon(Icons.telegram),
                label: const Text("Telegram"),
                style: ElevatedButton.styleFrom(
                  minimumSize: const Size(180, 50),
                  primary: const Color(0xFF54A9E9),
                  onPrimary: Colors.white,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(32.0),
                  ),
                ),
              ),
              const SizedBox(width: 20),
              ElevatedButton.icon(
                onPressed: () async => await launchUrl(discordBotUrl),
                icon: const Icon(Icons.discord),
                label: const Text("Discord"),
                style: ElevatedButton.styleFrom(
                  minimumSize: const Size(180, 50),
                  primary: const Color(0xFF6C89E0),
                  onPrimary: Colors.white,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(32.0),
                  ),
                ),
              ),
            ],
          ),
          subscriptionButton(),
          const Spacer(flex: 4),
          const FooterWidget(),
        ],
      ),
    );
  }
}
