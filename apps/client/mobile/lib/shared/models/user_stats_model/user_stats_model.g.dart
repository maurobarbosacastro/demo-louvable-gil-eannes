// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user_stats_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$UserStatsModelImpl _$$UserStatsModelImplFromJson(Map<String, dynamic> json) =>
    _$UserStatsModelImpl(
      level: json['level'] as String,
      valueSpent: (json['valueSpent'] as num).toDouble(),
    );

Map<String, dynamic> _$$UserStatsModelImplToJson(
        _$UserStatsModelImpl instance) =>
    <String, dynamic>{
      'level': instance.level,
      'valueSpent': instance.valueSpent,
    };
