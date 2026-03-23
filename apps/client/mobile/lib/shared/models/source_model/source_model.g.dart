// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'source_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$PartnerImpl _$$PartnerImplFromJson(Map<String, dynamic> json) =>
    _$PartnerImpl(
      uuid: json['uuid'] as String?,
      name: json['name'] as String,
      code: json['code'] as String,
      eCommercePlatform: json['eCommercePlatform'] as String?,
      commissionRate: (json['commissionRate'] as num?)?.toDouble(),
      validationPeriod: (json['validationPeriod'] as num?)?.toInt(),
      deepLink: json['deepLink'] as String?,
      deepLinkIdentifier: json['deepLinkIdentifier'] as String?,
      subIdentifier: json['subIdentifier'] as String,
      percentageTagpeak: (json['percentageTagpeak'] as num?)?.toDouble(),
      percentageInvested: (json['percentageInvested'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$$PartnerImplToJson(_$PartnerImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'name': instance.name,
      'code': instance.code,
      'eCommercePlatform': instance.eCommercePlatform,
      'commissionRate': instance.commissionRate,
      'validationPeriod': instance.validationPeriod,
      'deepLink': instance.deepLink,
      'deepLinkIdentifier': instance.deepLinkIdentifier,
      'subIdentifier': instance.subIdentifier,
      'percentageTagpeak': instance.percentageTagpeak,
      'percentageInvested': instance.percentageInvested,
    };
