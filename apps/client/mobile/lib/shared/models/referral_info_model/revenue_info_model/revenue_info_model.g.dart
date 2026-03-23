// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'revenue_info_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$RevenueInfoModelImpl _$$RevenueInfoModelImplFromJson(
        Map<String, dynamic> json) =>
    _$RevenueInfoModelImpl(
      totalRevenue: (json['totalRevenue'] as num).toInt(),
      revenueByMonth: (json['revenueByMonth'] as List<dynamic>)
          .map((e) => MonthDataModel.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$$RevenueInfoModelImplToJson(
        _$RevenueInfoModelImpl instance) =>
    <String, dynamic>{
      'totalRevenue': instance.totalRevenue,
      'revenueByMonth': instance.revenueByMonth,
    };
