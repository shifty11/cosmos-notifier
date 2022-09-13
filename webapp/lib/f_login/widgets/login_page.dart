import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:webapp/f_home/services/auth_provider.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({Key? key}) : super(key: key);

  bool get hasError {
    return Uri.base.queryParameters['error'] == "true";
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        // child: Container(),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
              final state = ref.read(authStateValueProvider);
              return Text(
                  state.value.whenOrNull(
                        expired: () => "Your session expired",
                        userNotFound: () => "User was not found. Use the Telegram or Discord bot to register",
                      ) ??
                      "There was an unknown error",
                  style: Theme.of(context).textTheme.headline3);
            }),
            Text("Please login again", style: Theme.of(context).textTheme.bodyLarge),
          ],
        ),
      ),
    );
  }
}
