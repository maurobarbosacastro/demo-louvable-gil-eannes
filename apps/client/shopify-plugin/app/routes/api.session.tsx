import type {LoaderFunctionArgs} from '@remix-run/node';
import dbServer from '../db.server';

export const loader = async ({request}: LoaderFunctionArgs) => {
	let offlineToken = null;

	try {
		const session = await dbServer.session.findFirst({
			where: {
				shop: request.headers.get('x-shopify-shop-domain')!,
				isOnline: false
			}
		})

		if (session){
			offlineToken = session.accessToken;
		}

	} catch (error) {
		return new Response(JSON.stringify(error), {status: 500})
	}

	return new Response(JSON.stringify({offlineToken}), {
		status: 200,
	});
}
