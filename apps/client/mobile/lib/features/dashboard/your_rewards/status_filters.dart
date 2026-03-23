import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/enums/status_enum.dart';

class Status {
  final String label;
  final StatusEnum value;

  const Status({
    required this.label,
    required this.value,
  });
}

List<Status> transactionStatus(BuildContext context) => [
  Status(
    label: translation(context).tracked,
    value: StatusEnum.tracked,
  ),
  Status(
    label: translation(context).validated,
    value: StatusEnum.validated,
  ),
  Status(
    label: translation(context).rejected,
    value: StatusEnum.rejected,
  ),
  Status(
    label: translation(context).live,
    value: StatusEnum.live,
  ),
  Status(
    label: translation(context).stopped,
    value: StatusEnum.stopped,
  ),
  Status(
    label: translation(context).expired,
    value: StatusEnum.expired,
  ),
  Status(
    label: translation(context).finished,
    value: StatusEnum.finished,
  ),
  Status(
    label: translation(context).paid,
    value: StatusEnum.paid,
  ),
  Status(
    label: translation(context).completed,
    value: StatusEnum.completed,
  ),
];

List<Status> balanceStatus(BuildContext context) {
  final mappedStatuses = transactionStatus(context).map((status) {
    if (status.value == StatusEnum.completed) {
      return Status(
        label: translation(context).paid,
        value: status.value,
      );
    } else if (status.value == StatusEnum.tracked) {
      return Status(
        label: translation(context).pending,
        value: status.value,
      );
    }
    return status;
  }).toList();

  mappedStatuses.add(Status(
    label: translation(context).pending,
    value: StatusEnum.pending,
  ));

  return mappedStatuses;
}
