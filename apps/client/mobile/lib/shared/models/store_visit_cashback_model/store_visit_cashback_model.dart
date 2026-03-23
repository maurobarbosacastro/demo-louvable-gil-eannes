import 'package:freezed_annotation/freezed_annotation.dart';

part 'store_visit_cashback_model.freezed.dart';
part 'store_visit_cashback_model.g.dart';

@freezed
class StoreVisitCashbackModel with _$StoreVisitCashbackModel {
  const factory StoreVisitCashbackModel({
    required String uuid,
    required String user,
    required String reference,
    required bool purchase,
    required String storeUuid,
    required DateTime createdAt,
  }) = _StoreVisitCashbackModel;

  factory StoreVisitCashbackModel.fromJson(Map<String, dynamic> json) =>
      _$StoreVisitCashbackModelFromJson(json);
}
