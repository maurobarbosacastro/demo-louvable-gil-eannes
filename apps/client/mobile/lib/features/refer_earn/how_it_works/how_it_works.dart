import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/svg.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/refer_earn/how_it_works/how_it_works_card.dart';
import 'package:tagpeak/features/refer_earn/your_referrals/membership_referral_card.dart';
import 'package:tagpeak/shared/models/referral_info_model/revenue_info_model/revenue_info_model.dart';
import 'package:tagpeak/shared/models/user_stats_model/user_stats_model.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/shared/services/referrals_service.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/cards/information_card.dart';
import 'package:tagpeak/shared/widgets/cards/stats_cards/stats_card.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';

class HowItWorks extends ConsumerStatefulWidget {
  const HowItWorks({super.key});

  @override
  ConsumerState<HowItWorks> createState() => _HowItWorksState();
}

class _HowItWorksState extends ConsumerState<HowItWorks> {
  late UserModel? loggedUser;
  late double totalRevenue = 0;
  RevenueInfoModel? revenueInfo;
  UserStatsModel userStats = const UserStatsModel(level: '/membership_levels/base', valueSpent: 0);

  @override
  void initState() {
    super.initState();

    loggedUser = ref.read(userProvider);

    _fetchTotalRevenue();
    _fetchRevenueInfo();
    _fetchUserStats();
  }

  void _fetchTotalRevenue() {
    ref
        .read(referralsServiceProvider)
        .getRevenueTotalValue(loggedUser!.uuid)
        .then((value) => setState(() => totalRevenue = value));
  }

  void _fetchRevenueInfo() {
    ref.read(referralsServiceProvider).getRevenueInfo(loggedUser!.uuid).then(
            (revenueInfo) => setState(() => this.revenueInfo = revenueInfo));
  }

  void _fetchUserStats() async {
    ref.read(userNotifier).getUserStats(loggedUser!.uuid).then((stats) {
      setState(() {
        userStats = stats;
      });
    });
  }

  void copyReferralCode() async {
    String? referralCode = ref.read(userProvider)?.referralCode;
    String backofficeUrl = '${FlavorConfig.instance.variables["backofficeUrl"]}/#/sign-up/$referralCode';

    // Copy to clipboard
    await Clipboard.setData(ClipboardData(text: backofficeUrl));

    // ToDO style SnackBar (future work
    scaffoldKey.currentState?.showSnackBar(
      SnackBar(
        backgroundColor: const Color(0xFF4237DA),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(4),
        ),
        content: Row(
          children: [
            SvgPicture.asset(Assets.circleCheck, colorFilter: const ColorFilter.mode(AppColors.wildSand, BlendMode.srcIn)),
            const SizedBox(width: 12),
            Expanded(
              child: Text(
                translation(context).copiedToClipboard,
                style: Theme.of(context)
                    .textTheme
                    .titleSmall
                    ?.copyWith(color: AppColors.wildSand),
                softWrap: true,
              ),
            ),
          ],
        ),
        behavior: SnackBarBehavior.floating,
        margin: const EdgeInsets.only(top: 40, left: 20, right: 20, bottom: 20),
        duration: const Duration(seconds: 2),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return ScrollableBody(
        child: Column(
          spacing: 20,
          children: [
            SizedBox(
              width: double.infinity,
              child: StatsCard(
                  title: translation(context).earningsReferrals,
                  icon: SvgPicture.asset(Assets.people),
                  value: totalRevenue.toCurrency(loggedUser?.currency),
                  mode: StatsCardMode.yellow),
            ),

            const HowItWorksCard(),

            revenueInfo?.totalRevenue != null ?
            InformationCard(
                backgroundColor: AppColors.youngBlue,
                color: Colors.white,
                title: translation(context)
                    .referYourFriends(10000.00.toCurrency(loggedUser?.currency)),
                description: translation(context).winUp10,
                question: translation(context).howDoesItWork,
                content: SizedBox(
                    width: 144,
                    child: Button(
                        label: translation(context).copyReferralLink,
                        icon: SvgPicture.asset(Assets.copy),
                        mode: Mode.white,
                        onPressed: copyReferralCode)))
                : MembershipReferralCard(membership: userStats.level.split('/')[2],
                spent: userStats.valueSpent,
                currency: loggedUser!.currency ?? 'EUR',
                extra: totalRevenue)
          ],
        )
    );
  }
}
