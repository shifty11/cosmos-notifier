import 'package:dao_dao_notifier/f_home/services/theme_provider.dart';
import 'package:dao_dao_notifier/routes.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:responsive_framework/responsive_framework.dart';

void main() {
  runApp(ProviderScope(child: Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
    final router = ref.read(routerProvider).router;
    final state = ref.watch(themeProvider);
    return MaterialApp.router(
      title: 'DAO DAO Notifier',
      theme: state.map(initial: (theme) => theme.lightStyle, custom: (theme) => theme.style),
      darkTheme: state.map(initial: (theme) => theme.darkStyle, custom: (theme) => theme.style),
      themeMode: ThemeMode.system,
      builder: (context, widget) => ResponsiveWrapper.builder(
        BouncingScrollWrapper.builder(context, widget!),
        maxWidth: 1200,
        minWidth: 450,
        defaultScale: true,
        breakpoints: [
          const ResponsiveBreakpoint.resize(450, name: MOBILE),
          const ResponsiveBreakpoint.autoScale(800, name: TABLET),
          const ResponsiveBreakpoint.autoScale(1000, name: TABLET),
          const ResponsiveBreakpoint.resize(1200, name: DESKTOP),
        ],
        background: Container(color: MediaQuery.of(context).platformBrightness == Brightness.dark ? Colors.black : Colors.white),
      ),
      routeInformationProvider: router.routeInformationProvider,
      routeInformationParser: router.routeInformationParser,
      routerDelegate: router.routerDelegate,
    );
  })));
}
