// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'user_payment_method_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

UserPaymentMethodModel _$UserPaymentMethodModelFromJson(
    Map<String, dynamic> json) {
  return _UserPaymentMethodModel.fromJson(json);
}

/// @nodoc
mixin _$UserPaymentMethodModel {
  String get uuid => throw _privateConstructorUsedError;
  String get paymentMethod => throw _privateConstructorUsedError;
  String get bankName => throw _privateConstructorUsedError;
  String get bankAddress => throw _privateConstructorUsedError;
  String get bankCountry => throw _privateConstructorUsedError;
  String get country => throw _privateConstructorUsedError;
  String get vat => throw _privateConstructorUsedError;
  String get bankAccountTitle => throw _privateConstructorUsedError;
  String get iban => throw _privateConstructorUsedError;
  IbanStatementModel get ibanStatement => throw _privateConstructorUsedError;
  String get state => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $UserPaymentMethodModelCopyWith<UserPaymentMethodModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $UserPaymentMethodModelCopyWith<$Res> {
  factory $UserPaymentMethodModelCopyWith(UserPaymentMethodModel value,
          $Res Function(UserPaymentMethodModel) then) =
      _$UserPaymentMethodModelCopyWithImpl<$Res, UserPaymentMethodModel>;
  @useResult
  $Res call(
      {String uuid,
      String paymentMethod,
      String bankName,
      String bankAddress,
      String bankCountry,
      String country,
      String vat,
      String bankAccountTitle,
      String iban,
      IbanStatementModel ibanStatement,
      String state});

  $IbanStatementModelCopyWith<$Res> get ibanStatement;
}

/// @nodoc
class _$UserPaymentMethodModelCopyWithImpl<$Res,
        $Val extends UserPaymentMethodModel>
    implements $UserPaymentMethodModelCopyWith<$Res> {
  _$UserPaymentMethodModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? paymentMethod = null,
    Object? bankName = null,
    Object? bankAddress = null,
    Object? bankCountry = null,
    Object? country = null,
    Object? vat = null,
    Object? bankAccountTitle = null,
    Object? iban = null,
    Object? ibanStatement = null,
    Object? state = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
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
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
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
abstract class _$$UserPaymentMethodModelImplCopyWith<$Res>
    implements $UserPaymentMethodModelCopyWith<$Res> {
  factory _$$UserPaymentMethodModelImplCopyWith(
          _$UserPaymentMethodModelImpl value,
          $Res Function(_$UserPaymentMethodModelImpl) then) =
      __$$UserPaymentMethodModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      String paymentMethod,
      String bankName,
      String bankAddress,
      String bankCountry,
      String country,
      String vat,
      String bankAccountTitle,
      String iban,
      IbanStatementModel ibanStatement,
      String state});

  @override
  $IbanStatementModelCopyWith<$Res> get ibanStatement;
}

/// @nodoc
class __$$UserPaymentMethodModelImplCopyWithImpl<$Res>
    extends _$UserPaymentMethodModelCopyWithImpl<$Res,
        _$UserPaymentMethodModelImpl>
    implements _$$UserPaymentMethodModelImplCopyWith<$Res> {
  __$$UserPaymentMethodModelImplCopyWithImpl(
      _$UserPaymentMethodModelImpl _value,
      $Res Function(_$UserPaymentMethodModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? paymentMethod = null,
    Object? bankName = null,
    Object? bankAddress = null,
    Object? bankCountry = null,
    Object? country = null,
    Object? vat = null,
    Object? bankAccountTitle = null,
    Object? iban = null,
    Object? ibanStatement = null,
    Object? state = null,
  }) {
    return _then(_$UserPaymentMethodModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
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
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$UserPaymentMethodModelImpl implements _UserPaymentMethodModel {
  const _$UserPaymentMethodModelImpl(
      {required this.uuid,
      required this.paymentMethod,
      required this.bankName,
      required this.bankAddress,
      required this.bankCountry,
      required this.country,
      required this.vat,
      required this.bankAccountTitle,
      required this.iban,
      required this.ibanStatement,
      required this.state});

  factory _$UserPaymentMethodModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$UserPaymentMethodModelImplFromJson(json);

  @override
  final String uuid;
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
  final String state;

  @override
  String toString() {
    return 'UserPaymentMethodModel(uuid: $uuid, paymentMethod: $paymentMethod, bankName: $bankName, bankAddress: $bankAddress, bankCountry: $bankCountry, country: $country, vat: $vat, bankAccountTitle: $bankAccountTitle, iban: $iban, ibanStatement: $ibanStatement, state: $state)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$UserPaymentMethodModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
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
      uuid,
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
  _$$UserPaymentMethodModelImplCopyWith<_$UserPaymentMethodModelImpl>
      get copyWith => __$$UserPaymentMethodModelImplCopyWithImpl<
          _$UserPaymentMethodModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$UserPaymentMethodModelImplToJson(
      this,
    );
  }
}

abstract class _UserPaymentMethodModel implements UserPaymentMethodModel {
  const factory _UserPaymentMethodModel(
      {required final String uuid,
      required final String paymentMethod,
      required final String bankName,
      required final String bankAddress,
      required final String bankCountry,
      required final String country,
      required final String vat,
      required final String bankAccountTitle,
      required final String iban,
      required final IbanStatementModel ibanStatement,
      required final String state}) = _$UserPaymentMethodModelImpl;

  factory _UserPaymentMethodModel.fromJson(Map<String, dynamic> json) =
      _$UserPaymentMethodModelImpl.fromJson;

  @override
  String get uuid;
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
  String get state;
  @override
  @JsonKey(ignore: true)
  _$$UserPaymentMethodModelImplCopyWith<_$UserPaymentMethodModelImpl>
      get copyWith => throw _privateConstructorUsedError;
}
