import {Injectable} from '@angular/core';
import {FuseMockApiService} from '@fuse/lib/mock-api';
import {cloneDeep} from 'lodash-es';
import {cashbacks, cashbacksDetails} from '@app/mock-api/common/cashback/data';
import {
	ManageCashbackDetails,
	ManageCashbackInterface,
} from '@app/modules/admin/cashback/models/manage-cashback.interface';

@Injectable({providedIn: 'root'})
export class CashbackMockApi {
	private _cashbacks: ManageCashbackInterface[] = cashbacks;
	private _cashbackUnique: ManageCashbackDetails[] = cashbacksDetails;

	/**
	 * Constructor
	 */
	constructor(private _fuseMockApiService: FuseMockApiService) {
		// Register Mock API handlers
		this.registerHandlers();
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Register Mock API handlers
	 */
	registerHandlers(): void {
		// -----------------------------------------------------------------------------------------------------
		// @ cashback - GET
		// -----------------------------------------------------------------------------------------------------
		this._fuseMockApiService.onGet('api/ms-tagpeak/cashbacks').reply(({request}) => {
			const status: string[] = [request.params.get('status')];
			const statusFormatted: string[] = status?.[0]?.split(',');

			// Clone the cashbacks
			let cashbacks: ManageCashbackInterface[] = cloneDeep(this._cashbacks);

			if (statusFormatted?.length > 0) {
				cashbacks = cashbacks.filter((cashback: any) => statusFormatted.includes(cashback.status));
			}

			const totalSize: number = cashbacks.length;

			return [200, {cashbacks: cashbacks, totalSize: totalSize}];
		});
		// -----------------------------------------------------------------------------------------------------
		// @ cashbackById - GET
		// -----------------------------------------------------------------------------------------------------
		this._fuseMockApiService.onGet('api/ms-tagpeak/cashback').reply(({request}) => {
			const id: string = request.params.get('id');

			const cashback: ManageCashbackDetails = cloneDeep(
				this._cashbackUnique.find((cashbackUnique) => cashbackUnique.id === id),
			);

			return [200, cashback];
		});
		// -----------------------------------------------------------------------------------------------------
		// @ updateCashback - PATCH
		// -----------------------------------------------------------------------------------------------------
		this._fuseMockApiService.onPatch('api/ms-tagpeak/cashback').reply(({request}) => {
			const id: string = request.params.get('id');

			const cashbackDetails: ManageCashbackDetails = this._cashbackUnique.find(
				(cashbackUnique) => cashbackUnique.id === id,
			);

			cashbackDetails.refId = request.body.exitClick;
			cashbackDetails.currency = request.body.currency;
			cashbackDetails.priceDayZero = Number(request.body.priceDayZero);
			cashbackDetails.initialDate = request.body.initialDate;
			cashbackDetails.endDate = request.body.endDate;
			cashbackDetails.isin = request.body.isin;
			cashbackDetails.title = request.body.title;
			cashbackDetails.details = request.body.details;
			cashbackDetails.orderValue = request.body.orderValue;
			cashbackDetails.networkCommission = request.body.networkCommission;
			cashbackDetails.cashback = request.body.cashback;
			cashbackDetails.orderDate = request.body.orderDate;
			cashbackDetails.status = request.body.status;
			cashbackDetails.overridePrice = request.body.overridePrice;

			const cashback: ManageCashbackInterface = this._cashbacks.find(
				(cashbackUnique) => cashbackUnique.id === id,
			);
			cashback.status = request.body.status;
			cashback.networkCommission = request.body.networkCommission;
			cashback.isin = request.body.isin;
			cashback.title = request.body.title;
			cashback.orderValue = request.body.orderValue;
			cashback.priceDayZero = request.body.priceDayZero;
			cashback.date = request.body.orderDate;

			return [200, cashbackDetails];
		});
		// -----------------------------------------------------------------------------------------------------
		// @ updateCashback - DELETE
		// -----------------------------------------------------------------------------------------------------
		this._fuseMockApiService.onDelete('api/ms-tagpeak/cashback').reply(({request}) => {
			const id: string = request.params.get('id');

			this._cashbacks = this._cashbacks.filter((cashbackUnique) => cashbackUnique.id !== id);

			return [200, this._cashbacks];
		});

		// -----------------------------------------------------------------------------------------------------
		// @ cashbackHistory - GET
		// -----------------------------------------------------------------------------------------------------
		this._fuseMockApiService.onPatch(`api/ms-tagpeak/cashback/bulk`).reply(({request}) => {
			console.log(request);

			request.body.ids.forEach((id: string) => {
				const cashbackDetails: ManageCashbackDetails = this._cashbackUnique.find(
					(cashbackUnique) => cashbackUnique.id === id,
				);

				cashbackDetails.priceDayZero = request.body.priceDayZero
					? Number(request.body.priceDayZero)
					: cashbackDetails.priceDayZero;
				cashbackDetails.initialDate = request.body.initialDate
					? request.body.initialDate
					: cashbackDetails.initialDate;
				cashbackDetails.endDate = request.body.endDate
					? request.body.endDate
					: cashbackDetails.endDate;
				cashbackDetails.isin = request.body.isin ? request.body.isin : cashbackDetails.isin;
				cashbackDetails.status = request.body.status ? request.body.status : cashbackDetails.status;
				cashbackDetails.overridePrice = request.body.overridePrice
					? request.body.overridePrice
					: cashbackDetails.overridePrice;

				const cashback: ManageCashbackInterface = this._cashbacks.find(
					(cashbackUnique) => cashbackUnique.id === id,
				);
				cashback.status = request.body.status ? request.body.status : cashback.status;
				cashback.isin = request.body.isin ? request.body.isin : cashback.isin;
				cashback.priceDayZero = request.body.priceDayZero
					? request.body.priceDayZero
					: cashback.priceDayZero;
			});

			return [200, []];
		});
	}
}
