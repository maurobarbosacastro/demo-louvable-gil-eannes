import 'dart:math';

import 'package:flutter/material.dart';
import 'package:tagpeak/features/reward_details/tgpk_line_chart.dart';
import 'package:tagpeak/shared/models/referral_info_model/month_data_model/month_data_model.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class HistoryChart extends StatelessWidget {
  final int value;
  final List<MonthDataModel> history;
  final String description;

  const HistoryChart({super.key, required this.value, required this.history, required this.description});

  List<Map<String, dynamic>> _getHistoryData() {
    if (history.isEmpty) return [];

    return history
        .map((item) => {'x': item.month, 'y': item.value})
        .toList()
        .reversed
        .toList();
  }

  double get maxValue {
    if (history.isEmpty) return 0;

    return history.map((item) {
      return item.value?.toDouble() ?? 0;
    }).reduce(max);
  }

  double get minValue {
    if (history.isEmpty) return 0;

    return history.map((item) {
      return item.value?.toDouble() ?? 0;
    }).reduce(min);
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        // Link clicks
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          spacing: 5,
          children: [
            Text(description,
                style: TextTheme.of(context).titleSmall),
            Text(value.toString(),
                style: TextTheme.of(context).headlineMedium)
          ],
        ),
        const SizedBox(height: 20),
        ConstrainedBox(
          constraints: const BoxConstraints(maxHeight: 130),
          child: TgpkLineChart(
            horizontalLineColor: AppColors.grandis,
            trackballColor: AppColors.grandis,
            hiddenYAxisLabels: true,
            datasets: [
              TgpkLineChartDataset(spots: _getHistoryData(), color: AppColors.grandis)
            ],
            minY: minValue,
            maxY: maxValue,
            tooltipBuilder: (context, trackballDetails) {
              return Container(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(4),
                    border: Border.all(color: AppColors.wildSand)
                ),
                child: Column(
                  spacing: 5,
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(trackballDetails.point?.y?.toString() ?? '0',
                        style: const TextStyle(
                            fontSize: 16,
                            color: AppColors.licorice
                        )),
                    Text(trackballDetails.point?.x, style: Theme.of(context).textTheme.bodySmall),
                  ],
                ),
              );
            },
          ),
        ),
      ],
    );
  }
}
