import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart';
import 'package:http_interceptor/http/intercepted_client.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/email_verified_model/email_verified_model.dart';
import 'package:tagpeak/shared/models/user_stats_model/user_stats_model.dart';
import 'package:tagpeak/shared/models/user_update_model/user_update_model.dart';
import 'package:tagpeak/shared/models/user_model/user_model.dart' as user_data_model;

class UserService {
  late final Client httpInterceptor;

  UserService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  Future<void> setEmailVerified(EmailVerifiedModel body) async {
    final response = await httpInterceptor.post(
        Uri.parse('$_ipAddress/auth/email-verified'),
        body: jsonEncode(body));

    if (response.statusCode != 200) {
      throw Exception('Error: setEmailVerified - Failed to set isVerified to true for user with UUID: ${body.userId}');
    }
  }

  Future<user_data_model.UserModel> updateUser(String id, UserUpdateModel body) async {
    final response = await httpInterceptor.patch(
        Uri.parse('$_ipAddress/auth/$id'),
        body: jsonEncode(body)
    );

    if (response.statusCode != 200) {
      throw Exception('Error: updateUser - Failed to update the user with UUID: ${response.body}');
    }

    Map<String, dynamic> jsonData = jsonDecode(utf8.decode(response.bodyBytes));
    jsonData['newsletter'] = jsonData['newsletter'] ?? false;

    return user_data_model.UserModel.fromJson(jsonData);
  }

  Future<user_data_model.UserModel> getUser() async {
    final response = await httpInterceptor.get(Uri.parse('$_ipAddress/auth/me'));

    if (response.statusCode != 200) {
      throw Exception('Error: getUser - Failed to get the logged user');
    } else {
      return user_data_model.UserModel.fromJson(jsonDecode(response.body));
    }
  }

  Future<Response> resetPassword(String email) async {
    try {
      return await httpInterceptor.post(
          Uri.parse('$_ipAddress/auth/reset'),
          body: jsonEncode({'email': email})
      );
    } catch (e) {
      throw Exception([e]);
    }
  }

  Future<UserStatsModel> getUserStats(String id) async {
    final response = await httpInterceptor.get(Uri.parse('$_ipAddress/auth/$id/stats'));

    if (response.statusCode != 200) {
      throw Exception('Error: getUserStats - Failed to get statistics for user with UUID: $id');
    } else {
      return UserStatsModel.fromJson(jsonDecode(response.body));
    }
  }
}

final userService = Provider<UserService>((ref) {
  final interceptor = ref.watch(httpInterceptorProvider);
  return UserService(interceptor);
});
