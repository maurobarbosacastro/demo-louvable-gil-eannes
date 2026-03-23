import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../utils/constants/colors.dart';

class TAppTheme {
  TAppTheme._();

  static ThemeData theme = ThemeData(
    fontFamily: GoogleFonts.ibmPlexSans().fontFamily,
    textTheme: TextTheme(
      headlineLarge: const TextStyle(fontSize: 32),
      headlineMedium: const TextStyle(fontSize: 25, color: AppColors.licorice),
      headlineSmall: const TextStyle(fontSize: 20, color: AppColors.licorice),
      titleLarge: GoogleFonts.ibmPlexSans(
        fontSize: 16,
        color: AppColors.licorice,
        fontWeight: FontWeight.w700,
      ),
      titleSmall: const TextStyle(fontSize: 13, color: AppColors.waterloo),
      bodyMedium: const TextStyle(fontSize: 15, color: AppColors.waterloo),
      bodySmall: const TextStyle(fontSize: 12, color: AppColors.waterloo),
    ),
    colorScheme: const ColorScheme(
      brightness: Brightness.light,
      primary: Color(0xFF32334C),
      onPrimary: Colors.white,
      secondary: Color(0xFF7A7B92),
      onSecondary: Colors.white,
      surface: Colors.white,
      onSurface: Color(0xFF32334C),
      error: Color(0xFFE53935),
      onError: Colors.white,
    ),
  );
}
