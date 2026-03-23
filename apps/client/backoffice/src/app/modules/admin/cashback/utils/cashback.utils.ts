import {CashbackTableResponse} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {StatusEnum} from '@app/utils/status.enum';

export const showHistoryBtn = (cashback: CashbackTableResponse): boolean => {
    if (cashback.status === StatusEnum.TRACKED || cashback.status === StatusEnum.REJECTED) {
        return false;
    }

    if (cashback.status === StatusEnum.VALIDATED && !cashback.reward) {
        return false;
    }
    return true;
}
