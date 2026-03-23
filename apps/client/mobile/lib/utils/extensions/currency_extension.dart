import 'package:intl/intl.dart';

extension CurrencyExtension on double {
  String toCurrency(String? currency) {
    final Map<String, String> currencies = {
      'USD': '\$',
      'EUR': '€'
    };

    if (currency != null && currencies.containsKey(currency)) {
      final symbol = currencies[currency];

      return NumberFormat.currency(
        locale: 'en_US',
        symbol: symbol,
        decimalDigits: 2
      ).format(this);
    }

    return NumberFormat.currency(
      locale: 'en_US',
      symbol: '€',
        decimalDigits: 2
    ).format(this);
  }
}
