// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'public_store_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$PublicStoreModelImpl _$$PublicStoreModelImplFromJson(
        Map<String, dynamic> json) =>
    _$PublicStoreModelImpl(
      uuid: json['uuid'] as String,
      name: json['name'] as String,
      logo: json['logo'] as String?,
      banner: json['banner'] as String?,
      shortDescription: json['shortDescription'] as String?,
      description: json['description'] as String?,
      averageRewardActivationTime:
          json['averageRewardActivationTime'] as String?,
      storeUrl: json['storeUrl'] as String?,
      termsAndConditions: json['termsAndConditions'] as String?,
      percentageCashout: (json['percentageCashout'] as num?)?.toDouble(),
      metaTitle: json['metaTitle'] as String?,
      metaKeywords: json['metaKeywords'] as String?,
      metaDescription: json['metaDescription'] as String?,
    );

Map<String, dynamic> _$$PublicStoreModelImplToJson(
        _$PublicStoreModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'name': instance.name,
      'logo': instance.logo,
      'banner': instance.banner,
      'shortDescription': instance.shortDescription,
      'description': instance.description,
      'averageRewardActivationTime': instance.averageRewardActivationTime,
      'storeUrl': instance.storeUrl,
      'termsAndConditions': instance.termsAndConditions,
      'percentageCashout': instance.percentageCashout,
      'metaTitle': instance.metaTitle,
      'metaKeywords': instance.metaKeywords,
      'metaDescription': instance.metaDescription,
    };
