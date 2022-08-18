import 'package:flutter/material.dart';

class NoAnimationTransitionsBuilder extends PageTransitionsBuilder {
  const NoAnimationTransitionsBuilder();

  @override
  Widget buildTransitions<T>(
    PageRoute<T> route,
    BuildContext context,
    Animation<double> animation,
    Animation<double> secondaryAnimation,
    Widget child,
  ) {
    return _NoAnimationTransitionsBuilder(routeAnimation: animation, child: child);
  }
}

class _NoAnimationTransitionsBuilder extends StatelessWidget {
  const _NoAnimationTransitionsBuilder({Key? key, required Animation<double> routeAnimation, required this.child}) : super(key: key);

  final Widget child;

  @override
  Widget build(BuildContext context) {
    return child;
  }
}
