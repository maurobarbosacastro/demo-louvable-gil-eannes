import 'package:flutter/material.dart';
import '../../../../../utils/constants/colors.dart';

class CheckboxWidget extends StatelessWidget {
  final bool value;
  final String label;
  final Function(bool?)? onChanged;

  const CheckboxWidget({super.key, required this.value, required this.onChanged, required this.label});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onChanged != null
          ? () {
        onChanged!(!value);
      }
          : null,
      child: Row(
        children: [
          SizedBox(
            width: 20,
            height: 20,
            child: Transform.scale(
              scale: 0.8,
              child: Checkbox(
                  visualDensity: VisualDensity.compact,
                  materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
                  value: value,
                  side: WidgetStateBorderSide.resolveWith((states) => BorderSide(width: 1.0, color: value ? AppColors.youngBlue : AppColors.waterloo)),
                  fillColor: WidgetStateColor.resolveWith((states) => value ? AppColors.youngBlue : Colors.transparent),
                  checkColor: Colors.white,
                  shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(2)),
                  onChanged: onChanged),
            ),
          ),
          const SizedBox(width: 4, height: 4),
          Text(label, style: Theme.of(context).textTheme.titleSmall)
        ],
      ),
    );
  }
}
