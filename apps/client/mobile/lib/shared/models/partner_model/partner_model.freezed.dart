// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'partner_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

PartnerModel _$PartnerModelFromJson(Map<String, dynamic> json) {
  return _PartnerModel.fromJson(json);
}

/// @nodoc
mixin _$PartnerModel {
  String get uuid => throw _privateConstructorUsedError;
  String get name => throw _privateConstructorUsedError;
  String get code => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $PartnerModelCopyWith<PartnerModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $PartnerModelCopyWith<$Res> {
  factory $PartnerModelCopyWith(
          PartnerModel value, $Res Function(PartnerModel) then) =
      _$PartnerModelCopyWithImpl<$Res, PartnerModel>;
  @useResult
  $Res call({String uuid, String name, String code});
}

/// @nodoc
class _$PartnerModelCopyWithImpl<$Res, $Val extends PartnerModel>
    implements $PartnerModelCopyWith<$Res> {
  _$PartnerModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? name = null,
    Object? code = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      code: null == code
          ? _value.code
          : code // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$PartnerModelImplCopyWith<$Res>
    implements $PartnerModelCopyWith<$Res> {
  factory _$$PartnerModelImplCopyWith(
          _$PartnerModelImpl value, $Res Function(_$PartnerModelImpl) then) =
      __$$PartnerModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String uuid, String name, String code});
}

/// @nodoc
class __$$PartnerModelImplCopyWithImpl<$Res>
    extends _$PartnerModelCopyWithImpl<$Res, _$PartnerModelImpl>
    implements _$$PartnerModelImplCopyWith<$Res> {
  __$$PartnerModelImplCopyWithImpl(
      _$PartnerModelImpl _value, $Res Function(_$PartnerModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? name = null,
    Object? code = null,
  }) {
    return _then(_$PartnerModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      code: null == code
          ? _value.code
          : code // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$PartnerModelImpl implements _PartnerModel {
  const _$PartnerModelImpl(
      {required this.uuid, required this.name, required this.code});

  factory _$PartnerModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$PartnerModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String name;
  @override
  final String code;

  @override
  String toString() {
    return 'PartnerModel(uuid: $uuid, name: $name, code: $code)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$PartnerModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.name, name) || other.name == name) &&
            (identical(other.code, code) || other.code == code));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, uuid, name, code);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$PartnerModelImplCopyWith<_$PartnerModelImpl> get copyWith =>
      __$$PartnerModelImplCopyWithImpl<_$PartnerModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$PartnerModelImplToJson(
      this,
    );
  }
}

abstract class _PartnerModel implements PartnerModel {
  const factory _PartnerModel(
      {required final String uuid,
      required final String name,
      required final String code}) = _$PartnerModelImpl;

  factory _PartnerModel.fromJson(Map<String, dynamic> json) =
      _$PartnerModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get name;
  @override
  String get code;
  @override
  @JsonKey(ignore: true)
  _$$PartnerModelImplCopyWith<_$PartnerModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
