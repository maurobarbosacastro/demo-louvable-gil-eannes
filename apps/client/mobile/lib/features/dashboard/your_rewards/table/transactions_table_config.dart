import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/cashback_table_model/cashback_table_model.dart';
import 'package:tagpeak/shared/widgets/status_container.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/enums/status_enum.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';

Widget storeCell(BuildContext context, CashbackTableModel row, BoxConstraints constraints, String? currency) {

  var isReferralOrCommission = false;
  var referralText = translation(context).referralBonus;
  if (row.reward != null) {
    if (row.reward!.origin == 'REFERRAL') {
      isReferralOrCommission = true;
    }
    if(row.reward!.origin == 'COMMISSION'){
      isReferralOrCommission = true;
      referralText = translation(context).commission;
    }
    if (row.reward!.origin == 'DUPLICATE' && row.reward!.title.toLowerCase().contains('reward')) {
      isReferralOrCommission = true;
    }
  }

  if (isReferralOrCommission) {
    return SizedBox(
      width: constraints.maxWidth * 0.45,
      child: Row(
        spacing: 10,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          SvgPicture.asset(Assets.tgpk, height: 30, width: 30,),
          Flexible(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(referralText, style: const TextStyle(fontSize: 16, fontWeight: FontWeight.w700, color: AppColors.licorice),
                    overflow: TextOverflow.ellipsis,
                    maxLines: 2)
              ],
            ),
          )
        ],
      ),
    );
  } else {
    return SizedBox(
      width: constraints.maxWidth * 0.45,
      child: Row(
        spacing: 10,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Image.network(
            row.store?.logo ?? '',
            height: 30,
            width: 30,
            fit: BoxFit.contain,
            errorBuilder: (context, error, stackTrace) {
              return Image.asset(Assets.defaultImage, height: 30,);
            },
          ),
          Flexible(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  row.store?.name ?? '',
                  style: const TextStyle(
                      fontSize: 16,
                      fontWeight: FontWeight.w700,
                      color: AppColors.licorice
                  ),
                  overflow: TextOverflow.ellipsis,
                  maxLines: 1,
                ),
                Text('${translation(context).paid} ${row.amountUser.toCurrency(currency)}', style: const TextStyle(fontSize: 12, color: AppColors.waterloo))
              ],
            ),
          )
        ],
      ),
    );
  }
}

Widget rewardCell(CashbackTableModel row, String? currency) {
  if (row.reward != null) {
    return Text(row.reward!.currentRewardUser.toCurrency(currency),
        style: const TextStyle(fontSize: 16,
            fontWeight: FontWeight.w700,
            color: AppColors.licorice));
  } else if (row.status == 'TRACKED' || row.status == 'VALIDATED') {
    var unvalidatedCurrentReward = 0.0;
    if (row.store != null && row.store?.percentageCashout != null ) {
      unvalidatedCurrentReward = row.amountUser * ( row.store!.percentageCashout! / 100);
    }
    return Text(unvalidatedCurrentReward.toCurrency(currency), style: const TextStyle(fontSize: 16,
        fontWeight: FontWeight.w700,
        color: AppColors.slateGray));
  } else {
    return const Text(
        '-', style: TextStyle(fontSize: 16,
        fontWeight: FontWeight.w700,
        color: AppColors.licorice)
    );
  }
}

Widget statusCell(CashbackTableModel row) {
  return Align(
    alignment: Alignment.centerRight,
      child: StatusContainer(status: StatusEnum.fromString(row.reward?.state ?? row.status))
  );
}
