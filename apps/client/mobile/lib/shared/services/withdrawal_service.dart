import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http_interceptor/http_interceptor.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/balance_model/balance_model.dart';
import 'package:tagpeak/shared/models/pagination_model/pagination_model.dart';
import 'package:tagpeak/shared/models/withdrawal_model/withdrawal_model.dart';
import 'package:tagpeak/utils/api_exception.dart';

class WithdrawalService {
  late final Client httpInterceptor;

  WithdrawalService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  Future<PaginationModel<WithdrawalModel>> getPersonalWithdrawals({
    Map<String, String?>? filtersBody,
    int page = 1,
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
        Uri.parse('$_ipAddress/withdrawal/me')
            .replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getPersonalWithdrawals - Failed to get withdrawals');
    }

    final pagination = PaginationModel<WithdrawalModel>.fromJson(
      jsonDecode(response.body),
          (json) => WithdrawalModel.fromJson(json),
    );

    return pagination;
  }

  Future<BalanceModel> getBalance() async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/withdrawal/me/stats'));

    if (response.statusCode != 200) {
      throw Exception('Error: getBalance - Failed to get balance');
    }

    return BalanceModel.fromJson(jsonDecode(response.body));
  }

  Future<void> createNewWithdrawal() async {
    final response = await httpInterceptor.post(
        Uri.parse('$_ipAddress/withdrawal'));

    if (response.statusCode != 201) {
      throw ApiException(response.statusCode, jsonDecode(response.body));
    }
  }
}

final withdrawalServiceProvider = Provider<WithdrawalService>((ref) {
  final interceptor = ref.read(httpInterceptorProvider);
  return WithdrawalService(interceptor);
});
