import 'package:flutter/material.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({Key? key}) : super(key: key);

  bool get hasError {
    return Uri.base.queryParameters['error'] == "true";
  }

  @override
  Widget build(BuildContext context) {
    final String text = hasError ? "There was an error" : "Your session expired";
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Text(text, style: Theme.of(context).textTheme.headline3),
            Text("Please login again", style: Theme.of(context).textTheme.bodyLarge),
          ],
        ),
      ),
    );
  }
}
