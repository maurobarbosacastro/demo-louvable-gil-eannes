import {MatTableDataSource} from '@angular/material/table';

export interface ITableConfiguration {
	dataSource: MatTableDataSource<any>;
	pageSize: number;
	css?: string;
	sortStart?: string;
	columns: IColumn[];
	styles?: IStyles;
	expanded?: boolean;
	showPagination?: boolean;
}

export interface IColumn {
	id: string;
	name: string;
	hasSort: boolean;
	hasTooltip: boolean;
	sortActionDescription?: string;
	valueTransform?: (row: any, val: any) => string;
	actions?: IActions[];
	colWidth?: string;
	css?: (row?: any, column?: any) => string;
	stickyEnd?: boolean;
	sticky?: boolean;
}

export interface IActions {
	type: string;
	iconUrl: string;
	css?: string;
	disabled?: (row: any) => boolean;
}

export interface IStyles {
	header?: string;
	content?: string;
	paginator?: string;
}
