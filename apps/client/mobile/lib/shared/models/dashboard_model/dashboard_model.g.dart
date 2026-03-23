// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'dashboard_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$DashboardModelImpl _$$DashboardModelImplFromJson(Map<String, dynamic> json) =>
    _$DashboardModelImpl(
      cashbackSection: json['cashbackSection'] == null
          ? null
          : CashbackSection.fromJson(
              json['cashbackSection'] as Map<String, dynamic>),
      indicatorsSection: json['indicatorsSection'] == null
          ? null
          : IndicatorsSection.fromJson(
              json['indicatorsSection'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$DashboardModelImplToJson(
        _$DashboardModelImpl instance) =>
    <String, dynamic>{
      'cashbackSection': instance.cashbackSection,
      'indicatorsSection': instance.indicatorsSection,
    };

_$CashbackSectionImpl _$$CashbackSectionImplFromJson(
        Map<String, dynamic> json) =>
    _$CashbackSectionImpl(
      totalValidatedCashbacks: (json['totalValidatedCashbacks'] as num).toInt(),
      totalStoppedCashbacks: (json['totalStoppedCashbacks'] as num).toInt(),
      totalPaidCashbacks: (json['totalPaidCashbacks'] as num).toInt(),
      totalRequestedCashbacks: (json['totalRequestedCashbacks'] as num).toInt(),
    );

Map<String, dynamic> _$$CashbackSectionImplToJson(
        _$CashbackSectionImpl instance) =>
    <String, dynamic>{
      'totalValidatedCashbacks': instance.totalValidatedCashbacks,
      'totalStoppedCashbacks': instance.totalStoppedCashbacks,
      'totalPaidCashbacks': instance.totalPaidCashbacks,
      'totalRequestedCashbacks': instance.totalRequestedCashbacks,
    };

_$IndicatorsSectionImpl _$$IndicatorsSectionImplFromJson(
        Map<String, dynamic> json) =>
    _$IndicatorsSectionImpl(
      totalUsers: json['totalUsers'],
      activeUsers: json['activeUsers'],
      numTransactions: json['numTransactions'],
      totalGMV: json['totalGMV'],
      averageTransactionAmount: json['averageTransactionAmount'],
      totalRevenue: json['totalRevenue'],
    );

Map<String, dynamic> _$$IndicatorsSectionImplToJson(
        _$IndicatorsSectionImpl instance) =>
    <String, dynamic>{
      'totalUsers': instance.totalUsers,
      'activeUsers': instance.activeUsers,
      'numTransactions': instance.numTransactions,
      'totalGMV': instance.totalGMV,
      'averageTransactionAmount': instance.averageTransactionAmount,
      'totalRevenue': instance.totalRevenue,
    };

_$MonthCountsModelImpl _$$MonthCountsModelImplFromJson(
        Map<String, dynamic> json) =>
    _$MonthCountsModelImpl(
      totalUsers: (json['totalUsers'] as num?)?.toDouble(),
      activeUsers: (json['activeUsers'] as num?)?.toDouble(),
      numTransaction: (json['numTransaction'] as num?)?.toDouble(),
      totalGMV: (json['totalGMV'] as num?)?.toDouble(),
      avgTransactionAmount: (json['avgTransactionAmount'] as num?)?.toDouble(),
      totalRevenue: (json['totalRevenue'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$$MonthCountsModelImplToJson(
        _$MonthCountsModelImpl instance) =>
    <String, dynamic>{
      'totalUsers': instance.totalUsers,
      'activeUsers': instance.activeUsers,
      'numTransaction': instance.numTransaction,
      'totalGMV': instance.totalGMV,
      'avgTransactionAmount': instance.avgTransactionAmount,
      'totalRevenue': instance.totalRevenue,
    };

_$RewardByCurrenciesImpl _$$RewardByCurrenciesImplFromJson(
        Map<String, dynamic> json) =>
    _$RewardByCurrenciesImpl(
      currency: json['currency'] as String,
      state: json['state'] as String,
      totalRewards: json['totalRewards'] as num,
    );

Map<String, dynamic> _$$RewardByCurrenciesImplToJson(
        _$RewardByCurrenciesImpl instance) =>
    <String, dynamic>{
      'currency': instance.currency,
      'state': instance.state,
      'totalRewards': instance.totalRewards,
    };
