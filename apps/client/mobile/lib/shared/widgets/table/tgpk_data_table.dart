import 'package:flutter/material.dart';
import 'package:number_paginator/number_paginator.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/dashboard/store_visits/without_results.dart';
import 'package:tagpeak/shared/widgets/table/tgpk_paginator.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class TgpkDataTable<T> extends StatefulWidget {
  final List<DataColumn> columns;

  final List<T> items;

  final void Function(T)? onCellClick;
  final double rowHeight;

  final List<Widget> Function(T item, BoxConstraints constraints) rowBuilder;

  final Function(int) onPageChanged;

  final int pages;

  final bool isLoading;
  final bool hasActiveFilters;
  /// Displayed when data loading fails
  final String? errorStateMessage;
  /// Displayed when filtered results are empty
  final String? emptyStateMessage;

  const TgpkDataTable(
      {super.key,
      required this.columns,
      required this.items,
      required this.rowBuilder,
      required this.onPageChanged,
      required this.pages,
      this.onCellClick,
      this.rowHeight = 65,
      this.isLoading = false,
      this.hasActiveFilters = false,
      this.errorStateMessage,
      this.emptyStateMessage});

  @override
  State<TgpkDataTable<T>> createState() => _TgpkDataTableState<T>();
}

class _TgpkDataTableState<T> extends State<TgpkDataTable<T>> {
  final double _headingHeight = 30;

  final double _numberPaginatorHeight = 35;

  final controller = NumberPaginatorController();

  @override
  Widget build(BuildContext context) {
    return widget.isLoading
        ? const Center(
            child: CircularProgressIndicator(color: AppColors.waterloo))
        : widget.items.isNotEmpty
            ? Column(
                spacing: 25,
                crossAxisAlignment: CrossAxisAlignment.end,
                children: [
                  Theme(
                    data: Theme.of(context).copyWith(
                      dividerTheme: const DividerThemeData(
                        color: Colors.transparent,
                        thickness: 0,
                      ),
                      cardTheme: const CardTheme(
                        shape: RoundedRectangleBorder(
                          side: BorderSide.none,
                          borderRadius: BorderRadius.zero,
                        ),
                        elevation: 0,
                      ),
                    ),
                    child: LayoutBuilder(
                      builder: (context, constraints) {
                        return ConstrainedBox(
                          constraints: BoxConstraints(
                              minWidth: MediaQuery.of(context).size.width),
                          child: DataTable(
                            showCheckboxColumn: false,
                            horizontalMargin: 0,
                            columnSpacing: 0,
                            dataRowMaxHeight: widget.rowHeight,
                            headingRowHeight: _headingHeight,
                            columns: widget.columns,
                            rows: widget.items
                                .map((item) => DataRow(
                                      cells: widget
                                          .rowBuilder(item, constraints)
                                          .map((widget) => DataCell(widget))
                                          .toList(),
                                      onSelectChanged: widget.onCellClick !=
                                              null
                                          ? (bool? selected) {
                                              if (selected ?? false) {
                                                widget.onCellClick?.call(item);
                                              }
                                            }
                                          : null,
                                    ))
                                .toList(),
                          ),
                        );
                      },
                    ),
                  ),
                  widget.pages > 1 && widget.items.isNotEmpty
                      ? TgpkPaginator(
                          controller: controller,
                          numberPaginatorHeight: _numberPaginatorHeight,
                          pages: widget.pages,
                          onPageChanged: widget.onPageChanged,
                        )
                      : const SizedBox.shrink()
                ],
              )
            : Align(
              alignment: Alignment.centerLeft,
              child: !widget.hasActiveFilters
                  ? WithoutResults(
                      title: translation(context).nothingHere,
                      description: widget.errorStateMessage ?? translation(context).defaultWithoutResults)
                  : WithoutResults(
                      title: translation(context).noResultsFound,
                      description: widget.emptyStateMessage ?? translation(context).defaultWithoutResultsAfterSearch),
            );
  }
}
