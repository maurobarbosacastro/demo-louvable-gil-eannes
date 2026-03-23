import 'package:flutter/material.dart';

class WithoutResults extends StatelessWidget {
  final String title;
  final String description;
  const WithoutResults({super.key, required this.title, required this.description});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      spacing: 10,
      children: [
        Text(title, style: Theme.of(context).textTheme.titleLarge),
        Text(description, style: Theme.of(context).textTheme.titleSmall)
      ]
    );
  }
}
