// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'cashback_table_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

CashbackTableModel _$CashbackTableModelFromJson(Map<String, dynamic> json) {
  return _CashbackTableModel.fromJson(json);
}

/// @nodoc
mixin _$CashbackTableModel {
  String get uuid => throw _privateConstructorUsedError;
  String? get exitId => throw _privateConstructorUsedError;
  CashbackStoreModel? get store => throw _privateConstructorUsedError;
  String? get email => throw _privateConstructorUsedError;
  String get date => throw _privateConstructorUsedError;
  double get amountSource => throw _privateConstructorUsedError;
  double get amountTarget => throw _privateConstructorUsedError;
  double get amountUser => throw _privateConstructorUsedError;
  String get currencySource => throw _privateConstructorUsedError;
  String get currencyTarget => throw _privateConstructorUsedError;
  double? get networkCommission => throw _privateConstructorUsedError;
  String get status => throw _privateConstructorUsedError;
  double get cashback => throw _privateConstructorUsedError;
  @JsonKey(fromJson: _rewardFromJson, toJson: _rewardToJson)
  RewardCashbackModel? get reward => throw _privateConstructorUsedError;
  double? get unvalidatedCurrentReward => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $CashbackTableModelCopyWith<CashbackTableModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $CashbackTableModelCopyWith<$Res> {
  factory $CashbackTableModelCopyWith(
          CashbackTableModel value, $Res Function(CashbackTableModel) then) =
      _$CashbackTableModelCopyWithImpl<$Res, CashbackTableModel>;
  @useResult
  $Res call(
      {String uuid,
      String? exitId,
      CashbackStoreModel? store,
      String? email,
      String date,
      double amountSource,
      double amountTarget,
      double amountUser,
      String currencySource,
      String currencyTarget,
      double? networkCommission,
      String status,
      double cashback,
      @JsonKey(fromJson: _rewardFromJson, toJson: _rewardToJson)
      RewardCashbackModel? reward,
      double? unvalidatedCurrentReward});

  $CashbackStoreModelCopyWith<$Res>? get store;
  $RewardCashbackModelCopyWith<$Res>? get reward;
}

/// @nodoc
class _$CashbackTableModelCopyWithImpl<$Res, $Val extends CashbackTableModel>
    implements $CashbackTableModelCopyWith<$Res> {
  _$CashbackTableModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? exitId = freezed,
    Object? store = freezed,
    Object? email = freezed,
    Object? date = null,
    Object? amountSource = null,
    Object? amountTarget = null,
    Object? amountUser = null,
    Object? currencySource = null,
    Object? currencyTarget = null,
    Object? networkCommission = freezed,
    Object? status = null,
    Object? cashback = null,
    Object? reward = freezed,
    Object? unvalidatedCurrentReward = freezed,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      exitId: freezed == exitId
          ? _value.exitId
          : exitId // ignore: cast_nullable_to_non_nullable
              as String?,
      store: freezed == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as CashbackStoreModel?,
      email: freezed == email
          ? _value.email
          : email // ignore: cast_nullable_to_non_nullable
              as String?,
      date: null == date
          ? _value.date
          : date // ignore: cast_nullable_to_non_nullable
              as String,
      amountSource: null == amountSource
          ? _value.amountSource
          : amountSource // ignore: cast_nullable_to_non_nullable
              as double,
      amountTarget: null == amountTarget
          ? _value.amountTarget
          : amountTarget // ignore: cast_nullable_to_non_nullable
              as double,
      amountUser: null == amountUser
          ? _value.amountUser
          : amountUser // ignore: cast_nullable_to_non_nullable
              as double,
      currencySource: null == currencySource
          ? _value.currencySource
          : currencySource // ignore: cast_nullable_to_non_nullable
              as String,
      currencyTarget: null == currencyTarget
          ? _value.currencyTarget
          : currencyTarget // ignore: cast_nullable_to_non_nullable
              as String,
      networkCommission: freezed == networkCommission
          ? _value.networkCommission
          : networkCommission // ignore: cast_nullable_to_non_nullable
              as double?,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String,
      cashback: null == cashback
          ? _value.cashback
          : cashback // ignore: cast_nullable_to_non_nullable
              as double,
      reward: freezed == reward
          ? _value.reward
          : reward // ignore: cast_nullable_to_non_nullable
              as RewardCashbackModel?,
      unvalidatedCurrentReward: freezed == unvalidatedCurrentReward
          ? _value.unvalidatedCurrentReward
          : unvalidatedCurrentReward // ignore: cast_nullable_to_non_nullable
              as double?,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $CashbackStoreModelCopyWith<$Res>? get store {
    if (_value.store == null) {
      return null;
    }

    return $CashbackStoreModelCopyWith<$Res>(_value.store!, (value) {
      return _then(_value.copyWith(store: value) as $Val);
    });
  }

  @override
  @pragma('vm:prefer-inline')
  $RewardCashbackModelCopyWith<$Res>? get reward {
    if (_value.reward == null) {
      return null;
    }

    return $RewardCashbackModelCopyWith<$Res>(_value.reward!, (value) {
      return _then(_value.copyWith(reward: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$CashbackTableModelImplCopyWith<$Res>
    implements $CashbackTableModelCopyWith<$Res> {
  factory _$$CashbackTableModelImplCopyWith(_$CashbackTableModelImpl value,
          $Res Function(_$CashbackTableModelImpl) then) =
      __$$CashbackTableModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      String? exitId,
      CashbackStoreModel? store,
      String? email,
      String date,
      double amountSource,
      double amountTarget,
      double amountUser,
      String currencySource,
      String currencyTarget,
      double? networkCommission,
      String status,
      double cashback,
      @JsonKey(fromJson: _rewardFromJson, toJson: _rewardToJson)
      RewardCashbackModel? reward,
      double? unvalidatedCurrentReward});

  @override
  $CashbackStoreModelCopyWith<$Res>? get store;
  @override
  $RewardCashbackModelCopyWith<$Res>? get reward;
}

/// @nodoc
class __$$CashbackTableModelImplCopyWithImpl<$Res>
    extends _$CashbackTableModelCopyWithImpl<$Res, _$CashbackTableModelImpl>
    implements _$$CashbackTableModelImplCopyWith<$Res> {
  __$$CashbackTableModelImplCopyWithImpl(_$CashbackTableModelImpl _value,
      $Res Function(_$CashbackTableModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? exitId = freezed,
    Object? store = freezed,
    Object? email = freezed,
    Object? date = null,
    Object? amountSource = null,
    Object? amountTarget = null,
    Object? amountUser = null,
    Object? currencySource = null,
    Object? currencyTarget = null,
    Object? networkCommission = freezed,
    Object? status = null,
    Object? cashback = null,
    Object? reward = freezed,
    Object? unvalidatedCurrentReward = freezed,
  }) {
    return _then(_$CashbackTableModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      exitId: freezed == exitId
          ? _value.exitId
          : exitId // ignore: cast_nullable_to_non_nullable
              as String?,
      store: freezed == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as CashbackStoreModel?,
      email: freezed == email
          ? _value.email
          : email // ignore: cast_nullable_to_non_nullable
              as String?,
      date: null == date
          ? _value.date
          : date // ignore: cast_nullable_to_non_nullable
              as String,
      amountSource: null == amountSource
          ? _value.amountSource
          : amountSource // ignore: cast_nullable_to_non_nullable
              as double,
      amountTarget: null == amountTarget
          ? _value.amountTarget
          : amountTarget // ignore: cast_nullable_to_non_nullable
              as double,
      amountUser: null == amountUser
          ? _value.amountUser
          : amountUser // ignore: cast_nullable_to_non_nullable
              as double,
      currencySource: null == currencySource
          ? _value.currencySource
          : currencySource // ignore: cast_nullable_to_non_nullable
              as String,
      currencyTarget: null == currencyTarget
          ? _value.currencyTarget
          : currencyTarget // ignore: cast_nullable_to_non_nullable
              as String,
      networkCommission: freezed == networkCommission
          ? _value.networkCommission
          : networkCommission // ignore: cast_nullable_to_non_nullable
              as double?,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String,
      cashback: null == cashback
          ? _value.cashback
          : cashback // ignore: cast_nullable_to_non_nullable
              as double,
      reward: freezed == reward
          ? _value.reward
          : reward // ignore: cast_nullable_to_non_nullable
              as RewardCashbackModel?,
      unvalidatedCurrentReward: freezed == unvalidatedCurrentReward
          ? _value.unvalidatedCurrentReward
          : unvalidatedCurrentReward // ignore: cast_nullable_to_non_nullable
              as double?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$CashbackTableModelImpl implements _CashbackTableModel {
  const _$CashbackTableModelImpl(
      {required this.uuid,
      this.exitId,
      this.store,
      this.email,
      required this.date,
      required this.amountSource,
      required this.amountTarget,
      required this.amountUser,
      required this.currencySource,
      required this.currencyTarget,
      this.networkCommission,
      required this.status,
      required this.cashback,
      @JsonKey(fromJson: _rewardFromJson, toJson: _rewardToJson) this.reward,
      this.unvalidatedCurrentReward});

  factory _$CashbackTableModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$CashbackTableModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String? exitId;
  @override
  final CashbackStoreModel? store;
  @override
  final String? email;
  @override
  final String date;
  @override
  final double amountSource;
  @override
  final double amountTarget;
  @override
  final double amountUser;
  @override
  final String currencySource;
  @override
  final String currencyTarget;
  @override
  final double? networkCommission;
  @override
  final String status;
  @override
  final double cashback;
  @override
  @JsonKey(fromJson: _rewardFromJson, toJson: _rewardToJson)
  final RewardCashbackModel? reward;
  @override
  final double? unvalidatedCurrentReward;

  @override
  String toString() {
    return 'CashbackTableModel(uuid: $uuid, exitId: $exitId, store: $store, email: $email, date: $date, amountSource: $amountSource, amountTarget: $amountTarget, amountUser: $amountUser, currencySource: $currencySource, currencyTarget: $currencyTarget, networkCommission: $networkCommission, status: $status, cashback: $cashback, reward: $reward, unvalidatedCurrentReward: $unvalidatedCurrentReward)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$CashbackTableModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.exitId, exitId) || other.exitId == exitId) &&
            (identical(other.store, store) || other.store == store) &&
            (identical(other.email, email) || other.email == email) &&
            (identical(other.date, date) || other.date == date) &&
            (identical(other.amountSource, amountSource) ||
                other.amountSource == amountSource) &&
            (identical(other.amountTarget, amountTarget) ||
                other.amountTarget == amountTarget) &&
            (identical(other.amountUser, amountUser) ||
                other.amountUser == amountUser) &&
            (identical(other.currencySource, currencySource) ||
                other.currencySource == currencySource) &&
            (identical(other.currencyTarget, currencyTarget) ||
                other.currencyTarget == currencyTarget) &&
            (identical(other.networkCommission, networkCommission) ||
                other.networkCommission == networkCommission) &&
            (identical(other.status, status) || other.status == status) &&
            (identical(other.cashback, cashback) ||
                other.cashback == cashback) &&
            (identical(other.reward, reward) || other.reward == reward) &&
            (identical(
                    other.unvalidatedCurrentReward, unvalidatedCurrentReward) ||
                other.unvalidatedCurrentReward == unvalidatedCurrentReward));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      uuid,
      exitId,
      store,
      email,
      date,
      amountSource,
      amountTarget,
      amountUser,
      currencySource,
      currencyTarget,
      networkCommission,
      status,
      cashback,
      reward,
      unvalidatedCurrentReward);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$CashbackTableModelImplCopyWith<_$CashbackTableModelImpl> get copyWith =>
      __$$CashbackTableModelImplCopyWithImpl<_$CashbackTableModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$CashbackTableModelImplToJson(
      this,
    );
  }
}

abstract class _CashbackTableModel implements CashbackTableModel {
  const factory _CashbackTableModel(
      {required final String uuid,
      final String? exitId,
      final CashbackStoreModel? store,
      final String? email,
      required final String date,
      required final double amountSource,
      required final double amountTarget,
      required final double amountUser,
      required final String currencySource,
      required final String currencyTarget,
      final double? networkCommission,
      required final String status,
      required final double cashback,
      @JsonKey(fromJson: _rewardFromJson, toJson: _rewardToJson)
      final RewardCashbackModel? reward,
      final double? unvalidatedCurrentReward}) = _$CashbackTableModelImpl;

  factory _CashbackTableModel.fromJson(Map<String, dynamic> json) =
      _$CashbackTableModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String? get exitId;
  @override
  CashbackStoreModel? get store;
  @override
  String? get email;
  @override
  String get date;
  @override
  double get amountSource;
  @override
  double get amountTarget;
  @override
  double get amountUser;
  @override
  String get currencySource;
  @override
  String get currencyTarget;
  @override
  double? get networkCommission;
  @override
  String get status;
  @override
  double get cashback;
  @override
  @JsonKey(fromJson: _rewardFromJson, toJson: _rewardToJson)
  RewardCashbackModel? get reward;
  @override
  double? get unvalidatedCurrentReward;
  @override
  @JsonKey(ignore: true)
  _$$CashbackTableModelImplCopyWith<_$CashbackTableModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
