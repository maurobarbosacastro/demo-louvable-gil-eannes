export const environment = {
    production: true,
    version: '%version%',
    company: 'Atlanse Portugal',
    keycloak: {
        host: 'https://qa.keycloak.atlanse.ddns.net',
        realm: 'qa-tagpeak',
        clientId: 'tagpeak-client',
        clientSecret: 'IkL2av5teY2amuMldLgZoawD9eQJo8OV',
        tokenUrl: '',
        authUrl: '',
        logoutUrl: '',
        userInfo: '',
    },
    host: 'https://qa.tagpeak.atlanse.ddns.net',
    email: '/email',
    tagpeak: '/tagpeak',
    firebase: '/messaging',
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
        notifications: true
    },
    imageSupportedFormats: '.jpg, .jpeg, .png, .webp, image/png, image/temp',
    minimumAmountWithdraw: 20,
    shopify: {
        appName: 'tp-qa',
    },
    firebaseConfig: {
        apiKey: "AIzaSyBryVQZZ8MsnvGGhq96XMlgEafvQ7wuJ4U",
        authDomain: "tagpeak-test.firebaseapp.com",
        projectId: "tagpeak-test",
        storageBucket: "tagpeak-test.firebasestorage.app",
        messagingSenderId: "509240225527",
        appId: "1:509240225527:web:e015a320db03bde4466df2",
        measurementId: "G-ML325X3LFV",
        vapidKey: "BHQhLgnyXVVGH1oRtgyNUi7cjPm36OGbJ8POo8x76kiODUmuEmzdRxOVUccr9RMqjd_TAPFvUSZ9knkxj6A7zSw"
    },
	commercial:{
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
		instagram: 'https://www.instagram.com/tagpeak.official/'
	}
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
