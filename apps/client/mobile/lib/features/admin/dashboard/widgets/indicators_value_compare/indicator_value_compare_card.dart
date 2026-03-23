import 'package:flutter/material.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class IndicatorValueCompareCard extends StatelessWidget {
  final String title;
  final Widget content;
  final String trailing;

  const IndicatorValueCompareCard({
    super.key,
    required this.title,
    required this.content,
    required this.trailing
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(left: 10),
      decoration: const BoxDecoration(
        border: Border(
          left: BorderSide(color: AppColors.wildSand),
        ),
      ),
      child: Column(
        spacing: 10,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(trailing, style: Theme.of(context).textTheme.headlineSmall?.copyWith(color: AppColors.waterloo)),
          content,
          Text(title, style: Theme.of(context).textTheme.titleSmall)
        ],
      ),
    );
  }
}
