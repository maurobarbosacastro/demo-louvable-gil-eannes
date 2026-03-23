import 'dart:io';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter_flavor/flutter_flavor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:http_interceptor/http/intercepted_client.dart';
import 'dart:convert';
import 'package:tagpeak/features/notifications/services/local_notification_service.dart';
import 'package:tagpeak/features/core/application/http_interceptor.dart';

import '../../core/application/core_provider.dart';


// For Firebase ios, I had to add the following for the build settings: https://github.com/firebase/flutterfire/issues/13323#issuecomment-2355013360

class FCMService {
  late http.Client httpInterceptor;
  final FirebaseMessaging _firebaseMessaging = FirebaseMessaging.instance;
  final String _baseUrl = FlavorConfig.instance.variables["baseUrl"];

  FCMService(ApiInterceptor interceptor) {
    httpInterceptor = InterceptedClient.build(interceptors: [interceptor]);
    }

  Future<String?> requestPermissionAndGetToken(String userUuid) async {
    try {
      // Request permission
      NotificationSettings settings = await _firebaseMessaging.requestPermission(
        alert: true,
        badge: true,
        sound: true,
        announcement: false,
        carPlay: false,
        criticalAlert: false,
        provisional: false,
      );

      if (settings.authorizationStatus != AuthorizationStatus.authorized) {
        print('FCM permission denied');
        throw Exception('Permission denied');
      }

      // Get FCM token
      String? fcmToken = await _firebaseMessaging.getToken();

      if (fcmToken == null) {
        throw Exception('Failed to get FCM token');
      }

      print('FCM token obtained for user: $userUuid');

      // iOS-specific: Verify APN token
      if (Platform.isIOS) {
        String? apnToken = await _firebaseMessaging.getAPNSToken();
        if (apnToken == null) {
          throw Exception('APN token not available');
        }
      }

      // Save token to backend
      bool success = await _saveTokenToBackend(userUuid, fcmToken);
      if (!success) {
        throw Exception('Failed to save token to backend');
      }

      print('FCM token saved to backend');
      return fcmToken;
    } catch (e) {
      print('FCM error: $e');
      // TODO: Implement comprehensive error handling for permission denial
      // TODO: Add retry mechanisms and user-friendly error messages
      // TODO: Implement persistent permission state management
      // TODO: Add proper network failure handling
      rethrow;
    }
  }

  Future<bool> _saveTokenToBackend(String userUuid, String token) async {
    try {
      final response = await httpInterceptor.post(
        Uri.parse('$_baseUrl/notifications/token'),
        headers: {
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          'userUuid': userUuid,
          'token': token,
        }),
      );

      if (response.statusCode == 200 || response.statusCode == 201) {
        return true;
      } else {
        print('FCM backend error: ${response.statusCode} - ${response.body}');
        if (response.statusCode == 409) {
          return true;
        }
        return false;
      }
    } catch (e) {
      print('FCM network error: $e');
      return false;
    }
  }

  Future<bool> clearUserToken(String userUuid) async {
    try {
      final response = await httpInterceptor.delete(
        Uri.parse('$_baseUrl/notifications/token'),
      );

      if (response.statusCode == 200) {
        return true;
      } else {
        print('FCM clear token error: ${response.statusCode}');
        return false;
      }
    } catch (e) {
      print('FCM clear token network error: $e');
      return false;
    }
  }

  void setupForegroundMessageHandler() {
    print('setupForegroundMessageHandler');
    FirebaseMessaging.onMessage.listen((RemoteMessage message) {
      print('FCM message received: ${message.notification?.title ?? message.messageId}');

      // Show local notification when app is in foreground
      if (message.notification != null) {
        LocalNotificationService.showNotification(
          id: message.messageId?.hashCode ?? 0,
          title: message.notification!.title,
          body: message.notification!.body,
          payload: message.data.toString(),
        );
      }

      // TODO: Add notification payload processing for app-specific data
    });
  }

  /// Listen for token refresh events
  void setupTokenRefreshHandler() {
    _firebaseMessaging.onTokenRefresh.listen((String token) {
      print('FCM token refreshed');
      // TODO: Implement token refresh handling
      // TODO: Update backend with new token automatically
    });
  }
}
