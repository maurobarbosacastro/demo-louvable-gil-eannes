import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/refer_earn/how_it_works/how_it_works.dart';
import 'package:tagpeak/features/refer_earn/your_referrals/your_referals.dart';
import 'package:tagpeak/shared/widgets/slide_tabs/slide_tabs.dart';

class ReferEarn extends StatelessWidget {
  const ReferEarn({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 100),
          Text(
            translation(context).referEarn,
            style: Theme.of(context).textTheme.headlineLarge,
          ),
          const SizedBox(height: 5),
          SlideTabs(
            tabs: [
              Tab(text: translation(context).yourReferrals),
              Tab(text: translation(context).howItWorks)
            ],
            views: const [
              YourReferrals(),
              HowItWorks()
            ],
          ),
        ],
      ),
    );
  }
}
