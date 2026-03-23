import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class SimpleAlertExample extends StatelessWidget {
  final String title;
  final String body;
  final String buttonText;

  const SimpleAlertExample({super.key, required this.title, required this.body, required this.buttonText});

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Column(
        children: [
          Text(
            title,
            style: const TextStyle(
              fontSize: 24,
              fontWeight: FontWeight.w600,
            ),
          ),
          const SizedBox(height: 8.0),
          Text(body),
        ],
      ),
      actions: [
        ElevatedButton(
          onPressed: () => context.pop(),
          child: Text(buttonText),
        ),
      ],
      actionsAlignment: MainAxisAlignment.center,
    );
  }
}
