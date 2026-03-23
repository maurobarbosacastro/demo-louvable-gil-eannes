import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_dialog.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class StopRewardDialog extends StatelessWidget {
  final void Function(BuildContext ctx) onConfirm;

  const StopRewardDialog({super.key, required this.onConfirm});

  @override
    Widget build(BuildContext context) {
      return Column(
        spacing: 10,
        children: [
          Image.asset(Assets.stopReward),
          Text(translation(context).stopReward,
              style: Theme.of(context).textTheme.titleLarge),
          Text(
              translation(context).stopRewardExplanation,
              style: Theme.of(context).textTheme.titleSmall, textAlign: TextAlign.center),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            spacing: 10,
            children: [
              Button(
                  label: translation(context).cancel,
                  mode: Mode.outlinedDark,
                  onPressed: () => GenericDialog.closeDialog(context)),
              Button(label: translation(context).yesStop, mode: Mode.red, onPressed: () => onConfirm(context))
            ],
          )
        ],
      );
    }
}
