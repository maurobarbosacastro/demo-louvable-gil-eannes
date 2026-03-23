import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class IconOffset {
  final double top;
  final double right;

  const IconOffset({required this.top, required this.right});
}

class GenericBottomDialog extends StatelessWidget {
  final Widget content;
  final bool hasCloseIcon;
  final EdgeInsetsGeometry padding;
  final IconOffset iconOffset;

  const GenericBottomDialog({
    super.key,
    required this.content,
    this.hasCloseIcon = false,
    this.padding = const EdgeInsets.symmetric(horizontal: 20, vertical: 21),
    this.iconOffset = const IconOffset(top: 0, right: 0),
  });

  static Future<void> openBottomDialog(BuildContext context,
      {
        required Widget content,
        bool canDismissOnTapOutside = true,
        bool hasCloseIcon = false,
        EdgeInsetsGeometry padding = const EdgeInsets.symmetric(horizontal: 20, vertical: 21),
        IconOffset iconOffset = const IconOffset(top: 0, right: 0),
      }
      ) {
    return showModalBottomSheet(
      context: context,
      enableDrag: false,
      isDismissible: canDismissOnTapOutside,
      isScrollControlled: true,
      useRootNavigator: true,
      builder: (BuildContext dialogContext) {
        return GenericBottomDialog(
            content: content,
            hasCloseIcon: hasCloseIcon,
            padding: padding,
            iconOffset: iconOffset
        );
      },
    );
  }

  static void closeBottomDialog(BuildContext context) {
    if (Navigator.canPop(context)) {
      Navigator.pop(context);
    }
  }

  @override
  Widget build(BuildContext context) {
    final maxHeight = MediaQuery.of(context).size.height * 0.9;

    return Container(
      padding: padding,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(20),
        color: Colors.white,
      ),
      constraints: BoxConstraints(
        maxHeight: maxHeight,
      ),
      child: IntrinsicHeight(
        child: Stack(
          children: [
            content,
            Visibility(
              visible: hasCloseIcon,
              child: Positioned(
                top: iconOffset.top,
                right: iconOffset.right,
                child: GestureDetector(
                  child: SvgPicture.asset(Assets.close),
                  onTap: () => closeBottomDialog(context),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
