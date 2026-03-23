import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart';
import 'package:http_interceptor/http/intercepted_client.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/cashback_history_model/cashback_history_model.dart';
import 'package:tagpeak/shared/models/cashback_model/cashback_model.dart';
import 'package:tagpeak/shared/models/cashback_table_model/cashback_table_model.dart';
import 'package:tagpeak/shared/models/pagination_model/pagination_model.dart';
import 'package:tagpeak/shared/models/reward_model/reward_model.dart';

class RewardsService {
  late final Client httpInterceptor;

  RewardsService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  Future<double> getUserLiveRewards(String userId) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/reward/$userId/sum-live'));

    if (response.statusCode != 200) {
      throw Exception('Error: getUserLiveRewards - Failed to get live rewards for user with UUID: $userId');
    }

    return double.parse(response.body.toString());
  }

  Future<PaginationModel<CashbackTableModel>> getMyTransactions({
    required Map<String, String> filtersBody,
    required int page,
    required int size,
    String sort = 'date desc',
  }) async {
    final params = <String, String>{
      ...filtersBody,
      'page': page.toString(),
      'limit': size.toString(),
      'sort': sort,
    };

    final response = await httpInterceptor.get(Uri.parse('$_ipAddress/transaction/me').replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getMyTransactions - Failed to get my transactions');
    }

    final pagination = PaginationModel<CashbackTableModel>.fromJson(
      jsonDecode(response.body),
          (json) => CashbackTableModel.fromJson(json),
    );

    return pagination;
  }

  Future<CashbackModel> getTransactionById(String id) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/transaction/$id'));

    if (response.statusCode != 200) {
      throw Exception('Error: getTransactionById - Failed to get live transaction with id: $id');
    }

    return CashbackModel.fromJson(jsonDecode(response.body));
  }

  Future<RewardModel> getTransactionReward(String id) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/transaction/$id/reward'));

    if (response.statusCode != 200) {
      throw Exception('Error: getTransactionById - Failed to get reward for transaction with UUID: $id');
    }

    return RewardModel.fromJson(jsonDecode(response.body));
  }

  Future<List<CashbackHistoryModel>> getRewardHistory(String id) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/reward/$id/history/graph'));

    if (response.statusCode != 200) {
      throw Exception('Error: getRewardHistory - Failed to get history for reward with UUID: $id');
    }

    return (jsonDecode(response.body) as List)
        .map((json) => CashbackHistoryModel.fromJson(json))
        .toList();
  }

  Future<void> stopReward(String id, String status) async {
    await httpInterceptor.patch(
        Uri.parse('$_ipAddress/reward/$id/stop'),
        body: utf8.encode(jsonEncode({'id': id, 'status': status})));

    // ToDO uncommented this when the backend is fixed
    /* if (response.statusCode != 200) {
      throw Exception('Error: stopReward - Failed to stop reward with UUID: $id');
    } */
  }
}

final rewardsServiceProvider = Provider<RewardsService>((ref) {
  final interceptor = ref.read(httpInterceptorProvider);
  return RewardsService(interceptor);
});
