// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'form_error_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

FormErrorModel _$FormErrorModelFromJson(Map<String, dynamic> json) {
  return _FormErrorModel.fromJson(json);
}

/// @nodoc
mixin _$FormErrorModel {
  bool get isError => throw _privateConstructorUsedError;
  String? get errorMessage => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $FormErrorModelCopyWith<FormErrorModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $FormErrorModelCopyWith<$Res> {
  factory $FormErrorModelCopyWith(
          FormErrorModel value, $Res Function(FormErrorModel) then) =
      _$FormErrorModelCopyWithImpl<$Res, FormErrorModel>;
  @useResult
  $Res call({bool isError, String? errorMessage});
}

/// @nodoc
class _$FormErrorModelCopyWithImpl<$Res, $Val extends FormErrorModel>
    implements $FormErrorModelCopyWith<$Res> {
  _$FormErrorModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isError = null,
    Object? errorMessage = freezed,
  }) {
    return _then(_value.copyWith(
      isError: null == isError
          ? _value.isError
          : isError // ignore: cast_nullable_to_non_nullable
              as bool,
      errorMessage: freezed == errorMessage
          ? _value.errorMessage
          : errorMessage // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$FormErrorModelImplCopyWith<$Res>
    implements $FormErrorModelCopyWith<$Res> {
  factory _$$FormErrorModelImplCopyWith(_$FormErrorModelImpl value,
          $Res Function(_$FormErrorModelImpl) then) =
      __$$FormErrorModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({bool isError, String? errorMessage});
}

/// @nodoc
class __$$FormErrorModelImplCopyWithImpl<$Res>
    extends _$FormErrorModelCopyWithImpl<$Res, _$FormErrorModelImpl>
    implements _$$FormErrorModelImplCopyWith<$Res> {
  __$$FormErrorModelImplCopyWithImpl(
      _$FormErrorModelImpl _value, $Res Function(_$FormErrorModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? isError = null,
    Object? errorMessage = freezed,
  }) {
    return _then(_$FormErrorModelImpl(
      isError: null == isError
          ? _value.isError
          : isError // ignore: cast_nullable_to_non_nullable
              as bool,
      errorMessage: freezed == errorMessage
          ? _value.errorMessage
          : errorMessage // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$FormErrorModelImpl implements _FormErrorModel {
  _$FormErrorModelImpl({this.isError = false, this.errorMessage});

  factory _$FormErrorModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$FormErrorModelImplFromJson(json);

  @override
  @JsonKey()
  final bool isError;
  @override
  final String? errorMessage;

  @override
  String toString() {
    return 'FormErrorModel(isError: $isError, errorMessage: $errorMessage)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$FormErrorModelImpl &&
            (identical(other.isError, isError) || other.isError == isError) &&
            (identical(other.errorMessage, errorMessage) ||
                other.errorMessage == errorMessage));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, isError, errorMessage);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$FormErrorModelImplCopyWith<_$FormErrorModelImpl> get copyWith =>
      __$$FormErrorModelImplCopyWithImpl<_$FormErrorModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$FormErrorModelImplToJson(
      this,
    );
  }
}

abstract class _FormErrorModel implements FormErrorModel {
  factory _FormErrorModel({final bool isError, final String? errorMessage}) =
      _$FormErrorModelImpl;

  factory _FormErrorModel.fromJson(Map<String, dynamic> json) =
      _$FormErrorModelImpl.fromJson;

  @override
  bool get isError;
  @override
  String? get errorMessage;
  @override
  @JsonKey(ignore: true)
  _$$FormErrorModelImplCopyWith<_$FormErrorModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
