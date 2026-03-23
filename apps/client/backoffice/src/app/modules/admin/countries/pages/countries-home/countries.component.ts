import {Component, inject, OnInit, Signal, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {TableComponent} from '@app/shared/components/table/table.component';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {BehaviorSubject, combineLatest, filter, switchMap, tap} from 'rxjs';
import {MatTableDataSource} from '@angular/material/table';
import {PageEvent} from '@angular/material/paginator';
import {CountriesService} from '@app/modules/admin/countries/services/countries.service';
import {CountryInterface} from '@app/modules/admin/countries/interfaces/country.interface';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {EditCountryComponent} from '@app/modules/admin/countries/components/edit-country/edit-country.component';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {AppConstants} from '@app/app.constants';
import {ScreenService} from '@app/core/services/screen.service';
import {toSignal} from '@angular/core/rxjs-interop';

@Component({
	selector: 'tagpeak-countries',
	standalone: true,
	imports: [CommonModule, MatIcon, TableComponent, TranslocoPipe],
	templateUrl: './countries.component.html',
	styleUrl: '../../../../../../styles/style-table.scss',
})
export class CountriesComponent implements OnInit {
	private _countriesService = inject(CountriesService);
	private _dialog = inject(MatDialog);
	private _toaster = inject(ToasterService);
	private transloco = inject(TranslocoService);
	private _screenService: ScreenService = inject(ScreenService);

	page$: BehaviorSubject<number> = new BehaviorSubject(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject('abbreviation');
	tableConfiguration: ITableConfiguration;
	totalSize: number = 1;
	pageSize = 10;
	page: number;
	total: number;

	$data: WritableSignal<CountryInterface[]> = signal([]);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);
	ngOnInit(): void {
		this.getCountries(this.pageSize);
		this.setTableConfiguration();
	}

	openEditCountryDialog(countryUuid?: string): void {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: {countryUuid: countryUuid},
		};

		if (this.$isMobile()) {
			modalConfig.width = '100%';
			modalConfig.height = '95%';
			modalConfig.panelClass = 'dialog-panel-details-mobile';
		}

		const dialogRef = this._dialog.open(EditCountryComponent, modalConfig);

		dialogRef.afterClosed().subscribe((result) => {
			if (!result.skipFetch) {
				this.page$.next(1);
			}
		});
	}

	clickTableAction(event: any): void {
		switch (event.type) {
			case 'edit': {
				this.openEditCountryDialog(event.element.uuid);
				break;
			}
			case 'delete': {
				this.deleteCountry(event.element.uuid);
				break;
			}
		}
	}

	deleteCountry(id: string): void {
		this._dialog
			.open(DeleteModalConfirmationComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '26rem',
			})
			.afterClosed()
			.pipe(
				filter((result: {skipFetch: boolean}) => result && !result.skipFetch),
				switchMap(() => this._countriesService.deleteCountry(id)),
			)
			.subscribe(
				(_) => {
					this.page$.next(1);
					this._toaster.showToast('success', this.transloco.translate('country.country-deleted'));
				},
				(error) => {
					this._toaster.showToast(
						'error',
						this.transloco.translate('country.country-delete-error'),
					);
				},
			);
	}

	setTableConfiguration(): void {
		this.tableConfiguration = {
			dataSource: new MatTableDataSource([]),
			css: 'rounded-t-[0.625rem] overflow-hidden shadow-table',
			pageSize: this.pageSize,
			styles: {
				header: 'font-plex text-project-waterloo text-md font-normal',
				content: 'font-plex text-project-licorice-900 text-md',
				paginator: 'font-plex text-black',
			},
			columns: [
				{
					id: 'abbreviation',
					name: 'Code',
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'id',
					colWidth: '15%',
				},
				{
					id: 'name',
					name: 'Country Name',
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'Country Name',
					colWidth: '20%',
				},
				{
					id: 'currency',
					name: 'Currency',
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'Currency',
					colWidth: '15%',
				},
				{
					id: 'enabled',
					name: 'Enabled',
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'Enabled',
					colWidth: '15%',
				},
				{
					id: 'actions',
					name: 'Actions',
					hasSort: false,
					hasTooltip: false,
					colWidth: '15%',
					actions: [
						{
							type: 'edit',
							iconUrl: 'pen',
							css: 'scale-75',
						},
						{
							type: 'delete',
							iconUrl: 'tag-delete',
						},
					],
				},
			],
		};

		this.setTableData();
	}

	setTableData(): void {
		this.tableConfiguration = {
			...this.tableConfiguration,
			dataSource: new MatTableDataSource(this.$data()),
		};
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			this.sort$.next(evt.sort.split(',').join(' '));
		}
		this.page$.next(evt.page.pageIndex);
	}

	private getCountries(pageSize: number) {
		combineLatest([this.sort$, this.page$])
			.pipe(
				switchMap(([sort, page]) => {
					return this._countriesService.getCountries(pageSize, page, sort);
				}),
				tap((res) => {
					this.page = res.page;
					this.total = res.totalPages;
					this.totalSize = res.totalRows;
				}),
			)
			.subscribe((res) => {
				this.$data.set(res.data as CountryInterface[]);
				this.setTableData();
			});
	}
}
