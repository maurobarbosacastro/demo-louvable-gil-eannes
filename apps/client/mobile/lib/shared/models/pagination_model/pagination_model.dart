import 'package:freezed_annotation/freezed_annotation.dart';

part 'pagination_model.freezed.dart';

@freezed
class PaginationModel<T> with _$PaginationModel<T> {
  const factory PaginationModel({
    required int limit,
    required int page,
    required String sort,
    required int totalRows,
    required int totalPages,
    required List<T> data
  }) = _PaginationModel<T>;

  factory PaginationModel.fromJson(
      Map<String, dynamic> json,
      T Function(Map<String, dynamic>) fromJsonT,
      ) {
    return PaginationModel<T>(
      limit: json['limit'] as int,
      page: json['page'] as int,
      sort: json['sort'] as String,
      totalRows: json['totalRows'] as int,
      totalPages: json['totalPages'] as int,
      data: (json['data'] as List<dynamic>?)
          ?.map((item) => fromJsonT(item as Map<String, dynamic>))
          .toList() ?? [],
    );
  }
}
