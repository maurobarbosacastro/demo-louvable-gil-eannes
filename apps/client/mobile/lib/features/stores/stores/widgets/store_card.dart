import 'package:flutter/material.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class StoreCard extends StatelessWidget {
  final PublicStoreModel store;
  final VoidCallback onClick;

  const StoreCard({super.key, required this.store, required this.onClick});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onClick,
      child: Container(
        clipBehavior: Clip.hardEdge,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(10),
          color: AppColors.wildSand,
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            ClipRRect(
              borderRadius:
                  const BorderRadius.vertical(top: Radius.circular(10)),
              child: store.logo != null && store.logo!.isNotEmpty
                  ? Image.network(
                      store.logo!,
                      height: 115,
                      fit: BoxFit.cover,
                      errorBuilder: (context, error, stackTrace) => Image.asset(
                        Assets.defaultStoreLogo,
                        height: 115,
                        fit: BoxFit.cover,
                      ),
                    )
                  : Image.asset(
                      Assets.defaultStoreLogo,
                      height: 115,
                      fit: BoxFit.cover,
                    ),
            ),
            Padding(
              padding: const EdgeInsets.all(12.0),
              child: Text(
                store.name,
                style: const TextStyle(
                  color: AppColors.licorice,
                  fontSize: 15,
                ),
                textAlign: TextAlign.center,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
