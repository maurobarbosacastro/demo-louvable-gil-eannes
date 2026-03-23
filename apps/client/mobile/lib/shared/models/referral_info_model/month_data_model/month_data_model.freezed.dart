// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'month_data_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

MonthDataModel _$MonthDataModelFromJson(Map<String, dynamic> json) {
  return _MonthDataModel.fromJson(json);
}

/// @nodoc
mixin _$MonthDataModel {
  String get month => throw _privateConstructorUsedError;
  @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson)
  int? get value => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $MonthDataModelCopyWith<MonthDataModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $MonthDataModelCopyWith<$Res> {
  factory $MonthDataModelCopyWith(
          MonthDataModel value, $Res Function(MonthDataModel) then) =
      _$MonthDataModelCopyWithImpl<$Res, MonthDataModel>;
  @useResult
  $Res call(
      {String month,
      @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson) int? value});
}

/// @nodoc
class _$MonthDataModelCopyWithImpl<$Res, $Val extends MonthDataModel>
    implements $MonthDataModelCopyWith<$Res> {
  _$MonthDataModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? month = null,
    Object? value = freezed,
  }) {
    return _then(_value.copyWith(
      month: null == month
          ? _value.month
          : month // ignore: cast_nullable_to_non_nullable
              as String,
      value: freezed == value
          ? _value.value
          : value // ignore: cast_nullable_to_non_nullable
              as int?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$MonthDataModelImplCopyWith<$Res>
    implements $MonthDataModelCopyWith<$Res> {
  factory _$$MonthDataModelImplCopyWith(_$MonthDataModelImpl value,
          $Res Function(_$MonthDataModelImpl) then) =
      __$$MonthDataModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String month,
      @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson) int? value});
}

/// @nodoc
class __$$MonthDataModelImplCopyWithImpl<$Res>
    extends _$MonthDataModelCopyWithImpl<$Res, _$MonthDataModelImpl>
    implements _$$MonthDataModelImplCopyWith<$Res> {
  __$$MonthDataModelImplCopyWithImpl(
      _$MonthDataModelImpl _value, $Res Function(_$MonthDataModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? month = null,
    Object? value = freezed,
  }) {
    return _then(_$MonthDataModelImpl(
      month: null == month
          ? _value.month
          : month // ignore: cast_nullable_to_non_nullable
              as String,
      value: freezed == value
          ? _value.value
          : value // ignore: cast_nullable_to_non_nullable
              as int?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$MonthDataModelImpl implements _MonthDataModel {
  const _$MonthDataModelImpl(
      {required this.month,
      @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson) this.value});

  factory _$MonthDataModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$MonthDataModelImplFromJson(json);

  @override
  final String month;
  @override
  @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson)
  final int? value;

  @override
  String toString() {
    return 'MonthDataModel(month: $month, value: $value)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$MonthDataModelImpl &&
            (identical(other.month, month) || other.month == month) &&
            (identical(other.value, value) || other.value == value));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, month, value);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$MonthDataModelImplCopyWith<_$MonthDataModelImpl> get copyWith =>
      __$$MonthDataModelImplCopyWithImpl<_$MonthDataModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$MonthDataModelImplToJson(
      this,
    );
  }
}

abstract class _MonthDataModel implements MonthDataModel {
  const factory _MonthDataModel(
      {required final String month,
      @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson)
      final int? value}) = _$MonthDataModelImpl;

  factory _MonthDataModel.fromJson(Map<String, dynamic> json) =
      _$MonthDataModelImpl.fromJson;

  @override
  String get month;
  @override
  @JsonKey(fromJson: _numberFromJson, toJson: _numberToJson)
  int? get value;
  @override
  @JsonKey(ignore: true)
  _$$MonthDataModelImplCopyWith<_$MonthDataModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
