import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/models/user_model.dart' as logged_user_model;
import 'package:tagpeak/features/auth/provider/user_info_provider.dart' as logged_user_provider;
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/features/dashboard/balance/withdrawal_dialog.dart';
import 'package:tagpeak/features/dashboard/balance/withdrawals_table.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/balance_model/balance_model.dart';
import 'package:tagpeak/shared/providers/config_provider.dart';
import 'package:tagpeak/shared/widgets/date_picker/date_picker.dart';
import 'package:tagpeak/shared/models/withdrawal_model/withdrawal_model.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/shared/services/withdrawal_service.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/cards/stats_cards/stats_card.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_dialog.dart';
import 'package:tagpeak/shared/widgets/snack_bar/snack_bar.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/shared/models/user_model/user_model.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';

class Balance extends ConsumerStatefulWidget {
  const Balance({super.key});

  @override
  ConsumerState<Balance> createState() => _BalanceState();
}

class _BalanceState extends ConsumerState<Balance> {
  List<WithdrawalModel> withdrawals = [];
  DateTime? startDate;
  DateTime? endDate;
  int numPages = 0;
  int selectedPage = 0;
  bool isLoading = true;
  bool isSearched = false;
  logged_user_model.UserModel? loggedUser;
  UserModel? user;
  BalanceModel? balance;
  Map<String, dynamic>? configs;
  TextEditingController startDateController = TextEditingController();
  TextEditingController endDateController = TextEditingController();

  @override
  void initState() {
    loggedUser = ref.read(logged_user_provider.userProvider);
    ref.read(userNotifier).getUser().then((user) => this.user = user);
    configs = ref.read(configsProvider) as Map<String, dynamic>?;
    _fetchWithdrawals();
    _fetchBalance();

    super.initState();
  }

  void _fetchWithdrawals() {
    if (!isLoading) setState(() => isLoading = true);

    ref.read(withdrawalServiceProvider).getPersonalWithdrawals(
      page: selectedPage + 1,
      filtersBody: {
        if (startDate != null)
          'startDate': startDate!.toUtc().toIso8601String(),
        if (endDate != null) 'endDate': endDate!.toUtc().toIso8601String(),
      },
    ).then((paginatedWithdrawals) {
      setState(() {
        withdrawals = paginatedWithdrawals.data;
        numPages = paginatedWithdrawals.totalPages;

        isLoading = false;
      });
    }).catchError((error) {
      setState(() {
        isLoading = false;
      });
    });
  }

  void _fetchBalance() {
    ref.read(withdrawalServiceProvider).getBalance().then((balance) {
      this.balance = balance;
    });
  }

  void handlePageChanged(int page) {
    setState(() {
      selectedPage = page;
    });
    _fetchWithdrawals();
  }

  @override
  Widget build(BuildContext context) {
    return ScrollableBody(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        spacing: 10,
        children: [
          IntrinsicHeight(
            child: Row(
              spacing: 20,
              crossAxisAlignment: CrossAxisAlignment.stretch,
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Expanded(
                    flex: 0,
                    child: StatsCard(
                        title: translation(context).rewardsAvailable,
                        icon: SvgPicture.asset(Assets.creditCard, colorFilter: const ColorFilter.mode(AppColors.wildSand, BlendMode.srcIn)),
                        value: (user?.balance ?? 0.00).toCurrency(loggedUser?.currency),
                        mode: StatsCardMode.blue,
                        content: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            RichText(
                              text: TextSpan(
                                children: [
                                  TextSpan(text: '${(balance?.amountRewards ?? 0.00).toCurrency(loggedUser?.currency)} ', style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.lavender)),
                                  TextSpan(text: translation(context).rewards, style: Theme.of(context).textTheme.bodySmall?.copyWith(color: AppColors.lavender))
                                ]
                              )
                            ),
                            RichText(
                                text: TextSpan(
                                    children: [
                                      TextSpan(text: '${(balance?.amountReferrals ?? 0.00).toCurrency(loggedUser?.currency)} ', style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.lavender)),
                                      TextSpan(text: translation(context).referrals, style: Theme.of(context).textTheme.bodySmall?.copyWith(color: AppColors.lavender))
                                    ]
                                )
                            ),
                            const SizedBox(height: 5),
                            Container(height: 1, color: AppColors.wildSand),
                            const SizedBox(height: 10),
                            user?.balance != null && user!.balance > 0
                                ? Row(
                                  spacing: 5,
                                  children: [
                                    Button(label: translation(context).withdraw, mode: Mode.outlinedWhite, onPressed: () {
                                      if (user!.balance < double.parse(configs?['withdrawal_balance_minimum']?['value'])) {
                                        showSnackBar(context, translation(context).minWithdrawal(double.parse(configs?['withdrawal_balance_minimum']?['value']).toCurrency(loggedUser?.currency)), SnackBarType.info);
                                        return;
                                      }
                                      ref.read(withdrawalServiceProvider).createNewWithdrawal().then((_) {
                                        if (context.mounted) {
                                          GenericDialog.openDialog(
                                              context,
                                              content: const WithdrawalDialog()
                                          );
                                        }
                                        _fetchWithdrawals();
                                      }).catchError((error) {
                                        if (!context.mounted) return;

                                        if (error.statusCode == 422 || error.statusCode == 409 || error.statusCode == 404) {
                                          showSnackBar(context, error.message, SnackBarType.error);
                                        } else {
                                          showSnackBar(context, translation(context).somethingWentWrong, SnackBarType.error);
                                        }
                                      });
                                    }, height: 29),
                                    Button(
                                        label: translation(context).settings,
                                        mode: Mode.white,
                                        onPressed: () => context.goNamed(RouteNames.settingsRouteName, queryParameters: {'tab': 'withdrawal'}),
                                        height: 29)
                                  ],
                                )
                                : Button(label: translation(context).withdrawalSettings, mode: Mode.white, onPressed: () => context.goNamed(RouteNames.settingsRouteName, queryParameters: {'tab': 'withdrawal'}), height: 29)
                          ],
                        ),
                    )),
                Expanded(
                    child: StatsCard(
                        title: translation(context).paidRewards,
                        icon: SvgPicture.asset(Assets.wallet),
                        value: (balance?.paidWithdrawals ?? 0.00).toCurrency(loggedUser?.currency))),
              ],
            ),
          ),
          const SizedBox(height: 5),
          Row(
            spacing: 10,
            children: [
              Expanded(
                child: DatePicker(
                  controller: startDateController,
                  placeholder: translation(context).dateFrom,
                  prefixIcon: SvgPicture.asset(Assets.calendar),
                  onSelectedDate: (date) {
                    setState(() {
                      startDate = date;
                    });
                  },
                  selectedDate: startDate,
                  onClickClear: () => setState(() {
                    startDate = null;

                    if (startDate == null && endDate == null) {
                      isSearched = false;
                    }

                    _fetchWithdrawals();
                  }),
                ),
              ),
              Expanded(
                child: DatePicker(
                  controller: endDateController,
                  placeholder: translation(context).dateTo,
                  prefixIcon: SvgPicture.asset(Assets.arrowRight),
                  onSelectedDate: (date) {
                    setState(() {
                      endDate = date;
                    });
                  },
                  selectedDate: endDate,
                  onClickClear: () => setState(() {
                    startDate = null;

                    if (startDate == null && endDate == null) {
                      isSearched = false;
                    }

                    _fetchWithdrawals();
                  }),
                ),
              ),
            ],
          ),
          Button(
            label: translation(context).search,
            mode: Mode.outlinedDark,
            onPressed: () {
              _fetchWithdrawals();
              setState(() => isSearched = true);

              if (startDate == null && endDate == null) {
                setState(() => isSearched = false);
              }
            },
          ),
          const SizedBox(height: 20),
          WithdrawalsTable(
            withdrawals: withdrawals,
            onPageChanged: handlePageChanged,
            numPages: numPages,
            isLoading: isLoading,
            isSearched: isSearched,
          )
        ],
      ),
    );
  }
}
