import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class AuthButtonWidget extends StatelessWidget {
  final VoidCallback? onPressed;
  final bool isLoading;
  final String label;

  const AuthButtonWidget({super.key, required this.onPressed, required this.isLoading, required this.label});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 100,
      padding: const EdgeInsets.symmetric(vertical: 24),
      child: ElevatedButton(
        onPressed: onPressed,
        style: TextButton.styleFrom(shape: const StadiumBorder(), backgroundColor: const Color(0XFF1C1C1C)),
        child: isLoading
            ? const CircularProgressIndicator(color: Colors.white)
            : Text(label, style: GoogleFonts.poppins(fontWeight: FontWeight.w700, fontSize: 12)),
      ),
    );
  }
}
