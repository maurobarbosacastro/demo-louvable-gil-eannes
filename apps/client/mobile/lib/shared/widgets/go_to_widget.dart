import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:google_fonts/google_fonts.dart';

class GoToWidget extends StatelessWidget {
  final String text;
  final String buttonText;
  final String route;
  const GoToWidget({super.key, required this.text, required this.buttonText, required this.route});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 40,
      width: MediaQuery.of(context).size.width,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(text),
          TextButton(
            onPressed: () => context.goNamed(route),
            style: TextButton.styleFrom(foregroundColor: Colors.yellow),
            child: Text(
              buttonText,
              style: GoogleFonts.lexend(fontSize: 14, fontWeight: FontWeight.w700, color: Colors.black),
            ),
          ),
        ],
      ),
    );
  }
}


