import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:logging/logging.dart';
import 'package:tagpeak/shared/theme/theme.dart';

import 'features/core/application/core_provider.dart';
import 'features/core/application/lifecycle.dart';
import 'features/core/routes/route_notifier.dart';
import 'utils/constants/strings.dart';
import 'utils/logging.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

Future<void> main() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);

  //Insipired by https://codewithandrea.com/articles/riverpod-initialize-listener-app-startup/#4-read-the-service-class-provider-with-providercontainer
  final ProviderContainer container = ProviderContainer();
  await Log.init();
  Log.setLevel(Level.OFF);
  await container.read(hiveProvider).init();

  FlavorConfig(name: "QA", color: Colors.red, location: BannerLocation.topStart, variables: {
    "keycloakUrl": "https://qa.keycloak.atlanse.ddns.net",
    "baseUrl": "https://qa.tagpeak.atlanse.ddns.net/tagpeak",
    "realm": "qa-tagpeak",
    "client": "tagpeak-mobile-client",
    "client_secret": "NJykM2HM1lwpvTDSxGUjTrdXBO8I3bJG",
    "backofficeUrl": "https://qa.tagpeak.backoffice.atlanse.ddns.net",
    "imageUrl": "https://qa.tagpeak.atlanse.ddns.net/images",
    "firebaseUrl": "https://qa.tagpeak.backoffice.atlanse.ddns.net/messaging",
    "vapidKey": "BHQhLgnyXVVGH1oRtgyNUi7cjPm36OGbJ8POo8x76kiODUmuEmzdRxOVUccr9RMqjd_TAPFvUSZ9knkxj6A7zSw"
  });

  runApp(
    UncontrolledProviderScope(
      // observers: [ProviderLogger()],
      container: container,
      child: const FlavorBanner(child: MyApp(/*theme: theme*/)),
    ),
  );
}

class MyApp extends ConsumerWidget {
  const MyApp({
    super.key,
    /*required this.theme*/
  });

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    SystemChrome.setPreferredOrientations([DeviceOrientation.portraitUp]);
    WidgetsFlutterBinding.ensureInitialized();

    final router = ref.watch(routerProvider);
    final scaffoldKey = ref.watch(scaffoldKeyProvider);

    SystemChrome.setSystemUIOverlayStyle(SystemUiOverlayStyle.dark);

    return Lifecycle(
      child: MaterialApp.router(
          routerConfig: router,
          title: Strings.appName,
          scaffoldMessengerKey: scaffoldKey,
          themeMode: ThemeMode.light,
          theme: TAppTheme.theme,
          debugShowCheckedModeBanner: false,
          localizationsDelegates: AppLocalizations.localizationsDelegates,
          supportedLocales: AppLocalizations.supportedLocales),
    );
  }
}
