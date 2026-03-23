import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:dropdown_button2/dropdown_button2.dart';

class DropdownOption {
  final String label;
  final String value;
  bool selected;
  final bool disabled;

  DropdownOption({
    required this.label,
    required this.value,
    this.selected = false,
    this.disabled = false,
  });
}

enum DropdownBorderStyle {
  outline,
  underline,
}

class Dropdown extends StatefulWidget {
  final List<DropdownOption> entries;
  final ValueChanged<String?> onChanged;
  final bool isMultiSelect;
  final String label;
  final void Function(String value)? onClickClear;
  final bool hasClearIcon;

  /// style of dropdown items
  final Offset dropdownOffset;
  final double maxHeight;
  final double width;

  final bool isExpanded;

  final DropdownBorderStyle borderStyle;

  final String? errorText;

  const Dropdown({
    super.key,
    required this.entries,
    required this.onChanged,
    this.dropdownOffset = const Offset(0, -5),
    this.maxHeight = 250,
    this.width = 150,
    this.isExpanded = false,
    this.isMultiSelect = false,
    required this.label,
    this.onClickClear,
    this.borderStyle = DropdownBorderStyle.outline,
    this.hasClearIcon = true,
    this.errorText
  });

  @override
  State<Dropdown> createState() => _DropdownState();
}

class _DropdownState extends State<Dropdown> {
  String? selected;

  @override
  void initState() {
    super.initState();
    if (widget.entries.isNotEmpty) {
      _setSelectedFromEntries();
    }
  }

  @override
  void didUpdateWidget(covariant Dropdown oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (widget.entries != oldWidget.entries && widget.entries.isNotEmpty) {
      _setSelectedFromEntries();
    }
  }

  void _setSelectedFromEntries() {
    final initiallySelectedOption = widget.entries.firstWhere(
            (e) => e.selected,
        orElse: () => DropdownOption(label: '', value: ''));
    if (initiallySelectedOption.value.isNotEmpty && initiallySelectedOption.value != selected) {
      setState(() {
        selected = initiallySelectedOption.value;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    bool hasError = widget.errorText != null && widget.errorText!.isNotEmpty;

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Container(
          width: widget.isExpanded ? null : 84,
          height: 35,
          decoration: widget.borderStyle == DropdownBorderStyle.outline
              ? BoxDecoration(
            border: Border.all(
              color: hasError ? Colors.red : AppColors.waterloo.withAlpha(153),
              width: 1,
            ),
            borderRadius: BorderRadius.circular(4),
            color: Colors.white,
          )
              : BoxDecoration(
            border: Border(
              bottom: BorderSide(
                color: hasError ? Colors.red : AppColors.waterloo.withAlpha(153),
                width: 1,
              ),
            ),
            color: Colors.white,
          ),
          child: DropdownButtonHideUnderline(
            child: DropdownButton2<String>(
              dropdownStyleData: DropdownStyleData(
                offset: widget.dropdownOffset,
                maxHeight: widget.maxHeight,
                width: widget.isExpanded ? null : widget.width,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(8),
                  color: Colors.white,
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black.withAlpha(25),
                      blurRadius: 4,
                      offset: const Offset(0, 2),
                    ),
                  ],
                ),
                scrollbarTheme: ScrollbarThemeData(
                  radius: const Radius.circular(40),
                  thickness: WidgetStateProperty.all(6),
                  thumbVisibility: WidgetStateProperty.all(true),
                  thumbColor: WidgetStateProperty.all(AppColors.waterloo.withAlpha(127)),
                ),
              ),
              isDense: true,
              isExpanded: widget.isExpanded,
              customButton: Row(
                children: [
                  Expanded(
                    child: Padding(
                      padding: widget.borderStyle == DropdownBorderStyle.outline
                          ? const EdgeInsets.only(left: 10)
                          : const EdgeInsets.only(left: 0),
                      child: () {
                        final label = !widget.isMultiSelect
                            ? widget.entries
                            .where((item) => item.value == selected)
                            .cast<DropdownOption?>()
                            .firstWhere((e) => e != null, orElse: () => null)
                            ?.label ??
                            widget.label
                            : widget.label;

                        return Text(
                          label,
                          style: TextStyle(
                            fontSize: 15,
                            color: !widget.isMultiSelect
                                ? (selected != null ? AppColors.licorice : AppColors.waterloo)
                                : AppColors.waterloo,
                          ),
                        );
                      }(),
                    ),
                  ),
                  Visibility(
                    visible: selected != null && !widget.isMultiSelect && widget.hasClearIcon,
                    child: Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 5),
                      child: GestureDetector(
                        onTap: () {
                          widget.onClickClear?.call(selected!);
                          setState(() {
                            selected = null;
                          });
                        },
                        child: SvgPicture.asset(Assets.circleClose, colorFilter: const ColorFilter.mode(AppColors.waterloo, BlendMode.srcIn)),
                      ),
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.only(right: 5),
                    child: Icon(
                      Icons.keyboard_arrow_down,
                      size: 15,
                      color: AppColors.waterloo.withAlpha(153),
                    ),
                  ),
                ],
              ),
              items: widget.entries.map<DropdownMenuItem<String>>((DropdownOption option) {
                return DropdownMenuItem<String>(
                  value: option.value,
                  enabled: !widget.isMultiSelect && !option.disabled,
                  child: StatefulBuilder(
                    builder: (BuildContext context, StateSetter setState) {
                      if (widget.isMultiSelect) {
                        return InkWell(
                          onTap: option.disabled
                              ? null
                              : () {
                            setState(() {
                              selected = option.value;
                              option.selected = !option.selected;
                            });
                            widget.onChanged(option.value);
                          },
                          child: SizedBox(
                            height: 40,
                            child: Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Text(
                                  option.label,
                                  style: TextStyle(
                                    fontSize: 13,
                                    color: option.disabled ? Colors.grey : AppColors.waterloo,
                                  ),
                                ),
                                Visibility(
                                  visible: option.selected,
                                  child: SvgPicture.asset(
                                    Assets.check,
                                    colorFilter: const ColorFilter.mode(AppColors.youngBlue, BlendMode.srcIn),
                                  ),
                                )
                              ],
                            ),
                          ),
                        );
                      }
                      return SizedBox(
                        height: 40,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            Text(
                              option.label,
                              style: TextStyle(
                                fontSize: 13,
                                color: option.disabled ? Colors.grey : AppColors.waterloo,
                              ),
                            ),
                            Visibility(
                              visible: selected == option.value,
                              child: SvgPicture.asset(
                                Assets.check,
                                colorFilter: const ColorFilter.mode(AppColors.youngBlue, BlendMode.srcIn),
                              ),
                            ),
                          ],
                        ),
                      );
                    },
                  ),
                );
              }).toList(),
              onChanged: (String? value) {
                if (value != null) {
                  final option = widget.entries.firstWhere(
                        (element) => element.value == value,
                    orElse: () => DropdownOption(label: '', value: ''),
                  );
                  if (option.disabled) return;
                }
                setState(() {
                  selected = value;
                });
                widget.onChanged(value);
              },
              menuItemStyleData: const MenuItemStyleData(
                height: 40,
                padding: EdgeInsets.only(left: 15, right: 20),
              ),
            ),
          ),
        ),
        if (hasError)
          Padding(
            padding: const EdgeInsets.only(top: 2),
            child: Text(
              widget.errorText!,
              style: const TextStyle(
                color: Colors.red,
                fontSize: 12,
              ),
            ),
          ),
      ],
    );
  }
}
