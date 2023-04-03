import 'package:cosmos_notifier/config.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:responsive_framework/responsive_framework.dart';

class HeaderWidget extends StatefulWidget {
  const HeaderWidget({Key? key}) : super(key: key);

  @override
  _HeaderWidgetState createState() => _HeaderWidgetState();
}

class MenuButtonData {
  final String title;
  final IconData icon;
  final MyRouteData routeData;

  MenuButtonData(this.title, this.icon, this.routeData);
}

class _HeaderWidgetState extends State<HeaderWidget> {
  bool isCollapsed = false;

  Iterable<MenuButtonData> getRoutes() {
    return [
      MenuButtonData("Subscriptions", Icons.notifications, rSubscriptions),
      MenuButtonData("Reminders", Icons.alarm, rTracking),
      if (jwtManager.isAdmin) MenuButtonData("Admin", Icons.settings, rAdmin),
    ];
  }

  Widget homeButton() {
    return IconButton(
      icon: Image.asset("images/dove_round.png", width: 36, height: 36),
      onPressed: () => context.pushNamed(rHome.name),
    );
  }

  Widget getPopupMenu() {
    var location = GoRouter.of(context).location;
    final menuItems = getRoutes().map((data) {
      final color = location == data.routeData.path ? Theme.of(context).primaryColor : Theme.of(context).disabledColor;
      return PopupMenuItem(
        value: data.routeData,
        onTap: () {
          context.pushNamed(data.routeData.name);
        },
        child: ListTile(
          leading: Icon(data.icon, color: color),
          title: Text(data.title, style: TextStyle(color: color)),
        ),
      );
    }).toList();
    return Column(
      children: [
        Row(
          children: [
            homeButton(),
            const SizedBox(width: 8),
            PopupMenuButton(
              itemBuilder: (_) => menuItems,
              child: Row(
                children: const [
                  Icon(Icons.menu),
                ],
              ),
            ),
          ],
        ),
        const Divider(),
      ],
    );
  }

  Widget getExpandedMenu() {
    var location = GoRouter.of(context).location;
    final menuItems = getRoutes().map((data) {
      final color = location == data.routeData.path ? Theme.of(context).primaryColor : Theme.of(context).disabledColor;
      return TextButton.icon(
        onPressed: () => context.pushNamed(data.routeData.name),
        icon: Icon(data.icon, color: color),
        label: Text(data.title, style: TextStyle(color: color)),
      );
    }).toList();
    return Column(
      children: [
        Row(
          children: [
            ...menuItems,
            const Spacer(),
            homeButton(),
          ],
        ),
        const Divider(),
      ],
    );
  }

  bool get shouldCollapse {
    if (ResponsiveWrapper.of(context).isSmallerThan(TABLET)) {
      return getRoutes().length > 2;
    }
    return false;
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints constraints) {
        return shouldCollapse
            ? GestureDetector(
                onTap: () {
                  setState(() {
                    isCollapsed = !isCollapsed;
                  });
                },
                child: getPopupMenu(),
              )
            : getExpandedMenu();
      },
    );
  }
}
