import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

import '../widgets/password_text_line_widget.dart';

class PasswordRequirementsComponent extends StatelessWidget {
  final List list;
  final bool isVisible;
  const PasswordRequirementsComponent({super.key, required this.list, required this.isVisible});

  @override
  Widget build(BuildContext context) {
    return Visibility(
      visible: isVisible, //ref.watch(isVisibleProvider),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          Padding(
            padding: const EdgeInsets.only(top: 6, bottom: 6),
            child: Text(
              AppLocalizations.of(context)!.passwordRequirements,
              style: const TextStyle(
                fontSize: 14,
                fontWeight: FontWeight.w500,
              ),
            ),
          ),
          PasswordTextLineWidget(check: list[0], errorMessage: AppLocalizations.of(context)!.passwordSize),
          PasswordTextLineWidget(check: list[1], errorMessage: AppLocalizations.of(context)!.passwordSpecialChar),
          PasswordTextLineWidget(check: list[2], errorMessage: AppLocalizations.of(context)!.passwordNumber),
          PasswordTextLineWidget(check: list[3], errorMessage: AppLocalizations.of(context)!.passwordCapital),
          const SizedBox(height: 16),
        ],
      ),
    );
  }
}
