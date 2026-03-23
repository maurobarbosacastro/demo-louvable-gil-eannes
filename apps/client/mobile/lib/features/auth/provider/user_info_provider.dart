import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/hive_cache.dart';
import 'package:jwt_decode/jwt_decode.dart';
import 'package:tagpeak/utils/enums/user_roles_enum.dart';

class UserClass {
  final HiveDatabase storage;
  final StateController<UserModel?> userNotifier;

  UserClass({required this.storage, required this.userNotifier});

  saveUserData() {
    if (storage.token != null) {
      final jwtPayload = Jwt.parseJwt(storage.token!);

      userNotifier.update((state) => state = UserModel(
        uuid: jwtPayload['sub'] as String, // temporary user uuid
        email: jwtPayload['email'] as String,
        given_name: jwtPayload['given_name'] as String,
        family_name: jwtPayload['family_name'] as String,
        preferred_username: jwtPayload['preferred_username'] as String,
        roles: [UserRolesEnum.client.name]
      ));
    }
  }

  removeUser() {
    userNotifier.update((state) => null);
  }
}

final userProvider = StateProvider<UserModel?>((ref) => null);

final userClassProvider = Provider(
  (ref) => UserClass(
    storage: ref.watch(hiveProvider),
    userNotifier: ref.watch(userProvider.notifier),
  ),
);
