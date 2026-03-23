import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class UserProfileWidget extends ConsumerWidget {
  const UserProfileWidget({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(userProvider);

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        ClipRRect(
          borderRadius: BorderRadius.circular(100),
          child: FadeInImage.assetNetwork(
            fadeInDuration: const Duration(milliseconds: 1),
            placeholder: Assets.userDefaultImage,
            image: user?.profilePicture ?? '',
            fit: BoxFit.cover,
            width: 80,
            height: 80,
            imageErrorBuilder: (context, error, stackTrace) =>
                Image.asset(Assets.userDefaultImage,
                    fit: BoxFit.cover, width: 80, height: 80),
          ),
        ),
        const SizedBox(height: 10),
        Text('${user?.given_name} ${user?.family_name}',
            style: Theme.of(context).textTheme.headlineSmall),
        Text(user?.email ?? '', style: Theme.of(context).textTheme.titleSmall),
      ],
    );
  }
}
