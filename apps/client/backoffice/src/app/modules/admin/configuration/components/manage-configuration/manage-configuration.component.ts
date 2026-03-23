import {Component, Inject} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {TranslocoPipe} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ConfigurationService} from '@app/modules/admin/configuration/services/configuration.service';
import {
	Configuration,
	PatchConfiguration,
} from '@app/modules/admin/configuration/interfaces/configuration.interface';
import _ from 'lodash';

@Component({
	selector: 'tagpeak-manage-configuration',
	standalone: true,
	imports: [CommonModule, MatIcon, ReactiveFormsModule, TranslocoPipe],
	templateUrl: './manage-configuration.component.html',
})
export class ManageConfigurationComponent {
	form: FormGroup = new FormGroup({
		name: new FormControl(this.receivedData.name),
		value: new FormControl(this.receivedData.value),
	});

	constructor(
		private readonly dialogRef: MatDialogRef<ManageConfigurationComponent, {skipFetch: boolean}>,
		@Inject(MAT_DIALOG_DATA)
		private readonly receivedData: Configuration,
		private readonly configurationService: ConfigurationService,
	) {}

	closeModal() {
		this.dialogRef.close({skipFetch: true});
	}

	save() {
		if (this.form.invalid) {
			this.form.markAllAsTouched();
			return;
		}

		let body: PatchConfiguration = {
			...(this.nameControl.valid &&
				this.nameControl.touched &&
				this.nameControl.dirty && {name: this.nameControl.value}),
			...(this.valueControl.valid &&
				this.valueControl.touched &&
				this.valueControl.dirty && {value: this.valueControl.value}),
		};

		body = _.omitBy(body, _.isNil);
		this.configurationService.updateConfiguration(this.receivedData.id, body).subscribe((_) => {
			this.dialogRef.close({skipFetch: false});
		});
	}

	get nameControl() {
		return this.form.get('name') as FormControl;
	}

	get valueControl() {
		return this.form.get('value') as FormControl;
	}
}
