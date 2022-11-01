import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:cosmos_notifier/config.dart';

class SidebarWidget extends StatelessWidget {
  final double sideBarWith = 300;

  const SidebarWidget({Key? key}) : super(key: key);

  Widget button(BuildContext context, String name, RouteData routeData) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 8.0),
      child: OutlinedButton(
        style: OutlinedButton.styleFrom(
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(30.0)),
          minimumSize: const Size.fromHeight(40),
          backgroundColor: GoRouter.of(context).location == routeData.path ? Colors.grey[400] : null,
        ),
        onPressed: () {
          context.goNamed(routeData.name);
        },
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Text(name, style: Theme.of(context).textTheme.titleLarge),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: sideBarWith,
      color: Colors.grey[300],
      padding: const EdgeInsets.only(top: 200, right: 50, left: 50),
      child: Column(
        children: [
          button(context, 'Subscriptions', rSubscriptions),
        ],
      ),
    );
  }
}
