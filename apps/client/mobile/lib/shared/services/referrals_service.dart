import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart';
import 'package:http_interceptor/http/intercepted_client.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/referral_info_model/referral_info_model.dart';
import 'package:tagpeak/shared/models/referral_info_model/revenue_info_model/revenue_info_model.dart';
import 'package:tagpeak/shared/models/referral_info_model/user_referral_model/user_referral_model.dart';

class ReferralsService {
  late final Client httpInterceptor;

  ReferralsService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  Future<double> getRevenueTotalValue(String userId) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/auth/$userId/referral/revenue/total-revenue'));

    if (response.statusCode != 200) {
      throw Exception('Error: getRevenueTotalValue - Failed to get revenue total value for user with UUID: $userId');
    }

    return double.parse(response.body.toString());
  }

  Future<ReferralInfoModel> getReferralInfo(String userId) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/auth/$userId/referral/clicks'));

    if (response.statusCode != 200) {
      throw Exception('Error: getReferralInfo - Failed to get referral info for user with UUID: $userId');
    }

    return ReferralInfoModel.fromJson(jsonDecode(response.body));
  }

  Future<RevenueInfoModel> getRevenueInfo(String userId) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/auth/$userId/referral/revenue'));

    if (response.statusCode != 200) {
      throw Exception('Error: getRevenueInfo - Failed to get revenue info for user with UUID: $userId');
    }

    return RevenueInfoModel.fromJson(jsonDecode(response.body));
  }

  Future<List<UserReferralModel>> getUsersRevenueInfo(String userId) async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/auth/$userId/referral/revenue/users-info'));

    if (response.statusCode != 200) {
      throw Exception('Error: getUsersRevenueInfo - Failed to get revenue users info for user with UUID: $userId');
    }

    return (jsonDecode(response.body) as List)
        .map((json) => UserReferralModel.fromJson(json))
        .toList();
  }
}

final referralsServiceProvider = Provider<ReferralsService>((ref) {
  final interceptor = ref.read(httpInterceptorProvider);
  return ReferralsService(interceptor);
});
