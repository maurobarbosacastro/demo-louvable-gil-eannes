import {provideHttpClient} from '@angular/common/http';
import {APP_INITIALIZER, ApplicationConfig, inject} from '@angular/core';
import {LuxonDateAdapter} from '@angular/material-luxon-adapter';
import {DateAdapter, MAT_DATE_FORMATS} from '@angular/material/core';
import {provideAnimations} from '@angular/platform-browser/animations';
import {
	PreloadAllModules,
	provideRouter,
	withComponentInputBinding,
	withHashLocation,
	withInMemoryScrolling,
	withPreloading,
} from '@angular/router';
import {provideFuse} from '@fuse';
import {TranslocoService, provideTransloco} from '@ngneat/transloco';
import {firstValueFrom} from 'rxjs';
import {TranslocoHttpLoader} from './core/transloco/transloco.http-loader';
import {appRoutes} from './app.routing';
import {mockApiServices} from '@app/mock-api';
import {provideAuth} from '@app/core/auth/auth.provider';
import {provideIcons} from '@app/core/icons/icons.provider';
import {provideQuillConfig} from 'ngx-quill';
import {provideFirebaseApp, initializeApp} from '@angular/fire/app';
import {getMessaging, provideMessaging} from '@angular/fire/messaging';
import {environment} from '@environments/environment';

const notificationsProviders = environment.features.notifications
	? [
			provideFirebaseApp(() => initializeApp(environment.firebaseConfig)),
			provideMessaging(() => getMessaging()),
		]
	: [];

export const appConfig: ApplicationConfig = {
	providers: [
		provideAnimations(),
		provideHttpClient(),
		provideRouter(
			appRoutes,
			withPreloading(PreloadAllModules),
			withInMemoryScrolling({scrollPositionRestoration: 'enabled'}),
			withComponentInputBinding(),
			withHashLocation(),
		),
		...notificationsProviders,

		// Material Date Adapter
		{
			provide: DateAdapter,
			useClass: LuxonDateAdapter,
		},
		{
			provide: MAT_DATE_FORMATS,
			useValue: {
				parse: {
					dateInput: 'dd/MM/yyyy',
				},
				display: {
					dateInput: 'dd/MM/yyyy',
					monthYearLabel: 'MMM yyyy',
					dateA11yLabel: 'DD/MM/YYYY',
					monthYearA11yLabel: 'MMMM yyyy',
				},
			},
		},

		// Transloco Config
		provideTransloco({
			config: {
				availableLangs: [
					{
						id: 'en',
						label: 'English',
					},
					{
						id: 'tr',
						label: 'Turkish',
					},
				],
				defaultLang: 'en',
				fallbackLang: 'en',
				reRenderOnLangChange: true,
				prodMode: true,
			},
			loader: TranslocoHttpLoader,
		}),
		{
			// Preload the default language before the app starts to prevent empty/jumping content
			provide: APP_INITIALIZER,
			useFactory: () => {
				const translocoService = inject(TranslocoService);
				const defaultLang = translocoService.getDefaultLang();
				translocoService.setActiveLang(defaultLang);

				return () => firstValueFrom(translocoService.load(defaultLang));
			},
			multi: true,
		},

		// Fuse
		provideAuth(),
		provideIcons(),
		provideFuse({
			mockApi: {
				delay: 0,
				services: mockApiServices,
			},
			fuse: {
				layout: 'classy',
				scheme: 'light',
				screens: {
					sm: '600px',
					md: '960px',
					lg: '1280px',
					xl: '1440px',
				},
				theme: 'theme-default',
				themes: [
					{
						id: 'theme-default',
						name: 'Default',
					},
					{
						id: 'theme-brand',
						name: 'Brand',
					},
					{
						id: 'theme-teal',
						name: 'Teal',
					},
					{
						id: 'theme-rose',
						name: 'Rose',
					},
					{
						id: 'theme-purple',
						name: 'Purple',
					},
					{
						id: 'theme-amber',
						name: 'Amber',
					},
				],
			},
		}),

		// Config fonts to Quill editor
		provideQuillConfig({
			customOptions: [
				{
					import: 'formats/font',
					whitelist: ['IBM Plex Sans'],
				},
			],
		}),
	],
};
