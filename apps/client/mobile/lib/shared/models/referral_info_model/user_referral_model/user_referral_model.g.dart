// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user_referral_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$UserReferralModelImpl _$$UserReferralModelImplFromJson(
        Map<String, dynamic> json) =>
    _$UserReferralModelImpl(
      uuid: json['uuid'] as String,
      firstName: json['firstName'] as String,
      lastName: json['lastName'] as String,
      profilePicture: json['profilePicture'] as String,
      referredValue: (json['referredValue'] as num).toDouble(),
      firstTransactionSuccessful: json['firstTransactionSuccessful'] as bool,
      displayName: json['displayName'] as String?,
    );

Map<String, dynamic> _$$UserReferralModelImplToJson(
        _$UserReferralModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'firstName': instance.firstName,
      'lastName': instance.lastName,
      'profilePicture': instance.profilePicture,
      'referredValue': instance.referredValue,
      'firstTransactionSuccessful': instance.firstTransactionSuccessful,
      'displayName': instance.displayName,
    };
