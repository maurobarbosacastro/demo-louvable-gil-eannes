import 'package:flutter/material.dart';
import 'package:reactive_forms/reactive_forms.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class TgpkReactiveTextField extends StatelessWidget {
  final FormControl<String> formControl;
  final String? label;
  final TextInputType keyboardType;
  final Map<String, String Function(Object)>? validationMessages;

  const TgpkReactiveTextField({
    super.key,
    required this.formControl,
    this.label,
    this.keyboardType = TextInputType.text,
    this.validationMessages
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        if (label != null) Text(label!, style: Theme.of(context).textTheme.titleSmall),
        ReactiveValueListenableBuilder<String>(
          formControl: formControl,
          builder: (context, value, child) {
            final isEmpty = value.isNullOrEmpty;
            return ReactiveTextField<String>(
              formControl: formControl,
              decoration: InputDecoration(
                contentPadding: const EdgeInsets.only(bottom: 10.0),
                isCollapsed: true,
                focusedBorder: UnderlineInputBorder(
                  borderSide: BorderSide(
                    color: isEmpty ? AppColors.wildSand : AppColors.licorice,
                  ),
                ),
                enabledBorder: UnderlineInputBorder(
                  borderSide: BorderSide(
                    color: isEmpty ? AppColors.wildSand : AppColors.licorice,
                  ),
                ),
              ),
              keyboardType: keyboardType,
              validationMessages: validationMessages,
            );
          },
        )
      ],
    );
  }
}
