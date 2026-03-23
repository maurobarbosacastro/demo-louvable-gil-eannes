import 'package:awesome_bottom_bar/awesome_bottom_bar.dart';
import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/main_layouts/navigation_bar/navigation_bar_tab.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class SimpleNavigationBar extends StatelessWidget {
  final void Function(int) onItemSelected;
  final int currentIndex;
  final List<NavigationBarTab> tabs;

  const SimpleNavigationBar({
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
          Container(
            decoration: BoxDecoration(
              border: Border(
                top: BorderSide(
                  color: AppColors.licorice.withAlpha(30),
                  width: 1,
                ),
              ),
            ),
            child: BottomBarDefault(
              top: 15,
              bottom: isCompactSize ? 25 : 0,
              items: tabs,
              backgroundColor: Colors.white,
              color: AppColors.licorice.withAlpha(125),
              colorSelected: AppColors.licorice,
              iconSize: 22,
              titleStyle: const TextStyle(fontSize: 13),
              boxShadow: [],
              indexSelected: currentIndex,
              onTap: onItemSelected,
              animated: false,
            ),
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
        ],
      ),
    );
  }
}
