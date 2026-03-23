import 'package:freezed_annotation/freezed_annotation.dart';

part 'create_payment_method_model.freezed.dart';
part 'create_payment_method_model.g.dart';

@freezed
class CreatePaymentMethodModel with _$CreatePaymentMethodModel {
  const factory CreatePaymentMethodModel({
    required String paymentMethod,
    required String bankName,
    required String bankAddress,
    required String bankCountry,
    required String country,
    required String vat,
    required String bankAccountTitle,
    required String iban,
    required String ibanStatement,
    String? state
  }) = _CreatePaymentMethodModel;

  factory CreatePaymentMethodModel.fromJson(Map<String, dynamic> json) => _$CreatePaymentMethodModelFromJson(json);
}
