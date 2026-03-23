// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'month_data_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$MonthDataModelImpl _$$MonthDataModelImplFromJson(Map<String, dynamic> json) =>
    _$MonthDataModelImpl(
      month: json['month'] as String,
      value: _numberFromJson(json['value']),
    );

Map<String, dynamic> _$$MonthDataModelImplToJson(
        _$MonthDataModelImpl instance) =>
    <String, dynamic>{
      'month': instance.month,
      'value': _numberToJson(instance.value),
    };
