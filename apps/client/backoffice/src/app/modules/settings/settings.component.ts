import {Component, inject, input, OnInit, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ReactiveFormsModule} from '@angular/forms';
import {TranslocoService} from '@ngneat/transloco';
import {Router} from '@angular/router';
import {ProfileComponent} from '@app/modules/settings/profile/profile.component';
import {SettingsTabEnum} from '@app/modules/settings/models/settings-tab.enum';
import {WithdrawalSettingsComponent} from '@app/modules/settings/withdrawals-settings/withdrawal-settings.component';
import {PrivacyDetailsComponent} from '@app/modules/settings/privacy-details/privacy-details.component';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {CountryInterface} from '@app/modules/admin/countries/interfaces/country.interface';

@Component({
	selector: 'tagpeak-settings',
	standalone: true,
	imports: [
		CommonModule,
		ReactiveFormsModule,
		ProfileComponent,
		WithdrawalSettingsComponent,
		PrivacyDetailsComponent,
	],
	templateUrl: './settings.component.html',
})
export class SettingsComponent implements OnInit {
	_router = inject(Router);
	transloco = inject(TranslocoService);
	private queryService: UrlQueryService = inject(UrlQueryService);

	//Tabs with the route name
	tabs: SettingsTabEnum[] = [
		SettingsTabEnum.PROFILE,
		SettingsTabEnum.WITHDRAWAL_SETTINGS,
		SettingsTabEnum.PRIVACY,
	];

	$selectedTab: WritableSignal<SettingsTabEnum> = signal(SettingsTabEnum.PROFILE);
	$preLoadedCountry = input<CountryInterface[]>();

	ngOnInit() {
		if (this.queryService.getParam('tab')) {
			this.$selectedTab.set(this.queryService.getParam('tab') as SettingsTabEnum);
		} else {
			this.queryService.addParam('tab', this.tabs[0]);
		}
	}

	getTabName(tab: string): string {
		return this.transloco.translate(`settings.${tab}`);
	}

	goToTab(tab: SettingsTabEnum) {
		this.$selectedTab.update((): SettingsTabEnum => tab);
		this.queryService.addParam('tab', tab);
	}

	protected readonly SettingsTabEnum = SettingsTabEnum;
}
