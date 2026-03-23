// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'user_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

UserModel _$UserModelFromJson(Map<String, dynamic> json) {
  return _UserModel.fromJson(json);
}

/// @nodoc
mixin _$UserModel {
  String get uuid => throw _privateConstructorUsedError;
  String get email => throw _privateConstructorUsedError;
  String get country => throw _privateConstructorUsedError;
  String get referralCode => throw _privateConstructorUsedError;
  double get balance => throw _privateConstructorUsedError;
  int get createdAt => throw _privateConstructorUsedError;
  String get firstName => throw _privateConstructorUsedError;
  String get lastName => throw _privateConstructorUsedError;
  String get currency => throw _privateConstructorUsedError;
  String get displayName => throw _privateConstructorUsedError;
  String get birthDate => throw _privateConstructorUsedError;
  List<String> get groups => throw _privateConstructorUsedError;
  bool get isVerified => throw _privateConstructorUsedError;
  bool get onboardingFinished => throw _privateConstructorUsedError;
  String? get profilePicture => throw _privateConstructorUsedError;
  double? get transactionPercentage => throw _privateConstructorUsedError;
  double? get rewardPercentage => throw _privateConstructorUsedError;
  bool get newsletter => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $UserModelCopyWith<UserModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $UserModelCopyWith<$Res> {
  factory $UserModelCopyWith(UserModel value, $Res Function(UserModel) then) =
      _$UserModelCopyWithImpl<$Res, UserModel>;
  @useResult
  $Res call(
      {String uuid,
      String email,
      String country,
      String referralCode,
      double balance,
      int createdAt,
      String firstName,
      String lastName,
      String currency,
      String displayName,
      String birthDate,
      List<String> groups,
      bool isVerified,
      bool onboardingFinished,
      String? profilePicture,
      double? transactionPercentage,
      double? rewardPercentage,
      bool newsletter});
}

/// @nodoc
class _$UserModelCopyWithImpl<$Res, $Val extends UserModel>
    implements $UserModelCopyWith<$Res> {
  _$UserModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? email = null,
    Object? country = null,
    Object? referralCode = null,
    Object? balance = null,
    Object? createdAt = null,
    Object? firstName = null,
    Object? lastName = null,
    Object? currency = null,
    Object? displayName = null,
    Object? birthDate = null,
    Object? groups = null,
    Object? isVerified = null,
    Object? onboardingFinished = null,
    Object? profilePicture = freezed,
    Object? transactionPercentage = freezed,
    Object? rewardPercentage = freezed,
    Object? newsletter = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      email: null == email
          ? _value.email
          : email // ignore: cast_nullable_to_non_nullable
              as String,
      country: null == country
          ? _value.country
          : country // ignore: cast_nullable_to_non_nullable
              as String,
      referralCode: null == referralCode
          ? _value.referralCode
          : referralCode // ignore: cast_nullable_to_non_nullable
              as String,
      balance: null == balance
          ? _value.balance
          : balance // ignore: cast_nullable_to_non_nullable
              as double,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as int,
      firstName: null == firstName
          ? _value.firstName
          : firstName // ignore: cast_nullable_to_non_nullable
              as String,
      lastName: null == lastName
          ? _value.lastName
          : lastName // ignore: cast_nullable_to_non_nullable
              as String,
      currency: null == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String,
      displayName: null == displayName
          ? _value.displayName
          : displayName // ignore: cast_nullable_to_non_nullable
              as String,
      birthDate: null == birthDate
          ? _value.birthDate
          : birthDate // ignore: cast_nullable_to_non_nullable
              as String,
      groups: null == groups
          ? _value.groups
          : groups // ignore: cast_nullable_to_non_nullable
              as List<String>,
      isVerified: null == isVerified
          ? _value.isVerified
          : isVerified // ignore: cast_nullable_to_non_nullable
              as bool,
      onboardingFinished: null == onboardingFinished
          ? _value.onboardingFinished
          : onboardingFinished // ignore: cast_nullable_to_non_nullable
              as bool,
      profilePicture: freezed == profilePicture
          ? _value.profilePicture
          : profilePicture // ignore: cast_nullable_to_non_nullable
              as String?,
      transactionPercentage: freezed == transactionPercentage
          ? _value.transactionPercentage
          : transactionPercentage // ignore: cast_nullable_to_non_nullable
              as double?,
      rewardPercentage: freezed == rewardPercentage
          ? _value.rewardPercentage
          : rewardPercentage // ignore: cast_nullable_to_non_nullable
              as double?,
      newsletter: null == newsletter
          ? _value.newsletter
          : newsletter // ignore: cast_nullable_to_non_nullable
              as bool,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$UserModelImplCopyWith<$Res>
    implements $UserModelCopyWith<$Res> {
  factory _$$UserModelImplCopyWith(
          _$UserModelImpl value, $Res Function(_$UserModelImpl) then) =
      __$$UserModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      String email,
      String country,
      String referralCode,
      double balance,
      int createdAt,
      String firstName,
      String lastName,
      String currency,
      String displayName,
      String birthDate,
      List<String> groups,
      bool isVerified,
      bool onboardingFinished,
      String? profilePicture,
      double? transactionPercentage,
      double? rewardPercentage,
      bool newsletter});
}

/// @nodoc
class __$$UserModelImplCopyWithImpl<$Res>
    extends _$UserModelCopyWithImpl<$Res, _$UserModelImpl>
    implements _$$UserModelImplCopyWith<$Res> {
  __$$UserModelImplCopyWithImpl(
      _$UserModelImpl _value, $Res Function(_$UserModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? email = null,
    Object? country = null,
    Object? referralCode = null,
    Object? balance = null,
    Object? createdAt = null,
    Object? firstName = null,
    Object? lastName = null,
    Object? currency = null,
    Object? displayName = null,
    Object? birthDate = null,
    Object? groups = null,
    Object? isVerified = null,
    Object? onboardingFinished = null,
    Object? profilePicture = freezed,
    Object? transactionPercentage = freezed,
    Object? rewardPercentage = freezed,
    Object? newsletter = null,
  }) {
    return _then(_$UserModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      email: null == email
          ? _value.email
          : email // ignore: cast_nullable_to_non_nullable
              as String,
      country: null == country
          ? _value.country
          : country // ignore: cast_nullable_to_non_nullable
              as String,
      referralCode: null == referralCode
          ? _value.referralCode
          : referralCode // ignore: cast_nullable_to_non_nullable
              as String,
      balance: null == balance
          ? _value.balance
          : balance // ignore: cast_nullable_to_non_nullable
              as double,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as int,
      firstName: null == firstName
          ? _value.firstName
          : firstName // ignore: cast_nullable_to_non_nullable
              as String,
      lastName: null == lastName
          ? _value.lastName
          : lastName // ignore: cast_nullable_to_non_nullable
              as String,
      currency: null == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String,
      displayName: null == displayName
          ? _value.displayName
          : displayName // ignore: cast_nullable_to_non_nullable
              as String,
      birthDate: null == birthDate
          ? _value.birthDate
          : birthDate // ignore: cast_nullable_to_non_nullable
              as String,
      groups: null == groups
          ? _value._groups
          : groups // ignore: cast_nullable_to_non_nullable
              as List<String>,
      isVerified: null == isVerified
          ? _value.isVerified
          : isVerified // ignore: cast_nullable_to_non_nullable
              as bool,
      onboardingFinished: null == onboardingFinished
          ? _value.onboardingFinished
          : onboardingFinished // ignore: cast_nullable_to_non_nullable
              as bool,
      profilePicture: freezed == profilePicture
          ? _value.profilePicture
          : profilePicture // ignore: cast_nullable_to_non_nullable
              as String?,
      transactionPercentage: freezed == transactionPercentage
          ? _value.transactionPercentage
          : transactionPercentage // ignore: cast_nullable_to_non_nullable
              as double?,
      rewardPercentage: freezed == rewardPercentage
          ? _value.rewardPercentage
          : rewardPercentage // ignore: cast_nullable_to_non_nullable
              as double?,
      newsletter: null == newsletter
          ? _value.newsletter
          : newsletter // ignore: cast_nullable_to_non_nullable
              as bool,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$UserModelImpl implements _UserModel {
  const _$UserModelImpl(
      {required this.uuid,
      required this.email,
      required this.country,
      required this.referralCode,
      required this.balance,
      required this.createdAt,
      required this.firstName,
      required this.lastName,
      required this.currency,
      required this.displayName,
      required this.birthDate,
      required final List<String> groups,
      required this.isVerified,
      required this.onboardingFinished,
      this.profilePicture,
      this.transactionPercentage,
      this.rewardPercentage,
      required this.newsletter})
      : _groups = groups;

  factory _$UserModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$UserModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String email;
  @override
  final String country;
  @override
  final String referralCode;
  @override
  final double balance;
  @override
  final int createdAt;
  @override
  final String firstName;
  @override
  final String lastName;
  @override
  final String currency;
  @override
  final String displayName;
  @override
  final String birthDate;
  final List<String> _groups;
  @override
  List<String> get groups {
    if (_groups is EqualUnmodifiableListView) return _groups;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_groups);
  }

  @override
  final bool isVerified;
  @override
  final bool onboardingFinished;
  @override
  final String? profilePicture;
  @override
  final double? transactionPercentage;
  @override
  final double? rewardPercentage;
  @override
  final bool newsletter;

  @override
  String toString() {
    return 'UserModel(uuid: $uuid, email: $email, country: $country, referralCode: $referralCode, balance: $balance, createdAt: $createdAt, firstName: $firstName, lastName: $lastName, currency: $currency, displayName: $displayName, birthDate: $birthDate, groups: $groups, isVerified: $isVerified, onboardingFinished: $onboardingFinished, profilePicture: $profilePicture, transactionPercentage: $transactionPercentage, rewardPercentage: $rewardPercentage, newsletter: $newsletter)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$UserModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.email, email) || other.email == email) &&
            (identical(other.country, country) || other.country == country) &&
            (identical(other.referralCode, referralCode) ||
                other.referralCode == referralCode) &&
            (identical(other.balance, balance) || other.balance == balance) &&
            (identical(other.createdAt, createdAt) ||
                other.createdAt == createdAt) &&
            (identical(other.firstName, firstName) ||
                other.firstName == firstName) &&
            (identical(other.lastName, lastName) ||
                other.lastName == lastName) &&
            (identical(other.currency, currency) ||
                other.currency == currency) &&
            (identical(other.displayName, displayName) ||
                other.displayName == displayName) &&
            (identical(other.birthDate, birthDate) ||
                other.birthDate == birthDate) &&
            const DeepCollectionEquality().equals(other._groups, _groups) &&
            (identical(other.isVerified, isVerified) ||
                other.isVerified == isVerified) &&
            (identical(other.onboardingFinished, onboardingFinished) ||
                other.onboardingFinished == onboardingFinished) &&
            (identical(other.profilePicture, profilePicture) ||
                other.profilePicture == profilePicture) &&
            (identical(other.transactionPercentage, transactionPercentage) ||
                other.transactionPercentage == transactionPercentage) &&
            (identical(other.rewardPercentage, rewardPercentage) ||
                other.rewardPercentage == rewardPercentage) &&
            (identical(other.newsletter, newsletter) ||
                other.newsletter == newsletter));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      uuid,
      email,
      country,
      referralCode,
      balance,
      createdAt,
      firstName,
      lastName,
      currency,
      displayName,
      birthDate,
      const DeepCollectionEquality().hash(_groups),
      isVerified,
      onboardingFinished,
      profilePicture,
      transactionPercentage,
      rewardPercentage,
      newsletter);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$UserModelImplCopyWith<_$UserModelImpl> get copyWith =>
      __$$UserModelImplCopyWithImpl<_$UserModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$UserModelImplToJson(
      this,
    );
  }
}

abstract class _UserModel implements UserModel {
  const factory _UserModel(
      {required final String uuid,
      required final String email,
      required final String country,
      required final String referralCode,
      required final double balance,
      required final int createdAt,
      required final String firstName,
      required final String lastName,
      required final String currency,
      required final String displayName,
      required final String birthDate,
      required final List<String> groups,
      required final bool isVerified,
      required final bool onboardingFinished,
      final String? profilePicture,
      final double? transactionPercentage,
      final double? rewardPercentage,
      required final bool newsletter}) = _$UserModelImpl;

  factory _UserModel.fromJson(Map<String, dynamic> json) =
      _$UserModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get email;
  @override
  String get country;
  @override
  String get referralCode;
  @override
  double get balance;
  @override
  int get createdAt;
  @override
  String get firstName;
  @override
  String get lastName;
  @override
  String get currency;
  @override
  String get displayName;
  @override
  String get birthDate;
  @override
  List<String> get groups;
  @override
  bool get isVerified;
  @override
  bool get onboardingFinished;
  @override
  String? get profilePicture;
  @override
  double? get transactionPercentage;
  @override
  double? get rewardPercentage;
  @override
  bool get newsletter;
  @override
  @JsonKey(ignore: true)
  _$$UserModelImplCopyWith<_$UserModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
