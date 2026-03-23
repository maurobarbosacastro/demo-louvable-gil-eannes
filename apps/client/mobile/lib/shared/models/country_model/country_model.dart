import 'package:freezed_annotation/freezed_annotation.dart';

part 'country_model.freezed.dart';
part 'country_model.g.dart';

@freezed
class CountryModel with _$CountryModel {
  const factory CountryModel({
    required String uuid,
    required String abbreviation,
    required DateTime createdAt,
    required String createdBy,
    required String currency,
    required bool deleted,
    DateTime? deletedAt,
    String? deletedBy,
    required bool enabled,
    String? flag,
    required String name,
    required DateTime updatedAt,
    required String updatedBy,
  }) = _CountryModel;

  factory CountryModel.fromJson(Map<String, dynamic> json) => _$CountryModelFromJson(json);
}
