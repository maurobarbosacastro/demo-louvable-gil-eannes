import {
	Component,
	computed,
	Inject,
	inject,
	Input,
	input,
	InputSignal,
	LOCALE_ID,
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
import {
	ReferralInfo,
	ReferralInterface,
	RevenueInfo,
	UserReferralInterface,
} from '@app/modules/client/referrals/models/referrals.interface';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {LineChartComponent} from '@app/shared/components/line-chart/line-chart.component';
import {ChartOptions} from '@app/shared/components/line-chart/models/line-chart';
import {Router} from '@angular/router';
import {MembershipLevelInfoInterface} from '@app/shared/interfaces/membership-level.interface';
import {ScreenService} from '@app/core/services/screen.service';
import {toSignal} from '@angular/core/rxjs-interop';

@Component({
	selector: 'tagpeak-your-referrals',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		TranslocoPipe,
		CdkCopyToClipboard,
		NgOptimizedImage,
		LineChartComponent,
	],
	templateUrl: './your-referrals.component.html',
})
export class YourReferralsComponent implements OnInit {
	constructor(@Inject(LOCALE_ID) private locale: string) {}

	$user = input<TagpeakUser>();
	@Input() getNextMembershipLevel: (level: string) => string;
	$userLevel: InputSignal<MembershipLevelInterface> = input.required<MembershipLevelInterface>();
	$goalValue: InputSignal<number> = input.required<number>();

	private _referralsService = inject(ReferralsService);
	private _toasterService = inject(ToasterService);
	private router = inject(Router);
	private _screenService: ScreenService = inject(ScreenService);

	$referrals: WritableSignal<ReferralInterface[]> = signal(null);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	$chartOptionsClicks: WritableSignal<Partial<ChartOptions>> = signal(null);
	$chartOptionsRegistered: WritableSignal<Partial<ChartOptions>> = signal(null);
	$chartOptionsPurchases: WritableSignal<Partial<ChartOptions>> = signal(null);
	$chartOptionsRevenue: WritableSignal<Partial<ChartOptions>> = signal(null);

	$users: WritableSignal<UserReferralInterface[]> = signal([]);
	$referralInfo: WritableSignal<ReferralInfo> = signal(null);
	$revenueInfo: WritableSignal<RevenueInfo> = signal(null);
	$totalRevenue: WritableSignal<number> = signal(0.0);
	$doubleValue: WritableSignal<number> = signal(0.0);

	page: number = 1;
	pageSize: number = 10;
	sort: string = 'created_at desc';

	dataClicks: number[] = [];
	xDataClicks: string[] = [];
	dataRegistered: number[] = [];
	xDataRegistered: string[] = [];
	dataPurchases: number[] = [];
	xDataPurchases: string[] = [];
	dataRevenue: number[] = [];
	xDataRevenue: string[] = [];

	progressMinValue: number = 10;
	progressMaxValue: number = 100;

	$referralLink: Signal<string> = computed(() => {
		const origin = new URL(window.location.href).origin;
		const internalRoute = this.router.createUrlTree(['sign-up', this.$user().referralCode]);
		return origin + '/#' + this.router.serializeUrl(internalRoute);
	});

	ngOnInit(): void {
		this._referralsService.getUsersRevenueInfo(this.$user().uuid).subscribe((res) => {
			this.$users.set(res);
		});

		this._referralsService.getReferralInfo(this.$user().uuid).subscribe((res) => {
			this.$referralInfo.set(res);
			this.setDataReferralInfoGraph();
		});

		this._referralsService.getRevenueInfo(this.$user().uuid).subscribe((res) => {
			this.$revenueInfo.set(res);
			this.setDataRevenueInfoGraph();
		});

		this._referralsService.getRevenueTotalValue(this.$user().uuid).subscribe((res) => {
			this.$totalRevenue.set(res);
			this.$doubleValue.set(res * 2);
		});

		this._referralsService
			.getAllReferralsByUserUuid(this.$user().uuid, this.page, this.pageSize, this.sort)
			.subscribe((referrals) => {
				this.$referrals.set(referrals.data);
			});
	}

	setDataReferralInfoGraph() {
		this.$referralInfo().clicksByMonth.forEach((value) => {
			this.xDataClicks.push(value.month);
			this.dataClicks.push(value.value);
		});
		this.xDataClicks.reverse();
		this.dataClicks.reverse();

		this.$referralInfo().registeredByMonth.forEach((value) => {
			this.xDataRegistered.push(value.month);
			this.dataRegistered.push(value.value);
		});
		this.xDataRegistered.reverse();
		this.dataRegistered.reverse();

		this.$referralInfo().firstPurchaseByMonth.forEach((value) => {
			this.xDataPurchases.push(value.month);
			this.dataPurchases.push(value.value);
		});

		this.xDataPurchases.reverse();
		this.dataPurchases.reverse();

		this.$chartOptionsClicks.set(
			this.setupChartOptions(
				this.dataClicks,
				this.xDataClicks,
				this.minReferralClicksValue,
				this.maxReferralClicksValue,
			),
		);

		this.$chartOptionsRegistered.set(
			this.setupChartOptions(
				this.dataRegistered,
				this.xDataRegistered,
				this.minReferralRegisteredValue,
				this.maxReferralRegisteredValue,
			),
		);
		this.$chartOptionsPurchases.set(
			this.setupChartOptions(
				this.dataPurchases,
				this.xDataPurchases,
				this.minReferralFirstPurchaseValue,
				this.maxReferralFirstPurchaseValue,
			),
		);
	}

	setDataRevenueInfoGraph() {
		this.$revenueInfo().revenueByMonth.forEach((value) => {
			this.xDataRevenue.push(value.month);
			this.dataRevenue.push(value.value);
		});

		this.xDataRevenue.reverse();
		this.dataRevenue.reverse();

		this.$chartOptionsRevenue.set(
			this.setupChartOptions(
				this.dataRevenue,
				this.xDataRevenue,
				this.minRevenueValue,
				this.maxRevenueValue,
				true,
				'#4237DA',
				350,
			),
		);
	}

	informCopy() {
		this._toasterService.showToast('info', 'Copied to clipboard');
	}

	setupChartOptions(
		data: number[],
		xData: string[],
		min: number,
		max: number,
		showYLabel: boolean = false,
		lineColor: string = '#FDCA65',
		height: number = 200,
	): Partial<ChartOptions> {
		return {
			series: [
				{
					name: '',
					data: data,
					color: lineColor,
				},
			],
			chart: {
				width: '100%',
				type: 'line',
				height: height,
				zoom: {
					enabled: false,
				},
				toolbar: {
					show: false,
				},
			},
			xAxis: {
				type: 'category',
				categories: xData,
				labels: {
					offsetX: 3,
					style: {
						colors: '#637381',
						fontSize: '11px',
					},
				},
				tooltip: {
					enabled: false,
				},
			},
			yAxis: {
				opposite: true,
				show: showYLabel,
				min: min > 0 ? min - 1 : 0,
				max: max + 1,
				tooltip: {
					enabled: false,
				},
			},
			stroke: {
				curve: 'straight',
				width: 2,
				colors: ['#4237DA'],
			},
			tooltip: {
				custom: ({series, seriesIndex, dataPointIndex, w}) => {
					return `
			            <div class="flex flex-col p-2.5 gap-2.5 leading-none">
			                <span class="text-project-licorice-900 text-xl">${series[seriesIndex][dataPointIndex]}</span>
			                <span class="text-project-waterloo text-sm">${w.globals.categoryLabels[dataPointIndex]}</span>
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

	get minReferralClicksValue(): number {
		return Math.min(
			...(this.$referralInfo()?.clicksByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	get maxReferralClicksValue(): number {
		return Math.max(
			...(this.$referralInfo()?.clicksByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	get minReferralRegisteredValue(): number {
		return Math.min(
			...(this.$referralInfo()?.registeredByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	get maxReferralRegisteredValue(): number {
		return Math.max(
			...(this.$referralInfo()?.registeredByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	get minReferralFirstPurchaseValue(): number {
		return Math.min(
			...(this.$referralInfo()?.firstPurchaseByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	get maxReferralFirstPurchaseValue(): number {
		return Math.max(
			...(this.$referralInfo()?.firstPurchaseByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	get minRevenueValue(): number {
		return Math.min(
			...(this.$revenueInfo()?.revenueByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	get maxRevenueValue(): number {
		return Math.max(
			...(this.$revenueInfo()?.revenueByMonth.map((item) =>
				parseFloat(Number(item.value).toFixed(2)),
			) || [0]),
		);
	}

	getCurrencySymbol(currencyCode: string): string {
		return (
			new Intl.NumberFormat(this.locale, {style: 'currency', currency: currencyCode})
				.formatToParts(0)
				.find((part) => part.type === 'currency')?.value || ''
		);
	}

	howItWorks(): void {
		this.router.navigate(['client/referrals'], {
			queryParams: {tab: 'how-it-works'},
		});
	}

	getInfoBaseOnMembershipLevel(level: string): Partial<MembershipLevelInfoInterface> {
		return this._referralsService.getInfoBaseOnMembershipLevel(level);
	}
}
