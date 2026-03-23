import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:jwt_decode/jwt_decode.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/shared/providers/user_provider.dart' as users_provider;
import 'core_provider.dart';

class Lifecycle extends ConsumerStatefulWidget {
  const Lifecycle({super.key, required this.child});

  final Widget child;

  @override
  ConsumerState<Lifecycle> createState() => _LifecycleState();
}

class _LifecycleState extends ConsumerState<Lifecycle> with WidgetsBindingObserver {
  bool isLoggedIn = false;

  @override
  void initState() {
    WidgetsBinding.instance.addObserver(this);
    super.initState();
  }

  @override
  void dispose() {
    WidgetsBinding.instance.removeObserver(this);
    super.dispose();
  }

  @override
  Future<void> didChangeAppLifecycleState(AppLifecycleState state) async {
    if (state == AppLifecycleState.resumed) {
      dynamic token = ref.read(hiveProvider).token;

      bool isExpired = token != null ? Jwt.isExpired(token as String) : false;

      if (isExpired) {
        await ref.read(httpInterceptorProvider).verifyToken();
        /*Future.delayed(const Duration(seconds: 1), () {
          final snackBar = SnackBar(
            behavior: SnackBarBehavior.floating,
            margin: EdgeInsets.only(bottom: MediaQuery.of(scaffoldKey.currentState!.context).size.height - 150),
            content: Text(AppLocalizations.of(scaffoldKey.currentState!.context)!.sessionExpired),
            duration: const Duration(seconds: 5),
          );
          scaffoldKey.currentState?.hideCurrentSnackBar();
          scaffoldKey.currentState?.showSnackBar(snackBar);

          ref.read(lifecycleRouteProvider.notifier).update((state) => RouteNames.homeRouteLocation);
        });*/
      }

      await ref.read(users_provider.userNotifier).getUser().then((user) {
        ref.read(userProvider.notifier).update((state) => state = state?.copyWith(
            profilePicture: user.profilePicture));
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return widget.child;
  }
}
