
export enum WithdrawalStateEnum {
    PENDING = 'PENDING',
    COMPLETED = 'COMPLETED',
    REJECTED = 'REJECTED'
}

export interface WithdrawalsInterface {
    uuid: string;
    user: {
        uuid: string
        email: string
    };
    amountSource: number;
    amountTarget: number;
    details: string;
    state: WithdrawalStateEnum;
    completionDate: string;
    createdAt: string;
    createdBy: string;
    updatedAt: string;
    updatedBy: string;
}

export interface BalanceInterface {
    amountRewards: number;
    amountReferrals: number;
    paidWithdrawals: number;
}

export interface WithdrawalsSearchParams {
    startDate: string;
    endDate: string;
}


export const WithdrawalsSortingFields = {
    uuid: 'uuid',
    amountSource: 'amount_source',
    state: 'state',
    details: 'details',
    completionDate: 'completion_date',
    createdAt: 'created_at',
}
