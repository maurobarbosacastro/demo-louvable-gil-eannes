import { ActivatedRouteSnapshot, createUrlTreeFromSnapshot, Router, Routes } from '@angular/router';
import { AuthGuard } from '@app/core/auth/guards/auth.guard';
import { NoAuthGuard } from '@app/core/auth/guards/noAuth.guard';
import { LayoutComponent } from '@app/layout/layout.component';
import { initialDataResolver, manageStoredUtmParams } from '@app/app.resolvers';
import { CallbackComponent } from '@app/modules/auth/callback/callback.component';
import { AuthService } from '@app/core/auth/auth.service';
import { inject } from '@angular/core';
import { AppConstants, UserTypesEnum } from '@app/app.constants';
import { clientGuard, rolesGuard, shopGuard } from '@app/core/auth/guards/roles.guard';
import { StorageService } from '@app/shared/services/storage.service';
import { PublicStoreService } from '@app/modules/stores/services/public-store.service';
import { ConfigurationService } from '@app/modules/admin/configuration/services/configuration.service';
import { catchError, map, of, tap } from 'rxjs';
import { CountriesService } from '@app/modules/admin/countries/services/countries.service';
import { CategoriesService } from '@app/modules/admin/categories/services/categories.service';
import { ShopService } from '@app/modules/shop/services/shop.service';
import { UserService } from '@app/core/user/user.service';
import { NotificationService } from './modules/notification/notifications.service';
import {NoAuthShopGuard} from '@app/core/auth/guards/noAuth.shop.guard';

// @formatter:off
/* eslint-disable max-len */
/* eslint-disable @typescript-eslint/explicit-function-return-type */
export const appRoutes: Routes = [
    {
        path: '',
        pathMatch: 'full',
        redirectTo: (_) => {
            const authService = inject(AuthService);
            const userService = inject(UserService);

            if (!userService.$isProfileCompleted()) {
                return 'profile-setup';
            }

            if (authService.userGroups && authService.userGroups.includes(UserTypesEnum.admin)) {
                return 'admin';
            }

            if (authService.userGroups && authService.userGroups.includes(UserTypesEnum.shop)) {
                return 'shop';
            }

            return 'client';
        },
    },

    {
        path: 'auth',
        component: LayoutComponent,
        data: { layout: 'empty' },
        children: [
            { path: '', pathMatch: 'full', redirectTo: 'callback' },
            { path: 'callback', component: CallbackComponent },
            {
                path: 'utm',
                redirectTo: (redirectData: any) => {
                    const activatedRoute = redirectData as ActivatedRouteSnapshot;
                    const storageService = inject(StorageService);
                    const utmParams = Object.fromEntries(
                        Object.entries(activatedRoute.queryParams).filter(([key]) => key.startsWith('utm_')),
                    );
                    storageService.save(AppConstants.STORAGE_KEYS.UTM, JSON.stringify(utmParams));
                    return '/client';
                },
            },
        ],
    },

    // Redirect signed in user to the '/example'
    //
    // After the user signs in, the sign in page will redirect the user to the 'signed-in-redirect'
    // path. Below is another redirection for that path to redirect the user to the desired
    // location. This is a small convenience to keep all main routes together here on this file.
    {
        path: 'signed-in-redirect',
        pathMatch: 'full',
        redirectTo: (_) => {
            const authService = inject(AuthService);
            return authService.userGroups.includes(UserTypesEnum.admin) ? '/admin' : '/client';
        },
    },

    // Auth routes for guests
    {
        path: '',
        canActivate: [NoAuthGuard],
        canActivateChild: [NoAuthGuard],
        component: LayoutComponent,
        data: {
            layout: 'empty',
        },
        children: [
            {
                path: 'confirmation-required',
                loadComponent: () =>
                    import('@app/modules/auth/confirmation-required/confirmation-required.component').then(
                        (m) => m.AuthConfirmationRequiredComponent,
                    ),
            },
            {
                path: 'forgot-password',
                loadComponent: () =>
                    import('@app/modules/auth/forgot-password/forgot-password.component').then(
                        (m) => m.AuthForgotPasswordComponent,
                    ),
            },
            {
                path: 'reset-password',
                loadComponent: () =>
                    import('@app/modules/auth/reset-password/reset-password.component').then(
                        (m) => m.AuthResetPasswordComponent,
                    ),
            },
            {
                path: 'sign-in',
                loadComponent: () =>
                    import('@app/modules/auth/sign-in/sign-in.component').then((m) => m.AuthSignInComponent),
            },
            {
                path: 'sign-up',
                loadComponent: () =>
                    import('@app/modules/auth/sign-up/sign-up.component').then((m) => m.AuthSignUpComponent),
            },
            {
                path: 'sign-up/:referralCode',
                loadComponent: () =>
                    import('@app/modules/auth/sign-up/sign-up.component').then((m) => m.AuthSignUpComponent),
            }
        ],
    },

    {
        path: 'profile-setup',
        canActivate: [
            AuthGuard,
            (_: ActivatedRouteSnapshot) => {
                const userService = inject(UserService);
                const router = inject(Router);

                if (!userService.$isProfileCompleted()) {
                    return true;
                }

                router.navigate(['/']);
                return false;
            }
        ],
        loadComponent: () =>
            import('@app/modules/auth/auth-profile-completion/auth-profile-completion.component').then(m => m.AuthProfileCompletionComponent),
    },

    {
        path: 'shopify',
        component: LayoutComponent,
        data: { layout: 'empty' },
        children: [
            {
                path: 'shop-sign-up',
                loadComponent: () =>
                    import('@app/modules/auth/shop-sign-up/shop-sign-up.component').then((m) => m.AuthShopSignUpComponent),
            },
            {
                path: 'shop-sign-in',
                loadComponent: () =>
                    import('@app/modules/auth/shop-sign-in/shop-sign-in.component').then((m) => m.AuthShopSignInComponent),
            },
        ]
    },

    {
        path: 'shop',
        canActivate: [AuthGuard, shopGuard],
        canActivateChild: [AuthGuard, shopGuard],
        component: LayoutComponent,
        data: {
            layout: 'thin',
        },
        resolve: {
            initialData: initialDataResolver,
            manageStoredUtmParams: manageStoredUtmParams,
            configs: (_) => {
                const configService = inject(ConfigurationService);
                configService.loadConfigurations();
                return of(true);
            },
            currentShop: () => {
                const shopService = inject(ShopService);
                const userService = inject(UserService);

                return shopService.getShopByUserUuid(userService.$user().uuid)
                    .pipe(
                        tap((data) => {
                            shopService.$currentShop.set(data);
                        })
                    )
            }
        },
        children: [
            { path: '', pathMatch: 'full', redirectTo: 'home' },
            {
                path: 'home',
                loadComponent: () =>
                    import('@app/modules/shop/pages/shop-home/shop-home.component').then(
                        (m) => m.ShopHomeComponent,
                    ),
            },
        ]
    },

    // Auth routes for authenticated users
    {
        path: '',
        canActivate: [AuthGuard],
        canActivateChild: [AuthGuard],
        component: LayoutComponent,
        data: {
            layout: 'empty',
        },
        resolve: {
            config: (_) => {
                const config = inject(ConfigurationService);
                config.loadConfigurations();
                return of(true);
            },
        },
        children: [
            {
                path: 'sign-out',
                loadComponent: () =>
                    import('@app/modules/auth/sign-out/sign-out.component').then(
                        (m) => m.AuthSignOutComponent,
                    ),
            },
            {
                path: 'unlock-session',
                loadComponent: () =>
                    import('@app/modules/auth/unlock-session/unlock-session.component').then(
                        (m) => m.AuthUnlockSessionComponent,
                    ),
            },
        ],
    },

    // Public routes
    {
        path: '',
        component: LayoutComponent,
        data: {
            layout: 'empty',
        },
        children: [
            // Landing route
            {
                path: 'home',
                loadComponent: () =>
                    import('@app/modules/landing/home/home.component').then((m) => m.LandingHomeComponent),
            },
            {
                path: 'stores',
                children: [
                    {
                        path: '',
                        resolve: {
                            $preLoadedCountry: (route: ActivatedRouteSnapshot) => {
                                const countryService = inject(CountriesService);
                                const code = route.queryParamMap.get('country');

                                if (!code) {
                                    return of(null);
                                }

                                return countryService.getCountryByCode(code).pipe(
                                    catchError((_) => {
                                        return of(null);
                                    }),
                                );
                            },
                            $preLoadedCategory: (route: ActivatedRouteSnapshot) => {
                                const categoriesService = inject(CategoriesService);
                                const code = route.queryParamMap.get('category');

                                if (!code) {
                                    return of(null);
                                }

                                return categoriesService.getCategoryByCode(code).pipe(
                                    catchError((_) => {
                                        return of(null);
                                    }),
                                );
                            },
                            $preLoadedSearchText: (route: ActivatedRouteSnapshot) => {
                                const searchText = route.queryParamMap.get('name');

                                if (!searchText) {
                                    return of(null);
                                }

                                return of(searchText);
                            },
                        },
                        loadComponent: () =>
                            import('@app/modules/stores/stores.component').then((m) => m.StoresComponent),
                    },
                    {
                        path: ':id',
                        children: [
                            {
                                path: '',
                                resolve: {
                                    $store: (route: ActivatedRouteSnapshot) => {
                                        const publicStoreService = inject(PublicStoreService);
                                        return publicStoreService.getStorePublic(route.params.id);
                                    },
                                },
                                loadComponent: () =>
                                    import('@app/modules/stores/stores-detail/pages/stores-detail.component').then(
                                        (m) => m.StoresDetailComponent,
                                    ),
                            },
                            {
                                path: 'redirect',
                                canActivate: [
                                    (route: ActivatedRouteSnapshot) => {
                                        if (window.history.length > 2) {
                                            return createUrlTreeFromSnapshot(route, ['/', 'stores', route.params.id]);
                                        }
                                        return true;
                                    },
                                ],
                                resolve: {
                                    $store: (route: ActivatedRouteSnapshot) => {
                                        const publicStoreService = inject(PublicStoreService);
                                        return publicStoreService.getStorePublic(route.params.id);
                                    },
                                },
                                loadComponent: () =>
                                    import('@app/modules/stores/pages/redirect/redirect.component').then(
                                        (m) => m.RedirectComponent,
                                    ),
                            },
                        ],
                    },
                ],
            },
        ],
    },

    // Admin routes
    {
        path: 'admin',
        canActivate: [AuthGuard, rolesGuard],
        canActivateChild: [AuthGuard, rolesGuard],
        component: LayoutComponent,
        data: {
            layout: 'thin',
        },
        resolve: {
            initialData: initialDataResolver,
            manageStoredUtmParams: manageStoredUtmParams,
            configs: (_) => {
                const configService = inject(ConfigurationService);
                configService.loadConfigurations();
                return of(true);
            },
        },
        children: [
            { path: '', pathMatch: 'full', redirectTo: 'dashboard' },
            {
                path: 'dashboard',
                loadComponent: () =>
                    import('@app/modules/admin/dashboard/pages/dashboard.component').then(
                        (m) => m.DashboardComponent,
                    ),
            },
            {
                path: 'cashbacks',
                loadComponent: () =>
                    import('@app/modules/admin/cashback/components/cashback-header/cashback.component').then(
                        (m) => m.CashbackComponent,
                    ),
            },
            {
                path: 'email',
                children: [
                    {
                        path: '',
                        loadComponent: () =>
                            import('@app/modules/admin/email/listing/listing.component').then(
                                (m) => m.ListingComponent,
                            ),
                    },
                    {
                        path: ':id',
                        loadComponent: () =>
                            import('@app/modules/admin/email/email-generator/email-generator.component').then(
                                (m) => m.EmailGeneratorComponent,
                            ),
                    },
                    {
                        path: 'new',
                        loadComponent: () =>
                            import('@app/modules/admin/email/email-generator/email-generator.component').then(
                                (m) => m.EmailGeneratorComponent,
                            ),
                    },
                ],
            },
            {
                path: 'countries',
                loadComponent: () =>
                    import('@app/modules/admin/countries/pages/countries-home/countries.component').then(
                        (m) => m.CountriesComponent,
                    ),
            },
            {
                path: 'sources',
                loadComponent: () =>
                    import('@app/modules/admin/partners/pages/partners-home/partners.component').then(
                        (m) => m.PartnersComponent,
                    ),
            },
            {
                path: 'languages',
                loadComponent: () =>
                    import(
                        '@app/modules/admin/languages/components/languages-management/languages-management.component'
                    ).then((m) => m.LanguageManagementComponent),
            },
            {
                path: 'categories',
                loadComponent: () =>
                    import(
                        '@app/modules/admin/categories/components/categories-management/categories-management.component'
                    ).then((m) => m.CategoriesManagementComponent),
            },
            {
                path: 'stores',
                loadComponent: () =>
                    import('@app/modules/admin/stores/pages/home-stores/home-stores.component').then((m) => m.HomeStoresComponent),
            },
            {
                path: 'store-visits',
                loadComponent: () =>
                    import('@app/modules/admin/store-visits/pages/store-visits.component').then(
                        (m) => m.StoreVisitsComponent,
                    ),
            },
            {
                path: 'configurations',
                loadComponent: () =>
                    import(
                        '@app/modules/admin/configuration/pages/configurations/configurations.component'
                    ).then((m) => m.ConfigurationsComponent),
            },
        ],
    },

    // Client routes
    {
        path: 'client',
        canActivate: [AuthGuard, clientGuard],
        canActivateChild: [AuthGuard, clientGuard],
        component: LayoutComponent,
        data: {
            layout: 'centered',
        },
        resolve: {
            initialData: initialDataResolver,
            manageStoredUtmParams: manageStoredUtmParams,
            configs: (_) => {
                const configService = inject(ConfigurationService);
                configService.loadConfigurations();
                return of(true);
            },
            setupNotifications: (_) => {
                const notifService = inject(NotificationService)
                notifService.requestPermission()

                return of(true)
            }
        },
        children: [
	        {
		        path: '',
		        pathMatch: 'full',
		        redirectTo: (_) => {
			        const userService = inject(UserService);

			        if (!userService.$isProfileCompleted()) {
				        return '/profile-setup';
			        }

			        return 'dashboard';
		        },
	        },
            {
                path: 'settings',
                resolve: {
                    $preLoadedCountry: () => {
                        const countryService = inject(CountriesService);

                        return countryService.getCountries(0, 0, 'abbreviation asc').pipe(
                            catchError((_) => {
                                return of(null);
                            }),
                            map((value) => {
                                return value.data;
                            }),
                        );
                    },
                    configs: (_) => {
                        const configService = inject(ConfigurationService);
                        configService.loadConfigurations();
                        return of(true);
                    },
                },
                loadComponent: () =>
                    import('@app/modules/settings/settings.component').then((m) => m.SettingsComponent),
            },
            {
                path: 'dashboard',
                resolve: {
                    configs: (_) => {
                        const configService = inject(ConfigurationService);
                        configService.loadConfigurations();
                        return of(true);
                    },
                },
                loadComponent: () =>
                    import('@app/modules/client/dashboard/pages/dashboard/dashboard.component').then(
                        (m) => m.DashboardComponent,
                    ),
            },
            {
                path: 'referrals',
                resolve: {
                    configs: (_) => {
                        const configService = inject(ConfigurationService);
                        configService.loadConfigurations();
                        return of(true);
                    },
                },
                loadComponent: () =>
                    import('@app/modules/client/referrals/pages/referrals.component').then(
                        (m) => m.ReferralsComponent,
                    ),
            },
        ],
    },
];
