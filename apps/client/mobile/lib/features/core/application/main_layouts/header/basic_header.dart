import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/shared/models/user_stats_model/user_stats_model.dart';
import 'package:tagpeak/shared/widgets/header.dart';
import 'package:tagpeak/shared/widgets/membership_pill/membership_pill.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class BasicHeader extends StatelessWidget implements PreferredSizeWidget {
  final String? profilePicture;
  final bool isAdmin;
  final UserStatsModel? stats;

  const BasicHeader({super.key, this.profilePicture = '', this.isAdmin = false, this.stats});

  @override
  Widget build(BuildContext context) {
    final level = stats?.level ?? '/membership_levels/base';
    final levelSegment = level.split('/')[2];

    return Header(
        endContent: Row(
      children: [
        if (!isAdmin) MembershipPill(membership: levelSegment),
        const SizedBox(width: 8),
        GestureDetector(
          onTap: () => context.goNamed(RouteNames.settingsRouteName),
          child: ClipRRect(
            borderRadius: BorderRadius.circular(100),
            child: FadeInImage.assetNetwork(
              fadeInDuration: const Duration(milliseconds: 1),
              placeholder: Assets.userDefaultImage,
              image: profilePicture ?? '',
              fit: BoxFit.cover,
              width: 32,
              height: 32,
              imageErrorBuilder: (context, error, stackTrace) => Image.asset(
                Assets.userDefaultImage,
                fit: BoxFit.cover,
                width: 32,
                height: 32,
              ),
            ),
          ),
        ),
      ],
    ));
  }

  @override
  Size get preferredSize => const Size.fromHeight(70.0);
}
