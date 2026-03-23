// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'store_visit_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$StoreVisitModelImpl _$$StoreVisitModelImplFromJson(
        Map<String, dynamic> json) =>
    _$StoreVisitModelImpl(
      uuid: json['uuid'] as String,
      user: json['user'] as String?,
      reference: json['reference'] as String,
      store: StoreModel.fromJson(json['store'] as Map<String, dynamic>),
      purchased: json['purchased'] as bool,
      dateTime: json['dateTime'] as String,
    );

Map<String, dynamic> _$$StoreVisitModelImplToJson(
        _$StoreVisitModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'user': instance.user,
      'reference': instance.reference,
      'store': instance.store,
      'purchased': instance.purchased,
      'dateTime': instance.dateTime,
    };

_$StoreVisitAdminModelImpl _$$StoreVisitAdminModelImplFromJson(
        Map<String, dynamic> json) =>
    _$StoreVisitAdminModelImpl(
      uuid: json['uuid'] as String,
      user: json['user'] == null
          ? null
          : UserStoreVisitModel.fromJson(json['user'] as Map<String, dynamic>),
      reference: json['reference'] as String,
      store: StoreModel.fromJson(json['store'] as Map<String, dynamic>),
      purchased: json['purchased'] as bool,
      dateTime: json['dateTime'] as String,
    );

Map<String, dynamic> _$$StoreVisitAdminModelImplToJson(
        _$StoreVisitAdminModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'user': instance.user,
      'reference': instance.reference,
      'store': instance.store,
      'purchased': instance.purchased,
      'dateTime': instance.dateTime,
    };
