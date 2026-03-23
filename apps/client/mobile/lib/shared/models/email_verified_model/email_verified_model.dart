import 'package:freezed_annotation/freezed_annotation.dart';

part 'email_verified_model.freezed.dart';
part 'email_verified_model.g.dart';

@freezed
class EmailVerifiedModel with _$EmailVerifiedModel {
  const factory EmailVerifiedModel({
    required String userId,
    required String isVerified
  }) = _EmailVerifiedModel;

  factory EmailVerifiedModel.fromJson(Map<String, dynamic> json) => _$EmailVerifiedModelFromJson(json);
}
