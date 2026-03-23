import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:http_interceptor/http/intercepted_client.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/country_model/country_model.dart';
import 'package:tagpeak/shared/models/pagination_model/pagination_model.dart';

class CountryService {
  late http.Client httpInterceptor;
  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  CountryService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  Future<PaginationModel<CountryModel>> getCountries() async {
    final params = <String, String>{
      'sort': 'abbreviation,asc',
    };

    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/countries').replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getCountries - Failed to get countries');
    }

    final pagination = PaginationModel<CountryModel>.fromJson(
      jsonDecode(response.body),
          (json) => CountryModel.fromJson(json),
    );

    return pagination;
  }
}

final countryServiceProvider = Provider(
      (ref) => CountryService(ref.watch(httpInterceptorProvider)),
);
