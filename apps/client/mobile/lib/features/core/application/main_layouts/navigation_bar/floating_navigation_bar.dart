import 'package:awesome_bottom_bar/awesome_bottom_bar.dart';
import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/main_layouts/navigation_bar/navigation_bar_tab.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class FloatingNavigationBar extends StatelessWidget {
  final void Function(int) onItemSelected;
  final int currentIndex;
  final List<NavigationBarTab> tabs;

  const FloatingNavigationBar({
    super.key,
    required this.onItemSelected,
    required this.currentIndex,
    required this.tabs
  });

  @override
  Widget build(BuildContext context) {
    final bool isCompactSize = MediaQuery.of(context).size.height <= 667;

    return Material(
      elevation: 0,
      type: MaterialType.transparency,
      child: Stack(
        alignment: Alignment.bottomCenter,
        children: [
          Positioned(
            bottom: kBottomNavigationBarHeight + 10,
            left: 0,
            right: 0,
            child: Container(
              height: 1,
              color: AppColors.licorice.withAlpha(30),
            ),
          ),

          BottomBarCreative(
            top: 15,
            bottom: isCompactSize ? 25 : 0,
            items: tabs,
            backgroundColor: Colors.white,
            color: AppColors.licorice.withAlpha(125),
            colorSelected: AppColors.licorice,
            highlightStyle: HighlightStyle(
              sizeLarge: true,
              background: AppColors.licorice,
              color: currentIndex == 1 ? Colors.white : Colors.white.withAlpha(140),
            ),
            iconSize: 22,
            isFloating: true,
            titleStyle: const TextStyle(fontSize: 13),
            indexSelected: currentIndex,
            onTap: onItemSelected,
            boxShadow: [],
          ),

          Positioned.fill(
            child: LayoutBuilder(
              builder: (context, constraints) {
                final itemWidth = constraints.maxWidth / tabs.length;

                return Row(
                  children: List.generate(
                    tabs.length,
                        (index) {
                      final bool isCenterButton = index == 1;

                      return Container(
                        width: itemWidth,
                        height: constraints.maxHeight,
                        padding: const EdgeInsets.symmetric(horizontal: 10),
                        child: Align(
                          alignment: Alignment.topCenter,
                          child: GestureDetector(
                            behavior: HitTestBehavior.translucent,
                            onTap: () {
                              onItemSelected(index);
                            },
                            child: Container(
                              width: double.infinity,
                              height: isCenterButton
                                  ? constraints.maxHeight - 40
                                  : constraints.maxHeight,
                              color: Colors.transparent,
                            ),
                          ),
                        ),
                      );
                    },
                  ),
                );
              },
            ),
          ),
          Positioned(
            top: 62.5,
            child: Text(
              'Shop',
              style: TextStyle(
                color: currentIndex == 1
                    ? AppColors.licorice
                    : AppColors.licorice.withAlpha(125),
                fontSize: 13,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
