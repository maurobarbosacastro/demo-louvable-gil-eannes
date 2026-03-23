import {
	Component,
	DestroyRef,
	inject,
	Inject,
	OnInit,
	signal,
	Signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule, CurrencyPipe} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {LineChartComponent} from '@app/shared/components/line-chart/line-chart.component';
import {StatusContainerComponent} from '@app/shared/components/status-container/status-container.component';
import {ChartOptions} from '@app/shared/components/line-chart/models/line-chart';
import {StopConfirmationComponent} from '@app/modules/client/rewards/components/stop-confirmation/stop-confirmation.component';
import {filter, switchMap, tap} from 'rxjs';
import {ModalActionResult} from '@app/utils/modal-action-result.enum';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {differenceInCalendarDays, differenceInDays, format, parse} from 'date-fns';
import {StatusEnum} from '@app/utils/status.enum';
import {CashbackService} from '@app/modules/admin/cashback/services/cashback.service';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {
	CashbackHistory,
	CashbackResponse,
	RewardResponse,
} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {UserService} from '@app/core/user/user.service';
import {RewardService} from '@app/modules/admin/cashback/services/reward.service';
import {Router} from '@angular/router';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-reward-details',
	standalone: true,
	imports: [CommonModule, MatIcon, TranslocoPipe, LineChartComponent, StatusContainerComponent],
	templateUrl: './reward-details.component.html',
	providers: [CurrencyPipe],
})
export class RewardDetailsComponent implements OnInit {
	today = new Date();
	$user: Signal<TagpeakUser> = toSignal(this.userService.get());
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	protected readonly StatusEnum = StatusEnum;
	protected readonly differenceInDays = differenceInDays;
	private readonly router: Router = inject(Router);

	constructor(
		private readonly myDialog: MatDialogRef<RewardDetailsComponent>,
		@Inject(MAT_DIALOG_DATA)
		private readonly receivedData: string,
		private readonly rewardService: RewardService,
		private readonly userService: UserService,
		private readonly cashbackService: CashbackService,
		private readonly dialog: MatDialog,
		private readonly toasterService: ToasterService,
		private readonly translocoService: TranslocoService,
		private destroyRef: DestroyRef,
		private readonly currencyPipe: CurrencyPipe,
		private readonly _screenService: ScreenService,
	) {}

	$transaction: Signal<CashbackResponse> = toSignal(
		this.cashbackService
			.getCashbackById(this.receivedData)
			.pipe(takeUntilDestroyed(this.destroyRef)),
	);

	$reward: WritableSignal<RewardResponse> = signal(null);
	$chartOptions: WritableSignal<Partial<ChartOptions>> = signal(null);
	$rewardHistory: WritableSignal<CashbackHistory[]> = signal(null);

	ngOnInit(): void {
		this.cashbackService
			.getCashbackReward(this.receivedData)
			.pipe(
				tap((res: RewardResponse) => this.$reward.set(res)),
				switchMap((res: RewardResponse) => this.cashbackService.getCashbackHistory(res.uuid)),
				takeUntilDestroyed(this.destroyRef),
			)
			.subscribe((res: CashbackHistory[]) => {
				this.$rewardHistory.set(res);
				this.$chartOptions.set(this.setupChartOptions(res));
			});
	}

	closeModal(): void {
		this.myDialog.close();
	}

	setupChartOptions(data: CashbackHistory[]): Partial<ChartOptions> {
		return {
			series: [
				{
					name: 'Invested',
					data: this.getRewardHistoryData(data),
					color: '#4237DA',
				},
				{
					name: 'Minimum Guaranteed',
					data: this.getMinimumGuaranteedReward(data),
					color: '#32A852',
				},
			],
			chart: {
				width: '100%',
				type: 'line',
				height: 330,
				zoom: {
					enabled: false,
				},
				toolbar: {
					show: false,
				},
			},
			xAxis: {
				type: 'datetime',
				min: this.minDate,
				max: this.maxDate,
				labels: {
					rotate: 0,
					offsetX: 3,
					style: {
						colors: '#637381',
						fontSize: '11px',
					},
					formatter: (value, _, __) => format(value, 'dd.MM'),
				},
				axisTicks: {
					show: false,
				},
				tooltip: {
					enabled: false,
				},
			},
			yAxis: {
				opposite: true,
				min: 0,
				max: this.maxRewardValue + 1,
				tooltip: {
					enabled: false,
				},
				labels: {
					style: {
						colors: '#7A7B92',
						fontSize: '12px',
					},
				},
			},
			stroke: {
				curve: 'smooth',
				width: 2,
				colors: ['#4237DA'],
				show: true,
			},
			tooltip: {
				custom: ({series, seriesIndex, dataPointIndex, w}) => {
					// Current reward value
					const formattedValue: string = this.currencyPipe.transform(
						series[0][dataPointIndex],
						this.$user().currency,
						'symbol-narrow',
						'1.2-2',
					);
					// Guarantee value is the same as the current reward
					const formattedValueGuarantee: string = this.currencyPipe.transform(
						series[1][dataPointIndex],
						this.$user().currency,
						'symbol-narrow',
						'1.2-2',
					);
					const formattedDate: string = format(
						this.$chartOptions().series.at(seriesIndex).data.at(dataPointIndex)['x'],
						'dd MMMM',
					);
					return `
			            <div class="flex flex-col p-2.5 gap-2.5 leading-none">
			                <span class="text-[#4237DA] text-xl">${formattedValue}</span>
			                <span class="text-[#32A852] text-lg">${formattedValueGuarantee}</span>
			                <span class="text-project-waterloo text-sm">${formattedDate}</span>
			            </div>
			        `;
				},
			},
			dataLabels: {
				enabled: false,
			},
			grid: {
				padding: {
					bottom: 30, // Increase this until the cut-off disappears
				},
			},
		};
	}

	visitStore(): void {
		this.router.navigate(['/stores', this.$transaction().store.uuid]);
		this.myDialog.close();
	}

	stopReward(): void {
		const dialogRef = this.dialog.open(StopConfirmationComponent, {
			panelClass: 'dialog-panel',
			height: 'fit-content',
			width: '26rem',
		});

		dialogRef
			.afterClosed()
			.pipe(
				filter((result) => result.action === ModalActionResult.OK),
				switchMap(() =>
					this.rewardService.stopReward({id: this.$reward().uuid, status: StatusEnum.STOPPED}),
				),
			)
			.subscribe({
				next: () => {
					this.toasterService.showToast(
						'success',
						this.translocoService.translate('dashboard.stop-confirmation.success'),
						'top',
					);
				},
				error: (error) => {
					this.toasterService.showToast(
						'error',
						this.translocoService.translate('dashboard.stop-confirmation.error'),
					);
				},
			});
	}

	getRewardHistoryData(data: CashbackHistory[]): {x: number; y: number}[] {
		return (
			data?.map((item: CashbackHistory) => {
				return {
					x: new Date(item.createdAt).getTime(),
					y: parseFloat(Number(item.cashReward).toFixed(2)),
				};
			}) || []
		);
	}

	getMinimumGuaranteedReward(data: CashbackHistory[]): {x: number; y: number}[] {
		return (
			data?.map((item: CashbackHistory) => {
				return {
					x: new Date(item.createdAt).getTime(),
					y: parseFloat(this.$reward().minimumReward.toFixed(2)),
				};
			}) || []
		);
	}

	get maxRewardValue(): number {
		return Math.max(
			...(this.$rewardHistory()?.map((item) => parseFloat(Number(item.cashReward).toFixed(2))) || [
				0,
			]),
		);
	}

	protected readonly differenceInCalendarDays = differenceInCalendarDays;

	get minDate(): number {
		return this.$rewardHistory && this.$rewardHistory()?.length > 0
			? new Date(this.$rewardHistory()[0]?.createdAt).getTime()
			: new Date().getTime();
	}

	get maxDate(): number {
		return new Date(this.$rewardHistory().at(-1)?.createdAt).getTime();
	}
}
