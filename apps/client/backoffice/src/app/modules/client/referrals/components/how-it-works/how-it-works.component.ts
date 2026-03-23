import {
	Component,
	computed,
	inject,
	input,
	Input,
	InputSignal,
	OnInit,
	Signal,
	signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule, NgOptimizedImage} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe} from '@ngneat/transloco';
import {ReferralsService} from '@app/modules/client/referrals/services/referrals.service';
import {MembershipLevelInterface, TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {ConfigurationService} from '@app/modules/admin/configuration/services/configuration.service';
import {MembershipLevelInfoInterface} from '@app/shared/interfaces/membership-level.interface';
import {ScreenService} from '@app/core/services/screen.service';
import {toSignal} from '@angular/core/rxjs-interop';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {Router} from '@angular/router';

@Component({
	selector: 'tagpeak-how-it-works',
	standalone: true,
	imports: [CommonModule, MatIcon, NgOptimizedImage, TranslocoPipe, CdkCopyToClipboard],
	templateUrl: './how-it-works.component.html',
})
export class HowItWorksComponent implements OnInit {
	private readonly _screenService: ScreenService = inject(ScreenService);
	private readonly _toasterService: ToasterService = inject(ToasterService);
	private router: Router = inject(Router);

	@Input() $user: WritableSignal<TagpeakUser> = signal(null);
	@Input() getNextMembershipLevel: (level: string) => string;
	$userLevel: InputSignal<MembershipLevelInterface> = input.required<MembershipLevelInterface>();
	$goalValue: InputSignal<number> = input.required<number>();

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	_referralsService = inject(ReferralsService);
	_configService = inject(ConfigurationService);

	$totalRevenue: WritableSignal<number> = signal(0.0);
	$doubleValue: WritableSignal<number> = signal(0.0);

	$referralLink: Signal<string> = computed(() => {
		const origin = new URL(window.location.href).origin;
		const internalRoute = this.router.createUrlTree(['sign-up', this.$user().referralCode]);
		return origin + '/#' + this.router.serializeUrl(internalRoute);
	});

	silverGoal: string = '';
	silverGoalAmount: string = '';
	goldGoal: string = '';
	goldGoalAmount: string = '';
	memberReward: string = '';
	silverPercentage: string = '';
	silverFriendsPercentage: string = '';
	goldPercentage: string = '';
	goldFriendsPercentage: string = '';

	ngOnInit(): void {
		this._referralsService.getRevenueTotalValue(this.$user().uuid).subscribe((res) => {
			this.$totalRevenue.set(res);
			this.$doubleValue.set(res * 2);
		});
		this.setConfigs();
	}

	setConfigs() {
		this.silverGoal = this._configService.configs['referral_silver_status_goal'].value;
		this.silverGoalAmount = this._configService.configs['referral_silver_status_goal_amount'].value;
		this.goldGoal = this._configService.configs['referral_gold_status_goal'].value;
		this.goldGoalAmount = this._configService.configs['referral_gold_status_goal_amount'].value;
		this.memberReward = this._configService.configs['referral_member_friend_cash_reward'].value;
		this.silverPercentage =
			this._configService.configs['referral_silver_transaction_cash_reward'].value;
		this.silverFriendsPercentage =
			this._configService.configs['referral_silver_friend_reward_share'].value;
		this.goldPercentage =
			this._configService.configs['referral_gold_transaction_cash_reward'].value;
		this.goldFriendsPercentage =
			this._configService.configs['referral_gold_friend_reward_share'].value;
	}

	getInfoBaseOnMembershipLevel(level: string): Partial<MembershipLevelInfoInterface> {
		return this._referralsService.getInfoBaseOnMembershipLevel(level);
	}

	informCopy(): void {
		this._toasterService.showToast('info', 'Copied to clipboard');
	}
}
