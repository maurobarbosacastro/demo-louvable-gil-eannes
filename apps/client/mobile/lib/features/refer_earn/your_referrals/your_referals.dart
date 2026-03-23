import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/refer_earn/your_referrals/membership_referral_card.dart';
import 'package:tagpeak/features/refer_earn/your_referrals/my_revenue_chart/my_revenue_chart.dart';
import 'package:tagpeak/features/refer_earn/your_referrals/users_table/users_table.dart';
import 'package:tagpeak/shared/models/referral_info_model/referral_info_model.dart';
import 'package:tagpeak/shared/models/referral_info_model/revenue_info_model/revenue_info_model.dart';
import 'package:tagpeak/shared/models/referral_info_model/user_referral_model/user_referral_model.dart';
import 'package:tagpeak/shared/models/user_stats_model/user_stats_model.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/shared/services/referrals_service.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/cards/information_card.dart';
import 'package:tagpeak/shared/widgets/cards/stats_cards/stats_card.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';

import 'history_chart/history_chart.dart';

class YourReferrals extends ConsumerStatefulWidget {
  const YourReferrals({super.key});

  @override
  ConsumerState<YourReferrals> createState() => _YourReferralsState();
}

class _YourReferralsState extends ConsumerState<YourReferrals> {
  late UserModel? loggedUser;
  late double totalRevenue = 0;
  ReferralInfoModel? referralInfo;
  RevenueInfoModel? revenueInfo;
  List<UserReferralModel>? users;
  UserStatsModel userStats = const UserStatsModel(level: '/membership_levels/base', valueSpent: 0);
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();

    loggedUser = ref.read(userProvider);

    _fetchTotalRevenue();
    _fetchReferralInfo();
    _fetchRevenueInfo();
    _fetchUsers();
    _fetchUserStats();
  }

  void _fetchTotalRevenue() {
    ref
        .read(referralsServiceProvider)
        .getRevenueTotalValue(loggedUser!.uuid)
        .then((value) => setState(() => totalRevenue = value));
  }

  void _fetchReferralInfo() {
    ref
        .read(referralsServiceProvider)
        .getReferralInfo(loggedUser!.uuid)
        .then((referralInfo) => setState(() {
              this.referralInfo = referralInfo;
              _isLoading = false;
            }))
        .catchError((error) {
      setState(() {
        _isLoading = false;
      });
    });
  }

  void _fetchRevenueInfo() {
    ref.read(referralsServiceProvider).getRevenueInfo(loggedUser!.uuid).then(
        (revenueInfo) => setState(() => this.revenueInfo = revenueInfo));
  }

  void _fetchUsers() {
    ref.read(referralsServiceProvider).getUsersRevenueInfo(loggedUser!.uuid).then(
            (users) => setState(() => this.users = users));
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
      children: [
        SizedBox(
          width: double.infinity,
          child: StatsCard(
              title: translation(context).earningsReferrals,
              icon: SvgPicture.asset(Assets.people),
              value: totalRevenue.toCurrency(loggedUser?.currency),
              mode: StatsCardMode.yellow),
        ),
        const SizedBox(height: 50),
        _isLoading ? const SizedBox(
          height: 300,
          child: Center(
            child: CircularProgressIndicator(
              color: Colors.black,
            ),
          ),
        ) :
        referralInfo?.totalClicks == null
            ? Padding(
                padding: const EdgeInsets.symmetric(horizontal: 30),
                child: Column(
                  spacing: 10,
                  children: [
                    Image.asset(Assets.connectedPeople, height: 194),
                    const SizedBox(height: 2.5),
                    Text(translation(context).nothingHere,
                        style: TextTheme.of(context).titleLarge),
                    Text(translation(context).shareReferralLink,
                        style: TextTheme.of(context).titleSmall,
                        textAlign: TextAlign.center)
                  ],
                ),
              )
            : Column(
              spacing: 40,
              children: [
                HistoryChart(value: referralInfo?.totalClicks ?? 0, history: referralInfo?.clicksByMonth ?? [], description: translation(context).linkClicks),
                HistoryChart(value: referralInfo?.totalUserRegistered ?? 0, history: referralInfo?.registeredByMonth ?? [], description: translation(context).registeredReferrals),
                HistoryChart(value: referralInfo?.totalFirstPurchase ?? 0, history: referralInfo?.firstPurchaseByMonth ?? [], description: translation(context).referralsWhoMadeAPurchase),
                MyRevenueChart(totalRevenue: revenueInfo?.totalRevenue ?? 0, revenueByMonth: revenueInfo?.revenueByMonth ?? []),
                if (users != null && users!.isNotEmpty)
                    UsersTable(users: users, currency: loggedUser!.currency ?? 'EUR')
                ],
            ),
        const SizedBox(height: 50),
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
    ));
  }
}
