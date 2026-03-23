// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'store_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

StoreModel _$StoreModelFromJson(Map<String, dynamic> json) {
  return _StoreModel.fromJson(json);
}

/// @nodoc
mixin _$StoreModel {
  String get uuid => throw _privateConstructorUsedError;
  String get name => throw _privateConstructorUsedError;
  String? get logo => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $StoreModelCopyWith<StoreModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $StoreModelCopyWith<$Res> {
  factory $StoreModelCopyWith(
          StoreModel value, $Res Function(StoreModel) then) =
      _$StoreModelCopyWithImpl<$Res, StoreModel>;
  @useResult
  $Res call({String uuid, String name, String? logo});
}

/// @nodoc
class _$StoreModelCopyWithImpl<$Res, $Val extends StoreModel>
    implements $StoreModelCopyWith<$Res> {
  _$StoreModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? name = null,
    Object? logo = freezed,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      logo: freezed == logo
          ? _value.logo
          : logo // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$StoreModelImplCopyWith<$Res>
    implements $StoreModelCopyWith<$Res> {
  factory _$$StoreModelImplCopyWith(
          _$StoreModelImpl value, $Res Function(_$StoreModelImpl) then) =
      __$$StoreModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String uuid, String name, String? logo});
}

/// @nodoc
class __$$StoreModelImplCopyWithImpl<$Res>
    extends _$StoreModelCopyWithImpl<$Res, _$StoreModelImpl>
    implements _$$StoreModelImplCopyWith<$Res> {
  __$$StoreModelImplCopyWithImpl(
      _$StoreModelImpl _value, $Res Function(_$StoreModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? name = null,
    Object? logo = freezed,
  }) {
    return _then(_$StoreModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      logo: freezed == logo
          ? _value.logo
          : logo // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$StoreModelImpl implements _StoreModel {
  const _$StoreModelImpl({required this.uuid, required this.name, this.logo});

  factory _$StoreModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$StoreModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String name;
  @override
  final String? logo;

  @override
  String toString() {
    return 'StoreModel(uuid: $uuid, name: $name, logo: $logo)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$StoreModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.name, name) || other.name == name) &&
            (identical(other.logo, logo) || other.logo == logo));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, uuid, name, logo);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$StoreModelImplCopyWith<_$StoreModelImpl> get copyWith =>
      __$$StoreModelImplCopyWithImpl<_$StoreModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$StoreModelImplToJson(
      this,
    );
  }
}

abstract class _StoreModel implements StoreModel {
  const factory _StoreModel(
      {required final String uuid,
      required final String name,
      final String? logo}) = _$StoreModelImpl;

  factory _StoreModel.fromJson(Map<String, dynamic> json) =
      _$StoreModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get name;
  @override
  String? get logo;
  @override
  @JsonKey(ignore: true)
  _$$StoreModelImplCopyWith<_$StoreModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

CashbackStoreModel _$CashbackStoreModelFromJson(Map<String, dynamic> json) {
  return _CashbackStoreModel.fromJson(json);
}

/// @nodoc
mixin _$CashbackStoreModel {
  String get uuid => throw _privateConstructorUsedError;
  String get name => throw _privateConstructorUsedError;
  String? get logo => throw _privateConstructorUsedError;
  double? get percentageCashout => throw _privateConstructorUsedError;
  double? get cashbackValue => throw _privateConstructorUsedError;
  String? get cashbackType => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $CashbackStoreModelCopyWith<CashbackStoreModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $CashbackStoreModelCopyWith<$Res> {
  factory $CashbackStoreModelCopyWith(
          CashbackStoreModel value, $Res Function(CashbackStoreModel) then) =
      _$CashbackStoreModelCopyWithImpl<$Res, CashbackStoreModel>;
  @useResult
  $Res call(
      {String uuid,
      String name,
      String? logo,
      double? percentageCashout,
      double? cashbackValue,
      String? cashbackType});
}

/// @nodoc
class _$CashbackStoreModelCopyWithImpl<$Res, $Val extends CashbackStoreModel>
    implements $CashbackStoreModelCopyWith<$Res> {
  _$CashbackStoreModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? name = null,
    Object? logo = freezed,
    Object? percentageCashout = freezed,
    Object? cashbackValue = freezed,
    Object? cashbackType = freezed,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      logo: freezed == logo
          ? _value.logo
          : logo // ignore: cast_nullable_to_non_nullable
              as String?,
      percentageCashout: freezed == percentageCashout
          ? _value.percentageCashout
          : percentageCashout // ignore: cast_nullable_to_non_nullable
              as double?,
      cashbackValue: freezed == cashbackValue
          ? _value.cashbackValue
          : cashbackValue // ignore: cast_nullable_to_non_nullable
              as double?,
      cashbackType: freezed == cashbackType
          ? _value.cashbackType
          : cashbackType // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$CashbackStoreModelImplCopyWith<$Res>
    implements $CashbackStoreModelCopyWith<$Res> {
  factory _$$CashbackStoreModelImplCopyWith(_$CashbackStoreModelImpl value,
          $Res Function(_$CashbackStoreModelImpl) then) =
      __$$CashbackStoreModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      String name,
      String? logo,
      double? percentageCashout,
      double? cashbackValue,
      String? cashbackType});
}

/// @nodoc
class __$$CashbackStoreModelImplCopyWithImpl<$Res>
    extends _$CashbackStoreModelCopyWithImpl<$Res, _$CashbackStoreModelImpl>
    implements _$$CashbackStoreModelImplCopyWith<$Res> {
  __$$CashbackStoreModelImplCopyWithImpl(_$CashbackStoreModelImpl _value,
      $Res Function(_$CashbackStoreModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? name = null,
    Object? logo = freezed,
    Object? percentageCashout = freezed,
    Object? cashbackValue = freezed,
    Object? cashbackType = freezed,
  }) {
    return _then(_$CashbackStoreModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      logo: freezed == logo
          ? _value.logo
          : logo // ignore: cast_nullable_to_non_nullable
              as String?,
      percentageCashout: freezed == percentageCashout
          ? _value.percentageCashout
          : percentageCashout // ignore: cast_nullable_to_non_nullable
              as double?,
      cashbackValue: freezed == cashbackValue
          ? _value.cashbackValue
          : cashbackValue // ignore: cast_nullable_to_non_nullable
              as double?,
      cashbackType: freezed == cashbackType
          ? _value.cashbackType
          : cashbackType // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$CashbackStoreModelImpl implements _CashbackStoreModel {
  const _$CashbackStoreModelImpl(
      {required this.uuid,
      required this.name,
      this.logo,
      this.percentageCashout,
      this.cashbackValue,
      this.cashbackType});

  factory _$CashbackStoreModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$CashbackStoreModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String name;
  @override
  final String? logo;
  @override
  final double? percentageCashout;
  @override
  final double? cashbackValue;
  @override
  final String? cashbackType;

  @override
  String toString() {
    return 'CashbackStoreModel(uuid: $uuid, name: $name, logo: $logo, percentageCashout: $percentageCashout, cashbackValue: $cashbackValue, cashbackType: $cashbackType)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$CashbackStoreModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.name, name) || other.name == name) &&
            (identical(other.logo, logo) || other.logo == logo) &&
            (identical(other.percentageCashout, percentageCashout) ||
                other.percentageCashout == percentageCashout) &&
            (identical(other.cashbackValue, cashbackValue) ||
                other.cashbackValue == cashbackValue) &&
            (identical(other.cashbackType, cashbackType) ||
                other.cashbackType == cashbackType));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, uuid, name, logo,
      percentageCashout, cashbackValue, cashbackType);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$CashbackStoreModelImplCopyWith<_$CashbackStoreModelImpl> get copyWith =>
      __$$CashbackStoreModelImplCopyWithImpl<_$CashbackStoreModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$CashbackStoreModelImplToJson(
      this,
    );
  }
}

abstract class _CashbackStoreModel implements CashbackStoreModel {
  const factory _CashbackStoreModel(
      {required final String uuid,
      required final String name,
      final String? logo,
      final double? percentageCashout,
      final double? cashbackValue,
      final String? cashbackType}) = _$CashbackStoreModelImpl;

  factory _CashbackStoreModel.fromJson(Map<String, dynamic> json) =
      _$CashbackStoreModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get name;
  @override
  String? get logo;
  @override
  double? get percentageCashout;
  @override
  double? get cashbackValue;
  @override
  String? get cashbackType;
  @override
  @JsonKey(ignore: true)
  _$$CashbackStoreModelImplCopyWith<_$CashbackStoreModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
