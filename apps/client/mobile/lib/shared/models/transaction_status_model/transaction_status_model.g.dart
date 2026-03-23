// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'transaction_status_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$TransactionStatusModelImpl _$$TransactionStatusModelImplFromJson(
        Map<String, dynamic> json) =>
    _$TransactionStatusModelImpl(
      status: json['status'] as String?,
      count: (json['count'] as num?)?.toInt(),
      warning: (json['warning'] as num?)?.toInt(),
      value: (json['value'] as num?)?.toInt(),
    );

Map<String, dynamic> _$$TransactionStatusModelImplToJson(
        _$TransactionStatusModelImpl instance) =>
    <String, dynamic>{
      'status': instance.status,
      'count': instance.count,
      'warning': instance.warning,
      'value': instance.value,
    };
