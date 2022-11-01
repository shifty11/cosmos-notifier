import 'dart:math';

import 'package:flutter/material.dart';
import 'package:dao_dao_notifier/style.dart';

class LogoWidget extends StatelessWidget {
  final double sideBarWith = 300;

  const LogoWidget({Key? key}) : super(key: key);

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

  @override
  Widget build(BuildContext context) {
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
              Text("Cosmos", style: Theme.of(context).textTheme.headline3?.copyWith(fontSize: fontSize, fontFamily: "Alienated")),
              Text("Notifier", style: Theme.of(context).textTheme.headline3!.copyWith(fontSize: fontSize, fontFamily: "Alien Robot")),
            ],
          )),
        ),
      ),
    ]);
  }
}
