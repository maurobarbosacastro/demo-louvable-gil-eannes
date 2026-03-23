import 'dart:convert';
import 'dart:io';

import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:http_interceptor/http/intercepted_client.dart';
import 'package:tagpeak/features/core/application/core_provider.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';
import 'package:tagpeak/shared/models/create_payment_method_model/create_payment_method_model.dart';
import 'package:tagpeak/shared/models/pagination_model/pagination_model.dart';
import 'package:tagpeak/shared/models/payment_method_model/payment_method_model.dart';
import 'package:tagpeak/shared/models/user_payment_method_model/user_payment_method_model.dart';

class PaymentMethodService {
  late http.Client httpInterceptor;
  final String _ipAddress = FlavorConfig.instance.variables["baseUrl"];

  PaymentMethodService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
  }

  Future<PaginationModel<UserPaymentMethodModel>> getPaymentMethods() async {
    final params = <String, String>{
      'number': '0',
      'size': '10',
      'sort': 'state desc',
    };

    final response = await httpInterceptor.get(
        Uri.parse('$_ipAddress/user/payment-method').replace(queryParameters: params));

    if (response.statusCode != 200) {
      throw Exception('Error: getPaymentMethods - Failed to get payment methods');
    }

    final pagination = PaginationModel<UserPaymentMethodModel>.fromJson(
      jsonDecode(response.body),
          (json) => UserPaymentMethodModel.fromJson(json),
    );

    return pagination;
  }

  Future<void> createPaymentMethod(CreatePaymentMethodModel body) async {
    final response = await httpInterceptor.post(
        Uri.parse('$_ipAddress/user/payment-method'), body: utf8.encode(jsonEncode(body)));

    if (response.statusCode != 201) {
      throw Exception('Error: createPaymentMethod - Failed to create payment method');
    }
  }

  Future<List<PaymentMethodModel>> getPaymentMethodsAvailable() async {
    final response = await httpInterceptor.get(Uri.parse('$_ipAddress/payment-method'));

    if (response.statusCode != 200) {
      throw Exception('Error: getPaymentMethodsAvailable - Failed to get payment methods');
    }

    final List<dynamic> jsonList = jsonDecode(response.body) as List<dynamic>;

    final paymentMethods = jsonList
        .map((json) => PaymentMethodModel.fromJson(json as Map<String, dynamic>))
        .toList();

    return paymentMethods;
  }

  Future<String> loadIbanStatement(File file) async {
    var uri = Uri.parse('$_ipAddress/user/payment-method/file');
    var request = http.MultipartRequest('POST', uri);

    request.files.add(await http.MultipartFile.fromPath('file', file.path));

    final streamedResponse = await httpInterceptor.send(request);

    final responseBody = await streamedResponse.stream.bytesToString();

    if (streamedResponse.statusCode == 200) {
      return responseBody;
    } else {
      throw Exception('Failed to upload file: ${streamedResponse.statusCode} - Response: $responseBody');
    }
  }

  Future<void> deletePaymentMethod(String uuid) async {
    final response = await httpInterceptor.delete(Uri.parse('$_ipAddress/user/payment-method/$uuid'));

    if (response.statusCode != 202) {
      throw Exception('Error: deletePaymentMethod - Failed to delete payment method with id $uuid');
    }
  }

  Future<PaymentMethodModel> getPaymentMethodById(String uuid) async {
    final response = await httpInterceptor.get(Uri.parse('$_ipAddress/user/payment-method/$uuid'));

    if (response.statusCode != 200) {
      throw Exception('Error: getPaymentMethodById - Failed to get payment method with id $uuid');
    }

    return PaymentMethodModel.fromJson(jsonDecode(response.body));
  }
}

final paymentMethodServiceProvider = Provider(
      (ref) => PaymentMethodService(ref.watch(httpInterceptorProvider)),
);
