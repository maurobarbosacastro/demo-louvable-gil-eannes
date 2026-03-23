import {
	Component,
	EventEmitter,
	input,
	Input,
	InputSignal,
	Output,
	TemplateRef,
	ViewEncapsulation,
} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, ReactiveFormsModule} from '@angular/forms';
import {TranslocoPipe} from '@ngneat/transloco';
import {MatIcon} from '@angular/material/icon';
import {OverlayModule} from '@angular/cdk/overlay';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	IconInputDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {AppConstants} from '@app/app.constants';

@Component({
	selector: 'tagpeak-custom-dropdown',
	standalone: true,
	imports: [CommonModule, TranslocoPipe, ReactiveFormsModule, MatIcon, OverlayModule],
	templateUrl: './custom-dropdown.component.html',
	encapsulation: ViewEncapsulation.None,
})
export class CustomDropdownComponent {
	$options: InputSignal<OptionDropdownInterface[]> = input.required();
	@Input({required: false}) placeholder: string = 'Select an option...';
	@Input({required: true}) control: FormControl;
	@Input({required: true}) value: any;
	@Input({required: false}) iconInput: IconInputDropdownInterface =
		AppConstants.DEFAULT_ICON_DROPDOWN;
	@Input({required: false}) css: CssDropdownInterface;
	@Input({required: false}) iconCheckOption: IconDropdownInterface;
	@Input({required: false}) showDefaultValue: boolean = true;
	@Input({required: false}) iconResetValue: boolean = true;
	@Input({required: false}) isDisabled: boolean = false;
	@Input({required: false}) isReadonly: boolean = true;

	@Input({required: false}) customListTemplate: TemplateRef<any>;

	@Output() valueChangesEvent: EventEmitter<any> = new EventEmitter<any>();

	isOpen: boolean = false;
	wasClean: boolean = false;

	setValue(value: any): void {
		this.valueChangesEvent.emit(value);
		if (value === null) {
			this.wasClean = true;
		}
		this.isOpen = false;
	}

	toggleDropdown(event: Event): void {
		event.stopPropagation();
		if (!this.isDisabled) {
			if (this.wasClean && this.isOpen) {
				this.isOpen = true;
				this.wasClean = false;
			} else {
				this.isOpen = !this.isOpen;
			}
		}
	}

	validateOptionCheck(option: OptionDropdownInterface): boolean {
		return (
			this.value?.toUpperCase() === option?.value?.toUpperCase() ||
			this.value?.toUpperCase() === option?.label?.toUpperCase()
		);
	}
}
