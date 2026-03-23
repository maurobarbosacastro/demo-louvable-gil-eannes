import {Component, inject, OnInit} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogClose, MatDialogRef} from '@angular/material/dialog';
import {BugReportRequest} from '@app/modules/client/bug-report/models/bug-report.interface';
import {BugReportService} from '@app/modules/client/bug-report/services/bug-report.service';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {UserService} from '@app/core/user/user.service';

@Component({
	selector: 'tagpeak-bug-report',
	standalone: true,
	imports: [CommonModule, ReactiveFormsModule, TranslocoPipe, MatDialogClose],
	templateUrl: './bug-report.component.html',
})
export class BugReportComponent implements OnInit {
	bugForm = new FormGroup({
		name: new FormControl('', [Validators.required]),
		type: new FormControl('technical', [Validators.required]), // Default to "technical"
		email: new FormControl('', [Validators.required]),
		description: new FormControl('', [Validators.required]),
		file: new FormControl<File | null>(null),
	});

	private dialogData = inject(MAT_DIALOG_DATA);
	private dialogRef = inject(MatDialogRef);
	private bugReportService = inject(BugReportService);
	private transloco = inject(TranslocoService);
	private toasterService = inject(ToasterService);
    private userService = inject(UserService);

	selectedFile: File | null = null;
	base64File: string = '';
	mimeType: string = '';

	ngOnInit(): void {
		this.setForm();
	}

	setForm(): void {
		this.bugForm.patchValue({
			name: this.userService.$user().firstName + ' ' + this.userService.$user().lastName,
			email: this.userService.$user().email,
		});
	}

	// Handle file selection
	fileSelected(event: Event): void {
		const input = event.target as HTMLInputElement;
		if (input.files?.length) {
			this.selectedFile = input.files[0]; // Store file separately (NOT in FormControl)

			const reader = new FileReader();
			reader.readAsDataURL(this.selectedFile);
			reader.onload = () => {
				const base64String = reader.result as string;
				this.base64File = base64String.split(',')[1];
				this.mimeType = this.selectedFile!.type;
			};
		}
	}

	// Submit the form
	submitForm(): void {
		const body: BugReportRequest = {
			name: this.bugForm.value.name!,
			type: this.bugForm.value.type!,
			email: this.bugForm.value.email!,
			description: this.bugForm.value.description!,
		};

		if (this.selectedFile) {
			body.attachment = {
				filename: this.selectedFile.name,
				data: this.base64File,
				mimeType: this.mimeType,
			};
		}

		this.bugReportService.submitBugReport(body).subscribe(
			(_) => {
				this.toasterService.showToast('success', this.transloco.translate('bug-report.success'));
			},
			(_) => {
				this.toasterService.showToast('error', this.transloco.translate('bug-report.error'));
			},
		);

		this.dialogRef.close();
	}

	protected readonly ReportType = ReportType;
}

export enum ReportType {
	Technical = 'Technical Issue',
	Question = 'Question',
}
