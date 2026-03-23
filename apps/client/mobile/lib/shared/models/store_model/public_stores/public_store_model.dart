import 'package:freezed_annotation/freezed_annotation.dart';

part 'public_store_model.freezed.dart';
part 'public_store_model.g.dart';

@freezed
class PublicStoreModel with _$PublicStoreModel {
  const factory PublicStoreModel({
    required String uuid,
    required String name,
    String? logo,
    String? banner,
    String? shortDescription,
    String? description,
    String? averageRewardActivationTime,
    String? storeUrl,
    String? termsAndConditions,
    double? percentageCashout,
    String? metaTitle,
    String? metaKeywords,
    String? metaDescription
  }) = _PublicStoreModel;

  factory PublicStoreModel.fromJson(Map<String, dynamic> json) =>
      _$PublicStoreModelFromJson(json);
}
