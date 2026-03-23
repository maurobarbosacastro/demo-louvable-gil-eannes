// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'withdrawal_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$WithdrawalModelImpl _$$WithdrawalModelImplFromJson(
        Map<String, dynamic> json) =>
    _$WithdrawalModelImpl(
      uuid: json['uuid'] as String,
      user: json['user'] as Object,
      amountSource: (json['amountSource'] as num).toDouble(),
      amountTarget: (json['amountTarget'] as num).toDouble(),
      details: json['details'] as String?,
      state: json['state'] as String,
      completionDate: json['completionDate'] as String?,
      createdAt: json['createdAt'] as String,
      createdBy: json['createdBy'] as String,
      updatedAt: json['updatedAt'] as String,
      updatedBy: json['updatedBy'] as String,
    );

Map<String, dynamic> _$$WithdrawalModelImplToJson(
        _$WithdrawalModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'user': instance.user,
      'amountSource': instance.amountSource,
      'amountTarget': instance.amountTarget,
      'details': instance.details,
      'state': instance.state,
      'completionDate': instance.completionDate,
      'createdAt': instance.createdAt,
      'createdBy': instance.createdBy,
      'updatedAt': instance.updatedAt,
      'updatedBy': instance.updatedBy,
    };

_$WithdrawalsRequestImpl _$$WithdrawalsRequestImplFromJson(
        Map<String, dynamic> json) =>
    _$WithdrawalsRequestImpl(
      uuid: json['uuid'] as String,
      user: json['user'],
      state: json['state'] as String,
      amountTarget: (json['amountTarget'] as num).toDouble(),
      currencyTarget: json['currencyTarget'] as String?,
      createdAt: json['createdAt'] as String,
      completionDate: json['completionDate'] as String?,
      details: json['details'] as String?,
      paymentMethod: PaymentMethodModel.fromJson(
          json['paymentMethod'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$WithdrawalsRequestImplToJson(
        _$WithdrawalsRequestImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'user': instance.user,
      'state': instance.state,
      'amountTarget': instance.amountTarget,
      'currencyTarget': instance.currencyTarget,
      'createdAt': instance.createdAt,
      'completionDate': instance.completionDate,
      'details': instance.details,
      'paymentMethod': instance.paymentMethod,
    };

_$PaymentMethodModelImpl _$$PaymentMethodModelImplFromJson(
        Map<String, dynamic> json) =>
    _$PaymentMethodModelImpl(
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
      state: json['state'] as String?,
    );

Map<String, dynamic> _$$PaymentMethodModelImplToJson(
        _$PaymentMethodModelImpl instance) =>
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

_$IbanStatementModelImpl _$$IbanStatementModelImplFromJson(
        Map<String, dynamic> json) =>
    _$IbanStatementModelImpl(
      uuid: json['uuid'] as String,
      fileName: json['fileName'] as String,
    );

Map<String, dynamic> _$$IbanStatementModelImplToJson(
        _$IbanStatementModelImpl instance) =>
    <String, dynamic>{
      'uuid': instance.uuid,
      'fileName': instance.fileName,
    };
