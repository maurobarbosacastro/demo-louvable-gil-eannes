import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/screens/login/provider/login_provider.dart';
import 'package:tagpeak/features/auth/screens/login/widgets/text_form_field_widget.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/logging.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';

class ResetPasswordScreen extends ConsumerStatefulWidget {
  const ResetPasswordScreen({super.key});

  @override
  ConsumerState<ResetPasswordScreen> createState() => _ResetPasswordScreenState();
}

class _ResetPasswordScreenState extends ConsumerState<ResetPasswordScreen> {
  final TextEditingController _emailController = TextEditingController();
  final _formKey = GlobalKey<FormState>();

  String? getResetPasswordError(String? message) {
    switch(message) {
      case 'emailNotFound':
        return translation(context).emailNotFound;
      case 'defaultError':
        return translation(context).defaultError;
      default:
        return message;
    }
  }

  Future<void> resetPassword(String email, GlobalKey<FormState> formKey) async {
    ref.read(resetPasswordErrorProvider.notifier).update((state) => FormErrorModel());
    if (formKey.currentState != null && formKey.currentState!.validate()) {
      ref.read(userNotifier).resetPassword(email).then((isSuccess) async {
        if (isSuccess) {
          context.goNamed(RouteNames.confirmEmailRouteName, queryParameters: {'isPasswordReset': 'true'});
        } else {
          Log.warning("[login_provider > resetPassword] Something went wrong.");
        }
      });
    } else {
      ref.read(resetPasswordErrorProvider.notifier).update(
            (state) => FormErrorModel(
          isError: true,
          errorMessage: null,
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20.0, vertical: 40.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(translation(context).forgotPassword, style: Theme.of(context).textTheme.headlineMedium),
          const SizedBox(height: 5.0),
          Text(
            translation(context).noWorries,
            style: Theme.of(context).textTheme.bodyMedium,
          ),
          const SizedBox(height: 45),
          Form(
            key: _formKey,
            child: TextFormFieldWidget(
              isError: ref.watch(resetPasswordErrorProvider).isError,
              errorLabel: getResetPasswordError(ref.watch(resetPasswordErrorProvider).errorMessage),
              controller: _emailController,
              validator: (value) => value!.isEmpty ? '' : null,
              placeholder: translation(context).yourEmail,
              label: translation(context).email,
            ),
          ),
          const SizedBox(height: 15),
          Button(
            onPressed: () =>
              resetPassword(_emailController.text, _formKey),
            label: translation(context).reset,
          ),
        ],
      ),
    );
  }
}
