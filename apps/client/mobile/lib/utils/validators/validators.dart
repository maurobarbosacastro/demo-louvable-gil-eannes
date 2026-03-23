class Validators {
  static bool validateEmail(String? value) {
    String pattern =
        r'^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$';
    RegExp regex = RegExp(pattern);
    if (!regex.hasMatch(value!)) {
      return true;
    } else {
      return false;
    }
  }

  static bool validatePassword(String? value) {
    String pattern =
        '''^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[!"#\$'%&()*+,-./:;<=>?@[\\]^_`{|}~]).{8,20}\$''';
    RegExp regex = RegExp(pattern);
    if (!regex.hasMatch(value!)) {
      return true;
    } else {
      return false;
    }
  }

  static bool validateField(String? value) {
    return value!.isEmpty;
  }
}
