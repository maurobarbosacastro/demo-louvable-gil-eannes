import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import '../../core/application/core_provider.dart';
import '../services/fcm_service.dart';
import '../models/notification_permission_state.dart';

final fcmServiceProvider = Provider<FCMService>((ref) => FCMService(ref.watch(httpInterceptorProvider)));

final notificationPermissionProvider = StateNotifierProvider<NotificationPermissionNotifier, NotificationPermissionState>(
  (ref) => NotificationPermissionNotifier(),
);

class NotificationPermissionNotifier extends StateNotifier<NotificationPermissionState> {
  NotificationPermissionNotifier() : super(const NotificationPermissionState());

  Future<void> requestPermissionAndSaveToken(String userUuid, FCMService fcmService, ref) async {
    if (state.hasRequestedPermission) {
      return;
    }

    state = state.copyWith(hasRequestedPermission: true);

    try {
      String? token = await fcmService.requestPermissionAndGetToken(userUuid);

      if (token != null) {
        state = state.copyWith(
          isPermissionGranted: true,
          isTokenSaved: true,
          fcmToken: token,
        );

        // Setup foreground message handler
        fcmService.setupForegroundMessageHandler();

        // Setup token refresh handler
        fcmService.setupTokenRefreshHandler();
      }
    } catch (e) {
      print('FCM provider error: $e');

      state = state.copyWith(
        isPermissionGranted: false,
        isTokenSaved: false,
      );

      // TODO: Handle different types of errors appropriately
      // TODO: Show user-friendly error messages
      // TODO: Implement retry logic for network failures
    }
  }

  void resetPermissionState() {
    state = const NotificationPermissionState();
  }
}

// Provider to get user info for FCM setup (read-only, no side effects)
final fcmUserInfoProvider = Provider<Map<String, dynamic>>((ref) {
  final user = ref.watch(userProvider);

  return {
    'user': user,
    'isClient': user != null && user.roles.contains('default'),
  };
});

// Function to manually trigger FCM permission request (called from widgets)
void triggerFCMPermissionRequest(WidgetRef ref) {
  final userInfo = ref.read(fcmUserInfoProvider);
  final user = userInfo['user'];
  final isClient = userInfo['isClient'] as bool;

  if (isClient && user != null) {
    print('FCM permission request for user: ${user.uuid}');
    final fcmService = ref.read(fcmServiceProvider);
    final permissionNotifier = ref.read(notificationPermissionProvider.notifier);

    // Delay the execution to avoid provider initialization issues
    Future.microtask(() {
      permissionNotifier.requestPermissionAndSaveToken(user.uuid, fcmService, ref);
    });
  }
}
