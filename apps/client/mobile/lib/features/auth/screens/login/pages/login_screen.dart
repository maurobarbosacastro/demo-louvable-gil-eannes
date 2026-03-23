import 'package:flutter/gestures.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/auth/screens/login/widgets/social_media_widget.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/enums/user_roles_enum.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';
import '../../../../../shared/widgets/button.dart';
import '../provider/login_provider.dart';
import '../widgets/checkbox_widget.dart';
import '../widgets/password_form_field_widget.dart';
import '../widgets/text_form_field_widget.dart';

class LoginScreen extends ConsumerWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Padding(
      padding: const EdgeInsetsDirectional.symmetric(
          horizontal: 20.0, vertical: 35.0),
      child: SingleChildScrollView(
        child: Column(
          spacing: 2.5,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              translation(context).welcomeBack,
              style: Theme.of(context).textTheme.headlineMedium,
            ),
            RichText(
              text: TextSpan(children: [
                TextSpan(
                  text: translation(context).needAnAccount,
                  style: Theme.of(context).textTheme.bodyMedium,
                ),
                TextSpan(
                  text: translation(context).signUpHere,
                  style: Theme.of(context).textTheme.bodyMedium?.copyWith(
                        color: AppColors.youngBlue,
                        decoration: TextDecoration.underline,
                        decorationColor: AppColors.youngBlue,
                      ),
                  recognizer: TapGestureRecognizer()
                    ..onTap = () {
                      ref.watch(loginErrorProvider.notifier).update((state) => FormErrorModel(isError: false));
                      context.go('/register');
                    },
                ),
              ]),
            ),
            const SizedBox(height: 20.0),
            const _FormBuild(),
          ],
        ),
      ),
    );
  }
}

class _FormBuild extends ConsumerStatefulWidget {
  const _FormBuild();

  @override
  ConsumerState<_FormBuild> createState() => _FormBuildState();
}

class _FormBuildState extends ConsumerState<_FormBuild> {
  //Controllers to save the user input
  final TextEditingController _emailController = TextEditingController();

  final TextEditingController _passwordController = TextEditingController();

  //Key to validate the form
  final _formKey = GlobalKey<FormState>();

  final _storage = const FlutterSecureStorage();

  String getAuthError() {
    return ref.watch(loginErrorProvider).errorMessage == 'wrongCredentials'
        ? translation(context).wrongCredentials
        : translation(context).incompleteSetup;
  }

  void resetLogInErrorState() {
    ref.watch(loginErrorProvider.notifier).update((state) => FormErrorModel(isError: false));
  }

  _saveCredentials() async {
    await _storage.write(key: 'KEY_EMAIL', value: _emailController.text);
    await _storage.write(key: 'KEY_PASSWORD', value: _passwordController.text);
  }

  _readCredentials() async {
    _emailController.text = await _storage.read(key: 'KEY_EMAIL') ?? '';
    _passwordController.text = await _storage.read(key: 'KEY_PASSWORD') ?? '';
  }

  @override
  void initState() {
    super.initState();

    _readCredentials();

    // Clear any previous login errors when entering the screen
    WidgetsBinding.instance.addPostFrameCallback((_) {
      resetLogInErrorState();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: [
          Column(
            spacing: 0.0,
            children: [
              TextFormFieldWidget(
                isError: ref.watch(loginErrorProvider).isError,
                controller: _emailController,
                validator: (value) => value!.isEmpty ? '' : null,
                placeholder: translation(context).yourEmail,
                errorStyle: const TextStyle(height: 0, fontSize: 0),
                label: translation(context).email,
              ),
              PasswordFormFieldWidget(
                isError: ref.watch(loginErrorProvider).isError
                    ? getAuthError()
                    : null,
                controller: _passwordController,
                validator: (value) =>
                    value!.isEmpty ? translation(context).errorLogin : null,
                placeholder: translation(context).yourPassword,
                obscureText: ref.watch(showPasswordProvider),
                showPassword: () => ref
                    .read(showPasswordProvider.notifier)
                    .update((state) => !state),
                label: translation(context).password,
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  CheckboxWidget(
                    key: const Key("SavePasswordMobile"),
                    value: ref.watch(savePasswordProvider),
                    onChanged: (value) => ref
                        .read(savePasswordProvider.notifier)
                        .update((state) => value ?? false),
                    label: translation(context).rememberMe,
                  ),
                  TextButton(
                    onPressed: () {
                      resetLogInErrorState();
                      context.go('/reset-password');
                    },
                    style: TextButton.styleFrom(
                        padding: EdgeInsets.zero,
                        textStyle: Theme.of(context)
                            .textTheme
                            .titleSmall
                            ?.copyWith(color: AppColors.licorice)),
                    child: Text(
                      translation(context).forgotPassword,
                    ),
                  ),
                ],
              ),
            ],
          ),
          const SizedBox(height: 20.0),
          Button(
            onPressed: () {
              ref.watch(loginNotifier).login(
                  _emailController.text, _passwordController.text, _formKey).then((success) {
                    if (!context.mounted) return;
                    if(success) {
                      UserModel? loggedUser = ref.read(userProvider);
                      if (loggedUser != null) {
                        loggedUser.roles.contains(UserRolesEnum.admin.name)
                            ? context.go(RouteNames.adminRouteLocation)
                            : context.go(RouteNames.clientRouteLocation);
                      }
                    }

                    // Save credentials
                    if (ref.read(savePasswordProvider)) {
                      _saveCredentials();
                    }
              });
              resetLogInErrorState();
            },
            label: translation(context).signIn,
          ),
          const SizedBox(height: 22.0),
           Column(
            spacing: 10,
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                spacing: 5,
                children: [
                  Expanded(
                    child: Container(
                      height: 1,
                      color: AppColors.waterloo,
                    ),
                  ),
                  Text(translation(context).orSignInWith,
                      style: Theme.of(context).textTheme.bodySmall),
                  Expanded(
                    child: Container(
                      height: 1,
                      color: AppColors.waterloo,
                    ),
                  ),
                ],
              ),
              SocialMediaWidget()
            ],
          )
        ],
      ),
    );
  }
}
