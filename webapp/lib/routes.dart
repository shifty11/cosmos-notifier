import 'package:fixnum/fixnum.dart';
import 'package:webapp/config.dart';
import 'package:webapp/f_home/services/auth_provider.dart';
import 'package:webapp/f_home/services/chat_id_provider.dart';
import 'package:webapp/f_home/services/state/auth_state.dart';
import 'package:webapp/f_home/widgets/loading_page.dart';
import 'package:webapp/f_login/widgets/login_page.dart';
import 'package:webapp/f_subscription/widgets/subscription_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';


// chatIdProvider is needed here to read the queryParams after the first page load. Later they will be striped away.
final routerProvider = Provider<MyRouter>((ref) => MyRouter(ref.watch(authStateValueProvider), ref.read(chatIdProvider)));

class MyRouter {
  final ValueNotifier<AuthState> authStateListener;

  MyRouter(this.authStateListener, Int64? _);

  late final router = GoRouter(
    refreshListenable: authStateListener,
    // debugLogDiagnostics: cDebugMode,
    urlPathStrategy: UrlPathStrategy.hash,
    routes: [
      GoRoute(
        name: rRoot.name,
        path: rRoot.path,
        pageBuilder: (context, state) => MaterialPage<void>(
          key: state.pageKey,
          child: const LoadingPage(),
        ),
      ),
      GoRoute(
        name: rUnauthenticated.name,
        path: rUnauthenticated.path,
        pageBuilder: (context, state) => MaterialPage<void>(
          key: state.pageKey,
          child: const LoginPage(),
        ),
      ),
      GoRoute(
        name: rSubscriptions.name,
        path: rSubscriptions.path,
        pageBuilder: (context, state) => MaterialPage<void>(
          key: state.pageKey,
          child: const SubscriptionPage(),
        ),
      ),
    ],
    errorPageBuilder: (context, state) => MaterialPage<void>(
      key: state.pageKey,
      // child: ErrorPage(error: state.error),
      child: Scaffold(body: Center(child: Text("There was an error ${state.error}")),),
    ),
    redirect: (state) {
      return authStateListener.value.when(
        loading: () => state.subloc != rRoot.path ? state.namedLocation(rRoot.name, queryParams: {"from": state.subloc}) : null,
        authorized: () {
          // if `from` is set redirect there; if current page is / or /login redirect to /subscriptions
          final from = state.queryParams["from"] ?? "";
          if (from.isNotEmpty && state.subloc != from && from != rUnauthenticated.path) {
            try {
              return state.namedLocation(from.replaceAll("/", ""));
            } catch (e) {
              if (cDebugMode) {
                print("error in router redirect: $e");
              }
            }
          }
          if (state.subloc == rRoot.path || state.subloc == rUnauthenticated.path) {
            return state.namedLocation(rSubscriptions.name);
          }
          return null;
        },
        expired: () => state.subloc == rUnauthenticated.path ? null : state.namedLocation(rUnauthenticated.name),
        userNotFound: () => state.subloc == rUnauthenticated.path ? null : state.namedLocation(rUnauthenticated.name),
        error: () => state.subloc == rUnauthenticated.path ? null : state.namedLocation(rUnauthenticated.name),
      );
    },
  );
}
