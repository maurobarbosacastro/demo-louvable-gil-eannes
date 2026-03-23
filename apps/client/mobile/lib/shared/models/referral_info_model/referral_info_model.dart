import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tagpeak/shared/models/referral_info_model/month_data_model/month_data_model.dart';

part 'referral_info_model.freezed.dart';
part 'referral_info_model.g.dart';

@freezed
class ReferralInfoModel with _$ReferralInfoModel {
  const factory ReferralInfoModel({
    int? totalClicks,
    required int totalUserRegistered,
    required int totalFirstPurchase,
    required List<MonthDataModel> clicksByMonth,
    required List<MonthDataModel> registeredByMonth,
    required List<MonthDataModel> firstPurchaseByMonth
  }) = _ReferralInfoModel;

  factory ReferralInfoModel.fromJson(Map<String, dynamic> json) =>
      _$ReferralInfoModelFromJson(json);
}
