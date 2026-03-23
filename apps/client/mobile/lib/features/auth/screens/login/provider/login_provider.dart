import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/auth/repositories/auth_repo.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/features/core/routes/route_notifier.dart';
import 'package:tagpeak/shared/providers/config_provider.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/shared/providers/user_provider.dart' as users_provider;

enum LoginError {
  wrongCredentials,
  incompleteSetup
}

final showPasswordProvider = StateProvider<bool>((ref) => true);

final isSigningInProvider = StateProvider<bool>((ref) => false);

final loginErrorProvider = StateProvider<FormErrorModel>((ref) => FormErrorModel());

final savePasswordProvider = StateProvider((ref) => true);

final loginLoading = StateProvider<bool>((ref) => false);

final resetPasswordErrorProvider = StateProvider<FormErrorModel>((ref) => FormErrorModel());

class LoginClass {
  final AuthRepo authRepo;
  final GoRouter goRouter;
  final UserClass userClass;
  final StateController<FormErrorModel> loginError;
  final StateController<bool> isLoading;
  final StateController<bool> savePassword;
  final StateController<FormErrorModel> resetPasswordError;
  final users_provider.UserProvider userProvider;

  final ConfigProvider configProvider;

  LoginClass({required this.authRepo, required this.goRouter, required this.isLoading, required this.loginError, required this.savePassword, required this.userClass, required this.resetPasswordError, required this.userProvider, required this.configProvider});

  Future<bool> login(String username, String password, GlobalKey<FormState> formKey) async {
    isLoading.update((state) => true);
    loginError.update((state) => FormErrorModel());
    if (formKey.currentState != null && formKey.currentState!.validate()) {
      return authRepo.login(username, password).then((errorDescription) async {
        if (errorDescription == null) {
          /// Load basic user data from JWT (local, instant)
          userClass.saveUserData();
          /// Sync and update user data from backend (latest info)
          await userProvider.getUser().then((user) {
            List<String> roles = user.groups
                .where((element) => element.contains('user_type'))
                .map((element) => element.split('/')[2])
                .toList();

            userClass.userNotifier.update((state) => state = state?.copyWith(
                uuid: user.uuid,
                profilePicture: user.profilePicture,
                referralCode: user.referralCode,
                currency: user.currency,
                preferred_username: user.displayName,
                newsletter: user.newsletter,
                birthDate: user.birthDate,
                country: user.country,
                roles: roles
            ));
          });
          // Load configs
          configProvider.loadConfigurations();
          isLoading.update((state) => false);
          return true;
        } else {
          loginError.update(
            (state) => FormErrorModel(
              isError: true,
              errorMessage: errorDescription == 'Invalid user credentials' ?
                LoginError.wrongCredentials.name : LoginError.incompleteSetup.name
            ),
          );
          formKey.currentState != null && formKey.currentState!.validate();
          isLoading.update((state) => false);
          return false;
        }
      });
    } else {
      isLoading.update((state) => false);
      loginError.update(
        (state) => FormErrorModel(
          isError: true,
          errorMessage: "Wrong email or password.",
        ),
      );
    }
    return false;
  }

  Future signInWithSocial(String provider) async {
    isLoading.update((state) => true);
    loginError.update((state) => FormErrorModel());

    authRepo.signInWithSocial(provider).then((value) async {
      if (value) {
        /// Load basic user data from JWT (local, instant)
        userClass.saveUserData();
        /// Sync and update user data from backend (latest info)
        await userProvider.getUser().then((user) {
          List<String> roles = user.groups
              .where((element) => element.contains('user_type'))
              .map((element) => element.split('/')[2])
              .toList();

          userClass.userNotifier.update((state) => state = state?.copyWith(
              uuid: user.uuid,
              profilePicture: user.profilePicture,
              referralCode: user.referralCode,
              currency: user.currency,
              preferred_username: user.displayName,
              newsletter: user.newsletter,
              birthDate: user.birthDate,
              country: user.country,
              roles: roles
          ));
        });
        // Load configs
        configProvider.loadConfigurations();
        isLoading.update((state) => false);
        goRouter.goNamed(RouteNames.clientDashboardRouteName);
      } else {
        isLoading.update((state) => false);
        loginError.update((state) => FormErrorModel(
              isError: true,
              errorMessage: "Social login failed.",
            ));
      }
    }).catchError((error) {
      isLoading.update((state) => false);
      loginError.update((state) => FormErrorModel(
            isError: true,
            errorMessage: "Social login failed: ${error.toString()}",
          ));
    });
  }

  Future logout() async {
    goRouter.goNamed(RouteNames.loginRouteName);
    Future.delayed(const Duration(milliseconds: 100), () {
      authRepo.logout();
      userClass.removeUser();
    });
  }
}

final loginNotifier = Provider(
  (ref) => LoginClass(
    authRepo: ref.watch(authRepo),
    goRouter: ref.watch(routerProvider),
    isLoading: ref.watch(isSigningInProvider.notifier),
    loginError: ref.watch(loginErrorProvider.notifier),
    userClass: ref.read(userClassProvider),
    savePassword: ref.watch(savePasswordProvider.notifier),
    resetPasswordError: ref.watch(resetPasswordErrorProvider.notifier),
    userProvider: ref.watch(users_provider.userNotifier),
    configProvider: ref.watch(configClassProvider)
  ),
);
