import {StatusEnum} from "@app/utils/status.enum";

export interface Reward {
    uuid: string;
    store: Store;
    pricePaid: number;
    stopDate?: StopDate;
    currentReward?: number;
    status: StatusEnum;
}

export interface Store {
    uuid: string;
    name: string;
    logo: string;
    purchaseDate: string;
    link?: string;
}

export interface StopDate {
    date: string;
    daysLeft: number;
}

export interface RewardSummary {
    liveRewards: number;
    availableForWithdrawal: number;
    valueSpent: number;
}

export interface RewardDetails extends Reward {
    refId: number;
    investments: { x: string; y: number }[];
}
