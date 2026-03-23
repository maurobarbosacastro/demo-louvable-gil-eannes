export const environment = {
	production: true,
	version: '%version%',
	company: 'Atlanse Portugal',
	keycloak: {
		host: 'https://app.tagpeak.com/keycloak/auth',
		realm: 'tagpeak',
		clientId: 'tagpeak-client',
		clientSecret: 'ZHWJK7kbK9gxP2XAUHb37rwD3jp75A3M',
		tokenUrl: '',
		authUrl: '',
		logoutUrl: '',
		userInfo: '',
	},
	host: 'https://app.tagpeak.com',
	email: '/email',
	tagpeak: '/tagpeak',
	image: {
		host: '/images',
		path: '/api/image/',
		origin: '',
	},
	features: {
		currency: true,
		authSocials: true,
		dashboard: true,
		bugReport: true,
		notifications: false,
	},
	imageSupportedFormats: '.jpg, .jpeg, .png, .webp, image/png, image/temp',
	minimumAmountWithdraw: 20,
	shopify: {
		appName: '',
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
	tokenUrl: `${environment.keycloak.host}/realms/${environment.keycloak.realm}/protocol/openid-connect/token`,
	authUrl: `${environment.keycloak.host}/realms/${environment.keycloak.realm}/protocol/openid-connect/auth`,
	logoutUrl: `${environment.keycloak.host}/realms/${environment.keycloak.realm}/protocol/openid-connect/logout`,
	userInfo: `${environment.keycloak.host}/realms/${environment.keycloak.realm}/protocol/openid-connect/userinfo`,
};

environment.image = {
	...environment.image,
	origin: `${environment.image.host}/images/`,
};
