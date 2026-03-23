import 'package:flutter/material.dart';

import '../../utils/constants/colors.dart';

enum Mode {
  primary,
  blue,
  outlinedDark,
  outlinedWhite,
  red,
  white,
  error,
  outlinedRed,
  youngBlue
}

class Button extends StatelessWidget {
  final VoidCallback? onPressed;
  final String label;
  final Mode mode;
  final Widget? icon;
  final double width;
  final double height;
  final bool isRounded;
  final double? fontSize;
  final EdgeInsetsGeometry? padding;

  const Button(
      {super.key,
      this.onPressed,
      required this.label,
      this.mode = Mode.primary,
      this.icon,
      this.width = double.infinity,
      this.height = 34,
      this.isRounded = false,
      this.fontSize,
      this.padding});

  static final Map<Mode, ButtonStyle> _buttonStyles = {
    Mode.primary: ButtonStyles.primaryButtonStyle,
    Mode.blue: ButtonStyles.blueButtonStyle,
    Mode.outlinedDark: ButtonStyles.outlinedDarkButtonStyle,
    Mode.outlinedWhite: ButtonStyles.outlinedWhiteButtonStyle,
    Mode.red: ButtonStyles.red,
    Mode.white: ButtonStyles.whiteButtonStyle,
    Mode.error: ButtonStyles.errorStyle,
    Mode.outlinedRed: ButtonStyles.outlinedRedStyle,
    Mode.youngBlue: ButtonStyles.youngBlueStyle,
  };

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: onPressed,
        style: _buttonStyles[mode]?.copyWith(
          minimumSize: WidgetStateProperty.all( Size(0, height)),
          maximumSize: WidgetStateProperty.all(Size(width, height)),
          alignment: Alignment.center,
          shape: WidgetStateProperty.all(
            RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(isRounded ? 27.0 : 4.0),
            )
          ),
          padding: WidgetStateProperty.all(
            padding ?? _buttonStyles[mode]?.padding?.resolve({}),
          ),
          tapTargetSize: MaterialTapTargetSize.shrinkWrap,
          textStyle: WidgetStateProperty.all(
            TextStyle(
              fontSize: fontSize ?? _buttonStyles[mode]?.textStyle?.resolve(<WidgetState>{})?.fontSize,
            ),
          ),
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          spacing: 4,
          children: [
            if (icon != null) icon!,
            Text(label)
          ],
        ),
    );
  }
}

// Button styles
class ButtonStyles {
  static final ButtonStyle primaryButtonStyle = ElevatedButton.styleFrom(
    backgroundColor: AppColors.licorice,
    foregroundColor: Colors.white,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 20.0),
    textStyle: const TextStyle(fontSize: 15.0),
  );

  static final ButtonStyle blueButtonStyle = ElevatedButton.styleFrom(
    backgroundColor: AppColors.blue,
    foregroundColor: Colors.white,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 15.0),
    textStyle: const TextStyle(fontSize: 15.0),
  );

  static final ButtonStyle outlinedDarkButtonStyle = ElevatedButton.styleFrom(
    backgroundColor: Colors.white,
    foregroundColor: AppColors.licorice,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 10.0),
    textStyle: const TextStyle(fontSize: 15.0),
    side: const BorderSide(
      color: AppColors.licorice,
      width: 1,
    ),
  );

  static final ButtonStyle outlinedWhiteButtonStyle = ElevatedButton.styleFrom(
    backgroundColor: Colors.transparent,
    foregroundColor: Colors.white,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 10.0),
    textStyle: const TextStyle(fontSize: 13.0),
    side: const BorderSide(
      color: Colors.white,
      width: 1,
    ),
  );

  static final ButtonStyle red = ElevatedButton.styleFrom(
    backgroundColor: AppColors.radicalRed,
    foregroundColor: Colors.white,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 11.0),
    textStyle: const TextStyle(fontSize: 15.0),
  );

  static final ButtonStyle whiteButtonStyle = ElevatedButton.styleFrom(
    backgroundColor: Colors.white,
    foregroundColor: AppColors.licorice,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 10.0),
    textStyle: const TextStyle(fontSize: 13.0),
  );

  static final ButtonStyle errorStyle = ElevatedButton.styleFrom(
    backgroundColor: Colors.white,
    foregroundColor: Colors.red,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 10.0),
    textStyle: const TextStyle(fontSize: 13.0),
    side: const BorderSide(
      color: Colors.red,
      width: 1,
    ),
  );

  static final ButtonStyle outlinedRedStyle = ElevatedButton.styleFrom(
    backgroundColor: Colors.white,
    foregroundColor: AppColors.radicalRed,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 10.0),
    textStyle: const TextStyle(fontSize: 15.0),
    side: const BorderSide(
      color: AppColors.radicalRed,
      width: 1,
    ),
  );

  static final ButtonStyle youngBlueStyle = ElevatedButton.styleFrom(
    backgroundColor: AppColors.youngBlue,
    foregroundColor: Colors.white,
    shadowColor: Colors.transparent,
    padding: const EdgeInsets.symmetric(horizontal: 10.0),
    textStyle: const TextStyle(fontSize: 13.0),
  );
}
