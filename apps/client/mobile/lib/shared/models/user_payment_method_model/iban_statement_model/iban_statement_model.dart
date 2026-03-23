import 'package:freezed_annotation/freezed_annotation.dart';

part 'iban_statement_model.freezed.dart';
part 'iban_statement_model.g.dart';

@freezed
class IbanStatementModel with _$IbanStatementModel {
  const factory IbanStatementModel({
    required String uuid,
    required String fileName
  }) = _IbanStatementModel;

  factory IbanStatementModel.fromJson(Map<String, dynamic> json) => _$IbanStatementModelFromJson(json);
}
