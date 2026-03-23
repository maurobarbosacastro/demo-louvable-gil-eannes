import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/dashboard_model/dashboard_model.dart';
import 'package:tagpeak/shared/widgets/table/tgpk_scrollable_table.dart';
import 'package:tagpeak/shared/widgets/toggle_panel/toggle_panel.dart';
import 'package:tagpeak/utils/constants/currencies.dart';

class RewardRow {
  final Map<String, String> rewardsByState;

  RewardRow(this.rewardsByState);

  factory RewardRow.fromJson(Map<String, dynamic> json) {
    return RewardRow(
      json.map((key, value) => MapEntry(key, value.toString())),
    );
  }

  Map<String, dynamic> toJson() => rewardsByState;

  @override
  String toString() => rewardsByState.toString();
}


class RewardsByCurrency extends StatelessWidget {
  final Map<String, Map<String, RewardByCurrencies>> rewardByCurrencies;

  final columns = currencies.map((currency) {
    return TgpkScrollableTableColumn(currency.key, '${currency.label} (${currency.symbol})');
  }).toList();

  buildMetrics(BuildContext context) {
    return [
      TgpkScrollableTableMetric('tracked', translation(context).tracked),
      TgpkScrollableTableMetric('confirmed', translation(context).confirmed),
      TgpkScrollableTableMetric('rejected', translation(context).rejected),
      TgpkScrollableTableMetric('live', translation(context).live),
      TgpkScrollableTableMetric('expired', translation(context).expired),
      TgpkScrollableTableMetric('stopped', translation(context).stopped),
      TgpkScrollableTableMetric('paid', translation(context).paid),
      TgpkScrollableTableMetric('finished', translation(context).finished)
    ];
  }

  Map<String, RewardRow> transformRewards(BuildContext context) {
    final Map<String, Map<String, String>> tempResult = {};
    final metrics = buildMetrics(context);

    final states = metrics.map((state) => state.key).toList();

    for (final currency in currencies) {
      final upperCurrency = currency.key.toUpperCase();
      tempResult[upperCurrency] = {for (var state in states) state: '0'};
    }

    rewardByCurrencies.forEach((state, currencyMap) {
      currencyMap.forEach((currency, reward) {
        final upperCurrency = currency;
        tempResult.putIfAbsent(upperCurrency, () {
          return { for (var state in states) state: '0' };
        });
        tempResult[upperCurrency]![state.toLowerCase()] = reward.totalRewards.toString();
      });
    });

    final Map<String, RewardRow> parsed = tempResult.map((currency, rowMap) {
      return MapEntry(currency, RewardRow.fromJson(rowMap));
    });

    return parsed;
  }

  RewardsByCurrency({super.key, required this.rewardByCurrencies});

  @override
  Widget build(BuildContext context) {
    return TogglePanel(
      title: translation(context).rewardsByCurrency,
      content: TgpkScrollableTable(
          dataSource: transformRewards(context),
          columns: columns,
          metrics: buildMetrics(context)
      )
    );
  }
}
