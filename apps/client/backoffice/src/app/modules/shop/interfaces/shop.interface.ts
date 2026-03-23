import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {CashbackTableResponse} from '@app/modules/admin/cashback/models/manage-cashback.interface';

export type TagpeakShopUser = TagpeakUser & {
    shop: string;
}

export type CashbackShopTableResponse = Omit<CashbackTableResponse, 'reward'>

export interface DashboardStats {
    totalOrders: number;
    totalAmount: number;
    currency: string;
}

export interface ShopifyShop {
    uuid: string;
    shopUuid: string;
    userUuid: string;
    storeUuid: string;
}
