import { Component, input, InputSignal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ChartComponent } from "ng-apexcharts";
import { ChartOptions } from "@app/shared/components/line-chart/models/line-chart";

@Component({
    selector: 'tagpeak-line-chart',
    standalone: true,
    imports: [CommonModule, ChartComponent],
    templateUrl: './line-chart.component.html',
})
export class LineChartComponent {

    $chartOptions: InputSignal<Partial<ChartOptions>> = input.required();

}
