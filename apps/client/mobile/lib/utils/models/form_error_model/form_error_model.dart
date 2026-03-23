import 'package:freezed_annotation/freezed_annotation.dart';

part 'form_error_model.freezed.dart';

part 'form_error_model.g.dart';

@freezed
class FormErrorModel with _$FormErrorModel {
  factory FormErrorModel({
    @Default(false) bool isError,
    String? errorMessage,
  }) = _FormErrorModel;

  factory FormErrorModel.fromJson(Map<String, dynamic> json) => _$FormErrorModelFromJson(json);
// Map<String, dynamic> toJson() => _$OnBoardingModelToJson(this);
}
