import 'package:freezed_annotation/freezed_annotation.dart';

part 'reward_model.freezed.dart';
part 'reward_model.g.dart';

@freezed
class RewardModel with _$RewardModel {
  const factory RewardModel({
    required String uuid,
    required String user,
    required String transactionUuid,
    required String isin,
    required String conid,
    required double initialReward,
    required double currentRewardSource,
    required double currentRewardTarget,
    required double currentRewardUser,
    required String currencyExchangeRateUuid,
    required String currencySource,
    required String currencyTarget,
    required String currencyUser,
    required String state,
    required double initialPrice,
    required String endDate,
    required double assetUnits,
    dynamic history,
    required String type,
    required String title,
    required String details,
    required String createdAt,
    required String createdBy,
    required String updatedAt,
    dynamic updatedBy,
    required bool deleted,
    dynamic deletedAt,
    dynamic deletedBy,
    double? overridePrice,
    String? withdrawalUuid,
    required String origin,
    double? minimumReward
  }) = _RewardModel;

  factory RewardModel.fromJson(Map<String, dynamic> json) =>
      _$RewardModelFromJson(json);
}
