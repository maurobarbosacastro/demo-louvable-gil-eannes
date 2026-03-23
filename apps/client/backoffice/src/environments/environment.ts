// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
	production: false,
	version: '00.00.00',
	company: 'Atlanse Portugal',
	keycloak: {
		host: 'http://localhost:8081/',
		realm: 'tagpeak',
		clientId: 'tagpeak-client',
		clientSecret: '5BXrDobSxJaMPzUqXKlVaWscI9wn5AfF',
		tokenUrl: '',
		authUrl: '',
		logoutUrl: '',
		userInfo: '',
	},
	host: 'http://localhost',
	tagpeak: ':8097',
	email: ':8093',
	firebase: ':8092',
	image: {
		host: ':8099',
		path: '/api/image/',
		origin: '',
	},
	features: {
		currency: true,
		authSocials: true,
		dashboard: true,
		bugReport: true,
		notifications: true,
	},
	imageSupportedFormats: '.jpg, .jpeg, .png, .webp, image/png, image/temp',
	shopify: {
		appName: 'diogo-tp',
	},
	firebaseConfig: {
		apiKey: 'AIzaSyBryVQZZ8MsnvGGhq96XMlgEafvQ7wuJ4U',
		authDomain: 'tagpeak-test.firebaseapp.com',
		projectId: 'tagpeak-test',
		storageBucket: 'tagpeak-test.firebasestorage.app',
		messagingSenderId: '509240225527',
		appId: '1:509240225527:web:e015a320db03bde4466df2',
		measurementId: 'G-ML325X3LFV',
		vapidKey:
			'BHQhLgnyXVVGH1oRtgyNUi7cjPm36OGbJ8POo8x76kiODUmuEmzdRxOVUccr9RMqjd_TAPFvUSZ9knkxj6A7zSw',
	},
	commercial: {
		base: 'https://tagpeak.com',
		faqUrl: 'https://tagpeak.com/faqs',
		about: 'https://tagpeak.com/about',
		contacts: 'https://tagpeak.com/contacts',
		forPartners: 'https://tagpeak.com/for-partners',
		forEmployees: 'https://tagpeak.com/for-employees',
		news: 'https://tagpeak.com/news',
		termsAndConditions: 'https://tagpeak.com/legal/terms-and-conditions',
		privacyPolicy: 'https://tagpeak.com/legal/privacy-policy',
		linkedIn: 'https://www.linkedin.com/company/tagpeak-technologies-ltd/',
		x: 'https://x.com/tagpeak',
		instagram: 'https://www.instagram.com/tagpeak.official/',
	},
};

environment.keycloak = {
	...environment.keycloak,
	tokenUrl: `${environment.keycloak.host}realms/${environment.keycloak.realm}/protocol/openid-connect/token`,
	authUrl: `${environment.keycloak.host}realms/${environment.keycloak.realm}/protocol/openid-connect/auth`,
	logoutUrl: `${environment.keycloak.host}realms/${environment.keycloak.realm}/protocol/openid-connect/logout`,
	userInfo: `${environment.keycloak.host}realms/${environment.keycloak.realm}/protocol/openid-connect/userinfo`,
};

environment.image = {
	...environment.image,
	origin: `${environment.image.host}/images/`,
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.
