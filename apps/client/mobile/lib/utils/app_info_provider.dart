import 'package:flutter_riverpod/flutter_riverpod.dart';

final appInfoProvider = StateProvider<AppInfoModel>((ref) => AppInfoModel());

class AppInfoModel {
  late String version;
  late String packageName;
  late String buildNumber;
  late String appName;

  AppInfoModel();

  AppInfoModel.init({
    required this.version,
    required this.packageName,
    required this.buildNumber,
    required this.appName,
  });
}
