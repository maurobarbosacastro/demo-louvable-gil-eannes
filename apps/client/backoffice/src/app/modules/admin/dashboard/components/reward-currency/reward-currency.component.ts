import {Component, computed, input, InputSignal, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {CustomAccordionComponent} from '@app/shared/components/custom-accordion/custom-accordion.component';
import {MatIcon} from '@angular/material/icon';
import {TableComponent} from '@app/shared/components/table/table.component';
import {TranslocoPipe} from '@ngneat/transloco';
import {IColumn, ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {RewardByCurrenciesInterface} from '@app/modules/admin/dashboard/models/dashboard.interface';
import {rewardStatuses} from '@app/utils/status.enum';
import {AppConstants} from '@app/app.constants';

@Component({
	selector: 'tagpeak-reward-currency',
	standalone: true,
	imports: [CommonModule, CustomAccordionComponent, MatIcon, TableComponent, TranslocoPipe],
	templateUrl: './reward-currency.component.html',
})
export class RewardCurrencyComponent {
	$dashboardRewards: InputSignal<RewardByCurrenciesInterface> =
		input.required<RewardByCurrenciesInterface>();

	$dataSource: Signal<any> = computed(() => {
		const currencies: string[] = AppConstants.CURRENCIES.map((currency): string => currency.key);

		const metrics = rewardStatuses.map((status: string) => ({key: status, label: status}));

		return metrics.map(({key, label}) => {
			const rowData: Record<string, any> = {rowValue: label};

			currencies.forEach((currency: string) => {
				rowData[currency] = this.$dashboardRewards()[key]?.[currency]?.totalRewards ?? 0;
			});

			return rowData;
		});
	});

	$tableConfig: Signal<ITableConfiguration> = computed(
		(): ITableConfiguration => this.setupTableConfigurationRewards(),
	);

	setupTableConfigurationRewards(): ITableConfiguration {
		const columns: IColumn[] = [
			{
				id: 'rowValue',
				name: '',
				hasSort: false,
				hasTooltip: false,
			},
		];

		columns.push(
			...AppConstants.CURRENCIES.map((currency) => ({
				id: currency.key,
				name: `${currency.label} (${currency.symbol})`,
				hasSort: false,
				hasTooltip: false,
			})),
		);

		return {
			dataSource: new MatTableDataSource(this.$dataSource()),
			css: 'rounded-t-[0.625rem] overflow-hidden shadow-table',
			pageSize: 8,
			showPagination: false,
			styles: {
				header: 'font-plex text-project-waterloo text-md font-normal',
				content: 'font-plex text-project-licorice-900 text-md',
			},
			columns: columns,
		};
	}
}
