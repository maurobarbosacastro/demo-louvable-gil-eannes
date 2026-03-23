import {Component, Input, TemplateRef} from '@angular/core';
import {CommonModule} from '@angular/common';

@Component({
	selector: 'tagpeak-info-card',
	standalone: true,
	imports: [CommonModule],
	template: `
        <div class="w-full h-fit" [ngClass]="styleClass">
            <ng-container *ngTemplateOutlet="templateCardBody"></ng-container>
        </div>
    `,
})
export class InfoCardComponent {
	@Input({required: true}) templateCardBody: TemplateRef<any>;
	@Input({required: false}) styleClass: string = 'bg-project-youngBlue rounded-[20px]';
}
