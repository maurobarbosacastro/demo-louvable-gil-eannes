import {Component, effect, input, output, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule} from '@angular/forms';
import {TranslocoPipe} from '@ngneat/transloco';
import {MatIcon} from '@angular/material/icon';

export interface CheckboxGroupInterface {
	key: string;
	name: string;
	selected: boolean;
}

@Component({
	selector: 'tagpeak-checkbox-group',
	standalone: true,
	imports: [CommonModule, FormsModule, TranslocoPipe, MatIcon],
	templateUrl: './checkbox-group.component.html',
})
export class CheckboxGroupComponent {
	$data = input.required<CheckboxGroupInterface[]>();
	$placeholder = input.required<string>();
	$error = input.required<boolean>();
	$onSelectedValues = output<string[]>();

	$isDropdownOpen: WritableSignal<boolean> = signal(false); // Tracks whether the dropdown is open
	$areAllSelected: WritableSignal<boolean> = signal(false);
	$filteredValues: WritableSignal<CheckboxGroupInterface[]> = signal([]); // Filtered list
	$selectedValues: WritableSignal<CheckboxGroupInterface[]> = signal([]); // Filtered list
	$hideSelectAll: WritableSignal<boolean> = signal(false);

	searchQuery: string = ''; // Search query

	constructor() {
		effect(
			() => {
				if (this.$data()) {
					this.$filteredValues.set([...this.$data()]);
					this.$selectedValues.set(...[this.$data().filter((f) => f.selected === true)]);
				}
			},
			{allowSignalWrites: true},
		);
	}

	toggleDropdown(): void {
		this.$isDropdownOpen.set(!this.$isDropdownOpen());
	}

	// Filters the list of values based on the search query
	filterValues(): void {
		this.$hideSelectAll.set(this.searchQuery.length !== 0);
		const query = this.searchQuery.toLowerCase();
		this.$filteredValues.set(this.$data().filter((val) => val.name.toLowerCase().includes(query)));
	}

	// Toggles the selection state of all values
	toggleSelectAll(checked: boolean): void {
		this.$filteredValues().forEach((val) => (val.selected = checked));
		this.$areAllSelected.set(checked);
		this.$selectedValues.set(this.$filteredValues().filter((v) => v.selected == true));
		this.$onSelectedValues.emit(this.$selectedValues().map((v) => v.key));
	}

	selectValue(val: CheckboxGroupInterface) {
		let value = this.$filteredValues().find((v) => v.key === val.key);
		if (value.selected) {
			value.selected = false;
			this.$selectedValues.set(this.$selectedValues().filter((v) => v.key !== val.key));
		} else {
			value.selected = true;
			this.$selectedValues.set(this.$selectedValues().concat(value));
		}
		this.$onSelectedValues.emit(this.$selectedValues().map((v) => v.key));
	}
}
