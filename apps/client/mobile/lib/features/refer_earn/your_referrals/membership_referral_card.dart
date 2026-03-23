import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/shared/providers/config_provider.dart';
import 'package:tagpeak/shared/widgets/cards/information_card.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';
import 'package:tagpeak/utils/extensions/strings_extension.dart';
import 'package:tagpeak/utils/multi_styled_text.dart';

enum Membership {
  base,
  silver,
  gold
}
class MembershipReferralCard extends ConsumerWidget {
  final String membership;
  final double spent;
  final String currency;
  final double extra;
  const MembershipReferralCard({super.key, required this.membership, required this.spent, required this.currency, required this.extra});

  /// Get next (or last) membership
  String get _membership => _getNextMembership().toCapitalize();

  String _getNextMembership() {
    return switch (membership) {
      'base' => Membership.silver.name,
      'silver' => Membership.gold.name,
      _ => Membership.gold.name,
    };
  }

  Map<String, dynamic> _setupMembershipLevels(String level, WidgetRef ref) {
    final configs = ref.read(configsProvider) as Map<String, dynamic>?;

    if (level == Membership.base.name) {
      return {
        'percentageOnReward': configs?['referral_member_friend_cash_reward']?['value'],
        'maxValue': configs?['referral_silver_status_goal_amount']?['value'] ?? '0',
      };
    } else if (level == Membership.silver.name) {
      return {
        'percentageOnTransaction': configs?['referral_silver_transaction_cash_reward']?['value'],
        'percentageOnReward': configs?['referral_gold_status_goal']?['value'],
        'maxValue': configs?['referral_gold_status_goal_amount']?['value'] ?? '0',
      };
    } else {
      return {
        'percentageOnTransaction': configs?['referral_gold_transaction_cash_reward']?['value'],
        'percentageOnReward': configs?['referral_gold_friend_reward_share']?['value'],
        'maxValue': configs?['referral_gold_status_goal_amount']?['value'] ?? '0',
      };
    }
  }

  Map<String, String?> _getMembershipLabels(BuildContext context, double percentageOnReward, double nextPercentageOnReward) {
    if (membership != Membership.gold.name) {
      return {
        'title': translation(context).reachedNextLevel(_membership),
        'description': translation(context).lastMembershipLevel
      };
    }
    return {
      'title': translation(context).reachStatusReferral(_membership, extra.toCurrency(currency)),
      'description': translation(context).reachStatusReferralDescription((extra * 2).toCurrency(currency), _membership, percentageOnReward, nextPercentageOnReward),
      'question': translation(context).howToReachNextLevel
    };
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final String maxValue = _setupMembershipLevels(membership, ref)['maxValue'];

    final double percentageOnReward = double.parse(_setupMembershipLevels(membership, ref)['percentageOnReward'].toString().replaceAll('%', ''));
    final double nextPercentageOnReward = double.parse(_setupMembershipLevels(_getNextMembership(), ref)['percentageOnReward'].toString().replaceAll('%', ''));

    final double goal = double.parse(maxValue.replaceAll(RegExp(r'[^\d.]'), ''));

    final double percentageSpent = spent / goal;

    final Map<String, String?> labels = _getMembershipLabels(context, percentageOnReward, nextPercentageOnReward);

    return InformationCard(
        backgroundColor: AppColors.youngBlue,
        color: Colors.white,
        title: MultiStyledText(
          fullText: labels['title']!,
          defaultStyle: Theme.of(context).textTheme.headlineSmall?.copyWith(color: Colors.white),
          styledPlaceholders: {
            _membership: const TextStyle(color: AppColors.grandis)
          },
        ),
        description: labels['description']!,
        question: labels['question'],
        handleQuestionClick: () => context.pushNamed(RouteNames.referEarnRouteName),
        content: Column(
          spacing: 5,
          children: [
            Container(
              margin: const EdgeInsets.only(top: 15),
              height: 8,
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  colors: const [AppColors.grandis, AppColors.grandis, Colors.white, Colors.white],
                  stops: [0.0, percentageSpent, percentageSpent, 1.0]
                ),
                borderRadius: BorderRadius.circular(30)
              ),
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text('${spent.toCurrency(currency)} ${translation(context).spent}',
                    style: Theme.of(context).textTheme.titleSmall?.copyWith(color: Colors.white)),
                Text('${goal.toCurrency(currency)} ${translation(context).goal.toCapitalize()}',
                    style: Theme.of(context).textTheme.titleSmall?.copyWith(color: Colors.white))
              ],
            )
          ],
        )
    );
  }
}
