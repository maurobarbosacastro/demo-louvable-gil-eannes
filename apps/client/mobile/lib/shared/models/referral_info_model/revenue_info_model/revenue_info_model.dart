import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tagpeak/shared/models/referral_info_model/month_data_model/month_data_model.dart';

part 'revenue_info_model.freezed.dart';
part 'revenue_info_model.g.dart';

@freezed
class RevenueInfoModel with _$RevenueInfoModel {
  const factory RevenueInfoModel({
    required int totalRevenue,
    required List<MonthDataModel> revenueByMonth,
  }) = _RevenueInfoModel;

  factory RevenueInfoModel.fromJson(Map<String, dynamic> json) =>
      _$RevenueInfoModelFromJson(json);
}
