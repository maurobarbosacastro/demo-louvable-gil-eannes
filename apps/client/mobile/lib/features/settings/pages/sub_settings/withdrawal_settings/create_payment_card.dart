import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class CreatePaymentCard extends StatelessWidget {
  final VoidCallback? onAddPaymentMethodPressed;
  const CreatePaymentCard({super.key, this.onAddPaymentMethodPressed});

  @override
  Widget build(BuildContext context) {
    return Container(
      clipBehavior: Clip.hardEdge,
      decoration: BoxDecoration(
          color: AppColors.grandis,
          borderRadius: BorderRadius.circular(10)
      ),
      padding: const EdgeInsets.only(left: 20, top: 28),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          Padding(
            padding: const EdgeInsets.only(bottom: 20),
            child: Button(
              label: translation(context).createPaymentMethod,
              fontSize: 13,
              padding: const EdgeInsets.symmetric(horizontal: 10),
              onPressed: onAddPaymentMethodPressed,
            ),
          ),
          Image.asset(Assets.brandsImage, width: 150, opacity: const AlwaysStoppedAnimation(.6))
        ],
      ),
    );
  }
}
