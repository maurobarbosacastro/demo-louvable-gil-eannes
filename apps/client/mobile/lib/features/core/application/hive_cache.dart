import 'package:hive_flutter/hive_flutter.dart';

import '../../../utils/constants/strings.dart';
import '../../../utils/logging.dart';

class HiveDatabase {
  late Box<String> _tokenBox;
  late Box<String> _refreshTokenBox;
  late Box<String> _idTokenBox;
  late Box<String> _loginMethodBox;
  late Box<bool> _biometricsBox;
  late Box<dynamic> _appConfig;

  Box<String>? get tokenBox => _hasBeenInitialized ? _tokenBox : null;
  Box<String>? get refreshTokenBox => _hasBeenInitialized ? _refreshTokenBox : null;
  Box<String>? get idTokenBox => _hasBeenInitialized ? _idTokenBox : null;
  Box<String>? get loginMethodBox => _hasBeenInitialized ? _loginMethodBox : null;
  Box<bool>? get biometricsBox => _hasBeenInitialized ? _biometricsBox : null;

  Box<dynamic>? get configBox => _hasBeenInitialized ? _appConfig : null;

  bool _hasBeenInitialized = false;

  bool get hasBeenInitialized => _hasBeenInitialized;

  String boxName(String box) => '${Strings.appCodeName}.$box';

  Future<void> init() async {
    Log.info('Init Hive START');
    if (_hasBeenInitialized) return;

    _hasBeenInitialized = true;

    await Hive.initFlutter();
    await openBoxes();
    Log.info('Init Hive END');
  }

  Future<void> openBoxes() async {
    _tokenBox = await Hive.openBox<String>(boxName(Strings.token));
    _refreshTokenBox = await Hive.openBox<String>(boxName(Strings.refreshToken));
    _idTokenBox = await Hive.openBox<String>(boxName(Strings.idToken));
    _loginMethodBox = await Hive.openBox<String>(boxName(Strings.loginMethod));
    _biometricsBox = await Hive.openBox<bool>(boxName(Strings.biometrics));
    _appConfig = await Hive.openBox<dynamic>(boxName(Strings.config));
  }

  Future<void> closeBoxes() async {
    await tokenBox?.close();
    await refreshTokenBox?.close();
    await idTokenBox?.close();
    await loginMethodBox?.close();
    await configBox?.close();
  }

  String? get token => tokenBox?.get('token');
  String? get refreshToken => refreshTokenBox?.get('refresh');
  String? get idToken => idTokenBox?.get('idToken');
  String? get loginMethod => loginMethodBox?.get('loginMethod');
  bool? get biometrics => biometricsBox?.get('biometrics');
}
