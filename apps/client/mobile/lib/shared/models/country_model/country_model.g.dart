// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'country_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$CountryModelImpl _$$CountryModelImplFromJson(Map<String, dynamic> json) =>
    _$CountryModelImpl(
      uuid: json['uuid'] as String,
      abbreviation: json['abbreviation'] as String,
      createdAt: DateTime.parse(json['createdAt'] as String),
      createdBy: json['createdBy'] as String,
      currency: json['currency'] as String,
      deleted: json['deleted'] as bool,
      deletedAt: json['deletedAt'] == null
          ? null
          : DateTime.parse(json['deletedAt'] as String),
      deletedBy: json['deletedBy'] as String?,
      enabled: json['enabled'] as bool,
      flag: json['flag'] as String?,
      name: json['name'] as String,
      updatedAt: json['updatedAt'] == null
          ? null
          : DateTime.parse(json['updatedAt'] as String),
      updatedBy: json['updatedBy'] as String?,
    );

Map<String, dynamic> _$$CountryModelImplToJson(_$CountryModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'abbreviation': instance.abbreviation,
      'createdAt': instance.createdAt.toIso8601String(),
      'createdBy': instance.createdBy,
      'currency': instance.currency,
      'deleted': instance.deleted,
      'deletedAt': instance.deletedAt?.toIso8601String(),
      'deletedBy': instance.deletedBy,
      'enabled': instance.enabled,
      'flag': instance.flag,
      'name': instance.name,
      'updatedAt': instance.updatedAt?.toIso8601String(),
      'updatedBy': instance.updatedBy,
    };

_$CountryDTOImpl _$$CountryDTOImplFromJson(Map<String, dynamic> json) =>
    _$CountryDTOImpl(
      abbreviation: json['abbreviation'] as String,
      currency: json['currency'] as String,
      name: json['name'] as String,
      enabled: json['enabled'] as bool,
    );

Map<String, dynamic> _$$CountryDTOImplToJson(_$CountryDTOImpl instance) =>
    <String, dynamic>{
      'abbreviation': instance.abbreviation,
      'currency': instance.currency,
      'name': instance.name,
      'enabled': instance.enabled,
    };
