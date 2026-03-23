import 'package:flutter/material.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_dialog.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';

class WithdrawalDialog extends StatelessWidget {
  const WithdrawalDialog({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisSize: MainAxisSize.min,
        spacing: 20,
        children: [
          SvgPicture.asset(Assets.info, height: 50),
          Text(translation(context).infoWithdraw, textAlign: TextAlign.center),
          Center(
            child: IntrinsicWidth(
              child: Button(
                label: translation(context).gotIt,
                mode: Mode.outlinedDark,
                onPressed: () => GenericDialog.closeDialog(context),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
