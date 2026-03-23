import 'dart:convert';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:http_interceptor/http/intercepted_client.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';

class ConfigService {

  late http.Client httpInterceptor;

  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  ConfigService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  Future<Object> loadConfigurations() async {
    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/configuration/latest'));

    if (response.statusCode != 200) {
      throw Exception('Error: loadConfigurations - Failed to load configs');
    }

    return jsonDecode(utf8.decode(response.bodyBytes));
  }
}

final configService = Provider(
  (ref) => ConfigService(ref.watch(httpInterceptorProvider)),
);
