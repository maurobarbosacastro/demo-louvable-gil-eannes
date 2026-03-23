import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:geocoding/geocoding.dart';
import 'package:geolocator/geolocator.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/features/stores/stores/widgets/stores_filters.dart';
import 'package:tagpeak/features/stores/stores/widgets/store_card.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';
import 'package:tagpeak/shared/providers/stores_provider.dart';
import 'package:tagpeak/shared/providers/visit_store_provider.dart';
import 'package:tagpeak/shared/services/country_service.dart';
import 'package:tagpeak/shared/services/stores_service.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/debounce.dart';
import 'package:tagpeak/utils/logging.dart';

class StoresFiltersCriteria {
  String? country;
  String? category;
  String? name;
  String? sort;

  StoresFiltersCriteria(this.country, this.category, this.name, this.sort);
}

class Stores extends ConsumerStatefulWidget {
  const Stores({Key? key}) : super(key: key);

  @override
  _StoresState createState() => _StoresState();
}

class _StoresState extends ConsumerState<Stores> {
  List<PublicStoreModel> stores = [];
  List<DropdownOption> countries = [];
  List<DropdownOption> categories = [];

  final TextEditingController countryController = TextEditingController();
  final TextEditingController categoryController = TextEditingController();
  final TextEditingController nameController = TextEditingController();
  final ScrollController _scrollController = ScrollController();

  final debounce = Debounce(milliseconds: 1000);

  StoresFiltersCriteria storesFiltersCriteria =
      StoresFiltersCriteria(null, null, null, 'most-popular');

  bool isLoadingLocation = true;
  bool isLoadingStores = true;
  bool isLoadingMore = false;
  bool hasMoreStores = true;

  int currentPage = 0;
  final int storesPerPage = 10;

  Future<String?> getUserCountryCode() async {
    try {
      bool serviceEnabled = await Geolocator.isLocationServiceEnabled();
      if (!serviceEnabled) {
        return null;
      }

      LocationPermission permission = await Geolocator.checkPermission();
      if (permission == LocationPermission.denied) {
        permission = await Geolocator.requestPermission();
        if (permission == LocationPermission.denied) {
          return null;
        }
      }

      if (permission == LocationPermission.deniedForever) {
        return null;
      }

      Position position = await Geolocator.getCurrentPosition(
        desiredAccuracy: LocationAccuracy.low,
        timeLimit: const Duration(seconds: 5),
      );

      List<Placemark> placemarks = await placemarkFromCoordinates(
        position.latitude,
        position.longitude,
      );

      if (placemarks.isNotEmpty) {
        String? countryCode = placemarks.first.isoCountryCode;
        return countryCode;
      }

      return null;
    } catch (e) {
      print('Error getting user location: $e');
      return null;
    }
  }

  Future<void> _initializeLocation() async {
    String? countryCode = ref.read(StoreCountryProvider);
    countryCode ??= await getUserCountryCode();

    setState(() {
      if (countryCode != null) {
        storesFiltersCriteria.country = countryCode;
        ref.read(StoreCountryProvider.notifier).update( (_) => countryCode);
      }

      isLoadingLocation = false;
    });

    await ref.read(countryServiceProvider).getCountries().then((countries) {
      setState(() {
        this.countries = countries.data.map((country) {
          return DropdownOption(
              label: country.name,
              value: country.abbreviation,
              selected: storesFiltersCriteria.country != null ? country.abbreviation == storesFiltersCriteria.country : false);
        }).toList();

        final selectedCountries = this.countries.where((country) => country.value == countryCode);
        if (selectedCountries.isNotEmpty) {
          countryController.text = selectedCountries.first.label;
        } else {
          countryController.text = '';
        }

      });
    });

    fetchPublicStores();
  }

  void fetchPublicStores({bool isRefresh = true}) async {
    if (isLoadingMore) return;

    setState(() {
      if (isRefresh) {
        isLoadingStores = true;
        currentPage = 0;
        hasMoreStores = true;
      } else {
        isLoadingMore = true;
      }
    });

    try {
      final response = await ref.read(storesServiceProvider).getPublicStores(
            countryCode: storesFiltersCriteria.country,
            categoryCode: storesFiltersCriteria.category,
            sort: storesFiltersCriteria.sort,
            name: storesFiltersCriteria.name,
            page: currentPage,
            size: storesPerPage,
          );

      setState(() {
        if (isRefresh) {
          stores = List<PublicStoreModel>.from(response.data);
        } else {
          stores = [...stores, ...response.data];
        }

        hasMoreStores = response.data.length >= storesPerPage;

        if (!isRefresh && response.data.isNotEmpty) {
          currentPage++;
        }

        isLoadingStores = false;
        isLoadingMore = false;
      });
    } catch (e) {
      setState(() {
        isLoadingStores = false;
        isLoadingMore = false;
      });
      print('Error fetching stores: $e');
    }
  }

  void _onScroll() {
    if (_scrollController.position.pixels >=
        _scrollController.position.maxScrollExtent - 200) {
      if (!isLoadingMore && hasMoreStores) {
        fetchPublicStores(isRefresh: false);
      }
    }
  }

  void goToStoreDetails(String storeID) {
    context.pushNamed(RouteNames.storeDetailsRouteName, pathParameters: {'id': storeID}, extra: { 'previousRoute': "/client/${RouteNames.shopRouteName}" });
  }

  void _resetVisitStore() {
    Future.microtask(() {
      ref.read(visitStoreNotifierProvider).removeVisitStore();
    });
  }

  void _setCategories() async {
    ref.read(storesServiceProvider).getPublicCategories().then((categories) {
      setState(() {
        this.categories = categories.data.map((category) {
          return DropdownOption(label: category.name, value: category.code);
        }).toList();
      });
    });
  }

  @override
  void dispose() {
    debounce.cancel();
    nameController.dispose();
    _scrollController.dispose();
    super.dispose();
  }

  @override
  void initState() {
    super.initState();

    _resetVisitStore();

    _scrollController.addListener(_onScroll);

    _initializeLocation();

    _setCategories();
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(left: 20, right: 20, top: 100),
      child: CustomScrollView(
        controller: _scrollController,
        slivers: <Widget>[
          SliverAppBar(
            automaticallyImplyLeading: false,
            scrolledUnderElevation: 0,
            floating: true,
            snap: true,
            expandedHeight: 75,
            flexibleSpace: FlexibleSpaceBar(
              background: StoresFilters(
                countries: countries,
                onSelectCountry: (String? country) {
                  setState(() {
                    storesFiltersCriteria.country = country;
                  });

                  fetchPublicStores();
                },
                onClearCountry: () {
                  setState(() {
                    storesFiltersCriteria.country = null;
                  });

                  fetchPublicStores();
                },
                countryController: countryController,
                categories: categories,
                onSelectCategory: (String? category) {
                  setState(() {
                    storesFiltersCriteria.category = category;
                  });

                  fetchPublicStores();
                },
                onClearCategory: () {
                  setState(() {
                    storesFiltersCriteria.category = null;
                  });

                  fetchPublicStores();
                },
                categoryController: categoryController,
                nameController: nameController,
                onSearch: () {
                  setState(() => storesFiltersCriteria.name = nameController.text);
                  fetchPublicStores();
                },
                sortValue: storesFiltersCriteria.sort,
                onChangeSortDropdown: (value) {
                  setState(() {
                    storesFiltersCriteria.sort = value;
                  });

                  fetchPublicStores();
                }
              ),
              collapseMode: CollapseMode.pin,
            ),
          ),
          if (isLoadingStores)
            const SliverToBoxAdapter(
              child: Center(
                child: Padding(
                  padding: EdgeInsets.all(20),
                  child: CircularProgressIndicator(),
                ),
              ),
            ),
          if (!isLoadingStores && stores.isNotEmpty)
            SliverPadding(
              padding: const EdgeInsets.only(bottom: kBottomNavigationBarHeight + 24, top: 15),
              sliver: SliverList(
                delegate: SliverChildBuilderDelegate(
                  (context, rowIndex) {
                    final int itemIndex1 = rowIndex * 2;
                    final int itemIndex2 = itemIndex1 + 1;

                    final firstStore = stores[itemIndex1];
                    final secondStore =
                        itemIndex2 < stores.length ? stores[itemIndex2] : null;

                    return IntrinsicHeight(
                      child: Padding(
                        padding: const EdgeInsets.symmetric(vertical: 5),
                        child: Row(
                          children: [
                            Expanded(
                              child: StoreCard(store: firstStore, onClick: () => goToStoreDetails(firstStore.uuid)),
                            ),
                            const SizedBox(width: 10),
                            Expanded(
                              child: secondStore != null
                                  ? StoreCard(store: secondStore, onClick: () => goToStoreDetails(secondStore.uuid))
                                  : const SizedBox(),
                            ),
                          ],
                        ),
                      ),
                    );
                  },
                  childCount: (stores.length / 2).ceil(),
                ),
              ),
            ),
          if (isLoadingMore)
            const SliverToBoxAdapter(
              child: Center(
                child: Padding(
                  padding: EdgeInsets.symmetric(horizontal: 20, vertical: 20),
                  child: CircularProgressIndicator(),
                ),
              ),
            ),
          if (!isLoadingStores && stores.isEmpty)
            SliverToBoxAdapter(
              child: Center(
                child: Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 40),
                  child: storesFiltersCriteria.country == null
                      ? Column(
                          mainAxisSize: MainAxisSize.min,
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            SvgPicture.asset(
                              Assets.info,
                              colorFilter: const ColorFilter.mode(
                                  AppColors.radicalRed, BlendMode.srcIn),
                            ),
                            const SizedBox(height: 10),
                            Text(
                              translation(context).selectCountryToSeeAvailableStores,
                              style: Theme.of(context).textTheme.bodyMedium,
                              softWrap: true,
                              textAlign: TextAlign.center,
                            ),
                          ],
                        )
                      : (stores.isEmpty
                          ? Text(
                              translation(context).noStoresFound,
                              style: Theme.of(context).textTheme.bodyMedium,
                            )
                          : const SizedBox.shrink()),
                ),
              ),
            ),
          // ToDO talk to the designer to figure out where to put this
          /* const SliverPadding(
            padding: EdgeInsets.only(top: 40, bottom: kBottomNavigationBarHeight + 24),
            sliver: SliverToBoxAdapter(
              child: ProcessStepCard()
            )
          ), */
        ],
      ),
    );
  }
}
