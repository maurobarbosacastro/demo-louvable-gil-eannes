// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'store_visit_cashback_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

StoreVisitCashbackModel _$StoreVisitCashbackModelFromJson(
    Map<String, dynamic> json) {
  return _StoreVisitCashbackModel.fromJson(json);
}

/// @nodoc
mixin _$StoreVisitCashbackModel {
  String get uuid => throw _privateConstructorUsedError;
  String get user => throw _privateConstructorUsedError;
  String? get reference => throw _privateConstructorUsedError;
  bool? get purchase => throw _privateConstructorUsedError;
  String? get storeUuid => throw _privateConstructorUsedError;
  DateTime? get createdAt => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $StoreVisitCashbackModelCopyWith<StoreVisitCashbackModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $StoreVisitCashbackModelCopyWith<$Res> {
  factory $StoreVisitCashbackModelCopyWith(StoreVisitCashbackModel value,
          $Res Function(StoreVisitCashbackModel) then) =
      _$StoreVisitCashbackModelCopyWithImpl<$Res, StoreVisitCashbackModel>;
  @useResult
  $Res call(
      {String uuid,
      String user,
      String? reference,
      bool? purchase,
      String? storeUuid,
      DateTime? createdAt});
}

/// @nodoc
class _$StoreVisitCashbackModelCopyWithImpl<$Res,
        $Val extends StoreVisitCashbackModel>
    implements $StoreVisitCashbackModelCopyWith<$Res> {
  _$StoreVisitCashbackModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = null,
    Object? reference = freezed,
    Object? purchase = freezed,
    Object? storeUuid = freezed,
    Object? createdAt = freezed,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: null == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as String,
      reference: freezed == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String?,
      purchase: freezed == purchase
          ? _value.purchase
          : purchase // ignore: cast_nullable_to_non_nullable
              as bool?,
      storeUuid: freezed == storeUuid
          ? _value.storeUuid
          : storeUuid // ignore: cast_nullable_to_non_nullable
              as String?,
      createdAt: freezed == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as DateTime?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$StoreVisitCashbackModelImplCopyWith<$Res>
    implements $StoreVisitCashbackModelCopyWith<$Res> {
  factory _$$StoreVisitCashbackModelImplCopyWith(
          _$StoreVisitCashbackModelImpl value,
          $Res Function(_$StoreVisitCashbackModelImpl) then) =
      __$$StoreVisitCashbackModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      String user,
      String? reference,
      bool? purchase,
      String? storeUuid,
      DateTime? createdAt});
}

/// @nodoc
class __$$StoreVisitCashbackModelImplCopyWithImpl<$Res>
    extends _$StoreVisitCashbackModelCopyWithImpl<$Res,
        _$StoreVisitCashbackModelImpl>
    implements _$$StoreVisitCashbackModelImplCopyWith<$Res> {
  __$$StoreVisitCashbackModelImplCopyWithImpl(
      _$StoreVisitCashbackModelImpl _value,
      $Res Function(_$StoreVisitCashbackModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = null,
    Object? reference = freezed,
    Object? purchase = freezed,
    Object? storeUuid = freezed,
    Object? createdAt = freezed,
  }) {
    return _then(_$StoreVisitCashbackModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: null == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as String,
      reference: freezed == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String?,
      purchase: freezed == purchase
          ? _value.purchase
          : purchase // ignore: cast_nullable_to_non_nullable
              as bool?,
      storeUuid: freezed == storeUuid
          ? _value.storeUuid
          : storeUuid // ignore: cast_nullable_to_non_nullable
              as String?,
      createdAt: freezed == createdAt
          ? _value.createdAt
          : createdAt // ignore: cast_nullable_to_non_nullable
              as DateTime?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$StoreVisitCashbackModelImpl implements _StoreVisitCashbackModel {
  const _$StoreVisitCashbackModelImpl(
      {required this.uuid,
      required this.user,
      this.reference,
      this.purchase,
      this.storeUuid,
      this.createdAt});

  factory _$StoreVisitCashbackModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$StoreVisitCashbackModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String user;
  @override
  final String? reference;
  @override
  final bool? purchase;
  @override
  final String? storeUuid;
  @override
  final DateTime? createdAt;

  @override
  String toString() {
    return 'StoreVisitCashbackModel(uuid: $uuid, user: $user, reference: $reference, purchase: $purchase, storeUuid: $storeUuid, createdAt: $createdAt)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$StoreVisitCashbackModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.user, user) || other.user == user) &&
            (identical(other.reference, reference) ||
                other.reference == reference) &&
            (identical(other.purchase, purchase) ||
                other.purchase == purchase) &&
            (identical(other.storeUuid, storeUuid) ||
                other.storeUuid == storeUuid) &&
            (identical(other.createdAt, createdAt) ||
                other.createdAt == createdAt));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType, uuid, user, reference, purchase, storeUuid, createdAt);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$StoreVisitCashbackModelImplCopyWith<_$StoreVisitCashbackModelImpl>
      get copyWith => __$$StoreVisitCashbackModelImplCopyWithImpl<
          _$StoreVisitCashbackModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$StoreVisitCashbackModelImplToJson(
      this,
    );
  }
}

abstract class _StoreVisitCashbackModel implements StoreVisitCashbackModel {
  const factory _StoreVisitCashbackModel(
      {required final String uuid,
      required final String user,
      final String? reference,
      final bool? purchase,
      final String? storeUuid,
      final DateTime? createdAt}) = _$StoreVisitCashbackModelImpl;

  factory _StoreVisitCashbackModel.fromJson(Map<String, dynamic> json) =
      _$StoreVisitCashbackModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String get user;
  @override
  String? get reference;
  @override
  bool? get purchase;
  @override
  String? get storeUuid;
  @override
  DateTime? get createdAt;
  @override
  @JsonKey(ignore: true)
  _$$StoreVisitCashbackModelImplCopyWith<_$StoreVisitCashbackModelImpl>
      get copyWith => throw _privateConstructorUsedError;
}
