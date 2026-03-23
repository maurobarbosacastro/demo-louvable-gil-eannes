import 'package:flutter/material.dart';

class ScrollableBody extends StatelessWidget {
  final Widget child;
  final bool bodyBottomSpacing;

  const ScrollableBody({super.key, required this.child, this.bodyBottomSpacing = true});

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      padding: EdgeInsets.only(
          bottom: kBottomNavigationBarHeight + (bodyBottomSpacing ? 24 : - 24),
          top: 20
      ),
      child: child
    );
  }
}
