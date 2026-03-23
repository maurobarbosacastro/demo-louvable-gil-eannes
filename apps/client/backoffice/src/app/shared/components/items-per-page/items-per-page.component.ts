import {Component, EventEmitter, Input, Output} from '@angular/core';
import {setValue, TranslocoModule} from "@ngneat/transloco";
import {ItemsPerPageConstants} from "@app/shared/components/items-per-page/items-per-page.constants";
import {Option} from "@app/shared/components/items-per-page/items-per-page.interface";
import {SelectComponent} from "@app/shared/components/select/select.component";
import {NgIf} from "@angular/common";

@Component({
    selector: 'atl-items-per-page',
    standalone: true,
    imports: [
        TranslocoModule,
        SelectComponent,
        NgIf
    ],
    templateUrl: './items-per-page.component.html',
    styleUrls: ['./items-per-page.component.scss'],
})
export class ItemsPerPageComponent {

    defaultItemsPerPage: string = ItemsPerPageConstants.PAGE_SIZE.toString();
    @Input() totalItems: string;
    @Input() minItemNumber: string;
    @Input() maxItemNumber: string;
    @Output() onSelect = new EventEmitter();

    default: Option = ({
        value: ItemsPerPageConstants.PAGE_SIZE.toString(),
        label: [ItemsPerPageConstants.PAGE_SIZE.toString()]
    })

    options: Option[] = [
        ({
            value: "5",
            label: ["5"]
        }),
        ({
            value: "10",
            label: ["10"]
        }),
        ({
            value: "20",
            label: ["20"]
        }),
        ({
            value: "30",
            label: ["30"]
        }),
    ];

    protected readonly setValue = setValue;

    selectItem($event) {
        this.onSelect.emit($event.value)
    }
}
