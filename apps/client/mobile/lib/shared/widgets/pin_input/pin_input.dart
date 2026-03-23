import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../features/auth/repositories/auth_repo.dart';
import '../../../features/core/application/local_auth/biometric_provider.dart';
import '../../../features/core/routes/route_names.dart';
import '../../../utils/constants/assets.dart';
import 'unlock_providers.dart';

//Based in: https://medium.com/@henryifebunandu/create-custom-keyboard-for-your-flutter-app-20926a0aaf19
//          https://github.com/maykhid/custom_keyboard
class NumericKeypad extends ConsumerStatefulWidget {
  final int pinMaxChars;
  final VoidCallback callback;

  final VoidCallback? confirm;

  const NumericKeypad(this.callback, {super.key, this.pinMaxChars = 4, this.confirm});

  @override
  ConsumerState<NumericKeypad> createState() => _NumericKeypadState();
}

class _NumericKeypadState extends ConsumerState<NumericKeypad> {
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 16),
      child: Column(
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              _InputButton(text: '1', onPressed: () => _input('1')),
              _InputButton(text: '2', onPressed: () => _input('2')),
              _InputButton(text: '3', onPressed: () => _input('3')),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              _InputButton(text: '4', onPressed: () => _input('4')),
              _InputButton(text: '5', onPressed: () => _input('5')),
              _InputButton(text: '6', onPressed: () => _input('6')),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              _InputButton(text: '7', onPressed: () => _input('7')),
              _InputButton(text: '8', onPressed: () => _input('8')),
              _InputButton(text: '9', onPressed: () => _input('9')),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              ref.watch(canCheckBiometricsProvider) ?? false ? _InputButton(icon: Assets.fingerprint, onPressed: _openBiometrics) : widget.confirm != null ? _InputButton(icon: Assets.checkUnderlined, onPressed: widget.confirm!) : const SizedBox(width: 80),

              _InputButton(text: '0', onPressed: () => _input('0')),
              _InputButton(icon: Assets.backspace, onPressed: _backspace),
            ],
          ),
        ],
      ),
    );
  }

  void _input(String text) async {
    final List pin = ref.watch(pinValuesProvider);
    if (pin.length < widget.pinMaxChars) {
      ref.read(pinErrorProvider.notifier).update((state) => false);
      ref.read(pinValuesProvider.notifier).update((state) => [...state, text]);
      ref.read(pinShowValuesProvider.notifier).update((state) => [...state.map((e) => false), true]);
      Future.delayed(
        const Duration(milliseconds: 500),
        () {
          ref.read(pinShowValuesProvider.notifier).update((state) => [...state.map((e) => false)]);
        },
      );
    }
    widget.callback();
  }

  void _backspace() {
    ref.read(pinValuesProvider.notifier).update((state) => []);
    ref.read(pinShowValuesProvider.notifier).update((state) => []);
  }

  void _openBiometrics() async {
    await ref.watch(localAuthProvider).authenticateWithBiometrics().then((value) async {
      if (value) {
        final bool isAuthed = await ref.watch(authRepo).refreshToken();
        if (isAuthed) {
          if (mounted) context.goNamed(RouteNames.clientDashboardRouteName);
        } else {
          if (mounted) context.goNamed(RouteNames.loginRouteName);
        }
      }
      return;
    });
  }
}

class _InputButton extends StatelessWidget {
  final String? text;
  final String? icon;
  final VoidCallback onPressed;

  const _InputButton({
    this.text,
    this.icon,
    required this.onPressed,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(vertical: 8),
      height: 80,
      width: 80,
      child: TextButton(
        onPressed: onPressed,
        style: ElevatedButton.styleFrom(
          foregroundColor: const Color(0xFF6B6B6B),
          backgroundColor: const Color(0x3AC7C7C7),
          shape: const CircleBorder(),
          //  padding: const EdgeInsets.all(30),
        ),
        child: text != null
            ? Text(
                text!,
                textAlign: TextAlign.justify,
                style: GoogleFonts.nunito(
                  fontSize: 26,
                  fontWeight: FontWeight.w400,
                  color: const Color(0xFF252525),
                ),
              )
            : Image.asset(icon!, width: 26),
      ),
    );
  }
}
