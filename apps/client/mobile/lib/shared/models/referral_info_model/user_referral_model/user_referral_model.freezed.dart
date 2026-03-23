// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'user_referral_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

UserReferralModel _$UserReferralModelFromJson(Map<String, dynamic> json) {
  return _UserReferralModel.fromJson(json);
}

/// @nodoc
mixin _$UserReferralModel {
  String get uuid => throw _privateConstructorUsedError;
  String get firstName => throw _privateConstructorUsedError;
  String get lastName => throw _privateConstructorUsedError;
  String get profilePicture => throw _privateConstructorUsedError;
  double get referredValue => throw _privateConstructorUsedError;
  bool get firstTransactionSuccessful => throw _privateConstructorUsedError;
  String? get displayName => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $UserReferralModelCopyWith<UserReferralModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $UserReferralModelCopyWith<$Res> {
  factory $UserReferralModelCopyWith(
          UserReferralModel value, $Res Function(UserReferralModel) then) =
      _$UserReferralModelCopyWithImpl<$Res, UserReferralModel>;
  @useResult
  $Res call(
      {String uuid,
      String firstName,
      String lastName,
      String profilePicture,
      double referredValue,
      bool firstTransactionSuccessful,
      String? displayName});
}

/// @nodoc
class _$UserReferralModelCopyWithImpl<$Res, $Val extends UserReferralModel>
    implements $UserReferralModelCopyWith<$Res> {
  _$UserReferralModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? firstName = null,
    Object? lastName = null,
    Object? profilePicture = null,
    Object? referredValue = null,
    Object? firstTransactionSuccessful = null,
    Object? displayName = freezed,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      firstName: null == firstName
          ? _value.firstName
          : firstName // ignore: cast_nullable_to_non_nullable
              as String,
      lastName: null == lastName
          ? _value.lastName
          : lastName // ignore: cast_nullable_to_non_nullable
              as String,
      profilePicture: null == profilePicture
          ? _value.profilePicture
          : profilePicture // ignore: cast_nullable_to_non_nullable
              as String,
      referredValue: null == referredValue
          ? _value.referredValue
          : referredValue // ignore: cast_nullable_to_non_nullable
              as double,
      firstTransactionSuccessful: null == firstTransactionSuccessful
          ? _value.firstTransactionSuccessful
          : firstTransactionSuccessful // ignore: cast_nullable_to_non_nullable
              as bool,
      displayName: freezed == displayName
          ? _value.displayName
          : displayName // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$UserReferralModelImplCopyWith<$Res>
    implements $UserReferralModelCopyWith<$Res> {
  factory _$$UserReferralModelImplCopyWith(_$UserReferralModelImpl value,
          $Res Function(_$UserReferralModelImpl) then) =
      __$$UserReferralModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      String firstName,
      String lastName,
      String profilePicture,
      double referredValue,
      bool firstTransactionSuccessful,
      String? displayName});
}

/// @nodoc
class __$$UserReferralModelImplCopyWithImpl<$Res>
    extends _$UserReferralModelCopyWithImpl<$Res, _$UserReferralModelImpl>
    implements _$$UserReferralModelImplCopyWith<$Res> {
  __$$UserReferralModelImplCopyWithImpl(_$UserReferralModelImpl _value,
      $Res Function(_$UserReferralModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? firstName = null,
    Object? lastName = null,
    Object? profilePicture = null,
    Object? referredValue = null,
    Object? firstTransactionSuccessful = null,
    Object? displayName = freezed,
  }) {
    return _then(_$UserReferralModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      firstName: null == firstName
          ? _value.firstName
          : firstName // ignore: cast_nullable_to_non_nullable
              as String,
      lastName: null == lastName
          ? _value.lastName
          : lastName // ignore: cast_nullable_to_non_nullable
              as String,
      profilePicture: null == profilePicture
          ? _value.profilePicture
          : profilePicture // ignore: cast_nullable_to_non_nullable
              as String,
      referredValue: null == referredValue
          ? _value.referredValue
          : referredValue // ignore: cast_nullable_to_non_nullable
              as double,
      firstTransactionSuccessful: null == firstTransactionSuccessful
          ? _value.firstTransactionSuccessful
          : firstTransactionSuccessful // ignore: cast_nullable_to_non_nullable
              as bool,
      displayName: freezed == displayName
          ? _value.displayName
          : displayName // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$UserReferralModelImpl implements _UserReferralModel {
  const _$UserReferralModelImpl(
      {required this.uuid,
      required this.firstName,
      required this.lastName,
      required this.profilePicture,
      required this.referredValue,
      required this.firstTransactionSuccessful,
      this.displayName});

  factory _$UserReferralModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$UserReferralModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String firstName;
  @override
  final String lastName;
  @override
  final String profilePicture;
  @override
  final double referredValue;
  @override
  final bool firstTransactionSuccessful;
  @override
  final String? displayName;

  @override
  String toString() {
    return 'UserReferralModel(uuid: $uuid, firstName: $firstName, lastName: $lastName, profilePicture: $profilePicture, referredValue: $referredValue, firstTransactionSuccessful: $firstTransactionSuccessful, displayName: $displayName)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$UserReferralModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.firstName, firstName) ||
                other.firstName == firstName) &&
            (identical(other.lastName, lastName) ||
                other.lastName == lastName) &&
            (identical(other.profilePicture, profilePicture) ||
                other.profilePicture == profilePicture) &&
            (identical(other.referredValue, referredValue) ||
                other.referredValue == referredValue) &&
            (identical(other.firstTransactionSuccessful,
                    firstTransactionSuccessful) ||
                other.firstTransactionSuccessful ==
                    firstTransactionSuccessful) &&
            (identical(other.displayName, displayName) ||
                other.displayName == displayName));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, uuid, firstName, lastName,
      profilePicture, referredValue, firstTransactionSuccessful, displayName);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$UserReferralModelImplCopyWith<_$UserReferralModelImpl> get copyWith =>
      __$$UserReferralModelImplCopyWithImpl<_$UserReferralModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$UserReferralModelImplToJson(
      this,
    );
  }
}

abstract class _UserReferralModel implements UserReferralModel {
  const factory _UserReferralModel(
      {required final String uuid,
      required final String firstName,
      required final String lastName,
      required final String profilePicture,
      required final double referredValue,
      required final bool firstTransactionSuccessful,
      final String? displayName}) = _$UserReferralModelImpl;

  factory _UserReferralModel.fromJson(Map<String, dynamic> json) =
      _$UserReferralModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get firstName;
  @override
  String get lastName;
  @override
  String get profilePicture;
  @override
  double get referredValue;
  @override
  bool get firstTransactionSuccessful;
  @override
  String? get displayName;
  @override
  @JsonKey(ignore: true)
  _$$UserReferralModelImplCopyWith<_$UserReferralModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
