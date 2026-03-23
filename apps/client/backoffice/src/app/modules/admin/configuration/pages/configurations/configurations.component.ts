import {Component, computed, DestroyRef, inject, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ConfigurationService} from '../../services/configuration.service';
import {BehaviorSubject, combineLatest, map, switchMap, tap} from 'rxjs';
import {AppConstants} from '@app/app.constants';
import {
	Configuration,
	ConfigurationsSortingFields,
} from '@app/modules/admin/configuration/interfaces/configuration.interface';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {TableComponent} from '@app/shared/components/table/table.component';
import {PageEvent} from '@angular/material/paginator';
import {ManageConfigurationComponent} from '@app/modules/admin/configuration/components/manage-configuration/manage-configuration.component';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-configurations',
	standalone: true,
	imports: [CommonModule, TableComponent, TranslocoPipe],
	templateUrl: './configurations.component.html',
	styleUrl: '../../../../../../styles/style-table.scss',
})
export class ConfigurationsComponent {
	private configurationService: ConfigurationService = inject(ConfigurationService);
	private destroyRef: DestroyRef = inject(DestroyRef);
	private translocoService: TranslocoService = inject(TranslocoService);
	private matDialog = inject(MatDialog);
	private _screenService: ScreenService = inject(ScreenService);

	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject<string>('id asc');

	pageSize: number = AppConstants.PAGE_SIZE;
	totalSize: number = 1;
	page: number;
	totalPages: number;

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);
	$dataNew: Signal<Configuration[]> = toSignal(
		combineLatest([this.page$, this.sort$]).pipe(
			takeUntilDestroyed(this.destroyRef),
			switchMap(([page, sort]) =>
				this.configurationService.getConfigurations(page, this.pageSize, sort),
			),
			tap((res) => {
				this.page = res.page;
				this.totalPages = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			map((res) => res.data ?? []),
		),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => {
		return {
			dataSource: new MatTableDataSource(this.$dataNew()),
			pageSize: this.pageSize,
			styles: {
				header: 'font-nunitoBlack text-project-waterloo',
				content: 'font-nunitoRegular text-project-licorice-900',
				paginator: 'font-nunitoRegular text-black',
			},
			columns: [
				{
					id: 'id',
					name: this.translocoService.translate('configurations.table.id'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'name',
					name: this.translocoService.translate('configurations.table.name'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'code',
					name: this.translocoService.translate('configurations.table.code'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'value',
					name: this.translocoService.translate('configurations.table.value'),
					hasSort: false,
					hasTooltip: false,
				},
				{
					id: 'type',
					name: this.translocoService.translate('configurations.table.type'),
					hasSort: true,
					hasTooltip: false,
					valueTransform: (row) => row.dataType,
				},
				{
					id: 'actions',
					name: this.translocoService.translate('configurations.table.actions'),
					hasSort: false,
					hasTooltip: false,
					actions: [
						{
							type: 'edit',
							iconUrl: 'edit',
							css: 'text-project-licorice-500 hover:text-white',
						},
					],
				},
			],
		};
	});

	onSortAndPageChange(event: {sort: string; page: PageEvent}) {
		if (event.sort) {
			const field = ConfigurationsSortingFields[event.sort.split(',')[0]];
			const direction = event.sort.split(',')[1];
			this.sort$.next(`${field} ${direction}`);
		}
		this.page$.next(event.page.pageIndex);
	}

	handleActionClick(event: {type: string; element: Configuration; event: any}) {
		switch (event.type) {
			case 'edit': {
				this.editConfiguration(event.element);
			}
		}
	}

	editConfiguration(config: Configuration) {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: config,
		};

		if (this.$isMobile()) {
			modalConfig.width = '100%';
			modalConfig.height = '95%';
			modalConfig.panelClass = 'dialog-panel-details-mobile';
		}

		this.matDialog
			.open(ManageConfigurationComponent, modalConfig)
			.afterClosed()
			.subscribe((result) => {
				if (!result.skipFetch) {
					this.page$.next(1);
				}
			});
	}
}
