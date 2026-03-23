import 'package:intl/intl.dart';

String formatDate(String dateString, {String pattern = 'MMM. d, yyyy'}) {
  DateTime dateTime = DateTime.parse(dateString).toLocal();

  DateFormat formatter = DateFormat(pattern);
  return formatter.format(dateTime);
}

String dateTimeToString(DateTime date, {String pattern = 'dd/MM/yyyy'}) {
  DateFormat formatter = DateFormat(pattern);
  return formatter.format(date);
}

int differenceBetweenDates(String startDateString, String endDateString) {
  DateTime startDate = DateTime.parse(startDateString);
  DateTime endDate = DateTime.parse(endDateString);

  return startDate.difference(endDate).inDays;
}
