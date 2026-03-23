// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'store_visit_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

StoreVisitModel _$StoreVisitModelFromJson(Map<String, dynamic> json) {
  return _StoreVisitModel.fromJson(json);
}

/// @nodoc
mixin _$StoreVisitModel {
  String get uuid => throw _privateConstructorUsedError;
  String? get user => throw _privateConstructorUsedError;
  String get reference => throw _privateConstructorUsedError;
  StoreModel get store => throw _privateConstructorUsedError;
  bool get purchased => throw _privateConstructorUsedError;
  String get dateTime => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $StoreVisitModelCopyWith<StoreVisitModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $StoreVisitModelCopyWith<$Res> {
  factory $StoreVisitModelCopyWith(
          StoreVisitModel value, $Res Function(StoreVisitModel) then) =
      _$StoreVisitModelCopyWithImpl<$Res, StoreVisitModel>;
  @useResult
  $Res call(
      {String uuid,
      String? user,
      String reference,
      StoreModel store,
      bool purchased,
      String dateTime});

  $StoreModelCopyWith<$Res> get store;
}

/// @nodoc
class _$StoreVisitModelCopyWithImpl<$Res, $Val extends StoreVisitModel>
    implements $StoreVisitModelCopyWith<$Res> {
  _$StoreVisitModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = freezed,
    Object? reference = null,
    Object? store = null,
    Object? purchased = null,
    Object? dateTime = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: freezed == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as String?,
      reference: null == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String,
      store: null == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as StoreModel,
      purchased: null == purchased
          ? _value.purchased
          : purchased // ignore: cast_nullable_to_non_nullable
              as bool,
      dateTime: null == dateTime
          ? _value.dateTime
          : dateTime // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $StoreModelCopyWith<$Res> get store {
    return $StoreModelCopyWith<$Res>(_value.store, (value) {
      return _then(_value.copyWith(store: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$StoreVisitModelImplCopyWith<$Res>
    implements $StoreVisitModelCopyWith<$Res> {
  factory _$$StoreVisitModelImplCopyWith(_$StoreVisitModelImpl value,
          $Res Function(_$StoreVisitModelImpl) then) =
      __$$StoreVisitModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      String? user,
      String reference,
      StoreModel store,
      bool purchased,
      String dateTime});

  @override
  $StoreModelCopyWith<$Res> get store;
}

/// @nodoc
class __$$StoreVisitModelImplCopyWithImpl<$Res>
    extends _$StoreVisitModelCopyWithImpl<$Res, _$StoreVisitModelImpl>
    implements _$$StoreVisitModelImplCopyWith<$Res> {
  __$$StoreVisitModelImplCopyWithImpl(
      _$StoreVisitModelImpl _value, $Res Function(_$StoreVisitModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = freezed,
    Object? reference = null,
    Object? store = null,
    Object? purchased = null,
    Object? dateTime = null,
  }) {
    return _then(_$StoreVisitModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: freezed == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as String?,
      reference: null == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String,
      store: null == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as StoreModel,
      purchased: null == purchased
          ? _value.purchased
          : purchased // ignore: cast_nullable_to_non_nullable
              as bool,
      dateTime: null == dateTime
          ? _value.dateTime
          : dateTime // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$StoreVisitModelImpl implements _StoreVisitModel {
  const _$StoreVisitModelImpl(
      {required this.uuid,
      this.user,
      required this.reference,
      required this.store,
      required this.purchased,
      required this.dateTime});

  factory _$StoreVisitModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$StoreVisitModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final String? user;
  @override
  final String reference;
  @override
  final StoreModel store;
  @override
  final bool purchased;
  @override
  final String dateTime;

  @override
  String toString() {
    return 'StoreVisitModel(uuid: $uuid, user: $user, reference: $reference, store: $store, purchased: $purchased, dateTime: $dateTime)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$StoreVisitModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.user, user) || other.user == user) &&
            (identical(other.reference, reference) ||
                other.reference == reference) &&
            (identical(other.store, store) || other.store == store) &&
            (identical(other.purchased, purchased) ||
                other.purchased == purchased) &&
            (identical(other.dateTime, dateTime) ||
                other.dateTime == dateTime));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType, uuid, user, reference, store, purchased, dateTime);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$StoreVisitModelImplCopyWith<_$StoreVisitModelImpl> get copyWith =>
      __$$StoreVisitModelImplCopyWithImpl<_$StoreVisitModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$StoreVisitModelImplToJson(
      this,
    );
  }
}

abstract class _StoreVisitModel implements StoreVisitModel {
  const factory _StoreVisitModel(
      {required final String uuid,
      final String? user,
      required final String reference,
      required final StoreModel store,
      required final bool purchased,
      required final String dateTime}) = _$StoreVisitModelImpl;

  factory _StoreVisitModel.fromJson(Map<String, dynamic> json) =
      _$StoreVisitModelImpl.fromJson;

  @override
  String get uuid;
  @override
  String? get user;
  @override
  String get reference;
  @override
  StoreModel get store;
  @override
  bool get purchased;
  @override
  String get dateTime;
  @override
  @JsonKey(ignore: true)
  _$$StoreVisitModelImplCopyWith<_$StoreVisitModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

StoreVisitAdminModel _$StoreVisitAdminModelFromJson(Map<String, dynamic> json) {
  return _StoreVisitAdminModel.fromJson(json);
}

/// @nodoc
mixin _$StoreVisitAdminModel {
  String get uuid => throw _privateConstructorUsedError;
  UserStoreVisitModel? get user => throw _privateConstructorUsedError;
  String get reference => throw _privateConstructorUsedError;
  StoreModel get store => throw _privateConstructorUsedError;
  bool get purchased => throw _privateConstructorUsedError;
  String get dateTime => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $StoreVisitAdminModelCopyWith<StoreVisitAdminModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $StoreVisitAdminModelCopyWith<$Res> {
  factory $StoreVisitAdminModelCopyWith(StoreVisitAdminModel value,
          $Res Function(StoreVisitAdminModel) then) =
      _$StoreVisitAdminModelCopyWithImpl<$Res, StoreVisitAdminModel>;
  @useResult
  $Res call(
      {String uuid,
      UserStoreVisitModel? user,
      String reference,
      StoreModel store,
      bool purchased,
      String dateTime});

  $UserStoreVisitModelCopyWith<$Res>? get user;
  $StoreModelCopyWith<$Res> get store;
}

/// @nodoc
class _$StoreVisitAdminModelCopyWithImpl<$Res,
        $Val extends StoreVisitAdminModel>
    implements $StoreVisitAdminModelCopyWith<$Res> {
  _$StoreVisitAdminModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = freezed,
    Object? reference = null,
    Object? store = null,
    Object? purchased = null,
    Object? dateTime = null,
  }) {
    return _then(_value.copyWith(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: freezed == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as UserStoreVisitModel?,
      reference: null == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String,
      store: null == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as StoreModel,
      purchased: null == purchased
          ? _value.purchased
          : purchased // ignore: cast_nullable_to_non_nullable
              as bool,
      dateTime: null == dateTime
          ? _value.dateTime
          : dateTime // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $UserStoreVisitModelCopyWith<$Res>? get user {
    if (_value.user == null) {
      return null;
    }

    return $UserStoreVisitModelCopyWith<$Res>(_value.user!, (value) {
      return _then(_value.copyWith(user: value) as $Val);
    });
  }

  @override
  @pragma('vm:prefer-inline')
  $StoreModelCopyWith<$Res> get store {
    return $StoreModelCopyWith<$Res>(_value.store, (value) {
      return _then(_value.copyWith(store: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$StoreVisitAdminModelImplCopyWith<$Res>
    implements $StoreVisitAdminModelCopyWith<$Res> {
  factory _$$StoreVisitAdminModelImplCopyWith(_$StoreVisitAdminModelImpl value,
          $Res Function(_$StoreVisitAdminModelImpl) then) =
      __$$StoreVisitAdminModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String uuid,
      UserStoreVisitModel? user,
      String reference,
      StoreModel store,
      bool purchased,
      String dateTime});

  @override
  $UserStoreVisitModelCopyWith<$Res>? get user;
  @override
  $StoreModelCopyWith<$Res> get store;
}

/// @nodoc
class __$$StoreVisitAdminModelImplCopyWithImpl<$Res>
    extends _$StoreVisitAdminModelCopyWithImpl<$Res, _$StoreVisitAdminModelImpl>
    implements _$$StoreVisitAdminModelImplCopyWith<$Res> {
  __$$StoreVisitAdminModelImplCopyWithImpl(_$StoreVisitAdminModelImpl _value,
      $Res Function(_$StoreVisitAdminModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? uuid = null,
    Object? user = freezed,
    Object? reference = null,
    Object? store = null,
    Object? purchased = null,
    Object? dateTime = null,
  }) {
    return _then(_$StoreVisitAdminModelImpl(
      uuid: null == uuid
          ? _value.uuid
          : uuid // ignore: cast_nullable_to_non_nullable
              as String,
      user: freezed == user
          ? _value.user
          : user // ignore: cast_nullable_to_non_nullable
              as UserStoreVisitModel?,
      reference: null == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String,
      store: null == store
          ? _value.store
          : store // ignore: cast_nullable_to_non_nullable
              as StoreModel,
      purchased: null == purchased
          ? _value.purchased
          : purchased // ignore: cast_nullable_to_non_nullable
              as bool,
      dateTime: null == dateTime
          ? _value.dateTime
          : dateTime // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$StoreVisitAdminModelImpl implements _StoreVisitAdminModel {
  const _$StoreVisitAdminModelImpl(
      {required this.uuid,
      this.user,
      required this.reference,
      required this.store,
      required this.purchased,
      required this.dateTime});

  factory _$StoreVisitAdminModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$StoreVisitAdminModelImplFromJson(json);

  @override
  final String uuid;
  @override
  final UserStoreVisitModel? user;
  @override
  final String reference;
  @override
  final StoreModel store;
  @override
  final bool purchased;
  @override
  final String dateTime;

  @override
  String toString() {
    return 'StoreVisitAdminModel(uuid: $uuid, user: $user, reference: $reference, store: $store, purchased: $purchased, dateTime: $dateTime)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$StoreVisitAdminModelImpl &&
            (identical(other.uuid, uuid) || other.uuid == uuid) &&
            (identical(other.user, user) || other.user == user) &&
            (identical(other.reference, reference) ||
                other.reference == reference) &&
            (identical(other.store, store) || other.store == store) &&
            (identical(other.purchased, purchased) ||
                other.purchased == purchased) &&
            (identical(other.dateTime, dateTime) ||
                other.dateTime == dateTime));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType, uuid, user, reference, store, purchased, dateTime);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$StoreVisitAdminModelImplCopyWith<_$StoreVisitAdminModelImpl>
      get copyWith =>
          __$$StoreVisitAdminModelImplCopyWithImpl<_$StoreVisitAdminModelImpl>(
              this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$StoreVisitAdminModelImplToJson(
      this,
    );
  }
}

abstract class _StoreVisitAdminModel implements StoreVisitAdminModel {
  const factory _StoreVisitAdminModel(
      {required final String uuid,
      final UserStoreVisitModel? user,
      required final String reference,
      required final StoreModel store,
      required final bool purchased,
      required final String dateTime}) = _$StoreVisitAdminModelImpl;

  factory _StoreVisitAdminModel.fromJson(Map<String, dynamic> json) =
      _$StoreVisitAdminModelImpl.fromJson;

  @override
  String get uuid;
  @override
  UserStoreVisitModel? get user;
  @override
  String get reference;
  @override
  StoreModel get store;
  @override
  bool get purchased;
  @override
  String get dateTime;
  @override
  @JsonKey(ignore: true)
  _$$StoreVisitAdminModelImplCopyWith<_$StoreVisitAdminModelImpl>
      get copyWith => throw _privateConstructorUsedError;
}
