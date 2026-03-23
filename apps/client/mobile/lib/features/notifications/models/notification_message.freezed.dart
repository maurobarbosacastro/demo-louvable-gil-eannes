// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'notification_message.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

NotificationMessage _$NotificationMessageFromJson(Map<String, dynamic> json) {
  return _NotificationMessage.fromJson(json);
}

/// @nodoc
mixin _$NotificationMessage {
  String get id => throw _privateConstructorUsedError;
  String get title => throw _privateConstructorUsedError;
  String get body => throw _privateConstructorUsedError;
  int get duration => throw _privateConstructorUsedError;
  bool get dismissible => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $NotificationMessageCopyWith<NotificationMessage> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $NotificationMessageCopyWith<$Res> {
  factory $NotificationMessageCopyWith(
          NotificationMessage value, $Res Function(NotificationMessage) then) =
      _$NotificationMessageCopyWithImpl<$Res, NotificationMessage>;
  @useResult
  $Res call(
      {String id, String title, String body, int duration, bool dismissible});
}

/// @nodoc
class _$NotificationMessageCopyWithImpl<$Res, $Val extends NotificationMessage>
    implements $NotificationMessageCopyWith<$Res> {
  _$NotificationMessageCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? title = null,
    Object? body = null,
    Object? duration = null,
    Object? dismissible = null,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as String,
      title: null == title
          ? _value.title
          : title // ignore: cast_nullable_to_non_nullable
              as String,
      body: null == body
          ? _value.body
          : body // ignore: cast_nullable_to_non_nullable
              as String,
      duration: null == duration
          ? _value.duration
          : duration // ignore: cast_nullable_to_non_nullable
              as int,
      dismissible: null == dismissible
          ? _value.dismissible
          : dismissible // ignore: cast_nullable_to_non_nullable
              as bool,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$NotificationMessageImplCopyWith<$Res>
    implements $NotificationMessageCopyWith<$Res> {
  factory _$$NotificationMessageImplCopyWith(_$NotificationMessageImpl value,
          $Res Function(_$NotificationMessageImpl) then) =
      __$$NotificationMessageImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String id, String title, String body, int duration, bool dismissible});
}

/// @nodoc
class __$$NotificationMessageImplCopyWithImpl<$Res>
    extends _$NotificationMessageCopyWithImpl<$Res, _$NotificationMessageImpl>
    implements _$$NotificationMessageImplCopyWith<$Res> {
  __$$NotificationMessageImplCopyWithImpl(_$NotificationMessageImpl _value,
      $Res Function(_$NotificationMessageImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? title = null,
    Object? body = null,
    Object? duration = null,
    Object? dismissible = null,
  }) {
    return _then(_$NotificationMessageImpl(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as String,
      title: null == title
          ? _value.title
          : title // ignore: cast_nullable_to_non_nullable
              as String,
      body: null == body
          ? _value.body
          : body // ignore: cast_nullable_to_non_nullable
              as String,
      duration: null == duration
          ? _value.duration
          : duration // ignore: cast_nullable_to_non_nullable
              as int,
      dismissible: null == dismissible
          ? _value.dismissible
          : dismissible // ignore: cast_nullable_to_non_nullable
              as bool,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$NotificationMessageImpl implements _NotificationMessage {
  const _$NotificationMessageImpl(
      {required this.id,
      required this.title,
      required this.body,
      this.duration = 5000,
      this.dismissible = true});

  factory _$NotificationMessageImpl.fromJson(Map<String, dynamic> json) =>
      _$$NotificationMessageImplFromJson(json);

  @override
  final String id;
  @override
  final String title;
  @override
  final String body;
  @override
  @JsonKey()
  final int duration;
  @override
  @JsonKey()
  final bool dismissible;

  @override
  String toString() {
    return 'NotificationMessage(id: $id, title: $title, body: $body, duration: $duration, dismissible: $dismissible)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$NotificationMessageImpl &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.title, title) || other.title == title) &&
            (identical(other.body, body) || other.body == body) &&
            (identical(other.duration, duration) ||
                other.duration == duration) &&
            (identical(other.dismissible, dismissible) ||
                other.dismissible == dismissible));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode =>
      Object.hash(runtimeType, id, title, body, duration, dismissible);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$NotificationMessageImplCopyWith<_$NotificationMessageImpl> get copyWith =>
      __$$NotificationMessageImplCopyWithImpl<_$NotificationMessageImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$NotificationMessageImplToJson(
      this,
    );
  }
}

abstract class _NotificationMessage implements NotificationMessage {
  const factory _NotificationMessage(
      {required final String id,
      required final String title,
      required final String body,
      final int duration,
      final bool dismissible}) = _$NotificationMessageImpl;

  factory _NotificationMessage.fromJson(Map<String, dynamic> json) =
      _$NotificationMessageImpl.fromJson;

  @override
  String get id;
  @override
  String get title;
  @override
  String get body;
  @override
  int get duration;
  @override
  bool get dismissible;
  @override
  @JsonKey(ignore: true)
  _$$NotificationMessageImplCopyWith<_$NotificationMessageImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
