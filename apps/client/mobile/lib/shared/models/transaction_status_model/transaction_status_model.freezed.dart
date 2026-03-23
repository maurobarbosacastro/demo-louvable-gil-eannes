// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'transaction_status_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

TransactionStatusModel _$TransactionStatusModelFromJson(
    Map<String, dynamic> json) {
  return _TransactionStatusModel.fromJson(json);
}

/// @nodoc
mixin _$TransactionStatusModel {
  String? get status => throw _privateConstructorUsedError;
  int? get count => throw _privateConstructorUsedError;
  int? get warning => throw _privateConstructorUsedError;
  int? get value => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $TransactionStatusModelCopyWith<TransactionStatusModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TransactionStatusModelCopyWith<$Res> {
  factory $TransactionStatusModelCopyWith(TransactionStatusModel value,
          $Res Function(TransactionStatusModel) then) =
      _$TransactionStatusModelCopyWithImpl<$Res, TransactionStatusModel>;
  @useResult
  $Res call({String? status, int? count, int? warning, int? value});
}

/// @nodoc
class _$TransactionStatusModelCopyWithImpl<$Res,
        $Val extends TransactionStatusModel>
    implements $TransactionStatusModelCopyWith<$Res> {
  _$TransactionStatusModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? status = freezed,
    Object? count = freezed,
    Object? warning = freezed,
    Object? value = freezed,
  }) {
    return _then(_value.copyWith(
      status: freezed == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String?,
      count: freezed == count
          ? _value.count
          : count // ignore: cast_nullable_to_non_nullable
              as int?,
      warning: freezed == warning
          ? _value.warning
          : warning // ignore: cast_nullable_to_non_nullable
              as int?,
      value: freezed == value
          ? _value.value
          : value // ignore: cast_nullable_to_non_nullable
              as int?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$TransactionStatusModelImplCopyWith<$Res>
    implements $TransactionStatusModelCopyWith<$Res> {
  factory _$$TransactionStatusModelImplCopyWith(
          _$TransactionStatusModelImpl value,
          $Res Function(_$TransactionStatusModelImpl) then) =
      __$$TransactionStatusModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String? status, int? count, int? warning, int? value});
}

/// @nodoc
class __$$TransactionStatusModelImplCopyWithImpl<$Res>
    extends _$TransactionStatusModelCopyWithImpl<$Res,
        _$TransactionStatusModelImpl>
    implements _$$TransactionStatusModelImplCopyWith<$Res> {
  __$$TransactionStatusModelImplCopyWithImpl(
      _$TransactionStatusModelImpl _value,
      $Res Function(_$TransactionStatusModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? status = freezed,
    Object? count = freezed,
    Object? warning = freezed,
    Object? value = freezed,
  }) {
    return _then(_$TransactionStatusModelImpl(
      status: freezed == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String?,
      count: freezed == count
          ? _value.count
          : count // ignore: cast_nullable_to_non_nullable
              as int?,
      warning: freezed == warning
          ? _value.warning
          : warning // ignore: cast_nullable_to_non_nullable
              as int?,
      value: freezed == value
          ? _value.value
          : value // ignore: cast_nullable_to_non_nullable
              as int?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$TransactionStatusModelImpl implements _TransactionStatusModel {
  const _$TransactionStatusModelImpl(
      {this.status, this.count, this.warning, this.value});

  factory _$TransactionStatusModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$TransactionStatusModelImplFromJson(json);

  @override
  final String? status;
  @override
  final int? count;
  @override
  final int? warning;
  @override
  final int? value;

  @override
  String toString() {
    return 'TransactionStatusModel(status: $status, count: $count, warning: $warning, value: $value)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$TransactionStatusModelImpl &&
            (identical(other.status, status) || other.status == status) &&
            (identical(other.count, count) || other.count == count) &&
            (identical(other.warning, warning) || other.warning == warning) &&
            (identical(other.value, value) || other.value == value));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, status, count, warning, value);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$TransactionStatusModelImplCopyWith<_$TransactionStatusModelImpl>
      get copyWith => __$$TransactionStatusModelImplCopyWithImpl<
          _$TransactionStatusModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$TransactionStatusModelImplToJson(
      this,
    );
  }
}

abstract class _TransactionStatusModel implements TransactionStatusModel {
  const factory _TransactionStatusModel(
      {final String? status,
      final int? count,
      final int? warning,
      final int? value}) = _$TransactionStatusModelImpl;

  factory _TransactionStatusModel.fromJson(Map<String, dynamic> json) =
      _$TransactionStatusModelImpl.fromJson;

  @override
  String? get status;
  @override
  int? get count;
  @override
  int? get warning;
  @override
  int? get value;
  @override
  @JsonKey(ignore: true)
  _$$TransactionStatusModelImplCopyWith<_$TransactionStatusModelImpl>
      get copyWith => throw _privateConstructorUsedError;
}
