import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class MembershipPill extends StatelessWidget {
  final String membership;

  const MembershipPill({super.key, required this.membership});

  Map<String, (Color, String)> _styling(BuildContext context) => {
    'base': (AppColors.radicalRed, translation(context).base),
    'silver': (AppColors.slateGray, translation(context).silver),
    'gold': (AppColors.goldenOrange, translation(context).gold),
  };

  @override
  Widget build(BuildContext context) {
    final (backgroundColor, label) = _styling(context)[membership]!;

    return Container(
      decoration: BoxDecoration(
        color: backgroundColor,
        borderRadius: BorderRadius.circular(25),
      ),
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 2.5),
      child: Text(
        label,
        style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: Colors.white),
      ),
    );
  }
}
