import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/admin_widget.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:url_launcher/url_launcher.dart';

class FooterWidget extends StatelessWidget {
  const FooterWidget({Key? key}) : super(key: key);

  static const iconSize = 30.0;
  static const spaceBetween = 15.0;

  @override
  Widget build(BuildContext context) {
    if (MediaQuery.of(context).size.height < 400) {
      return Container();
    }
    return Column(
      children: [
        const SizedBox(height: 10),
        const Divider(),
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            IconButton(
              padding: const EdgeInsets.all(0),
              onPressed: () async => await launchUrl(tgContactUrl),
              icon: const Icon(Icons.telegram, size: iconSize),
            ),
            const SizedBox(width: spaceBetween),
            IconButton(
              padding: const EdgeInsets.all(0),
              onPressed: () async => await launchUrl(discordContactUrl),
              icon: const Icon(Icons.discord, size: iconSize),
            ),
            const SizedBox(width: spaceBetween),
            IconButton(
              padding: const EdgeInsets.all(0),
              onPressed: () async => await launchUrl(twitterUrl),
              icon: const FaIcon(FontAwesomeIcons.twitter, size: iconSize),
            ),
            const SizedBox(width: spaceBetween),
            IconButton(
              padding: const EdgeInsets.all(0),
              onPressed: () async => await launchUrl(githubUrl),
              icon: const FaIcon(FontAwesomeIcons.github, size: iconSize),
            ),
          ],
        ),
        const SizedBox(height: 10),
        kReleaseMode ? Container() : const AdminWidget(),
      ],
    );
  }
}
