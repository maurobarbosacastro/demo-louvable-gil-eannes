import 'package:freezed_annotation/freezed_annotation.dart';

part 'month_data_model.freezed.dart';
part 'month_data_model.g.dart';

@freezed
class MonthDataModel with _$MonthDataModel {
  const factory MonthDataModel({
    required String month,
    @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson) int? value,
  }) = _MonthDataModel;

  factory MonthDataModel.fromJson(Map<String, dynamic> json) =>
      _$MonthDataModelFromJson(json);
}

int? _numberFromJson(dynamic json) {
  if (json == null) return null;
  if (json is int) return json;
  if (json is String) return int.tryParse(json);
  return null;
}

dynamic _numberToJson(int? number) => number;
