import { BaseInterface } from '@app/shared/interfaces/general.interface';

export interface PartnersInterface extends BaseInterface{
    uuid: string;
    name: string;
    code: string;
    eCommercePlatform: string;
    commissionRate: number;
    validationPeriod: number;
    deepLink: string;
    deepLinkIdentifier: string;
    subIdentifier: string;
    percentageTagpeak: number;
    percentageInvested: number;
}

export const PartnerSortingFields = {
    name: 'name',
    code: 'code',
    eCommercePlatform: 'e_commerce_platform',
    commissionRate: 'commission_rate',
    validationPeriod: 'validation_period',
    deepLink: 'deep_link',
    deepLinkIdentifier: 'deep_link_identifier',
    subIdentifier: 'sub_identifier',
    percentageTagpeak: 'percentage_tagpeak',
    percentageInvested: 'percentage_invested',
}
