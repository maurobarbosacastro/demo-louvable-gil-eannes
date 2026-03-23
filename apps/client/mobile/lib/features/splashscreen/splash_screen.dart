import 'package:tagpeak/features/auth/repositories/auth_repo.dart';
import 'package:tagpeak/features/core/application/local_auth/biometric_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:jwt_decode/jwt_decode.dart';
import 'package:package_info_plus/package_info_plus.dart';
import 'package:tagpeak/utils/constants/colors.dart';

import '../../utils/app_info_provider.dart';
import '../../utils/logging.dart';
import '../core/application/core_provider.dart';
import '../core/routes/route_names.dart';
import '../settings/widgets/dropdown_widget.dart';

class SplashScreen extends ConsumerStatefulWidget {
  const SplashScreen({super.key});

  @override
  ConsumerState<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends ConsumerState<SplashScreen> {
  @override
  Widget build(BuildContext context) {
    return Container(
      color: AppColors.wildSand,
    );
  }

  @override
  void initState() {
    super.initState();
    checkToken();
    appInfo();
  }

  void appInfo() async {
    PackageInfo packageInfo = await PackageInfo.fromPlatform();
    ref.read(appInfoProvider.notifier).update(
          (state) => AppInfoModel.init(
            version: packageInfo.version,
            packageName: packageInfo.packageName,
            buildNumber: packageInfo.buildNumber,
            appName: packageInfo.appName,
          ),
        );
  }

  Future<void> checkToken() async {
    dynamic token = ref.read(hiveProvider).token;
    dynamic refreshToken = ref.read(hiveProvider).refreshToken;
    bool biometricsEnabled = ref.read(hiveProvider).biometrics ?? false;

    bool isExpired = token != null ? Jwt.isExpired(token as String) : false;
    bool refreshIsExpired =
        refreshToken != null ? Jwt.isExpired(refreshToken as String) : false;

    Log.info('Token expired : $isExpired');
    Log.info('Refresh expired : $refreshIsExpired');

    Future(() {
      if (token != null && refreshToken != null) {
        navigateToPage(isExpired, refreshIsExpired, biometricsEnabled);
      } else {
        context.goNamed(RouteNames.loginRouteName);
      }
      FlutterNativeSplash.remove();
    });
  }

  void navigateToPage(
      bool tokenExpired, bool refreshTokenExpired, bool biometrics) async {
    if (tokenExpired && !refreshTokenExpired) {
      unlockApp();
    } else if (!tokenExpired || !refreshTokenExpired) {
      if (biometrics) {
        ref.watch(enableBiometricsProvider.notifier).update((state) => true);
        unlockApp();
      } else {
        ref.watch(enableBiometricsProvider.notifier).update((state) => false);
        context.goNamed(RouteNames.loginRouteName);
      }
    }
  }

  void unlockApp() async {
    try {
      //Verify if have Biometrics
      await ref.read(localAuthProvider).checkBiometrics();
      final bool haveBiometrics =
          ref.watch(canCheckBiometricsProvider) ?? false;
      //Unlock with Biometrics
      if (haveBiometrics) {
        bool isAuthed =
            await ref.watch(localAuthProvider).authenticateWithBiometrics();
        if (isAuthed) {
          await ref.watch(authRepo).refreshToken().then(
            (value) {
              if (value) {
                if (mounted) context.goNamed(RouteNames.clientDashboardRouteName);
              } else {
                if (mounted) context.goNamed(RouteNames.loginRouteName);
              }
            },
          );
        }
      } else {
        //Unlock with app pin
        // if (mounted) context.goNamed(RouteNames.loginRouteName);
      }
    } catch (e) {
      print(e);
      context.go(RouteNames.loginRouteLocation);
      return;
    }
  }
}
