class Currency {
  final String key;
  final String label;
  final String icon;
  final String symbol;

  const Currency({
    required this.key,
    required this.label,
    required this.icon,
    required this.symbol,
  });
}

const List<Currency> currencies = [
  Currency(key: 'EUR', label: 'EUR', icon: 'EUR', symbol: '€'),
  Currency(key: 'GBP', label: 'GBP', icon: 'GBP', symbol: '£'),
  Currency(key: 'CAD', label: 'CAD', icon: 'CAD', symbol: 'CA\$'),
  Currency(key: 'AUD', label: 'AUD', icon: 'AUD', symbol: 'A\$'),
  Currency(key: 'USD', label: 'USD', icon: 'USD', symbol: 'US\$'),
  Currency(key: 'BRL', label: 'BRL', icon: 'BRL', symbol: 'R\$'),
  Currency(key: 'MXN', label: 'MXN', icon: 'MXN', symbol: 'MX\$'),
];
