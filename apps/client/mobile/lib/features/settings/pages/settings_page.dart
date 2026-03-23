import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/privacy/privacy.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/profile/profile.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/withdrawal_settings/withdrawal_settings.dart';
import 'package:tagpeak/shared/widgets/slide_tabs/slide_tabs.dart';

class SettingsPage extends ConsumerWidget {
  final String? initialTab;
  const SettingsPage({super.key, this.initialTab});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    List<Tab> tabGroup = [
      Tab(key: const ValueKey('profile'), text: translation(context).profile),
      Tab(key: const ValueKey('withdrawal'), text: translation(context).withdrawalSettings),
      Tab(key: const ValueKey('privacy'), text: translation(context).privacy)
    ];

    final int initialIndex = tabGroup.indexWhere((element) => (element.key as ValueKey<String>).value == initialTab);
    final int safeInitialIndex = initialIndex == -1 ? 0 : initialIndex;

    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 100),
          Text(
            translation(context).settings,
            style: Theme.of(context).textTheme.headlineLarge,
          ),
          const SizedBox(height: 5),
          SlideTabs(
            initialIndex: safeInitialIndex,
            tabs: tabGroup,
            views: const [
              Profile(),
              WithdrawalSettings(),
              Privacy()
            ],
          ),
        ],
      ),
    );
  }
}
