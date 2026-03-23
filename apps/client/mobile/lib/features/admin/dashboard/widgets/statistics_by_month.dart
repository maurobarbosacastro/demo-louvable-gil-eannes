import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/dashboard_model/dashboard_model.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/shared/widgets/table/tgpk_scrollable_table.dart';
import 'package:tagpeak/shared/widgets/toggle_panel/toggle_panel.dart';

class StatisticsByMonth extends StatelessWidget {
  final Map<String, MonthCountsModel> dashboardStatistics;
  final Function(String?) onYearChanged;
  final List<DropdownOption> years;
  final int selectedYear;

  const StatisticsByMonth({super.key, required this.dashboardStatistics, required this.onYearChanged, required this.years, required this.selectedYear});

  @override
  Widget build(BuildContext context) {
    final columns = [
      TgpkScrollableTableColumn('1/$selectedYear', translation(context).january),
      TgpkScrollableTableColumn('2/$selectedYear', translation(context).february),
      TgpkScrollableTableColumn('3/$selectedYear', translation(context).march),
      TgpkScrollableTableColumn('4/$selectedYear', translation(context).april),
      TgpkScrollableTableColumn('5/$selectedYear', translation(context).may),
      TgpkScrollableTableColumn('6/$selectedYear', translation(context).june),
      TgpkScrollableTableColumn('7/$selectedYear', translation(context).july),
      TgpkScrollableTableColumn('8/$selectedYear', translation(context).august),
      TgpkScrollableTableColumn('9/$selectedYear', translation(context).september),
      TgpkScrollableTableColumn('10/$selectedYear', translation(context).october),
      TgpkScrollableTableColumn('11/$selectedYear', translation(context).november),
      TgpkScrollableTableColumn('12/$selectedYear', translation(context).december)
    ];

    final metrics = [
      TgpkScrollableTableMetric('totalUsers', translation(context).totalUsers),
      TgpkScrollableTableMetric('activeUsers', translation(context).activeUsers),
      TgpkScrollableTableMetric('numTransaction', translation(context).numTransaction),
      TgpkScrollableTableMetric('totalGMV', translation(context).totalGMV),
      TgpkScrollableTableMetric('avgTransactionAmount', translation(context).avgTransactionAmount),
      TgpkScrollableTableMetric('totalRevenue', translation(context).totalRevenue)
    ];

    return TogglePanel(
      title: translation(context).statisticsByMonth,
      content: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            spacing: 10,
            children: [
              Text(translation(context).year, style: TextTheme.of(context).titleSmall),
              Dropdown(entries: years, onChanged: (year) => onYearChanged(year), label: translation(context).filterByYear, hasClearIcon: false),
            ],
          ),
          TgpkScrollableTable(
            metricsWrapperWidth: 110,
            rowWrapperHeight: 65,
            dataSource: dashboardStatistics,
            columns: columns,
            metrics: metrics
          )
        ],
      ),
    );
  }
}
