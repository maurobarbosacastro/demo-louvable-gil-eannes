import 'package:freezed_annotation/freezed_annotation.dart';

part 'dashboard_model.freezed.dart';
part 'dashboard_model.g.dart';

@freezed
class DashboardModel with _$DashboardModel {
  const factory DashboardModel({
    CashbackSection? cashbackSection,
    IndicatorsSection? indicatorsSection
  }) = _DashboardModel;

  factory DashboardModel.fromJson(Map<String, dynamic> json) =>
      _$DashboardModelFromJson(json);
}

@freezed
class CashbackSection with _$CashbackSection {
  const factory CashbackSection({
    required int totalValidatedCashbacks,
    required int totalStoppedCashbacks,
    required int totalPaidCashbacks,
    required int totalRequestedCashbacks
  }) = _CashbackSection;

  factory CashbackSection.fromJson(Map<String, dynamic> json) => _$CashbackSectionFromJson(json);
}

@freezed
class IndicatorsSection with _$IndicatorsSection {
  const factory IndicatorsSection({
    required dynamic totalUsers,
    required dynamic activeUsers,
    required dynamic numTransactions,
    required dynamic totalGMV,
    required dynamic averageTransactionAmount,
    required dynamic totalRevenue
  }) = _IndicatorsSection;

  factory IndicatorsSection.fromJson(Map<String, dynamic> json) => _$IndicatorsSectionFromJson(json);
}

@freezed
class MonthCountsModel with _$MonthCountsModel {
  const factory MonthCountsModel({
    double? totalUsers,
    double? activeUsers,
    double? numTransaction,
    double? totalGMV,
    double? avgTransactionAmount,
    double? totalRevenue
  }) = _MonthCountsModel;

  factory MonthCountsModel.fromJson(Map<String, dynamic> json) => _$MonthCountsModelFromJson(json);
}

@freezed
class RewardByCurrencies with _$RewardByCurrencies {
  const factory RewardByCurrencies({
    required String currency,
    required String state,
    required num totalRewards,
  }) = _RewardByCurrencies;

  factory RewardByCurrencies.fromJson(Map<String, dynamic> json) =>
      _$RewardByCurrenciesFromJson(json);
}


