import 'package:tagpeak/features/core/application/local_auth/biometric_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class BiometricAuth extends ConsumerWidget {
  const BiometricAuth({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return  ElevatedButton(
      onPressed: () async => await ref.watch(localAuthProvider).authenticateWithBiometrics(),
      child: const Text('Get available biometrics'),
    );
  }
}
