import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http_interceptor/http_interceptor.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/pagination_model/pagination_model.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/category_model/category_model.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';
import 'package:tagpeak/shared/models/store_model/store_model.dart';
import 'package:tagpeak/shared/models/store_visit_model/store_visit_model.dart';

class StoresService {
  late final Client httpInterceptor;

  StoresService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  Future<PaginationModel<StoreModel>> getStoresbyUser({
    Map<String, String>? filtersBody,
    String sort = 'date desc',
  }) async {
    final params = <String, String>{
      ...?filtersBody,
      'sort': sort,
    };

    final response = await httpInterceptor.get(Uri.parse('$_ipAddress/store-visit/stores').replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getAllStores - Failed to get stores');
    }

    final pagination = PaginationModel<StoreModel>.fromJson(
      jsonDecode(response.body),
          (json) => StoreModel.fromJson(json),
    );

    return pagination;
  }

  Future<PaginationModel<StoreVisitModel>> getStoresVisitedByUser({
    Map<String, String?>? filtersBody,
    int page = 0,
    int size = 10,
    String sort = 'created_at desc',
  }) async {
    final params = <String, dynamic>{
      ...?filtersBody,
      'page': page.toString(),
      'limit': size.toString(),
      'sort': sort,
    };

    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/store-visit/user')
            .replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getStoresVisitedByUser - Failed to get stores visits');
    }

    final pagination = PaginationModel<StoreVisitModel>.fromJson(
      jsonDecode(response.body),
          (json) => StoreVisitModel.fromJson(json),
    );

    return pagination;
  }

  // Public stores

  Future<PaginationModel<PublicStoreModel>> getPublicStores({
    int page = 0,
    int size = 10,
    String? sort = 'most-popular',
    String? countryCode,
    String? categoryCode,
    String? name
  }) async {
    final params = <String, dynamic>{
      'page': page.toString(),
      'limit': size.toString(),
      'sort': sort,
      'countryCode': countryCode,
      'categoryCode': categoryCode,
      'name': name
    };

    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/public/store')
            .replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getPublicStores - Failed to get public stores');
    }

    final pagination = PaginationModel<PublicStoreModel>.fromJson(
      jsonDecode(response.body),
          (json) => PublicStoreModel.fromJson(json),
    );

    return pagination;
  }

  Future<PaginationModel<CategoryModel>> getPublicCategories({
    int page = 0,
    int size = 10,
    String sort = 'code'
  }) async {
    final params = <String, dynamic>{
      'page': page.toString(),
      'limit': size.toString(),
      'sort': sort,
    };

    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/public/category')
            .replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getPublicCategories - Failed to get public categories');
    }

    final pagination = PaginationModel<CategoryModel>.fromJson(
      jsonDecode(response.body),
          (json) => CategoryModel.fromJson(json),
    );

    return pagination;
  }

  Future<PublicStoreModel> getPublicStoreByID(String id) async {
    final response = await httpInterceptor.get(Uri.parse('$_ipAddress/public/store/$id'));

    if (response.statusCode != 200) {
      throw Exception('Error: getPublicStoreByID - Failed to get public stores');
    }

    return PublicStoreModel.fromJson(jsonDecode(response.body));
  }

  Future<String> getStoreRedirectUrl(String id) async {
    final url = Uri.parse('$_ipAddress/store/$id/redirect');

    final response = await httpInterceptor.get(url);

    if (response.statusCode != 200) {
      throw Exception('Error: getStoreRedirectUrl - Failed to get store redirect url');
    }

    return json.decode(response.body);
  }
}

final storesServiceProvider = Provider<StoresService>((ref) {
  final interceptor = ref.read(httpInterceptorProvider);
  return StoresService(interceptor);
});
