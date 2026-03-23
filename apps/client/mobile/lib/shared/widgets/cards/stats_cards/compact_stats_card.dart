import 'package:flutter/material.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class CompactStatsCard extends StatelessWidget {
  final String title;
  final Widget icon;
  final String value;
  final bool isRowHeader;

  const CompactStatsCard({Key? key, required this.title, required this.icon, required this.value, this.isRowHeader = false}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(bottom: 10),
      decoration: const BoxDecoration(
        border: Border(
          bottom: BorderSide(color: AppColors.wildSand)
        )
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          isRowHeader ? Row(
            spacing: 5,
            children: [
              icon,
              Text(
                title,
                style: Theme.of(context).textTheme.titleSmall?.copyWith(color: AppColors.licorice),
              ),
            ]
          ) : Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            spacing: 5,
              children: [
                icon,
                Text(
                  title,
                  style: Theme.of(context).textTheme.titleSmall?.copyWith(color: AppColors.licorice),
                ),
              ]
          ),
          const SizedBox(height: 7),
          Text(
            value,
            style: Theme.of(context).textTheme.headlineLarge,
          ),
        ],
      ),
    );
  }
}
