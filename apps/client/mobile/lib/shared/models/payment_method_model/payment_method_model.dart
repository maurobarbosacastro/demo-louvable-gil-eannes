import 'package:freezed_annotation/freezed_annotation.dart';

part 'payment_method_model.freezed.dart';
part 'payment_method_model.g.dart';

@freezed
class PaymentMethodModel with _$PaymentMethodModel {
  const factory PaymentMethodModel({
    required String uuid,
    required String name,
    required String code
  }) = _PaymentMethodModel;

  factory PaymentMethodModel.fromJson(Map<String, dynamic> json) => _$PaymentMethodModelFromJson(json);
}
