import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class WithoutTransactions extends StatelessWidget {
  const WithoutTransactions({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 80),
        child: Column(children: [
          Image.asset(Assets.priceVsRewards),
          const SizedBox(height: 10),
          Text(
            translation(context).nothingHere,
            style: Theme.of(context).textTheme.titleLarge,
          ),
          const SizedBox(height: 10),
          Text(translation(context).makeYourFirstPurchase,
              style: Theme.of(context).textTheme.titleSmall,
              textAlign: TextAlign.center),
          const SizedBox(height: 10),
          Button(
            label: translation(context).newPurchase,
            onPressed: () {},
          )
        ])
    );
  }
}
