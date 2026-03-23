import 'package:flutter/material.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class TgpkScrollableTableColumn {
  final String key;
  final String name;

  const TgpkScrollableTableColumn(this.key, this.name);
}

class TgpkScrollableTableMetric {
  final String key;
  final String name;

  const TgpkScrollableTableMetric(this.key, this.name);
}

class TgpkScrollableTable extends StatefulWidget {
  final Map<String, dynamic>? dataSource;
  final List<TgpkScrollableTableColumn> columns;
  final List<TgpkScrollableTableMetric> metrics;
  final double metricsWrapperWidth;
  final double rowWrapperHeight;

  const TgpkScrollableTable({
    super.key,
    required this.dataSource,
    required this.columns,
    required this.metrics,
    this.metricsWrapperWidth = 80,
    this.rowWrapperHeight = 48
  });

  @override
  State<TgpkScrollableTable> createState() => _TgpkScrollableTableState();
}

class _TgpkScrollableTableState extends State<TgpkScrollableTable> {
  late final ScrollController _horizontalScrollController;

  @override
  void initState() {
    super.initState();
    _horizontalScrollController = ScrollController();
  }

  @override
  void dispose() {
    _horizontalScrollController.dispose();
    super.dispose();
  }

  String _formatValue(dynamic value) {
    if (value == null) return '-';

    if (value is num) {
      return value % 1 == 0 ? value.toInt().toString() : value.toString();
    }

    return value.toString();
  }

  @override
  Widget build(BuildContext context) {
    final List<TgpkScrollableTableMetric> metrics = widget.metrics;

    return Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        // Metrics Column
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            SizedBox(
              width: widget.metricsWrapperWidth,
              height: widget.rowWrapperHeight,
            ),
            ...metrics.asMap().entries.map((entry) {
              final int idx = entry.key;
              final String label = entry.value.name;
              return Container(
                width: widget.metricsWrapperWidth,
                height: widget.rowWrapperHeight,
                alignment: Alignment.centerLeft,
                padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 8),
                decoration: BoxDecoration(
                  border: Border(
                    top: const BorderSide(color: AppColors.wildSand),
                    bottom: idx < metrics.length - 1
                        ? const BorderSide(color: AppColors.wildSand)
                        : BorderSide.none,
                  ),
                ),
                child: Text(
                  label,
                  softWrap: true,
                ),
              );
            }),
          ],
        ),

        // Grid Content
        Expanded(
          child: SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            controller: _horizontalScrollController,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                // Header Row
                Row(
                  children: widget.columns.map((column) {
                    return Container(
                      width: 100,
                      height: widget.rowWrapperHeight,
                      alignment: Alignment.center,
                      padding: const EdgeInsets.symmetric(vertical: 10),
                      decoration: const BoxDecoration(
                        border: Border(bottom: BorderSide(color: AppColors.wildSand)),
                      ),
                      child: Text(column.name),
                    );
                  }).toList(),
                ),

                // Data Rows
                ...List.generate(metrics.length, (rowIndex) {
                  final metricKey = metrics[rowIndex];

                  return Row(
                    children: widget.columns.map((column) {
                      final item = widget.dataSource?[column.key.toUpperCase()];

                      final value = item?.toJson()[metricKey.key];

                      return Container(
                        width: 100,
                        height: widget.rowWrapperHeight,
                        alignment: Alignment.center,
                        padding: const EdgeInsets.symmetric(vertical: 10),
                        decoration: BoxDecoration(
                          border: Border(
                            bottom: rowIndex < metrics.length - 1
                                ? const BorderSide(color: AppColors.wildSand)
                                : BorderSide.none,
                          ),
                        ),
                        child: Text(
                          _formatValue(value),
                        ),
                      );
                    }).toList(),
                  );
                }),
              ],
            ),
          ),
        ),
      ],
    );
  }
}
