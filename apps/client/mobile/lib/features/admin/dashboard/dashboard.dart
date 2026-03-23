import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/admin/dashboard/widgets/indicators_value_compare/indicators_compare_per_values_section.dart';
import 'package:tagpeak/features/admin/dashboard/widgets/rewards_by_currency.dart';
import 'package:tagpeak/features/admin/dashboard/widgets/statistics_by_month.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/dashboard_model/dashboard_model.dart';
import 'package:tagpeak/shared/models/transaction_status_model/transaction_status_model.dart';
import 'package:tagpeak/shared/services/dashboard_service.dart';
import 'package:tagpeak/shared/widgets/cards/stats_cards/stats_card.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/shared/widgets/table/tgpk_scrollable_table.dart';

class Dashboard extends ConsumerStatefulWidget {
  const Dashboard({super.key});

  @override
  ConsumerState<Dashboard> createState() => _DashboardState();
}

class _DashboardState extends ConsumerState<Dashboard> {

  DashboardModel? dashboard;

  Map<String, TransactionStatusModel>? transactionsStatus;
  Map<String, MonthCountsModel>? dashboardStatistics;
  Map<String, Map<String, RewardByCurrencies>>? rewardByCurrencies;

  final int currentYear = DateTime.now().year;
  int selectedYear = DateTime.now().year;

  bool isLoading = true;

  @override
  void initState() {
    super.initState();

    _fetchDashboardTransactions();
    _fetchDashboardValues();
    _fetchDashboardStatistics(selectedYear.toString());
    _fetchRewardByCurrencies();
  }

  void _checkLoadingStatus() {
    if (dashboard != null &&
        transactionsStatus != null &&
        dashboardStatistics != null &&
        rewardByCurrencies != null) {
      setState(() {
        isLoading = false;
      });
    }
  }

  _fetchDashboardTransactions() {
    ref.read(dashboardService).getDashboardTransactions().then((transactionsStatus) {
      setState(() {
        this.transactionsStatus = transactionsStatus;
      });
      _checkLoadingStatus();
    });
  }

  _fetchDashboardValues() {
    ref.read(dashboardService).getDashboardValues().then((dashboard) {
      setState(() {
        this.dashboard = dashboard;
      });
      _checkLoadingStatus();
    });
  }

  _fetchDashboardStatistics(String year) {
    ref.read(dashboardService).getDashboardStatistics(year).then((dashboardStatistics) {
      setState(() {
        this.dashboardStatistics = dashboardStatistics;
      });
      _checkLoadingStatus();
    });
  }

  _fetchRewardByCurrencies() {
    ref.read(dashboardService).getRewardCountByCurrencies().then((rewardByCurrencies) {
      setState(() {
        this.rewardByCurrencies = rewardByCurrencies;
      });
      _checkLoadingStatus();
    });
  }

  List<DropdownOption> generateYears() {
    return List.generate(
      currentYear - 2022 + 1,
          (index) => currentYear - index,
    ).map((year) {
      final yearStr = year.toString();
      return DropdownOption(label: yearStr, value: yearStr, selected: yearStr == selectedYear.toString());
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    if (isLoading) {
      return const Center(
        child: CircularProgressIndicator(
          color: Colors.black,
        ),
      );
    }
    if (dashboard != null && dashboardStatistics != null && rewardByCurrencies != null) {
      return Container(
        color: Colors.white,
        child: Padding(
          padding: const EdgeInsets.only(top: 80),
          child: ScrollableBody(
            bodyBottomSpacing: false,
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child: Column(
                spacing: 20,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(translation(context).dashboard, style: Theme.of(context).textTheme.headlineLarge),

                  // Overview
                  TgpkScrollableTable(
                      metrics: [
                        TgpkScrollableTableMetric('value', translation(context).value),
                        TgpkScrollableTableMetric('count', translation(context).count),
                        TgpkScrollableTableMetric('warning', translation(context).warning)
                      ],
                      dataSource: transactionsStatus,
                      columns: [
                        TgpkScrollableTableColumn('tracked', translation(context).tracked),
                        TgpkScrollableTableColumn('validated', translation(context).validated),
                        TgpkScrollableTableColumn('live', translation(context).live),
                        TgpkScrollableTableColumn('stopped', translation(context).stopped),
                        TgpkScrollableTableColumn('expired', translation(context).expired),
                        TgpkScrollableTableColumn('finished', translation(context).finished)
                      ]
                  ),
                  Column(
                    spacing: 20,
                    children: [
                      IntrinsicHeight(
                        child: Row(
                          spacing: 5,
                          children: [
                            Expanded(child: StatsCard(title: translation(context).totalValidatedCashback, value: dashboard?.cashbackSection?.totalValidatedCashbacks.toString() ?? '0', mode: StatsCardMode.transparentSideline)),
                            Expanded(child: StatsCard(title: translation(context).totalStoppedCashback, value: dashboard?.cashbackSection?.totalStoppedCashbacks.toString() ?? '0', mode: StatsCardMode.transparentSideline)),
                          ],
                        ),
                      ),
                      IntrinsicHeight(
                        child: Row(
                          children: [
                            Expanded(child: StatsCard(title: translation(context).totalPaidCashback, value: dashboard?.cashbackSection?.totalPaidCashbacks.toString() ?? '0', mode: StatsCardMode.transparentSideline)),
                            Expanded(child: StatsCard(title: translation(context).totalRequestedCashback, value: dashboard?.cashbackSection?.totalRequestedCashbacks.toString() ?? '0', mode: StatsCardMode.transparentSideline)),
                          ],
                        ),
                      ),
                    ],
                  ),

                  const SizedBox(height: 5),

                  // Indicators compare per values
                  IndicatorsComparePerValuesSection(indicatorsSection: dashboard?.indicatorsSection),

                  const SizedBox(height: 5),

                  // Statistics by Month
                  StatisticsByMonth(
                    dashboardStatistics: dashboardStatistics!,
                    onYearChanged: (year) {
                      setState(() {
                        selectedYear =
                            year != null ? int.parse(year) : currentYear;
                      });
                      _fetchDashboardStatistics(selectedYear.toString());
                    },
                    years: generateYears(),
                    selectedYear: selectedYear,
                  ),

                  const SizedBox(height: 5),

                  // Rewards by currency
                  RewardsByCurrency(rewardByCurrencies: rewardByCurrencies!)

                ],
              ),
            ),
          ),
        ),
      );
    }
    return const SizedBox.shrink();
  }
}
