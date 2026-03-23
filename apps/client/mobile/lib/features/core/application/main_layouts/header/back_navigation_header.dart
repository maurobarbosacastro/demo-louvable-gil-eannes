import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class BackNavigationHeader extends StatelessWidget implements PreferredSizeWidget {
  final String title;
  final String? backPathLocation;
  final String? backPathParentLocation;
  const BackNavigationHeader({super.key, required this.title, this.backPathLocation, this.backPathParentLocation });

  @override
  Widget build(BuildContext context) {
    return AppBar(
        systemOverlayStyle: SystemUiOverlayStyle.dark,
        backgroundColor: Colors.transparent,
        forceMaterialTransparency: true,
        elevation: 0,
        centerTitle: true,
        title: Padding(
          padding: const EdgeInsets.only(left: 4.0, top: 30.0, right: 4.0, bottom: 20.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              GestureDetector(
                onTap: () => context.go(backPathLocation!, extra: { 'previousRoute': backPathParentLocation }),
                child: SvgPicture.asset(
                  Assets.simpleArrowLeft,
                  colorFilter: const ColorFilter.mode(AppColors.licorice, BlendMode.srcIn),
                  width: 30,
                ),
              ),
              Text(title, style: Theme.of(context).textTheme.titleLarge),
              const SizedBox(width: 30),
            ],
          ),
        )
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(70.0);
}
