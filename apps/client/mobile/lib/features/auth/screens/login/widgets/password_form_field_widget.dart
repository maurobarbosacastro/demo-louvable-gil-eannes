import 'package:flutter/material.dart';
import '../../../../../utils/constants/colors.dart';

class PasswordFormFieldWidget extends StatelessWidget {
  final TextEditingController controller;
  final String? Function(String? value) validator;
  final String? isError;
  final String placeholder;
  final bool obscureText;
  final bool enabled;
  final VoidCallback? showPassword;
  final Function(String value)? onChanged;
  final bool suffixIcon;
  final String label;
  final bool readOnly;

  const PasswordFormFieldWidget({
    super.key,
    required this.controller,
    required this.validator,
    required this.placeholder,
    required this.obscureText,
    this.showPassword,
    this.isError,
    this.onChanged,
    this.suffixIcon = true,
    this.enabled = true,
    required this.label,
    this.readOnly = false
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
            controller: controller,
            enabled: enabled,
            readOnly: readOnly,
            cursorColor: Colors.black,
            obscureText: obscureText,
            keyboardType: TextInputType.text,
            decoration: InputDecoration(
              contentPadding: const EdgeInsets.only(bottom: 10.0),
              isCollapsed: true,
              errorStyle: isError == "" ? const TextStyle(height: 0, fontSize: 0): null,
              errorText: isError,
              errorMaxLines: 2,
              hintText: placeholder,
              suffixIconConstraints: const BoxConstraints(
                minHeight: 10,
                minWidth: 10,
              ),
              suffixIcon: suffixIcon
                  ? GestureDetector(
                onTap: showPassword,
                child: obscureText
                    ? const Icon(
                  Icons.visibility_off_outlined,
                  color: Color(0xFF3F3F3F),
                  size: 15.0,
                )
                    : const Icon(
                  Icons.visibility_outlined,
                  color: Color(0xFF3F3F3F),
                  size: 15.0,
                ),
              )
                  : null,
              floatingLabelBehavior: FloatingLabelBehavior.never,
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
            ),
            onChanged: (value) {
              (context as Element).markNeedsBuild();
              onChanged?.call;
            },
            validator: validator,
          ),
        ],
      ),
    );
  }
}
