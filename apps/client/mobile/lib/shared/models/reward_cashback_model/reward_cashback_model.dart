import 'package:freezed_annotation/freezed_annotation.dart';

part 'reward_cashback_model.freezed.dart';
part 'reward_cashback_model.g.dart';

@freezed
class RewardCashbackModel with _$RewardCashbackModel {
  const factory RewardCashbackModel({
    required String uuid,
    required String isin,
    required String conid,
    required double currentRewardSource,
    required double currentRewardTarget,
    required double currentRewardUser,
    required String state,
    required double initialPrice,
    required String title,
    required String endDate,
    required String origin
  }) = _RewardCashbackModel;

  factory RewardCashbackModel.fromJson(Map<String, dynamic> json) => _$RewardCashbackModelFromJson(json);
}
