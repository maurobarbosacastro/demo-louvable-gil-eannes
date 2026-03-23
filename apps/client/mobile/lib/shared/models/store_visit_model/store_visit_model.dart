import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tagpeak/shared/models/store_model/store_model.dart';

part 'store_visit_model.freezed.dart';
part 'store_visit_model.g.dart';

@freezed
class StoreVisitModel with _$StoreVisitModel {
  const factory StoreVisitModel({
    required String uuid,
    required String user,
    required String reference,
    required StoreModel store,
    required bool purchased,
    required String dateTime
  }) = _StoreVisitModel;

  factory StoreVisitModel.fromJson(Map<String, dynamic> json) =>
      _$StoreVisitModelFromJson(json);
}
