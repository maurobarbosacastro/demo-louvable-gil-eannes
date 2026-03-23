export async function checkTPTokenValidity(): Promise<boolean> {
//@ts-ignore
  const env = window.ENV;

  try {
    const token = sessionStorage.getItem('tp_accessToken');
    if (!token) return false;

    const response = await fetch(`${env.BE_URL}/tagpeak/auth/valid`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': token,
      },
    });

    return response.ok;
  } catch (error) {
    return false;
  }
}


export function logout(shop: string) {
//@ts-ignore
  const env = window.ENV;

  sessionStorage.removeItem('tp_accessToken');
  sessionStorage.removeItem('shop_uuid');
  sessionStorage.removeItem('tp_shop');
  sessionStorage.removeItem('shop');
  window.open(`${env.BO_URL}/#/shopify/shop-sign-in?shop=${shop}`, '_top');
}
