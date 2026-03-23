import {
	Component,
	computed,
	EventEmitter,
	inject,
	input,
	InputSignal,
	Output,
	Signal,
} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MatIcon} from '@angular/material/icon';
import {MatTooltip} from '@angular/material/tooltip';

@Component({
	selector: 'tagpeak-status-container',
	standalone: true,
	imports: [CommonModule, TranslocoPipe, MatIcon, MatTooltip],
	template: `
        <div class="p-2.5 rounded text-center flex items-center gap-1 !h-[{{$height()}}] !min-h-[{{$height()}}] w-fit"
             [matTooltip]="$description()"
             [matTooltipDisabled]="$disableTooltip()"
             [class.status-pending]="$status() === 'TRACKED' || $status() === 'PENDING' || $status() === 'REQUESTED' || $status() === 'INACTIVE'"
             [class.status-finished-paid]="$status() === 'PAID' || $status() === 'FINISHED' || $status() === 'VALIDATED' || $status() === 'COMPLETED' || $status() === 'CONFIRMED' || $status() === 'ACTIVE'"
             [class.status-stopped]="$status() === 'STOPPED' || $status() === 'REJECTED' || $status() === 'CANCELLED'"
             [class.status-live]="$status() === 'LIVE'">
            <span class="text-md font-medium leading-none">
                @if ($alt()) {
                    {{ ('status.alt.' + $status()) | transloco }}
                } @else if ($withdrawal()) {
                    {{ ('status.withdrawal.' + $status()) | transloco }}
                } @else {
                    {{ ('status.' + $status()) | transloco }}
                }
            </span>
            @if ($icon()) {
                <mat-icon [svgIcon]="$icon()" class="icon-size-4 cursor-pointer icon-color" (click)="iconClick()"></mat-icon>
            }
        </div>`,
	styles: `
        .status-pending {
            background: #7A7B921A;
            color: #7A7B92;

            .icon-color {
                color: #7A7B92;
            }
        }

        .status-live {
            background: #FDCA6566;
            color: #CF8E0C;

            .icon-color {
                color: #CF8E0C;
            }
        }

        .status-finished-paid {
            background: #4237DA33;
            color: #4237DA;

            .icon-color {
                color: #4237DA;
            }
        }

        .status-stopped {
            background: #F2295B33;
            color: #F2295B;

            .icon-color {
                color: #F2295B;
            }
        }
    `,
})
export class StatusContainerComponent {
	$status: InputSignal<string> = input.required();
	$icon: InputSignal<string> = input(null);
	$disableTooltip: InputSignal<boolean> = input(true);
	$alt: InputSignal<boolean> = input(false);
	$withdrawal: InputSignal<boolean> = input(false);

	transloco = inject(TranslocoService);
	$description: Signal<string> = computed(() => {
		return this.transloco.translate(`status.description.${this.$status()}`);
	});
	$height: InputSignal<string> = input('34px');

	@Output() iconEvent: EventEmitter<string> = new EventEmitter<string>();

	iconClick(): void {
		this.iconEvent.emit(this.$status());
	}
}
