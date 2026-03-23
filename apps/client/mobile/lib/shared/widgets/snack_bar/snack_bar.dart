import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

enum SnackBarType {
  success, error, info
}

Map<SnackBarType, (BoxDecoration, Color, SvgPicture)> _styles = {
  SnackBarType.error: (
  BoxDecoration(
    color: const Color(0xFFfcd4de),
    borderRadius: BorderRadius.circular(4),
  ),
  AppColors.radicalRed,
  SvgPicture.asset(
    Assets.info,
    colorFilter: const ColorFilter.mode(AppColors.radicalRed, BlendMode.srcIn),
  )),
  SnackBarType.info: (
  BoxDecoration(
    color: const Color(0xFFe4e5e9),
    borderRadius: BorderRadius.circular(4),
  ),
  AppColors.waterloo,
  SvgPicture.asset(Assets.info, colorFilter: const ColorFilter.mode(AppColors.waterloo, BlendMode.srcIn)),
  ),
  SnackBarType.success: (
  BoxDecoration(
    color: const Color(0xFF4237DA),
    borderRadius: BorderRadius.circular(4),
  ),
  AppColors.wildSand,
  SvgPicture.asset(Assets.circleCheck, colorFilter: const ColorFilter.mode(AppColors.wildSand, BlendMode.srcIn))),
};

void showSnackBar(BuildContext context, String message, SnackBarType type) {
  final overlay = Overlay.of(context);
  late OverlayEntry overlayEntry;

  final (boxDecoration, textColor, icon) = _styles[type]!;

  overlayEntry = OverlayEntry(
    builder: (context) => Positioned(
      bottom: 100,
      left: 20,
      right: 20,
      child: Material(
        color: Colors.transparent,
        child: Container(
          padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
          decoration: boxDecoration,
          child: Row(
            children: [
              icon,
              const SizedBox(width: 12),
              Expanded(
                child: Text(
                  message,
                  style: Theme.of(context)
                      .textTheme
                      .titleSmall
                      ?.copyWith(color: textColor),
                  softWrap: true,
                ),
              ),
            ],
          ),
        ),
      ),
    ),
  );

  overlay.insert(overlayEntry);

  Timer(const Duration(milliseconds: 1500), () {
    overlayEntry.remove();
  });
}
