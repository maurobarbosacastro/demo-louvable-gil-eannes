import {Component, input, InputSignal} from '@angular/core';
import {CommonModule} from '@angular/common';

@Component({
	selector: 'tagpeak-stat-card',
	standalone: true,
	imports: [CommonModule],
	template: `
        <div class="card bg-base-100 w-full shadow">
            <div class="card-body">
                <h2 class="card-title">{{ $title() }}</h2>
                <div class="card-actions justify-end">
                    @if ($currency()) {
                        <span class="text-xl font-bold text-base-content">
                            {{ $value() | currency: $currency() }}
                        </span>
                    }
                    @else {
                        <span class="text-xl font-bold text-base-content">
                            {{ $value() }}
                        </span>
                    }
                </div>

            </div>
        </div>
    `,
	styles: ``,
})
export class StatCardComponent {

    $title: InputSignal<string> = input.required<string>();
    $value: InputSignal<number> = input.required<number>();
    $currency: InputSignal<string> = input<string>();

}
