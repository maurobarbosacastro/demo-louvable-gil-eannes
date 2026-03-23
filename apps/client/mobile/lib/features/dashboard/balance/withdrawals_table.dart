import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/dashboard/your_rewards/status_filters.dart';
import 'package:tagpeak/shared/models/withdrawal_model/withdrawal_model.dart';
import 'package:tagpeak/shared/widgets/status_container.dart';
import 'package:tagpeak/shared/widgets/table/tgpk_data_table.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/date_helpers.dart';
import 'package:tagpeak/utils/enums/status_enum.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';

class WithdrawalsTable extends ConsumerWidget {
  final List<WithdrawalModel> withdrawals;
  final Function(int) onPageChanged;
  final int numPages;

  final bool isLoading;
  final bool isSearched;
  const WithdrawalsTable({super.key, required this.withdrawals, required this.onPageChanged, required this.numPages, this.isLoading = false, this.isSearched = false});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    UserModel? user = ref.read(userProvider);

    String truncateUuid(String uuid) {
      return '${uuid.substring(0, 1)}..${uuid.substring(uuid.length - 2)}';
    }

    return TgpkDataTable<WithdrawalModel>(
      rowHeight: 80,
      columns: [
        DataColumn(label: Text(translation(context).id, style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
        DataColumn(label: Text(translation(context).cashoutAmount, style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
        DataColumn(label: Text(translation(context).status, style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
        DataColumn(headingRowAlignment: MainAxisAlignment.end, label: Text(translation(context).dateRequested, style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
      ],
      rowBuilder: (withdrawal, constraints) => [
        SizedBox(
            width: constraints.maxWidth * 0.13,
            child: Text(truncateUuid(withdrawal.uuid), style: const TextStyle(fontSize: 15, color: AppColors.licorice))
        ),
        SizedBox(
          width: constraints.maxWidth * 0.33,
            child: Text(withdrawal.amountSource.toCurrency(user?.currency),
                style: const TextStyle(fontSize: 15, color: AppColors.licorice, fontWeight: FontWeight.w700))),
        StatusContainer(status: StatusEnum.fromString(withdrawal.state), statusList: balanceStatus(context)),
        SizedBox(
          width: double.infinity,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.end,
            children: [
              Text(formatDate(withdrawal.createdAt), style: const TextStyle(fontSize: 15, color: AppColors.waterloo), textAlign: TextAlign.end),
              Text(formatDate(withdrawal.createdAt, pattern: 'HH:mm:ss'), style: const TextStyle(fontSize: 15, color: AppColors.waterloo), textAlign: TextAlign.end)
            ],
          ),
        )
      ],
      items: withdrawals,
      onPageChanged: onPageChanged,
      pages: numPages,
      isLoading: isLoading,
      hasActiveFilters: isSearched,
      errorStateMessage: translation(context).withoutResultsWithdrawals,
      emptyStateMessage: translation(context).withoutResultsAfterSearchWithdrawals
    );
  }
}
