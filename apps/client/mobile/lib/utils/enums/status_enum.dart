enum StatusEnum {
  live('LIVE'),
  tracked('TRACKED'),
  paid('PAID'),
  finished('FINISHED'),
  stopped('STOPPED'),
  rejected('REJECTED'),
  validated('VALIDATED'),
  expired('EXPIRED'),
  pending('PENDING'),
  completed('COMPLETED'),
  confirmed('CONFIRMED'),
  cancelled('CANCELLED'),
  requested('REQUESTED'),

  active('ACTIVE'),

  inactive('INACTIVE');

  const StatusEnum(this.value);

  final String value;

  static StatusEnum fromString(String value) {
    for (StatusEnum status in StatusEnum.values) {
      if (status.value == value) {
        return status;
      }
    }
    return StatusEnum.live;
  }

  @override
  String toString() => value;
}
