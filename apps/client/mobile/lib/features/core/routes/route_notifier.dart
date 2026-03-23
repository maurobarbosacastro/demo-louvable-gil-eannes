import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:tagpeak/features/admin/more_features/more_features.dart';
import 'package:tagpeak/features/auth/provider/user_info_provider.dart';
import 'package:tagpeak/features/auth/screens/login/provider/login_provider.dart';
import 'package:tagpeak/features/auth/screens/register/pages/confirm_email_screen.dart';
import 'package:tagpeak/features/core/application/main_layouts/scaffold_without_bottom_nav_bar.dart';
import 'package:tagpeak/features/dashboard/dashboard.dart';
import 'package:tagpeak/features/refer_earn/refer_earn.dart';
import 'package:tagpeak/features/reward_details/reward_details.dart';
import 'package:tagpeak/features/stores/stores/store_details/store_details.dart';
import 'package:tagpeak/features/stores/stores/store_redirect/store_redirect.dart';
import 'package:tagpeak/features/stores/stores/stores.dart';
import 'package:tagpeak/utils/enums/user_roles_enum.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';
import 'package:tagpeak/features/admin/dashboard/dashboard.dart' as admin_dashboard;
import '../../auth/screens/login/pages/login_screen.dart';
import '../../auth/screens/login/pages/reset_password_screen.dart';
import '../../auth/screens/register/pages/register_screen.dart';
import '../../settings/pages/settings_page.dart';
import '../../splashscreen/splash_screen.dart';
import '../application/main_layouts/scaffold_with_bottom_nav_bar.dart';
import '../application/core_provider.dart';
import 'route_names.dart';

CustomTransitionPage<void> buildPageWithDefaultTransition<T>({
  required BuildContext context,
  required GoRouterState state,
  required Widget child
}) {
  return CustomTransitionPage<void>(
    key: state.pageKey,
    child: child,
    transitionDuration: const Duration(milliseconds: 0),
    transitionsBuilder: (BuildContext context, Animation<double> animation,
        Animation<double> secondaryAnimation, Widget child) {
      return FadeTransition(
        opacity: CurveTween(curve: Curves.easeInBack).animate(animation),
        child: child,
      );
    },
  );
}

final rootNavigatorKey = GlobalKey<NavigatorState>();
final shellNavigatorKey = GlobalKey<NavigatorState>();

final routerProvider = Provider<GoRouter>(
      (ref) {
    return GoRouter(
      navigatorKey: rootNavigatorKey,
      debugLogDiagnostics: true,
      initialLocation: ref.watch(lifecycleRouteProvider),
      routes: [
        GoRoute(
          name: RouteNames.splashRouteName,
          path: RouteNames.splashRouteLocation,
          builder: (context, state) => const SplashScreen(),
        ),
        GoRoute(
          name: RouteNames.loginRouteName,
          path: RouteNames.loginRouteLocation,
          builder: (context, state) =>
          const ScaffoldWithoutBottomNavBar(child: LoginScreen()),
        ),
        GoRoute(
          name: RouteNames.registerRouteName,
          path: RouteNames.registerRouteLocation,
          builder: (context, state) => const ScaffoldWithoutBottomNavBar(
              extendBodyBehindAppBar: true, child: RegisterScreen()),
        ),
        GoRoute(
          name: RouteNames.confirmEmailRouteName,
          path: RouteNames.confirmEmailRouteLocation,
          builder: (context, state) {
            final bool isPasswordReset =
                state.uri.queryParameters['isPasswordReset'] == 'true';

            return ScaffoldWithoutBottomNavBar(
              extendBodyBehindAppBar: true,
              trailingHeaderButton: !isPasswordReset
                  ? TrailingHeaderButton.whiteMode
                  : TrailingHeaderButton.darkMode,
              child: ConfirmEmailScreen(
                isPasswordReset: isPasswordReset,
              ),
            );
          },
        ),
        GoRoute(
          name: RouteNames.resetPasswordRouteName,
          path: RouteNames.resetPasswordRouteLocation,
          builder: (context, state) => ScaffoldWithoutBottomNavBar(
            trailingHeaderButton: TrailingHeaderButton.darkMode,
            child: const ResetPasswordScreen(),
            onPressedTrailingButton: () {
              ref
                  .read(resetPasswordErrorProvider.notifier)
                  .update((state) => FormErrorModel(isError: false));
            },
          ),
        ),
        ShellRoute(
          navigatorKey: shellNavigatorKey,
          builder: (context, state, child) {
            final myData = state.extra as Map<String, dynamic>?;

            return ScaffoldWithBottomNavBar(
                location: state.uri.toString(),
                previousPath: myData?['previousRoute'],
                previousRouteParent: myData?['previousRouteParent'],
                child: child,
            );
          },
          routes: [
            // Client side
            GoRoute(
              name: RouteNames.clientRouteName,
              path: RouteNames.clientRouteLocation,
              redirect: (context, state) {
                final loggedUser = ref.read(userProvider);
                if (loggedUser == null) return '/';

                final userRole = loggedUser.roles;

                if (userRole.contains('default')) {
                  if (state.uri.path == RouteNames.clientRouteLocation) {
                    return '${RouteNames.clientRouteLocation}${RouteNames.dashboardRouteLocation}';
                  }
                  return null;
                } else {
                  return '/';
                }
              },
              routes: [
                GoRoute(
                  name: RouteNames.clientDashboardRouteName,
                  path: RouteNames.dashboardRouteLocation,
                  pageBuilder: (context, state) => buildPageWithDefaultTransition<void>(
                    child: const Dashboard(),
                    context: context,
                    state: state,
                  ),
                  routes: [
                    GoRoute(
                      name: RouteNames.rewardRouteName,
                      path: '${RouteNames.rewardRouteLocation}/:id',
                      pageBuilder: (context, state) {
                        final String id = state.pathParameters['id']!;
                        return buildPageWithDefaultTransition<void>(
                          child: RewardDetails(id: id),
                          context: context,
                          state: state,
                        );
                      },
                    ),
                  ]
                ),
                GoRoute(
                  name: RouteNames.shopRouteName,
                  path: RouteNames.shopRouteLocation,
                  pageBuilder: (context, state) =>
                      buildPageWithDefaultTransition<void>(
                    child: const Stores(),
                    context: context,
                    state: state,
                  ),
                ),
                GoRoute(
                  name: RouteNames.referEarnRouteName,
                  path: RouteNames.referEarnRouteLocation,
                  pageBuilder: (context, state) => buildPageWithDefaultTransition<void>(
                    child: const ReferEarn(),
                    context: context,
                    state: state,
                  ),
                ),
                GoRoute(
                  name: RouteNames.settingsRouteName,
                  path: RouteNames.settingsRouteLocation,
                  pageBuilder: (context, state) {
                    final tab = state.uri.queryParameters['tab'];
                    return buildPageWithDefaultTransition<void>(
                      child: SettingsPage(initialTab: tab),
                      context: context,
                      state: state,
                    );
                  },
                ),
                GoRoute(
                    name: RouteNames.shopRouteName,
                    path: RouteNames.shopRouteLocation,
                    pageBuilder: (context, state) =>
                        buildPageWithDefaultTransition<void>(
                          child: const Stores(),
                          context: context,
                          state: state,
                        ),
                    routes: [
                      GoRoute(
                          name: RouteNames.storeDetailsRouteName,
                          path: '${RouteNames.storeDetailsRouteLocation}/:id',
                          pageBuilder: (context, state) {
                            final String storeID = state.pathParameters['id']!;
                            return buildPageWithDefaultTransition<void>(
                              child: StoreDetails(storeID: storeID),
                              context: context,
                              state: state
                            );
                          }
                      ),
                      GoRoute(
                          name: RouteNames.redirectDetailsRouteName,
                          path: RouteNames.redirectRouteLocation,
                          pageBuilder: (context, state) =>
                             buildPageWithDefaultTransition<void>(
                              child: const StoreRedirect(),
                              context: context,
                              state: state,
                            )
                      )
                    ]
                ),
              ],
            ),
            // Admin side
            GoRoute(
              name: RouteNames.adminRouteName,
              path: RouteNames.adminRouteLocation,
              redirect: (context, state) {
                final loggedUser = ref.read(userProvider);
                if (loggedUser == null) return '/';

                final userRole = loggedUser.roles;

                if (userRole.contains(UserRolesEnum.admin.name)) {
                  if (state.uri.path == RouteNames.adminRouteLocation) {
                    return '${RouteNames.adminRouteLocation}${RouteNames.dashboardRouteLocation}';
                  }
                  return null;
                } else {
                  return '/';
                }
              },
              routes: [
                GoRoute(
                  name: RouteNames.adminDashboardRouteName,
                  path: RouteNames.dashboardRouteLocation,
                  pageBuilder: (context, state) => buildPageWithDefaultTransition<void>(
                    child: const admin_dashboard.Dashboard(),
                    context: context,
                    state: state,
                  ),
                ),
                GoRoute(
                  name: RouteNames.cashbackRouteName,
                  path: RouteNames.cashbackRouteLocation,
                  pageBuilder: (context, state) => buildPageWithDefaultTransition<void>(
                    child: const Center(child: Text('cashback admin')),
                    context: context,
                    state: state,
                  ),
                ),
                GoRoute(
                  name: RouteNames.storesRouteName,
                  path: RouteNames.storesRouteLocation,
                  pageBuilder: (context, state) => buildPageWithDefaultTransition<void>(
                    child: const Center(child: Text('stores admin')),
                    context: context,
                    state: state,
                  ),
                ),
                GoRoute(
                  name: RouteNames.moreRouteName,
                  path: RouteNames.moreRouteLocation,
                  pageBuilder: (context, state) => buildPageWithDefaultTransition<void>(
                    child: MoreFeatures(),
                    context: context,
                    state: state,
                  ),
                ),
              ],
            ),
          ],
        ),
      ],
    );
  },
);
