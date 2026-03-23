import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:go_router/go_router.dart';
import 'package:intl/intl.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/features/reward_details/stop_reward_dialog.dart';
import 'package:tagpeak/features/reward_details/tgpk_line_chart.dart';
import 'package:tagpeak/shared/models/cashback_history_model/cashback_history_model.dart';
import 'package:tagpeak/shared/models/cashback_model/cashback_model.dart';
import 'package:tagpeak/shared/models/reward_model/reward_model.dart';
import 'package:tagpeak/shared/services/rewards_service.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/cards/stats_cards/stats_card.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_dialog.dart';
import 'package:tagpeak/shared/widgets/snack_bar/success_snack_bar.dart';
import 'package:tagpeak/shared/widgets/status_container.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/date_helpers.dart';
import 'package:tagpeak/utils/enums/status_enum.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';

class RewardDetails extends ConsumerStatefulWidget {
  final String id;
  const RewardDetails({super.key, required this.id});

  @override
  ConsumerState<RewardDetails> createState() => _RewardDetailsState();
}

class _RewardDetailsState extends ConsumerState<RewardDetails> {
  CashbackModel? _transaction;
  RewardModel? _reward;
  List<CashbackHistoryModel> _history = [];
  UserModel? user;
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();

    user = ref.read(userProvider.notifier).state;

    _fetchTransaction();
  }

  void _fetchTransaction() async {
    ref.read(rewardsServiceProvider).getTransactionById(widget.id).then((transaction) {
      setState(() {
        _transaction = transaction;
      });
      _fetchReward(transaction.uuid);
    }).catchError((error) {
      setState(() {
        _isLoading = false;
      });
    });
  }

  void _fetchReward(String id) {
    ref.read(rewardsServiceProvider).getTransactionReward(id).then((reward) {
      setState(() {
        _reward = reward;
      });
      _fetchRewardHistory(reward.uuid);
    }).catchError((error) {
      setState(() {
        _isLoading = false;
      });
    });
  }

  void _fetchRewardHistory(String id) {
    ref.read(rewardsServiceProvider).getGraphRewardHistory(id).then((history) {
      setState(() {
        _history = history;
        _isLoading = false;
      });
    }).catchError((error) {
      setState(() {
        _isLoading = false;
      });
    });
  }

  void goToStoreDetails() {
    final String? uuid = _transaction?.store?.uuid;
    if (uuid != null) {
      context.pushNamed(
        RouteNames.storeDetailsRouteName,
        pathParameters: {'id': uuid},
        extra: { 'previousRoute': "/client${RouteNames.dashboardRouteLocation}${RouteNames.rewardRouteLocation}/${_transaction!.uuid}" }
      );
    }
  }

  /// chart methods
  double get maxRewardValue {
    if (_history.isEmpty) return 0;

    return _history.map((item) {
      return item.cashReward;
    }).reduce(max);
  }

  double get minRewardValue {
    return _reward?.minimumReward ?? 0;
  }

  List<Map<String, dynamic>> _getRewardHistoryData() {
    if (_history.isEmpty) return [];

    return _history.map((item) {
      double cashReward = item.cashReward;
      DateTime date = DateTime.parse(item.createdAt);
      return {'x': DateFormat('dd.MM').format(date), 'y': cashReward.roundToDouble()};
    }).toList();
  }

  /// open stop dialog
  void openStopDialog() {
    GenericDialog.openDialog(context,
        content: StopRewardDialog(onConfirm: (BuildContext ctx) {
      ref
          .read(rewardsServiceProvider)
          .stopReward(_reward!.uuid, StatusEnum.stopped.toString())
          .then((_) {
        GenericDialog.closeDialog(ctx);
        _fetchReward(_transaction!.uuid);

        // show snack bar
        showTopSuccessSnackBar(context, translation(context).successStopping);
      });
    }));
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return const Center(
        child: CircularProgressIndicator(
          color: Colors.black,
        ),
      );
    }
    if (_transaction != null && _reward != null) {
      return Container(
        color: Colors.white,
        child: Padding(
          padding: const EdgeInsets.only(top: 80),
          child: SingleChildScrollView(
            padding: const EdgeInsets.only(
                bottom: kBottomNavigationBarHeight + 24,
                top: 20,
                right: 20,
                left: 20
            ),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                // store details
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  spacing: 5,
                  children: [
                    Image.network(
                      _transaction?.store?.logo ?? '',
                      height: 60,
                      errorBuilder: (context, error, stackTrace) {
                        return Image.asset(Assets.defaultImage, height: 60);
                      },
                    ),
                    Text(
                      _transaction!.store?.name ?? '-',
                      style: Theme.of(context).textTheme.headlineMedium,
                    ),
                    Text(
                      '${translation(context).purchasedOn} ${formatDate(_transaction!.orderDate)}',
                      style: Theme.of(context).textTheme.bodySmall,
                    )
                  ],
                ),
                const SizedBox(height: 30),
                // visit store & stop buttons
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Button(label: translation(context).visitStore, mode: Mode.outlinedDark, onPressed: goToStoreDetails),
                    _reward!.state == StatusEnum.live.toString()
                        ? Button(
                        label: translation(context).stop,
                        mode: Mode.red,
                        onPressed: openStopDialog,
                        icon: SvgPicture.asset(Assets.stop)
                    ) : const SizedBox.shrink()
                  ],
                ),
                const SizedBox(height: 15),
                Container(
                    height: 1,
                    color: AppColors.wildSand
                ),
                const SizedBox(height: 15),
                // status
                Column(
                  spacing: 20,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      spacing: 5,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        SvgPicture.asset(Assets.timer),
                        Text(translation(context).status, style: Theme.of(context).textTheme.titleSmall?.copyWith(color: AppColors.licorice))
                      ],
                    ),
                    Row(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        StatusContainer(status: StatusEnum.fromString(_reward!.state)),
                        Column(
                          crossAxisAlignment: CrossAxisAlignment.end,
                          children: [
                            Text(formatDate(_reward!.endDate), style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice)),
                            Text('${differenceBetweenDates(_reward!.endDate, DateTime.now().toString())} ${translation(context).daysLeft}', style: Theme.of(context).textTheme.bodySmall),
                          ],
                        )
                      ],
                    )
                  ],
                ),
                const SizedBox(height: 15),
                Container(
                    height: 1,
                    color: AppColors.wildSand
                ),
                const SizedBox(height: 15),
                // stats cards
                Row(
                  spacing: 20,
                  children: [
                    Expanded(child: StatsCard(
                      isRowHeader: true,
                      title: translation(context).currentReward,
                      icon: SvgPicture.asset(Assets.star),
                      value: double.parse(_reward!.currentRewardUser.toString()).toCurrency(user!.currency),
                    )),
                    Expanded(child: StatsCard(
                      isRowHeader: true,
                      title: translation(context).orderTotalPrice,
                      icon: SvgPicture.asset(Assets.creditCard),
                      value: double.parse(_transaction!.amountUser.toString()).toCurrency(user!.currency),
                    ))
                  ],
                ),
                const SizedBox(height: 45),
                // reward history
                TgpkLineChart(
                  datasets: [
                    TgpkLineChartDataset(spots: _getRewardHistoryData(), color: AppColors.youngBlue)
                  ],
                  minY: minRewardValue,
                  maxY: maxRewardValue,
                  constants: [
                    TgpkConstantLineChart(y: minRewardValue, color: Colors.green)
                  ],
                  tooltipBuilder: (context, trackballDetails) {
                    List<String> parts = trackballDetails.point?.x.split('.');
                    DateTime date = DateTime(
                      DateTime.now().year,
                      int.parse(parts[1]),
                      int.parse(parts[0]),
                    );
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
                          Text(trackballDetails.point?.y?.toDouble().toCurrency(user?.currency) ?? '0.00',
                              style: const TextStyle(
                                  fontSize: 16,
                                  color: AppColors.licorice
                              )),
                          Text(DateFormat('d MMMM').format(date).toString(), style: Theme.of(context).textTheme.bodySmall),
                        ],
                      ),
                    );
                  },
                ),
                const SizedBox(height: 10),
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  spacing: 5,
                  children: [
                    Text(translation(context).stopRewardTitle, style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice)),
                    Text(translation(context).stopRewardDescription, style: Theme.of(context).textTheme.bodySmall),
                    const SizedBox(height: 10),
                    Row(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(translation(context).refId, style: Theme.of(context).textTheme.titleSmall),
                        Text(_transaction!.storeVisit?.reference ?? '-', style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice)),
                      ]
                    ),
                    Row(
                        crossAxisAlignment: CrossAxisAlignment.center,
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(translation(context).purchasedOn, style: Theme.of(context).textTheme.titleSmall),
                          Text(formatDate(_transaction!.orderDate), style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice)),
                        ]
                    ),
                    Row(
                        crossAxisAlignment: CrossAxisAlignment.center,
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(translation(context).stopDate, style: Theme.of(context).textTheme.titleSmall),
                          Text(formatDate(_reward!.endDate), style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice)),
                        ]
                    )
                  ],
                )
              ],
            ),
          ),
        ),
      );
    }
    return const SizedBox.shrink();
  }
}
