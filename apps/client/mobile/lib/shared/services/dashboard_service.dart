import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart';
import 'package:http_interceptor/http/intercepted_client.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/dashboard_model/dashboard_model.dart';
import 'package:tagpeak/shared/models/transaction_status_model/transaction_status_model.dart';

class DashboardService {
  late Client httpInterceptor;

  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  DashboardService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  Future<DashboardModel> getDashboardValues() async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/dashboard'));

    if (response.statusCode != 200) {
      throw Exception('Error: getDashboardValues - Failed to get dashboard values');
    }

    return DashboardModel.fromJson(jsonDecode(response.body));
  }

  Future<Map<String, TransactionStatusModel>> getDashboardTransactions() async {
    final response = await httpInterceptor.get(
      Uri.parse('$_ipAddress/dashboard/transactions'),
    );

    if (response.statusCode != 200) {
      throw Exception('Error: getDashboardTransactions - Failed to get dashboard values');
    }

    final jsonMap = jsonDecode(response.body) as Map<String, dynamic>;

    return jsonMap.map((key, value) => MapEntry(
      key,
      TransactionStatusModel.fromJson(value as Map<String, dynamic>),
    ));
  }

  Future<Map<String, MonthCountsModel>> getDashboardStatistics(String year) async {
    final response = await httpInterceptor.get(
      Uri.parse('$_ipAddress/dashboard/statistics').replace(queryParameters: {'year': year})
    );

    if (response.statusCode != 200) {
      throw Exception('Error: getDashboardStatistics - Failed to get dashboard values');
    }

    final jsonMap = jsonDecode(response.body) as Map<String, dynamic>;

    return jsonMap.map((key, value) => MapEntry(
      key,
      MonthCountsModel.fromJson(value as Map<String, dynamic>),
    ));
  }

  Future<Map<String, Map<String, RewardByCurrencies>>> getRewardCountByCurrencies() async {
    final response = await httpInterceptor.get(
      Uri.parse('$_ipAddress/dashboard/rewards/currencies/count'),
    );

    if (response.statusCode != 200) {
      throw Exception('Error: getRewardCountByCurrencies - Failed to get reward count by currencies');
    }

    final jsonMap = jsonDecode(response.body) as Map<String, dynamic>;

    return jsonMap.map((key, value) {
      final innerMap = (value as Map<String, dynamic>).map((innerKey, innerValue) {
        return MapEntry(
          innerKey,
          RewardByCurrencies.fromJson(innerValue as Map<String, dynamic>),
        );
      });
      return MapEntry(key, innerMap);
    });
  }
}

final dashboardService = Provider(
      (ref) => DashboardService(ref.watch(httpInterceptorProvider)),
);
