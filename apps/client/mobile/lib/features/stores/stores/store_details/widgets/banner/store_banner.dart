import 'package:flutter/material.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class StoreBanner extends StatelessWidget {
  final PublicStoreModel storeDetails;
  const StoreBanner({super.key, required this.storeDetails});

  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      child: (storeDetails.banner != null && storeDetails.banner!.isNotEmpty)
          ? Stack(
        children: [
          ConstrainedBox(
            constraints: const BoxConstraints(
              maxHeight: 250,
              maxWidth: double.infinity,
            ),
            child: Image.network(
              storeDetails.banner!,
              width: double.infinity,
              fit: BoxFit.fitWidth,
              errorBuilder: (context, error, stackTrace) => const DefaultBanner(),
            ),
          ),
          Positioned.fill(
            child: Container(
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  begin: Alignment.topCenter,
                  end: Alignment.bottomCenter,
                  colors: [
                    Colors.transparent,
                    Colors.black.withAlpha(178),
                  ],
                ),
              ),
            ),
          ),
        ],
      ) : const DefaultBanner()
    );
  }
}

class DefaultBanner extends StatelessWidget {
  const DefaultBanner({super.key});

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        ConstrainedBox(
          constraints: const BoxConstraints(
            maxHeight: 250,
            maxWidth: double.infinity,
          ),
          child: Image.asset(
            Assets.defaultBanner,
            width: double.infinity,
            fit: BoxFit.fitWidth,
          ),
        ),
        Positioned.fill(
          child: Container(
            decoration: BoxDecoration(
              gradient: LinearGradient(
                begin: Alignment.topCenter,
                end: Alignment.bottomCenter,
                colors: [
                  Colors.transparent,
                  Colors.black.withAlpha(178),
                ],
              ),
            ),
          ),
        ),
      ],
    );
  }
}
