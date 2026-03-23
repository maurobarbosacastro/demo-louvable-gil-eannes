// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'cashback_table_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$CashbackTableModelImpl _$$CashbackTableModelImplFromJson(
        Map<String, dynamic> json) =>
    _$CashbackTableModelImpl(
      uuid: json['uuid'] as String,
      exitId: json['exitId'] as String?,
      store: json['store'] == null
          ? null
          : CashbackStoreModel.fromJson(json['store'] as Map<String, dynamic>),
      email: json['email'] as String?,
      date: json['date'] as String,
      amountSource: (json['amountSource'] as num).toDouble(),
      amountTarget: (json['amountTarget'] as num).toDouble(),
      amountUser: (json['amountUser'] as num).toDouble(),
      currencySource: json['currencySource'] as String,
      currencyTarget: json['currencyTarget'] as String,
      networkCommission: (json['networkCommission'] as num?)?.toDouble(),
      status: json['status'] as String,
      cashback: (json['cashback'] as num).toDouble(),
      reward: _rewardFromJson(json['reward']),
      unvalidatedCurrentReward:
          (json['unvalidatedCurrentReward'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$$CashbackTableModelImplToJson(
        _$CashbackTableModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'exitId': instance.exitId,
      'store': instance.store,
      'email': instance.email,
      'date': instance.date,
      'amountSource': instance.amountSource,
      'amountTarget': instance.amountTarget,
      'amountUser': instance.amountUser,
      'currencySource': instance.currencySource,
      'currencyTarget': instance.currencyTarget,
      'networkCommission': instance.networkCommission,
      'status': instance.status,
      'cashback': instance.cashback,
      'reward': _rewardToJson(instance.reward),
      'unvalidatedCurrentReward': instance.unvalidatedCurrentReward,
    };
