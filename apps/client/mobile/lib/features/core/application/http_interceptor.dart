import 'dart:convert';
import 'dart:io';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http_interceptor/http_interceptor.dart';
import 'package:jwt_decode/jwt_decode.dart';
import 'package:tagpeak/features/auth/screens/login/provider/login_provider.dart';
import 'package:tagpeak/features/auth/services/auth_service.dart';

import '../../../utils/logging.dart';
import 'hive_cache.dart';

class ApiInterceptor implements InterceptorContract {
  final HiveDatabase storage;
  final Ref ref;

  ApiInterceptor(this.storage, this.ref);

  final _client = FlavorConfig.instance.variables["client"];
  final _clientSecret = FlavorConfig.instance.variables["client_secret"];

  @override
  Future<BaseRequest> interceptRequest({required BaseRequest request}) async {
    await verifyToken();
    try {
      request.headers[HttpHeaders.contentTypeHeader] = "application/json";
      request.headers[HttpHeaders.acceptHeader] = "application/json";
      request.headers[HttpHeaders.authorizationHeader] =
          'Bearer ${storage.token}';
    } catch (e) {
      print(e);
    }
    return request;
  }

  Future<void> verifyToken() async {
    dynamic token = storage.token;
    bool isExpired = token != null ? Jwt.isExpired(token as String) : false;
    Log.info('Jwt expired : $isExpired');

    if (isExpired) {
      bool refreshed = await _refreshToken();
      if (!refreshed) {
        Log.warning('Token refresh failed');
      }
    }
  }

  Future<bool> _refreshToken() async {
    final Map<String, String> model = {
      "grant_type": "refresh_token",
      "client_id": "$_client",
      "refresh_token": storage.refreshToken.toString(),
      "client_secret": "$_clientSecret",
    };
    try {
      final response = await ref.read(authServiceProvider).login(body: model);
      if (response.statusCode == 200) {
        storage.tokenBox?.put('token', jsonDecode(response.body)["access_token"] as String);
        storage.refreshTokenBox?.put('refresh', jsonDecode(response.body)["refresh_token"] as String);
        ref.read(isSigningInProvider.notifier).update((state) => false);
        return true;
      } else {
        return false;
      }
    } catch (e) {
      print(e);
      rethrow;
    }
  }

  @override
  Future<bool> shouldInterceptRequest() async {
    // TODO: implement shouldInterceptRequest
    return true;
  }

  @override
  Future<bool> shouldInterceptResponse() async {
    // TODO: implement shouldInterceptResponse
    return true;
  }

  @override
  Future<BaseResponse> interceptResponse(
      {required BaseResponse response}) async {
    return response;
  }
}

class LoggerInterceptor implements InterceptorContract {
  @override
  Future<bool> shouldInterceptRequest() {
    // TODO: implement shouldInterceptRequest
    throw UnimplementedError();
  }

  @override
  Future<bool> shouldInterceptResponse() {
    // TODO: implement shouldInterceptResponse
    throw UnimplementedError();
  }

  @override
  Future<BaseRequest> interceptRequest({required BaseRequest request}) async {
    return request;
  }

  @override
  Future<BaseResponse> interceptResponse(
      {required BaseResponse response}) async {
    return response;
  }
}
