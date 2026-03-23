// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'category_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$CategoryModelImpl _$$CategoryModelImplFromJson(Map<String, dynamic> json) =>
    _$CategoryModelImpl(
      uuid: json['uuid'] as String,
      name: json['name'] as String,
      code: json['code'] as String,
    );

Map<String, dynamic> _$$CategoryModelImplToJson(_$CategoryModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'name': instance.name,
      'code': instance.code,
    };

_$CategoryDTOImpl _$$CategoryDTOImplFromJson(Map<String, dynamic> json) =>
    _$CategoryDTOImpl(
      name: json['name'] as String,
      code: json['code'] as String,
    );

Map<String, dynamic> _$$CategoryDTOImplToJson(_$CategoryDTOImpl instance) =>
    <String, dynamic>{
      'name': instance.name,
      'code': instance.code,
    };
