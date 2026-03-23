import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:logging/logging.dart';
import 'package:tagpeak/shared/theme/theme.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:tagpeak/features/notifications/services/local_notification_service.dart';
import 'firebase_options.dart';

import 'features/core/application/core_provider.dart';
import 'features/core/application/lifecycle.dart';
import 'features/core/routes/route_notifier.dart';
import 'utils/constants/strings.dart';
import 'utils/logging.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

@pragma('vm:entry-point')
Future<void> _firebaseMessagingBackgroundHandler(RemoteMessage message) async {
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  print('FCM background message: ${message.messageId}');
}

Future<void> main() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);

  //Insipired by https://codewithandrea.com/articles/riverpod-initialize-listener-app-startup/#4-read-the-service-class-provider-with-providercontainer
  final ProviderContainer container = ProviderContainer();
  await Log.init();
  Log.setLevel(Level.OFF);
  await container.read(hiveProvider).init();

  FlavorConfig(variables: {
    "keycloakUrl": "https://sheepdog-living-treefrog.ngrok-free.app/keycloak",
    "baseUrl": "https://sheepdog-living-treefrog.ngrok-free.app/ms-tagpeak",
    "realm": "tagpeak",
    "client": "tagpeak-mobile-client",
    "client_secret": "k0PyXfKIbdvCpH9Pwoyp9GvKFXe3xp4x",
    "backofficeUrl": "https://sheepdog-living-treefrog.ngrok-free.app",
    "imageUrl": "https://sheepdog-living-treefrog.ngrok-free.app/ms-image",
    "firebaseUrl": "https://sheepdog-living-treefrog.ngrok-free.app/ms-firebase",
    "vapidKey": "BHQhLgnyXVVGH1oRtgyNUi7cjPm36OGbJ8POo8x76kiODUmuEmzdRxOVUccr9RMqjd_TAPFvUSZ9knkxj6A7zSw"
  });

  try {
    // Initialize Firebase with options
    await Firebase.initializeApp(
      options: DefaultFirebaseOptions.currentPlatform,
    );
    print('Firebase initialized: ${DefaultFirebaseOptions.currentPlatform.projectId}');

    // Initialize local notifications
    await LocalNotificationService.initialize();
    print('Local notifications initialized');

    // Background message handler
    FirebaseMessaging.onBackgroundMessage(_firebaseMessagingBackgroundHandler);
  } catch (e) {
    print('Firebase initialization failed: $e');
    rethrow;
  }

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
