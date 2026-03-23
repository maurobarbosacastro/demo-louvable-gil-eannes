import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

// ToDO make it more reusable
void showTopSuccessSnackBar(BuildContext context, String message) {
  ScaffoldMessenger.of(context).showSnackBar(SnackBar(
    duration: const Duration(milliseconds: 1000),
    backgroundColor: Colors.transparent,
    elevation: 0,
    content: Container(
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
      decoration: BoxDecoration(
        gradient: LinearGradient(
          begin: Alignment.topCenter,
          end: Alignment.bottomCenter,
          colors: [
            const Color(0xFF4237DA).withAlpha(51),
            const Color(0xFF4237DA).withAlpha(51),
          ],
        ),
        color: Colors.white,
        borderRadius: BorderRadius.circular(4),
      ),
      child: Row(spacing: 6, children: [
        SvgPicture.asset(Assets.circleCheck),
        Expanded(
          child: Text(
            message,
            style: Theme.of(context)
                .textTheme
                .bodyMedium
                ?.copyWith(color: AppColors.youngBlue),
            softWrap: true,
          ),
        ),
      ]),
    ),
  ));
}
