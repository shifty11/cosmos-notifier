import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/jwt_manager.dart';

class NavigationData {
  final int index;
  final RouteData routerData;
  final Icon icon;
  final String label;

  const NavigationData(this.index, this.routerData, this.icon, this.label);
}

List<NavigationData> buildMenu(JwtManager jwtManager) {
  List<NavigationData> menu = [
    const NavigationData(0, rSubscriptions, Icon(Icons.bookmarks), 'Subscriptions'),
  ];
  if (jwtManager.isAdmin) {}
  return menu;
}

class BottomNavigationBarWidget extends StatelessWidget {
  final double sideBarWith = 300;
  final JwtManager jwtManager;

  final List<NavigationData> menu;

  BottomNavigationBarWidget({Key? key, required this.jwtManager})
      : menu = buildMenu(jwtManager),
        super(key: key);

  int getIndex(BuildContext context) {
    final location = GoRouter.of(context).location; // format: /subscriptions?chat_id=123
    for (var d in menu) {
      if (location.startsWith(d.routerData.path)) {
        return d.index;
      }
    }
    return -1;
  }

  goTo(BuildContext context, int index) {
    final routerData = menu.firstWhere((d) => d.index == index).routerData;
    context.goNamed(routerData.name);
  }

  BottomNavigationBarItem navigationItem(Icon icon, String label) {
    return BottomNavigationBarItem(icon: icon, label: label);
  }

  @override
  Widget build(BuildContext context) {
    // if (!jwtManager.isAdmin) {
      return const SizedBox.shrink();
    // }
    return BottomNavigationBar(
        onTap: (index) => goTo(context, index),
        currentIndex: getIndex(context),
        items: menu
            .map((d) => BottomNavigationBarItem(
                  icon: d.icon,
                  label: d.label,
                  tooltip: '',
                ))
            .toList());
  }
}
