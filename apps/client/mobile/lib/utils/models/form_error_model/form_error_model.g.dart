// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'form_error_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$FormErrorModelImpl _$$FormErrorModelImplFromJson(Map<String, dynamic> json) =>
    _$FormErrorModelImpl(
      isError: json['isError'] as bool? ?? false,
      errorMessage: json['errorMessage'] as String?,
    );

Map<String, dynamic> _$$FormErrorModelImplToJson(
        _$FormErrorModelImpl instance) =>
    <String, dynamic>{
      'isError': instance.isError,
      'errorMessage': instance.errorMessage,
    };
