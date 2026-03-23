import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/dashboard/store_visits/store_visits_table.dart';
import 'package:tagpeak/shared/widgets/date_picker/date_picker.dart';
import 'package:tagpeak/shared/models/store_model/store_model.dart';
import 'package:tagpeak/shared/models/store_visit_model/store_visit_model.dart';
import 'package:tagpeak/shared/services/stores_service.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/utils/constants/assets.dart';

class StoreVisits extends ConsumerStatefulWidget {
  const StoreVisits({super.key});

  @override
  ConsumerState<StoreVisits> createState() => _StoreVisitsState();
}

class _StoreVisitsState extends ConsumerState<StoreVisits> {
  List<StoreModel> stores = [];
  List<StoreVisitModel> storesVisits = [];

  // properties to filter store visits
  String? store;
  DateTime? startDate;
  DateTime? endDate;

  int numPages = 0;
  int selectedPage = 0;

  bool _isLoading = true;
  bool _isSearched = false;
  TextEditingController startDateController = TextEditingController();
  TextEditingController endDateController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _fetchStores();
    _fetchStoresVisits();
  }

  void _fetchStores() {
    ref.read(storesServiceProvider).getStoresbyUser().then((paginatedStores) {
      setState(() {
        stores = paginatedStores.data;
      });
    });
  }

  void _fetchStoresVisits() async {
    if (!_isLoading) setState(() => _isLoading = true);

    ref.read(storesServiceProvider).getStoresVisitedByUser(
      page: selectedPage + 1,
      filtersBody: {
        if (store != null) 'store': store,
        if (startDate != null)
          'startDate': startDate!.toUtc().toIso8601String(),
        if (endDate != null) 'endDate': endDate!.toUtc().toIso8601String(),
      },
    ).then((paginatedStores) {
      setState(() {
        storesVisits = paginatedStores.data;
        numPages = paginatedStores.totalPages;

        _isLoading = false;
      });
    }).catchError((error) {
      setState(() {
        _isLoading = false;
      });
    });


  }

  List<DropdownOption> get dropdownEntries => stores
      .map((store) => DropdownOption(label: store.name, value: store.uuid))
      .toList();

  void handlePageChanged(int page) {
    setState(() {
      selectedPage = page;
    });
    _fetchStoresVisits();
  }

  @override
  Widget build(BuildContext context) {
    return ScrollableBody(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        spacing: 10,
        children: [
          SizedBox(
            width: double.infinity,
            child: Dropdown(
                entries: dropdownEntries,
                onChanged: (value) {
                  store = value;
                },
                isExpanded: true,
                label: translation(context).selectStore,
                onClickClear: (_) {
                  setState(() {
                    store = null;
                  });
                  _fetchStoresVisits();

                  if (store == null && startDate == null && endDate == null) {
                    setState(() => _isSearched = false);
                  }
                },
            ),
          ),
          Row(
            spacing: 10,
            children: [
              Expanded(
                child: DatePicker(
                  controller: startDateController,
                  placeholder: translation(context).dateFrom,
                  prefixIcon: SvgPicture.asset(Assets.calendar),
                  onSelectedDate: (date) {
                    setState(() {
                      startDate = date;
                    });
                  },
                  selectedDate: startDate,
                  onClickClear: () => setState(() {
                    startDate = null;

                    if (store == null && startDate == null && endDate == null) {
                      _isSearched = false;
                    }

                    _fetchStoresVisits();
                  }),
                ),
              ),
              Expanded(
                child: DatePicker(
                  controller: endDateController,
                  placeholder: translation(context).dateTo,
                  prefixIcon: SvgPicture.asset(Assets.arrowRight),
                  onSelectedDate: (date) {
                    setState(() {
                       endDate = date;
                       _fetchStoresVisits();
                    });
                  },
                  selectedDate: endDate,
                  onClickClear: () => setState(() {
                    endDate = null;

                    if (store == null && startDate == null && endDate == null) {
                      _isSearched = false;
                    }

                    _fetchStoresVisits();
                  }),
                ),
              ),
            ],
          ),
          Button(
            label: translation(context).search,
            mode: Mode.outlinedDark,
            onPressed: () {
              _fetchStoresVisits();
              setState(() => _isSearched = true);

              if (store == null && startDate == null && endDate == null) {
                setState(() => _isSearched = false);
              }
            },
          ),
          const SizedBox(height: 20),
          StoreVisitsTable(
              visits: storesVisits,
              onPageChanged: handlePageChanged,
              numPages: numPages,
              isLoading: _isLoading,
              isSearched: _isSearched
          )
        ],
      ),
    );
  }
}
