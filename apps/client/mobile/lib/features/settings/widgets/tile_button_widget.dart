import 'package:flutter/material.dart';

class TileButton extends StatelessWidget {
  final String label;
  final VoidCallback onTap;
  final IconData icon;

  final IconData endIcon;

  const TileButton({super.key, required this.label, required this.onTap, required this.icon, this.endIcon = Icons.arrow_forward_ios});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 50,
      padding: const EdgeInsets.only(top: 4.0),
      child: InkWell(
        splashColor: Colors.grey,
        highlightColor: Colors.grey,
        onTap: onTap,
        child: Row(
          children: [
            Icon(icon),
            const SizedBox(
              width: 12,
            ),
            Expanded(
              child: Text(
                label,
                style: const TextStyle(fontWeight: FontWeight.w500, fontSize: 16),
              ),
            ),
            Icon(endIcon),
          ],
        ),
      ),
    );
  }
}
