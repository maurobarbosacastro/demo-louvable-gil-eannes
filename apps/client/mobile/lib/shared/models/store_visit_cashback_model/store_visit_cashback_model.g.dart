// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'store_visit_cashback_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$StoreVisitCashbackModelImpl _$$StoreVisitCashbackModelImplFromJson(
        Map<String, dynamic> json) =>
    _$StoreVisitCashbackModelImpl(
      uuid: json['uuid'] as String,
      user: json['user'] as String,
      reference: json['reference'] as String?,
      purchase: json['purchase'] as bool?,
      storeUuid: json['storeUuid'] as String?,
      createdAt: json['createdAt'] == null
          ? null
          : DateTime.parse(json['createdAt'] as String),
    );

Map<String, dynamic> _$$StoreVisitCashbackModelImplToJson(
        _$StoreVisitCashbackModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'user': instance.user,
      'reference': instance.reference,
      'purchase': instance.purchase,
      'storeUuid': instance.storeUuid,
      'createdAt': instance.createdAt?.toIso8601String(),
    };
