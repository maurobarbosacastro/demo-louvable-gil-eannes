import {Component, inject, OnInit, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from "@angular/material/icon";
import {TableComponent} from "@app/shared/components/table/table.component";
import {TranslocoPipe, TranslocoService} from "@ngneat/transloco";
import {ITableConfiguration} from "@app/shared/components/table/table.interface";
import {BehaviorSubject, combineLatest, switchMap, tap} from "rxjs";
import {MatTableDataSource} from "@angular/material/table";
import {PageEvent} from "@angular/material/paginator";
import {CountriesService} from "@app/modules/admin/countries/countries.service";
import {CountryInterface} from "@app/shared/interfaces/country.interface";
import {MatDialog} from "@angular/material/dialog";
import {EditCountryComponent} from "@app/modules/admin/countries/edit-country/edit-country.component";
import {ToasterService} from "@app/shared/components/toaster/toaster.service";

@Component({
    selector: 'tagpeak-countries',
    standalone: true,
    imports: [CommonModule, MatIcon, TableComponent, TranslocoPipe],
    templateUrl: './countries.component.html',
})
export class CountriesComponent implements OnInit {

    private _countriesService = inject(CountriesService);
    private _dialog = inject(MatDialog);
    private _toaster = inject(ToasterService);
    private transloco = inject(TranslocoService);

    page$: BehaviorSubject<number> = new BehaviorSubject(1);
    sort$: BehaviorSubject<string> = new BehaviorSubject("abbreviation");
    tableConfiguration: ITableConfiguration;
    totalSize: number = 1;
    pageSize = 10
    page: number;
    total: number;

    $data: WritableSignal<CountryInterface[]> = signal([]);

    ngOnInit(): void {
        this.getCountries(this.pageSize);
        this.setTableConfiguration();
    }

    openEditCountryDialog(country_uuid?: string): void {

        const dialogRef = this._dialog.open(EditCountryComponent,
            {
                panelClass: 'onboarding-dialog-panel',
                position: {right: '2rem'},
                data: {country_uuid: country_uuid}
            });

        dialogRef.afterClosed().subscribe(result => {
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
        this._countriesService.deleteCountry(id).subscribe(
            value => {
                this.getCountries(this.pageSize);
                this._toaster.showToast('success', this.transloco.translate('country.country-deleted'));
            },
            error => {
                this._toaster.showToast('error', this.transloco.translate('country.country-delete-error'));

            }
        );
    }

    setTableConfiguration(): void {
        this.tableConfiguration = {
            dataSource: new MatTableDataSource([]),
            css: 'rounded-t-[0.625rem] overflow-hidden shadow-table',
            pageSize: this.pageSize,
            styles: {
                header: 'font-nunitoBlack text-black bg-project-tableGray',
                content: 'text-black',
                paginator: 'text-black'
            },
            columns: [
                {
                    id: 'abbreviation',
                    name: 'Code',
                    hasSort: true,
                    hasTooltip: false,
                    sortActionDescription: 'id',
                    colWidth: '15%'
                },
                {
                    id: 'name',
                    name: 'Country Name',
                    hasSort: true,
                    hasTooltip: false,
                    sortActionDescription: 'Country Name',
                    colWidth: '20%'
                },
                {
                    id: 'currency',
                    name: 'Currency',
                    hasSort: true,
                    hasTooltip: false,
                    sortActionDescription: 'Currency',
                    colWidth: '15%'
                },
                {
                    id: 'enabled',
                    name: 'Enabled',
                    hasSort: true,
                    hasTooltip: false,
                    sortActionDescription: 'Enabled',
                    colWidth: '15%'
                },
                {
                    id: 'actions',
                    name: 'Actions',
                    hasSort: false,
                    hasTooltip: false,
                    colWidth: '10%',
                    actions: [
                        {
                            type: 'edit',
                            iconUrl: 'pen',
                            css: 'scale-75'
                        },
                        {
                            type: 'delete',
                            iconUrl: 'tag-delete'
                        }
                    ]
                }
            ]
        };

        this.setTableData();
    }

    setTableData(): void {
        this.tableConfiguration = {...this.tableConfiguration, dataSource: new MatTableDataSource(this.$data())};
    }

    onSortAndPageChange(evt: { sort: string, page: PageEvent }): void {
        if (evt.sort) {
            this.sort$.next(evt.sort.split(',').join(' '));
        }
        this.page$.next(evt.page.pageIndex);
    }

    private getCountries(pageSize: number) {
        combineLatest([
            this.sort$,
            this.page$
        ]).pipe(
            switchMap(([sort, page]) => {
                return this._countriesService.getCountries(pageSize, page, sort);
            })
            , tap(res => {
                this.page = res.page;
                this.total = res.total_pages;
                this.totalSize = res.total_rows;
            })
        ).subscribe(res => {
            this.$data.set(res.data as CountryInterface[]);
            this.setTableData();
        });
    }


}
