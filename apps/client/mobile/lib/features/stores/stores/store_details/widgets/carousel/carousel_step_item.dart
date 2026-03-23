import 'package:flutter/material.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class CarouselStepItem extends StatelessWidget {
  final int index;
  final String label;
  const CarouselStepItem({super.key, required this.index, required this.label});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: AppColors.wildSand,
        borderRadius: BorderRadius.circular(20)
      ),
      padding: const EdgeInsets.symmetric(horizontal: 15),
      child: Row(
        spacing: 10,
        mainAxisSize: MainAxisSize.min,
        children: [
          Container(
            height: 40,
            width: 40,
            decoration: const BoxDecoration(
              shape: BoxShape.circle,
              color: AppColors.grandis,
            ),
            child: Center(child: Text('0$index', style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))),
          ),
          Expanded(child: Text(label, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)))
        ],
      ),
    );
  }
}
