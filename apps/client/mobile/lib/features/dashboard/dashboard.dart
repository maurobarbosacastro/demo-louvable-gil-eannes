import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/dashboard/balance/balance.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/dashboard/store_visits/store_visits.dart';
import 'package:tagpeak/features/dashboard/your_rewards/your_rewards.dart';
import 'package:tagpeak/shared/models/user_model/user_model.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_bottom_dialog.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_dialog.dart';
import 'package:tagpeak/shared/widgets/dialogs/on_boarding_dialog.dart';
import 'package:tagpeak/shared/widgets/slide_tabs/slide_tabs.dart';
import 'package:tagpeak/features/notifications/providers/fcm_provider.dart';

import '../../shared/widgets/dialogs/user_verification_dialog.dart';

class Dashboard extends ConsumerStatefulWidget {
  const Dashboard({super.key});

  @override
  ConsumerState<Dashboard> createState() => _DashboardState();
}

class _DashboardState extends ConsumerState<Dashboard> {
  late UserModel user;

  ///
  /// User verification and onboarding dialog keys / vars
  ///
  final userVerificationDialogKey = GlobalKey();
  final onBoardingDialogKey = GlobalKey();

  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) async {
      user = await ref.read(userNotifier).getUser();
      if (!mounted) return;

      if (!user.isVerified) {
        showUserVerification();
        return;
      }

      // ToDO remove this comment after confirming this behavior
      /* if (!user.onboardingFinished) {
        showOnBoarding();
      } */

      // Trigger FCM permission request for client users
      triggerFCMPermissionRequest(ref);
    });
  }

  ///
  /// User verification and onboarding dialog methods
  ///

  void submitUserVerification() {
    ref.read(userNotifier).setEmailVerified(user.uuid, true, userVerificationDialogKey.currentContext!);
  }

  void startShopping() {
    final context = userVerificationDialogKey.currentContext ?? onBoardingDialogKey.currentContext;
    if (context != null) {
      ref.read(userNotifier).setEmailVerified(user.uuid, true, context).then((_) {
        final onBoardingContext = onBoardingDialogKey.currentContext;
        if (userVerificationDialogKey.currentContext != null &&
            onBoardingContext != null &&
            onBoardingContext.mounted) {
          Navigator.pop(onBoardingContext);
        }
      });
    }
  }

  void showOnBoarding() {
    GenericBottomDialog.openBottomDialog(
        padding: const EdgeInsets.symmetric(horizontal: 20),
        context,
        content: OnBoardingDialog(
            key: onBoardingDialogKey,
            onClickStartShopping: startShopping
        ),
        canDismissOnTapOutside: false,
        hasCloseIcon: userVerificationDialogKey.currentContext != null,
        iconOffset: const IconOffset(top: 21, right: 0)
    );
  }

  void showUserVerification() {
    GenericDialog.openDialog(
        context,
        content: UserVerificationDialog(
          key: userVerificationDialogKey,
          onClickLetsGo: submitUserVerification,
          showOnBoarding: showOnBoarding,
        ),
        canDismissOnTapOutside: false,
        hasCloseIcon: false
    );
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 100),
          Text(
            translation(context).dashboard,
            style: Theme.of(context).textTheme.headlineLarge,
          ),
          const SizedBox(height: 5),
          SlideTabs(
            tabs: [
              Tab(text: translation(context).yourRewards),
              Tab(text: translation(context).storeVisits),
              Tab(text: translation(context).balance)
            ],
            views: const [
              YourRewards(),
              StoreVisits(),
              Balance()
            ],
          ),
        ],
      ),
    );
  }
}
