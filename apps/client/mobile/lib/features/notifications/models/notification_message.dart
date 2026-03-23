import 'package:freezed_annotation/freezed_annotation.dart';

part 'notification_message.freezed.dart';
part 'notification_message.g.dart';

@freezed
class NotificationMessage with _$NotificationMessage {
  const factory NotificationMessage({
    required String id,
    required String title,
    required String body,
    @Default(5000) int duration,
    @Default(true) bool dismissible,
  }) = _NotificationMessage;

  factory NotificationMessage.fromJson(Map<String, Object?> json) =>
      _$NotificationMessageFromJson(json);
}