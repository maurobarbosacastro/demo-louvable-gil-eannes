import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class UserVerificationDialog extends StatelessWidget {
  final void Function() onClickLetsGo;
  final void Function() showOnBoarding;

  const UserVerificationDialog({super.key, required this.onClickLetsGo, required this.showOnBoarding});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Image.asset(Assets.priceVsRewards),
        const SizedBox(height: 15),
        Text(translation(context).userVerified, style: Theme.of(context).textTheme.titleLarge),
        const SizedBox(height: 2),
        Text(translation(context).welcomeToTagpeak, style: Theme.of(context).textTheme.titleSmall, textAlign: TextAlign.center),
        const SizedBox(height: 10),
        SizedBox(
          width: 120,
          child: Button(
              label: translation(context).letsGo,
            onPressed: onClickLetsGo
          ),
        ),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const SizedBox(height: 15),
            Container(
                width: double.infinity,
                height: 1,
                color: AppColors.wildSand
            ),
            const SizedBox(height: 10),
            GestureDetector(
              onTap: showOnBoarding,
              child: Text(translation(context).howDoesItWork, style: Theme
                  .of(context)
                  .textTheme
                  .bodySmall),
            )
          ],
        )
      ],
    );
  }
}
