import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/providers/config_provider.dart';
import 'package:tagpeak/shared/widgets/membership_pill/membership_pill.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class HowItWorksCard extends ConsumerWidget {
  const HowItWorksCard({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final configs = ref.read(configsProvider) as Map<String, dynamic>?;


    return Container(
        padding: const EdgeInsets.symmetric(vertical: 30, horizontal: 20),
        decoration: BoxDecoration(
            color: AppColors.wildSand,
            borderRadius: BorderRadius.circular(20)
        ),
        child: Column(
          spacing: 20,
          children: [
            SingleChildScrollView(
              scrollDirection: Axis.horizontal,

              child: Column(
                spacing: 20,
                children: [

                  const HowItWorksCardRow(
                    children: [
                      SizedBox.shrink(),
                      Align(
                        alignment: Alignment.centerLeft,
                        child: MembershipPill(membership: 'base')
                      ),
                      Align(
                          alignment: Alignment.centerLeft,
                          child: MembershipPill(membership: 'silver')
                      ),
                      Align(
                          alignment: Alignment.centerLeft,
                          child: MembershipPill(membership: 'gold')
                      )
                    ]
                  ),

                  HowItWorksCardRow(
                    children: [
                      Container(
                        padding: const EdgeInsets.only(bottom: 20),
                        decoration: const BoxDecoration(
                            border: Border(
                                bottom: BorderSide(color: AppColors.lavender)
                            )
                        ),
                        child: Column(
                          spacing: 5,
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(translation(context).status, style: TextTheme.of(context).titleSmall),
                            Text(translation(context).whatDoYouNeedToDo, style: TextTheme.of(context).titleSmall?.copyWith(color: AppColors.licorice))
                          ],
                        ),
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          decoration: const BoxDecoration(
                              border: Border(
                                  bottom: BorderSide(color: AppColors.radicalRed)
                              )
                          ),
                          child: Text(translation(context).signUp, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          decoration: const BoxDecoration(
                              border: Border(
                                  bottom: BorderSide(color: AppColors.radicalRed)
                              )
                          ),
                          child: Text(translation(context).oneReferralOrSpend(configs?['referral_silver_status_goal_amount']?['value']), style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          decoration: const BoxDecoration(
                              border: Border(
                                  bottom: BorderSide(color: AppColors.radicalRed)
                              )
                          ),
                          child: Text(translation(context).multipleReferralOrSpend(configs?['referral_gold_status_goal']?['value'], configs?['referral_gold_status_goal_amount']?['value']), style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
                      ),
                    ]
                  ),

                  HowItWorksCardRow(
                    children: [
                      Container(
                        padding: const EdgeInsets.only(bottom: 20),
                        decoration: const BoxDecoration(
                            border: Border(
                                bottom: BorderSide(color: AppColors.lavender)
                            )
                        ),
                        child: Column(
                          spacing: 5,
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(translation(context).yourBenefits, style: TextTheme.of(context).titleSmall),
                            Text(translation(context).moreCashRewardsFromPurchases, style: TextTheme.of(context).titleSmall?.copyWith(color: AppColors.licorice)),
                            const SizedBox(height: 5),
                            Text(translation(context).shareOfYourFriendsCashRewards, style: TextTheme.of(context).titleSmall?.copyWith(color: AppColors.licorice))
                          ],
                        ),
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          decoration: const BoxDecoration(
                              border: Border(
                                  bottom: BorderSide(color: AppColors.radicalRed)
                              )
                          ),
                          child: Column(
                            spacing: 10,
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text(translation(context).cashRewardsAsListedOnBrands, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                              Text(
                                  '${double.parse(configs?['referral_member_friend_cash_reward']?['value'].toString().replaceAll('%', '') ?? '0').round()}%',
                                  style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                            ],
                          )
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          decoration: const BoxDecoration(
                              border: Border(
                                  bottom: BorderSide(color: AppColors.radicalRed)
                              )
                          ),
                          child: Column(
                            spacing: 10,
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(translation(context).memberRewards, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                                  Text(
                                      '+ ${double.parse(configs?['referral_silver_transaction_cash_reward']?['value'].toString().replaceAll('%', '') ?? '0').round()}%',
                                      style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                                ],
                              ),
                              Text(
                                  '${double.parse(configs?['referral_silver_friend_reward_share']?['value'].toString().replaceAll('%', '') ?? '0').round()}%',
                                  style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                            ],
                          )
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          decoration: const BoxDecoration(
                              border: Border(
                                  bottom: BorderSide(color: AppColors.radicalRed)
                              )
                          ),
                          child: Column(
                            spacing: 10,
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(translation(context).memberRewards, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                                  Text(
                                      '+ ${double.parse(configs?['referral_gold_transaction_cash_reward']?['value'].toString().replaceAll('%', '') ?? '0').round()}%',
                                      style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                                ],
                              ),
                              Text(
                                  '${double.parse(configs?['referral_gold_friend_reward_share']?['value'].toString().replaceAll('%', '') ?? '0').round()}%',
                                  style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                            ],
                          )
                      ),
                    ],
                  ),

                  HowItWorksCardRow(
                    border: const Border(
                      bottom: BorderSide(color: AppColors.lavender)
                    ),
                    children: [
                      Container(
                        padding: const EdgeInsets.only(bottom: 20),
                        child: Column(
                          spacing: 5,
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(translation(context).yourFriendsBenefits, style: TextTheme.of(context).titleSmall),
                            Text(translation(context).statusTheyGetWhenRecommendedYou, style: TextTheme.of(context).titleSmall?.copyWith(color: AppColors.licorice))
                          ],
                        ),
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          child: Text(translation(context).silver, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          child: Text(translation(context).silver, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
                      ),
                      Container(
                          padding: const EdgeInsets.only(bottom: 28),
                          child: Text(translation(context).silver, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
                      ),
                    ],
                  ),

                ],
              ),
            ),

            Column(
              spacing: 5,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(translation(context).notes, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
                Text(translation(context).notesHowItWorks, style: TextTheme.of(context).bodySmall)
              ],
            )
          ],
        )
    );
  }
}

class HowItWorksCardRow extends StatelessWidget {
  final List<Widget> children;
  final Border? border;

  const HowItWorksCardRow({super.key, required this.children, this.border});

  @override
  Widget build(BuildContext context) {
    final double width = MediaQuery.of(context).size.width * 0.74;

    return Container(
      decoration: BoxDecoration(
          border: border
      ),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.end,
        spacing: 20,
        children: [
          ...children.map((child) {
            return SizedBox(
              width: width * 0.5,
              child: child,
            );
          })
        ],
      ),
    );
  }
}
