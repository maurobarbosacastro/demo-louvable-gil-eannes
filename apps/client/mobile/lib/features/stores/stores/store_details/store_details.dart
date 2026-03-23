import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_html/flutter_html.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/features/stores/stores/store_details/widgets/banner/store_overview_header.dart';
import 'package:tagpeak/features/stores/stores/store_details/widgets/redirect_to_store_dialog.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';
import 'package:tagpeak/shared/providers/visit_store_provider.dart';
import 'package:tagpeak/shared/services/stores_service.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_bottom_dialog.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class StoreDetails extends ConsumerStatefulWidget {
  final String storeID;

  const StoreDetails({super.key, required this.storeID});

  @override
  ConsumerState<StoreDetails> createState() => _StoreDetailsState();
}

class _StoreDetailsState extends ConsumerState<StoreDetails> {
  PublicStoreModel? storeDetails;
  bool isHiddenUntilNextMonth = false;
  final _storage = const FlutterSecureStorage();
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();

    _fetchStoreDetails();

    _loadIsHiddenUntilNextMonth();
  }

  void _fetchStoreDetails() {
    ref
        .read(storesServiceProvider)
        .getPublicStoreByID(widget.storeID)
        .then((store) {
          setState(() {
            storeDetails = store;
            _isLoading = false;
          });
          ref.read(visitStoreNotifierProvider).saveVisitStore(store);
        }).catchError((error) {
          setState(() {
            _isLoading = false;
          });
        });
  }

  void onClickGoToStore(String uuid) {
    if (!isHiddenUntilNextMonth) {
      GenericBottomDialog.openBottomDialog(context,
          content: RedirectToStoreDialog(
              onSubmit: () {
                _saveIsHiddenUntilNextMonth();
                context.pop();
                context.pushNamed(RouteNames.redirectDetailsRouteName, extra: { 'previousRoute': "/client/${RouteNames.shopRouteName}/${RouteNames.storeDetailsRouteName}/$uuid" });
              }),
          hasCloseIcon: true,
          canDismissOnTapOutside: true,
          padding:
          const EdgeInsets.only(right: 20, left: 20, top: 20, bottom: 50));
    } else {
      context.pushNamed(RouteNames.redirectDetailsRouteName, extra: {
        "previousRoute": "/client/${RouteNames.shopRouteName}/${RouteNames.storeDetailsRouteName}/$uuid",
        "previousRouteParent": '/client/${RouteNames.shopRouteName}'
      });
    }
  }

  _saveIsHiddenUntilNextMonth() async {
    final isHidden = ref.read(isHiddenUntilNextMonthProvider);
    final data = {
      'checked': isHidden,
      'checkedDate': DateTime.now().toIso8601String(),
    };
    await _storage.write(
      key: 'KEY_HIDDEN_UNTIL_NEXT_MONTH',
      value: jsonEncode(data),
    );
  }

  Future<void> _loadIsHiddenUntilNextMonth() async {
    final jsonString = await _storage.read(key: 'KEY_HIDDEN_UNTIL_NEXT_MONTH');

    if (jsonString == null) {
      isHiddenUntilNextMonth = false;
      return;
    }

    final Map<String, dynamic> data = jsonDecode(jsonString);
    final bool checked = data['checked'] ?? false;
    final String? checkedDateStr = data['checkedDate'];

    if (checkedDateStr == null) {
      isHiddenUntilNextMonth = false;
      return;
    }

    final DateTime checkedDate = DateTime.parse(checkedDateStr);
    final int difference = DateTime.now().difference(checkedDate).inDays;

    if (difference > 30) {
      isHiddenUntilNextMonth = false;
    } else {
      isHiddenUntilNextMonth = checked;
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return const Center(
        child: CircularProgressIndicator(
          color: Colors.black,
        ),
      );
    }
    if (storeDetails != null) {
      return Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          StoreOverviewHeader(storeDetails: storeDetails!, goToStore: () => onClickGoToStore(storeDetails!.uuid)),
          Expanded(
            child: ScrollableBody(
              child: Padding(
                padding: const EdgeInsets.symmetric(horizontal: 20),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    IntrinsicHeight(
                      child: Row(
                        crossAxisAlignment: CrossAxisAlignment.stretch,
                        children: [
                          Expanded(
                            child: Container(
                              padding: const EdgeInsets.only(left: 10, bottom: 20),
                              decoration: BoxDecoration(
                                border: Border(
                                  left: BorderSide(color: AppColors.licorice.withAlpha(50), width: 0.5),
                                ),
                              ),
                              child: Column(
                                spacing: 10,
                                mainAxisSize: MainAxisSize.max,
                                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(
                                    translation(context).timeToActivateReward,
                                    style: Theme.of(context)
                                        .textTheme
                                        .titleSmall
                                        ?.copyWith(color: AppColors.licorice),
                                  ),
                                  Text(
                                    '${storeDetails?.averageRewardActivationTime}',
                                    style: Theme.of(context)
                                        .textTheme
                                        .bodyLarge
                                        ?.copyWith(color: AppColors.licorice),
                                  ),
                                ],
                              ),
                            ),
                          ),
                          const SizedBox(width: 30),
                          Expanded(
                            child: Container(
                              padding: const EdgeInsets.only(left: 10, bottom: 20),
                              decoration: BoxDecoration(
                                border: Border(
                                  left: BorderSide(color: AppColors.licorice.withAlpha(50), width: 0.5),
                                ),
                              ),
                              child: Column(
                                mainAxisSize: MainAxisSize.max,
                                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(
                                    translation(context).initialRewardAmount,
                                    style: Theme.of(context)
                                        .textTheme
                                        .titleSmall
                                        ?.copyWith(color: AppColors.licorice),
                                  ),
                                  Column(
                                    crossAxisAlignment: CrossAxisAlignment.start,
                                    children: [
                                      Text(
                                        '${storeDetails?.percentageCashout}%',
                                        style: Theme.of(context)
                                            .textTheme
                                            .headlineMedium
                                            ?.copyWith(color: AppColors.licorice),
                                      ),
                                      Text(
                                        translation(context).potentialToGrow,
                                        style: Theme.of(context)
                                            .textTheme
                                            .titleSmall
                                            ?.copyWith(color: AppColors.licorice),
                                      ),
                                    ],
                                  ),
                                ],
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    const SizedBox(height: 20),

                    Text(
                      translation(context).brand,
                      style: Theme.of(context).textTheme.headlineSmall,
                    ),
                    const SizedBox(height: 10),
                    Text(
                      storeDetails?.shortDescription ?? '',
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                    const SizedBox(height: 10),
                    Html(data: storeDetails?.termsAndConditions ?? '')
                  ],
                ),
              ),
            ),
          ),
        ],
      );
    }
    return const SizedBox.shrink();
  }
}
