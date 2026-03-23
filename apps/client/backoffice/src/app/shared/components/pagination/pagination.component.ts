import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIconModule} from '@angular/material/icon';

@Component({
    selector: 'app-pagination',
    standalone: true,
    imports: [CommonModule, MatIconModule,],
    templateUrl: './pagination.component.html',
    styleUrls: ['./pagination.component.scss']
})
export class PaginationComponent implements OnInit {


    @Input() page: number;
    @Input() total: number;
    @Output() next = new EventEmitter();
    @Output() previous = new EventEmitter();


    constructor() {
    }

    ngOnInit(): void {
    }

    nextPage(): void {
        if (this.page <= this.total) {
            this.next.emit();
        }
    }

    previousPage(): void {
        if (this.page >= 1) {
            this.previous.emit();
        }
    }

}
