import {
	AfterViewInit,
	Component,
	EventEmitter,
	Input,
	OnChanges,
	Output,
	SimpleChanges,
	TemplateRef,
	ViewChild,
} from '@angular/core';
import {MatSort, MatSortModule, Sort} from '@angular/material/sort';
import {MatTable, MatTableModule} from '@angular/material/table';
import {MatPaginator, MatPaginatorModule, PageEvent} from '@angular/material/paginator';
import {MatCheckboxChange, MatCheckboxModule} from '@angular/material/checkbox';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import _ from 'lodash';
import {animate, state, style, transition, trigger} from '@angular/animations';
import {MatFormFieldModule} from '@angular/material/form-field';
import {TranslocoModule} from '@ngneat/transloco';
import {NgClass, NgForOf, NgIf, NgTemplateOutlet} from '@angular/common';
import {MatButtonModule} from '@angular/material/button';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatIconModule} from '@angular/material/icon';
import {FormsModule} from '@angular/forms';

@Component({
	selector: 'atl-ng-table',
	standalone: true,
	imports: [
		MatPaginatorModule,
		MatTableModule,
		MatSortModule,
		MatFormFieldModule,
		TranslocoModule,
		NgClass,
		NgForOf,
		NgIf,
		MatButtonModule,
		MatTooltipModule,
		MatIconModule,
		MatCheckboxModule,
		FormsModule,
		NgTemplateOutlet,
	],
	templateUrl: './table.component.html',
	styleUrls: ['./table.component.scss'],
	animations: [
		trigger('detailExpand', [
			state('collapsed', style({height: '0px', minHeight: '0'})),
			state('expanded', style({height: '*'})),
			transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
		]),
	],
})
export class TableComponent implements AfterViewInit, OnChanges {
	@Input() tableConfiguration: ITableConfiguration;
	@Input() headerCustomTemplate?: TemplateRef<any>;
	@Input() bodyCustomTemplate?: TemplateRef<any>;
	@Input() bodyCustomExpandedTemplate?: TemplateRef<any>;
	@Input() totalPages: number = 0;
	@Input() totalSize: number = 0;
	@Input() resetPaginator: number = 1;

	@Output() onRowClick: EventEmitter<{element: any}> = new EventEmitter();
	@Output() onActionClick: EventEmitter<{
		type: string;
		element: any;
		event: MouseEvent | MatCheckboxChange;
	}> = new EventEmitter();
	@Output() onSortAndPageChange: EventEmitter<{sort: string; page: PageEvent}> = new EventEmitter();
	@ViewChild(MatSort) sort: MatSort;
	@ViewChild(MatPaginator) paginator: MatPaginator;

	columnList: string[] = [];
	columnsToDisplayWithExpand: string[];
	currentPage = 0;

	expandedElement = null;

	/*
        Allow access the data in table to modify it. Example:
        (declare in the parent component a ViewChild of the component:  @ViewChild(TableComponent) table: TableComponent;)
        let dataSource: MatTableDataSource<IItem> = this.table.table.dataSource as MatTableDataSource<any>;
        const idxItem = dataSource.data.findIndex((itemDS) => itemDS.id === item.id);
        dataSource.data[idxItem].enable = res.enable; */
	@ViewChild(MatTable, {static: true}) table!: MatTable<any>;

	protected readonly Number = Number;

	ngOnChanges(changes: SimpleChanges): void {
		// If showPagination is not set, set it to true
		if (changes.tableConfiguration && this.tableConfiguration) {
			this.tableConfiguration.showPagination ??= true;
		}

		if (
			changes.tableConfiguration &&
			!_.isEqual(changes.tableConfiguration.currentValue, changes.tableConfiguration.previousValue)
		) {
			this.columnList = this.tableConfiguration.columns.map((col) => col.id);
			this.columnsToDisplayWithExpand = [...this.columnList, 'expand'];
		}
		if (this.paginator && this.resetPaginator === 1) {
			this.paginator.firstPage();
		}
	}

	ngAfterViewInit() {
		this.setSortAndPaginator();
		this.expandedElement = this.tableConfiguration.expanded;
	}

	setSortAndPaginator(): void {
		this.tableConfiguration.dataSource.sort = this.sort;
		this.tableConfiguration.dataSource.paginator = this.paginator;
	}

	/** Announce the change in sort state for assistive technology. */
	announceSortChange(sortState: Sort) {
		// This example uses English messages. If your application supports
		// multiple language, you would internationalize these strings.
		// Furthermore, you can customize the message to add additional
		// details about the values being sorted.
		if (sortState.direction) {
			this.emitPageChange(`${sortState.active},${sortState.direction}`);
		} else {
			this.emitPageChange();
		}
	}

	hasDataToDisplay(): boolean {
		return this.totalResults() > 0;
	}

	totalResults(): number {
		return this.tableConfiguration ? this.tableConfiguration.dataSource.data.length : 0;
	}

	emitActionClick(row: any, event: MouseEvent | MatCheckboxChange, action: any): void {
		if (!(event instanceof MatCheckboxChange)) {
			event.stopPropagation();
		}
		this.onActionClick.emit({type: action.type, element: row, event});
	}

	rowClick(row: any): void {
		this.onRowClick.emit({element: row});
	}

	getRowValue(row: any, column: any): string {
		return column.valueTransform ? column.valueTransform(row, column) : row[column.id];
	}

	previousPage(): void {
		this.paginator.previousPage();
	}

	nextPage(): void {
		this.paginator.nextPage();
	}

	emitPageChange(sort?: string, nextPage?: number) {
		const sortString = this.getSortString();
		const pageEvent = new PageEvent();
		pageEvent.pageIndex = this.paginator.pageIndex + 1;
		pageEvent.pageSize = this.paginator.pageSize;
		pageEvent.length = this.paginator.length;

		this.onSortAndPageChange.emit({sort: sort ? sort : sortString, page: pageEvent});
	}

	getSortString(): string {
		const activeSort = this.sort.active;
		const direction = this.sort.direction;

		if (activeSort && direction) {
			return `${activeSort},${direction}`;
		}
		return null;
	}

	onPaginateChange(event: PageEvent): void {
		this.currentPage = event.pageIndex;
		this.emitPageChange();
	}

	getCssValue(row, column: any): string {
		if (column.css) {
			return column.css(row, column) ?? this.tableConfiguration.styles.content + column.css;
		}
		return this.tableConfiguration.styles.content;
	}

	visiblePages(): (number | string)[] {
		const pages: (number | string)[] = [];

		if (this.totalPages <= 7) {
			return Array.from({length: this.totalPages}, (_, i) => i);
		}

		pages.push(0);
		if (this.currentPage > 2) {
			pages.push('...');
		}

		const start = Math.max(1, this.currentPage - 1);
		const end = Math.min(this.totalPages - 2, this.currentPage + 1);

		for (let i = start; i <= end; i++) {
			pages.push(i);
		}

		if (this.currentPage < this.totalPages - 3) {
			pages.push('...');
		}
		pages.push(this.totalPages - 1);

		return pages;
	}

	changePage(page: number | string) {
		this.paginator.pageIndex = Number(page);
		this.currentPage = Number(page);
		this.emitPageChange();
	}
}
