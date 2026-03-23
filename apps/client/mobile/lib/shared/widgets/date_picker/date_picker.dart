import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:calendar_date_picker2/calendar_date_picker2.dart';
import 'package:tagpeak/utils/date_helpers.dart';

enum InputType {
  underline,
  outline
}

class DatePicker extends StatefulWidget {
  final TextEditingController controller;
  final String placeholder;
  final Widget? prefixIcon;
  final DateTime? selectedDate;
  final DateTime? initialDate;
  final DateTime? endDate;
  final void Function(DateTime date)? onSelectedDate;
  final void Function()? onClickClear;
  final InputType type;
  final String? label;
  final Widget? suffix;
  final bool readOnly;

  const DatePicker({
    super.key,
    required this.controller,
    required this.placeholder,
    this.prefixIcon,
    this.selectedDate,
    this.initialDate,
    this.endDate,
    this.onSelectedDate,
    this.onClickClear,
    this.type = InputType.outline,
    this.label,
    this.suffix,
    this.readOnly = false
  });

  @override
  State<DatePicker> createState() => _DatePickerState();
}

class _DatePickerState extends State<DatePicker> {
  @override
  void initState() {
    super.initState();
    _updateControllerText();
  }

  @override
  void didUpdateWidget(DatePicker oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (widget.selectedDate != oldWidget.selectedDate) {
      _updateControllerText();
    }
  }

  void _updateControllerText() {
    final displayDate = widget.selectedDate;
    widget.controller.text = displayDate != null ? dateTimeToString(displayDate) : '';
  }

  @override
  void dispose() {
    widget.controller.dispose();
    super.dispose();
  }

  Future<List<DateTime?>?> showCustomCalendarDialog(BuildContext context) async {
    return await showDialog<List<DateTime?>>(
      context: context,
      builder: (BuildContext context) {
        return Dialog(
          shape: RoundedRectangleBorder(
            side: const BorderSide(color: AppColors.grey),
            borderRadius: BorderRadius.circular(6),
          ),
          child: Container(
            padding: const EdgeInsets.all(6),
            constraints: const BoxConstraints(
              minWidth: 350,
              maxWidth: 500,
              minHeight: 280,
            ),
            child: CalendarDatePicker2(
              config: CalendarDatePicker2Config(
                customModePickerIcon: const SizedBox.shrink(),
                selectedDayHighlightColor: AppColors.youngBlue,
                weekdayLabels: [
                  translation(context).monday,
                  translation(context).tuesday,
                  translation(context).wednesday,
                  translation(context).thursday,
                  translation(context).friday,
                  translation(context).saturday,
                  translation(context).sunday,
                ],
                weekdayLabelTextStyle: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice),
                firstDayOfWeek: 0,
                controlsHeight: 40,
                controlsTextStyle: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice),
                dayTextStyle: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice),
                disabledDayTextStyle: Theme.of(context).textTheme.bodyMedium,
                selectableDayPredicate: (day) => true,
              ),
              onValueChanged: (dates) {
                if (dates.isNotEmpty) {
                  widget.onSelectedDate?.call(dates[0]);
                  Future.delayed(const Duration(milliseconds: 200), () {
                    if (context.mounted) {
                      Navigator.pop(context);
                    }
                  });
                }
              },
              value: [widget.selectedDate ?? widget.initialDate],
            ),
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        if (widget.type == InputType.underline)
        Text(widget.label ?? translation(context).date, style: Theme.of(context).textTheme.titleSmall),
        TextField(
          controller: widget.controller,
          readOnly: true,
          decoration: InputDecoration(
            enabledBorder: widget.type == InputType.outline
                ? OutlineInputBorder(
              borderSide: BorderSide(
                  color: AppColors.waterloo.withAlpha(153), width: 1),
              borderRadius: BorderRadius.circular(4),
            )
                : UnderlineInputBorder(
              borderSide: BorderSide(
                  color: widget.selectedDate == null
                      ? AppColors.wildSand
                      : AppColors.licorice),
            ),
            focusedBorder: widget.type == InputType.outline
                ? OutlineInputBorder(
              borderSide: BorderSide(
                  color: AppColors.waterloo.withAlpha(153), width: 1),
              borderRadius: BorderRadius.circular(4),
            )
                : UnderlineInputBorder(
              borderSide: BorderSide(
                  color: widget.selectedDate == null
                      ? AppColors.wildSand
                      : AppColors.licorice),
            ),
            contentPadding: widget.type == InputType.outline ?
            const EdgeInsets.symmetric(vertical: 5, horizontal: 10) : const EdgeInsets.only(bottom: 10.0),
            isCollapsed: true,
            hintText: widget.placeholder,
            floatingLabelBehavior: FloatingLabelBehavior.never,
            hintStyle: Theme.of(context).textTheme.bodyMedium,
            prefixIcon: widget.prefixIcon,
            prefixIconConstraints: const BoxConstraints(
              minWidth: 30,
              minHeight: 0,
            ),
            suffixIcon: widget.suffix ?? (widget.controller.text.isNotEmpty ? GestureDetector(
                onTap: () {
                  widget.controller.text = '';
                  widget.onClickClear?.call();
                },
                child: SvgPicture.asset(Assets.circleClose, colorFilter: const ColorFilter.mode(AppColors.waterloo, BlendMode.srcIn))
            ) : null),
            suffixIconConstraints: const BoxConstraints(
              minWidth: 30,
              minHeight: 0,
            ),
          ),
          onTap: () {
            if(widget.readOnly) {
              return;
            }
            showCustomCalendarDialog(context);
          },
        ),
      ],
    );
  }
}
