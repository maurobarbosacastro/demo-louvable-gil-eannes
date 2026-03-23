import 'package:flutter/material.dart';

class PasswordTextLineWidget extends StatelessWidget {
  const PasswordTextLineWidget({super.key, required this.check, required this.errorMessage});
  final bool check;
  final String errorMessage;
  @override
  Widget build(BuildContext context) {
    return Row(children: [
      check
          ? const Icon(
              Icons.check_circle,
              color: Colors.green,
              size: 16,
            )
          : const Icon(
              Icons.error,
              color: Colors.grey,
              size: 16,
            ),
      const SizedBox(
        width: 5,
      ),
      Text(
        errorMessage,
        style: TextStyle(
          fontSize: 12,
          fontWeight: FontWeight.w500,
          color: check ? Colors.green : Colors.grey,
        ),
      ),
    ]);
  }
}
