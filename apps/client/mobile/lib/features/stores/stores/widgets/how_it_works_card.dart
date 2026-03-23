import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/extensions/strings_extension.dart';

class ProcessStepCard extends StatelessWidget {
  const ProcessStepCard({super.key});

  @override
  Widget build(BuildContext context) {
    final List<String> steps = [
      translation(context).makeAPurchase,
      translation(context).invest,
      translation(context).followInvestment,
      translation(context).receiveReward
    ];

    return Container(
      decoration: BoxDecoration(
          color: AppColors.grandis,
          borderRadius: BorderRadius.circular(10.0)
      ),
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 25),
      child: Column(
        spacing: 10,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(translation(context).howItWorks.toCapitalize(), style: TextTheme
              .of(context)
              .headlineSmall
              ?.copyWith(color: AppColors.licorice)),

          const SizedBox(height: 10),

          ...List.generate(steps.length, (index) {
            return ProcessStep(index: index + 1, label: steps[index]);
          })
        ],
      ),
    );
  }
}

class ProcessStep extends StatelessWidget {
  final int index;
  final String label;

  const ProcessStep({super.key, required this.index, required this.label});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 15),
      width: double.infinity,
      decoration: BoxDecoration(
          color: AppColors.wildSand,
          borderRadius: BorderRadius.circular(15.0)
      ),
      child: Column(
        spacing: 20,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            height: 40,
            width: 40,
            decoration: const BoxDecoration(
              shape: BoxShape.circle,
              color: AppColors.grandis,
            ),
            child: Center(child: Text('0$index', style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))),
          ),
          Text(label, style: TextTheme.of(context).bodyMedium?.copyWith(color: AppColors.licorice))
        ],
      ),
    );
  }
}
