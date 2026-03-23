import {Component, DestroyRef, inject, signal, Signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
import {MatIcon} from '@angular/material/icon';
import {UserPaymentMethodInterface} from '@app/modules/settings/withdrawals-settings/models/user-payment-method.interface';
import {takeUntilDestroyed, toObservable, toSignal} from '@angular/core/rxjs-interop';
import {PaymentMethodsService} from '@app/modules/settings/withdrawals-settings/services/payment-methods.service';
import {filter, map, switchMap} from 'rxjs';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-payment-method-details',
	standalone: true,
	imports: [CommonModule, MatIcon, TranslocoPipe],
	templateUrl: './payment-method-details.component.html',
	styleUrl: './payment-method-details.component.scss',
})
export class PaymentMethodDetailsComponent {
	private readonly _dialogData = inject(MAT_DIALOG_DATA);
	private readonly _myDialog = inject(
		MatDialogRef<PaymentMethodDetailsComponent, {skipFetch: boolean}>,
	);
	private readonly _paymentMethodsService: PaymentMethodsService = inject(PaymentMethodsService);
	private readonly _destroyRef = inject(DestroyRef);
	private readonly _toaster: ToasterService = inject(ToasterService);
	private readonly _translocoService: TranslocoService = inject(TranslocoService);
	private readonly _matDialog: MatDialog = inject(MatDialog);
	private readonly _screenService: ScreenService = inject(ScreenService);

	$havePayment: WritableSignal<boolean> = signal(this._dialogData && !!this._dialogData.uuid);
	$paymentMethod: Signal<UserPaymentMethodInterface> = toSignal(
		toObservable(this.$havePayment).pipe(
			takeUntilDestroyed(this._destroyRef),
			filter((val: boolean) => !!val),
			switchMap(() => this._paymentMethodsService.getPaymentMethodById(this._dialogData.uuid)),
			map((paymentMethod: UserPaymentMethodInterface) => paymentMethod),
		),
	);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	ibanFormat(iban: string): string {
		let formattedValue: string = '';
		for (let i: number = 0; i < iban?.length; i++) {
			if (i > 0 && i % 4 === 0) {
				formattedValue += ' ';
			}
			formattedValue += iban[i];
		}
		return formattedValue;
	}

	closeModal(): void {
		this._myDialog.close({skipFetch: true});
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

	downloadFile(): void {
		this._paymentMethodsService
			.downloadFile(this.$paymentMethod().ibanStatement.uuid)
			.subscribe((file: File) => {
				window.open(URL.createObjectURL(file));
			});
	}

	handleSuccessRequest(message: string): void {
		this._toaster.showToast('success', this._translocoService.translate(message));
		this._myDialog.close({skipFetch: false});
	}

	handleErrorRequest(message: string): void {
		this._toaster.showToast('error', this._translocoService.translate(message));
	}
}
