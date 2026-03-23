// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$UserModelImpl _$$UserModelImplFromJson(Map<String, dynamic> json) =>
    _$UserModelImpl(
      uuid: json['uuid'] as String,
      email: json['email'] as String,
      country: json['country'] as String,
      referralCode: json['referralCode'] as String,
      balance: (json['balance'] as num).toDouble(),
      createdAt: (json['createdAt'] as num).toInt(),
      firstName: json['firstName'] as String,
      lastName: json['lastName'] as String,
      currency: json['currency'] as String,
      displayName: json['displayName'] as String,
      birthDate: json['birthDate'] as String,
      groups:
          (json['groups'] as List<dynamic>).map((e) => e as String).toList(),
      isVerified: json['isVerified'] as bool,
      onboardingFinished: json['onboardingFinished'] as bool,
      profilePicture: json['profilePicture'] as String?,
      transactionPercentage:
          (json['transactionPercentage'] as num?)?.toDouble(),
      rewardPercentage: (json['rewardPercentage'] as num?)?.toDouble(),
      newsletter: json['newsletter'] as bool,
    );

Map<String, dynamic> _$$UserModelImplToJson(_$UserModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'email': instance.email,
      'country': instance.country,
      'referralCode': instance.referralCode,
      'balance': instance.balance,
      'createdAt': instance.createdAt,
      'firstName': instance.firstName,
      'lastName': instance.lastName,
      'currency': instance.currency,
      'displayName': instance.displayName,
      'birthDate': instance.birthDate,
      'groups': instance.groups,
      'isVerified': instance.isVerified,
      'onboardingFinished': instance.onboardingFinished,
      'profilePicture': instance.profilePicture,
      'transactionPercentage': instance.transactionPercentage,
      'rewardPercentage': instance.rewardPercentage,
      'newsletter': instance.newsletter,
    };
