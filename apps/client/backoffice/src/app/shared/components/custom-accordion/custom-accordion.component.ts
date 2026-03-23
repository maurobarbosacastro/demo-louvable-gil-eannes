import {Component, Input, signal, TemplateRef, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe} from '@ngneat/transloco';

@Component({
	selector: 'tagpeak-custom-accordion',
	standalone: true,
	imports: [CommonModule, MatIcon, TranslocoPipe],
	templateUrl: './custom-accordion.component.html',
})
export class CustomAccordionComponent {
	@Input({required: false}) title: string;
	@Input({required: false}) headerTemplate: TemplateRef<any>;
	@Input({required: true}) bodyTemplate: TemplateRef<any>;
	@Input({required: false}) cssHeader: string;

	$showAccordion: WritableSignal<boolean> = signal(true);

	clickAccordion(): void {
		this.$showAccordion.update((state: boolean) => !state);
	}
}
