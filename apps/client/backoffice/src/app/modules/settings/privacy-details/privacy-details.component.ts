import {Component} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe} from "@ngneat/transloco";

@Component({
	selector: 'tagpeak-privacy-details',
	standalone: true,
	imports: [CommonModule, TranslocoPipe],
	templateUrl: './privacy-details.component.html',
})
export class PrivacyDetailsComponent {}
