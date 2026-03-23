import 'package:flutter/material.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_bottom_dialog.dart';

class GenericDialog extends StatelessWidget {
  final Widget content;
  final bool hasCloseIcon;

  const GenericDialog({
    super.key,
    required this.content,
    this.hasCloseIcon = false
  });

  static Future<void> openDialog(BuildContext context,
      {
        required Widget content,
        bool canDismissOnTapOutside = true,
        bool hasCloseIcon = false
      }) {
    return showDialog(
      context: context,
      barrierDismissible: canDismissOnTapOutside,
      builder: (BuildContext dialogContext) {
        return GenericBottomDialog(
            content: content,
            hasCloseIcon: hasCloseIcon
        );
      },
    );
  }

  static void closeDialog(BuildContext context) {
    if (Navigator.canPop(context)) {
      Navigator.pop(context);
    }
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      backgroundColor: Colors.white,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
      title: hasCloseIcon ?
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              IconButton(
                icon: const Icon(Icons.close, size: 20),
                onPressed: () => closeDialog(context)
              ),
            ],
          ) : null,
      content: IntrinsicHeight(child: content),
    );
  }
}
