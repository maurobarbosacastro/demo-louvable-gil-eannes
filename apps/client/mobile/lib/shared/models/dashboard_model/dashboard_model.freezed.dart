// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'dashboard_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

DashboardModel _$DashboardModelFromJson(Map<String, dynamic> json) {
  return _DashboardModel.fromJson(json);
}

/// @nodoc
mixin _$DashboardModel {
  CashbackSection? get cashbackSection => throw _privateConstructorUsedError;
  IndicatorsSection? get indicatorsSection =>
      throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $DashboardModelCopyWith<DashboardModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $DashboardModelCopyWith<$Res> {
  factory $DashboardModelCopyWith(
          DashboardModel value, $Res Function(DashboardModel) then) =
      _$DashboardModelCopyWithImpl<$Res, DashboardModel>;
  @useResult
  $Res call(
      {CashbackSection? cashbackSection, IndicatorsSection? indicatorsSection});

  $CashbackSectionCopyWith<$Res>? get cashbackSection;
  $IndicatorsSectionCopyWith<$Res>? get indicatorsSection;
}

/// @nodoc
class _$DashboardModelCopyWithImpl<$Res, $Val extends DashboardModel>
    implements $DashboardModelCopyWith<$Res> {
  _$DashboardModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? cashbackSection = freezed,
    Object? indicatorsSection = freezed,
  }) {
    return _then(_value.copyWith(
      cashbackSection: freezed == cashbackSection
          ? _value.cashbackSection
          : cashbackSection // ignore: cast_nullable_to_non_nullable
              as CashbackSection?,
      indicatorsSection: freezed == indicatorsSection
          ? _value.indicatorsSection
          : indicatorsSection // ignore: cast_nullable_to_non_nullable
              as IndicatorsSection?,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $CashbackSectionCopyWith<$Res>? get cashbackSection {
    if (_value.cashbackSection == null) {
      return null;
    }

    return $CashbackSectionCopyWith<$Res>(_value.cashbackSection!, (value) {
      return _then(_value.copyWith(cashbackSection: value) as $Val);
    });
  }

  @override
  @pragma('vm:prefer-inline')
  $IndicatorsSectionCopyWith<$Res>? get indicatorsSection {
    if (_value.indicatorsSection == null) {
      return null;
    }

    return $IndicatorsSectionCopyWith<$Res>(_value.indicatorsSection!, (value) {
      return _then(_value.copyWith(indicatorsSection: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$DashboardModelImplCopyWith<$Res>
    implements $DashboardModelCopyWith<$Res> {
  factory _$$DashboardModelImplCopyWith(_$DashboardModelImpl value,
          $Res Function(_$DashboardModelImpl) then) =
      __$$DashboardModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {CashbackSection? cashbackSection, IndicatorsSection? indicatorsSection});

  @override
  $CashbackSectionCopyWith<$Res>? get cashbackSection;
  @override
  $IndicatorsSectionCopyWith<$Res>? get indicatorsSection;
}

/// @nodoc
class __$$DashboardModelImplCopyWithImpl<$Res>
    extends _$DashboardModelCopyWithImpl<$Res, _$DashboardModelImpl>
    implements _$$DashboardModelImplCopyWith<$Res> {
  __$$DashboardModelImplCopyWithImpl(
      _$DashboardModelImpl _value, $Res Function(_$DashboardModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? cashbackSection = freezed,
    Object? indicatorsSection = freezed,
  }) {
    return _then(_$DashboardModelImpl(
      cashbackSection: freezed == cashbackSection
          ? _value.cashbackSection
          : cashbackSection // ignore: cast_nullable_to_non_nullable
              as CashbackSection?,
      indicatorsSection: freezed == indicatorsSection
          ? _value.indicatorsSection
          : indicatorsSection // ignore: cast_nullable_to_non_nullable
              as IndicatorsSection?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$DashboardModelImpl implements _DashboardModel {
  const _$DashboardModelImpl({this.cashbackSection, this.indicatorsSection});

  factory _$DashboardModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$DashboardModelImplFromJson(json);

  @override
  final CashbackSection? cashbackSection;
  @override
  final IndicatorsSection? indicatorsSection;

  @override
  String toString() {
    return 'DashboardModel(cashbackSection: $cashbackSection, indicatorsSection: $indicatorsSection)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$DashboardModelImpl &&
            (identical(other.cashbackSection, cashbackSection) ||
                other.cashbackSection == cashbackSection) &&
            (identical(other.indicatorsSection, indicatorsSection) ||
                other.indicatorsSection == indicatorsSection));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode =>
      Object.hash(runtimeType, cashbackSection, indicatorsSection);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$DashboardModelImplCopyWith<_$DashboardModelImpl> get copyWith =>
      __$$DashboardModelImplCopyWithImpl<_$DashboardModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$DashboardModelImplToJson(
      this,
    );
  }
}

abstract class _DashboardModel implements DashboardModel {
  const factory _DashboardModel(
      {final CashbackSection? cashbackSection,
      final IndicatorsSection? indicatorsSection}) = _$DashboardModelImpl;

  factory _DashboardModel.fromJson(Map<String, dynamic> json) =
      _$DashboardModelImpl.fromJson;

  @override
  CashbackSection? get cashbackSection;
  @override
  IndicatorsSection? get indicatorsSection;
  @override
  @JsonKey(ignore: true)
  _$$DashboardModelImplCopyWith<_$DashboardModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

CashbackSection _$CashbackSectionFromJson(Map<String, dynamic> json) {
  return _CashbackSection.fromJson(json);
}

/// @nodoc
mixin _$CashbackSection {
  int get totalValidatedCashbacks => throw _privateConstructorUsedError;
  int get totalStoppedCashbacks => throw _privateConstructorUsedError;
  int get totalPaidCashbacks => throw _privateConstructorUsedError;
  int get totalRequestedCashbacks => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $CashbackSectionCopyWith<CashbackSection> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $CashbackSectionCopyWith<$Res> {
  factory $CashbackSectionCopyWith(
          CashbackSection value, $Res Function(CashbackSection) then) =
      _$CashbackSectionCopyWithImpl<$Res, CashbackSection>;
  @useResult
  $Res call(
      {int totalValidatedCashbacks,
      int totalStoppedCashbacks,
      int totalPaidCashbacks,
      int totalRequestedCashbacks});
}

/// @nodoc
class _$CashbackSectionCopyWithImpl<$Res, $Val extends CashbackSection>
    implements $CashbackSectionCopyWith<$Res> {
  _$CashbackSectionCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalValidatedCashbacks = null,
    Object? totalStoppedCashbacks = null,
    Object? totalPaidCashbacks = null,
    Object? totalRequestedCashbacks = null,
  }) {
    return _then(_value.copyWith(
      totalValidatedCashbacks: null == totalValidatedCashbacks
          ? _value.totalValidatedCashbacks
          : totalValidatedCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
      totalStoppedCashbacks: null == totalStoppedCashbacks
          ? _value.totalStoppedCashbacks
          : totalStoppedCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
      totalPaidCashbacks: null == totalPaidCashbacks
          ? _value.totalPaidCashbacks
          : totalPaidCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
      totalRequestedCashbacks: null == totalRequestedCashbacks
          ? _value.totalRequestedCashbacks
          : totalRequestedCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$CashbackSectionImplCopyWith<$Res>
    implements $CashbackSectionCopyWith<$Res> {
  factory _$$CashbackSectionImplCopyWith(_$CashbackSectionImpl value,
          $Res Function(_$CashbackSectionImpl) then) =
      __$$CashbackSectionImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {int totalValidatedCashbacks,
      int totalStoppedCashbacks,
      int totalPaidCashbacks,
      int totalRequestedCashbacks});
}

/// @nodoc
class __$$CashbackSectionImplCopyWithImpl<$Res>
    extends _$CashbackSectionCopyWithImpl<$Res, _$CashbackSectionImpl>
    implements _$$CashbackSectionImplCopyWith<$Res> {
  __$$CashbackSectionImplCopyWithImpl(
      _$CashbackSectionImpl _value, $Res Function(_$CashbackSectionImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalValidatedCashbacks = null,
    Object? totalStoppedCashbacks = null,
    Object? totalPaidCashbacks = null,
    Object? totalRequestedCashbacks = null,
  }) {
    return _then(_$CashbackSectionImpl(
      totalValidatedCashbacks: null == totalValidatedCashbacks
          ? _value.totalValidatedCashbacks
          : totalValidatedCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
      totalStoppedCashbacks: null == totalStoppedCashbacks
          ? _value.totalStoppedCashbacks
          : totalStoppedCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
      totalPaidCashbacks: null == totalPaidCashbacks
          ? _value.totalPaidCashbacks
          : totalPaidCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
      totalRequestedCashbacks: null == totalRequestedCashbacks
          ? _value.totalRequestedCashbacks
          : totalRequestedCashbacks // ignore: cast_nullable_to_non_nullable
              as int,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$CashbackSectionImpl implements _CashbackSection {
  const _$CashbackSectionImpl(
      {required this.totalValidatedCashbacks,
      required this.totalStoppedCashbacks,
      required this.totalPaidCashbacks,
      required this.totalRequestedCashbacks});

  factory _$CashbackSectionImpl.fromJson(Map<String, dynamic> json) =>
      _$$CashbackSectionImplFromJson(json);

  @override
  final int totalValidatedCashbacks;
  @override
  final int totalStoppedCashbacks;
  @override
  final int totalPaidCashbacks;
  @override
  final int totalRequestedCashbacks;

  @override
  String toString() {
    return 'CashbackSection(totalValidatedCashbacks: $totalValidatedCashbacks, totalStoppedCashbacks: $totalStoppedCashbacks, totalPaidCashbacks: $totalPaidCashbacks, totalRequestedCashbacks: $totalRequestedCashbacks)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$CashbackSectionImpl &&
            (identical(
                    other.totalValidatedCashbacks, totalValidatedCashbacks) ||
                other.totalValidatedCashbacks == totalValidatedCashbacks) &&
            (identical(other.totalStoppedCashbacks, totalStoppedCashbacks) ||
                other.totalStoppedCashbacks == totalStoppedCashbacks) &&
            (identical(other.totalPaidCashbacks, totalPaidCashbacks) ||
                other.totalPaidCashbacks == totalPaidCashbacks) &&
            (identical(
                    other.totalRequestedCashbacks, totalRequestedCashbacks) ||
                other.totalRequestedCashbacks == totalRequestedCashbacks));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, totalValidatedCashbacks,
      totalStoppedCashbacks, totalPaidCashbacks, totalRequestedCashbacks);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$CashbackSectionImplCopyWith<_$CashbackSectionImpl> get copyWith =>
      __$$CashbackSectionImplCopyWithImpl<_$CashbackSectionImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$CashbackSectionImplToJson(
      this,
    );
  }
}

abstract class _CashbackSection implements CashbackSection {
  const factory _CashbackSection(
      {required final int totalValidatedCashbacks,
      required final int totalStoppedCashbacks,
      required final int totalPaidCashbacks,
      required final int totalRequestedCashbacks}) = _$CashbackSectionImpl;

  factory _CashbackSection.fromJson(Map<String, dynamic> json) =
      _$CashbackSectionImpl.fromJson;

  @override
  int get totalValidatedCashbacks;
  @override
  int get totalStoppedCashbacks;
  @override
  int get totalPaidCashbacks;
  @override
  int get totalRequestedCashbacks;
  @override
  @JsonKey(ignore: true)
  _$$CashbackSectionImplCopyWith<_$CashbackSectionImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

IndicatorsSection _$IndicatorsSectionFromJson(Map<String, dynamic> json) {
  return _IndicatorsSection.fromJson(json);
}

/// @nodoc
mixin _$IndicatorsSection {
  dynamic get totalUsers => throw _privateConstructorUsedError;
  dynamic get activeUsers => throw _privateConstructorUsedError;
  dynamic get numTransactions => throw _privateConstructorUsedError;
  dynamic get totalGMV => throw _privateConstructorUsedError;
  dynamic get averageTransactionAmount => throw _privateConstructorUsedError;
  dynamic get totalRevenue => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $IndicatorsSectionCopyWith<IndicatorsSection> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $IndicatorsSectionCopyWith<$Res> {
  factory $IndicatorsSectionCopyWith(
          IndicatorsSection value, $Res Function(IndicatorsSection) then) =
      _$IndicatorsSectionCopyWithImpl<$Res, IndicatorsSection>;
  @useResult
  $Res call(
      {dynamic totalUsers,
      dynamic activeUsers,
      dynamic numTransactions,
      dynamic totalGMV,
      dynamic averageTransactionAmount,
      dynamic totalRevenue});
}

/// @nodoc
class _$IndicatorsSectionCopyWithImpl<$Res, $Val extends IndicatorsSection>
    implements $IndicatorsSectionCopyWith<$Res> {
  _$IndicatorsSectionCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalUsers = freezed,
    Object? activeUsers = freezed,
    Object? numTransactions = freezed,
    Object? totalGMV = freezed,
    Object? averageTransactionAmount = freezed,
    Object? totalRevenue = freezed,
  }) {
    return _then(_value.copyWith(
      totalUsers: freezed == totalUsers
          ? _value.totalUsers
          : totalUsers // ignore: cast_nullable_to_non_nullable
              as dynamic,
      activeUsers: freezed == activeUsers
          ? _value.activeUsers
          : activeUsers // ignore: cast_nullable_to_non_nullable
              as dynamic,
      numTransactions: freezed == numTransactions
          ? _value.numTransactions
          : numTransactions // ignore: cast_nullable_to_non_nullable
              as dynamic,
      totalGMV: freezed == totalGMV
          ? _value.totalGMV
          : totalGMV // ignore: cast_nullable_to_non_nullable
              as dynamic,
      averageTransactionAmount: freezed == averageTransactionAmount
          ? _value.averageTransactionAmount
          : averageTransactionAmount // ignore: cast_nullable_to_non_nullable
              as dynamic,
      totalRevenue: freezed == totalRevenue
          ? _value.totalRevenue
          : totalRevenue // ignore: cast_nullable_to_non_nullable
              as dynamic,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$IndicatorsSectionImplCopyWith<$Res>
    implements $IndicatorsSectionCopyWith<$Res> {
  factory _$$IndicatorsSectionImplCopyWith(_$IndicatorsSectionImpl value,
          $Res Function(_$IndicatorsSectionImpl) then) =
      __$$IndicatorsSectionImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {dynamic totalUsers,
      dynamic activeUsers,
      dynamic numTransactions,
      dynamic totalGMV,
      dynamic averageTransactionAmount,
      dynamic totalRevenue});
}

/// @nodoc
class __$$IndicatorsSectionImplCopyWithImpl<$Res>
    extends _$IndicatorsSectionCopyWithImpl<$Res, _$IndicatorsSectionImpl>
    implements _$$IndicatorsSectionImplCopyWith<$Res> {
  __$$IndicatorsSectionImplCopyWithImpl(_$IndicatorsSectionImpl _value,
      $Res Function(_$IndicatorsSectionImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalUsers = freezed,
    Object? activeUsers = freezed,
    Object? numTransactions = freezed,
    Object? totalGMV = freezed,
    Object? averageTransactionAmount = freezed,
    Object? totalRevenue = freezed,
  }) {
    return _then(_$IndicatorsSectionImpl(
      totalUsers: freezed == totalUsers
          ? _value.totalUsers
          : totalUsers // ignore: cast_nullable_to_non_nullable
              as dynamic,
      activeUsers: freezed == activeUsers
          ? _value.activeUsers
          : activeUsers // ignore: cast_nullable_to_non_nullable
              as dynamic,
      numTransactions: freezed == numTransactions
          ? _value.numTransactions
          : numTransactions // ignore: cast_nullable_to_non_nullable
              as dynamic,
      totalGMV: freezed == totalGMV
          ? _value.totalGMV
          : totalGMV // ignore: cast_nullable_to_non_nullable
              as dynamic,
      averageTransactionAmount: freezed == averageTransactionAmount
          ? _value.averageTransactionAmount
          : averageTransactionAmount // ignore: cast_nullable_to_non_nullable
              as dynamic,
      totalRevenue: freezed == totalRevenue
          ? _value.totalRevenue
          : totalRevenue // ignore: cast_nullable_to_non_nullable
              as dynamic,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$IndicatorsSectionImpl implements _IndicatorsSection {
  const _$IndicatorsSectionImpl(
      {required this.totalUsers,
      required this.activeUsers,
      required this.numTransactions,
      required this.totalGMV,
      required this.averageTransactionAmount,
      required this.totalRevenue});

  factory _$IndicatorsSectionImpl.fromJson(Map<String, dynamic> json) =>
      _$$IndicatorsSectionImplFromJson(json);

  @override
  final dynamic totalUsers;
  @override
  final dynamic activeUsers;
  @override
  final dynamic numTransactions;
  @override
  final dynamic totalGMV;
  @override
  final dynamic averageTransactionAmount;
  @override
  final dynamic totalRevenue;

  @override
  String toString() {
    return 'IndicatorsSection(totalUsers: $totalUsers, activeUsers: $activeUsers, numTransactions: $numTransactions, totalGMV: $totalGMV, averageTransactionAmount: $averageTransactionAmount, totalRevenue: $totalRevenue)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$IndicatorsSectionImpl &&
            const DeepCollectionEquality()
                .equals(other.totalUsers, totalUsers) &&
            const DeepCollectionEquality()
                .equals(other.activeUsers, activeUsers) &&
            const DeepCollectionEquality()
                .equals(other.numTransactions, numTransactions) &&
            const DeepCollectionEquality().equals(other.totalGMV, totalGMV) &&
            const DeepCollectionEquality().equals(
                other.averageTransactionAmount, averageTransactionAmount) &&
            const DeepCollectionEquality()
                .equals(other.totalRevenue, totalRevenue));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(totalUsers),
      const DeepCollectionEquality().hash(activeUsers),
      const DeepCollectionEquality().hash(numTransactions),
      const DeepCollectionEquality().hash(totalGMV),
      const DeepCollectionEquality().hash(averageTransactionAmount),
      const DeepCollectionEquality().hash(totalRevenue));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$IndicatorsSectionImplCopyWith<_$IndicatorsSectionImpl> get copyWith =>
      __$$IndicatorsSectionImplCopyWithImpl<_$IndicatorsSectionImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$IndicatorsSectionImplToJson(
      this,
    );
  }
}

abstract class _IndicatorsSection implements IndicatorsSection {
  const factory _IndicatorsSection(
      {required final dynamic totalUsers,
      required final dynamic activeUsers,
      required final dynamic numTransactions,
      required final dynamic totalGMV,
      required final dynamic averageTransactionAmount,
      required final dynamic totalRevenue}) = _$IndicatorsSectionImpl;

  factory _IndicatorsSection.fromJson(Map<String, dynamic> json) =
      _$IndicatorsSectionImpl.fromJson;

  @override
  dynamic get totalUsers;
  @override
  dynamic get activeUsers;
  @override
  dynamic get numTransactions;
  @override
  dynamic get totalGMV;
  @override
  dynamic get averageTransactionAmount;
  @override
  dynamic get totalRevenue;
  @override
  @JsonKey(ignore: true)
  _$$IndicatorsSectionImplCopyWith<_$IndicatorsSectionImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

MonthCountsModel _$MonthCountsModelFromJson(Map<String, dynamic> json) {
  return _MonthCountsModel.fromJson(json);
}

/// @nodoc
mixin _$MonthCountsModel {
  double? get totalUsers => throw _privateConstructorUsedError;
  double? get activeUsers => throw _privateConstructorUsedError;
  double? get numTransaction => throw _privateConstructorUsedError;
  double? get totalGMV => throw _privateConstructorUsedError;
  double? get avgTransactionAmount => throw _privateConstructorUsedError;
  double? get totalRevenue => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $MonthCountsModelCopyWith<MonthCountsModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $MonthCountsModelCopyWith<$Res> {
  factory $MonthCountsModelCopyWith(
          MonthCountsModel value, $Res Function(MonthCountsModel) then) =
      _$MonthCountsModelCopyWithImpl<$Res, MonthCountsModel>;
  @useResult
  $Res call(
      {double? totalUsers,
      double? activeUsers,
      double? numTransaction,
      double? totalGMV,
      double? avgTransactionAmount,
      double? totalRevenue});
}

/// @nodoc
class _$MonthCountsModelCopyWithImpl<$Res, $Val extends MonthCountsModel>
    implements $MonthCountsModelCopyWith<$Res> {
  _$MonthCountsModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalUsers = freezed,
    Object? activeUsers = freezed,
    Object? numTransaction = freezed,
    Object? totalGMV = freezed,
    Object? avgTransactionAmount = freezed,
    Object? totalRevenue = freezed,
  }) {
    return _then(_value.copyWith(
      totalUsers: freezed == totalUsers
          ? _value.totalUsers
          : totalUsers // ignore: cast_nullable_to_non_nullable
              as double?,
      activeUsers: freezed == activeUsers
          ? _value.activeUsers
          : activeUsers // ignore: cast_nullable_to_non_nullable
              as double?,
      numTransaction: freezed == numTransaction
          ? _value.numTransaction
          : numTransaction // ignore: cast_nullable_to_non_nullable
              as double?,
      totalGMV: freezed == totalGMV
          ? _value.totalGMV
          : totalGMV // ignore: cast_nullable_to_non_nullable
              as double?,
      avgTransactionAmount: freezed == avgTransactionAmount
          ? _value.avgTransactionAmount
          : avgTransactionAmount // ignore: cast_nullable_to_non_nullable
              as double?,
      totalRevenue: freezed == totalRevenue
          ? _value.totalRevenue
          : totalRevenue // ignore: cast_nullable_to_non_nullable
              as double?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$MonthCountsModelImplCopyWith<$Res>
    implements $MonthCountsModelCopyWith<$Res> {
  factory _$$MonthCountsModelImplCopyWith(_$MonthCountsModelImpl value,
          $Res Function(_$MonthCountsModelImpl) then) =
      __$$MonthCountsModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {double? totalUsers,
      double? activeUsers,
      double? numTransaction,
      double? totalGMV,
      double? avgTransactionAmount,
      double? totalRevenue});
}

/// @nodoc
class __$$MonthCountsModelImplCopyWithImpl<$Res>
    extends _$MonthCountsModelCopyWithImpl<$Res, _$MonthCountsModelImpl>
    implements _$$MonthCountsModelImplCopyWith<$Res> {
  __$$MonthCountsModelImplCopyWithImpl(_$MonthCountsModelImpl _value,
      $Res Function(_$MonthCountsModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalUsers = freezed,
    Object? activeUsers = freezed,
    Object? numTransaction = freezed,
    Object? totalGMV = freezed,
    Object? avgTransactionAmount = freezed,
    Object? totalRevenue = freezed,
  }) {
    return _then(_$MonthCountsModelImpl(
      totalUsers: freezed == totalUsers
          ? _value.totalUsers
          : totalUsers // ignore: cast_nullable_to_non_nullable
              as double?,
      activeUsers: freezed == activeUsers
          ? _value.activeUsers
          : activeUsers // ignore: cast_nullable_to_non_nullable
              as double?,
      numTransaction: freezed == numTransaction
          ? _value.numTransaction
          : numTransaction // ignore: cast_nullable_to_non_nullable
              as double?,
      totalGMV: freezed == totalGMV
          ? _value.totalGMV
          : totalGMV // ignore: cast_nullable_to_non_nullable
              as double?,
      avgTransactionAmount: freezed == avgTransactionAmount
          ? _value.avgTransactionAmount
          : avgTransactionAmount // ignore: cast_nullable_to_non_nullable
              as double?,
      totalRevenue: freezed == totalRevenue
          ? _value.totalRevenue
          : totalRevenue // ignore: cast_nullable_to_non_nullable
              as double?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$MonthCountsModelImpl implements _MonthCountsModel {
  const _$MonthCountsModelImpl(
      {this.totalUsers,
      this.activeUsers,
      this.numTransaction,
      this.totalGMV,
      this.avgTransactionAmount,
      this.totalRevenue});

  factory _$MonthCountsModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$MonthCountsModelImplFromJson(json);

  @override
  final double? totalUsers;
  @override
  final double? activeUsers;
  @override
  final double? numTransaction;
  @override
  final double? totalGMV;
  @override
  final double? avgTransactionAmount;
  @override
  final double? totalRevenue;

  @override
  String toString() {
    return 'MonthCountsModel(totalUsers: $totalUsers, activeUsers: $activeUsers, numTransaction: $numTransaction, totalGMV: $totalGMV, avgTransactionAmount: $avgTransactionAmount, totalRevenue: $totalRevenue)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$MonthCountsModelImpl &&
            (identical(other.totalUsers, totalUsers) ||
                other.totalUsers == totalUsers) &&
            (identical(other.activeUsers, activeUsers) ||
                other.activeUsers == activeUsers) &&
            (identical(other.numTransaction, numTransaction) ||
                other.numTransaction == numTransaction) &&
            (identical(other.totalGMV, totalGMV) ||
                other.totalGMV == totalGMV) &&
            (identical(other.avgTransactionAmount, avgTransactionAmount) ||
                other.avgTransactionAmount == avgTransactionAmount) &&
            (identical(other.totalRevenue, totalRevenue) ||
                other.totalRevenue == totalRevenue));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, totalUsers, activeUsers,
      numTransaction, totalGMV, avgTransactionAmount, totalRevenue);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$MonthCountsModelImplCopyWith<_$MonthCountsModelImpl> get copyWith =>
      __$$MonthCountsModelImplCopyWithImpl<_$MonthCountsModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$MonthCountsModelImplToJson(
      this,
    );
  }
}

abstract class _MonthCountsModel implements MonthCountsModel {
  const factory _MonthCountsModel(
      {final double? totalUsers,
      final double? activeUsers,
      final double? numTransaction,
      final double? totalGMV,
      final double? avgTransactionAmount,
      final double? totalRevenue}) = _$MonthCountsModelImpl;

  factory _MonthCountsModel.fromJson(Map<String, dynamic> json) =
      _$MonthCountsModelImpl.fromJson;

  @override
  double? get totalUsers;
  @override
  double? get activeUsers;
  @override
  double? get numTransaction;
  @override
  double? get totalGMV;
  @override
  double? get avgTransactionAmount;
  @override
  double? get totalRevenue;
  @override
  @JsonKey(ignore: true)
  _$$MonthCountsModelImplCopyWith<_$MonthCountsModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

RewardByCurrencies _$RewardByCurrenciesFromJson(Map<String, dynamic> json) {
  return _RewardByCurrencies.fromJson(json);
}

/// @nodoc
mixin _$RewardByCurrencies {
  String get currency => throw _privateConstructorUsedError;
  String get state => throw _privateConstructorUsedError;
  num get totalRewards => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $RewardByCurrenciesCopyWith<RewardByCurrencies> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $RewardByCurrenciesCopyWith<$Res> {
  factory $RewardByCurrenciesCopyWith(
          RewardByCurrencies value, $Res Function(RewardByCurrencies) then) =
      _$RewardByCurrenciesCopyWithImpl<$Res, RewardByCurrencies>;
  @useResult
  $Res call({String currency, String state, num totalRewards});
}

/// @nodoc
class _$RewardByCurrenciesCopyWithImpl<$Res, $Val extends RewardByCurrencies>
    implements $RewardByCurrenciesCopyWith<$Res> {
  _$RewardByCurrenciesCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? currency = null,
    Object? state = null,
    Object? totalRewards = null,
  }) {
    return _then(_value.copyWith(
      currency: null == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String,
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      totalRewards: null == totalRewards
          ? _value.totalRewards
          : totalRewards // ignore: cast_nullable_to_non_nullable
              as num,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$RewardByCurrenciesImplCopyWith<$Res>
    implements $RewardByCurrenciesCopyWith<$Res> {
  factory _$$RewardByCurrenciesImplCopyWith(_$RewardByCurrenciesImpl value,
          $Res Function(_$RewardByCurrenciesImpl) then) =
      __$$RewardByCurrenciesImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String currency, String state, num totalRewards});
}

/// @nodoc
class __$$RewardByCurrenciesImplCopyWithImpl<$Res>
    extends _$RewardByCurrenciesCopyWithImpl<$Res, _$RewardByCurrenciesImpl>
    implements _$$RewardByCurrenciesImplCopyWith<$Res> {
  __$$RewardByCurrenciesImplCopyWithImpl(_$RewardByCurrenciesImpl _value,
      $Res Function(_$RewardByCurrenciesImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? currency = null,
    Object? state = null,
    Object? totalRewards = null,
  }) {
    return _then(_$RewardByCurrenciesImpl(
      currency: null == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String,
      state: null == state
          ? _value.state
          : state // ignore: cast_nullable_to_non_nullable
              as String,
      totalRewards: null == totalRewards
          ? _value.totalRewards
          : totalRewards // ignore: cast_nullable_to_non_nullable
              as num,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$RewardByCurrenciesImpl implements _RewardByCurrencies {
  const _$RewardByCurrenciesImpl(
      {required this.currency,
      required this.state,
      required this.totalRewards});

  factory _$RewardByCurrenciesImpl.fromJson(Map<String, dynamic> json) =>
      _$$RewardByCurrenciesImplFromJson(json);

  @override
  final String currency;
  @override
  final String state;
  @override
  final num totalRewards;

  @override
  String toString() {
    return 'RewardByCurrencies(currency: $currency, state: $state, totalRewards: $totalRewards)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$RewardByCurrenciesImpl &&
            (identical(other.currency, currency) ||
                other.currency == currency) &&
            (identical(other.state, state) || other.state == state) &&
            (identical(other.totalRewards, totalRewards) ||
                other.totalRewards == totalRewards));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, currency, state, totalRewards);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$RewardByCurrenciesImplCopyWith<_$RewardByCurrenciesImpl> get copyWith =>
      __$$RewardByCurrenciesImplCopyWithImpl<_$RewardByCurrenciesImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$RewardByCurrenciesImplToJson(
      this,
    );
  }
}

abstract class _RewardByCurrencies implements RewardByCurrencies {
  const factory _RewardByCurrencies(
      {required final String currency,
      required final String state,
      required final num totalRewards}) = _$RewardByCurrenciesImpl;

  factory _RewardByCurrencies.fromJson(Map<String, dynamic> json) =
      _$RewardByCurrenciesImpl.fromJson;

  @override
  String get currency;
  @override
  String get state;
  @override
  num get totalRewards;
  @override
  @JsonKey(ignore: true)
  _$$RewardByCurrenciesImplCopyWith<_$RewardByCurrenciesImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
