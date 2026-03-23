// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'email_verified_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

EmailVerifiedModel _$EmailVerifiedModelFromJson(Map<String, dynamic> json) {
  return _EmailVerifiedModel.fromJson(json);
}

/// @nodoc
mixin _$EmailVerifiedModel {
  String get userId => throw _privateConstructorUsedError;
  String get isVerified => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $EmailVerifiedModelCopyWith<EmailVerifiedModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $EmailVerifiedModelCopyWith<$Res> {
  factory $EmailVerifiedModelCopyWith(
          EmailVerifiedModel value, $Res Function(EmailVerifiedModel) then) =
      _$EmailVerifiedModelCopyWithImpl<$Res, EmailVerifiedModel>;
  @useResult
  $Res call({String userId, String isVerified});
}

/// @nodoc
class _$EmailVerifiedModelCopyWithImpl<$Res, $Val extends EmailVerifiedModel>
    implements $EmailVerifiedModelCopyWith<$Res> {
  _$EmailVerifiedModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? userId = null,
    Object? isVerified = null,
  }) {
    return _then(_value.copyWith(
      userId: null == userId
          ? _value.userId
          : userId // ignore: cast_nullable_to_non_nullable
              as String,
      isVerified: null == isVerified
          ? _value.isVerified
          : isVerified // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$EmailVerifiedModelImplCopyWith<$Res>
    implements $EmailVerifiedModelCopyWith<$Res> {
  factory _$$EmailVerifiedModelImplCopyWith(_$EmailVerifiedModelImpl value,
          $Res Function(_$EmailVerifiedModelImpl) then) =
      __$$EmailVerifiedModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String userId, String isVerified});
}

/// @nodoc
class __$$EmailVerifiedModelImplCopyWithImpl<$Res>
    extends _$EmailVerifiedModelCopyWithImpl<$Res, _$EmailVerifiedModelImpl>
    implements _$$EmailVerifiedModelImplCopyWith<$Res> {
  __$$EmailVerifiedModelImplCopyWithImpl(_$EmailVerifiedModelImpl _value,
      $Res Function(_$EmailVerifiedModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? userId = null,
    Object? isVerified = null,
  }) {
    return _then(_$EmailVerifiedModelImpl(
      userId: null == userId
          ? _value.userId
          : userId // ignore: cast_nullable_to_non_nullable
              as String,
      isVerified: null == isVerified
          ? _value.isVerified
          : isVerified // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$EmailVerifiedModelImpl implements _EmailVerifiedModel {
  const _$EmailVerifiedModelImpl(
      {required this.userId, required this.isVerified});

  factory _$EmailVerifiedModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$EmailVerifiedModelImplFromJson(json);

  @override
  final String userId;
  @override
  final String isVerified;

  @override
  String toString() {
    return 'EmailVerifiedModel(userId: $userId, isVerified: $isVerified)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$EmailVerifiedModelImpl &&
            (identical(other.userId, userId) || other.userId == userId) &&
            (identical(other.isVerified, isVerified) ||
                other.isVerified == isVerified));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, userId, isVerified);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$EmailVerifiedModelImplCopyWith<_$EmailVerifiedModelImpl> get copyWith =>
      __$$EmailVerifiedModelImplCopyWithImpl<_$EmailVerifiedModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$EmailVerifiedModelImplToJson(
      this,
    );
  }
}

abstract class _EmailVerifiedModel implements EmailVerifiedModel {
  const factory _EmailVerifiedModel(
      {required final String userId,
      required final String isVerified}) = _$EmailVerifiedModelImpl;

  factory _EmailVerifiedModel.fromJson(Map<String, dynamic> json) =
      _$EmailVerifiedModelImpl.fromJson;

  @override
  String get userId;
  @override
  String get isVerified;
  @override
  @JsonKey(ignore: true)
  _$$EmailVerifiedModelImplCopyWith<_$EmailVerifiedModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
