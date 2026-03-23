import 'package:flutter/material.dart';
import 'package:tagpeak/utils/constants/colors.dart';

enum StatsCardMode {
  transparentUnderline,
  blue,
  yellow,
  transparentSideline
}

class StatsCard extends StatelessWidget {
  final String title;
  final Widget? icon;
  final String value;
  final bool isRowHeader;
  final Widget? content;
  final StatsCardMode mode;

  const StatsCard({
    super.key,
    required this.title,
    this.icon,
    required this.value,
    this.isRowHeader = false,
    this.content,
    this.mode = StatsCardMode.transparentUnderline,
  });

  static final Map<StatsCardMode, (BoxDecoration, Color, EdgeInsets)> _styles = {
    StatsCardMode.transparentUnderline: (
      const BoxDecoration(
        border: Border(
          bottom: BorderSide(color: AppColors.wildSand),
        ),
      ),
      AppColors.licorice,
      const EdgeInsets.only(bottom: 10, top: 10),
    ),
    StatsCardMode.blue: (
      BoxDecoration(
        color: AppColors.youngBlue,
        borderRadius: BorderRadius.circular(10),
      ),
      AppColors.wildSand,
      const EdgeInsets.all(10),
    ),
    StatsCardMode.yellow: (
      BoxDecoration(
        color: AppColors.grandis,
        borderRadius: BorderRadius.circular(10),
      ),
      AppColors.licorice,
      const EdgeInsets.only(left: 10, top: 10, right: 10, bottom: 0),
    ),
    StatsCardMode.transparentSideline: (
    const BoxDecoration(
      border: Border(
        left: BorderSide(color: AppColors.wildSand),
      ),
    ),
    AppColors.licorice,
    const EdgeInsets.only(left: 10),
    ),
  };

  @override
  Widget build(BuildContext context) {
    final (boxDecoration, textColor, padding) = _styles[mode]!;
    return Container(
      padding: padding,
      decoration: boxDecoration,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          isRowHeader
              ? Row(
            children: [
              icon != null ? icon! : const SizedBox.shrink(),
              const SizedBox(width: 5),
              Text(
                title,
                style: Theme.of(context)
                    .textTheme
                    .titleSmall
                    ?.copyWith(color: textColor),
              ),
            ],
          )
              : Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              icon != null ? icon! : const SizedBox.shrink(),
              const SizedBox(height: 5),
              Text(
                title,
                style: Theme.of(context)
                    .textTheme
                    .titleSmall
                    ?.copyWith(color: textColor),
              ),
            ],
          ),
          const SizedBox(height: 10),
          Text(
            value,
            style: Theme.of(context)
                .textTheme
                .headlineLarge
                ?.copyWith(color: textColor),
          ),
          const SizedBox(height: 5),
          if (content != null) content!,
        ],
      ),
    );
  }
}
