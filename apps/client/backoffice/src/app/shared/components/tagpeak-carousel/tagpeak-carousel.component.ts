import {
	Component,
	DestroyRef,
	inject,
	input,
	InputSignal,
	OnInit,
	signal,
	TemplateRef,
	WritableSignal,
} from '@angular/core';
import {CommonModule} from '@angular/common';
import {interval} from 'rxjs';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';

@Component({
	selector: 'tagpeak-carousel',
	standalone: true,
	imports: [CommonModule],
	templateUrl: './tagpeak-carousel.component.html',
})
export class TagpeakCarouselComponent implements OnInit {
	destroyRef = inject(DestroyRef);

	$templates: InputSignal<TemplateRef<any>[]> = input([]);
	$action: InputSignal<TemplateRef<any>> = input(null);
	$image: InputSignal<TemplateRef<any>> = input(null);
	$currentIndex: WritableSignal<number> = signal(0);
	$slideAutomatically: InputSignal<boolean> = input(true);
	$bgClassSlideIndicator: InputSignal<string> = input('!bg-white');
	// Could be RECT -> rectangular or ROUNDED -> rounded
	$slideIndicatorType: InputSignal<string> = input('RECT');

	timerInterval = 10000; // Interval in milliseconds

	ngOnInit(): void {
		if (this.$slideAutomatically()) {
			this.startTimer();
		}
	}

	selectCard(index: number) {
		this.$currentIndex.set(index);
	}

	startTimer() {
		interval(this.timerInterval)
			.pipe(takeUntilDestroyed(this.destroyRef))
			.subscribe(() => {
				this.nextSlide();
			});
	}

	nextSlide() {
		this.$currentIndex.set((this.$currentIndex() + 1) % this.$templates().length);
	}

	slideSelectedColor(idx: number): string {
		return this.$bgClassSlideIndicator() + (idx === this.$currentIndex() ? '' : ' bg-opacity-20');
	}
}
