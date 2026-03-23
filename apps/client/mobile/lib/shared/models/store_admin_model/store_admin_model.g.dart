// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'store_admin_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$StoreAdminModelImpl _$$StoreAdminModelImplFromJson(
        Map<String, dynamic> json) =>
    _$StoreAdminModelImpl(
      uuid: json['uuid'] as String,
      name: json['name'] as String,
      logo: json['logo'] as String?,
      banner: json['banner'] as String?,
      shortDescription: json['shortDescription'] as String?,
      description: json['description'] as String?,
      urlSlug: json['urlSlug'] as String?,
      initialReward: (json['initialReward'] as num?)?.toDouble(),
      averageRewardActivationTime:
          json['averageRewardActivationTime'] as String?,
      state: json['state'] as String?,
      keywords: json['keywords'] as String?,
      affiliateLink: json['affiliateLink'] as String?,
      storeUrl: json['storeUrl'] as String?,
      termsAndConditions: json['termsAndConditions'] as String?,
      cashbackType: json['cashbackType'] as String?,
      cashbackValue: (json['cashbackValue'] as num?)?.toDouble(),
      percentageCashout: (json['percentageCashout'] as num?)?.toDouble(),
      metaTitle: json['metaTitle'] as String?,
      metaKeywords: json['metaKeywords'] as String?,
      metaDescription: json['metaDescription'] as String?,
      country:
          (json['country'] as List<dynamic>?)?.map((e) => e as String).toList(),
      category: (json['category'] as List<dynamic>?)
          ?.map((e) => e as String)
          .toList(),
      languageCode: json['languageCode'] as String?,
      affiliatePartnerCode: json['affiliatePartnerCode'] as String?,
      partner: json['partner'] as String?,
      transactionFee: (json['transactionFee'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$$StoreAdminModelImplToJson(
        _$StoreAdminModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'name': instance.name,
      'logo': instance.logo,
      'banner': instance.banner,
      'shortDescription': instance.shortDescription,
      'description': instance.description,
      'urlSlug': instance.urlSlug,
      'initialReward': instance.initialReward,
      'averageRewardActivationTime': instance.averageRewardActivationTime,
      'state': instance.state,
      'keywords': instance.keywords,
      'affiliateLink': instance.affiliateLink,
      'storeUrl': instance.storeUrl,
      'termsAndConditions': instance.termsAndConditions,
      'cashbackType': instance.cashbackType,
      'cashbackValue': instance.cashbackValue,
      'percentageCashout': instance.percentageCashout,
      'metaTitle': instance.metaTitle,
      'metaKeywords': instance.metaKeywords,
      'metaDescription': instance.metaDescription,
      'country': instance.country,
      'category': instance.category,
      'languageCode': instance.languageCode,
      'affiliatePartnerCode': instance.affiliatePartnerCode,
      'partner': instance.partner,
      'transactionFee': instance.transactionFee,
    };

_$StoreDTOImpl _$$StoreDTOImplFromJson(Map<String, dynamic> json) =>
    _$StoreDTOImpl(
      name: json['name'] as String,
      urlSlug: json['urlSlug'] as String,
      averageRewardActivationTime:
          json['averageRewardActivationTime'] as String,
      state: json['state'] as String,
      cashbackType: json['cashbackType'] as String,
      cashbackValue: (json['cashbackValue'] as num).toDouble(),
      percentageCashout: (json['percentageCashout'] as num).toDouble(),
      country:
          (json['country'] as List<dynamic>).map((e) => e as String).toList(),
      category:
          (json['category'] as List<dynamic>).map((e) => e as String).toList(),
      languageCode: json['languageCode'] as String?,
      logo: json['logo'] as String?,
      banner: json['banner'] as String?,
      shortDescription: json['shortDescription'] as String?,
      description: json['description'] as String?,
      keywords: json['keywords'] as String?,
      affiliateLink: json['affiliateLink'] as String?,
      storeUrl: json['storeUrl'] as String?,
      termsAndConditions: json['termsAndConditions'] as String?,
      metaTitle: json['metaTitle'] as String?,
      metaKeywords: json['metaKeywords'] as String?,
      metaDescription: json['metaDescription'] as String?,
      affiliatePartnerCode: json['affiliatePartnerCode'] as String?,
      partnerIdentity: json['partnerIdentity'] as String?,
      transactionFee: (json['transactionFee'] as num?)?.toDouble(),
    );

Map<String, dynamic> _$$StoreDTOImplToJson(_$StoreDTOImpl instance) =>
    <String, dynamic>{
      'name': instance.name,
      'urlSlug': instance.urlSlug,
      'averageRewardActivationTime': instance.averageRewardActivationTime,
      'state': instance.state,
      'cashbackType': instance.cashbackType,
      'cashbackValue': instance.cashbackValue,
      'percentageCashout': instance.percentageCashout,
      'country': instance.country,
      'category': instance.category,
      'languageCode': instance.languageCode,
      'logo': instance.logo,
      'banner': instance.banner,
      'shortDescription': instance.shortDescription,
      'description': instance.description,
      'keywords': instance.keywords,
      'affiliateLink': instance.affiliateLink,
      'storeUrl': instance.storeUrl,
      'termsAndConditions': instance.termsAndConditions,
      'metaTitle': instance.metaTitle,
      'metaKeywords': instance.metaKeywords,
      'metaDescription': instance.metaDescription,
      'affiliatePartnerCode': instance.affiliatePartnerCode,
      'partnerIdentity': instance.partnerIdentity,
      'transactionFee': instance.transactionFee,
    };
