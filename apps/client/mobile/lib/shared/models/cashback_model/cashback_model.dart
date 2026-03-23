import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tagpeak/shared/models/store_model/store_model.dart';
import 'package:tagpeak/shared/models/store_visit_cashback_model/store_visit_cashback_model.dart';

part 'cashback_model.freezed.dart';
part 'cashback_model.g.dart';

@freezed
class CashbackModel with _$CashbackModel {
  const factory CashbackModel({
    required String uuid,
    required double amountSource,
    required double amountTarget,
    required double amountUser,
    required String currencySource,
    required String currencyTarget,
    required String state,
    required double commissionSource,
    required double commissionTarget,
    required double commissionUser,
    required String orderDate,
    required String user,
    required StoreModel store,
    required StoreVisitCashbackModel storeVisit,
    required String currencyExchangeRateUuid,
    required double cashback,
    required String createdAt,
    required String createdBy,
    required String updatedAt,
    String? updatedBy,
  }) = _CashbackModel;

  factory CashbackModel.fromJson(Map<String, dynamic> json) =>
      _$CashbackModelFromJson(json);
}
