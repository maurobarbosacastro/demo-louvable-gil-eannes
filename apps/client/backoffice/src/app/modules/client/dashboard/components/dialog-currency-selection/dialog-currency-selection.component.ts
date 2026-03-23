import {Component, effect, inject, Signal} from '@angular/core';
import {UserService} from '@app/core/user/user.service';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe} from '@ngneat/transloco';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {CssDropdownInterface, IconDropdownInterface, OptionDropdownInterface} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {AppConstants} from '@app/app.constants';
import {FormControl, Validators} from '@angular/forms';
import {toSignal} from '@angular/core/rxjs-interop';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {MatDialogRef} from '@angular/material/dialog';

@Component({
	selector: 'tagpeak-dialog-currency-selection',
	templateUrl: 'dialog-currency-selection.component.html',
	imports: [MatIcon, TranslocoPipe, CustomDropdownComponent],
	standalone: true,
})
export class DialogCurrencySelectionComponent {
	userService = inject(UserService);
	toastService = inject(ToasterService)
	dialogRef = inject(MatDialogRef);

	currencies: OptionDropdownInterface[] = AppConstants.CURRENCIES.map((currency) => ({
		label: currency.label + ' (' + currency.symbol + ')',
		value: currency.key,
		iconName: currency.icon,
	}));

	defaultCurrency = AppConstants.CURRENCIES.find( c => c.key === 'EUR' )
	defaultOption: OptionDropdownInterface = {
		label: this.defaultCurrency.label + ' (' + this.defaultCurrency.symbol + ')',
		value: this.defaultCurrency.key,
		iconName: this.defaultCurrency.icon,
	}

	formControl: FormControl = new FormControl(this.defaultOption, [Validators.required]);

	styleDropdown: CssDropdownInterface = {
		...AppConstants.DEFAULT_STYLE_DROPDOWN,
		styleHeader: 'bg-white border-b border-project-brand-500 pb-2',
		styleInput: 'text-project-licorice-500',
	};
	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;

	setCurrencyValue(value: OptionDropdownInterface): void {
		this.formControl.setValue(value);
	}

	getCurrencyValue(value: OptionDropdownInterface): string {
		if (!value) return null;

		const currency: OptionDropdownInterface = this.currencies.find(
			(currency: OptionDropdownInterface): boolean => currency.value === value.value,
		);
		return currency.label;
	}

	save() {
		this.userService.setCurrency(this.userService.$user().uuid, this.formControl.getRawValue().value)
			.subscribe( _ => {
				this.toastService.showToast("success", "Currency saved", "top");
				this.dialogRef.close()
			}, _ => {
				this.toastService.showToast("error", "Something wrong happened, contact support", "top");
				this.dialogRef.close()
			})
	}
}
