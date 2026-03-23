// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user_payment_method_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$UserPaymentMethodModelImpl _$$UserPaymentMethodModelImplFromJson(
        Map<String, dynamic> json) =>
    _$UserPaymentMethodModelImpl(
      uuid: json['uuid'] as String,
      paymentMethod: json['paymentMethod'] as String,
      bankName: json['bankName'] as String,
      bankAddress: json['bankAddress'] as String,
      bankCountry: json['bankCountry'] as String,
      country: json['country'] as String,
      vat: json['vat'] as String,
      bankAccountTitle: json['bankAccountTitle'] as String,
      iban: json['iban'] as String,
      ibanStatement: IbanStatementModel.fromJson(
          json['ibanStatement'] as Map<String, dynamic>),
      state: json['state'] as String,
    );

Map<String, dynamic> _$$UserPaymentMethodModelImplToJson(
        _$UserPaymentMethodModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
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
