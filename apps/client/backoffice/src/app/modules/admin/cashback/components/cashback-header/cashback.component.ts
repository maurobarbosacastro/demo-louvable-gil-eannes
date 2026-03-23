import {Component, DestroyRef, inject, OnInit, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe, TranslocoService} from "@ngneat/transloco";
import {CashbackTabEnum} from "@app/modules/admin/cashback/models/cashback.enum";
import {
    ManageCashbackComponent
} from "@app/modules/admin/cashback/pages/manage-cashback/manage-cashback.component";
import {
    CashoutRequestComponent
} from "@app/modules/admin/cashout/components/cashout-request/cashout-request.component";
import {UrlQueryService} from '@app/shared/services/url-query.service';

@Component({
	selector: 'tagpeak-cashback',
	standalone: true,
	imports: [CommonModule, TranslocoPipe, ManageCashbackComponent, CashoutRequestComponent],
	templateUrl: './cashback.component.html',
})
export class CashbackComponent implements OnInit {
	private translocoService: TranslocoService = inject(TranslocoService);
	private queryService: UrlQueryService = inject(UrlQueryService);
	private destroyRef: DestroyRef = inject(DestroyRef);

	//Tab names
	tabs: CashbackTabEnum[] = [CashbackTabEnum.MANAGE_CASHBACK, CashbackTabEnum.CASHOUT_REQUEST];

	$selectedTab: WritableSignal<CashbackTabEnum> = signal(this.tabs[0]);

	ngOnInit() {
        if (this.queryService.getParam('tab')) {
            this.$selectedTab.set(this.queryService.getParam('tab') as CashbackTabEnum);
        } else {
            this.queryService.addParam('tab', this.tabs[0]);
        }
    }

	getTabName(tab: CashbackTabEnum): string {
		return this.translocoService.translate(`cashback.tabs.${tab}`);
	}

	selectTab(tab: CashbackTabEnum) {
		this.$selectedTab.update(() => tab);
		this.queryService.addParam('tab', tab);
	}

	protected readonly CashbackTabEnum = CashbackTabEnum;
}
