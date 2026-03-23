import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/features/dashboard/your_rewards/membership_card.dart';
import 'package:tagpeak/features/dashboard/your_rewards/status_filters.dart';
import 'package:tagpeak/features/dashboard/your_rewards/table/transactions_table.dart';
import 'package:tagpeak/features/dashboard/your_rewards/without_transactions.dart';
import 'package:tagpeak/shared/models/cashback_table_model/cashback_table_model.dart';
import 'package:tagpeak/shared/models/user_stats_model/user_stats_model.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/shared/services/rewards_service.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/cards/information_card.dart';
import 'package:tagpeak/shared/widgets/cards/stats_cards/stats_card.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/shared/widgets/status_container.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/enums/status_enum.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';

class YourRewards extends ConsumerStatefulWidget {
  const YourRewards({super.key});

  @override
  ConsumerState<YourRewards> createState() => _YourRewardsState();
}

class _YourRewardsState extends ConsumerState<YourRewards> {
  UserModel? user;
  double balance = 0;
  double liveRewards = 0;
  List<CashbackTableModel> rewards = [];
  UserStatsModel userStats = const UserStatsModel(level: 'base', valueSpent: 0);

  final List<StatusEnum> status = [];

  int numPages = 0;
  int selectedPage = 0;

  @override
  void initState() {
    super.initState();

    user = ref.read(userProvider.notifier).state;

    _fetchUserBalance();
    _fetchLiveRewards();
    _fetchMyTransactions();
    _fetchUserStats();
  }

  void _fetchUserBalance() async {
    ref.read(userNotifier).getUser().then((user) {
      setState(() {
        balance = user.balance;
      });
    });
  }

  void _fetchLiveRewards() async {
    ref.read(rewardsServiceProvider).getUserLiveRewards(user!.uuid).then((liveReward) {
      setState(() {
        liveRewards = liveReward;
      });
    });
  }

  void _fetchMyTransactions() async {
    ref.read(rewardsServiceProvider).getMyTransactions(
        filtersBody: {'state': status.join(',')}, page: selectedPage, size: 10).then((transactions) {
      setState(() {
        rewards = transactions.data;
        numPages = transactions.totalPages;
      });
    });
  }

  void _fetchUserStats() async {
    ref.read(userNotifier).getUserStats(user!.uuid).then((stats) {
      setState(() {
        userStats = stats;
      });
    });
  }

  void handlePageChanged(int page) {
    setState(() {
      selectedPage = page;
    });
    _fetchMyTransactions();
  }

  List<DropdownOption> _buildDropdownOptions() {
    return transactionStatus(context).map((item) =>
        DropdownOption(label: item.label, value: item.value.toString(), selected: status.contains(item.value))).toList();
  }

  void onDropdownStatusChange(StatusEnum value) {
    setState(() {
      status.contains(value) ? status.remove(value) : status.add(value);
    });
    _fetchMyTransactions();
  }

  void onStatusContainerClose(StatusEnum value) {
    setState(() {
      status.remove(value);
    });
    _fetchMyTransactions();
  }

  @override
  Widget build(BuildContext context) {
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

    return ScrollableBody(
      child: Column(
        children: [
          Row(
            spacing: 20,
            children: [
              Expanded(child: StatsCard(
                title: translation(context).liveRewards,
                icon: SvgPicture.asset(Assets.star),
                value: liveRewards.toCurrency(user?.currency),
              )),
              Expanded(child: StatsCard(
                title: translation(context).availableWithdrawal,
                icon: SvgPicture.asset(Assets.creditCard),
                value: balance.toCurrency(user?.currency),
              ))
            ],
          ),
          rewards.isEmpty && status.isEmpty ? const WithoutTransactions() : Padding(
            padding: const EdgeInsets.symmetric(horizontal: 0, vertical: 20),
            child: Column(
              spacing: 20,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(translation(context).filterBy, style: Theme.of(context).textTheme.titleSmall),
                    Dropdown(
                      entries: _buildDropdownOptions(),
                      onChanged: (value) {
                        if (value != null) onDropdownStatusChange(StatusEnum.fromString(value));
                      },
                      dropdownOffset: const Offset(-60, -5),
                      isMultiSelect: true,
                      label: translation(context).status
                    ),
                  ],
                ),
                status.isNotEmpty
                    ? SizedBox(
                  width: double.infinity,
                  child: Wrap(
                    spacing: 8.0,
                    runSpacing: 8.0,
                    children: status.map((item) => StatusContainer(status: item, hasCloseIcon: true, onCloseClick: onStatusContainerClose)).toList(),
                  ),
                ) : const SizedBox.shrink(),
                TransactionsTable(transactions: rewards, onPageChanged: handlePageChanged, numPages: numPages),
              ],
            ),
          ),
          userStats.valueSpent <= 0 ? InformationCard(
              title: translation(context).referYourFriendsAndWin,
              description: translation(context).winUp10,
              question: translation(context).howDoesItWork,
              handleQuestionClick: () => context.pushNamed(RouteNames.referEarnRouteName),
              content: SizedBox(
                width: 144,
                child: Button(
                    label: translation(context).copyReferralLink,
                    icon: SvgPicture.asset(Assets.copy),
                    mode: Mode.white,
                    onPressed: copyReferralCode
                ),
              )
          ) : Padding(
            padding: const EdgeInsets.only(top: 30),
            child: MembershipCard(membership: userStats.level.split('/')[2], spent: userStats.valueSpent, currency: user!.currency!),
          )
        ],
      ),
    );
  }
}
