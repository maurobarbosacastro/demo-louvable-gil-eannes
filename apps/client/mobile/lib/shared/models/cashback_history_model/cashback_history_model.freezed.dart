// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'cashback_history_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

CashbackHistoryModel _$CashbackHistoryModelFromJson(Map<String, dynamic> json) {
  return _CashbackHistoryModel.fromJson(json);
}

/// @nodoc
mixin _$CashbackHistoryModel {
  String get uuid => throw _privateConstructorUsedError;
  double get rate => throw _privateConstructorUsedError;
  double get units => throw _privateConstructorUsedError;
  double get cashReward => throw _privateConstructorUsedError;
  String get createdAt => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $CashbackHistoryModelCopyWith<CashbackHistoryModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $CashbackHistoryModelCopyWith<$Res> {
  factory $CashbackHistoryModelCopyWith(CashbackHistoryModel value,
          $Res Function(CashbackHistoryModel) then) =
      _$CashbackHistoryModelCopyWithImpl<$Res, CashbackHistoryModel>;
  @useResult
  $Res call(
      {String uuid,
      double rate,
      double units,
      double cashReward,
      String createdAt});
}

/// @nodoc
class _$CashbackHistoryModelCopyWithImpl<$Res,
        $Val extends CashbackHistoryModel>
    implements $CashbackHistoryModelCopyWith<$Res> {
  _$CashbackHistoryModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? rate = null,
    Object? units = null,
    Object? cashReward = null,
    Object? createdAt = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      rate: null == rate
          ? _value.rate
          : rate // ignore: cast_nullable_to_non_nullable
              as double,
      units: null == units
          ? _value.units
          : units // ignore: cast_nullable_to_non_nullable
              as double,
      cashReward: null == cashReward
          ? _value.cashReward
          : cashReward // ignore: cast_nullable_to_non_nullable
              as double,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$CashbackHistoryModelImplCopyWith<$Res>
    implements $CashbackHistoryModelCopyWith<$Res> {
  factory _$$CashbackHistoryModelImplCopyWith(_$CashbackHistoryModelImpl value,
          $Res Function(_$CashbackHistoryModelImpl) then) =
      __$$CashbackHistoryModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      double rate,
      double units,
      double cashReward,
      String createdAt});
}

/// @nodoc
class __$$CashbackHistoryModelImplCopyWithImpl<$Res>
    extends _$CashbackHistoryModelCopyWithImpl<$Res, _$CashbackHistoryModelImpl>
    implements _$$CashbackHistoryModelImplCopyWith<$Res> {
  __$$CashbackHistoryModelImplCopyWithImpl(_$CashbackHistoryModelImpl _value,
      $Res Function(_$CashbackHistoryModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? rate = null,
    Object? units = null,
    Object? cashReward = null,
    Object? createdAt = null,
  }) {
    return _then(_$CashbackHistoryModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      rate: null == rate
          ? _value.rate
          : rate // ignore: cast_nullable_to_non_nullable
              as double,
      units: null == units
          ? _value.units
          : units // ignore: cast_nullable_to_non_nullable
              as double,
      cashReward: null == cashReward
          ? _value.cashReward
          : cashReward // ignore: cast_nullable_to_non_nullable
              as double,
      createdAt: null == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$CashbackHistoryModelImpl implements _CashbackHistoryModel {
  const _$CashbackHistoryModelImpl(
      {required this.uuid,
      required this.rate,
      required this.units,
      required this.cashReward,
      required this.createdAt});

  factory _$CashbackHistoryModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$CashbackHistoryModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final double rate;
  @override
  final double units;
  @override
  final double cashReward;
  @override
  final String createdAt;

  @override
  String toString() {
    return 'CashbackHistoryModel(uuid: $uuid, rate: $rate, units: $units, cashReward: $cashReward, createdAt: $createdAt)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$CashbackHistoryModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.rate, rate) || other.rate == rate) &&
            (identical(other.units, units) || other.units == units) &&
            (identical(other.cashReward, cashReward) ||
                other.cashReward == cashReward) &&
            (identical(other.createdAt, createdAt) ||
                other.createdAt == createdAt));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode =>
      Object.hash(runtimeType, uuid, rate, units, cashReward, createdAt);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$CashbackHistoryModelImplCopyWith<_$CashbackHistoryModelImpl>
      get copyWith =>
          __$$CashbackHistoryModelImplCopyWithImpl<_$CashbackHistoryModelImpl>(
              this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$CashbackHistoryModelImplToJson(
      this,
    );
  }
}

abstract class _CashbackHistoryModel implements CashbackHistoryModel {
  const factory _CashbackHistoryModel(
      {required final String uuid,
      required final double rate,
      required final double units,
      required final double cashReward,
      required final String createdAt}) = _$CashbackHistoryModelImpl;

  factory _CashbackHistoryModel.fromJson(Map<String, dynamic> json) =
      _$CashbackHistoryModelImpl.fromJson;

  @override
  String get uuid;
  @override
  double get rate;
  @override
  double get units;
  @override
  double get cashReward;
  @override
  String get createdAt;
  @override
  @JsonKey(ignore: true)
  _$$CashbackHistoryModelImplCopyWith<_$CashbackHistoryModelImpl>
      get copyWith => throw _privateConstructorUsedError;
}
