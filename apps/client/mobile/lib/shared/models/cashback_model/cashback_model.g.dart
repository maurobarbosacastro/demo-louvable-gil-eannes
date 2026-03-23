// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'cashback_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$CashbackModelImpl _$$CashbackModelImplFromJson(Map<String, dynamic> json) =>
    _$CashbackModelImpl(
      uuid: json['uuid'] as String,
      amountSource: (json['amountSource'] as num).toDouble(),
      amountTarget: (json['amountTarget'] as num).toDouble(),
      amountUser: (json['amountUser'] as num).toDouble(),
      currencySource: json['currencySource'] as String,
      currencyTarget: json['currencyTarget'] as String,
      state: json['state'] as String,
      commissionSource: (json['commissionSource'] as num).toDouble(),
      commissionTarget: (json['commissionTarget'] as num).toDouble(),
      commissionUser: (json['commissionUser'] as num).toDouble(),
      orderDate: json['orderDate'] as String,
      user: json['user'] as String,
      store: json['store'] == null
          ? null
          : StoreModel.fromJson(json['store'] as Map<String, dynamic>),
      storeVisit: json['storeVisit'] == null
          ? null
          : StoreVisitCashbackModel.fromJson(
              json['storeVisit'] as Map<String, dynamic>),
      currencyExchangeRateUuid: json['currencyExchangeRateUuid'] as String,
      cashback: (json['cashback'] as num).toDouble(),
      createdAt: json['createdAt'] as String,
      createdBy: json['createdBy'] as String,
      updatedAt: json['updatedAt'] as String,
      updatedBy: json['updatedBy'] as String?,
      deleted: json['deleted'] as bool?,
      deletedAt: json['deletedAt'],
      deletedBy: json['deletedBy'],
    );

Map<String, dynamic> _$$CashbackModelImplToJson(_$CashbackModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'amountSource': instance.amountSource,
      'amountTarget': instance.amountTarget,
      'amountUser': instance.amountUser,
      'currencySource': instance.currencySource,
      'currencyTarget': instance.currencyTarget,
      'state': instance.state,
      'commissionSource': instance.commissionSource,
      'commissionTarget': instance.commissionTarget,
      'commissionUser': instance.commissionUser,
      'orderDate': instance.orderDate,
      'user': instance.user,
      'store': instance.store,
      'storeVisit': instance.storeVisit,
      'currencyExchangeRateUuid': instance.currencyExchangeRateUuid,
      'cashback': instance.cashback,
      'createdAt': instance.createdAt,
      'createdBy': instance.createdBy,
      'updatedAt': instance.updatedAt,
      'updatedBy': instance.updatedBy,
      'deleted': instance.deleted,
      'deletedAt': instance.deletedAt,
      'deletedBy': instance.deletedBy,
    };
