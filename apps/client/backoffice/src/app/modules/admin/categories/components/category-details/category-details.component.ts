import {Component, DestroyRef, effect, Inject, OnInit, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {SharedModule} from '@app/shared/shared.module';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {CategoriesInterface} from '@app/modules/admin/categories/models/categories.interface';
import {CategoriesService} from '@app/modules/admin/categories/services/categories.service';
import {MatIcon} from '@angular/material/icon';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {filter, map, of, Subject, switchMap, takeUntil} from 'rxjs';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {HttpErrorResponse} from '@angular/common/http';

@Component({
	selector: 'tagpeak-category-details',
	standalone: true,
	imports: [CommonModule, ReactiveFormsModule, SharedModule, TranslocoPipe, MatIcon],
	templateUrl: './category-details.component.html',
})
export class CategoryDetailsComponent {
	categoryForm = new FormGroup({
		code: new FormControl({value: '', disabled: true}, [Validators.required]),
		name: new FormControl('', [Validators.required]),
	});

	$editCategory: Signal<boolean> = toSignal(
		of(this.receivedData).pipe(
			filter((data: {id: string}) => !!data?.id),
			map((data: {id: string}) => data?.id),
			switchMap((id: string) => this.categoriesService.getCategoryById(id)),
			map((category: CategoriesInterface) => {
				this.setupForm(category);
				return true;
			}),
		),
	);

	_ = effect(() => {
		if (!this.$editCategory()) {
			this.categoryForm
				.get('name')
				.valueChanges.pipe(takeUntil(this.killIt$), takeUntilDestroyed(this.destroyRef))
				.subscribe((value) => {
					if (value) {
						this.categoryForm.get('code').setValue(value.toLowerCase().replace(/ /g, '_'));
					}
				});
		} else {
			this.killIt$.next();
		}
	});

	killIt$: Subject<void> = new Subject();

	constructor(
		private readonly myDialog: MatDialogRef<CategoryDetailsComponent, {skipFetch: boolean}>,
		@Inject(MAT_DIALOG_DATA)
		private readonly receivedData: {id: string},
		private readonly categoriesService: CategoriesService,
		private readonly toasterService: ToasterService,
		private readonly translocoService: TranslocoService,
		private readonly destroyRef: DestroyRef,
	) {}

	private setupForm(category: CategoriesInterface): void {
		this.categoryForm.patchValue({
			code: category.code,
			name: category.name,
		});
	}

	closeModal(): void {
		this.myDialog.close({skipFetch: true});
	}

	finalize(): void {
		if (this.categoryForm.invalid) {
			this.categoryForm.markAllAsTouched();
			return;
		}

		const body: Partial<CategoriesInterface> = {
			code: this.codeControl.value,
			name: this.nameControl.value,
		};

		if (this.$editCategory()) {
			this.categoriesService.updateCategory(body, this.receivedData.id).subscribe({
				next: () => {
					this.toasterService.showToast(
						'success',
						this.translocoService.translate('categories.success-update'),
					);
					this.myDialog.close({skipFetch: false});
				},
				error: (error: HttpErrorResponse) => {
					this.toasterService.showToast('error', error.error.ErrorMessage);
				},
			});
		} else {
			this.categoriesService.createCategory(body).subscribe({
				next: () => {
					this.toasterService.showToast(
						'success',
						this.translocoService.translate('categories.success-create'),
					);
					this.myDialog.close({skipFetch: false});
				},
				error: (error: HttpErrorResponse) => {
					this.toasterService.showToast('error', error.error.ErrorMessage);
				},
			});
		}
	}

	get nameControl(): FormControl {
		return this.categoryForm.get('name') as FormControl;
	}

	get codeControl(): FormControl {
		return this.categoryForm.get('code') as FormControl;
	}
}
