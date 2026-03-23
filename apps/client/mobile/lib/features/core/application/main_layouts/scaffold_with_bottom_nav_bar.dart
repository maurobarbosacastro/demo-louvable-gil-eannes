import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/auth/models/user_model.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/core/application/main_layouts/header/basic_header.dart';
import 'package:tagpeak/features/core/application/main_layouts/navigation_bar/floating_navigation_bar.dart';
import 'package:tagpeak/features/core/application/main_layouts/navigation_bar/navigation_bar_tab.dart';
import 'package:tagpeak/features/core/application/main_layouts/navigation_bar/simple_navigation_bar.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/core/routes/route_names.dart';
import 'package:tagpeak/shared/models/user_stats_model/user_stats_model.dart';
import 'package:tagpeak/shared/providers/user_provider.dart';
import 'package:tagpeak/utils/custom_icons.dart';
import 'package:tagpeak/utils/enums/user_roles_enum.dart';

import 'header/back_navigation_header.dart';

class RouteDescriptor {
  final String path;
  final String title;

  RouteDescriptor(this.path, this.title);
}

List<RouteDescriptor> ROUTES_WITHOUT_HEADER = [
  RouteDescriptor('/store-details', 'Store details'),
  RouteDescriptor('/redirect', '')
];

class ScaffoldWithBottomNavBar extends ConsumerStatefulWidget {
  final Widget child;
  final String location;
  final String? previousPath;
  final String? previousRouteParent;

  const ScaffoldWithBottomNavBar({super.key, required this.child, required this.location, this.previousPath, this.previousRouteParent });

  @override
  ConsumerState<ScaffoldWithBottomNavBar> createState() => _ScaffoldWithBottomNavBarState();
}

class _ScaffoldWithBottomNavBarState extends ConsumerState<ScaffoldWithBottomNavBar> {
  UserStatsModel? stats;
  late UserModel loggedUser;
  bool isAdmin = false;

  @override
  void initState() {
    super.initState();
    loggedUser = ref.read(userProvider)!;

    isAdmin = loggedUser.roles.contains(UserRolesEnum.admin.name);
    _fetchUserStatsIfNeeded(loggedUser.uuid);
  }

  int get _currentIndex {
    final uri = GoRouterState.of(context).uri.toString();
    final idx = navigationItems.indexWhere((tab) => uri.startsWith(tab.initialLocation));
    return idx == -1 ? 0 : idx;
  }

  List<NavigationBarTab> get navigationItems {
    return isAdmin ? _fetchAdminNavigationTabs() : _fetchClientNavigationTabs();
  }

  List<NavigationBarTab> _fetchAdminNavigationTabs() {
    return [
      NavigationBarTab(icon: CustomIcons.dashboard, title: translation(context).dashboard, initialLocation: '${RouteNames.adminRouteLocation}${RouteNames.dashboardRouteLocation}'),
      NavigationBarTab(icon: CustomIcons.cashback, title: translation(context).cashback, initialLocation: '${RouteNames.adminRouteLocation}${RouteNames.cashbackRouteLocation}'),
      NavigationBarTab(icon: CustomIcons.store_mall_directory, title: translation(context).stores, initialLocation: '${RouteNames.adminRouteLocation}${RouteNames.storesRouteLocation}'),
      NavigationBarTab(icon: CustomIcons.layout, title: translation(context).more, initialLocation: '${RouteNames.adminRouteLocation}${RouteNames.moreRouteLocation}'),
    ];
  }

  List<NavigationBarTab> _fetchClientNavigationTabs() {
    return [
      NavigationBarTab(icon: CustomIcons.dashboard, title: translation(context).dashboard, initialLocation: '${RouteNames.clientRouteLocation}${RouteNames.dashboardRouteLocation}'),
      NavigationBarTab(icon: CustomIcons.shopping_cart, title: translation(context).shop, initialLocation: '${RouteNames.clientRouteLocation}${RouteNames.shopRouteLocation}'),
      NavigationBarTab(icon: CustomIcons.money, title: translation(context).referEarn, initialLocation: '${RouteNames.clientRouteLocation}${RouteNames.referEarnRouteLocation}'),
    ];
  }

  void _onItemTapped(int index) {
    final location = navigationItems[index].initialLocation;
    if (location != 'not-ready') {
      context.push(location);
    }
  }

  void _fetchUserStatsIfNeeded(String uuid) {
    ref.read(userNotifier).getUserStats(uuid).then((value) {
      if (mounted) {
        setState(() {
          stats = value;
        });
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    final RouteDescriptor route = ROUTES_WITHOUT_HEADER.firstWhere(
          (route) => widget.location.contains(route.path),
      orElse: () => RouteDescriptor('', ''),
    );

    final bool hasBackNavigation = route.path.isNotEmpty;

    return Container(
      color: Colors.white,
      child: SafeArea(
        child: Scaffold(
          extendBodyBehindAppBar: true,
          resizeToAvoidBottomInset: false,
          extendBody: true,
          appBar: !hasBackNavigation ? BasicHeader(
              profilePicture: loggedUser.profilePicture,
              isAdmin: isAdmin,
              stats: stats,
          ) : BackNavigationHeader(title: route.title, backPathLocation: widget.previousPath, backPathParentLocation: widget.previousRouteParent) as PreferredSizeWidget,
          body: GestureDetector(
            onTap: () {
              FocusScope.of(context).requestFocus(FocusNode());
            },
            child: Padding(
              padding: const EdgeInsets.only(bottom: kBottomNavigationBarHeight + 10),
              child: widget.child,
            ),
          ),
          bottomNavigationBar: isAdmin
              ? SimpleNavigationBar(
            currentIndex: _currentIndex,
            onItemSelected: _onItemTapped,
            tabs: navigationItems,
          )
              : FloatingNavigationBar(
            currentIndex: _currentIndex,
            onItemSelected: _onItemTapped,
            tabs: navigationItems,
          ),
        ),
      ),
    );
  }
}
