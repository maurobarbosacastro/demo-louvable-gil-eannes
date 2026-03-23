import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/auth/screens/login/provider/login_provider.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/profile/profile_form.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/profile/user_profile.dart';
import 'package:tagpeak/shared/models/user_update_model/user_update_model.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/toggle.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:permission_handler/permission_handler.dart';

class Profile extends ConsumerStatefulWidget {
  const Profile({super.key});

  @override
  ConsumerState<Profile> createState() => _ProfileState();
}

class _ProfileState extends ConsumerState<Profile> {
  UserModel? user;
  bool newsletterSubscription = false;
  PermissionStatus notificationPermission = PermissionStatus.denied;

  // Profile form properties
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();
  late TextEditingController nameController;
  late TextEditingController displayNameController;
  late TextEditingController passwordController;
  late TextEditingController emailController;
  late TextEditingController dateOfBirthController;
  late TextEditingController countryController;
  late Map<String, ProfileFormProps> profileFormConfigs;

  @override
  void initState() {
    super.initState();

    user = ref.read(userProvider);
    _initFormConfigs();
    newsletterSubscription = user?.newsletter ?? false;
    _checkNotificationPermission();
  }

    // Init profile from props
    nameController = TextEditingController(text: '${user?.given_name} ${user?.family_name}');
    displayNameController = TextEditingController(text: user?.preferred_username);
    passwordController = TextEditingController(text: '***********');
    emailController = TextEditingController(text: user?.email);
    dateOfBirthController = TextEditingController(text: user?.birthDate);
    countryController = TextEditingController(text: user?.country);

    profileFormConfigs = {
      'name': ProfileFormProps(
          readOnly: true,
          controller: nameController,
      ),
      'displayName': ProfileFormProps(
        readOnly: true,
        controller: displayNameController,
      ),
      'password': ProfileFormProps(
          readOnly: true,
          controller: passwordController
      ),
      'email': ProfileFormProps(
          readOnly: true,
          controller: emailController
      ),
      'dateOfBirth': ProfileFormProps(
          readOnly: true,
          controller: dateOfBirthController,
      ),
      'country': ProfileFormProps(
          readOnly: true,
          controller: countryController,
      )
    };

    setState(() => newsletterSubscription = user?.newsletter ?? false);
  }

  void handleSuffixAction(String field, SuffixAction suffixAction) {
    if (suffixAction == SuffixAction.edit) {
      setState(() => profileFormConfigs[field]?.readOnly = false);
    } else if (suffixAction == SuffixAction.close) {
      setState(() => profileFormConfigs[field]?.readOnly = true);
    } else {
      updateUser(field, profileFormConfigs[field]?.controller.text);
    }
  }

  Future<void> updateUser(String field, dynamic value) async {
    UserUpdateModel body = const UserUpdateModel();

    if (field == 'password') return;

    if (!formKey.currentState!.validate()) {
      return;
    }

    switch (field) {
      case 'name':
        String firstName = value.split(' ').firstWhere((_) => true,
            orElse: () => '');
        List<String> lastNameParts = value.split(' ').skip(1).toList();

        body = UserUpdateModel(
            firstName: firstName,
            lastName: lastNameParts.join(' ')
        );
        break;
      case 'dateOfBirth':
        body = UserUpdateModel(
            birthDate: value
        );
        break;
      default:
        body = UserUpdateModel.fromJson({ field: value});
        break;
    }

    await ref.read(userNotifier).updateUser(user!.uuid, body).then((response) {
      ref.read(userClassProvider).userNotifier.update((state) {
        handleSuffixAction(field, SuffixAction.close);

        return state?.copyWith(
            preferred_username: response.displayName,
            given_name: response.firstName,
            family_name: response.lastName,
            referralCode: response.referralCode,
            currency: response.currency,
            newsletter: response.newsletter,
            birthDate: response.birthDate,
            country: response.country);
      });
    });
  }

  Future<void> _checkNotificationPermission() async {
    final status = await Permission.notification.status;
    if (mounted) {
      setState(() => notificationPermission = status);
    }
  }

  Future<void> _openNotificationSettings() async {
    await openAppSettings();
    // Recheck permission when user returns
    _checkNotificationPermission();
  }

  @override
  Widget build(BuildContext context) {
    return ScrollableBody(
        child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        spacing: 30,
        children: [
        // Profile picture, name & email
          const Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              UserProfileWidget(),
            ],
          ),
          ProfileForm(
            formKey: formKey,
            configs: profileFormConfigs,
            onSuffixAction: (field, suffixAction) => handleSuffixAction(field, suffixAction),
          ),
          const SizedBox(height: 30),
          const Divider(color: AppColors.wildSand),
          _buildSubscriptionSection(context),
          const SizedBox(height: 40),
          const Divider(color: AppColors.wildSand),
          _buildNotificationSection(context),
          const SizedBox(height: 60),
          Button(
            label: translation(context).logout,
            mode: Mode.outlinedDark,
            icon: const Icon(Icons.logout),
            onPressed: () => ref.read(loginNotifier).logout(),
          ),
        ],
      ),
    );
  }

  Widget _buildSubscriptionSection(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const SizedBox(height: 20),
        Text(translation(context).subscriptionDetail,
            style: const TextStyle(fontSize: 16, color: AppColors.licorice)),
        const SizedBox(height: 10),
        Toggle(
          label: translation(context).newsletterSubscription,
          value: newsletterSubscription,
          onChanged: (v) {
            setState(() => newsletterSubscription = v);
            _updateUser('newsletter', v);
          },
        ),
      ],
    ));
  }

  Widget _buildNotificationSection(BuildContext context) {
    String statusText = notificationPermission.isGranted ? 'Enabled' : 'Disabled';

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const SizedBox(height: 20),
        const Text(
          'Push Notifications',
          style: TextStyle(fontSize: 16, color: AppColors.licorice),
        ),
        const SizedBox(height: 10),
        Text(
          'Manage notification preferences ($statusText)',
          style: const TextStyle(fontSize: 14, color: AppColors.grey),
        ),
        const SizedBox(height: 10),
        Button(
          label: 'Open Settings',
          mode: Mode.outlinedDark,
          icon: const Icon(Icons.settings),
          onPressed: _openNotificationSettings,
        ),
      ],
    );
  }
}

