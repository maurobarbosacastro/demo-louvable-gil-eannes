import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/dashboard/your_rewards/status_filters.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/enums/status_enum.dart';

class StatusContainer extends StatelessWidget {
  final StatusEnum status;
  final bool hasCloseIcon;
  final void Function(StatusEnum)? onCloseClick;
  final List<Status>? statusList;

  const StatusContainer({super.key, required this.status, this.hasCloseIcon = false, this.onCloseClick, this.statusList});

  bool get isPending => [StatusEnum.tracked, StatusEnum.pending, StatusEnum.requested, StatusEnum.inactive].contains(status);
  bool get isFinishedPaid => [StatusEnum.paid, StatusEnum.finished, StatusEnum.validated, StatusEnum.completed, StatusEnum.confirmed, StatusEnum.active].contains(status);
  bool get isStopped => [StatusEnum.stopped, StatusEnum.rejected, StatusEnum.cancelled].contains(status);
  bool get isLive => status == StatusEnum.live;

  BoxDecoration _containerDecoration() {
    if (isPending) {
      return BoxDecoration(
        color: const Color.fromRGBO(122, 123, 146, 0.10),
        borderRadius: BorderRadius.circular(4),
      );
    }
    if (isFinishedPaid) {
      return BoxDecoration(
          color: const Color.fromRGBO(66, 55, 218, 0.20),
          borderRadius: BorderRadius.circular(4)
      );
    }
    if (isStopped) {
      return BoxDecoration(
          color: const Color.fromRGBO(242, 41, 91, 0.20),
          borderRadius: BorderRadius.circular(4)
      );
    }
    return BoxDecoration(
        color: const Color.fromRGBO(253, 202, 101, 0.40),
        borderRadius: BorderRadius.circular(4)
    );
  }

  TextStyle _textStyle() {
    if (isPending) {
      return const TextStyle(
          color: AppColors.waterloo,
          fontSize: 13,
          fontWeight: FontWeight.w500
      );
    }
    if (isFinishedPaid) {
      return const TextStyle(
          color: AppColors.youngBlue,
          fontSize: 13,
          fontWeight: FontWeight.w500
      );
    }
    if (isStopped) {
      return const TextStyle(
          color: AppColors.radicalRed,
          fontSize: 13,
          fontWeight: FontWeight.w500
      );
    }
    return const TextStyle(
        color: Color(0xFFCF8E0C),
        fontSize: 13,
        fontWeight: FontWeight.w500
    );
  }

  String _textTranslation(BuildContext context) {
    return (statusList ?? transactionStatus(context)).firstWhere(
            (item) => item.value.toString() == status.toString(),
        orElse: () => throw ArgumentError('No matching status filter found for value: ${status.toString()}')
    ).label;
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
      decoration: _containerDecoration(),
      child: Row(
        spacing: 5,
        mainAxisSize: MainAxisSize.min,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(_textTranslation(context), style: _textStyle()),
          Visibility(visible: hasCloseIcon,
              child: GestureDetector(child: SvgPicture.asset(Assets.circleClose, colorFilter: ColorFilter.mode(_textStyle().color!, BlendMode.srcIn)), onTap: () => onCloseClick?.call(status))
          )
        ],
      ),
    );
  }
}
