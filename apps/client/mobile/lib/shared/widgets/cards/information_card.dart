import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class InformationCard extends StatelessWidget {
  final Object title;
  final String description;
  final String? question;
  final VoidCallback? handleQuestionClick;
  final Widget content;

  // Styling properties
  final Color backgroundColor;
  final Color color;

  const InformationCard({super.key, required this.title, required this.description, this.question, this.handleQuestionClick, required this.content, this.backgroundColor = AppColors.grandis, this.color = AppColors.licorice});

  _buildTitle(BuildContext context) {
    if (title is String) {
      return Text(
        title.toString(),
        style: Theme.of(context).textTheme.headlineSmall?.copyWith(color: color),
      );
    }
    return title;
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 30),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(20),
        color: backgroundColor
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          _buildTitle(context),
          const SizedBox(height: 7),
          RichText(
            text: TextSpan(
              style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: color),
              children: [
                TextSpan(
                  text: description,
                ),
                TextSpan(
                  text: question,
                  style: TextStyle(
                    decoration: TextDecoration.underline,
                    decorationColor: color
                  ),
                  recognizer: TapGestureRecognizer()
                    ..onTap = () {
                      handleQuestionClick?.call();
                    },
                ),
              ],
            ),
          ),
          const SizedBox(height: 15),
          content
        ],
      ),
    );
  }
}
