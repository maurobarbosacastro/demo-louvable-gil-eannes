import 'package:json_annotation/json_annotation.dart';

enum UserRolesEnum {
  @JsonValue('default')
  client,

  @JsonValue('admin')
  admin,
}
