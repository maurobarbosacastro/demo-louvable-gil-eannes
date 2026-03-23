import {
	ApexAxisChartSeries,
	ApexChart,
	ApexDataLabels,
	ApexGrid,
	ApexLegend,
	ApexMarkers,
	ApexStroke,
	ApexTitleSubtitle,
	ApexTooltip,
	ApexXAxis,
	ApexYAxis,
} from 'ng-apexcharts';

export type ChartOptions = {
	series: ApexAxisChartSeries;
	chart: ApexChart;
	xAxis: ApexXAxis;
	dataLabels: ApexDataLabels;
	grid: ApexGrid;
	stroke: ApexStroke;
	title: ApexTitleSubtitle;
	yAxis: ApexYAxis;
	markers: ApexMarkers;
	tooltip: ApexTooltip;
	legends: ApexLegend;
};
