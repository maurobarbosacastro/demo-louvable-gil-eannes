import 'dart:io';

import 'package:tagpeak/shared/theme/theme.dart';
import 'package:tagpeak/utils/bad_certificate.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:logging/logging.dart';

import 'features/core/application/core_provider.dart';
import 'features/core/application/lifecycle.dart';
import 'features/core/routes/route_notifier.dart';
import 'utils/constants/strings.dart';
import 'utils/logging.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

Future<void> main() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);
  HttpOverrides.global = BadCertificate();
  //Insipired by https://codewithandrea.com/articles/riverpod-initialize-listener-app-startup/#4-read-the-service-class-provider-with-providercontainer
  final ProviderContainer container = ProviderContainer();
  await Log.init();
  Log.setLevel(Level.ALL);
  await container.read(hiveProvider).init();

  FlavorConfig(name: "PRE", color: Colors.blue, location: BannerLocation.topStart, variables: {
    "keycloakUrl": "https://pre.keycloak.atlanse-cloud.ddns.net",
    "baseUrl": "https://pre.tagpeak.atlanse-cloud.ddns.net/tagpeak",
    "realm": "pre-tagpeak",
    "client": "tagpeak-mobile-client",
    "client_secret": "IB4NdkGKmP0fOeZ8lxu3W7b2Pf7Cy8xT",
    "backofficeUrl": "https://pre.tagpeak.backoffice.atlanse-cloud.ddns.net",
    "imageUrl": "https://pre.tagpeak.atlanse-cloud.ddns.net/images",
    "firebaseUrl": "https://pre.tagpeak.backoffice.atlanse-cloud.ddns.net/messaging",
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

    final router = ref.watch(routerProvider);
    final scaffoldKey = ref.watch(scaffoldKeyProvider);

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
