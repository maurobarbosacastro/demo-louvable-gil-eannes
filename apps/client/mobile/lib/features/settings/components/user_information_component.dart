import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../utils/constants/assets.dart';

class UserInformationComponent extends ConsumerWidget {
  const UserInformationComponent({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final userData = ref.watch(userProvider);
    return Column(
      children: [
        const SizedBox(height: 16),
        const Center(
          child: CircleAvatar(
            radius: 55,
            backgroundImage: AssetImage(Assets.userDefaultImage),
            backgroundColor: Colors.transparent,
          ),
        ),
        Padding(
          padding: const EdgeInsets.only(top: 16.0),
          child: Text(
             "${userData?.given_name} ${userData?.family_name}" ,
            textAlign: TextAlign.center,
            style: const TextStyle(fontSize: 24),
          ),
        ),
        const SizedBox(height: 32),
      ],
    );
  }
}
