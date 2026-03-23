import 'package:freezed_annotation/freezed_annotation.dart';

part 'sign_up_model.freezed.dart';
part 'sign_up_model.g.dart';

@freezed
class SignUpModel with _$SignUpModel {
  const factory SignUpModel({
    required String firstName,
    required String lastName,
    required String email,
    required String password,
    required String currency,
    String? country
  }) = _SignUpModel;

  factory SignUpModel.fromJson(Map<String, dynamic> json) => _$SignUpModelFromJson(json);
}
