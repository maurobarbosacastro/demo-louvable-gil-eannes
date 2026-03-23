import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'logging.dart';

class ProviderLogger extends ProviderObserver {
  @override
  void didUpdateProvider(
    ProviderBase provider,
    Object? previousValue,
    Object? newValue,
    ProviderContainer container,
  ) {
    Log.info('''
    {
      "provider": "${provider.name ?? provider.runtimeType}",
      "newValue": "$newValue"
    }''');
  }

  @override
  void didAddProvider(
    ProviderBase provider,
    Object? value,
    ProviderContainer container,
  ) {
    Log.config('''
    {
      "provider": "${provider.name ?? provider.runtimeType}",
      "value": "$value"
    }''');
  }

  @override
  void didDisposeProvider(
    ProviderBase provider,
    ProviderContainer container,
  ) {
    Log.info('''
    {
      "provider": "${provider.name ?? provider.runtimeType}"
    }''');
  }

  @override
  void providerDidFail(
    ProviderBase provider,
    Object error,
    StackTrace stackTrace,
    ProviderContainer container,
  ) {
    Log.severe('''
    {
      "provider": "${provider.name ?? provider.runtimeType}"
    }''');
  }
}
