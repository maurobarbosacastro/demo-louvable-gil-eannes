// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'reward_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$RewardModelImpl _$$RewardModelImplFromJson(Map<String, dynamic> json) =>
    _$RewardModelImpl(
      uuid: json['uuid'] as String,
      user: json['user'] as String,
      transactionUuid: json['transactionUuid'] as String,
      isin: json['isin'] as String,
      conid: json['conid'] as String,
      initialReward: (json['initialReward'] as num).toDouble(),
      currentRewardSource: (json['currentRewardSource'] as num).toDouble(),
      currentRewardTarget: (json['currentRewardTarget'] as num).toDouble(),
      currentRewardUser: (json['currentRewardUser'] as num).toDouble(),
      currencyExchangeRateUuid: json['currencyExchangeRateUuid'] as String,
      currencySource: json['currencySource'] as String,
      currencyTarget: json['currencyTarget'] as String,
      currencyUser: json['currencyUser'] as String,
      state: json['state'] as String,
      initialPrice: (json['initialPrice'] as num).toDouble(),
      endDate: json['endDate'] as String,
      assetUnits: (json['assetUnits'] as num).toDouble(),
      history: json['history'],
      type: json['type'] as String,
      title: json['title'] as String,
      details: json['details'] as String,
      createdAt: json['createdAt'] as String,
      createdBy: json['createdBy'] as String,
      updatedAt: json['updatedAt'] as String,
      updatedBy: json['updatedBy'],
      deleted: json['deleted'] as bool,
      deletedAt: json['deletedAt'],
      deletedBy: json['deletedBy'],
      overridePrice: (json['overridePrice'] as num?)?.toDouble(),
      withdrawalUuid: json['withdrawalUuid'] as String?,
      origin: json['origin'] as String,
      minimumReward: (json['minimumReward'] as num?)?.toDouble(),
      unvalidatedCurrentReward:
          (json['unvalidatedCurrentReward'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$$RewardModelImplToJson(_$RewardModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'user': instance.user,
      'transactionUuid': instance.transactionUuid,
      'isin': instance.isin,
      'conid': instance.conid,
      'initialReward': instance.initialReward,
      'currentRewardSource': instance.currentRewardSource,
      'currentRewardTarget': instance.currentRewardTarget,
      'currentRewardUser': instance.currentRewardUser,
      'currencyExchangeRateUuid': instance.currencyExchangeRateUuid,
      'currencySource': instance.currencySource,
      'currencyTarget': instance.currencyTarget,
      'currencyUser': instance.currencyUser,
      'state': instance.state,
      'initialPrice': instance.initialPrice,
      'endDate': instance.endDate,
      'assetUnits': instance.assetUnits,
      'history': instance.history,
      'type': instance.type,
      'title': instance.title,
      'details': instance.details,
      'createdAt': instance.createdAt,
      'createdBy': instance.createdBy,
      'updatedAt': instance.updatedAt,
      'updatedBy': instance.updatedBy,
      'deleted': instance.deleted,
      'deletedAt': instance.deletedAt,
      'deletedBy': instance.deletedBy,
      'overridePrice': instance.overridePrice,
      'withdrawalUuid': instance.withdrawalUuid,
      'origin': instance.origin,
      'minimumReward': instance.minimumReward,
      'unvalidatedCurrentReward': instance.unvalidatedCurrentReward,
    };
