import {Component, computed, EventEmitter, input, InputSignal, Output, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TransactionsStatusInterface} from '@app/modules/admin/dashboard/models/dashboard.interface';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {TableComponent} from '@app/shared/components/table/table.component';
import {MatIcon} from '@angular/material/icon';

@Component({
	selector: 'tagpeak-mini-dashboard',
	standalone: true,
	imports: [CommonModule, TableComponent, MatIcon],
	templateUrl: './mini-dashboard.component.html',
})
export class MiniDashboardComponent {
	$transactionsDashboard: InputSignal<TransactionsStatusInterface> =
		input.required<TransactionsStatusInterface>();

	@Output() navigateToPage: EventEmitter<string> = new EventEmitter<string>();

	tableConfig: Signal<ITableConfiguration> = computed(() => this.setupTableConfiguration());

	$dataSource: Signal<any> = computed(() => {
		const status: string[] = ['LIVE', 'TRACKED', 'VALIDATED', 'FINISHED', 'STOPPED', 'EXPIRED'];

		const metrics = [
			{key: 'value', label: 'Value'},
			{key: 'count', label: 'Count'},
			{key: 'warning', label: 'Warning'},
		];

		return metrics.map(({key, label}) => {
			const rowData: Record<string, any> = {rowValue: label};

			status.forEach((status: string) => {
				rowData[status] = this.$transactionsDashboard()[status]?.[key] ?? this.defaultValue;
			});

			return rowData;
		});
	});

	defaultValue: string = '-';

	setupTableConfiguration(): ITableConfiguration {
		return {
			dataSource: new MatTableDataSource<any>(this.$dataSource()),
			pageSize: 3,
			showPagination: false,
			styles: {
				header: 'font-plex text-project-waterloo text-md font-normal',
				content: 'font-plex text-project-licorice-900 text-md',
			},
			columns: [
				{
					id: 'rowValue',
					name: '',
					hasSort: false,
					hasTooltip: false,
				},
				{
					id: 'TRACKED',
					name: 'Tracked',
					hasSort: false,
					hasTooltip: false,
				},
				{
					id: 'VALIDATED',
					name: 'Validated',
					hasSort: false,
					hasTooltip: false,
				},
				{
					id: 'LIVE',
					name: 'Live',
					hasSort: false,
					hasTooltip: false,
				},
				{
					id: 'STOPPED',
					name: 'Stopped',
					hasSort: false,
					hasTooltip: false,
				},
				{
					id: 'EXPIRED',
					name: 'Expired',
					hasSort: false,
					hasTooltip: false,
				},
				{
					id: 'FINISHED',
					name: 'Finished',
					hasSort: false,
					hasTooltip: false,
				},
			],
		};
	}

	navigateToCashback(implicit: any) {
		this.navigateToPage.emit(implicit.column.id);
	}
}
