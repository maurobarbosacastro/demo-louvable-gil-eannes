// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'user_store_visit_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

UserStoreVisitModel _$UserStoreVisitModelFromJson(Map<String, dynamic> json) {
  return _UserStoreVisitModel.fromJson(json);
}

/// @nodoc
mixin _$UserStoreVisitModel {
  String get uuid => throw _privateConstructorUsedError;
  String get firstName => throw _privateConstructorUsedError;
  String get lastName => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $UserStoreVisitModelCopyWith<UserStoreVisitModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $UserStoreVisitModelCopyWith<$Res> {
  factory $UserStoreVisitModelCopyWith(
          UserStoreVisitModel value, $Res Function(UserStoreVisitModel) then) =
      _$UserStoreVisitModelCopyWithImpl<$Res, UserStoreVisitModel>;
  @useResult
  $Res call({String uuid, String firstName, String lastName});
}

/// @nodoc
class _$UserStoreVisitModelCopyWithImpl<$Res, $Val extends UserStoreVisitModel>
    implements $UserStoreVisitModelCopyWith<$Res> {
  _$UserStoreVisitModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? firstName = null,
    Object? lastName = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      firstName: null == firstName
          ? _value.firstName
          : firstName // ignore: cast_nullable_to_non_nullable
              as String,
      lastName: null == lastName
          ? _value.lastName
          : lastName // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$UserStoreVisitModelImplCopyWith<$Res>
    implements $UserStoreVisitModelCopyWith<$Res> {
  factory _$$UserStoreVisitModelImplCopyWith(_$UserStoreVisitModelImpl value,
          $Res Function(_$UserStoreVisitModelImpl) then) =
      __$$UserStoreVisitModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String uuid, String firstName, String lastName});
}

/// @nodoc
class __$$UserStoreVisitModelImplCopyWithImpl<$Res>
    extends _$UserStoreVisitModelCopyWithImpl<$Res, _$UserStoreVisitModelImpl>
    implements _$$UserStoreVisitModelImplCopyWith<$Res> {
  __$$UserStoreVisitModelImplCopyWithImpl(_$UserStoreVisitModelImpl _value,
      $Res Function(_$UserStoreVisitModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? firstName = null,
    Object? lastName = null,
  }) {
    return _then(_$UserStoreVisitModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      firstName: null == firstName
          ? _value.firstName
          : firstName // ignore: cast_nullable_to_non_nullable
              as String,
      lastName: null == lastName
          ? _value.lastName
          : lastName // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$UserStoreVisitModelImpl implements _UserStoreVisitModel {
  const _$UserStoreVisitModelImpl(
      {required this.uuid, required this.firstName, required this.lastName});

  factory _$UserStoreVisitModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$UserStoreVisitModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String firstName;
  @override
  final String lastName;

  @override
  String toString() {
    return 'UserStoreVisitModel(uuid: $uuid, firstName: $firstName, lastName: $lastName)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$UserStoreVisitModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.firstName, firstName) ||
                other.firstName == firstName) &&
            (identical(other.lastName, lastName) ||
                other.lastName == lastName));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, uuid, firstName, lastName);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$UserStoreVisitModelImplCopyWith<_$UserStoreVisitModelImpl> get copyWith =>
      __$$UserStoreVisitModelImplCopyWithImpl<_$UserStoreVisitModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$UserStoreVisitModelImplToJson(
      this,
    );
  }
}

abstract class _UserStoreVisitModel implements UserStoreVisitModel {
  const factory _UserStoreVisitModel(
      {required final String uuid,
      required final String firstName,
      required final String lastName}) = _$UserStoreVisitModelImpl;

  factory _UserStoreVisitModel.fromJson(Map<String, dynamic> json) =
      _$UserStoreVisitModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get firstName;
  @override
  String get lastName;
  @override
  @JsonKey(ignore: true)
  _$$UserStoreVisitModelImplCopyWith<_$UserStoreVisitModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
