// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$UserModelImpl _$$UserModelImplFromJson(Map<String, dynamic> json) =>
    _$UserModelImpl(
      uuid: json['uuid'] as String,
      email: json['email'] as String,
      given_name: json['given_name'] as String,
      family_name: json['family_name'] as String,
      preferred_username: json['preferred_username'] as String,
      profilePicture: json['profilePicture'] as String?,
      referralCode: json['referralCode'] as String?,
      currency: json['currency'] as String?,
      country: json['country'] as String?,
      birthDate: json['birthDate'] as String?,
      newsletter: json['newsletter'] as bool?,
      roles: (json['roles'] as List<dynamic>).map((e) => e as String).toList(),
    );

Map<String, dynamic> _$$UserModelImplToJson(_$UserModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'email': instance.email,
      'given_name': instance.given_name,
      'family_name': instance.family_name,
      'preferred_username': instance.preferred_username,
      'profilePicture': instance.profilePicture,
      'referralCode': instance.referralCode,
      'currency': instance.currency,
      'country': instance.country,
      'birthDate': instance.birthDate,
      'newsletter': instance.newsletter,
      'roles': instance.roles,
    };
