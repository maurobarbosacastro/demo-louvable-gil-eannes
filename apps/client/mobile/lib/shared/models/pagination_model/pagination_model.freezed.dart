// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'pagination_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$PaginationModel<T> {
  int get limit => throw _privateConstructorUsedError;
  int get page => throw _privateConstructorUsedError;
  String get sort => throw _privateConstructorUsedError;
  int get totalRows => throw _privateConstructorUsedError;
  int get totalPages => throw _privateConstructorUsedError;
  List<T> get data => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $PaginationModelCopyWith<T, PaginationModel<T>> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $PaginationModelCopyWith<T, $Res> {
  factory $PaginationModelCopyWith(
          PaginationModel<T> value, $Res Function(PaginationModel<T>) then) =
      _$PaginationModelCopyWithImpl<T, $Res, PaginationModel<T>>;
  @useResult
  $Res call(
      {int limit,
      int page,
      String sort,
      int totalRows,
      int totalPages,
      List<T> data});
}

/// @nodoc
class _$PaginationModelCopyWithImpl<T, $Res, $Val extends PaginationModel<T>>
    implements $PaginationModelCopyWith<T, $Res> {
  _$PaginationModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? limit = null,
    Object? page = null,
    Object? sort = null,
    Object? totalRows = null,
    Object? totalPages = null,
    Object? data = null,
  }) {
    return _then(_value.copyWith(
      limit: null == limit
          ? _value.limit
          : limit // ignore: cast_nullable_to_non_nullable
              as int,
      page: null == page
          ? _value.page
          : page // ignore: cast_nullable_to_non_nullable
              as int,
      sort: null == sort
          ? _value.sort
          : sort // ignore: cast_nullable_to_non_nullable
              as String,
      totalRows: null == totalRows
          ? _value.totalRows
          : totalRows // ignore: cast_nullable_to_non_nullable
              as int,
      totalPages: null == totalPages
          ? _value.totalPages
          : totalPages // ignore: cast_nullable_to_non_nullable
              as int,
      data: null == data
          ? _value.data
          : data // ignore: cast_nullable_to_non_nullable
              as List<T>,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$PaginationModelImplCopyWith<T, $Res>
    implements $PaginationModelCopyWith<T, $Res> {
  factory _$$PaginationModelImplCopyWith(_$PaginationModelImpl<T> value,
          $Res Function(_$PaginationModelImpl<T>) then) =
      __$$PaginationModelImplCopyWithImpl<T, $Res>;
  @override
  @useResult
  $Res call(
      {int limit,
      int page,
      String sort,
      int totalRows,
      int totalPages,
      List<T> data});
}

/// @nodoc
class __$$PaginationModelImplCopyWithImpl<T, $Res>
    extends _$PaginationModelCopyWithImpl<T, $Res, _$PaginationModelImpl<T>>
    implements _$$PaginationModelImplCopyWith<T, $Res> {
  __$$PaginationModelImplCopyWithImpl(_$PaginationModelImpl<T> _value,
      $Res Function(_$PaginationModelImpl<T>) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? limit = null,
    Object? page = null,
    Object? sort = null,
    Object? totalRows = null,
    Object? totalPages = null,
    Object? data = null,
  }) {
    return _then(_$PaginationModelImpl<T>(
      limit: null == limit
          ? _value.limit
          : limit // ignore: cast_nullable_to_non_nullable
              as int,
      page: null == page
          ? _value.page
          : page // ignore: cast_nullable_to_non_nullable
              as int,
      sort: null == sort
          ? _value.sort
          : sort // ignore: cast_nullable_to_non_nullable
              as String,
      totalRows: null == totalRows
          ? _value.totalRows
          : totalRows // ignore: cast_nullable_to_non_nullable
              as int,
      totalPages: null == totalPages
          ? _value.totalPages
          : totalPages // ignore: cast_nullable_to_non_nullable
              as int,
      data: null == data
          ? _value._data
          : data // ignore: cast_nullable_to_non_nullable
              as List<T>,
    ));
  }
}

/// @nodoc

class _$PaginationModelImpl<T> implements _PaginationModel<T> {
  const _$PaginationModelImpl(
      {required this.limit,
      required this.page,
      required this.sort,
      required this.totalRows,
      required this.totalPages,
      required final List<T> data})
      : _data = data;

  @override
  final int limit;
  @override
  final int page;
  @override
  final String sort;
  @override
  final int totalRows;
  @override
  final int totalPages;
  final List<T> _data;
  @override
  List<T> get data {
    if (_data is EqualUnmodifiableListView) return _data;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_data);
  }

  @override
  String toString() {
    return 'PaginationModel<$T>(limit: $limit, page: $page, sort: $sort, totalRows: $totalRows, totalPages: $totalPages, data: $data)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$PaginationModelImpl<T> &&
            (identical(other.limit, limit) || other.limit == limit) &&
            (identical(other.page, page) || other.page == page) &&
            (identical(other.sort, sort) || other.sort == sort) &&
            (identical(other.totalRows, totalRows) ||
                other.totalRows == totalRows) &&
            (identical(other.totalPages, totalPages) ||
                other.totalPages == totalPages) &&
            const DeepCollectionEquality().equals(other._data, _data));
  }

  @override
  int get hashCode => Object.hash(runtimeType, limit, page, sort, totalRows,
      totalPages, const DeepCollectionEquality().hash(_data));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$PaginationModelImplCopyWith<T, _$PaginationModelImpl<T>> get copyWith =>
      __$$PaginationModelImplCopyWithImpl<T, _$PaginationModelImpl<T>>(
          this, _$identity);
}

abstract class _PaginationModel<T> implements PaginationModel<T> {
  const factory _PaginationModel(
      {required final int limit,
      required final int page,
      required final String sort,
      required final int totalRows,
      required final int totalPages,
      required final List<T> data}) = _$PaginationModelImpl<T>;

  @override
  int get limit;
  @override
  int get page;
  @override
  String get sort;
  @override
  int get totalRows;
  @override
  int get totalPages;
  @override
  List<T> get data;
  @override
  @JsonKey(ignore: true)
  _$$PaginationModelImplCopyWith<T, _$PaginationModelImpl<T>> get copyWith =>
      throw _privateConstructorUsedError;
}
