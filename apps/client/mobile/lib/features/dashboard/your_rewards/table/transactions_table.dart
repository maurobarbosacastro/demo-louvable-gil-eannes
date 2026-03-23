import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/features/dashboard/your_rewards/table/transactions_table_config.dart';
import 'package:tagpeak/shared/models/cashback_table_model/cashback_table_model.dart';
import 'package:tagpeak/shared/widgets/table/tgpk_data_table.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class TransactionsTable extends ConsumerWidget {
  final List<CashbackTableModel> transactions;
  final Function(int) onPageChanged;
  final int numPages;

  const TransactionsTable({super.key, required this.transactions, required this.onPageChanged, required this.numPages});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.read(userProvider);

    return TgpkDataTable<CashbackTableModel>(
        columns: [
          DataColumn(label: Text(translation(context).store, style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
          DataColumn(label: Text(translation(context).currentReward, style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
          DataColumn(headingRowAlignment: MainAxisAlignment.end, label: Text(translation(context).status, style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
        ],
        items: transactions,
        rowBuilder: (transaction, constraints) => [
            storeCell(context, transaction, constraints, user?.currency),
            rewardCell(transaction,  user?.currency),
            statusCell(transaction)
        ],
        onPageChanged: onPageChanged,
        pages: numPages,
        onCellClick: (item) {
          if (item.status == 'TRACKED' || item.status == 'REJECTED') {
            return;
          }

          if (item.status == 'VALIDATED' && item.reward == null) {
            return;
          }

          if (item.store == null) {
            return;
          }

          context.goNamed(RouteNames.rewardRouteName, pathParameters: {'id': item.uuid});
        },
    );
  }
}
