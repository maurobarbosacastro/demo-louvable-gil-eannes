import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_dialog.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class DeletePaymentMethodDialog extends StatelessWidget {
  final void Function(BuildContext ctx) onDelete;
  const DeletePaymentMethodDialog({super.key, required this.onDelete});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 300,
      padding: const EdgeInsets.only(top: 15),
      child: Column(
        children: [
          Text(translation(context).deletePayment, style: TextTheme.of(context).titleLarge),
          const SizedBox(height: 10),
          Text(translation(context).actionCanNotBeUndone, style: TextTheme.of(context).titleSmall),
          const SizedBox(height: 15),
          Row(
            spacing: 10,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Button(label: translation(context).cancel, mode: Mode.outlinedDark, onPressed: () => GenericDialog.closeDialog(context)),
              Button(label: translation(context).yesDelete, mode: Mode.red, onPressed: () => onDelete(context))
            ],
          ),
          const SizedBox(height: 30),
          Container(
            height: 1,
            color: AppColors.wildSand,
          ),
          const SizedBox(height: 15),
          Align(
            alignment: Alignment.centerLeft,
            child: Text(translation(context).readFaqs,
                style: TextTheme.of(context).bodySmall)
          ),
        ],
      ),
    );
  }
}
