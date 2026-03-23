import {Component, DestroyRef, inject, OnInit, Signal, signal, WritableSignal} from '@angular/core';
import {CommonModule, Location} from '@angular/common';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {YourRewardsComponent} from '@app/modules/client/rewards/components/your-rewards/your-rewards.component';
import {DashboardTabEnum} from '@app/modules/client/dashboard/models/dashboard-tab.enum';
import {StoreVisitsComponent} from '@app/modules/client/store-visits/components/store-visits.component';
import {BalanceWithdrawalsComponent} from '@app/modules/client/balance-withdrawals/components/balance-withdrawals.component';
import {Router} from '@angular/router';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {catchError, distinctUntilChanged, filter, map, of, throwError} from 'rxjs';
import {UserService} from '@app/core/user/user.service';
import {DialogCurrencySelectionComponent} from '@app/modules/client/dashboard/components/dialog-currency-selection/dialog-currency-selection.component';
import {MatDialog} from '@angular/material/dialog';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
    selector: 'tagpeak-dashboard',
    standalone: true,
    imports: [
        CommonModule,
        TranslocoPipe,
        YourRewardsComponent,
        StoreVisitsComponent,
        BalanceWithdrawalsComponent,
    ],
    templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {
	private translocoService: TranslocoService = inject(TranslocoService);
	private _router = inject(Router);
	private location = inject(Location);
	private queryService: UrlQueryService = inject(UrlQueryService);
	private destroyRef: DestroyRef = inject(DestroyRef);
	private _userService = inject(UserService);
	private dialog = inject(MatDialog);
	private toastr = inject(ToasterService);
	private _screenService: ScreenService = inject(ScreenService);

    //Tabs name
    tabs: DashboardTabEnum[] = [
        DashboardTabEnum.YOU_REWARDS,
        DashboardTabEnum.STORE_VISITS,
        DashboardTabEnum.BALANCE,
    ];

	$selectedTab: WritableSignal<DashboardTabEnum> = signal(this.tabs[0]);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);


    ngOnInit() {
        if (this.queryService.getParam('tab')) {
            this.$selectedTab.set(this.queryService.getParam('tab') as DashboardTabEnum);
        } else {
            this.queryService.addParam('tab', this.tabs[0]);
        }

        this.queryService.params$
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((params) => params.get('tab') !== null),
                map((params) => params.get('tab')),
                distinctUntilChanged(),
            )
            .subscribe((params) => {
                this.$selectedTab.set(params as DashboardTabEnum);
            });

		if (
			this._userService.$user().source === 'Shopify' &&
			!this._userService.$user().currencySelected
		) {
			this.openCurrencyDialog();
		}
	}

    getNameTabs(tab: DashboardTabEnum): string {
        return this.translocoService.translate(`dashboard.tabs.${tab}`);
    }

    selectTab(tab: DashboardTabEnum): void {
        this.$selectedTab.update((): DashboardTabEnum => tab);
        this.queryService.addParam('tab', tab);
    }

    goToStore(): void {
        this._router.navigate(['stores']);
    }

    protected readonly DashboardTabEnum = DashboardTabEnum;

    openCurrencyDialog(): void {
        const dialogRef = this.dialog.open(DialogCurrencySelectionComponent, {
            panelClass: 'verify-dialog-panel',
            disableClose: true,
            data: null,
        });

		dialogRef.afterClosed().subscribe((data) => {
			this.toastr.showToast(
				'info',
				this.translocoService.translate('settings.dialog-currency.reload_info'),
				'top',
				null,
				10000,
			);
			setTimeout(() => {
				window.location.reload();
			}, 10000);
		});
	}
}
