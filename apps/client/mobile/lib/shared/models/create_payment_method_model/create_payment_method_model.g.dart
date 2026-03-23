// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'create_payment_method_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$CreatePaymentMethodModelImpl _$$CreatePaymentMethodModelImplFromJson(
        Map<String, dynamic> json) =>
    _$CreatePaymentMethodModelImpl(
      paymentMethod: json['paymentMethod'] as String,
      bankName: json['bankName'] as String,
      bankAddress: json['bankAddress'] as String,
      bankCountry: json['bankCountry'] as String,
      country: json['country'] as String,
      vat: json['vat'] as String,
      bankAccountTitle: json['bankAccountTitle'] as String,
      iban: json['iban'] as String,
      ibanStatement: json['ibanStatement'] as String,
      state: json['state'] as String?,
    );

Map<String, dynamic> _$$CreatePaymentMethodModelImplToJson(
        _$CreatePaymentMethodModelImpl instance) =>
    <String, dynamic>{
      'paymentMethod': instance.paymentMethod,
      'bankName': instance.bankName,
      'bankAddress': instance.bankAddress,
      'bankCountry': instance.bankCountry,
      'country': instance.country,
      'vat': instance.vat,
      'bankAccountTitle': instance.bankAccountTitle,
      'iban': instance.iban,
      'ibanStatement': instance.ibanStatement,
      'state': instance.state,
    };
