// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'reward_cashback_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$RewardCashbackModelImpl _$$RewardCashbackModelImplFromJson(
        Map<String, dynamic> json) =>
    _$RewardCashbackModelImpl(
      uuid: json['uuid'] as String,
      isin: json['isin'] as String,
      conid: json['conid'] as String,
      currentRewardSource: (json['currentRewardSource'] as num).toDouble(),
      currentRewardTarget: (json['currentRewardTarget'] as num).toDouble(),
      currentRewardUser: (json['currentRewardUser'] as num).toDouble(),
      state: json['state'] as String,
      initialPrice: (json['initialPrice'] as num).toDouble(),
      title: json['title'] as String,
      endDate: json['endDate'] as String,
      origin: json['origin'] as String,
    );

Map<String, dynamic> _$$RewardCashbackModelImplToJson(
        _$RewardCashbackModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'isin': instance.isin,
      'conid': instance.conid,
      'currentRewardSource': instance.currentRewardSource,
      'currentRewardTarget': instance.currentRewardTarget,
      'currentRewardUser': instance.currentRewardUser,
      'state': instance.state,
      'initialPrice': instance.initialPrice,
      'title': instance.title,
      'endDate': instance.endDate,
      'origin': instance.origin,
    };
