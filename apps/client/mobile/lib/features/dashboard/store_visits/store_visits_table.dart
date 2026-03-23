import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/store_visit_model/store_visit_model.dart';
import 'package:tagpeak/shared/widgets/table/tgpk_data_table.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/date_helpers.dart';

class StoreVisitsTable extends StatelessWidget {
  final List<StoreVisitModel> visits;
  final Function(int) onPageChanged;
  final int numPages;
  final bool isLoading;
  final bool isSearched;

  const StoreVisitsTable(
      {super.key, required this.visits, required this.onPageChanged, required this.numPages, this.isLoading = false, this.isSearched = false});

  @override
  Widget build(BuildContext context) {
    return TgpkDataTable<StoreVisitModel>(
        rowHeight: 80,
        columns: [
          DataColumn(label: Text(translation(context).refId,
              style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
          DataColumn(label: Text(translation(context).storeName,
              style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
          DataColumn(label: Text(translation(context).purchased,
              style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
          DataColumn(headingRowAlignment: MainAxisAlignment.end,
              label: Text(translation(context).dateTime,
                  style: const TextStyle(fontSize: 13, color: AppColors.waterloo))),
        ],
        rowBuilder: (visit, constraints) =>
        [
          SizedBox(
              width: constraints.maxWidth * 0.15,
              child: Text(visit.reference,
                  style: const TextStyle(fontSize: 16, color: AppColors.licorice))
          ),
          SizedBox(
              width: constraints.maxWidth * 0.35,
              child: Text(visit.store.name,
                  style: const TextStyle(fontSize: 16,
                      color: AppColors.licorice,
                      fontWeight: FontWeight.w700))),
          visit.purchased ?
          SvgPicture.asset(
              Assets.circleCheck, colorFilter: const ColorFilter.mode(AppColors.licorice, BlendMode.srcIn), width: 20)
              : SvgPicture.asset(
              Assets.circleClose, colorFilter: const ColorFilter.mode(AppColors.licorice, BlendMode.srcIn), height: 20),
          SizedBox(
            width: double.infinity,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Text(formatDate(visit.dateTime),
                    style: const TextStyle(fontSize: 16, color: AppColors.waterloo), textAlign: TextAlign.end),
                Text(formatDate(visit.dateTime, pattern: 'HH:mm:ss'),
                    style: const TextStyle(fontSize: 16, color: AppColors.waterloo), textAlign: TextAlign.end)
              ],
            ),
          )
        ],
        items: visits,
        onPageChanged: onPageChanged,
        pages: numPages,
        isLoading: isLoading,
        hasActiveFilters: isSearched,
        errorStateMessage: translation(context).withoutResultsStoreVisits,
        emptyStateMessage: translation(context).withoutResultsAfterSearchStoreVisits
    );
  }
}
