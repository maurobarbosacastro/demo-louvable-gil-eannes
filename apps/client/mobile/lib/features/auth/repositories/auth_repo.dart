import 'dart:convert';

import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/auth/screens/login/provider/login_provider.dart';
import 'package:tagpeak/features/auth/screens/register/models/sign_up_model.dart';
import 'package:tagpeak/utils/logging.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart';

import '../../core/application/core_provider.dart';
import '../../core/application/hive_cache.dart';
import '../services/auth_service.dart';

final authRepo = Provider(
  (ref) => AuthRepo(
    ref.watch(authServiceProvider),
    ref.watch(hiveProvider),
    ref.watch(loginErrorProvider.notifier),
    ref.watch(isSigningInProvider.notifier),
    ref.watch(userProvider.notifier),
  ),
);

class AuthRepo {
  AuthRepo(this.service, this.storage, this.loginErrorProvider, this.isSigningInProvider, this.userProvider);

  final HiveDatabase storage;
  final AuthService service;

  final StateController<FormErrorModel> loginErrorProvider;
  final StateController<bool> isSigningInProvider;
  final StateController<UserModel?> userProvider;
  final _client = FlavorConfig.instance.variables["client"];
  final _clientSecret = FlavorConfig.instance.variables["client_secret"];

  Future<bool> createUser(String name, String email, String password) async {
    final SignUpModel model = SignUpModel(
        firstName: name.split(' ')[0],
        lastName: name.substring(name.indexOf(' ') + 1),
        email: email,
        password: base64Encode(password.codeUnits),
        currency: "EUR"
    );

    try {
      final response = await service.postUser(body: model);
      if (response.statusCode == 200) {
        return true;
      } else {
        return false;
      }
    } catch (e) {
      Log.warning("Error: $e");
      rethrow;
    }
  }

  Future<bool> refreshToken() async {
    final Map<String, String> model = {"grant_type": "refresh_token", "client_id": "$_client", "refresh_token": storage.refreshToken.toString()};
    try {
      final response = await service.login(body: model);
      if (response.statusCode == 200) {
        storage.tokenBox?.put('token', jsonDecode(response.body)["access_token"] as String);
        storage.refreshTokenBox?.put('refresh', jsonDecode(response.body)["refresh_token"] as String);
        isSigningInProvider.update((state) => false);
        return true;
      } else {
        return false;
      }
    } catch (e) {
      rethrow;
    }
  }

  Future<String?> login(String username, String password) async {
    Log.info("Login- Repo");
    final Map<String, String> model = {
      "username": username,
      "password": password,
      "grant_type": "password",
      "client_id": "$_client",
      "client_secret": "$_clientSecret",
      "scope": "email profile openid"
    };
    try {
      final response = await service.login(body: model);
      if (response.statusCode == 200) {
        saveTokens(response);
        return null;
      } else {
        Log.warning(response.body);
        Log.warning(response.statusCode.toString());
        return jsonDecode(response.body)['error_description'];
      }
    } catch (e) {
      Log.warning("Error: $e");
      rethrow;
    }
  }

  void saveTokens(Response response, {String loginMethod = 'email'}) {
    final responseBody = jsonDecode(response.body);
    storage.tokenBox?.put('token', responseBody["access_token"] as String);
    storage.refreshTokenBox?.put('refresh', responseBody["refresh_token"] as String);

    // Save id_token if present (for social auth)
    if (responseBody["id_token"] != null) {
      storage.idTokenBox?.put('idToken', responseBody["id_token"] as String);
    }

    // Save login method
    storage.loginMethodBox?.put('loginMethod', loginMethod);

    isSigningInProvider.update((state) => false);
  }

  Future<bool> signInWithSocial(String social) async {
    try {
      final response = await service.signInWithSocial(social: social);
      if (response.statusCode == 200) {
        saveTokens(response, loginMethod: 'socials');
        return true;
      } else {
        Log.warning(response.body);
        Log.warning(response.statusCode.toString());
        return false;
      }
    } catch (e) {
      Log.warning("Error: $e");
      rethrow;
    }
  }

  Future<bool> logout() async {
    // Check if this was a social login and call the social logout endpoint
    final loginMethod = storage.loginMethod;
    final idToken = storage.idToken;

    if (loginMethod == 'socials' && idToken != null) {
      try {
        // Call the Keycloak logout endpoint for social auth
        await service.logoutSocials(
          idToken: idToken,
          postLogoutRedirectUri: 'tagpeak-scheme://logout-callback',
        );
      } catch (e) {
        Log.warning("Error during social logout, continuing with local logout: $e");
      }
    }

    // Clear all stored tokens and login method
    if (storage.tokenBox != null) {
      await storage.tokenBox?.clear();
    }
    if (storage.refreshTokenBox != null) {
      await storage.refreshTokenBox?.clear();
    }
    if (storage.idTokenBox != null) {
      await storage.idTokenBox?.clear();
    }
    if (storage.loginMethodBox != null) {
      await storage.loginMethodBox?.clear();
    }

    return storage.refreshTokenBox != null && storage.refreshTokenBox!.isEmpty && storage.tokenBox != null && storage.tokenBox!.isEmpty ||
        storage.tokenBox == null ||
        storage.refreshTokenBox == null;
  }
}
