import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class TogglePanel extends StatefulWidget {
  final String title;
  final Widget content;
  const TogglePanel({super.key, required this.title, required this.content});

  @override
  State<TogglePanel> createState() => _TogglePanelState();
}

class _TogglePanelState extends State<TogglePanel> {
  bool isOpen = true;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      spacing: 20,
      children: [
        GestureDetector(
          onTap: () => setState(() => isOpen = !isOpen),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Text(widget.title, style: Theme.of(context).textTheme.headlineSmall),
              SvgPicture.asset(isOpen ? Assets.chevronUp : Assets.chevronDown)
            ],
          ),
        ),
        Visibility(
          visible: isOpen,
          child: widget.content
        )
      ],
    );
  }
}
