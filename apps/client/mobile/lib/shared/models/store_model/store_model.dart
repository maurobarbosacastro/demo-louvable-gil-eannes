import 'package:freezed_annotation/freezed_annotation.dart';

part 'store_model.freezed.dart';
part 'store_model.g.dart';

@freezed
class StoreModel with _$StoreModel {
  const factory StoreModel({
    required String uuid,
    required String name,
    String? logo,
  }) = _StoreModel;

  factory StoreModel.fromJson(Map<String, dynamic> json) =>
      _$StoreModelFromJson(json);
}


@freezed
class CashbackStoreModel with _$CashbackStoreModel {
  const factory CashbackStoreModel({
    required String uuid,
    required String name,
    String? logo,
    double? percentageCashout,
    double? cashbackValue,
    String? cashbackType
  }) = _CashbackStoreModel;

  factory CashbackStoreModel.fromJson(Map<String, dynamic> json) =>
      _$CashbackStoreModelFromJson(json);
}


