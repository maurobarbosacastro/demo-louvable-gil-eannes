import {Component, inject, OnInit, Signal, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {InfoCardComponent} from '@app/shared/components/info-card/info-card.component';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {AddPaymentMethodComponent} from '@app/modules/settings/withdrawals-settings/components/add-payment-method/add-payment-method.component';
import {AppConstants} from '@app/app.constants';
import {PaymentMethodsService} from '@app/modules/settings/withdrawals-settings/services/payment-methods.service';
import {UserPaymentMethodInterface} from '@app/modules/settings/withdrawals-settings/models/user-payment-method.interface';
import {filter, switchMap, tap} from 'rxjs';
import {PaymentMethodDetailsComponent} from '@app/modules/settings/withdrawals-settings/components/payment-method-details/payment-method-details.component';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {StatusContainerComponent} from '@app/shared/components/status-container/status-container.component';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {toSignal} from '@angular/core/rxjs-interop';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-withdrawal-settings',
	standalone: true,
	imports: [CommonModule, InfoCardComponent, MatIcon, TranslocoPipe, StatusContainerComponent],
	templateUrl: './withdrawal-settings.component.html',
	styleUrl: './withdrawal-settings.component.scss',
})
export class WithdrawalSettingsComponent implements OnInit {
	private readonly _matDialog: MatDialog = inject(MatDialog);
	private readonly _paymentMethodsService: PaymentMethodsService = inject(PaymentMethodsService);
	private readonly _toaster: ToasterService = inject(ToasterService);
	private readonly _translocoService: TranslocoService = inject(TranslocoService);
	private readonly _screenService: ScreenService = inject(ScreenService);

	$paymentMethods: WritableSignal<UserPaymentMethodInterface[]> = signal([]);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	ngOnInit(): void {
		this.getPaymentMethods();
	}

	getPaymentMethods(): void {
		this._paymentMethodsService
			.getPaymentMethods()
			.subscribe((paymentMethods: PaginationInterface<UserPaymentMethodInterface>) =>
				this.$paymentMethods.set(paymentMethods.data),
			);
	}

	formatIban(iban: string): string {
		let formattedValue: string = '';
		for (let i: number = 0; i < iban.length; i++) {
			if (i > 0 && i % 4 === 0) {
				formattedValue += ' ';
			}
			formattedValue += iban[i];
		}
		return formattedValue;
	}

	addPaymentMethod(): void {
		const config: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
		};

		if (this.$isMobile()) {
			config.width = '100%';
			config.height = '95%';
			config.panelClass = 'dialog-panel-details-mobile';
		}

		this._matDialog
			.open(AddPaymentMethodComponent, config)
			.afterClosed()
			.subscribe((result) => {
				if (!result.skipFetch) {
					this.getPaymentMethods();
				}
			});
	}

	openPaymentMethod(uuid: string): void {
		const config: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: {uuid: uuid},
		};

		if (this.$isMobile()) {
			config.width = '100%';
			config.height = '85%';
			config.panelClass = 'dialog-panel-details-mobile';
		}
		this._matDialog
			.open(PaymentMethodDetailsComponent, config)
			.afterClosed()
			.subscribe((result) => {
				if (!result.skipFetch) {
					this.getPaymentMethods();
				}
			});
	}

	deletePaymentMethod(uuid: string): void {
		const data = {
			title: 'withdrawal-settings.delete-payment',
			body: 'withdrawal-settings.delete-payment-body',
			showFAQ: true,
		};
		this._matDialog
			.open(DeleteModalConfirmationComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '26rem',
				data: data,
			})
			.afterClosed()
			.pipe(
				filter((result: {skipFetch: boolean}) => result && !result.skipFetch),
				switchMap(() => this._paymentMethodsService.deletePaymentMethod(uuid)),
			)
			.subscribe({
				next: () => {
					this.handleSuccessRequest('withdrawal-settings.success-delete');
				},
				error: (error) => {
					if (error) {
						this.handleErrorRequest('withdrawal-settings.error-delete');
					}
				},
			});
	}

	handleSuccessRequest(message: string): void {
		this._toaster.showToast('success', this._translocoService.translate(message));
		this.getPaymentMethods();
	}

	handleErrorRequest(message: string): void {
		this._toaster.showToast('error', this._translocoService.translate(message));
	}
}
