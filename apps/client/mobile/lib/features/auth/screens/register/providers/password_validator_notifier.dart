import 'package:flutter_riverpod/flutter_riverpod.dart';

class PasswordValidatorNotifier extends StateNotifier<bool> {
  PasswordValidatorNotifier(super.state);

  void validate(String text) {
    String size = r'^.{8,20}$';
    RegExp sizeRegex = RegExp(size);

    String special = '''^(.*[!"#\$'%&()*+,-./:;<=>?@[\\]^_`{|}~]+.*)\$''';
    RegExp specialRegex = RegExp(special);

    String number = r'^(.*\d+.*)$';
    RegExp numberRegex = RegExp(number);

    String upper = r'^(.*[A-Z].*)$';
    RegExp upperCaseRegex = RegExp(upper);

    state = sizeRegex.hasMatch(text) && specialRegex.hasMatch(text) && numberRegex.hasMatch(text) && upperCaseRegex.hasMatch(text);
  }

  void reset() => false;

  bool isValid() => state;
}
