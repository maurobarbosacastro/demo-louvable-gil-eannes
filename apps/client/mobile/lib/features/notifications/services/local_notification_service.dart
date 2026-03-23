import 'dart:io';
import 'package:flutter_local_notifications/flutter_local_notifications.dart';

enum NotificationPlatform {
  android,
  ios,
  both,
}

class LocalNotificationService {
  static final FlutterLocalNotificationsPlugin _notificationsPlugin =
      FlutterLocalNotificationsPlugin();

  static Future<void> initialize() async {
    const AndroidInitializationSettings initializationSettingsAndroid =
        AndroidInitializationSettings('@mipmap/launcher_icon');

    const DarwinInitializationSettings initializationSettingsDarwin =
        DarwinInitializationSettings(
      requestAlertPermission: true,
      requestBadgePermission: true,
      requestSoundPermission: true
    );

    const InitializationSettings initializationSettings =
        InitializationSettings(
      android: initializationSettingsAndroid,
      iOS: initializationSettingsDarwin,
    );

    await _notificationsPlugin.initialize(
      initializationSettings,
      onDidReceiveNotificationResponse: (NotificationResponse response) {
        print('Notification clicked: ${response.payload}');
      },
    );

    // Request permissions for Android 13+
    if (Platform.isAndroid) {
      await _notificationsPlugin
          .resolvePlatformSpecificImplementation<
              AndroidFlutterLocalNotificationsPlugin>()
          ?.requestNotificationsPermission();
    }
  }

  static Future<void> showNotification({
    required int id,
    String? title,
    String? body,
    String? payload,
    NotificationPlatform platform = NotificationPlatform.both,
  }) async {
    if (platform == NotificationPlatform.android && !Platform.isAndroid) {
      return;
    }

    if (platform == NotificationPlatform.ios && !Platform.isIOS) {
      return;
    }

    AndroidNotificationDetails? androidPlatformChannelSpecifics;
    DarwinNotificationDetails? darwinPlatformChannelSpecifics;

    if (Platform.isAndroid &&
        (platform == NotificationPlatform.android || platform == NotificationPlatform.both)) {
      androidPlatformChannelSpecifics = const AndroidNotificationDetails(
        'tagpeak_channel',
        'Tagpeak Notifications',
        channelDescription: 'General notifications for Tagpeak app',
        importance: Importance.high,
        priority: Priority.high,
        showWhen: true,
        icon: '@mipmap/launcher_icon',
      );
    }

    if (Platform.isIOS &&
        (platform == NotificationPlatform.ios || platform == NotificationPlatform.both)) {
      darwinPlatformChannelSpecifics = DarwinNotificationDetails(
        presentAlert: true,
        presentBadge: true,
        presentSound: true,
        subtitle: body,
        sound: 'default',
      );
    }

    final NotificationDetails platformChannelSpecifics = NotificationDetails(
      android: androidPlatformChannelSpecifics,
      iOS: darwinPlatformChannelSpecifics,
    );

    await _notificationsPlugin.show(
      id,
      title,
      body,
      platformChannelSpecifics,
      payload: payload,
    );
  }

  static Future<void> cancelNotification(int id) async {
    await _notificationsPlugin.cancel(id);
  }

  static Future<void> cancelAllNotifications() async {
    await _notificationsPlugin.cancelAll();
  }
}
