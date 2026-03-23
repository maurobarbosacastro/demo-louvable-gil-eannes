// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'user_stats_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

UserStatsModel _$UserStatsModelFromJson(Map<String, dynamic> json) {
  return _UserStatsModel.fromJson(json);
}

/// @nodoc
mixin _$UserStatsModel {
  String get level => throw _privateConstructorUsedError;
  double get valueSpent => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $UserStatsModelCopyWith<UserStatsModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $UserStatsModelCopyWith<$Res> {
  factory $UserStatsModelCopyWith(
          UserStatsModel value, $Res Function(UserStatsModel) then) =
      _$UserStatsModelCopyWithImpl<$Res, UserStatsModel>;
  @useResult
  $Res call({String level, double valueSpent});
}

/// @nodoc
class _$UserStatsModelCopyWithImpl<$Res, $Val extends UserStatsModel>
    implements $UserStatsModelCopyWith<$Res> {
  _$UserStatsModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? level = null,
    Object? valueSpent = null,
  }) {
    return _then(_value.copyWith(
      level: null == level
          ? _value.level
          : level // ignore: cast_nullable_to_non_nullable
              as String,
      valueSpent: null == valueSpent
          ? _value.valueSpent
          : valueSpent // ignore: cast_nullable_to_non_nullable
              as double,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$UserStatsModelImplCopyWith<$Res>
    implements $UserStatsModelCopyWith<$Res> {
  factory _$$UserStatsModelImplCopyWith(_$UserStatsModelImpl value,
          $Res Function(_$UserStatsModelImpl) then) =
      __$$UserStatsModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String level, double valueSpent});
}

/// @nodoc
class __$$UserStatsModelImplCopyWithImpl<$Res>
    extends _$UserStatsModelCopyWithImpl<$Res, _$UserStatsModelImpl>
    implements _$$UserStatsModelImplCopyWith<$Res> {
  __$$UserStatsModelImplCopyWithImpl(
      _$UserStatsModelImpl _value, $Res Function(_$UserStatsModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? level = null,
    Object? valueSpent = null,
  }) {
    return _then(_$UserStatsModelImpl(
      level: null == level
          ? _value.level
          : level // ignore: cast_nullable_to_non_nullable
              as String,
      valueSpent: null == valueSpent
          ? _value.valueSpent
          : valueSpent // ignore: cast_nullable_to_non_nullable
              as double,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$UserStatsModelImpl implements _UserStatsModel {
  const _$UserStatsModelImpl({required this.level, required this.valueSpent});

  factory _$UserStatsModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$UserStatsModelImplFromJson(json);

  @override
  final String level;
  @override
  final double valueSpent;

  @override
  String toString() {
    return 'UserStatsModel(level: $level, valueSpent: $valueSpent)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$UserStatsModelImpl &&
            (identical(other.level, level) || other.level == level) &&
            (identical(other.valueSpent, valueSpent) ||
                other.valueSpent == valueSpent));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, level, valueSpent);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$UserStatsModelImplCopyWith<_$UserStatsModelImpl> get copyWith =>
      __$$UserStatsModelImplCopyWithImpl<_$UserStatsModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$UserStatsModelImplToJson(
      this,
    );
  }
}

abstract class _UserStatsModel implements UserStatsModel {
  const factory _UserStatsModel(
      {required final String level,
      required final double valueSpent}) = _$UserStatsModelImpl;

  factory _UserStatsModel.fromJson(Map<String, dynamic> json) =
      _$UserStatsModelImpl.fromJson;

  @override
  String get level;
  @override
  double get valueSpent;
  @override
  @JsonKey(ignore: true)
  _$$UserStatsModelImplCopyWith<_$UserStatsModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
