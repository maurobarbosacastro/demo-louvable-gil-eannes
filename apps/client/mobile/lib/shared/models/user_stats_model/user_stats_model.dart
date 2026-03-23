import 'package:freezed_annotation/freezed_annotation.dart';

part 'user_stats_model.freezed.dart';
part 'user_stats_model.g.dart';

@freezed
class UserStatsModel with _$UserStatsModel {
  const factory UserStatsModel({
    required String level,
    required double valueSpent
  }) = _UserStatsModel;

  factory UserStatsModel.fromJson(Map<String, dynamic> json) => _$UserStatsModelFromJson(json);
}
