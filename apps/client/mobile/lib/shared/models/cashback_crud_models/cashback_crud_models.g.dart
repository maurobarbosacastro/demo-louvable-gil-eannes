// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'cashback_crud_models.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$RewardUpdateDTOImpl _$$RewardUpdateDTOImplFromJson(
        Map<String, dynamic> json) =>
    _$RewardUpdateDTOImpl(
      currentRewardSource: (json['currentRewardSource'] as num?)?.toInt(),
      endDate: json['endDate'] as String?,
      initialPrice: (json['initialPrice'] as num?)?.toDouble(),
      initialReward: (json['initialReward'] as num?)?.toDouble(),
      isin: json['isin'] as String?,
      conid: json['conid'] as String?,
      state: json['state'] as String?,
      initialDate: json['initialDate'] as String?,
      title: json['title'] as String?,
      details: json['details'] as String?,
      overridePrice: (json['overridePrice'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$$RewardUpdateDTOImplToJson(
        _$RewardUpdateDTOImpl instance) =>
    <String, dynamic>{
      'currentRewardSource': instance.currentRewardSource,
      'endDate': instance.endDate,
      'initialPrice': instance.initialPrice,
      'initialReward': instance.initialReward,
      'isin': instance.isin,
      'conid': instance.conid,
      'state': instance.state,
      'initialDate': instance.initialDate,
      'title': instance.title,
      'details': instance.details,
      'overridePrice': instance.overridePrice,
    };

_$CreateRewardDTOImpl _$$CreateRewardDTOImplFromJson(
        Map<String, dynamic> json) =>
    _$CreateRewardDTOImpl(
      currency: json['currency'] as String?,
      details: json['details'] as String?,
      initialDate: json['initialDate'] as String?,
      endDate: json['endDate'] as String?,
      initialPrice: (json['initialPrice'] as num?)?.toDouble(),
      isin: json['isin'] as String?,
      conid: json['conid'] as String?,
      state: json['state'] as String?,
      title: json['title'] as String?,
      transactionUuid: json['transactionUuid'] as String?,
      type: json['type'] as String?,
    );

Map<String, dynamic> _$$CreateRewardDTOImplToJson(
        _$CreateRewardDTOImpl instance) =>
    <String, dynamic>{
      'currency': instance.currency,
      'details': instance.details,
      'initialDate': instance.initialDate,
      'endDate': instance.endDate,
      'initialPrice': instance.initialPrice,
      'isin': instance.isin,
      'conid': instance.conid,
      'state': instance.state,
      'title': instance.title,
      'transactionUuid': instance.transactionUuid,
      'type': instance.type,
    };

_$TransactionUpdateDTOImpl _$$TransactionUpdateDTOImplFromJson(
        Map<String, dynamic> json) =>
    _$TransactionUpdateDTOImpl(
      amountTarget: (json['amountTarget'] as num?)?.toDouble(),
      commissionTarget: (json['commissionTarget'] as num?)?.toDouble(),
      currencySource: json['currencySource'] as String?,
      orderDate: json['orderDate'] as String?,
      state: json['state'] as String?,
      exitClick: json['exitClick'] as String?,
      uuids:
          (json['uuids'] as List<dynamic>?)?.map((e) => e as String).toList(),
    );

Map<String, dynamic> _$$TransactionUpdateDTOImplToJson(
        _$TransactionUpdateDTOImpl instance) =>
    <String, dynamic>{
      'amountTarget': instance.amountTarget,
      'commissionTarget': instance.commissionTarget,
      'currencySource': instance.currencySource,
      'orderDate': instance.orderDate,
      'state': instance.state,
      'exitClick': instance.exitClick,
      'uuids': instance.uuids,
    };
