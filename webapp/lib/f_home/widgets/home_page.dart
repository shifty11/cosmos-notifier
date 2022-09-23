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

  bool _isSmall(BuildContext context) {
    return MediaQuery.of(context).size.height <= 750;
  }

  Widget logo(BuildContext context) {
    final isSmall = _isSmall(context);
    final fontSize = isSmall ? 13.0 : Theme.of(context).textTheme.headline3!.fontSize;
    return Stack(children: [
      CircleAvatar(
        radius: isSmall ? 90 : 180,
        backgroundColor: Styles.isDarkTheme(context) ? Colors.white : Colors.black,
        child: ClipOval(
          child: Image.asset(
            "images/dove_square.png",
            width: isSmall ? 170 : 340,
            height: isSmall ? 170 : 340,
          ),
        ),
      ),
      Positioned(
        top: isSmall ? 30 : 70,
        child: SizedBox(
          width: isSmall ? 180 : 360,
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

  _rowOrColumn(BuildContext context) {
    return MediaQuery.of(context).size.width <= 400 ? Row : Column;
  }

  _createButtons(double buttonWith, double spaceBetween) {
    return [
      ElevatedButton.icon(
        onPressed: () async => await launchUrl(tgBotUrl),
        icon: const Icon(Icons.telegram),
        label: const Text("Telegram"),
        style: ElevatedButton.styleFrom(
          minimumSize: Size(buttonWith, 50),
          primary: const Color(0xFF54A9E9),
          onPrimary: Colors.white,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(32.0),
          ),
        ),
      ),
      SizedBox(width: spaceBetween, height: spaceBetween),
      ElevatedButton.icon(
        onPressed: () async => await launchUrl(discordBotUrl),
        icon: const Icon(Icons.discord),
        label: const Text("Discord"),
        style: ElevatedButton.styleFrom(
          minimumSize: Size(buttonWith, 50),
          primary: const Color(0xFF6C89E0),
          onPrimary: Colors.white,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(32.0),
          ),
        ),
      ),
    ];
  }

  Widget botButtons(BuildContext context) {
    const buttonWith = 180.0;
    const spaceBetween = 20.0;
    return MediaQuery.of(context).size.width <= 2 * buttonWith + spaceBetween + 2 * Styles.sidePadding
        ? Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: _createButtons(buttonWith, spaceBetween),
          )
        : Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: _createButtons(buttonWith, spaceBetween),
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
    final isSmall = _isSmall(context);
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: Styles.sidePadding),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(height: isSmall ? 20 : 50),
            logo(context),
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
            const FooterWidget(),
          ],
        ),
      ),
    );
  }
}
