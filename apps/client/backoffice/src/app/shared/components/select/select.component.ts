import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import {TranslocoModule} from "@ngneat/transloco";

@Component({
    selector: 'app-select',
    standalone: true,
    imports: [CommonModule, MatIconModule, TranslocoModule],
    templateUrl: './select.component.html',
    styleUrls: ['./select.component.scss']
})
export class SelectComponent implements OnInit, OnChanges {

    @Input() placeholder: string = 'Select a value';
    @Input() initValue: any;
    @Input() options: { value: string; label: string[]; icons?: string[] }[];
    @Input() icons: string[];
    @Input() reset: boolean;
    @Input() disable: boolean;
    @Input() required: boolean;
    @Input() emptyOption: boolean = true;
    @Input() openUpwards: boolean = false;

    @Output()
    valueChange: EventEmitter<any> = new EventEmitter<any>();

    open: boolean = false;
    currentValue: FormControl;

    constructor() {
    }

    ngOnInit(): void {
        this.currentValue = new FormControl<any>(this.initValue);
        this.currentValue.valueChanges
            .subscribe((res) => {
                this.valueChange.emit(res);
            });
    }

    ngOnChanges(changes: SimpleChanges) {
        if (this.initValue && this.currentValue) {
            this.currentValue.setValue(this.initValue);
        }
        if (changes.reset?.currentValue) {
            this.currentValue.patchValue(null);
        }
    }

    unfocus(target): void {
        this.currentValue.setValue(target);
        this.open = false;
        setTimeout(() => {
            (document.activeElement as any).blur();
        });
    }
}
