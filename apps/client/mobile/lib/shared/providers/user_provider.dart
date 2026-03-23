import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/auth/screens/login/provider/login_provider.dart';
import 'package:tagpeak/shared/models/user_stats_model/user_stats_model.dart';
import 'package:tagpeak/shared/models/user_update_model/user_update_model.dart';
import 'package:tagpeak/shared/services/user_service.dart';
import 'package:tagpeak/utils/logging.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';
import '../models/email_verified_model/email_verified_model.dart';
import 'package:tagpeak/shared/models/user_model/user_model.dart' as user_data_model;

class UserProvider {
  final UserService userService;
  final StateController<FormErrorModel> resetPasswordErrorProvider;

  UserProvider({required this.userService, required this.resetPasswordErrorProvider});

  Future<void> setEmailVerified(String userUuid, bool isVerified, BuildContext context) async {
    try {
      final EmailVerifiedModel body = EmailVerifiedModel(
          userId: userUuid,
          isVerified: isVerified.toString()
      );

      await userService.setEmailVerified(body);
      await updateUser(userUuid, const UserUpdateModel(onboardingFinished: 'true'));

      if (context.mounted) {
        Navigator.pop(context);
      }
    } catch (e) {
      throw Exception('Error: setEmailVerified - Something went wrong $e');
    }
  }

  Future<user_data_model.UserModel> updateUser(String userUuid, UserUpdateModel body) async {
    try {
      return await userService.updateUser(userUuid, body);
    } catch (e) {
      throw Exception('Error: updateUser - Something went wrong $e');
    }
  }

  Future<user_data_model.UserModel> getUser() async {
    try {
      return await userService.getUser();
    } catch (e) {
      throw Exception('Error: getUser - Something went wrong $e');
    }
  }

  Future<bool> resetPassword(String email) async {
    try {
      final response = await userService.resetPassword(email);
      if (response.statusCode == 200) {
        return true;
      } else if (response.statusCode == 404) {
        resetPasswordErrorProvider.update((state) => FormErrorModel(
            isError: true,
            errorMessage: 'emailNotFound'
        ));
        return false;
      } else {
        resetPasswordErrorProvider.update((state) => FormErrorModel(
            isError: true,
            errorMessage: 'defaultError'
        ));
        return false;
      }
    } catch (e) {
      Log.warning("Error: $e");
      rethrow;
    }
  }

  Future<UserStatsModel> getUserStats(String id) async {
    try {
      return await userService.getUserStats(id);
    } catch (e) {
      throw Exception('Error: getUserStats - Something went wrong $e');
    }
  }
}

final userNotifier = Provider(
    (ref) => UserProvider(
        userService: ref.watch(userService),
        resetPasswordErrorProvider: ref.watch(resetPasswordErrorProvider.notifier)
    )
);
