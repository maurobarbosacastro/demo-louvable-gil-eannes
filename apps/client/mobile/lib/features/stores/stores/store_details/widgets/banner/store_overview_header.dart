import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/stores/stores/store_details/widgets/banner/store_banner.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class StoreOverviewHeader extends StatelessWidget {
  final PublicStoreModel storeDetails;
  final VoidCallback goToStore;

  const StoreOverviewHeader({super.key, required this.storeDetails, required this.goToStore});

  @override
  Widget build(BuildContext context) {
    return Stack(children: [
      StoreBanner(storeDetails: storeDetails),
      Positioned(
          bottom: 30,
          left: 25,
          right: 25,
          child: Column(
            spacing: 15,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(storeDetails.name,
                  style: TextTheme.of(context)
                      .headlineMedium
                      ?.copyWith(color: AppColors.wildSand)),
              Button(
                  label: translation(context).goToStore,
                  mode: Mode.youngBlue,
                  onPressed: goToStore,
                  fontSize: 15,
                  width: 120,
                  height: 40),
            ],
          ))
    ]);
  }
}
