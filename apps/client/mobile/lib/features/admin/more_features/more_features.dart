import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class MoreFeatures extends StatelessWidget {
  const MoreFeatures({super.key});

  void handleNavigation(BuildContext context, String path) {
    context.go(path);
  }

  @override
  Widget build(BuildContext context) {
    // ToDO change path for each entry
    final List<Map<String, String>> entries = [
      {'name': translation(context).storeVisits, 'path': '/admin/stores', 'icon': Assets.eye},
      {'name': translation(context).countries, 'path': '/admin', 'icon': Assets.flag},
      {'name': translation(context).sources, 'path': '/admin', 'icon': Assets.heardHandshake},
      {'name': translation(context).categories, 'path': '/admin', 'icon': Assets.category}
    ];

    return Container(
      color: Colors.white,
      child: Padding(
        padding: const EdgeInsets.only(top: 80, left: 40, right: 40),
        child: Column(
          spacing: 20,
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            ...entries.map((entry) {
              return GestureDetector(
                onTap: () => handleNavigation(context, entry['path']!),
                child: Container(
                  padding: const EdgeInsets.only(top: 5, bottom: 10),
                  decoration: const BoxDecoration(
                    border: Border(
                      bottom: BorderSide(color: AppColors.wildSand),
                    )
                  ),
                  child: Row(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    spacing: 20,
                    children: [
                      SvgPicture.asset(entry['icon']!, width: 22, colorFilter: const ColorFilter.mode(AppColors.waterloo, BlendMode.srcIn)),
                      Text(entry['name']!, style: TextTheme.of(context).titleMedium)
                    ],
                  ),
                ),
              );
            })
          ],
        ),
      ),
    );
  }
}
