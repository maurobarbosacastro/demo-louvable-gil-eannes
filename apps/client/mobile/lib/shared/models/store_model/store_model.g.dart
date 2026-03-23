// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'store_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$StoreModelImpl _$$StoreModelImplFromJson(Map<String, dynamic> json) =>
    _$StoreModelImpl(
      uuid: json['uuid'] as String,
      name: json['name'] as String,
      logo: json['logo'] as String?,
    );

Map<String, dynamic> _$$StoreModelImplToJson(_$StoreModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'name': instance.name,
      'logo': instance.logo,
    };

_$CashbackStoreModelImpl _$$CashbackStoreModelImplFromJson(
        Map<String, dynamic> json) =>
    _$CashbackStoreModelImpl(
      uuid: json['uuid'] as String,
      name: json['name'] as String,
      logo: json['logo'] as String?,
      percentageCashout: (json['percentageCashout'] as num?)?.toDouble(),
      cashbackValue: (json['cashbackValue'] as num?)?.toDouble(),
      cashbackType: json['cashbackType'] as String?,
    );

Map<String, dynamic> _$$CashbackStoreModelImplToJson(
        _$CashbackStoreModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'name': instance.name,
      'logo': instance.logo,
      'percentageCashout': instance.percentageCashout,
      'cashbackValue': instance.cashbackValue,
      'cashbackType': instance.cashbackType,
    };
