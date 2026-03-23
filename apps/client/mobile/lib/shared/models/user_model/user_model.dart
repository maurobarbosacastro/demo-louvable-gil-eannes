import 'package:freezed_annotation/freezed_annotation.dart';

part 'user_model.freezed.dart';
part 'user_model.g.dart';

@freezed
class UserModel with _$UserModel {
  const factory UserModel({
    required String uuid,
    required String email,
    required String country,
    required String referralCode,
    required double balance,
    required int createdAt,
    required String firstName,
    required String lastName,
    required String currency,
    required String displayName,
    required String birthDate,
    required List<String> groups,
    required bool isVerified,
    required bool onboardingFinished,
    String? profilePicture,
    double? transactionPercentage,
    double? rewardPercentage,
    required bool newsletter,
  }) = _UserModel;

  factory UserModel.fromJson(Map<String, dynamic> json) => _$UserModelFromJson(json);
}
