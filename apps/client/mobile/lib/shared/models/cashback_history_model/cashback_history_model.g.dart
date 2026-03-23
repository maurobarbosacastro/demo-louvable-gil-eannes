// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'cashback_history_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$CashbackHistoryModelImpl _$$CashbackHistoryModelImplFromJson(
        Map<String, dynamic> json) =>
    _$CashbackHistoryModelImpl(
      uuid: json['uuid'] as String,
      rate: (json['rate'] as num).toDouble(),
      units: (json['units'] as num).toDouble(),
      cashReward: (json['cashReward'] as num).toDouble(),
      createdAt: json['createdAt'] as String,
    );

Map<String, dynamic> _$$CashbackHistoryModelImplToJson(
        _$CashbackHistoryModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'rate': instance.rate,
      'units': instance.units,
      'cashReward': instance.cashReward,
      'createdAt': instance.createdAt,
    };
