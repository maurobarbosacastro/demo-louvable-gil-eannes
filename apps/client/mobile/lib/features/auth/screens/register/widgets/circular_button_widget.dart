import 'package:flutter/material.dart';

class CircularButton extends StatelessWidget {
  final VoidCallback onTap;
  final Icon icon;
  final String label;
  const CircularButton({super.key, required this.onTap, required this.icon, required this.label});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        InkWell(
          onTap: onTap,
          borderRadius: const BorderRadius.all(Radius.circular(100)),
          child: Container(
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                border: Border.all(color: Colors.grey, width: 2),
              ),
              alignment: Alignment.center,
              width: 48,
              height: 48,
              padding: const EdgeInsets.all(5),
              child: icon),
        ),
        const SizedBox(height: 8.0),
        Text(
          label,
          style: const TextStyle(
            color: Colors.black,
            fontFamily: "Lexend",
            fontSize: 16,
            fontWeight: FontWeight.w500,
          ),
        ),
      ],
    );
  }
}
