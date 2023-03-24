import 'package:cosmos_notifier/config.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:responsive_framework/responsive_framework.dart';

class CollapsibleHeader extends StatefulWidget {
  const CollapsibleHeader({Key? key}) : super(key: key);

  @override
  _CollapsibleHeaderState createState() => _CollapsibleHeaderState();
}

class MenuButtonData {
  final String title;
  final IconData icon;
  final RouteData routeData;

  MenuButtonData(this.title, this.icon, this.routeData);
}

class _CollapsibleHeaderState extends State<CollapsibleHeader> {
  bool isCollapsed = false;

  Iterable<MenuButtonData> getRoutes() {
    return [
      MenuButtonData("Home", Icons.home, rRoot),
      MenuButtonData("Subscriptions", Icons.notifications, rSubscriptions),
      if (jwtManager.isAdmin) MenuButtonData("Tracking", Icons.my_location, rTracking),
      if (jwtManager.isAdmin) MenuButtonData("Admin", Icons.settings, rAdmin),
    ];
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
    return PopupMenuButton(
      itemBuilder: (_) => menuItems,
      child: Row(
        children: const [
          Icon(Icons.menu),
          SizedBox(width: 5),
          Text("Menu"),
        ],
      ),
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
          children: menuItems,
        ),
        const Divider(),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints constraints) {
        final shouldCollapse = ResponsiveWrapper.of(context).isSmallerThan(TABLET);
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
