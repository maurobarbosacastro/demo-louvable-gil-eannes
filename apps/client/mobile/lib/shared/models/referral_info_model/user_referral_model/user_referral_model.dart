import 'package:freezed_annotation/freezed_annotation.dart';

part 'user_referral_model.freezed.dart';
part 'user_referral_model.g.dart';

@freezed
class UserReferralModel with _$UserReferralModel {
  const factory UserReferralModel({
    required String uuid,
    required String firstName,
    required String lastName,
    required String profilePicture,
    required double referredValue,
    required bool firstTransactionSuccessful,
    String? displayName
  }) = _UserReferralModel;

  factory UserReferralModel.fromJson(Map<String, dynamic> json) =>
      _$UserReferralModelFromJson(json);
}
