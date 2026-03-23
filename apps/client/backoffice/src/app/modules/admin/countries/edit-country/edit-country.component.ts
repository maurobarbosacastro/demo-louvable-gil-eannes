import {Component, inject, OnInit, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {CountriesService} from "@app/modules/admin/countries/countries.service";
import {CountryInterface} from "@app/shared/interfaces/country.interface";
import {FormControl, ReactiveFormsModule, UntypedFormBuilder, UntypedFormGroup, Validators} from "@angular/forms";
import {MatIcon} from "@angular/material/icon";
import {TranslocoPipe, TranslocoService} from "@ngneat/transloco";
import {ToasterService} from "@app/shared/components/toaster/toaster.service";

@Component({
    selector: 'tagpeak-edit-country',
    standalone: true,
    imports: [CommonModule, MatIcon, ReactiveFormsModule, TranslocoPipe],
    templateUrl: './edit-country.component.html',
})
export class EditCountryComponent implements OnInit {

    private dialogData = inject(MAT_DIALOG_DATA);
    private dialogRef = inject(MatDialogRef);
    private countriesService = inject(CountriesService);
    private _formBuilder = inject(UntypedFormBuilder);
    private toaster = inject(ToasterService);
    private transloco = inject(TranslocoService);

    $country: WritableSignal<CountryInterface> = signal(null);

    countryForm: UntypedFormGroup;
    editCountry: boolean = false;

    ngOnInit(): void {

        this.countryForm = this._formBuilder.group({
            abbreviation: ['', Validators.required],
            name: ['', Validators.required],
            currency: ['', Validators.required],
            enabled: [true],
        });

        if (this.dialogData.country_uuid) {
            this.editCountry = true;
            this.countriesService.getCountry(this.dialogData.country_uuid).subscribe(country => {
                this.$country.set(country);
                this.countryForm.patchValue(country);
            });
        }
    }


    get abbreviationControl(): FormControl {
        return this.countryForm.get('abbreviation') as FormControl;
    }

    get nameControl(): FormControl {
        return this.countryForm.get('name') as FormControl;
    }

    get currencyControl(): FormControl {
        return this.countryForm.get('currency') as FormControl;
    }

    get enabledControl(): FormControl {
        return this.countryForm.get('enabled') as FormControl;
    }

    closeModal() {
        this.dialogRef.close({skipFetch: true});
    }

    saveCountry() {
        if (this.countryForm.invalid) {
            this.countryForm.markAllAsTouched();
            return;
        }

        let countryDto: CountryInterface = this.countryForm.value;
        if (this.editCountry) {

            this.countriesService.updateCountry(this.dialogData.country_uuid, countryDto).subscribe(country => {
                if (country) {
                    this.toaster.showToast('success', this.transloco.translate('country.country-updated'));
                    this.dialogRef.close({skipFetch: false});
                } else {
                    this.toaster.showToast('error', this.transloco.translate('country.country-update-error'));
                }
            });
        } else {
            this.countriesService.createCountry(countryDto).subscribe(country => {
                if (country) {
                    this.toaster.showToast('success', this.transloco.translate('country.country-created'));
                    this.dialogRef.close({skipFetch: false});
                } else {
                    this.toaster.showToast('error', this.transloco.translate('country.country-create-error'));
                }
            });
        }
    }
}
