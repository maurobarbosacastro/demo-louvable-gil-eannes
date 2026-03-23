import {Component, computed, inject, OnInit, Signal, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {ReferralsTabEnum} from '@app/modules/client/referrals/models/referrals.interface';
import {YourReferralsComponent} from '@app/modules/client/referrals/components/your-referrals/your-referrals.component';
import {HowItWorksComponent} from '@app/modules/client/referrals/components/how-it-works/how-it-works.component';
import {MatIcon} from '@angular/material/icon';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {MembershipLevelInterface, TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {UserService} from '@app/core/user/user.service';
import {Router} from '@angular/router';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {toSignal} from '@angular/core/rxjs-interop';
import {MembershipLevelsEnum} from '@app/app.constants';
import {ReferralsService} from '@app/modules/client/referrals/services/referrals.service';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-referrals',
	standalone: true,
	imports: [
		CommonModule,
		TranslocoPipe,
		YourReferralsComponent,
		HowItWorksComponent,
		MatIcon,
		CdkCopyToClipboard,
	],
	templateUrl: './referrals.component.html',
	styleUrl: './referrals.component.scss',
})
export class ReferralsComponent implements OnInit {
	private _translocoService = inject(TranslocoService);
	private _toasterService = inject(ToasterService);
	private _userService = inject(UserService);
	private router = inject(Router);
	private queryService: UrlQueryService = inject(UrlQueryService);
	private referralsService: ReferralsService = inject(ReferralsService);
	private _screenService: ScreenService = inject(ScreenService);

	protected readonly ReferralsTabEnum = ReferralsTabEnum;

	$user: WritableSignal<TagpeakUser> = signal(null);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	//Tabs name
	tabs: ReferralsTabEnum[] = [ReferralsTabEnum.YOUR_REFERRALS, ReferralsTabEnum.HOW_IT_WORKS];

	$referralLink: Signal<string> = computed(() => {
		const origin = new URL(window.location.href).origin;
		const internalRoute = this.router.createUrlTree(['sign-up', this.$user().referralCode]);
		return origin + '/#' + this.router.serializeUrl(internalRoute);
	});

	$selectedTab: WritableSignal<ReferralsTabEnum> = signal(this.tabs[0]);

	ngOnInit(): void {
		this._userService.user$.subscribe((user) => {
			this.$user.set(user);
		});

		if (this.queryService.getParam('tab')) {
			this.$selectedTab.set(this.queryService.getParam('tab') as ReferralsTabEnum);
		} else {
			this.queryService.addParam('tab', this.tabs[0]);
		}
	}

	getNameTabs(tab: ReferralsTabEnum): string {
		return this._translocoService.translate(`referrals.tabs.${tab}`);
	}

	selectTab(tab: ReferralsTabEnum): void {
		this.$selectedTab.update((): ReferralsTabEnum => tab);
		this.queryService.addParam('tab', tab);
	}

	informCopy() {
		this._toasterService.showToast('info', 'Copied to clipboard');
	}

	$userLevel: Signal<MembershipLevelInterface> = computed(
		toSignal(this._userService.getCurrentMembershipLevel(this.$user()?.uuid)),
	);

	$goalValue: Signal<number> = computed(
		() => this.referralsService.getInfoBaseOnMembershipLevel(this.$userLevel()?.level)?.maxValue,
	);

	getNextMembershipLevel(level: string): string {
		if (level === MembershipLevelsEnum.base) {
			return MembershipLevelsEnum.silver;
		}
		if (level === MembershipLevelsEnum.silver) {
			return MembershipLevelsEnum.gold;
		}

		return MembershipLevelsEnum.gold;
	}
}
