import {Component, computed, DestroyRef, inject, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MatIcon} from '@angular/material/icon';
import {DashboardService} from '@app/modules/admin/dashboard/services/dashboard.service';
import {
	DashboardInterface,
	RewardByCurrenciesInterface,
	StatisticsByMonthInterface,
	TransactionsStatusInterface,
} from '@app/modules/admin/dashboard/models/dashboard.interface';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {CustomAccordionComponent} from '@app/shared/components/custom-accordion/custom-accordion.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {TableComponent} from '@app/shared/components/table/table.component';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {combineLatest, debounceTime, distinctUntilChanged, map, startWith, switchMap} from 'rxjs';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {AppConstants} from '@app/app.constants';
import {MiniDashboardComponent} from '@app/modules/admin/dashboard/components/mini-dashboard/mini-dashboard.component';
import {CashbackTabEnum} from '@app/modules/admin/cashback/models/cashback.enum';
import {Router} from '@angular/router';
import {RewardCurrencyComponent} from '@app/modules/admin/dashboard/components/reward-currency/reward-currency.component';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-dashboard',
	standalone: true,
	imports: [
		CommonModule,
		TranslocoPipe,
		MatIcon,
		CustomAccordionComponent,
		TableComponent,
		ReactiveFormsModule,
		CustomDropdownComponent,
		MiniDashboardComponent,
		RewardCurrencyComponent,
	],
	templateUrl: './dashboard.component.html',
	styleUrl: './dashboard.component.scss',
})
export class DashboardComponent {
	private readonly _dashboardService: DashboardService = inject(DashboardService);
	private readonly _translocoService: TranslocoService = inject(TranslocoService);
	private readonly _destroyRef: DestroyRef = inject(DestroyRef);
	private readonly _router: Router = inject(Router);
	private readonly _screenService: ScreenService = inject(ScreenService);

	formFilter = new FormGroup({
		year: new FormControl(new Date().getFullYear().toString()),
	});

	$years: Signal<OptionDropdownInterface[]> = computed(() =>
		this.generateYears().map((year) => ({
			label: year.toString(),
			value: year.toString(),
		})),
	);

	styleDropdown: CssDropdownInterface = AppConstants.DEFAULT_STYLE_DROPDOWN;
	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);
	$dashboard: Signal<DashboardInterface> = toSignal(this._dashboardService.getValuesDashboard());

	$dashboardStatistics: Signal<StatisticsByMonthInterface> = toSignal(
		combineLatest([
			this.yearControl.valueChanges.pipe(
				startWith(new Date().getFullYear().toString()),
				debounceTime(500),
				distinctUntilChanged(),
				map((value) => value?.value ?? value),
			),
		]).pipe(
			switchMap(([year]: [string]) => this._dashboardService.getStatisticsDashboard(Number(year))),
			takeUntilDestroyed(this._destroyRef),
		),
	);

	$dashboardTransactions: Signal<TransactionsStatusInterface> = toSignal(
		this._dashboardService.getTransactionsDashboard(),
	);

	$dashboardRewards: Signal<RewardByCurrenciesInterface> = toSignal(
		this._dashboardService.getRewardCountByCurrencies(),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => this.setupTableConfiguration());

	$dataSource: Signal<any[]> = computed(() => {
		const months = [
			{id: `1/${this.yearControl.value}`, key: 'january'},
			{id: `2/${this.yearControl.value}`, key: 'february'},
			{id: `2/${this.yearControl.value}`, key: 'february'},
			{id: `3/${this.yearControl.value}`, key: 'march'},
			{id: `4/${this.yearControl.value}`, key: 'april'},
			{id: `5/${this.yearControl.value}`, key: 'may'},
			{id: `6/${this.yearControl.value}`, key: 'june'},
			{id: `7/${this.yearControl.value}`, key: 'july'},
			{id: `8/${this.yearControl.value}`, key: 'august'},
			{id: `9/${this.yearControl.value}`, key: 'september'},
			{id: `10/${this.yearControl.value}`, key: 'october'},
			{id: `11/${this.yearControl.value}`, key: 'november'},
			{id: `12/${this.yearControl.value}`, key: 'december'},
		];

		const metrics = [
			{key: 'totalUsers', label: 'Total Users'},
			{key: 'activeUsers', label: 'Active Users'},
			{key: 'numTransaction', label: 'Total Transactions'},
			{key: 'totalGMV', label: 'Total GMV'},
			{key: 'avgTransactionAmount', label: 'Average Transaction Amount'},
			{key: 'totalRevenue', label: 'Revenue'},
		];

		return metrics.map(({key, label}) => {
			const rowData: Record<string, any> = {valueRow: label};

			months.forEach((month) => {
				rowData[month.key] = this.$dashboardStatistics()[month.id]?.[key] ?? null;
			});

			return rowData;
		});
	});

	private setupTableConfiguration(): ITableConfiguration {
		// Helper function to generate month data
		const generateMonthColumns = () => {
			const months = [
				{id: 'january', key: 'january'},
				{id: 'february', key: 'february'},
				{id: 'march', key: 'march'},
				{id: 'april', key: 'april'},
				{id: 'may', key: 'may'},
				{id: 'june', key: 'june'},
				{id: 'july', key: 'july'},
				{id: 'august', key: 'august'},
				{id: 'september', key: 'september'},
				{id: 'october', key: 'october'},
				{id: 'november', key: 'november'},
				{id: 'december', key: 'december'},
			];

			return months.map((month) => ({
				id: month.id,
				name: this._translocoService.translate(`dashboard-admin.month-counts.${month.key}`),
				hasSort: false,
				hasTooltip: false,
				colWidth: '7%',
			}));
		};

		return {
			dataSource: new MatTableDataSource(this.$dataSource()),
			css: 'rounded-t-[0.625rem] overflow-hidden shadow-table',
			pageSize: 6,
			showPagination: false,
			styles: {
				header: 'font-plex text-project-waterloo text-md font-normal',
				content: 'font-plex text-project-licorice-900 text-md',
			},
			columns: [
				{
					id: 'valueRow',
					name: '',
					hasSort: false,
					hasTooltip: false,
					colWidth: '15%',
				},
				...generateMonthColumns(),
			],
		};
	}

	generateYears(): number[] {
		const currentYear: number = new Date().getFullYear();
		const years: number[] = [];
		for (let i: number = currentYear; i >= 2022; i--) {
			years.push(i);
		}
		return years;
	}

	setFormControlName(formControlName: string, value: any): void {
		if (value) {
			this.formFilter.get(formControlName).setValue(value?.value);
		} else {
			this.formFilter.get(formControlName).setValue(null);
		}
	}

	get yearControl(): FormControl {
		return this.formFilter.get('year') as FormControl;
	}

	navigateToCashback(columnId: string) {
		this._router.navigateByUrl(
			`/admin/cashbacks?tab=${CashbackTabEnum.MANAGE_CASHBACK}&status=${columnId}`,
		);
	}

	getCompareIcon(elementName: string): string {
		switch (this.$dashboard().indicatorsSection[elementName].compareLastMonth) {
			case 'IMPROVEMENT':
				return 'improvement';
			case 'DOWNGRADE':
				return 'downgrade';
			case 'EQUAL':
				return 'equal_compare';
			default:
				return 'equal_compare';
		}
	}

	getColorTextForPercentage(elementName: string): string {
		switch (this.$dashboard().indicatorsSection[elementName].compareLastMonth) {
			case 'IMPROVEMENT':
				return 'text-[#008000]';
			case 'DOWNGRADE':
				return 'text-project-red';
			case 'EQUAL':
				return 'text-project-waterloo';
			default:
				return 'text-project-waterloo';
		}
	}
}
