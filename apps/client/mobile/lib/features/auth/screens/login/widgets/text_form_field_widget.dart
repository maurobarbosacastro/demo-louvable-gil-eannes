import 'package:flutter/services.dart';
import 'package:tagpeak/utils/validators/lower_case_formatter.dart';
import 'package:flutter/material.dart';

import '../../../../../utils/constants/colors.dart';

class TextFormFieldWidget extends StatelessWidget {
  final TextEditingController controller;
  final String? Function(String? value) validator;
  final String placeholder;
  final TextStyle? errorStyle;
  final bool isError;
  final bool enabled;
  final TextInputType textInputType;
  final String label;
  final String? errorLabel;
  final bool readOnly;
  final Widget? suffix;
  final List<TextInputFormatter>? formatters;

  const TextFormFieldWidget({
    super.key,
    required this.controller,
    required this.validator,
    required this.placeholder,
    this.isError = false,
    this.errorStyle,
    this.enabled = true,
    this.textInputType = TextInputType.emailAddress,
    required this.label,
    this.errorLabel,
    this.readOnly = false,
    this.suffix,
    this.formatters
  });
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 10.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(label, style: Theme.of(context).textTheme.titleSmall),
          TextFormField(
            inputFormatters: formatters ?? [LowerCaseTextFormatter()],
            controller: controller,
            enabled: enabled,
            readOnly: readOnly,
            cursorColor: Colors.black,
            textInputAction: TextInputAction.next,
            keyboardType: textInputType,
            decoration: InputDecoration(
              contentPadding: const EdgeInsets.only(bottom: 10.0),
              isCollapsed: true,
              errorStyle: errorStyle,
              errorText: isError ? errorLabel : null,
              hintText: placeholder,
              hintStyle: Theme.of(context).textTheme.bodyMedium,
              focusedBorder: UnderlineInputBorder(
                borderSide: BorderSide(
                  color: controller.text.isEmpty ? AppColors.wildSand : AppColors.licorice,
                ),
              ),
              enabledBorder: UnderlineInputBorder(
                borderSide: BorderSide(
                  color: controller.text.isEmpty ? AppColors.wildSand : AppColors.licorice,
                ),
              ),
              floatingLabelBehavior: FloatingLabelBehavior.never,
              suffix: suffix
            ),
            validator: validator,
            onChanged: (value) {
              (context as Element).markNeedsBuild();
            },
          ),
        ],
      ),
    );
  }
}
