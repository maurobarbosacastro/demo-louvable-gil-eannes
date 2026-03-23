import 'package:carousel_slider/carousel_controller.dart';
import 'package:flutter/gestures.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/screens/register/providers/register_providers.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';
import '../../../../../shared/widgets/carousel_slider.dart';
import '../../../../../utils/constants/colors.dart';
import '../../login/provider/login_provider.dart';
import '../../login/widgets/password_form_field_widget.dart';
import '../../login/widgets/text_form_field_widget.dart';

class RegisterScreen extends ConsumerWidget {
  const RegisterScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final CarouselSliderController carouselSliderController = CarouselSliderController();

    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Carousel(
              carouselController: carouselSliderController,
              carouselItems: [
                ClipRRect(
                  borderRadius: BorderRadius.circular(10.0),
                  child: Image.asset(
                    Assets.dataTransferGraphic,
                    fit: BoxFit.fitWidth,
                    width: double.infinity,
                  ),
                ),
                ClipRRect(
                  borderRadius: BorderRadius.circular(10.0),
                  child: Image.asset(
                    Assets.dataTransferGraphic,
                    fit: BoxFit.fitWidth,
                    width: double.infinity,
                  ),
                ),
              ],
            ),
            const SizedBox(height: 20.0),
            Padding(
              padding:
                  const EdgeInsets.symmetric(horizontal: 11.0, vertical: 10.0),
              child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      translation(context).joinTagpeak,
                      style: Theme.of(context).textTheme.headlineMedium,
                    ),
                    RichText(
                      text: TextSpan(children: [
                        TextSpan(
                          text: translation(context).haveAnAccount,
                          style: Theme.of(context).textTheme.bodyMedium,
                        ),
                        TextSpan(
                          text: translation(context).logIn,
                          style:
                              Theme.of(context).textTheme.bodyMedium?.copyWith(
                                    color: AppColors.youngBlue,
                                    decoration: TextDecoration.underline,
                                    decorationColor: AppColors.youngBlue,
                                  ),
                          recognizer: TapGestureRecognizer()
                            ..onTap = () {
                              ref.watch(registerErrorProvider.notifier).update((state) => FormErrorModel(isError: false));
                              context.go('/login');
                            },
                        ),
                      ]),
                    ),
                    const SizedBox(height: 20.0),
                    const _FormBuild(),
                  ]),
            )
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
  final TextEditingController _nameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();

  final TextEditingController _passwordController = TextEditingController();

  //Key to validate the form
  final _formKey = GlobalKey<FormState>();

  String? getCreateAccountError(String? message) {
    switch(message) {
      case 'invalidPassword':
          return translation(context).invalidPassword;
      case 'error':
          return translation(context).invalidPassword;
      case 'verifyFields':
        return translation(context).verifyFields;
      default:
        return message;
    }
  }

  void resetSignUpErrorState() {
    ref.watch(registerErrorProvider.notifier).update((state) => FormErrorModel(isError: false));
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          Column(
            children: [
              TextFormFieldWidget(
                isError: ref.watch(registerErrorProvider).isError,
                controller: _nameController,
                validator: (value) => value!.isEmpty ? '' : null,
                placeholder: translation(context).yourName,
                errorStyle: const TextStyle(height: 0, fontSize: 0),
                label: translation(context).name,
              ),
              TextFormFieldWidget(
                isError: ref.watch(registerErrorProvider).isError,
                controller: _emailController,
                validator: (value) => value!.isEmpty ? '' : null,
                placeholder: translation(context).yourEmail,
                errorStyle: const TextStyle(height: 0, fontSize: 0),
                label: translation(context).email,
              ),
              PasswordFormFieldWidget(
                isError: getCreateAccountError(ref.watch(registerErrorProvider).errorMessage),
                controller: _passwordController,
                validator: (value) => value!.isEmpty ? '' : null,
                placeholder: translation(context).yourPassword,
                obscureText: ref.watch(showPasswordProvider),
                showPassword: () => ref.read(showPasswordProvider.notifier).update((state) => !state),
                label: translation(context).password
              ),
            ],
          ),
          const SizedBox(height: 10.0),
          Button(
            onPressed: () {
              ref.read(passwordValidatorProvider.notifier).validate(_passwordController.text);

              ref.read(signUpProvider).register(
                _nameController.text,
                _emailController.text,
                _passwordController.text,
                _formKey,
              );

              resetSignUpErrorState();
            },
            label: translation(context).createAccount,
          ),
          const SizedBox(height: 10.0),
          /* Column(
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
                  Text(translation(context).orSignUpWith, style: Theme.of(context).textTheme.bodySmall),
                  Expanded(
                    child: Container(
                      height: 1,
                      color: AppColors.waterloo,
                    ),
                  ),
                ],
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                spacing: 20,
                children: [
                  Expanded(
                    child: Button(
                      onPressed: () {
                        print('Click on Facebook!');
                        resetSignUpErrorState();
                      },
                      label: 'Facebook',
                      mode: Mode.blue,
                      icon: SvgPicture.asset('lib/assets/icons/facebook.svg'),
                    ),
                  ),
                  Expanded(
                      child: Button(
                        onPressed: () {
                          print('Click on Google!');
                          resetSignUpErrorState();
                        },
                        label: 'Google',
                        mode: Mode.outlinedDark,
                        icon: SvgPicture.asset('lib/assets/icons/google.svg'),
                      )
                  )
                ],
              )
            ],
          ) */
        ],
      ),
    );
  }
}
