// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'referral_info_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

ReferralInfoModel _$ReferralInfoModelFromJson(Map<String, dynamic> json) {
  return _ReferralInfoModel.fromJson(json);
}

/// @nodoc
mixin _$ReferralInfoModel {
  int? get totalClicks => throw _privateConstructorUsedError;
  int get totalUserRegistered => throw _privateConstructorUsedError;
  int get totalFirstPurchase => throw _privateConstructorUsedError;
  List<MonthDataModel> get clicksByMonth => throw _privateConstructorUsedError;
  List<MonthDataModel> get registeredByMonth =>
      throw _privateConstructorUsedError;
  List<MonthDataModel> get firstPurchaseByMonth =>
      throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $ReferralInfoModelCopyWith<ReferralInfoModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $ReferralInfoModelCopyWith<$Res> {
  factory $ReferralInfoModelCopyWith(
          ReferralInfoModel value, $Res Function(ReferralInfoModel) then) =
      _$ReferralInfoModelCopyWithImpl<$Res, ReferralInfoModel>;
  @useResult
  $Res call(
      {int? totalClicks,
      int totalUserRegistered,
      int totalFirstPurchase,
      List<MonthDataModel> clicksByMonth,
      List<MonthDataModel> registeredByMonth,
      List<MonthDataModel> firstPurchaseByMonth});
}

/// @nodoc
class _$ReferralInfoModelCopyWithImpl<$Res, $Val extends ReferralInfoModel>
    implements $ReferralInfoModelCopyWith<$Res> {
  _$ReferralInfoModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalClicks = freezed,
    Object? totalUserRegistered = null,
    Object? totalFirstPurchase = null,
    Object? clicksByMonth = null,
    Object? registeredByMonth = null,
    Object? firstPurchaseByMonth = null,
  }) {
    return _then(_value.copyWith(
      totalClicks: freezed == totalClicks
          ? _value.totalClicks
          : totalClicks // ignore: cast_nullable_to_non_nullable
              as int?,
      totalUserRegistered: null == totalUserRegistered
          ? _value.totalUserRegistered
          : totalUserRegistered // ignore: cast_nullable_to_non_nullable
              as int,
      totalFirstPurchase: null == totalFirstPurchase
          ? _value.totalFirstPurchase
          : totalFirstPurchase // ignore: cast_nullable_to_non_nullable
              as int,
      clicksByMonth: null == clicksByMonth
          ? _value.clicksByMonth
          : clicksByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
      registeredByMonth: null == registeredByMonth
          ? _value.registeredByMonth
          : registeredByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
      firstPurchaseByMonth: null == firstPurchaseByMonth
          ? _value.firstPurchaseByMonth
          : firstPurchaseByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$ReferralInfoModelImplCopyWith<$Res>
    implements $ReferralInfoModelCopyWith<$Res> {
  factory _$$ReferralInfoModelImplCopyWith(_$ReferralInfoModelImpl value,
          $Res Function(_$ReferralInfoModelImpl) then) =
      __$$ReferralInfoModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {int? totalClicks,
      int totalUserRegistered,
      int totalFirstPurchase,
      List<MonthDataModel> clicksByMonth,
      List<MonthDataModel> registeredByMonth,
      List<MonthDataModel> firstPurchaseByMonth});
}

/// @nodoc
class __$$ReferralInfoModelImplCopyWithImpl<$Res>
    extends _$ReferralInfoModelCopyWithImpl<$Res, _$ReferralInfoModelImpl>
    implements _$$ReferralInfoModelImplCopyWith<$Res> {
  __$$ReferralInfoModelImplCopyWithImpl(_$ReferralInfoModelImpl _value,
      $Res Function(_$ReferralInfoModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalClicks = freezed,
    Object? totalUserRegistered = null,
    Object? totalFirstPurchase = null,
    Object? clicksByMonth = null,
    Object? registeredByMonth = null,
    Object? firstPurchaseByMonth = null,
  }) {
    return _then(_$ReferralInfoModelImpl(
      totalClicks: freezed == totalClicks
          ? _value.totalClicks
          : totalClicks // ignore: cast_nullable_to_non_nullable
              as int?,
      totalUserRegistered: null == totalUserRegistered
          ? _value.totalUserRegistered
          : totalUserRegistered // ignore: cast_nullable_to_non_nullable
              as int,
      totalFirstPurchase: null == totalFirstPurchase
          ? _value.totalFirstPurchase
          : totalFirstPurchase // ignore: cast_nullable_to_non_nullable
              as int,
      clicksByMonth: null == clicksByMonth
          ? _value._clicksByMonth
          : clicksByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
      registeredByMonth: null == registeredByMonth
          ? _value._registeredByMonth
          : registeredByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
      firstPurchaseByMonth: null == firstPurchaseByMonth
          ? _value._firstPurchaseByMonth
          : firstPurchaseByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$ReferralInfoModelImpl implements _ReferralInfoModel {
  const _$ReferralInfoModelImpl(
      {this.totalClicks,
      required this.totalUserRegistered,
      required this.totalFirstPurchase,
      required final List<MonthDataModel> clicksByMonth,
      required final List<MonthDataModel> registeredByMonth,
      required final List<MonthDataModel> firstPurchaseByMonth})
      : _clicksByMonth = clicksByMonth,
        _registeredByMonth = registeredByMonth,
        _firstPurchaseByMonth = firstPurchaseByMonth;

  factory _$ReferralInfoModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$ReferralInfoModelImplFromJson(json);

  @override
  final int? totalClicks;
  @override
  final int totalUserRegistered;
  @override
  final int totalFirstPurchase;
  final List<MonthDataModel> _clicksByMonth;
  @override
  List<MonthDataModel> get clicksByMonth {
    if (_clicksByMonth is EqualUnmodifiableListView) return _clicksByMonth;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_clicksByMonth);
  }

  final List<MonthDataModel> _registeredByMonth;
  @override
  List<MonthDataModel> get registeredByMonth {
    if (_registeredByMonth is EqualUnmodifiableListView)
      return _registeredByMonth;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_registeredByMonth);
  }

  final List<MonthDataModel> _firstPurchaseByMonth;
  @override
  List<MonthDataModel> get firstPurchaseByMonth {
    if (_firstPurchaseByMonth is EqualUnmodifiableListView)
      return _firstPurchaseByMonth;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_firstPurchaseByMonth);
  }

  @override
  String toString() {
    return 'ReferralInfoModel(totalClicks: $totalClicks, totalUserRegistered: $totalUserRegistered, totalFirstPurchase: $totalFirstPurchase, clicksByMonth: $clicksByMonth, registeredByMonth: $registeredByMonth, firstPurchaseByMonth: $firstPurchaseByMonth)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ReferralInfoModelImpl &&
            (identical(other.totalClicks, totalClicks) ||
                other.totalClicks == totalClicks) &&
            (identical(other.totalUserRegistered, totalUserRegistered) ||
                other.totalUserRegistered == totalUserRegistered) &&
            (identical(other.totalFirstPurchase, totalFirstPurchase) ||
                other.totalFirstPurchase == totalFirstPurchase) &&
            const DeepCollectionEquality()
                .equals(other._clicksByMonth, _clicksByMonth) &&
            const DeepCollectionEquality()
                .equals(other._registeredByMonth, _registeredByMonth) &&
            const DeepCollectionEquality()
                .equals(other._firstPurchaseByMonth, _firstPurchaseByMonth));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      totalClicks,
      totalUserRegistered,
      totalFirstPurchase,
      const DeepCollectionEquality().hash(_clicksByMonth),
      const DeepCollectionEquality().hash(_registeredByMonth),
      const DeepCollectionEquality().hash(_firstPurchaseByMonth));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$ReferralInfoModelImplCopyWith<_$ReferralInfoModelImpl> get copyWith =>
      __$$ReferralInfoModelImplCopyWithImpl<_$ReferralInfoModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$ReferralInfoModelImplToJson(
      this,
    );
  }
}

abstract class _ReferralInfoModel implements ReferralInfoModel {
  const factory _ReferralInfoModel(
          {final int? totalClicks,
          required final int totalUserRegistered,
          required final int totalFirstPurchase,
          required final List<MonthDataModel> clicksByMonth,
          required final List<MonthDataModel> registeredByMonth,
          required final List<MonthDataModel> firstPurchaseByMonth}) =
      _$ReferralInfoModelImpl;

  factory _ReferralInfoModel.fromJson(Map<String, dynamic> json) =
      _$ReferralInfoModelImpl.fromJson;

  @override
  int? get totalClicks;
  @override
  int get totalUserRegistered;
  @override
  int get totalFirstPurchase;
  @override
  List<MonthDataModel> get clicksByMonth;
  @override
  List<MonthDataModel> get registeredByMonth;
  @override
  List<MonthDataModel> get firstPurchaseByMonth;
  @override
  @JsonKey(ignore: true)
  _$$ReferralInfoModelImplCopyWith<_$ReferralInfoModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
