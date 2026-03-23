import 'package:freezed_annotation/freezed_annotation.dart';

part 'user_model.freezed.dart';
part 'user_model.g.dart';

@freezed
class UserModel with _$UserModel {
  factory UserModel({
    required String uuid,
    required String email,
    required String given_name,
    required String family_name,
    required String preferred_username,
    String? profilePicture,
    String? referralCode,
    String? currency,
    String? country,
    String? birthDate,
    bool? newsletter,
    required List<String> roles
  }) = _UserModel;

  factory UserModel.fromJson(Map<String, dynamic> json) => _$UserModelFromJson(json);
}
