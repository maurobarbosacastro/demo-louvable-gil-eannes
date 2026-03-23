import 'package:flutter/material.dart';
import 'package:number_paginator/number_paginator.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class TgpkPaginator extends StatelessWidget {
  final double numberPaginatorHeight;

  final int pages;

  final Function(int) onPageChanged;

  final NumberPaginatorController controller;

  const TgpkPaginator(
      {super.key,
      required this.numberPaginatorHeight,
      required this.pages,
      required this.onPageChanged,
      required this.controller});

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: BoxConstraints(
        maxWidth: pages <= 2 ? 150 : 250,
      ),
      child: NumberPaginator(
        initialPage: controller.value,
        controller: controller,
        numberPages: pages,
        onPageChange: (int index) {
          onPageChanged(index);
        },
        child: SizedBox(
          height: numberPaginatorHeight,
          child: Row(
            children: [
              const PrevButton(),
              Expanded(
                child: NumberContent(
                  buttonBuilder: (context, index, isSelected) =>
                      _CustomItemPaginator(
                    onTap: () => controller.navigateToPage(index),
                    number: '${index + 1}',
                    isSelected: isSelected,
                  ),
                ),
              ),
              const NextButton(),
            ],
          ),
        ),
      ),
    );
  }
}

// Custom paginator item
class _CustomItemPaginator extends StatelessWidget {
  final VoidCallback onTap;
  final String number;
  final bool isSelected;

  const _CustomItemPaginator({
    required this.onTap,
    required this.number,
    this.isSelected = false,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: 25,
      height: 25,
      child: ElevatedButton(
        onPressed: onTap,
        style: ElevatedButton.styleFrom(
          elevation: 0,
          backgroundColor: isSelected ? AppColors.wildSand : Colors.transparent,
          foregroundColor: Colors.white,
          padding: EdgeInsets.zero,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(4),
          ),
        ),
        child: Text(number,
            style: const TextStyle(fontSize: 12, color: AppColors.waterloo)),
      ),
    );
  }
}
