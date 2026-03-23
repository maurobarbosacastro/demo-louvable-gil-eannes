import 'package:tagpeak/features/auth/repositories/auth_repo.dart';
import 'package:tagpeak/utils/logging.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:local_auth/local_auth.dart';

class LocalAuth {
  late LocalAuthentication authentication;
  final StateController<bool?> canCheckBiometrics;
  final StateController<bool?> isAuthenticating;
  final StateController<List<BiometricType>> availableBiometrics;
  final AuthRepo authRepo;

  LocalAuth({
    required this.canCheckBiometrics,
    required this.isAuthenticating,
    required this.availableBiometrics,
    required this.authRepo
  }) {
    authentication = LocalAuthentication();
  }

  Future<bool> get isSupported async => await authentication.isDeviceSupported();

  Future<void> checkBiometrics() async {
    late bool canCheck;
    try {
      canCheck = await authentication.canCheckBiometrics;
    } on PlatformException catch (e) {
      canCheck = false;
      Log.info(e.message ?? "");
    }
    canCheckBiometrics.update((state) => canCheck);
  }

  Future<void> getAvailableBiometrics() async {
    late List<BiometricType> available;
    try {
      available = await authentication.getAvailableBiometrics();
    } on PlatformException catch (e) {
      available = <BiometricType>[];
      Log.info(e.message ?? "");
    }
    availableBiometrics.update((state) => [...available]);
  }

  Future<bool> authenticateWithBiometrics() async {
    bool authenticated = false;
    try {
      isAuthenticating.update((state) => true);
      authenticated = await authentication.authenticate(
        localizedReason: 'Scan your fingerprint (or face or whatever) to authenticate',
        options: const AuthenticationOptions(
          stickyAuth: true,
          useErrorDialogs: true,
          biometricOnly: true,
        ),
      );
      isAuthenticating.update((state) => false);
      return authenticated;

    } on PlatformException catch (e) {
      Log.info(e.message ?? "");
      isAuthenticating.update((state) => false);
      return false;
    }
  }
}

final canCheckBiometricsProvider = StateProvider<bool?>((ref) => null);
final isAuthenticatingProvider = StateProvider<bool?>((ref) => null);
final availableBiometricsProvider = StateProvider<List<BiometricType>>((ref) => []);
final localAuthProvider = Provider(
  (ref) => LocalAuth(
    canCheckBiometrics: ref.read(canCheckBiometricsProvider.notifier),
    isAuthenticating: ref.read(isAuthenticatingProvider.notifier),
    availableBiometrics: ref.read(availableBiometricsProvider.notifier),
    authRepo: ref.read(authRepo)
  ),
);
