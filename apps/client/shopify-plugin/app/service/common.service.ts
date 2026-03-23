import type {StoreInterface} from '../interfaces/common.interface';
import type {ServiceResponse} from '../interfaces/dashbaord.interface';
import type {StoreUpdateDTO} from '../interfaces/dtos.interface';

export interface UserInterface {
	uuid: string;
	email: string;
	isVerified: boolean;
	firstName: string;
	lastName: string;
}

export async function getStoreDetails(storeUuid: string): Promise<ServiceResponse<StoreInterface>> {
	//@ts-ignore
	const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/tagpeak/store/${storeUuid}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
	});

	if (!call.ok) {
		if (call.status === 401) {
			return {
				status: 401,
				message: 'Unauthorized',
			};
		}
		return {
			status: call.status,
			message: 'Internal Server Error',
		};
	}

	const callJson: StoreInterface = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}

export async function updateStoreDetails(
	storeUuid: string,
	data: StoreUpdateDTO,
): Promise<ServiceResponse<StoreInterface>> {
	//@ts-ignore
	const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/tagpeak/store/${storeUuid}`, {
		method: 'PATCH',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
		body: JSON.stringify(data),
	});

	if (!call.ok) {
		if (call.status === 401) {
			return {
				status: 401,
				message: 'Unauthorized',
			};
		}
		return {
			status: call.status,
			message: 'Internal Server Error',
		};
	}

	const callJson: StoreInterface = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}

export async function activateShop(shopUuid: string): Promise<ServiceResponse<StoreInterface>> {
	//@ts-ignore
	const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/shopify/shop/${shopUuid}/activate`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
	});

	if (!call.ok) {
		if (call.status === 401) {
			return {
				status: 401,
				message: 'Unauthorized',
			};
		}
		return {
			status: call.status,
			message: 'Internal Server Error',
		};
	}

	const callJson: StoreInterface = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}

export async function deactivateShop(shopUuid: string): Promise<ServiceResponse<StoreInterface>> {
	//@ts-ignore
	const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/shopify/shop/${shopUuid}/deactivate`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
	});

	if (!call.ok) {
		if (call.status === 401) {
			return {
				status: 401,
				message: 'Unauthorized',
			};
		}
		return {
			status: call.status,
			message: 'Internal Server Error',
		};
	}

	const callJson: StoreInterface = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}

export async function getCurrentUser(): Promise<ServiceResponse<UserInterface>> {
	//@ts-ignore
	const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/tagpeak/auth/me`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
	});

	if (!call.ok) {
		if (call.status === 401) {
			return {
				status: 401,
				message: 'Unauthorized',
			};
		}
		return {
			status: call.status,
			message: 'Internal Server Error',
		};
	}

	const callJson: UserInterface = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}

export async function sendVerificationEmail(): Promise<ServiceResponse<void>> {
	//@ts-ignore
	const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/tagpeak/auth/send-verification-email`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
	});

	if (!call.ok) {
		if (call.status === 401) {
			return {
				status: 401,
				message: 'Unauthorized',
			};
		}
		return {
			status: call.status,
			message: 'Internal Server Error',
		};
	}

	return {
		status: 200,
		message: 'OK',
	};
}

export async function checkEmailVerificationStatus(): Promise<
	ServiceResponse<{isVerified: boolean}>
> {
	//@ts-ignore
	const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/tagpeak/auth/email-verification-status`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
	});

	if (!call.ok) {
		if (call.status === 401) {
			return {
				status: 401,
				message: 'Unauthorized',
			};
		}
		return {
			status: call.status,
			message: 'Internal Server Error',
		};
	}

	const callJson: {isVerified: boolean} = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}
