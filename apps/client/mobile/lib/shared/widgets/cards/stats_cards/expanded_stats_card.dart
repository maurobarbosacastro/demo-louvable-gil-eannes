import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/constants/colors.dart';

// ToDO make this component generic (future work
// ToDO create 3 cards: compact, regular & expanded (future work
class ExpandedStatsCard extends StatelessWidget {
  const ExpandedStatsCard({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(10.0),
      decoration: const BoxDecoration(
        color: AppColors.grandis,
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            spacing: 6,
            children: [
              SvgPicture.asset('lib/assets/icons/shopping_bag.svg'),
              Text(
                translation(context).purchasedSamsung,
                style: TextStyle(fontSize: 10, color: AppColors.licorice),
              ),
            ],
          ),
          const SizedBox(height: 15),
          Text(
            '345.99€',
            style: Theme.of(context).textTheme.headlineSmall?.copyWith(
              color: AppColors.licorice.withOpacity(0.7),
            ),
          ),
          const SizedBox(height: 5),
          Container(
            height: 1,
            color: AppColors.waterloo,
          ),
          const SizedBox(height: 5),
          Row(
            spacing: 6,
            children: [
              SvgPicture.asset('lib/assets/icons/trending_up.svg'),
              Text(
                translation(context).gotBack,
                style: TextStyle(fontSize: 10, color: AppColors.licorice),
              ),
            ],
          ),
          Text(
              '65.99€',
              style: Theme.of(context).textTheme.headlineMedium
          ),
        ],
      ),
    );
  }
}
