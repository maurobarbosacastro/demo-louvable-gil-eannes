// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'referral_info_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$ReferralInfoModelImpl _$$ReferralInfoModelImplFromJson(
        Map<String, dynamic> json) =>
    _$ReferralInfoModelImpl(
      totalClicks: (json['totalClicks'] as num?)?.toInt(),
      totalUserRegistered: (json['totalUserRegistered'] as num).toInt(),
      totalFirstPurchase: (json['totalFirstPurchase'] as num).toInt(),
      clicksByMonth: (json['clicksByMonth'] as List<dynamic>)
          .map((e) => MonthDataModel.fromJson(e as Map<String, dynamic>))
          .toList(),
      registeredByMonth: (json['registeredByMonth'] as List<dynamic>)
          .map((e) => MonthDataModel.fromJson(e as Map<String, dynamic>))
          .toList(),
      firstPurchaseByMonth: (json['firstPurchaseByMonth'] as List<dynamic>)
          .map((e) => MonthDataModel.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$$ReferralInfoModelImplToJson(
        _$ReferralInfoModelImpl instance) =>
    <String, dynamic>{
      'totalClicks': instance.totalClicks,
      'totalUserRegistered': instance.totalUserRegistered,
      'totalFirstPurchase': instance.totalFirstPurchase,
      'clicksByMonth': instance.clicksByMonth,
      'registeredByMonth': instance.registeredByMonth,
      'firstPurchaseByMonth': instance.firstPurchaseByMonth,
    };
