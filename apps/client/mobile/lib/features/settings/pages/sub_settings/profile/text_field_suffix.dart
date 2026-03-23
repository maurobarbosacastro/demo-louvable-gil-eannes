import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class TextFieldSuffix extends StatelessWidget {
  final bool inEditMode;
  final void Function()? handleClose;
  final void Function()? handleSave;
  final void Function()? handleEdit;

  const TextFieldSuffix({super.key, this.inEditMode = false, this.handleClose, this.handleSave, this.handleEdit});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 32,
      child: inEditMode ? Row(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          IconButton(
              padding: EdgeInsets.zero,
              onPressed: handleClose,
              icon: SvgPicture.asset(Assets.circleClose, colorFilter: const ColorFilter.mode(AppColors.licorice, BlendMode.srcIn))
          ),
          TextButton(
              style: TextButton.styleFrom(
                maximumSize: const Size(54, 32),
                minimumSize: const Size(0, 32),
                tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                padding: const EdgeInsets.symmetric(horizontal: 10)
              ),
              onPressed: handleSave,
              child: Text(translation(context).save, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
          )
        ],
      ) : TextButton(
          style: TextButton.styleFrom(
              maximumSize: const Size(54, 32),
              minimumSize: const Size(0, 32),
              tapTargetSize: MaterialTapTargetSize.shrinkWrap,
              padding: const EdgeInsets.symmetric(horizontal: 10)
          ),
          onPressed: handleEdit,
          child: Text(translation(context).edit, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
      ),
    );
  }
}
