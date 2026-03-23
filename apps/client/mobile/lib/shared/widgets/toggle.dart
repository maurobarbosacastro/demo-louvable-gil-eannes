import 'package:flutter/material.dart';
import 'package:flutter_switch/flutter_switch.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class Toggle extends StatelessWidget {
  final String label;
  final bool value;
  final Function(bool) onChanged;

  const Toggle({super.key, required this.label, required this.value, required this.onChanged});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(bottom: 15),
      decoration: const BoxDecoration(
        border: Border(
          bottom: BorderSide(
            width: 1,
            color: AppColors.wildSand,
          )
        ),
      ),
      child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(label, style: Theme.of(context).textTheme.bodyMedium),
            FlutterSwitch(
              value: value,
              onToggle: onChanged,
              inactiveColor: Colors.transparent,
              inactiveToggleColor: const Color(0xFF64758B),
              inactiveSwitchBorder: Border.all(
                width: 1.5,
                color: const Color(0xFF64758B),
              ),
              activeColor: AppColors.youngBlue,
              activeToggleColor: Colors.white,
              activeSwitchBorder: Border.all(
                width: 1.5,
                color: AppColors.youngBlue,
              ),
              toggleSize: 20,
              padding: 5,
              width: 60,
              height: 30,
            )
          ],
      ),
    );
  }
}
