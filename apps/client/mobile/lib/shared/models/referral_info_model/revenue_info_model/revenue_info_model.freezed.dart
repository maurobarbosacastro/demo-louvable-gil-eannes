// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'revenue_info_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

RevenueInfoModel _$RevenueInfoModelFromJson(Map<String, dynamic> json) {
  return _RevenueInfoModel.fromJson(json);
}

/// @nodoc
mixin _$RevenueInfoModel {
  int get totalRevenue => throw _privateConstructorUsedError;
  List<MonthDataModel> get revenueByMonth => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $RevenueInfoModelCopyWith<RevenueInfoModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $RevenueInfoModelCopyWith<$Res> {
  factory $RevenueInfoModelCopyWith(
          RevenueInfoModel value, $Res Function(RevenueInfoModel) then) =
      _$RevenueInfoModelCopyWithImpl<$Res, RevenueInfoModel>;
  @useResult
  $Res call({int totalRevenue, List<MonthDataModel> revenueByMonth});
}

/// @nodoc
class _$RevenueInfoModelCopyWithImpl<$Res, $Val extends RevenueInfoModel>
    implements $RevenueInfoModelCopyWith<$Res> {
  _$RevenueInfoModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalRevenue = null,
    Object? revenueByMonth = null,
  }) {
    return _then(_value.copyWith(
      totalRevenue: null == totalRevenue
          ? _value.totalRevenue
          : totalRevenue // ignore: cast_nullable_to_non_nullable
              as int,
      revenueByMonth: null == revenueByMonth
          ? _value.revenueByMonth
          : revenueByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$RevenueInfoModelImplCopyWith<$Res>
    implements $RevenueInfoModelCopyWith<$Res> {
  factory _$$RevenueInfoModelImplCopyWith(_$RevenueInfoModelImpl value,
          $Res Function(_$RevenueInfoModelImpl) then) =
      __$$RevenueInfoModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({int totalRevenue, List<MonthDataModel> revenueByMonth});
}

/// @nodoc
class __$$RevenueInfoModelImplCopyWithImpl<$Res>
    extends _$RevenueInfoModelCopyWithImpl<$Res, _$RevenueInfoModelImpl>
    implements _$$RevenueInfoModelImplCopyWith<$Res> {
  __$$RevenueInfoModelImplCopyWithImpl(_$RevenueInfoModelImpl _value,
      $Res Function(_$RevenueInfoModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? totalRevenue = null,
    Object? revenueByMonth = null,
  }) {
    return _then(_$RevenueInfoModelImpl(
      totalRevenue: null == totalRevenue
          ? _value.totalRevenue
          : totalRevenue // ignore: cast_nullable_to_non_nullable
              as int,
      revenueByMonth: null == revenueByMonth
          ? _value._revenueByMonth
          : revenueByMonth // ignore: cast_nullable_to_non_nullable
              as List<MonthDataModel>,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$RevenueInfoModelImpl implements _RevenueInfoModel {
  const _$RevenueInfoModelImpl(
      {required this.totalRevenue,
      required final List<MonthDataModel> revenueByMonth})
      : _revenueByMonth = revenueByMonth;

  factory _$RevenueInfoModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$RevenueInfoModelImplFromJson(json);

  @override
  final int totalRevenue;
  final List<MonthDataModel> _revenueByMonth;
  @override
  List<MonthDataModel> get revenueByMonth {
    if (_revenueByMonth is EqualUnmodifiableListView) return _revenueByMonth;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_revenueByMonth);
  }

  @override
  String toString() {
    return 'RevenueInfoModel(totalRevenue: $totalRevenue, revenueByMonth: $revenueByMonth)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$RevenueInfoModelImpl &&
            (identical(other.totalRevenue, totalRevenue) ||
                other.totalRevenue == totalRevenue) &&
            const DeepCollectionEquality()
                .equals(other._revenueByMonth, _revenueByMonth));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, totalRevenue,
      const DeepCollectionEquality().hash(_revenueByMonth));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$RevenueInfoModelImplCopyWith<_$RevenueInfoModelImpl> get copyWith =>
      __$$RevenueInfoModelImplCopyWithImpl<_$RevenueInfoModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$RevenueInfoModelImplToJson(
      this,
    );
  }
}

abstract class _RevenueInfoModel implements RevenueInfoModel {
  const factory _RevenueInfoModel(
          {required final int totalRevenue,
          required final List<MonthDataModel> revenueByMonth}) =
      _$RevenueInfoModelImpl;

  factory _RevenueInfoModel.fromJson(Map<String, dynamic> json) =
      _$RevenueInfoModelImpl.fromJson;

  @override
  int get totalRevenue;
  @override
  List<MonthDataModel> get revenueByMonth;
  @override
  @JsonKey(ignore: true)
  _$$RevenueInfoModelImplCopyWith<_$RevenueInfoModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
