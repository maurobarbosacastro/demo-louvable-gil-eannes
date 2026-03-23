import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tagpeak/shared/models/user_payment_method_model/iban_statement_model/iban_statement_model.dart';

part 'user_payment_method_model.freezed.dart';
part 'user_payment_method_model.g.dart';

@freezed
class UserPaymentMethodModel with _$UserPaymentMethodModel {
  const factory UserPaymentMethodModel({
    required String uuid,
    required String paymentMethod,
    required String bankName,
    required String bankAddress,
    required String bankCountry,
    required String country,
    required String vat,
    required String bankAccountTitle,
    required String iban,
    required IbanStatementModel ibanStatement,
    required String state
  }) = _UserPaymentMethodModel;

  factory UserPaymentMethodModel.fromJson(Map<String, dynamic> json) => _$UserPaymentMethodModelFromJson(json);
}
