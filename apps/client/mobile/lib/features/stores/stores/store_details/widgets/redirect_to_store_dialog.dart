import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/screens/login/widgets/checkbox_widget.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/providers/visit_store_provider.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class RedirectToStoreDialog extends ConsumerWidget {
  final VoidCallback onSubmit;
  const RedirectToStoreDialog({super.key, required this.onSubmit});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Column(
      spacing: 20,
      children: [
        const SizedBox(height: 40),
        Container(
          padding: const EdgeInsets.symmetric(horizontal: 15, vertical: 20),
          decoration: BoxDecoration(
            color: AppColors.wildSand,
            borderRadius: BorderRadius.circular(20)
          ),
          child: Column(
            spacing: 10,
            children: [
              RedirectToStoreStep(icon: Assets.redirect, label: translation(context).goToStoreFirstInstruction),
              RedirectToStoreStep(icon: Assets.cookie, label: translation(context).goToStoreSecondInstruction),
              RedirectToStoreStep(icon: Assets.blockers, label: translation(context).goToStoreThirdInstruction),
              RedirectToStoreStep(icon: Assets.shoppingCart, label: translation(context).goToStoreFourthInstruction)
            ],
          ),
        ),
        Container(
          width: double.infinity,
          color: AppColors.wildSand,
          height: 1
        ),
        Column(
          spacing: 15,
          children: [
            Text(translation(context).understoodGoToStoreInstructions,
              style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice)),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Button(label: 'Back', onPressed: () => context.pop(), mode: Mode.outlinedDark,
                    icon: SvgPicture.asset(Assets.arrowLeft, colorFilter: const ColorFilter.mode(AppColors.licorice, BlendMode.srcIn))),
                Button(label: 'Go to store', onPressed: () {
                  context.pop();
                  onSubmit();
                }, mode: Mode.youngBlue, fontSize: 15, height: 35),
              ],
            ),
            CheckboxWidget(
              key: const Key("HiddenUntilNextMonth"),
              value: ref.watch(isHiddenUntilNextMonthProvider),
              onChanged: (value) => ref
                  .read(isHiddenUntilNextMonthProvider.notifier)
                  .update((state) => value ?? false),
              label: translation(context).doNotShowForAnother30Days,
            ),
          ],
        )
      ],
    );
  }
}

class RedirectToStoreStep extends StatelessWidget {
  final String icon;
  final String label;
  const RedirectToStoreStep({super.key, required this.icon, required this.label});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(15),
          color: Colors.white
      ),
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 10),
      child: Row(
        spacing: 10,
        children: [
          Container(
            height: 40,
            width: 40,
            decoration: const BoxDecoration(
              shape: BoxShape.circle,
              color: AppColors.wildSand
            ),
            child: Center(child: SvgPicture.asset(icon)),
          ),
          Expanded(child: Text(label, style: TextTheme.of(context).titleSmall?.copyWith(color: AppColors.licorice)))
        ],
      ),
    );
  }
}
