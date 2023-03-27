import 'package:cosmos_notifier/f_home/services/type/custom_theme_data.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/transition_builder_widget.dart';
import 'package:flutter/material.dart';

extension ColorBrightness on Color {
  Color darken([double amount = .1]) {
    assert(amount >= 0 && amount <= 1);

    final hsl = HSLColor.fromColor(this);
    final hslDark = hsl.withLightness((hsl.lightness - amount).clamp(0.0, 1.0));

    return hslDark.toColor();
  }

  Color lighten([double amount = .1]) {
    assert(amount >= 0 && amount <= 1);

    final hsl = HSLColor.fromColor(this);
    final hslLight = hsl.withLightness((hsl.lightness + amount).clamp(0.0, 1.0));

    return hslLight.toColor();
  }

  Color intensify(BuildContext context, [double amount = .1]) {
    if (Styles.isDarkTheme(context)) {
      return lighten(amount);
    }
    return darken(amount);
  }

  Color intensifyBg(BuildContext context, [double amount = .1]) {
    if (Styles.isDarkTheme(context)) {
      return darken(amount);
    }
    return lighten(amount);
  }
}

class Styles {
  static const topPadding = 40.0;
  static const sidePadding = 40.0;
  static const selectCardBorderWidth = 1.5;

  static const dangerBgColor = Color(0xFFFE4A49);
  static const dangerTextColor = Colors.white;
  static const telegramColor = Color(0xFF54A9E9);
  static const discordColor = Color(0xFF6C89E0);
  static const cannyColor = Color(0xFF525DF9);

  static ThemeData customTheme(CustomThemeData themeParams, bool isDarkTheme) {
    return _defaultTheme(
      isDarkTheme,
      themeParams.bgColor,
      isDarkTheme ? themeParams.bgColor.lighten() : themeParams.bgColor.darken(),
      themeParams.textColor,
      themeParams.hintColor,
      themeParams.buttonColor,
    );
  }

  static ThemeData defaultTheme(bool isDarkTheme) {
    final bgColor = isDarkTheme ? const Color(0xff1d2733) : Colors.white.darken(0.025);
    final bgColorLight = isDarkTheme ? bgColor.lighten() : Colors.white.darken();
    final textColor = isDarkTheme ? Colors.white : Colors.black;
    final textColorHint = isDarkTheme ? const Color(0xff7d8b99) : Colors.black;
    const primaryColor = Color(0xff50a8eb);
    return _defaultTheme(isDarkTheme, bgColor, bgColorLight, textColor, textColorHint, primaryColor);
  }

  static isDarkTheme(BuildContext context) {
    return Theme.of(context).brightness == Brightness.dark;
  }

  static ThemeData _defaultTheme(
      bool isDarkTheme, Color bgColor, Color bgColorLight, Color textColor, Color textColorHint, Color primaryColor) {
    const borderColor = Colors.grey;
    var textTheme = isDarkTheme ? Typography().white : Typography().black;
    return ThemeData(
      brightness: isDarkTheme ? Brightness.dark : Brightness.light,
      fontFamily: "Montserrat",
      useMaterial3: true,
      primaryColor: primaryColor,
      primarySwatch: createMaterialColor(primaryColor),
      iconTheme: const IconThemeData(color: borderColor),
      scaffoldBackgroundColor: bgColor,
      canvasColor: bgColor,
      bottomNavigationBarTheme: BottomNavigationBarThemeData(backgroundColor: bgColorLight, unselectedItemColor: borderColor),
      unselectedWidgetColor: borderColor,
      inputDecorationTheme: InputDecorationTheme(
          labelStyle: TextStyle(color: textColor),
          hintStyle: const TextStyle(color: borderColor),
          iconColor: borderColor,
          enabledBorder: const OutlineInputBorder(borderSide: BorderSide(color: borderColor))),
      dialogBackgroundColor: bgColor,
      textTheme: textTheme,
      toggleButtonsTheme: ToggleButtonsThemeData(
        selectedColor: primaryColor,
      ),
      textButtonTheme: TextButtonThemeData(
        style: TextButton.styleFrom(
          foregroundColor: textColor,
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(5)),
        ),
      ),
      outlinedButtonTheme: OutlinedButtonThemeData(
        style: OutlinedButton.styleFrom(
            foregroundColor: primaryColor,
            side: BorderSide(color: primaryColor),
            shape: const RoundedRectangleBorder(
              borderRadius: BorderRadius.all(Radius.circular(5)),
            )),
      ),
      elevatedButtonTheme: ElevatedButtonThemeData(
        style: ElevatedButton.styleFrom(
            backgroundColor: primaryColor, // Button color
            foregroundColor: textColor, // Text color
            shape: const RoundedRectangleBorder(
              borderRadius: BorderRadius.all(Radius.circular(5)),
            )),
      ),
      popupMenuTheme: PopupMenuThemeData(
        color: bgColor,
      ),
      dialogTheme: DialogTheme(
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(5)),
      ),
      pageTransitionsTheme: const PageTransitionsTheme(
        builders: {
          TargetPlatform.linux: NoAnimationTransitionsBuilder(),
          TargetPlatform.android: NoAnimationTransitionsBuilder(),
          TargetPlatform.fuchsia: NoAnimationTransitionsBuilder(),
          TargetPlatform.iOS: NoAnimationTransitionsBuilder(),
          TargetPlatform.macOS: NoAnimationTransitionsBuilder(),
          TargetPlatform.windows: NoAnimationTransitionsBuilder(),
        },
      ),
    );
  }
}

MaterialColor createMaterialColor(Color color) {
  List strengths = <double>[.05];
  Map<int, Color> swatch = {};
  final int r = color.red, g = color.green, b = color.blue;

  for (int i = 1; i < 10; i++) {
    strengths.add(0.1 * i);
  }
  for (var strength in strengths) {
    final double ds = 0.5 - strength;
    swatch[(strength * 1000).round()] = Color.fromRGBO(
      r + ((ds < 0 ? r : (255 - r)) * ds).round(),
      g + ((ds < 0 ? g : (255 - g)) * ds).round(),
      b + ((ds < 0 ? b : (255 - b)) * ds).round(),
      1,
    );
  }
  return MaterialColor(color.value, swatch);
}
