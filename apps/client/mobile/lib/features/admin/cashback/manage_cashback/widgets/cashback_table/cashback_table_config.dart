import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/cashback_table_model/cashback_table_model.dart';
import 'package:tagpeak/shared/widgets/status_container.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/enums/status_enum.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';

Widget storeCell(BuildContext context, CashbackTableModel row, BoxConstraints constraints, String? currency) {
  if (row.reward != null && row.reward!.origin == 'REFERRAL') {
    return SizedBox(
      width: constraints.maxWidth * 0.33,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(translation(context).referralBonus, style: const TextStyle(fontSize: 16, fontWeight: FontWeight.w700, color: AppColors.licorice)),
          Text('${translation(context).paid} ${row.amountUser.toCurrency(currency)}', style: const TextStyle(fontSize: 12, color: AppColors.waterloo))
        ],
      )
    );
  } else {
    return SizedBox(
      width: constraints.maxWidth * 0.33,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            row.store?.name ?? '-',
            style: const TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.w700,
              color: AppColors.licorice,
            ),
            overflow: TextOverflow.ellipsis,
            maxLines: 1,
          ),
          Text(
            '${translation(context).paid} ${row.amountUser.toCurrency(currency)}',
            style: const TextStyle(fontSize: 12, color: AppColors.waterloo),
          ),
        ],
      ),
    );
  }
}

Widget rewardCell(CashbackTableModel row, String? currency, BoxConstraints constraints) {
  String rewardText;

  if (row.reward != null) {
    rewardText = row.reward!.currentRewardUser.toCurrency(currency);
  } else if (row.status == 'TRACKED') {
    final computedValue = (row.networkCommission ?? 0) * row.amountUser;
    rewardText = computedValue.toCurrency(currency);
  } else {
    rewardText = '-';
  }

  return SizedBox(
    width: constraints.maxWidth * 0.35,
    child: Text(
      rewardText,
      style: const TextStyle(
        fontSize: 16,
        fontWeight: FontWeight.w700,
        color: AppColors.licorice,
      ),
    ),
  );
}

Widget statusCell(CashbackTableModel row) {
  return StatusContainer(status: StatusEnum.fromString(row.reward?.state ?? row.status));
}
