// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'balance_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

BalanceModel _$BalanceModelFromJson(Map<String, dynamic> json) {
  return _BalanceModel.fromJson(json);
}

/// @nodoc
mixin _$BalanceModel {
  double get amountRewards => throw _privateConstructorUsedError;
  double get amountReferrals => throw _privateConstructorUsedError;
  double get paidWithdrawals => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $BalanceModelCopyWith<BalanceModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $BalanceModelCopyWith<$Res> {
  factory $BalanceModelCopyWith(
          BalanceModel value, $Res Function(BalanceModel) then) =
      _$BalanceModelCopyWithImpl<$Res, BalanceModel>;
  @useResult
  $Res call(
      {double amountRewards, double amountReferrals, double paidWithdrawals});
}

/// @nodoc
class _$BalanceModelCopyWithImpl<$Res, $Val extends BalanceModel>
    implements $BalanceModelCopyWith<$Res> {
  _$BalanceModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? amountRewards = null,
    Object? amountReferrals = null,
    Object? paidWithdrawals = null,
  }) {
    return _then(_value.copyWith(
      amountRewards: null == amountRewards
          ? _value.amountRewards
          : amountRewards // ignore: cast_nullable_to_non_nullable
              as double,
      amountReferrals: null == amountReferrals
          ? _value.amountReferrals
          : amountReferrals // ignore: cast_nullable_to_non_nullable
              as double,
      paidWithdrawals: null == paidWithdrawals
          ? _value.paidWithdrawals
          : paidWithdrawals // ignore: cast_nullable_to_non_nullable
              as double,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$BalanceModelImplCopyWith<$Res>
    implements $BalanceModelCopyWith<$Res> {
  factory _$$BalanceModelImplCopyWith(
          _$BalanceModelImpl value, $Res Function(_$BalanceModelImpl) then) =
      __$$BalanceModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {double amountRewards, double amountReferrals, double paidWithdrawals});
}

/// @nodoc
class __$$BalanceModelImplCopyWithImpl<$Res>
    extends _$BalanceModelCopyWithImpl<$Res, _$BalanceModelImpl>
    implements _$$BalanceModelImplCopyWith<$Res> {
  __$$BalanceModelImplCopyWithImpl(
      _$BalanceModelImpl _value, $Res Function(_$BalanceModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? amountRewards = null,
    Object? amountReferrals = null,
    Object? paidWithdrawals = null,
  }) {
    return _then(_$BalanceModelImpl(
      amountRewards: null == amountRewards
          ? _value.amountRewards
          : amountRewards // ignore: cast_nullable_to_non_nullable
              as double,
      amountReferrals: null == amountReferrals
          ? _value.amountReferrals
          : amountReferrals // ignore: cast_nullable_to_non_nullable
              as double,
      paidWithdrawals: null == paidWithdrawals
          ? _value.paidWithdrawals
          : paidWithdrawals // ignore: cast_nullable_to_non_nullable
              as double,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$BalanceModelImpl implements _BalanceModel {
  const _$BalanceModelImpl(
      {required this.amountRewards,
      required this.amountReferrals,
      required this.paidWithdrawals});

  factory _$BalanceModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$BalanceModelImplFromJson(json);

  @override
  final double amountRewards;
  @override
  final double amountReferrals;
  @override
  final double paidWithdrawals;

  @override
  String toString() {
    return 'BalanceModel(amountRewards: $amountRewards, amountReferrals: $amountReferrals, paidWithdrawals: $paidWithdrawals)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$BalanceModelImpl &&
            (identical(other.amountRewards, amountRewards) ||
                other.amountRewards == amountRewards) &&
            (identical(other.amountReferrals, amountReferrals) ||
                other.amountReferrals == amountReferrals) &&
            (identical(other.paidWithdrawals, paidWithdrawals) ||
                other.paidWithdrawals == paidWithdrawals));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode =>
      Object.hash(runtimeType, amountRewards, amountReferrals, paidWithdrawals);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$BalanceModelImplCopyWith<_$BalanceModelImpl> get copyWith =>
      __$$BalanceModelImplCopyWithImpl<_$BalanceModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$BalanceModelImplToJson(
      this,
    );
  }
}

abstract class _BalanceModel implements BalanceModel {
  const factory _BalanceModel(
      {required final double amountRewards,
      required final double amountReferrals,
      required final double paidWithdrawals}) = _$BalanceModelImpl;

  factory _BalanceModel.fromJson(Map<String, dynamic> json) =
      _$BalanceModelImpl.fromJson;

  @override
  double get amountRewards;
  @override
  double get amountReferrals;
  @override
  double get paidWithdrawals;
  @override
  @JsonKey(ignore: true)
  _$$BalanceModelImplCopyWith<_$BalanceModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
