import 'dart:math';

import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/reward_details/tgpk_line_chart.dart';
import 'package:tagpeak/shared/models/referral_info_model/month_data_model/month_data_model.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class MyRevenueChart extends StatelessWidget {
  final int totalRevenue;
  final List<MonthDataModel> revenueByMonth;

  const MyRevenueChart({super.key, required this.totalRevenue, required this.revenueByMonth});

  List<Map<String, dynamic>> _getRevenueHistoryData() {
    if (revenueByMonth.isEmpty) return [];

    return revenueByMonth
        .map((item) => {'x': item.month, 'y': item.value})
        .toList()
        .reversed
        .toList();
  }

  double get maxRevenueValue {
    if (revenueByMonth.isEmpty) return 0;

    return revenueByMonth.map((item) {
      return item.value?.toDouble() ?? 0;
    }).reduce(max);
  }

  double get minRevenueValue {
    if (revenueByMonth.isEmpty) return 0;

    return revenueByMonth.map((item) {
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
            Text(translation(context).myRevenue,
                style: TextTheme.of(context).headlineMedium),
            Text(translation(context).totalGeneratedByReferrals,
                style: TextTheme.of(context).titleSmall)
          ],
        ),
        const SizedBox(height: 20),
        ConstrainedBox(
          constraints: const BoxConstraints(maxHeight: 200),
          child: TgpkLineChart(
            datasets: [
              TgpkLineChartDataset(spots: _getRevenueHistoryData(), color: AppColors.youngBlue)
            ],
            minY: minRevenueValue,
            maxY: maxRevenueValue,
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
