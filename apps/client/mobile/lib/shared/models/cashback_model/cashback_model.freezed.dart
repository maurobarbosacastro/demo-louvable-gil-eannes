// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'cashback_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

CashbackModel _$CashbackModelFromJson(Map<String, dynamic> json) {
  return _CashbackModel.fromJson(json);
}

/// @nodoc
mixin _$CashbackModel {
  String get uuid => throw _privateConstructorUsedError;
  double get amountSource => throw _privateConstructorUsedError;
  double get amountTarget => throw _privateConstructorUsedError;
  double get amountUser => throw _privateConstructorUsedError;
  String get currencySource => throw _privateConstructorUsedError;
  String get currencyTarget => throw _privateConstructorUsedError;
  String get state => throw _privateConstructorUsedError;
  double get commissionSource => throw _privateConstructorUsedError;
  double get commissionTarget => throw _privateConstructorUsedError;
  double get commissionUser => throw _privateConstructorUsedError;
  String get orderDate => throw _privateConstructorUsedError;
  String get user => throw _privateConstructorUsedError;
  StoreModel? get store => throw _privateConstructorUsedError;
  StoreVisitCashbackModel? get storeVisit => throw _privateConstructorUsedError;
  String get currencyExchangeRateUuid => throw _privateConstructorUsedError;
  double get cashback => throw _privateConstructorUsedError;
  String get createdAt => throw _privateConstructorUsedError;
  String get createdBy => throw _privateConstructorUsedError;
  String get updatedAt => throw _privateConstructorUsedError;
  String? get updatedBy => throw _privateConstructorUsedError;
  bool? get deleted => throw _privateConstructorUsedError;
  dynamic get deletedAt => throw _privateConstructorUsedError;
  dynamic get deletedBy => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $CashbackModelCopyWith<CashbackModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $CashbackModelCopyWith<$Res> {
  factory $CashbackModelCopyWith(
          CashbackModel value, $Res Function(CashbackModel) then) =
      _$CashbackModelCopyWithImpl<$Res, CashbackModel>;
  @useResult
  $Res call(
      {String uuid,
      double amountSource,
      double amountTarget,
      double amountUser,
      String currencySource,
      String currencyTarget,
      String state,
      double commissionSource,
      double commissionTarget,
      double commissionUser,
      String orderDate,
      String user,
      StoreModel? store,
      StoreVisitCashbackModel? storeVisit,
      String currencyExchangeRateUuid,
      double cashback,
      String createdAt,
      String createdBy,
      String updatedAt,
      String? updatedBy,
      bool? deleted,
      dynamic deletedAt,
      dynamic deletedBy});

  $StoreModelCopyWith<$Res>? get store;
  $StoreVisitCashbackModelCopyWith<$Res>? get storeVisit;
}

/// @nodoc
class _$CashbackModelCopyWithImpl<$Res, $Val extends CashbackModel>
    implements $CashbackModelCopyWith<$Res> {
  _$CashbackModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? amountSource = null,
    Object? amountTarget = null,
    Object? amountUser = null,
    Object? currencySource = null,
    Object? currencyTarget = null,
    Object? state = null,
    Object? commissionSource = null,
    Object? commissionTarget = null,
    Object? commissionUser = null,
    Object? orderDate = null,
    Object? user = null,
    Object? store = freezed,
    Object? storeVisit = freezed,
    Object? currencyExchangeRateUuid = null,
    Object? cashback = null,
    Object? createdAt = null,
    Object? createdBy = null,
    Object? updatedAt = null,
    Object? updatedBy = freezed,
    Object? deleted = freezed,
    Object? deletedAt = freezed,
    Object? deletedBy = freezed,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
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
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      commissionSource: null == commissionSource
          ? _value.commissionSource
          : commissionSource // ignore: cast_nullable_to_non_nullable
              as double,
      commissionTarget: null == commissionTarget
          ? _value.commissionTarget
          : commissionTarget // ignore: cast_nullable_to_non_nullable
              as double,
      commissionUser: null == commissionUser
          ? _value.commissionUser
          : commissionUser // ignore: cast_nullable_to_non_nullable
              as double,
      orderDate: null == orderDate
          ? _value.orderDate
          : orderDate // ignore: cast_nullable_to_non_nullable
              as String,
      user: null == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as String,
      store: freezed == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as StoreModel?,
      storeVisit: freezed == storeVisit
          ? _value.storeVisit
          : storeVisit // ignore: cast_nullable_to_non_nullable
              as StoreVisitCashbackModel?,
      currencyExchangeRateUuid: null == currencyExchangeRateUuid
          ? _value.currencyExchangeRateUuid
          : currencyExchangeRateUuid // ignore: cast_nullable_to_non_nullable
              as String,
      cashback: null == cashback
          ? _value.cashback
          : cashback // ignore: cast_nullable_to_non_nullable
              as double,
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
      updatedBy: freezed == updatedBy
          ? _value.updatedBy
          : updatedBy // ignore: cast_nullable_to_non_nullable
              as String?,
      deleted: freezed == deleted
          ? _value.deleted
          : deleted // ignore: cast_nullable_to_non_nullable
              as bool?,
      deletedAt: freezed == deletedAt
          ? _value.deletedAt
          : deletedAt // ignore: cast_nullable_to_non_nullable
              as dynamic,
      deletedBy: freezed == deletedBy
          ? _value.deletedBy
          : deletedBy // ignore: cast_nullable_to_non_nullable
              as dynamic,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $StoreModelCopyWith<$Res>? get store {
    if (_value.store == null) {
      return null;
    }

    return $StoreModelCopyWith<$Res>(_value.store!, (value) {
      return _then(_value.copyWith(store: value) as $Val);
    });
  }

  @override
  @pragma('vm:prefer-inline')
  $StoreVisitCashbackModelCopyWith<$Res>? get storeVisit {
    if (_value.storeVisit == null) {
      return null;
    }

    return $StoreVisitCashbackModelCopyWith<$Res>(_value.storeVisit!, (value) {
      return _then(_value.copyWith(storeVisit: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$CashbackModelImplCopyWith<$Res>
    implements $CashbackModelCopyWith<$Res> {
  factory _$$CashbackModelImplCopyWith(
          _$CashbackModelImpl value, $Res Function(_$CashbackModelImpl) then) =
      __$$CashbackModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      double amountSource,
      double amountTarget,
      double amountUser,
      String currencySource,
      String currencyTarget,
      String state,
      double commissionSource,
      double commissionTarget,
      double commissionUser,
      String orderDate,
      String user,
      StoreModel? store,
      StoreVisitCashbackModel? storeVisit,
      String currencyExchangeRateUuid,
      double cashback,
      String createdAt,
      String createdBy,
      String updatedAt,
      String? updatedBy,
      bool? deleted,
      dynamic deletedAt,
      dynamic deletedBy});

  @override
  $StoreModelCopyWith<$Res>? get store;
  @override
  $StoreVisitCashbackModelCopyWith<$Res>? get storeVisit;
}

/// @nodoc
class __$$CashbackModelImplCopyWithImpl<$Res>
    extends _$CashbackModelCopyWithImpl<$Res, _$CashbackModelImpl>
    implements _$$CashbackModelImplCopyWith<$Res> {
  __$$CashbackModelImplCopyWithImpl(
      _$CashbackModelImpl _value, $Res Function(_$CashbackModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? amountSource = null,
    Object? amountTarget = null,
    Object? amountUser = null,
    Object? currencySource = null,
    Object? currencyTarget = null,
    Object? state = null,
    Object? commissionSource = null,
    Object? commissionTarget = null,
    Object? commissionUser = null,
    Object? orderDate = null,
    Object? user = null,
    Object? store = freezed,
    Object? storeVisit = freezed,
    Object? currencyExchangeRateUuid = null,
    Object? cashback = null,
    Object? createdAt = null,
    Object? createdBy = null,
    Object? updatedAt = null,
    Object? updatedBy = freezed,
    Object? deleted = freezed,
    Object? deletedAt = freezed,
    Object? deletedBy = freezed,
  }) {
    return _then(_$CashbackModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
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
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      commissionSource: null == commissionSource
          ? _value.commissionSource
          : commissionSource // ignore: cast_nullable_to_non_nullable
              as double,
      commissionTarget: null == commissionTarget
          ? _value.commissionTarget
          : commissionTarget // ignore: cast_nullable_to_non_nullable
              as double,
      commissionUser: null == commissionUser
          ? _value.commissionUser
          : commissionUser // ignore: cast_nullable_to_non_nullable
              as double,
      orderDate: null == orderDate
          ? _value.orderDate
          : orderDate // ignore: cast_nullable_to_non_nullable
              as String,
      user: null == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as String,
      store: freezed == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as StoreModel?,
      storeVisit: freezed == storeVisit
          ? _value.storeVisit
          : storeVisit // ignore: cast_nullable_to_non_nullable
              as StoreVisitCashbackModel?,
      currencyExchangeRateUuid: null == currencyExchangeRateUuid
          ? _value.currencyExchangeRateUuid
          : currencyExchangeRateUuid // ignore: cast_nullable_to_non_nullable
              as String,
      cashback: null == cashback
          ? _value.cashback
          : cashback // ignore: cast_nullable_to_non_nullable
              as double,
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
      updatedBy: freezed == updatedBy
          ? _value.updatedBy
          : updatedBy // ignore: cast_nullable_to_non_nullable
              as String?,
      deleted: freezed == deleted
          ? _value.deleted
          : deleted // ignore: cast_nullable_to_non_nullable
              as bool?,
      deletedAt: freezed == deletedAt
          ? _value.deletedAt
          : deletedAt // ignore: cast_nullable_to_non_nullable
              as dynamic,
      deletedBy: freezed == deletedBy
          ? _value.deletedBy
          : deletedBy // ignore: cast_nullable_to_non_nullable
              as dynamic,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$CashbackModelImpl implements _CashbackModel {
  const _$CashbackModelImpl(
      {required this.uuid,
      required this.amountSource,
      required this.amountTarget,
      required this.amountUser,
      required this.currencySource,
      required this.currencyTarget,
      required this.state,
      required this.commissionSource,
      required this.commissionTarget,
      required this.commissionUser,
      required this.orderDate,
      required this.user,
      this.store,
      this.storeVisit,
      required this.currencyExchangeRateUuid,
      required this.cashback,
      required this.createdAt,
      required this.createdBy,
      required this.updatedAt,
      this.updatedBy,
      this.deleted,
      this.deletedAt,
      this.deletedBy});

  factory _$CashbackModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$CashbackModelImplFromJson(json);

  @override
  final String uuid;
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
  final String state;
  @override
  final double commissionSource;
  @override
  final double commissionTarget;
  @override
  final double commissionUser;
  @override
  final String orderDate;
  @override
  final String user;
  @override
  final StoreModel? store;
  @override
  final StoreVisitCashbackModel? storeVisit;
  @override
  final String currencyExchangeRateUuid;
  @override
  final double cashback;
  @override
  final String createdAt;
  @override
  final String createdBy;
  @override
  final String updatedAt;
  @override
  final String? updatedBy;
  @override
  final bool? deleted;
  @override
  final dynamic deletedAt;
  @override
  final dynamic deletedBy;

  @override
  String toString() {
    return 'CashbackModel(uuid: $uuid, amountSource: $amountSource, amountTarget: $amountTarget, amountUser: $amountUser, currencySource: $currencySource, currencyTarget: $currencyTarget, state: $state, commissionSource: $commissionSource, commissionTarget: $commissionTarget, commissionUser: $commissionUser, orderDate: $orderDate, user: $user, store: $store, storeVisit: $storeVisit, currencyExchangeRateUuid: $currencyExchangeRateUuid, cashback: $cashback, createdAt: $createdAt, createdBy: $createdBy, updatedAt: $updatedAt, updatedBy: $updatedBy, deleted: $deleted, deletedAt: $deletedAt, deletedBy: $deletedBy)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$CashbackModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
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
            (identical(other.state, state) || other.state == state) &&
            (identical(other.commissionSource, commissionSource) ||
                other.commissionSource == commissionSource) &&
            (identical(other.commissionTarget, commissionTarget) ||
                other.commissionTarget == commissionTarget) &&
            (identical(other.commissionUser, commissionUser) ||
                other.commissionUser == commissionUser) &&
            (identical(other.orderDate, orderDate) ||
                other.orderDate == orderDate) &&
            (identical(other.user, user) || other.user == user) &&
            (identical(other.store, store) || other.store == store) &&
            (identical(other.storeVisit, storeVisit) ||
                other.storeVisit == storeVisit) &&
            (identical(
                    other.currencyExchangeRateUuid, currencyExchangeRateUuid) ||
                other.currencyExchangeRateUuid == currencyExchangeRateUuid) &&
            (identical(other.cashback, cashback) ||
                other.cashback == cashback) &&
            (identical(other.createdAt, createdAt) ||
                other.createdAt == createdAt) &&
            (identical(other.createdBy, createdBy) ||
                other.createdBy == createdBy) &&
            (identical(other.updatedAt, updatedAt) ||
                other.updatedAt == updatedAt) &&
            (identical(other.updatedBy, updatedBy) ||
                other.updatedBy == updatedBy) &&
            (identical(other.deleted, deleted) || other.deleted == deleted) &&
            const DeepCollectionEquality().equals(other.deletedAt, deletedAt) &&
            const DeepCollectionEquality().equals(other.deletedBy, deletedBy));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hashAll([
        runtimeType,
        uuid,
        amountSource,
        amountTarget,
        amountUser,
        currencySource,
        currencyTarget,
        state,
        commissionSource,
        commissionTarget,
        commissionUser,
        orderDate,
        user,
        store,
        storeVisit,
        currencyExchangeRateUuid,
        cashback,
        createdAt,
        createdBy,
        updatedAt,
        updatedBy,
        deleted,
        const DeepCollectionEquality().hash(deletedAt),
        const DeepCollectionEquality().hash(deletedBy)
      ]);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$CashbackModelImplCopyWith<_$CashbackModelImpl> get copyWith =>
      __$$CashbackModelImplCopyWithImpl<_$CashbackModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$CashbackModelImplToJson(
      this,
    );
  }
}

abstract class _CashbackModel implements CashbackModel {
  const factory _CashbackModel(
      {required final String uuid,
      required final double amountSource,
      required final double amountTarget,
      required final double amountUser,
      required final String currencySource,
      required final String currencyTarget,
      required final String state,
      required final double commissionSource,
      required final double commissionTarget,
      required final double commissionUser,
      required final String orderDate,
      required final String user,
      final StoreModel? store,
      final StoreVisitCashbackModel? storeVisit,
      required final String currencyExchangeRateUuid,
      required final double cashback,
      required final String createdAt,
      required final String createdBy,
      required final String updatedAt,
      final String? updatedBy,
      final bool? deleted,
      final dynamic deletedAt,
      final dynamic deletedBy}) = _$CashbackModelImpl;

  factory _CashbackModel.fromJson(Map<String, dynamic> json) =
      _$CashbackModelImpl.fromJson;

  @override
  String get uuid;
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
  String get state;
  @override
  double get commissionSource;
  @override
  double get commissionTarget;
  @override
  double get commissionUser;
  @override
  String get orderDate;
  @override
  String get user;
  @override
  StoreModel? get store;
  @override
  StoreVisitCashbackModel? get storeVisit;
  @override
  String get currencyExchangeRateUuid;
  @override
  double get cashback;
  @override
  String get createdAt;
  @override
  String get createdBy;
  @override
  String get updatedAt;
  @override
  String? get updatedBy;
  @override
  bool? get deleted;
  @override
  dynamic get deletedAt;
  @override
  dynamic get deletedBy;
  @override
  @JsonKey(ignore: true)
  _$$CashbackModelImplCopyWith<_$CashbackModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
