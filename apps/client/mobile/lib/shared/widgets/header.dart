import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_svg/svg.dart';

class Header extends StatelessWidget implements PreferredSizeWidget {
  final Widget? endContent;
  final Color backgroundColor;
  const Header({super.key, this.endContent, this.backgroundColor = Colors.transparent});

  @override
  Widget build(BuildContext context) {
    return AppBar(
      systemOverlayStyle: SystemUiOverlayStyle.dark,
      backgroundColor: backgroundColor,
      forceMaterialTransparency: true,
      elevation: 0,
      toolbarOpacity: 1,
      centerTitle: true,
      title: Padding(
        padding: const EdgeInsets.only(left: 4.0, top: 30.0, right: 4.0, bottom: 20.0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: <Widget>[
            SvgPicture.asset(
              'lib/assets/images/logo/logo_extended.svg',
            ),
            Visibility(
              visible: endContent != null,
              child: endContent ?? const SizedBox.shrink(),
            ),
          ],
        ),
      )
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(70.0);
}
