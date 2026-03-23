import {BaseInterface} from '@app/shared/interfaces/general.interface';

export interface Configuration extends BaseInterface{
    id: number;
    code: string;
    name: string;
    value: string;
    dataType: string;
}

export type PatchConfiguration = Partial<Pick<Configuration, 'name' | 'value'>>

export const ConfigurationsSortingFields = {
    id: 'id',
    name: 'name',
    code: 'code',
    type: 'data_type'
}
