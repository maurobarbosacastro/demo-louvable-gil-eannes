import 'package:freezed_annotation/freezed_annotation.dart';

part 'transaction_status_model.freezed.dart';
part 'transaction_status_model.g.dart';

@freezed
class TransactionStatusModel with _$TransactionStatusModel {
  const factory TransactionStatusModel({
    String? status,
    int? count,
    int? warning,
    int? value,
  }) = _TransactionStatusModel;

  factory TransactionStatusModel.fromJson(Map<String, dynamic> json) =>
      _$TransactionStatusModelFromJson(json);
}
