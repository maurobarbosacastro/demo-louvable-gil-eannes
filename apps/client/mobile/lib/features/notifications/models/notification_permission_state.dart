import 'package:freezed_annotation/freezed_annotation.dart';

part 'notification_permission_state.freezed.dart';

@freezed
class NotificationPermissionState with _$NotificationPermissionState {
  const factory NotificationPermissionState({
    @Default(false) bool isPermissionGranted,
    @Default(false) bool hasRequestedPermission,
    @Default(false) bool isTokenSaved,
    String? fcmToken,
    String? apnToken,
  }) = _NotificationPermissionState;
}