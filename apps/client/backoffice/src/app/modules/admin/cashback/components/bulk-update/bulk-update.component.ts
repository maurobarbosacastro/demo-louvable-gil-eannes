import {Component, EventEmitter, input, Input, InputSignal, Output} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {TranslocoPipe} from '@ngneat/transloco';
import {
	CashbackDto,
	CashbackTableResponse,
} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {statusFilters} from '@app/utils/constants';

@Component({
	selector: 'tagpeak-bulk-update',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		ReactiveFormsModule,
		TranslocoPipe,
		MatDatepicker,
		MatDatepickerInput,
	],
	templateUrl: './bulk-update.component.html',
	styleUrl: '../../../../../../styles/style-date-picker.scss',
})
export class BulkUpdateComponent {
	@Input() cashbacksSelected: CashbackTableResponse[];
	isMobile: InputSignal<boolean> = input.required<boolean>();

	@Output() updateClickEvent: EventEmitter<Partial<CashbackDto>> = new EventEmitter<
		Partial<CashbackDto>
	>();
	@Output() closeClickEvent: EventEmitter<void> = new EventEmitter<void>();

	@Input() lockButton: boolean = false;

	bulkForm = new FormGroup({
		priceDayZero: new FormControl(null),
		endDate: new FormControl(''),
		status: new FormControl(''),
		initialDate: new FormControl(''),
		isin: new FormControl(''),
		conid: new FormControl(''),
		overridePrice: new FormControl(null),
	});

	status: {label: string; value: string}[] = statusFilters.sort((a, b) =>
		a.label.localeCompare(b.label),
	);

	setFormControlValue(formControlName: string, value: string): void {
		this.bulkForm.get(formControlName).setValue(value);
	}

	getStatus(): string {
		return this.status.find((s) => s.value === this.bulkForm.value.status)?.label;
	}

	updateBulk(): void {
		if (this.lockButton) return;
		this.updateClickEvent.emit(this.bulkForm.value);
	}

	clearBulk(): void {
		this.closeClickEvent.emit();
	}

	getCashbacksSelectedAmount(): number {
		return this.cashbacksSelected.reduce(
			(acc: number, cur: CashbackTableResponse) => acc + cur.cashback,
			0,
		);
	}

	containStatusRequired(): boolean {
		return this.cashbacksSelected.some(
			(c: CashbackTableResponse) =>
				c.status === 'VALIDATED' || c.status === 'STOPPED' || c.status === 'FINISHED',
		);
	}
}
