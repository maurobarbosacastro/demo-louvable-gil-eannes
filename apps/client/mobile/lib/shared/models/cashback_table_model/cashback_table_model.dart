import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tagpeak/shared/models/reward_cashback_model/reward_cashback_model.dart';
import 'package:tagpeak/shared/models/store_model/store_model.dart';

part 'cashback_table_model.freezed.dart';
part 'cashback_table_model.g.dart';

@freezed
class CashbackTableModel with _$CashbackTableModel {
  const factory CashbackTableModel({
    required String uuid,
    String? exitId,
    CashbackStoreModel? store,
    String? email,
    required String date,
    required double amountSource,
    required double amountTarget,
    required double amountUser,
    required String currencySource,
    required String currencyTarget,
    double? networkCommission,
    required String status,
    required double cashback,
    @JsonKey(
      fromJson: _rewardFromJson,
      toJson: _rewardToJson,
    )
    RewardCashbackModel? reward,
    double? unvalidatedCurrentReward
  }) = _CashbackTableModel;

  factory CashbackTableModel.fromJson(Map<String, dynamic> json) =>
      _$CashbackTableModelFromJson(json);
}

RewardCashbackModel? _rewardFromJson(dynamic json) {
  if (json == null || (json is Map && json.isEmpty)) {
    return null;
  }
  return RewardCashbackModel.fromJson(json);
}

Map<String, dynamic>? _rewardToJson(RewardCashbackModel? reward) {
  return reward?.toJson();
}
