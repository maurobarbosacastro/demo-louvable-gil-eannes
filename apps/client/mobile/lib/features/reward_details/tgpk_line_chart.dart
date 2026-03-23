import 'package:flutter/material.dart';
import 'package:syncfusion_flutter_charts/charts.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class TgpkLineChartDataset {
  final List<dynamic> spots;
  final Color color;

  TgpkLineChartDataset({
    required this.spots,
    required this.color,
  });
}

class TgpkConstantLineChart {
  final double y;
  final Color color;
  final bool showLabel;

  TgpkConstantLineChart({
    required this.y,
    required this.color,
    this.showLabel = false
  });
}

class TgpkLineChart extends StatelessWidget {
  final List<TgpkLineChartDataset> datasets;
  final double minY;
  final double maxY;
  final Widget Function(BuildContext context, TrackballDetails trackballDetails)? tooltipBuilder;
  final List<TgpkConstantLineChart>? constants;
  final bool hiddenYAxisLabels;
  final Color? horizontalLineColor;
  final Color? trackballColor;

  const TgpkLineChart({
    super.key,
    required this.datasets,
    this.minY = 0,
    required this.maxY,
    this.tooltipBuilder,
    this.constants,
    this.hiddenYAxisLabels = false,
    this.horizontalLineColor,
    this.trackballColor
  });

  @override
  Widget build(BuildContext context) {
    return SfCartesianChart(
      primaryXAxis: const CategoryAxis(
        axisLine: AxisLine(width: 0),
        majorTickLines: MajorTickLines(width: 0),
        majorGridLines: MajorGridLines(width: 0),
      ),
      primaryYAxis: NumericAxis(
        opposedPosition: true,
        axisLine: const AxisLine(width: 0),
        majorTickLines: const MajorTickLines(width: 0),
        majorGridLines: const MajorGridLines(
            width: 1,
            color: AppColors.wildSand
        ),
        axisLabelFormatter: (AxisLabelRenderDetails details) {
          if (hiddenYAxisLabels) return ChartAxisLabel('', const TextStyle());
          return ChartAxisLabel(
            details.value.toStringAsFixed(2),
            details.textStyle,
          );
        },
        minimum: minY,
        maximum: maxY + 1,
        plotOffset: 0,
        edgeLabelPlacement: EdgeLabelPlacement.none,
        desiredIntervals: 4,
        plotBands: <PlotBand>[
          if (constants != null)
            ...constants!.map((cl) => PlotBand(
              isVisible: true,
              start: cl.y,
              end: cl.y,
              borderColor: cl.color,
              borderWidth: 3,
            )),
          PlotBand(
              start: minY,
              end: minY,
              borderWidth: 1,
              borderColor: Colors.grey[400]!
          ),
          PlotBand(
              start: maxY + 1,
              end: maxY + 1,
              borderWidth: 1,
              borderColor: Colors.grey[400]!
          ),
        ],
      ),
      trackballBehavior: TrackballBehavior(
        tooltipSettings: const InteractiveTooltip(
          arrowWidth: 0,
          arrowLength: 0,
          enable: false,
        ),
        enable: true,
        lineColor: horizontalLineColor ?? AppColors.youngBlue,
        lineDashArray: const [3, 3],
        activationMode: ActivationMode.singleTap,
        markerSettings: TrackballMarkerSettings(
          markerVisibility: TrackballVisibilityMode.visible,
          height: 10,
          width: 10,
          shape: DataMarkerType.circle,
          color: trackballColor ?? AppColors.youngBlue,
          borderColor: Colors.white,
          borderWidth: 1,
        ),
        builder: (context, trackballDetails) {
          if (trackballDetails.point == null) {
            return const SizedBox.shrink();
          }
          return tooltipBuilder?.call(context, trackballDetails) ?? const SizedBox.shrink();
        },
      ),
      series: datasets.map(
            (dataset) => LineSeries<dynamic, String>(
          color: dataset.color,
          dataSource: dataset.spots,
          xValueMapper: (dynamic spot, _) => spot['x'],
          yValueMapper: (dynamic spot, _) => spot['y'],
          dataLabelSettings: const DataLabelSettings(isVisible: false),
        ),
      ).toList(),
    );
  }
}
