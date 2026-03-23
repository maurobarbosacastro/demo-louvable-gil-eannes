import 'package:freezed_annotation/freezed_annotation.dart';

part 'withdrawal_model.freezed.dart';
part 'withdrawal_model.g.dart';

@freezed
class WithdrawalModel with _$WithdrawalModel {
  const factory WithdrawalModel({
    required String uuid,
    required Object user,
    required double amountSource,
    required double amountTarget,
    String? details,
    required String state,
    String? completionDate,
    required String createdAt,
    required String createdBy,
    required String updatedAt,
    required String updatedBy
  }) = _WithdrawalModel;

  factory WithdrawalModel.fromJson(Map<String, dynamic> json) =>
      _$WithdrawalModelFromJson(json);
}
