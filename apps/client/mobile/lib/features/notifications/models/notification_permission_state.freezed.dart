// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'notification_permission_state.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$NotificationPermissionState {
  bool get isPermissionGranted => throw _privateConstructorUsedError;
  bool get hasRequestedPermission => throw _privateConstructorUsedError;
  bool get isTokenSaved => throw _privateConstructorUsedError;
  String? get fcmToken => throw _privateConstructorUsedError;
  String? get apnToken => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $NotificationPermissionStateCopyWith<NotificationPermissionState>
      get copyWith => throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $NotificationPermissionStateCopyWith<$Res> {
  factory $NotificationPermissionStateCopyWith(
          NotificationPermissionState value,
          $Res Function(NotificationPermissionState) then) =
      _$NotificationPermissionStateCopyWithImpl<$Res,
          NotificationPermissionState>;
  @useResult
  $Res call(
      {bool isPermissionGranted,
      bool hasRequestedPermission,
      bool isTokenSaved,
      String? fcmToken,
      String? apnToken});
}

/// @nodoc
class _$NotificationPermissionStateCopyWithImpl<$Res,
        $Val extends NotificationPermissionState>
    implements $NotificationPermissionStateCopyWith<$Res> {
  _$NotificationPermissionStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isPermissionGranted = null,
    Object? hasRequestedPermission = null,
    Object? isTokenSaved = null,
    Object? fcmToken = freezed,
    Object? apnToken = freezed,
  }) {
    return _then(_value.copyWith(
      isPermissionGranted: null == isPermissionGranted
          ? _value.isPermissionGranted
          : isPermissionGranted // ignore: cast_nullable_to_non_nullable
              as bool,
      hasRequestedPermission: null == hasRequestedPermission
          ? _value.hasRequestedPermission
          : hasRequestedPermission // ignore: cast_nullable_to_non_nullable
              as bool,
      isTokenSaved: null == isTokenSaved
          ? _value.isTokenSaved
          : isTokenSaved // ignore: cast_nullable_to_non_nullable
              as bool,
      fcmToken: freezed == fcmToken
          ? _value.fcmToken
          : fcmToken // ignore: cast_nullable_to_non_nullable
              as String?,
      apnToken: freezed == apnToken
          ? _value.apnToken
          : apnToken // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$NotificationPermissionStateImplCopyWith<$Res>
    implements $NotificationPermissionStateCopyWith<$Res> {
  factory _$$NotificationPermissionStateImplCopyWith(
          _$NotificationPermissionStateImpl value,
          $Res Function(_$NotificationPermissionStateImpl) then) =
      __$$NotificationPermissionStateImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {bool isPermissionGranted,
      bool hasRequestedPermission,
      bool isTokenSaved,
      String? fcmToken,
      String? apnToken});
}

/// @nodoc
class __$$NotificationPermissionStateImplCopyWithImpl<$Res>
    extends _$NotificationPermissionStateCopyWithImpl<$Res,
        _$NotificationPermissionStateImpl>
    implements _$$NotificationPermissionStateImplCopyWith<$Res> {
  __$$NotificationPermissionStateImplCopyWithImpl(
      _$NotificationPermissionStateImpl _value,
      $Res Function(_$NotificationPermissionStateImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isPermissionGranted = null,
    Object? hasRequestedPermission = null,
    Object? isTokenSaved = null,
    Object? fcmToken = freezed,
    Object? apnToken = freezed,
  }) {
    return _then(_$NotificationPermissionStateImpl(
      isPermissionGranted: null == isPermissionGranted
          ? _value.isPermissionGranted
          : isPermissionGranted // ignore: cast_nullable_to_non_nullable
              as bool,
      hasRequestedPermission: null == hasRequestedPermission
          ? _value.hasRequestedPermission
          : hasRequestedPermission // ignore: cast_nullable_to_non_nullable
              as bool,
      isTokenSaved: null == isTokenSaved
          ? _value.isTokenSaved
          : isTokenSaved // ignore: cast_nullable_to_non_nullable
              as bool,
      fcmToken: freezed == fcmToken
          ? _value.fcmToken
          : fcmToken // ignore: cast_nullable_to_non_nullable
              as String?,
      apnToken: freezed == apnToken
          ? _value.apnToken
          : apnToken // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc

class _$NotificationPermissionStateImpl
    implements _NotificationPermissionState {
  const _$NotificationPermissionStateImpl(
      {this.isPermissionGranted = false,
      this.hasRequestedPermission = false,
      this.isTokenSaved = false,
      this.fcmToken,
      this.apnToken});

  @override
  @JsonKey()
  final bool isPermissionGranted;
  @override
  @JsonKey()
  final bool hasRequestedPermission;
  @override
  @JsonKey()
  final bool isTokenSaved;
  @override
  final String? fcmToken;
  @override
  final String? apnToken;

  @override
  String toString() {
    return 'NotificationPermissionState(isPermissionGranted: $isPermissionGranted, hasRequestedPermission: $hasRequestedPermission, isTokenSaved: $isTokenSaved, fcmToken: $fcmToken, apnToken: $apnToken)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$NotificationPermissionStateImpl &&
            (identical(other.isPermissionGranted, isPermissionGranted) ||
                other.isPermissionGranted == isPermissionGranted) &&
            (identical(other.hasRequestedPermission, hasRequestedPermission) ||
                other.hasRequestedPermission == hasRequestedPermission) &&
            (identical(other.isTokenSaved, isTokenSaved) ||
                other.isTokenSaved == isTokenSaved) &&
            (identical(other.fcmToken, fcmToken) ||
                other.fcmToken == fcmToken) &&
            (identical(other.apnToken, apnToken) ||
                other.apnToken == apnToken));
  }

  @override
  int get hashCode => Object.hash(runtimeType, isPermissionGranted,
      hasRequestedPermission, isTokenSaved, fcmToken, apnToken);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$NotificationPermissionStateImplCopyWith<_$NotificationPermissionStateImpl>
      get copyWith => __$$NotificationPermissionStateImplCopyWithImpl<
          _$NotificationPermissionStateImpl>(this, _$identity);
}

abstract class _NotificationPermissionState
    implements NotificationPermissionState {
  const factory _NotificationPermissionState(
      {final bool isPermissionGranted,
      final bool hasRequestedPermission,
      final bool isTokenSaved,
      final String? fcmToken,
      final String? apnToken}) = _$NotificationPermissionStateImpl;

  @override
  bool get isPermissionGranted;
  @override
  bool get hasRequestedPermission;
  @override
  bool get isTokenSaved;
  @override
  String? get fcmToken;
  @override
  String? get apnToken;
  @override
  @JsonKey(ignore: true)
  _$$NotificationPermissionStateImplCopyWith<_$NotificationPermissionStateImpl>
      get copyWith => throw _privateConstructorUsedError;
}
