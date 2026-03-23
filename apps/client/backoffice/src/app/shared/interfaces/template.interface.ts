export interface TemplateInterface {
    id: string;
    code: string;
    name: string;
    templateJson: string;
    templateHtml: string;
    createdAt: Date;
    updatedAt: Date;
    createdBy: string;
    updatedBy: string;
}

export type EmailTableFields = Pick<TemplateInterface, 'id' | 'name'>;

export type TemplateCreation = Pick<TemplateInterface, 'name' | 'templateJson' | 'templateHtml' | 'code'>;

export enum PageMode {
    new = 'new',
    edit = 'edit',
}

