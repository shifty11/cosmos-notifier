import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/auth_service.dart';
import 'package:cosmos_notifier/f_home/services/canny_query_param_provider.dart';
import 'package:cosmos_notifier/f_home/services/state/auth_state.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/footer_widget.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/logo_widget.dart';
import 'package:cosmos_notifier/f_login/services/login_state_provider.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:url_launcher/url_launcher.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({Key? key, errorCode}) : super(key: key);

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
            backgroundColor: Styles.telegramColor,
            foregroundColor: Colors.white,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(32.0),
            ),
          ),
        ),
        const SizedBox(width: spaceBetween, height: spaceBetween),
        ElevatedButton.icon(
          onPressed: () async => await launchUrl(discordOAuth2Url),
          icon: const Icon(Icons.discord),
          label: const Text("Discord"),
          style: ElevatedButton.styleFrom(
            minimumSize: const Size(buttonWith, 50),
            backgroundColor: Styles.discordColor,
            foregroundColor: Colors.white,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(32.0),
            ),
          ),
        ),
      ],
    );
  }

  _launchExternalURL(CannySSO cannySSO) async {
    var uri = Uri.parse(
        "https://canny.io/api/redirects/sso?companyID=${cannySSO.companyId}&ssoToken=${cannySSO.ssoToken}&redirect=${cannySSO.redirectUrl}");
    if (await canLaunchUrl(uri)) {
      if (cDebugMode) {
        print("Launching $uri");
      }
      await launchUrl(uri, webOnlyWindowName: '_self');
    } else {
      throw 'Could not launch $uri';
    }
  }

  Widget cannyButton(BuildContext context, WidgetRef ref) {
    final cannySSO = ref.watch(cannySSOProvider);
    return cannySSO.isCannySSO
        ? ElevatedButton.icon(
            onPressed: () async => await _launchExternalURL(cannySSO),
            icon: const Icon(Icons.reviews),
            label: const Text("Go back to Canny"),
            style: ElevatedButton.styleFrom(
              minimumSize: const Size(170, 50),
              foregroundColor: Colors.white,
              backgroundColor: Styles.cannyColor,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(32.0),
              ),
            ),
          )
        : Container();
  }

  Widget errorMsg() {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final state = ref.watch(loginStateProvider);
      return state.whenOrNull(
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
                  child: Text(text,
                      textAlign: TextAlign.center, style: Theme.of(context).textTheme.headlineSmall?.copyWith(color: Styles.dangerTextColor)));
            },
          ) ??
          Container();
    });
  }

  static const defaultHeadlineText = "Login to Cosmos Notifier";

  String getHeadlineText(WidgetRef ref) {
    return ref.watch(loginStateProvider).when(loading: () {
      return defaultHeadlineText;
    }, authenticated: (cannySSO) {
      if (cannySSO.ssoToken.isNotEmpty) {
        return "Already logged in. You can go back to Canny.";
      } else {
        return defaultHeadlineText;
      }
    }, unauthenticated: (cannySSO) {
      if (cannySSO.isCannySSO) {
        return "Login to send feedback on Canny";
      }
      return defaultHeadlineText;
    }, error: (err) {
      return defaultHeadlineText;
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
              child: Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
                return Text(getHeadlineText(ref), textAlign: TextAlign.center, style: Theme.of(context).textTheme.headlineSmall);
              }),
            ),
            // isCannySSO(context, ref) ? const Text("Canny SSO") : const Text("Not Canny SSO"),
            const Spacer(flex: 1),
            Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
              return ref.watch(loginStateProvider).when(loading: () {
                return const CircularProgressIndicator();
              }, authenticated: (cannySSO) {
                if (cannySSO.isCannySSO) {
                  return cannyButton(context, ref);
                } else {
                  return botButtons(context);
                }
              }, unauthenticated: (cannySSO) {
                return botButtons(context);
              }, error: (err) {
                return Container();
              });
            }),
            const Spacer(flex: 1),
            errorMsg(),
            const Spacer(flex: 4),
            const Flexible(flex: 0, child: FooterWidget()),
          ],
        ),
      ),
    );
  }
}
