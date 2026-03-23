import {
	Component,
	computed,
	DestroyRef,
	inject,
	OnInit,
	Signal,
	signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule, NgOptimizedImage} from '@angular/common';
import {TableComponent} from '@app/shared/components/table/table.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MatDialog} from '@angular/material/dialog';
import {RewardDetailsComponent} from '@app/modules/client/rewards/components/reward-details/reward-details.component';
import {StatusContainerComponent} from '@app/shared/components/status-container/status-container.component';
import {StopConfirmationComponent} from '@app/modules/client/rewards/components/stop-confirmation/stop-confirmation.component';
import {BehaviorSubject, combineLatest, filter, map, switchMap, tap} from 'rxjs';
import {ModalActionResult} from '@app/utils/modal-action-result.enum';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {InfoCardComponent} from '@app/shared/components/info-card/info-card.component';
import {MatButton} from '@angular/material/button';
import {StatusEnum} from '@app/utils/status.enum';
import {membershipLevel, statusFilters} from '@app/utils/constants';
import {CashbackService} from '@app/modules/admin/cashback/services/cashback.service';
import {
	CashbackSortingFields,
	CashbackTableResponse,
	RewardOrigins,
} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {UserService} from '@app/core/user/user.service';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {MembershipLevelInterface, TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {showHistoryBtn} from '@app/modules/admin/cashback/utils/cashback.utils';
import {differenceInCalendarDays} from 'date-fns';
import {PageEvent} from '@angular/material/paginator';
import {AppConstants, MembershipLevelsEnum} from '@app/app.constants';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {DashboardTabEnum} from '@app/modules/client/dashboard/models/dashboard-tab.enum';
import {Router} from '@angular/router';
import {RewardService} from '@app/modules/admin/cashback/services/reward.service';
import {ConfigurationService} from '@app/modules/admin/configuration/services/configuration.service';
import {MembershipLevelInfoInterface} from '@app/shared/interfaces/membership-level.interface';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-your-rewards',
	standalone: true,
	imports: [
		CommonModule,
		TableComponent,
		MatIcon,
		TranslocoPipe,
		StatusContainerComponent,
		ReactiveFormsModule,
		InfoCardComponent,
		MatButton,
		NgOptimizedImage,
		CdkCopyToClipboard,
	],
	templateUrl: './your-rewards.component.html',
	styleUrls: ['your-rewards.component.scss', '../../../../../../styles/style-table.scss'],
})
export class YourRewardsComponent implements OnInit {
	$user: WritableSignal<TagpeakUser> = signal(null);

	private _router: Router = inject(Router);
	private _configService: ConfigurationService = inject(ConfigurationService);
	private _screenService: ScreenService = inject(ScreenService);

	page$: BehaviorSubject<number> = new BehaviorSubject(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject('date desc');
	filters$: BehaviorSubject<string[]> = new BehaviorSubject([]);
	pageSize: number = 100; //AppConstants.PAGE_SIZE;
	totalSize: number = 1;
	page: number;
	totalPages: number;
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	$referralLink: Signal<string> = computed(() => {
		const origin = new URL(window.location.href).origin;
		const internalRoute = this._router.createUrlTree(['sign-up', this.$user()?.referralCode]);
		return origin + '/#' + this._router.serializeUrl(internalRoute);
	});

	$dataNew: Signal<CashbackTableResponse[]> = toSignal(
		combineLatest([this.sort$, this.page$, this.filters$]).pipe(
			takeUntilDestroyed(this.destroyRef),
			switchMap(([sort, page, filters]) => {
				return this.cashbackService.getMyCashback(page - 1, this.pageSize, sort, {
					state: filters.join(','),
				});
			}),
			tap((res) => {
				this.page = res.page;
				this.totalPages = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			map((res) => res.data ?? []),
			map((res) =>
				res.map(
					(c) =>
						({
							...c,
							reward: Object.keys(c.reward).length > 0 ? c.reward : null,
							unvalidatedCurrentReward:
								c.store && c.store.percentageCashout
									? c.amountUser * (c.store.percentageCashout / 100)
									: 0,
						}) as CashbackTableResponse,
				),
			),
		),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => {
		return {
			dataSource: new MatTableDataSource(this.$dataNew()),
			pageSize: this.pageSize,
			styles: {
				header: 'font-nunitoBlack text-project-waterloo whitespace-nowrap',
				content: 'font-nunitoRegular text-project-waterloo',
				paginator: 'font-nunitoRegular text-black',
			},
			columns: [
				{
					id: 'store',
					name: this.transloco.translate('dashboard.table.store'),
					hasSort: true,
					hasTooltip: false,
					colWidth: '30%',
				},
				{
					id: 'amountUser',
					name: this.transloco.translate('dashboard.table.price-paid'),
					hasSort: true,
					hasTooltip: false,
					colWidth: '13%',
				},
				{
					id: 'stopDate',
					name: this.transloco.translate('dashboard.table.stop-date'),
					hasSort: true,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'currentReward',
					name: this.transloco.translate('dashboard.table.current-reward'),
					hasSort: true,
					hasTooltip: false,
					colWidth: '13%',
				},
				{
					id: 'status',
					name: this.transloco.translate('dashboard.table.status'),
					hasSort: true,
					hasTooltip: false,
					valueTransform: (row) => {
						if (row.reward) {
							return row.reward.state;
						}
						return row.status;
					},
					colWidth: '20%',
				},
				// The actions column has this name so it can be customized in the bodyCustomTemplate
				{
					id: 'customActions',
					name: '',
					hasSort: false,
					hasTooltip: false,
					colWidth: '4%',
				},
			],
		};
	});

	$userLevel: Signal<MembershipLevelInterface> = computed(
		toSignal(this.userService.getCurrentMembershipLevel(this.$user()?.uuid)),
	);

	$goalValue: Signal<number> = computed(
		() => this.getInfoBaseOnMembershipLevel(this.$userLevel()?.level)?.maxValue,
	);

	$liveRewards: WritableSignal<number> = signal(0);

	getInfoBaseOnMembershipLevel(level: string) {
		return this.setupMembershipLevels(level)?.find((item) => item.level === level);
	}

	setupMembershipLevels(level: string) {
		if (level === MembershipLevelsEnum.base) {
			return membershipLevel.map((item: MembershipLevelInfoInterface) => ({
				...item,
				percentageOnReward: this._configService.configs['referral_member_friend_cash_reward'].value,
				maxValue: Number(this._configService.configs['referral_silver_status_goal_amount'].value),
			}));
		} else if (level === MembershipLevelsEnum.silver) {
			return membershipLevel.map((item: MembershipLevelInfoInterface) => ({
				...item,
				percentageOnTransaction:
					this._configService.configs['referral_silver_transaction_cash_reward'].value,
				percentageOnReward: this._configService.configs['referral_gold_status_goal'].value,
				maxValue: Number(this._configService.configs['referral_gold_status_goal_amount'].value),
			}));
		} else {
			return membershipLevel.map((item: MembershipLevelInfoInterface) => ({
				...item,
				percentageOnTransaction:
					this._configService.configs['referral_gold_transaction_cash_reward'].value,
				percentageOnReward: this._configService.configs['referral_gold_friend_reward_share'].value,
				maxValue: Number(this._configService.configs['referral_gold_status_goal_amount'].value),
			}));
		}
	}

	getNextMembershipLevel(level: string): string {
		if (level === MembershipLevelsEnum.base) {
			return MembershipLevelsEnum.silver;
		}
		if (level === MembershipLevelsEnum.silver) {
			return MembershipLevelsEnum.gold;
		}

		return MembershipLevelsEnum.gold;
	}

	filters: {label: string; value: string}[] = statusFilters.sort((a, b) =>
		a.label.localeCompare(b.label),
	);

	filterForm = new FormGroup({
		filter: new FormControl<string[]>([]),
	});

	showHistoryBtn = showHistoryBtn;
	today = new Date();
	protected readonly StatusEnum = StatusEnum;

	constructor(
		private readonly rewardService: RewardService,
		private readonly cashbackService: CashbackService,
		private readonly transloco: TranslocoService,
		private readonly dialog: MatDialog,
		private readonly toasterService: ToasterService,
		private readonly userService: UserService,
		private destroyRef: DestroyRef,
		private readonly urlQueryService: UrlQueryService,
	) {}

	ngOnInit(): void {
		this.userService.get().subscribe((user) => {
			this.$user.set(user);
			this.rewardService
				.getLiveRewardsValue(user.uuid)
				.subscribe((res) => this.$liveRewards.set(res));
		});
	}

	openDetailsDialog(element: CashbackTableResponse): void {
		const width: string = this.$isMobile() ? '100%' : '60%';
		const panelCss: string = this.$isMobile()
			? 'dialog-panel-details-mobile'
			: 'dialog-panel-details';
		const dialogRef = this.dialog.open(RewardDetailsComponent, {
			panelClass: panelCss,
			width: width,
			data: element.uuid,
			disableClose: true,
		});

		dialogRef.afterClosed().subscribe(() => {
			this.page$.next(1);
		});
	}

	stopReward(element: CashbackTableResponse): void {
		const dialogRef = this.dialog.open(StopConfirmationComponent, {
			panelClass: 'dialog-panel',
			height: 'fit-content',
			width: '26rem',
			disableClose: true,
		});

		dialogRef
			.afterClosed()
			.pipe(
				filter((result) => result.action === ModalActionResult.OK),
				switchMap(() =>
					this.rewardService.stopReward({id: element.reward.uuid, status: StatusEnum.STOPPED}),
				),
			)
			.subscribe({
				next: () => {
					this.page$.next(1);
					this.toasterService.showToast(
						'success',
						this.transloco.translate('dashboard.stop-confirmation.success'),
						'top',
					);
				},
				error: (error) => {
					if (error) {
						this.toasterService.showToast(
							'error',
							this.transloco.translate('dashboard.stop-confirmation.error'),
						);
					}
				},
			});
	}

	setFilters(filter: string): void {
		if (this.filterForm.controls.filter.value.includes(filter)) {
			this.removeFilter(filter);
			return;
		}
		const filters: string[] = Array.from(
			new Set([...this.filterForm.controls.filter.value, filter]),
		);
		this.filterForm.controls.filter.patchValue(filters);
		this.filters$.next(filters);
	}

	removeFilter(filter: string): void {
		const filters: string[] = this.filterForm.controls.filter.value.filter(
			(value: string) => value !== filter,
		);
		this.filterForm.controls.filter.patchValue(filters);
		this.filters$.next(filters);
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			const field = CashbackSortingFields[evt.sort.split(',')[0]];
			const direction = evt.sort.split(',')[1];
			this.sort$.next(`${field} ${direction}`);
			return;
		}
		this.page$.next(evt.page.pageIndex);
	}

	goToWithdraw() {
		this.urlQueryService.addParam('tab', DashboardTabEnum.BALANCE);
	}

	goToStore(): void {
		this._router.navigate(['stores']);
	}

	howItWorks(): void {
		this._router.navigate(['client/referrals'], {
			queryParams: {tab: 'how-it-works'},
		});
	}

	informCopy() {
		this.toasterService.showToast('info', 'Copied to clipboard');
	}

	protected readonly MembershipLevelsEnum = MembershipLevelsEnum;
	protected readonly RewardOrigins = RewardOrigins;
	protected readonly differenceInCalendarDays = differenceInCalendarDays;

	isReferralOrCommission(data: CashbackTableResponse){
		if (data.reward.origin === RewardOrigins.REFERRAL) return true;
		if (data.reward.origin === RewardOrigins.COMMISSION) return true;

		//Handle duplicate issue that were marked as "duplicate" in bd.
		return data.reward.origin === RewardOrigins.DUPLICATE && data.reward.title.toLowerCase().includes('reward');
	}
}
