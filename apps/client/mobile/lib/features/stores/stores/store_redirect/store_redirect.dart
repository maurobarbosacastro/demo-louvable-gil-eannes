import 'dart:async';

import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';
import 'package:tagpeak/shared/providers/visit_store_provider.dart';
import 'package:tagpeak/shared/services/stores_service.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:url_launcher/url_launcher.dart';

class StoreRedirect extends ConsumerStatefulWidget {
  const StoreRedirect({super.key });

  @override
  ConsumerState<StoreRedirect> createState() => _StoreRedirectState();
}

class _StoreRedirectState extends ConsumerState<StoreRedirect> {

  PublicStoreModel? store;
  String? redirectURL;

  // Redirect countdown properties
  int secondsRemaining = 5;
  Timer? timer;

  @override
  void initState() {
    super.initState();

    store = ref.read(visitStoreStateProvider);

    if (store != null) {
      _fetchRedirectURL();
    }

    startCountdown();
  }

  @override
  void dispose() {
    timer?.cancel();
    super.dispose();
  }

  Future<void> _fetchRedirectURL() async {
    ref.read(storesServiceProvider).getStoreRedirectUrl(store!.uuid).then((url) {
      redirectURL = url;
    });
  }

  Future<void> _redirectToStoreURL() async {
    if (redirectURL == null) return;
    final Uri url = Uri.parse(redirectURL!);
    if (!await launchUrl(url)) {
      throw Exception('Could not launch $url');
    }
  }

  void startCountdown() {
    timer = Timer.periodic(const Duration(seconds: 1), (Timer t) {
      if (secondsRemaining == 0) {
        t.cancel();
        _redirectToStoreURL();

        context.go('${RouteNames.clientRouteLocation}${RouteNames.shopRouteLocation}');
      } else {
        setState(() {
          secondsRemaining--;
        });
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    if (store == null) {
      return const Center(
        child: CircularProgressIndicator(
          color: Colors.black,
        ),
      );
    }
    return Center(
      child: Padding(
        padding: const EdgeInsets.only(left: 15, top: 124, right: 15, bottom: kBottomNavigationBarHeight + 24),
        child: Column(
          spacing: 20,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [

            Column(
              children: [
                ClipRRect(
                  borderRadius:
                  const BorderRadius.vertical(top: Radius.circular(10)),
                  child: store!.logo != null && store!.logo!.isNotEmpty
                      ? Image.network(
                    store!.logo!,
                    height: 90,
                    fit: BoxFit.cover,
                    errorBuilder: (context, error, stackTrace) => Image.asset(
                      Assets.defaultStoreLogo,
                      height: 90,
                      fit: BoxFit.cover,
                    ),
                  )
                      : Image.asset(
                    Assets.defaultStoreLogo,
                    height: 90,
                    fit: BoxFit.cover,
                  ),
                ),

                const SizedBox(height: 50),

                Text(translation(context).redirectToStore(secondsRemaining, store?.name ?? ''), style: TextTheme.of(context).headlineSmall, textAlign: TextAlign.center),
                const SizedBox(height: 10),
                Text(translation(context).importantAlwaysAcceptCookies, style: TextTheme.of(context).bodyMedium, textAlign: TextAlign.center),
              ],
            ),

            RichText(
              text: TextSpan(
                children: <InlineSpan>[
                  TextSpan(
                    text: translation(context).clickHere,
                    style: Theme.of(context).textTheme.bodyMedium?.copyWith(
                      color: AppColors.licorice,
                      decoration: TextDecoration.underline,
                    ),
                    recognizer: TapGestureRecognizer()
                      ..onTap = () {
                        timer?.cancel();
                        _redirectToStoreURL();

                        context.go('${RouteNames.clientRouteLocation}${RouteNames.shopRouteLocation}');
                      },
                  ),
                  TextSpan(
                    text: translation(context).toGoStraightToStore,
                    style: Theme.of(context).textTheme.bodyMedium,
                  ),
                ],
              ),
            )

          ],
        ),
      ),
    );
  }

}
