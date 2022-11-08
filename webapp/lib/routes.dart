import 'package:cosmos_notifier/f_admin/widget/admin_page.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/auth_provider.dart';
import 'package:cosmos_notifier/f_home/services/chat_id_provider.dart';
import 'package:cosmos_notifier/f_home/services/state/auth_state.dart';
import 'package:cosmos_notifier/f_home/widgets/loading_page.dart';
import 'package:cosmos_notifier/f_subscription/widgets/subscription_page.dart';

import 'f_home/widgets/home_page.dart';

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
          child: const HomePage(),
        ),
      ),
      GoRoute(
        name: rLoading.name,
        path: rLoading.path,
        pageBuilder: (context, state) => MaterialPage<void>(
          key: state.pageKey,
          child: const LoadingPage(),
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
      GoRoute(
        name: rAdmin.name,
        path: rAdmin.path,
        pageBuilder: (context, state) => MaterialPage<void>(
          key: state.pageKey,
          child: AdminPage(),
        ),
      ),
    ],
    errorPageBuilder: (context, state) => MaterialPage<void>(
      key: state.pageKey,
      // child: ErrorPage(error: state.error),
      child: Scaffold(
        body: Center(child: Text("There was an error ${state.error}")),
      ),
    ),
    redirect: (state) {
      return authStateListener.value.when(
        initial: () => null,
        loading: () => state.subloc != rLoading.path ? state.namedLocation(rLoading.name) : null,
        authenticated: (redirect) {
          // redirect to subscription if user becomes authenticated
          if (redirect && (state.subloc == rLoading.path || state.subloc == rLogin.path)) {
            return state.namedLocation(rSubscriptions.name);
          }
          // non admins can not access the admin page
          if (state.subloc == rAdmin.path && !jwtManager.isAdmin) {
            return state.namedLocation(rSubscriptions.name);
          }
          return null;
        },
        unauthenticated: () {
          if (state.subloc != rRoot.path && state.subloc != rLogin.path) {
            return state.namedLocation(rRoot.name);
          }
          return null;
        },
        error: (err) => state.subloc == rRoot.path ? null : state.namedLocation(rRoot.name),
      );
    },
  );
}
