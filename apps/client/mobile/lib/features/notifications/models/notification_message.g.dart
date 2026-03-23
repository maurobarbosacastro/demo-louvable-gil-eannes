// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'notification_message.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$NotificationMessageImpl _$$NotificationMessageImplFromJson(
        Map<String, dynamic> json) =>
    _$NotificationMessageImpl(
      id: json['id'] as String,
      title: json['title'] as String,
      body: json['body'] as String,
      duration: (json['duration'] as num?)?.toInt() ?? 5000,
      dismissible: json['dismissible'] as bool? ?? true,
    );

Map<String, dynamic> _$$NotificationMessageImplToJson(
        _$NotificationMessageImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'title': instance.title,
      'body': instance.body,
      'duration': instance.duration,
      'dismissible': instance.dismissible,
    };
