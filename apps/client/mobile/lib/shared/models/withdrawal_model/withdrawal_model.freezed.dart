// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'withdrawal_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

WithdrawalModel _$WithdrawalModelFromJson(Map<String, dynamic> json) {
  return _WithdrawalModel.fromJson(json);
}

/// @nodoc
mixin _$WithdrawalModel {
  String get uuid => throw _privateConstructorUsedError;
  Object get user => throw _privateConstructorUsedError;
  double get amountSource => throw _privateConstructorUsedError;
  double get amountTarget => throw _privateConstructorUsedError;
  String? get details => throw _privateConstructorUsedError;
  String get state => throw _privateConstructorUsedError;
  String? get completionDate => throw _privateConstructorUsedError;
  String get createdAt => throw _privateConstructorUsedError;
  String get createdBy => throw _privateConstructorUsedError;
  String get updatedAt => throw _privateConstructorUsedError;
  String get updatedBy => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $WithdrawalModelCopyWith<WithdrawalModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $WithdrawalModelCopyWith<$Res> {
  factory $WithdrawalModelCopyWith(
          WithdrawalModel value, $Res Function(WithdrawalModel) then) =
      _$WithdrawalModelCopyWithImpl<$Res, WithdrawalModel>;
  @useResult
  $Res call(
      {String uuid,
      Object user,
      double amountSource,
      double amountTarget,
      String? details,
      String state,
      String? completionDate,
      String createdAt,
      String createdBy,
      String updatedAt,
      String updatedBy});
}

/// @nodoc
class _$WithdrawalModelCopyWithImpl<$Res, $Val extends WithdrawalModel>
    implements $WithdrawalModelCopyWith<$Res> {
  _$WithdrawalModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = null,
    Object? amountSource = null,
    Object? amountTarget = null,
    Object? details = freezed,
    Object? state = null,
    Object? completionDate = freezed,
    Object? createdAt = null,
    Object? createdBy = null,
    Object? updatedAt = null,
    Object? updatedBy = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: null == user ? _value.user : user,
      amountSource: null == amountSource
          ? _value.amountSource
          : amountSource // ignore: cast_nullable_to_non_nullable
              as double,
      amountTarget: null == amountTarget
          ? _value.amountTarget
          : amountTarget // ignore: cast_nullable_to_non_nullable
              as double,
      details: freezed == details
          ? _value.details
          : details // ignore: cast_nullable_to_non_nullable
              as String?,
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      completionDate: freezed == completionDate
          ? _value.completionDate
          : completionDate // ignore: cast_nullable_to_non_nullable
              as String?,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as String,
      createdBy: null == createdBy
          ? _value.createdBy
          : createdBy // ignore: cast_nullable_to_non_nullable
              as String,
      updatedAt: null == updatedAt
          ? _value.updatedAt
          : updatedAt // ignore: cast_nullable_to_non_nullable
              as String,
      updatedBy: null == updatedBy
          ? _value.updatedBy
          : updatedBy // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$WithdrawalModelImplCopyWith<$Res>
    implements $WithdrawalModelCopyWith<$Res> {
  factory _$$WithdrawalModelImplCopyWith(_$WithdrawalModelImpl value,
          $Res Function(_$WithdrawalModelImpl) then) =
      __$$WithdrawalModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      Object user,
      double amountSource,
      double amountTarget,
      String? details,
      String state,
      String? completionDate,
      String createdAt,
      String createdBy,
      String updatedAt,
      String updatedBy});
}

/// @nodoc
class __$$WithdrawalModelImplCopyWithImpl<$Res>
    extends _$WithdrawalModelCopyWithImpl<$Res, _$WithdrawalModelImpl>
    implements _$$WithdrawalModelImplCopyWith<$Res> {
  __$$WithdrawalModelImplCopyWithImpl(
      _$WithdrawalModelImpl _value, $Res Function(_$WithdrawalModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = null,
    Object? amountSource = null,
    Object? amountTarget = null,
    Object? details = freezed,
    Object? state = null,
    Object? completionDate = freezed,
    Object? createdAt = null,
    Object? createdBy = null,
    Object? updatedAt = null,
    Object? updatedBy = null,
  }) {
    return _then(_$WithdrawalModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: null == user ? _value.user : user,
      amountSource: null == amountSource
          ? _value.amountSource
          : amountSource // ignore: cast_nullable_to_non_nullable
              as double,
      amountTarget: null == amountTarget
          ? _value.amountTarget
          : amountTarget // ignore: cast_nullable_to_non_nullable
              as double,
      details: freezed == details
          ? _value.details
          : details // ignore: cast_nullable_to_non_nullable
              as String?,
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      completionDate: freezed == completionDate
          ? _value.completionDate
          : completionDate // ignore: cast_nullable_to_non_nullable
              as String?,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as String,
      createdBy: null == createdBy
          ? _value.createdBy
          : createdBy // ignore: cast_nullable_to_non_nullable
              as String,
      updatedAt: null == updatedAt
          ? _value.updatedAt
          : updatedAt // ignore: cast_nullable_to_non_nullable
              as String,
      updatedBy: null == updatedBy
          ? _value.updatedBy
          : updatedBy // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$WithdrawalModelImpl implements _WithdrawalModel {
  const _$WithdrawalModelImpl(
      {required this.uuid,
      required this.user,
      required this.amountSource,
      required this.amountTarget,
      this.details,
      required this.state,
      this.completionDate,
      required this.createdAt,
      required this.createdBy,
      required this.updatedAt,
      required this.updatedBy});

  factory _$WithdrawalModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$WithdrawalModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final Object user;
  @override
  final double amountSource;
  @override
  final double amountTarget;
  @override
  final String? details;
  @override
  final String state;
  @override
  final String? completionDate;
  @override
  final String createdAt;
  @override
  final String createdBy;
  @override
  final String updatedAt;
  @override
  final String updatedBy;

  @override
  String toString() {
    return 'WithdrawalModel(uuid: $uuid, user: $user, amountSource: $amountSource, amountTarget: $amountTarget, details: $details, state: $state, completionDate: $completionDate, createdAt: $createdAt, createdBy: $createdBy, updatedAt: $updatedAt, updatedBy: $updatedBy)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$WithdrawalModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            const DeepCollectionEquality().equals(other.user, user) &&
            (identical(other.amountSource, amountSource) ||
                other.amountSource == amountSource) &&
            (identical(other.amountTarget, amountTarget) ||
                other.amountTarget == amountTarget) &&
            (identical(other.details, details) || other.details == details) &&
            (identical(other.state, state) || other.state == state) &&
            (identical(other.completionDate, completionDate) ||
                other.completionDate == completionDate) &&
            (identical(other.createdAt, createdAt) ||
                other.createdAt == createdAt) &&
            (identical(other.createdBy, createdBy) ||
                other.createdBy == createdBy) &&
            (identical(other.updatedAt, updatedAt) ||
                other.updatedAt == updatedAt) &&
            (identical(other.updatedBy, updatedBy) ||
                other.updatedBy == updatedBy));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      uuid,
      const DeepCollectionEquality().hash(user),
      amountSource,
      amountTarget,
      details,
      state,
      completionDate,
      createdAt,
      createdBy,
      updatedAt,
      updatedBy);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$WithdrawalModelImplCopyWith<_$WithdrawalModelImpl> get copyWith =>
      __$$WithdrawalModelImplCopyWithImpl<_$WithdrawalModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$WithdrawalModelImplToJson(
      this,
    );
  }
}

abstract class _WithdrawalModel implements WithdrawalModel {
  const factory _WithdrawalModel(
      {required final String uuid,
      required final Object user,
      required final double amountSource,
      required final double amountTarget,
      final String? details,
      required final String state,
      final String? completionDate,
      required final String createdAt,
      required final String createdBy,
      required final String updatedAt,
      required final String updatedBy}) = _$WithdrawalModelImpl;

  factory _WithdrawalModel.fromJson(Map<String, dynamic> json) =
      _$WithdrawalModelImpl.fromJson;

  @override
  String get uuid;
  @override
  Object get user;
  @override
  double get amountSource;
  @override
  double get amountTarget;
  @override
  String? get details;
  @override
  String get state;
  @override
  String? get completionDate;
  @override
  String get createdAt;
  @override
  String get createdBy;
  @override
  String get updatedAt;
  @override
  String get updatedBy;
  @override
  @JsonKey(ignore: true)
  _$$WithdrawalModelImplCopyWith<_$WithdrawalModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

WithdrawalsRequest _$WithdrawalsRequestFromJson(Map<String, dynamic> json) {
  return _WithdrawalsRequest.fromJson(json);
}

/// @nodoc
mixin _$WithdrawalsRequest {
  String get uuid => throw _privateConstructorUsedError;
  dynamic get user => throw _privateConstructorUsedError;
  String get state => throw _privateConstructorUsedError;
  double get amountTarget => throw _privateConstructorUsedError;
  String? get currencyTarget => throw _privateConstructorUsedError;
  String get createdAt => throw _privateConstructorUsedError;
  String? get completionDate => throw _privateConstructorUsedError;
  String? get details => throw _privateConstructorUsedError;
  PaymentMethodModel get paymentMethod => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $WithdrawalsRequestCopyWith<WithdrawalsRequest> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $WithdrawalsRequestCopyWith<$Res> {
  factory $WithdrawalsRequestCopyWith(
          WithdrawalsRequest value, $Res Function(WithdrawalsRequest) then) =
      _$WithdrawalsRequestCopyWithImpl<$Res, WithdrawalsRequest>;
  @useResult
  $Res call(
      {String uuid,
      dynamic user,
      String state,
      double amountTarget,
      String? currencyTarget,
      String createdAt,
      String? completionDate,
      String? details,
      PaymentMethodModel paymentMethod});

  $PaymentMethodModelCopyWith<$Res> get paymentMethod;
}

/// @nodoc
class _$WithdrawalsRequestCopyWithImpl<$Res, $Val extends WithdrawalsRequest>
    implements $WithdrawalsRequestCopyWith<$Res> {
  _$WithdrawalsRequestCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = freezed,
    Object? state = null,
    Object? amountTarget = null,
    Object? currencyTarget = freezed,
    Object? createdAt = null,
    Object? completionDate = freezed,
    Object? details = freezed,
    Object? paymentMethod = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: freezed == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as dynamic,
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      amountTarget: null == amountTarget
          ? _value.amountTarget
          : amountTarget // ignore: cast_nullable_to_non_nullable
              as double,
      currencyTarget: freezed == currencyTarget
          ? _value.currencyTarget
          : currencyTarget // ignore: cast_nullable_to_non_nullable
              as String?,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as String,
      completionDate: freezed == completionDate
          ? _value.completionDate
          : completionDate // ignore: cast_nullable_to_non_nullable
              as String?,
      details: freezed == details
          ? _value.details
          : details // ignore: cast_nullable_to_non_nullable
              as String?,
      paymentMethod: null == paymentMethod
          ? _value.paymentMethod
          : paymentMethod // ignore: cast_nullable_to_non_nullable
              as PaymentMethodModel,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $PaymentMethodModelCopyWith<$Res> get paymentMethod {
    return $PaymentMethodModelCopyWith<$Res>(_value.paymentMethod, (value) {
      return _then(_value.copyWith(paymentMethod: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$WithdrawalsRequestImplCopyWith<$Res>
    implements $WithdrawalsRequestCopyWith<$Res> {
  factory _$$WithdrawalsRequestImplCopyWith(_$WithdrawalsRequestImpl value,
          $Res Function(_$WithdrawalsRequestImpl) then) =
      __$$WithdrawalsRequestImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      dynamic user,
      String state,
      double amountTarget,
      String? currencyTarget,
      String createdAt,
      String? completionDate,
      String? details,
      PaymentMethodModel paymentMethod});

  @override
  $PaymentMethodModelCopyWith<$Res> get paymentMethod;
}

/// @nodoc
class __$$WithdrawalsRequestImplCopyWithImpl<$Res>
    extends _$WithdrawalsRequestCopyWithImpl<$Res, _$WithdrawalsRequestImpl>
    implements _$$WithdrawalsRequestImplCopyWith<$Res> {
  __$$WithdrawalsRequestImplCopyWithImpl(_$WithdrawalsRequestImpl _value,
      $Res Function(_$WithdrawalsRequestImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = freezed,
    Object? state = null,
    Object? amountTarget = null,
    Object? currencyTarget = freezed,
    Object? createdAt = null,
    Object? completionDate = freezed,
    Object? details = freezed,
    Object? paymentMethod = null,
  }) {
    return _then(_$WithdrawalsRequestImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: freezed == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as dynamic,
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      amountTarget: null == amountTarget
          ? _value.amountTarget
          : amountTarget // ignore: cast_nullable_to_non_nullable
              as double,
      currencyTarget: freezed == currencyTarget
          ? _value.currencyTarget
          : currencyTarget // ignore: cast_nullable_to_non_nullable
              as String?,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as String,
      completionDate: freezed == completionDate
          ? _value.completionDate
          : completionDate // ignore: cast_nullable_to_non_nullable
              as String?,
      details: freezed == details
          ? _value.details
          : details // ignore: cast_nullable_to_non_nullable
              as String?,
      paymentMethod: null == paymentMethod
          ? _value.paymentMethod
          : paymentMethod // ignore: cast_nullable_to_non_nullable
              as PaymentMethodModel,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$WithdrawalsRequestImpl implements _WithdrawalsRequest {
  const _$WithdrawalsRequestImpl(
      {required this.uuid,
      required this.user,
      required this.state,
      required this.amountTarget,
      this.currencyTarget,
      required this.createdAt,
      this.completionDate,
      this.details,
      required this.paymentMethod});

  factory _$WithdrawalsRequestImpl.fromJson(Map<String, dynamic> json) =>
      _$$WithdrawalsRequestImplFromJson(json);

  @override
  final String uuid;
  @override
  final dynamic user;
  @override
  final String state;
  @override
  final double amountTarget;
  @override
  final String? currencyTarget;
  @override
  final String createdAt;
  @override
  final String? completionDate;
  @override
  final String? details;
  @override
  final PaymentMethodModel paymentMethod;

  @override
  String toString() {
    return 'WithdrawalsRequest(uuid: $uuid, user: $user, state: $state, amountTarget: $amountTarget, currencyTarget: $currencyTarget, createdAt: $createdAt, completionDate: $completionDate, details: $details, paymentMethod: $paymentMethod)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$WithdrawalsRequestImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            const DeepCollectionEquality().equals(other.user, user) &&
            (identical(other.state, state) || other.state == state) &&
            (identical(other.amountTarget, amountTarget) ||
                other.amountTarget == amountTarget) &&
            (identical(other.currencyTarget, currencyTarget) ||
                other.currencyTarget == currencyTarget) &&
            (identical(other.createdAt, createdAt) ||
                other.createdAt == createdAt) &&
            (identical(other.completionDate, completionDate) ||
                other.completionDate == completionDate) &&
            (identical(other.details, details) || other.details == details) &&
            (identical(other.paymentMethod, paymentMethod) ||
                other.paymentMethod == paymentMethod));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      uuid,
      const DeepCollectionEquality().hash(user),
      state,
      amountTarget,
      currencyTarget,
      createdAt,
      completionDate,
      details,
      paymentMethod);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$WithdrawalsRequestImplCopyWith<_$WithdrawalsRequestImpl> get copyWith =>
      __$$WithdrawalsRequestImplCopyWithImpl<_$WithdrawalsRequestImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$WithdrawalsRequestImplToJson(
      this,
    );
  }
}

abstract class _WithdrawalsRequest implements WithdrawalsRequest {
  const factory _WithdrawalsRequest(
          {required final String uuid,
          required final dynamic user,
          required final String state,
          required final double amountTarget,
          final String? currencyTarget,
          required final String createdAt,
          final String? completionDate,
          final String? details,
          required final PaymentMethodModel paymentMethod}) =
      _$WithdrawalsRequestImpl;

  factory _WithdrawalsRequest.fromJson(Map<String, dynamic> json) =
      _$WithdrawalsRequestImpl.fromJson;

  @override
  String get uuid;
  @override
  dynamic get user;
  @override
  String get state;
  @override
  double get amountTarget;
  @override
  String? get currencyTarget;
  @override
  String get createdAt;
  @override
  String? get completionDate;
  @override
  String? get details;
  @override
  PaymentMethodModel get paymentMethod;
  @override
  @JsonKey(ignore: true)
  _$$WithdrawalsRequestImplCopyWith<_$WithdrawalsRequestImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

PaymentMethodModel _$PaymentMethodModelFromJson(Map<String, dynamic> json) {
  return _PaymentMethodModel.fromJson(json);
}

/// @nodoc
mixin _$PaymentMethodModel {
  String get paymentMethod => throw _privateConstructorUsedError;
  String get bankName => throw _privateConstructorUsedError;
  String get bankAddress => throw _privateConstructorUsedError;
  String get bankCountry => throw _privateConstructorUsedError;
  String get country => throw _privateConstructorUsedError;
  String get vat => throw _privateConstructorUsedError;
  String get bankAccountTitle => throw _privateConstructorUsedError;
  String get iban => throw _privateConstructorUsedError;
  IbanStatementModel get ibanStatement => throw _privateConstructorUsedError;
  String? get state => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $PaymentMethodModelCopyWith<PaymentMethodModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $PaymentMethodModelCopyWith<$Res> {
  factory $PaymentMethodModelCopyWith(
          PaymentMethodModel value, $Res Function(PaymentMethodModel) then) =
      _$PaymentMethodModelCopyWithImpl<$Res, PaymentMethodModel>;
  @useResult
  $Res call(
      {String paymentMethod,
      String bankName,
      String bankAddress,
      String bankCountry,
      String country,
      String vat,
      String bankAccountTitle,
      String iban,
      IbanStatementModel ibanStatement,
      String? state});

  $IbanStatementModelCopyWith<$Res> get ibanStatement;
}

/// @nodoc
class _$PaymentMethodModelCopyWithImpl<$Res, $Val extends PaymentMethodModel>
    implements $PaymentMethodModelCopyWith<$Res> {
  _$PaymentMethodModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? paymentMethod = null,
    Object? bankName = null,
    Object? bankAddress = null,
    Object? bankCountry = null,
    Object? country = null,
    Object? vat = null,
    Object? bankAccountTitle = null,
    Object? iban = null,
    Object? ibanStatement = null,
    Object? state = freezed,
  }) {
    return _then(_value.copyWith(
      paymentMethod: null == paymentMethod
          ? _value.paymentMethod
          : paymentMethod // ignore: cast_nullable_to_non_nullable
              as String,
      bankName: null == bankName
          ? _value.bankName
          : bankName // ignore: cast_nullable_to_non_nullable
              as String,
      bankAddress: null == bankAddress
          ? _value.bankAddress
          : bankAddress // ignore: cast_nullable_to_non_nullable
              as String,
      bankCountry: null == bankCountry
          ? _value.bankCountry
          : bankCountry // ignore: cast_nullable_to_non_nullable
              as String,
      country: null == country
          ? _value.country
          : country // ignore: cast_nullable_to_non_nullable
              as String,
      vat: null == vat
          ? _value.vat
          : vat // ignore: cast_nullable_to_non_nullable
              as String,
      bankAccountTitle: null == bankAccountTitle
          ? _value.bankAccountTitle
          : bankAccountTitle // ignore: cast_nullable_to_non_nullable
              as String,
      iban: null == iban
          ? _value.iban
          : iban // ignore: cast_nullable_to_non_nullable
              as String,
      ibanStatement: null == ibanStatement
          ? _value.ibanStatement
          : ibanStatement // ignore: cast_nullable_to_non_nullable
              as IbanStatementModel,
      state: freezed == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $IbanStatementModelCopyWith<$Res> get ibanStatement {
    return $IbanStatementModelCopyWith<$Res>(_value.ibanStatement, (value) {
      return _then(_value.copyWith(ibanStatement: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$PaymentMethodModelImplCopyWith<$Res>
    implements $PaymentMethodModelCopyWith<$Res> {
  factory _$$PaymentMethodModelImplCopyWith(_$PaymentMethodModelImpl value,
          $Res Function(_$PaymentMethodModelImpl) then) =
      __$$PaymentMethodModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String paymentMethod,
      String bankName,
      String bankAddress,
      String bankCountry,
      String country,
      String vat,
      String bankAccountTitle,
      String iban,
      IbanStatementModel ibanStatement,
      String? state});

  @override
  $IbanStatementModelCopyWith<$Res> get ibanStatement;
}

/// @nodoc
class __$$PaymentMethodModelImplCopyWithImpl<$Res>
    extends _$PaymentMethodModelCopyWithImpl<$Res, _$PaymentMethodModelImpl>
    implements _$$PaymentMethodModelImplCopyWith<$Res> {
  __$$PaymentMethodModelImplCopyWithImpl(_$PaymentMethodModelImpl _value,
      $Res Function(_$PaymentMethodModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? paymentMethod = null,
    Object? bankName = null,
    Object? bankAddress = null,
    Object? bankCountry = null,
    Object? country = null,
    Object? vat = null,
    Object? bankAccountTitle = null,
    Object? iban = null,
    Object? ibanStatement = null,
    Object? state = freezed,
  }) {
    return _then(_$PaymentMethodModelImpl(
      paymentMethod: null == paymentMethod
          ? _value.paymentMethod
          : paymentMethod // ignore: cast_nullable_to_non_nullable
              as String,
      bankName: null == bankName
          ? _value.bankName
          : bankName // ignore: cast_nullable_to_non_nullable
              as String,
      bankAddress: null == bankAddress
          ? _value.bankAddress
          : bankAddress // ignore: cast_nullable_to_non_nullable
              as String,
      bankCountry: null == bankCountry
          ? _value.bankCountry
          : bankCountry // ignore: cast_nullable_to_non_nullable
              as String,
      country: null == country
          ? _value.country
          : country // ignore: cast_nullable_to_non_nullable
              as String,
      vat: null == vat
          ? _value.vat
          : vat // ignore: cast_nullable_to_non_nullable
              as String,
      bankAccountTitle: null == bankAccountTitle
          ? _value.bankAccountTitle
          : bankAccountTitle // ignore: cast_nullable_to_non_nullable
              as String,
      iban: null == iban
          ? _value.iban
          : iban // ignore: cast_nullable_to_non_nullable
              as String,
      ibanStatement: null == ibanStatement
          ? _value.ibanStatement
          : ibanStatement // ignore: cast_nullable_to_non_nullable
              as IbanStatementModel,
      state: freezed == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$PaymentMethodModelImpl implements _PaymentMethodModel {
  const _$PaymentMethodModelImpl(
      {required this.paymentMethod,
      required this.bankName,
      required this.bankAddress,
      required this.bankCountry,
      required this.country,
      required this.vat,
      required this.bankAccountTitle,
      required this.iban,
      required this.ibanStatement,
      this.state});

  factory _$PaymentMethodModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$PaymentMethodModelImplFromJson(json);

  @override
  final String paymentMethod;
  @override
  final String bankName;
  @override
  final String bankAddress;
  @override
  final String bankCountry;
  @override
  final String country;
  @override
  final String vat;
  @override
  final String bankAccountTitle;
  @override
  final String iban;
  @override
  final IbanStatementModel ibanStatement;
  @override
  final String? state;

  @override
  String toString() {
    return 'PaymentMethodModel(paymentMethod: $paymentMethod, bankName: $bankName, bankAddress: $bankAddress, bankCountry: $bankCountry, country: $country, vat: $vat, bankAccountTitle: $bankAccountTitle, iban: $iban, ibanStatement: $ibanStatement, state: $state)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$PaymentMethodModelImpl &&
            (identical(other.paymentMethod, paymentMethod) ||
                other.paymentMethod == paymentMethod) &&
            (identical(other.bankName, bankName) ||
                other.bankName == bankName) &&
            (identical(other.bankAddress, bankAddress) ||
                other.bankAddress == bankAddress) &&
            (identical(other.bankCountry, bankCountry) ||
                other.bankCountry == bankCountry) &&
            (identical(other.country, country) || other.country == country) &&
            (identical(other.vat, vat) || other.vat == vat) &&
            (identical(other.bankAccountTitle, bankAccountTitle) ||
                other.bankAccountTitle == bankAccountTitle) &&
            (identical(other.iban, iban) || other.iban == iban) &&
            (identical(other.ibanStatement, ibanStatement) ||
                other.ibanStatement == ibanStatement) &&
            (identical(other.state, state) || other.state == state));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      paymentMethod,
      bankName,
      bankAddress,
      bankCountry,
      country,
      vat,
      bankAccountTitle,
      iban,
      ibanStatement,
      state);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$PaymentMethodModelImplCopyWith<_$PaymentMethodModelImpl> get copyWith =>
      __$$PaymentMethodModelImplCopyWithImpl<_$PaymentMethodModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$PaymentMethodModelImplToJson(
      this,
    );
  }
}

abstract class _PaymentMethodModel implements PaymentMethodModel {
  const factory _PaymentMethodModel(
      {required final String paymentMethod,
      required final String bankName,
      required final String bankAddress,
      required final String bankCountry,
      required final String country,
      required final String vat,
      required final String bankAccountTitle,
      required final String iban,
      required final IbanStatementModel ibanStatement,
      final String? state}) = _$PaymentMethodModelImpl;

  factory _PaymentMethodModel.fromJson(Map<String, dynamic> json) =
      _$PaymentMethodModelImpl.fromJson;

  @override
  String get paymentMethod;
  @override
  String get bankName;
  @override
  String get bankAddress;
  @override
  String get bankCountry;
  @override
  String get country;
  @override
  String get vat;
  @override
  String get bankAccountTitle;
  @override
  String get iban;
  @override
  IbanStatementModel get ibanStatement;
  @override
  String? get state;
  @override
  @JsonKey(ignore: true)
  _$$PaymentMethodModelImplCopyWith<_$PaymentMethodModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

IbanStatementModel _$IbanStatementModelFromJson(Map<String, dynamic> json) {
  return _IbanStatementModel.fromJson(json);
}

/// @nodoc
mixin _$IbanStatementModel {
  String get uuid => throw _privateConstructorUsedError;
  String get fileName => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $IbanStatementModelCopyWith<IbanStatementModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $IbanStatementModelCopyWith<$Res> {
  factory $IbanStatementModelCopyWith(
          IbanStatementModel value, $Res Function(IbanStatementModel) then) =
      _$IbanStatementModelCopyWithImpl<$Res, IbanStatementModel>;
  @useResult
  $Res call({String uuid, String fileName});
}

/// @nodoc
class _$IbanStatementModelCopyWithImpl<$Res, $Val extends IbanStatementModel>
    implements $IbanStatementModelCopyWith<$Res> {
  _$IbanStatementModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? fileName = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      fileName: null == fileName
          ? _value.fileName
          : fileName // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$IbanStatementModelImplCopyWith<$Res>
    implements $IbanStatementModelCopyWith<$Res> {
  factory _$$IbanStatementModelImplCopyWith(_$IbanStatementModelImpl value,
          $Res Function(_$IbanStatementModelImpl) then) =
      __$$IbanStatementModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String uuid, String fileName});
}

/// @nodoc
class __$$IbanStatementModelImplCopyWithImpl<$Res>
    extends _$IbanStatementModelCopyWithImpl<$Res, _$IbanStatementModelImpl>
    implements _$$IbanStatementModelImplCopyWith<$Res> {
  __$$IbanStatementModelImplCopyWithImpl(_$IbanStatementModelImpl _value,
      $Res Function(_$IbanStatementModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? fileName = null,
  }) {
    return _then(_$IbanStatementModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      fileName: null == fileName
          ? _value.fileName
          : fileName // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$IbanStatementModelImpl implements _IbanStatementModel {
  const _$IbanStatementModelImpl({required this.uuid, required this.fileName});

  factory _$IbanStatementModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$IbanStatementModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String fileName;

  @override
  String toString() {
    return 'IbanStatementModel(uuid: $uuid, fileName: $fileName)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$IbanStatementModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.fileName, fileName) ||
                other.fileName == fileName));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, uuid, fileName);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$IbanStatementModelImplCopyWith<_$IbanStatementModelImpl> get copyWith =>
      __$$IbanStatementModelImplCopyWithImpl<_$IbanStatementModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$IbanStatementModelImplToJson(
      this,
    );
  }
}

abstract class _IbanStatementModel implements IbanStatementModel {
  const factory _IbanStatementModel(
      {required final String uuid,
      required final String fileName}) = _$IbanStatementModelImpl;

  factory _IbanStatementModel.fromJson(Map<String, dynamic> json) =
      _$IbanStatementModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get fileName;
  @override
  @JsonKey(ignore: true)
  _$$IbanStatementModelImplCopyWith<_$IbanStatementModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
