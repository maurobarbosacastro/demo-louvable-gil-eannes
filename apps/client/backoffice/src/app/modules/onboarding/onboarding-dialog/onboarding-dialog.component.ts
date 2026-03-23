import { Component, ViewChild } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { MatIcon } from '@angular/material/icon';
import { TagpeakCarouselComponent } from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import { TranslocoPipe } from '@ngneat/transloco';
import { MatDialogClose } from '@angular/material/dialog';

@Component({
    selector: 'tagpeak-onboarding-dialog',
    standalone: true,
    imports: [CommonModule, MatIcon, NgOptimizedImage, TagpeakCarouselComponent, TranslocoPipe, MatDialogClose],
    templateUrl: './onboarding-dialog.component.html',
    styleUrl: './onboarding-dialog.component.scss'
})
export class OnboardingDialogComponent {

    @ViewChild('carousel') carousel!: TagpeakCarouselComponent;

    goNextStep(): void {
        this.carousel.nextSlide();
    }
}
