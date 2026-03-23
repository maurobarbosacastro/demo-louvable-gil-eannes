import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/shared/services/config_service.dart';

class ConfigProvider {
  final ConfigService service;
  final StateController<Object?> configNotifier;

  ConfigProvider({required this.service, required this.configNotifier});

  loadConfigurations() async {
    final configs = await service.loadConfigurations();

    configNotifier.update((_) => configs);
  }
}

final configsProvider = StateProvider<Object?>((ref) => null);

final configClassProvider = Provider(
      (ref) => ConfigProvider(
    service: ref.watch(configService),
    configNotifier: ref.watch(configsProvider.notifier)
  ),
);
