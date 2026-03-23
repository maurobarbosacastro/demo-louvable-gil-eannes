import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/admin/dashboard/widgets/indicators_value_compare/indicator_value_compare_card.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/dashboard_model/dashboard_model.dart';
import 'package:tagpeak/shared/widgets/toggle_panel/toggle_panel.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class IndicatorsComparePerValuesSection extends StatelessWidget {
  final IndicatorsSection? indicatorsSection;
  const IndicatorsComparePerValuesSection({super.key, this.indicatorsSection});

  @override
  Widget build(BuildContext context) {
    return TogglePanel(
      title: translation(context).indicatorsComparePerValues,
      content: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        spacing: 20,
        children: [
          IntrinsicHeight(
            child: Row(
              spacing: 5,
              children: [
                Expanded(
                    child: IndicatorValueCompareCard(
                  title: translation(context).totalUsersLastMonth,
                  content: Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(indicatorsSection?.totalUsers['lastMonth'].toString() ?? '0',
                          style: Theme.of(context).textTheme.headlineLarge),
                      StatChangeIndicator(
                            compareLastMonth: indicatorsSection
                                ?.totalUsers['compareLastMonth'],
                            percentageChange: indicatorsSection
                                ?.totalUsers['percentageChange'])
                    ],
                  ),
                  trailing: indicatorsSection?.totalUsers['allTime'].toString() ?? '0',
                )),
                Expanded(
                    child: IndicatorValueCompareCard(
                  title: translation(context).usersLast12Months,
                  content: Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(indicatorsSection?.activeUsers['last12Months'].toString() ?? '0',
                          style: Theme.of(context).textTheme.headlineLarge),
                      StatChangeIndicator(
                          compareLastMonth: indicatorsSection
                              ?.activeUsers['compareLastMonth'],
                          percentageChange: indicatorsSection
                              ?.activeUsers['percentageChange'])
                    ],
                  ),
                  trailing: indicatorsSection?.activeUsers['allTime'].toString() ?? '0',
                )),
              ],
            ),
          ),
          IntrinsicHeight(
            child: Row(
              spacing: 5,
              children: [
                Expanded(
                    child: IndicatorValueCompareCard(
                  title: translation(context).transactionsCurrentMonth,
                  content: Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(indicatorsSection?.numTransactions['currentMonth'].toString() ?? '0',
                          style: Theme.of(context).textTheme.headlineLarge),
                      StatChangeIndicator(
                          compareLastMonth: indicatorsSection
                              ?.numTransactions['compareLastMonth'],
                          percentageChange: indicatorsSection
                              ?.numTransactions['percentageChange'])
                    ],
                  ),
                  trailing: indicatorsSection?.numTransactions['allTime'].toString() ?? '0',
                )),
                Expanded(
                    child: IndicatorValueCompareCard(
                  title: translation(context).totalGMVCurrentMonth,
                  content: Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(indicatorsSection?.totalGMV['currentMonth'].toString() ?? '0',
                          style: Theme.of(context).textTheme.headlineLarge),
                      StatChangeIndicator(
                          compareLastMonth: indicatorsSection
                              ?.totalGMV['compareLastMonth'],
                          percentageChange: indicatorsSection
                              ?.totalGMV['percentageChange'])
                    ],
                  ),
                  trailing: indicatorsSection?.totalGMV['allTime'].toString() ?? '0',
                )),
              ],
            ),
          ),
          IntrinsicHeight(
            child: Row(
              spacing: 5,
              children: [
                Expanded(
                    child: IndicatorValueCompareCard(
                  title: translation(context).transactionsACGCurrentMonth,
                  content: Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(indicatorsSection?.averageTransactionAmount['currentMonth'].toString() ?? '0',
                          style: Theme.of(context).textTheme.headlineLarge),
                      StatChangeIndicator(
                          compareLastMonth: indicatorsSection
                              ?.averageTransactionAmount['compareLastMonth'],
                          percentageChange: indicatorsSection
                              ?.averageTransactionAmount['percentageChange'])
                    ],
                  ),
                  trailing: indicatorsSection?.averageTransactionAmount['allTime'].toString() ?? '0',
                )),
                Expanded(
                    child: IndicatorValueCompareCard(
                  title: translation(context).totalRevenueCurrentMonth,
                  content: Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(indicatorsSection?.totalRevenue['currentMonth'].toString() ?? '0',
                          style: Theme.of(context).textTheme.headlineLarge),
                      StatChangeIndicator(
                          compareLastMonth: indicatorsSection
                              ?.totalRevenue['compareLastMonth'],
                          percentageChange: indicatorsSection
                              ?.totalRevenue['percentageChange'])
                    ],
                  ),
                  trailing: indicatorsSection?.totalRevenue['allTime'].toString() ?? '0',
                )),
              ],
            ),
          ),
        ],
      )
    );
  }
}

class StatChangeIndicator extends StatelessWidget {
  final String compareLastMonth;
  final int percentageChange;

  const StatChangeIndicator({super.key, required this.compareLastMonth, this.percentageChange = 0});

  (SvgPicture, Color) getComparisonSymbol() {
    final value = compareLastMonth;

    return switch (value) {
      'IMPROVEMENT' => (
        SvgPicture.asset(Assets.improvement, width: 16),
        Colors.green
      ),
      'DOWNGRADE' => (
        SvgPicture.asset(Assets.downgrade, width: 16),
        Colors.red
      ),
      'EQUAL' => (
        SvgPicture.asset(Assets.equal, width: 14),
        AppColors.waterloo
      ),
      _ => (
        SvgPicture.asset(Assets.equal, width: 16),
        AppColors.waterloo
      ),
    };
  }

  @override
  Widget build(BuildContext context) {
    final (symbol, color) = getComparisonSymbol();

    return Padding(
      padding: const EdgeInsets.only(top: 2, left: 4),
      child: Row(
        spacing: 5,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          symbol,
          Text('$percentageChange%',
              style: Theme.of(context).textTheme.titleSmall?.copyWith(color: color))
        ],
      ),
    );
  }
}
