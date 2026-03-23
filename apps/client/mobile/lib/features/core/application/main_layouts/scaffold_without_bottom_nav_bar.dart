import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/header.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import '../translation.dart';

enum TrailingHeaderButton {
  darkMode,
  whiteMode,
}

class ScaffoldWithoutBottomNavBar extends StatelessWidget {
  final Widget child;
  final bool extendBodyBehindAppBar;
  final TrailingHeaderButton? trailingHeaderButton;
  final VoidCallback? onPressedTrailingButton;
  const ScaffoldWithoutBottomNavBar({super.key, required this.child, this.extendBodyBehindAppBar = false, this.trailingHeaderButton, this.onPressedTrailingButton});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: SafeArea(
          child: Scaffold(
            extendBodyBehindAppBar: extendBodyBehindAppBar,
            appBar: Header(
              endContent: Visibility(
                visible: trailingHeaderButton != null,
                child: Button(
                  label: translation(context).backToLogin,
                  icon: SvgPicture.asset('lib/assets/icons/simple_arrow_left.svg', colorFilter: ColorFilter.mode(trailingHeaderButton == TrailingHeaderButton.darkMode ? AppColors.licorice : Colors.white, BlendMode.srcIn)),
                  mode: trailingHeaderButton == TrailingHeaderButton.darkMode ? Mode.outlinedDark : Mode.outlinedWhite,
                  onPressed: () {
                    onPressedTrailingButton?.call();
                    context.go('/login');
                  },
                  height: 30,
                )
              )
            ),
            body: child
          )
      ),
    );
  }
}
