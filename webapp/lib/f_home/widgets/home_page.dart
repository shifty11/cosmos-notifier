import 'dart:math';

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

  // returns a value between 120 und 360 depending on height and width of screen
  double _getSize(BuildContext context) {
    const maxSize = 360;
    return max(
        [
          MediaQuery.of(context).size.width - (450 - maxSize),
          MediaQuery.of(context).size.height - (700 - maxSize),
          360.0,
        ].reduce(min),
        120);
  }

  Widget logo(BuildContext context) {
    final size = _getSize(context);
    final double fontSize = size / 360 * (Theme.of(context).textTheme.headline3!.fontSize ?? 1);
    return Stack(children: [
      CircleAvatar(
        radius: size / 2,
        backgroundColor: Styles.isDarkTheme(context) ? Colors.white : Colors.black,
        child: ClipOval(
          child: Image.asset(
            "images/dove_square.png",
            width: size - (20 * (size / 360)),
            height: size - (20 * (size / 360)),
          ),
        ),
      ),
      Positioned(
        top: size / 36 * 7,
        child: SizedBox(
          width: size,
          child: Center(
              child: Column(
            children: [
              Text("DAO DAO", style: Theme.of(context).textTheme.headline3?.copyWith(fontSize: fontSize, fontFamily: "Alienated")),
              Text("Notifier", style: Theme.of(context).textTheme.headline3!.copyWith(fontSize: fontSize, fontFamily: "Alien Robot")),
            ],
          )),
        ),
      ),
    ]);
  }

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
            primary: const Color(0xFF54A9E9),
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
            primary: const Color(0xFF6C89E0),
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

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: Styles.sidePadding),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Spacer(flex: 2),
            MediaQuery.of(context).size.height >= 400 ? logo(context) : Container(),
            const Spacer(flex: 2),
            Flexible(
              flex: 0,
              child: Text("Get notified about DAO DAO governance proposals",
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
            const Spacer(flex: 4),
            const Flexible(flex: 0, child: FooterWidget()),
          ],
        ),
      ),
    );
  }
}
