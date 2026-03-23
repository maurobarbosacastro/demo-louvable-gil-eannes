import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:http_interceptor/http_interceptor.dart';
import 'package:flutter_web_auth_2/flutter_web_auth_2.dart';
import 'package:tagpeak/utils/logging.dart';

final authServiceProvider = StateProvider<AuthService>((ref) => AuthService());

class AuthService {
  final _ipKeycloakAddress = FlavorConfig.instance.variables["keycloakUrl"];
  final _ipAddress = FlavorConfig.instance.variables["baseUrl"];
  final _userManagement = FlavorConfig.instance.variables["userManagement"];
  final _realm = FlavorConfig.instance.variables["realm"];
  final _client = FlavorConfig.instance.variables["client"];
  final _clientSecret = FlavorConfig.instance.variables["client_secret"];

  Future<http.Response> postUser({required dynamic body}) async {
    try {
      final url = Uri.parse('$_ipAddress/auth/');
      final http.Response response = await http.post(url, headers: {"Content-Type": "application/json"}, body: jsonEncode(body));
      switch (response.statusCode) {
        case 200:
          return response;
        default:
          throw Exception(["${response.statusCode}: ${response.reasonPhrase}"]);
      }
    } catch (e) {
      throw Exception([e]);
    }
  }

  Future<http.Response> login({dynamic body}) async {
    try {
      final url = Uri.parse('$_ipKeycloakAddress/realms/$_realm/protocol/openid-connect/token');
      return await http.post(
        url,
        headers: {"Content-Type": "application/x-www-form-urlencoded"},
        body: body,
      );
    } catch (e) {
      throw Exception([e]);
    }
  }

  Future<http.Response> validateAction({dynamic body}) async {
    try {
      final url = Uri.parse('$_ipAddress/public/auth/validate-action');

      final uriWithParams = url.replace(
        queryParameters: body,
      );

      return await http.get(
        uriWithParams,
      );
    } catch (e) {
      throw Exception([e]);
    }
  }

  Future<http.Response> signInWithSocial({required String social}) async {
    try {
      final url = '$_ipKeycloakAddress/realms/$_realm/protocol/openid-connect/auth';
      final authUri = Uri.parse(url).replace(queryParameters: {
        "kc_idp_hint": social,
        "redirect_uri": "tagpeak-scheme://login-callback",
        "response_type": "code",
        "scope": "openid profile email",
        "state": "8945bkgfs",
        "client_id": "$_client",
      });
      final result = await FlutterWebAuth2.authenticate(url: authUri.toString(), callbackUrlScheme: 'tagpeak-scheme', options: const FlutterWebAuth2Options());
      final code = Uri.parse(result).queryParameters['code'];
      Log.info("Code: $code");
      final Map<String, String> model = {
        "redirect_uri": "tagpeak-scheme://login-callback",
        "grant_type": "authorization_code",
        "code": code.toString(),
        "client_id": "$_client",
        "client_secret":"$_clientSecret",
        "scope":"openid"
      };

      return login(body: model);
    } catch (e) {

      throw Exception([e]);
    }
  }

  Future<http.Response> verifyEmail({required String email}) async {
    try {
      final url = Uri.parse('$_ipKeycloakAddress$_userManagement/api/users/availability?email=$email');
      return await http.get(url);
    } catch (e) {
      throw Exception([e]);
    }
  }

  Future<void> logoutSocials({required String idToken, required String postLogoutRedirectUri}) async {
    try {
      final logoutUrl = Uri.parse('$_ipKeycloakAddress/realms/$_realm/protocol/openid-connect/logout');
      final uriWithParams = logoutUrl.replace(
        queryParameters: {
          'id_token_hint': idToken,
          'post_logout_redirect_uri': postLogoutRedirectUri,
        },
      );

      await http.get(uriWithParams);
    } catch (e) {
      Log.warning("Error during social logout: $e");
      throw Exception([e]);
    }
  }
}
