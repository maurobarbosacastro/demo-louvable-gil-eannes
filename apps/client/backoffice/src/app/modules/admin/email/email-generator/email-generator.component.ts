import {
	Component,
	computed,
	Input,
	OnInit,
	Signal,
	signal,
	ViewChild,
	WritableSignal,
} from '@angular/core';
import {EmailEditorComponent, EmailEditorModule} from 'angular-email-editor';
import {HeaderService} from '@app/shared/components/header/header.service';
import {NgClass, NgForOf, NgIf} from '@angular/common';
import {MatTooltip} from '@angular/material/tooltip';
import {Router} from '@angular/router';
import {PageMode, TemplateCreation} from '@app/shared/interfaces/template.interface';
import {AppConstants} from '@app/app.constants';
import {MatIcon} from '@angular/material/icon';
import {EmailService} from '@app/modules/admin/email/email.service';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatHint} from '@angular/material/form-field';
import {NgxSpinnerModule} from 'ngx-spinner';
import {MatProgressSpinner} from '@angular/material/progress-spinner';
import {TranslocoPipe} from '@ngneat/transloco';

@Component({
	selector: 'tagpeak-ms-email-generator',
	templateUrl: './email-generator.component.html',
	styleUrl: './email-generator.component.scss',
	standalone: true,
	imports: [
		EmailEditorModule,
		NgForOf,
		MatTooltip,
		MatIcon,
		ReactiveFormsModule,
		NgClass,
		MatHint,
		NgIf,
		NgxSpinnerModule,
		MatProgressSpinner,
		TranslocoPipe,
	],
})
export class EmailGeneratorComponent implements OnInit {
	@Input() id: string;

	variables: string[] = ['username', 'firstName', 'lastName', 'refId', 'store', 'value'];

	options = {
		appearance: {
			theme: 'modern_light',
		},
	};
	scriptUrl = 'https://editor.unlayer.com/embed.js?2';

	$isLoading: WritableSignal<boolean> = signal(true);
	$hasError: WritableSignal<boolean> = signal(false);
	$isNew: Signal<boolean> = computed(() => this.id === PageMode.new);

	templateForm: FormGroup = new FormGroup({
		templateName: new FormControl('', [Validators.required]),
		templateCode: new FormControl('', [Validators.required]),
	});

	@ViewChild('editor')
	private emailEditor: EmailEditorComponent;

	constructor(
		private headerService: HeaderService,
		private readonly router: Router,
		private readonly emailService: EmailService,
	) {}

	ngOnInit(): void {
		this.headerService.setHeaderInfo({title: 'email', button: false});
	}

	editorLoaded($event) {
		if (!this.$isNew()) {
			this.emailService.getTemplate(this.id).subscribe((template) => {
				this.emailEditor.editor.loadDesign(JSON.parse(template.templateJson));

				this.templateCodeControl.setValue(template.code);
				this.templateCodeControl.disable();

				this.templateNameControl.setValue(template.name);
			});
		} else {
			if (this.emailService.jsonTemplateValue) {
				this.emailEditor.editor.loadDesign(JSON.parse(this.emailService.jsonTemplateValue));
			}
		}
	}

	editorReady($event) {
		this.$isLoading.set(false);
	}

	// Not used for now, could be used in the future
	/*
    saveDesign() {
        this.emailEditor.editor.saveDesign((data) =>
            console.log('saveDesign', data)
        );
    }*/

	saveTemplate() {
		if (!this.templateForm.valid) {
			this.$hasError.set(true);
			return;
		}

		// Using exportHtml function to create a new template because with this function we have the designJSON and designHTML
		this.emailEditor.editor.exportHtml((data) => {
			const body: TemplateCreation = {
				code: this.templateCodeControl.value,
				name: this.templateNameControl.value,
				templateJson: JSON.stringify(data.design),
				templateHtml: data.html.replace(/\n/g, '').replace(/\"/g, '\"'),
			};

			if (this.$isNew()) {
				this.emailService.createTemplate(body).subscribe((_) => {
					this.returnToTable();
				});
			} else {
				this.emailService.updateTemplate(this.id, body).subscribe((_) => {
					this.returnToTable();
				});
			}
		},{
            inlineStyles: true
        });
	}

	copyToClipboard(variable: string): void {
		const toBeCopied = '{{' + variable + '}}';
		navigator.clipboard.writeText(toBeCopied);
	}

	returnToTable(): void {
		this.router.navigate([AppConstants.ROUTES.admin + AppConstants.ROUTES.email]);
	}

	get templateNameControl(): FormControl {
		return this.templateForm.get('templateName') as FormControl;
	}

	get templateCodeControl(): FormControl {
		return this.templateForm.get('templateCode') as FormControl;
	}

	onTemplateNameChange(): void {
		this.templateCodeControl.setValue(
			this.templateNameControl.value.replaceAll(' ', '-').toLowerCase(),
		);
	}
}
