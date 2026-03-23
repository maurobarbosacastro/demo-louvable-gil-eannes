import 'package:tagpeak/features/auth/repositories/auth_repo.dart';
import 'package:tagpeak/features/core/routes/route_notifier.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../../../core/routes/route_names.dart';
import 'password_validator_notifier.dart';

//provider to save the state of the show/hide password
final showPasswordRegisterProvider = StateProvider<bool>((ref) => true);

final showConfirmPasswordProvider = StateProvider<bool>((ref) => true);

final isVisibleProvider = StateProvider<bool>((ref) => false);

final isSignUpLoading = StateProvider<bool>((ref) => false);

final registerErrorProvider = StateProvider<FormErrorModel>((ref) => FormErrorModel());

final passwordValidatorProvider = StateNotifierProvider<PasswordValidatorNotifier, bool>((ref) {
  return PasswordValidatorNotifier(false);
});

class SignUpClass {
  final AuthRepo authRepo;
  final GoRouter goRouter;
  final StateController<FormErrorModel> registerError;
  final StateController<bool> isLoading;
  final PasswordValidatorNotifier passwordValidator;

  SignUpClass({required this.authRepo, required this.goRouter, required this.isLoading, required this.registerError, required this.passwordValidator});

  Future register(String name, String email, String password, GlobalKey<FormState> formKey) async {
    isLoading.update((state) => true);
    registerError.update((state) => FormErrorModel());
    if (formKey.currentState != null && formKey.currentState!.validate()) {
      if (!passwordValidator.isValid()) {
        isLoading.update((state) => false);
        registerError.update(
              (state) => FormErrorModel(
            isError: true,
            errorMessage: 'invalidPassword',
          ),
        );
        return;
      }
      authRepo.createUser(name, email, password).then((isSuccess) {
        if (isSuccess) {
          isLoading.update((state) => false);
          goRouter.goNamed(RouteNames.confirmEmailRouteName);
        } else {
          registerError.update(
                (state) => FormErrorModel(
                isError: true,
                errorMessage: "error"
            ),
          );
          formKey.currentState != null && formKey.currentState!.validate();
          isLoading.update((state) => false);
        }
      });
    } else {
      isLoading.update((state) => false);
      registerError.update(
            (state) => FormErrorModel(
          isError: true,
          errorMessage: "verifyFields",
        ),
      );
    }
  }
}

final signUpProvider = Provider(
  (ref) => SignUpClass(
    authRepo: ref.watch(authRepo),
    goRouter: ref.watch(routerProvider),
    isLoading: ref.watch(isSignUpLoading.notifier),
    registerError: ref.watch(registerErrorProvider.notifier),
    passwordValidator: ref.watch(passwordValidatorProvider.notifier)
  ),
);
