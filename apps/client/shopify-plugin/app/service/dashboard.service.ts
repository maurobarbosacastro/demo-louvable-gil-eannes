import type {
  BulkEditTransaction, CashbackTableResponse,
  DashboardStats,
  ServiceResponse,
} from '../interfaces/dashbaord.interface';
import type {PaginationInterface} from '../interfaces/common.interface';


export async function getDashboardStats(
	uuid: string,
): Promise<ServiceResponse<DashboardStats>> {
  //@ts-ignore
  const env = window.ENV;

	const call = await fetch(`${env.BE_URL}/tagpeak/shopify/shop/${uuid}/stats`, {
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

	const callJson: DashboardStats = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}

export async function getCashbacks(
	uuid: string,
	limit: number = 30,
	page: number = 0,
	sort: string = 'date desc',
): Promise<ServiceResponse<PaginationInterface<CashbackTableResponse>>> {
  //@ts-ignore
  const env = window.ENV;

  const url = new URL(`${env.BE_URL}/tagpeak/shopify/store/${uuid}/transaction`);
	url.searchParams.append('limit', limit.toString());
	url.searchParams.append('sort', sort);
	url.searchParams.append('page', page.toString());

	const call = await fetch(url, {
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

	const callJson: any = await call.json();
	return {
		status: 200,
		message: 'OK',
		data: callJson,
	};
}

export async function bulkEditTransaction(
	body: BulkEditTransaction,
): Promise<ServiceResponse<any>> {
  //@ts-ignore
  const env = window.ENV;

  const call = await fetch(`${env.BE_URL}/tagpeak/transaction/bulk/edit`, {
		method: 'PATCH',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': sessionStorage.getItem('tp_accessToken') as string,
		},
		body: JSON.stringify(body),
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
		data: null,
	};
}
