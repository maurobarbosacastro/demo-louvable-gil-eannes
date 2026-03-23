// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'balance_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$BalanceModelImpl _$$BalanceModelImplFromJson(Map<String, dynamic> json) =>
    _$BalanceModelImpl(
      amountRewards: (json['amountRewards'] as num).toDouble(),
      amountReferrals: (json['amountReferrals'] as num).toDouble(),
      paidWithdrawals: (json['paidWithdrawals'] as num).toDouble(),
    );

Map<String, dynamic> _$$BalanceModelImplToJson(_$BalanceModelImpl instance) =>
    <String, dynamic>{
      'amountRewards': instance.amountRewards,
      'amountReferrals': instance.amountReferrals,
      'paidWithdrawals': instance.paidWithdrawals,
    };
