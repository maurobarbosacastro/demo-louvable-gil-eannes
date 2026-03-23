import 'package:freezed_annotation/freezed_annotation.dart';

part 'cashback_history_model.freezed.dart';
part 'cashback_history_model.g.dart';

@freezed
class CashbackHistoryModel with _$CashbackHistoryModel {
  const factory CashbackHistoryModel({
    required String uuid,
    required double rate,
    required double units,
    required double cashReward,
    required String createdAt,
  }) = _CashbackHistoryModel;

  factory CashbackHistoryModel.fromJson(Map<String, dynamic> json) =>
      _$CashbackHistoryModelFromJson(json);
}
