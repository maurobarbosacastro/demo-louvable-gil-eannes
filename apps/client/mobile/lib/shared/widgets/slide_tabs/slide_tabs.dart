import 'package:flutter/material.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class SlideTabs extends StatefulWidget {
  final List<Tab> tabs;
  final List<Widget> views;
  final bool resetOnSwitch;
  final int initialIndex;

  const SlideTabs({
    super.key,
    required this.tabs,
    required this.views,
    this.resetOnSwitch = true,
    this.initialIndex = 0
  });

  @override
  State<SlideTabs> createState() => _SlideTabsState();
}

class _SlideTabsState extends State<SlideTabs> with SingleTickerProviderStateMixin {
  late TabController _tabController;
  int _currentIndex = 0;

  @override
  void initState() {
    super.initState();
    _currentIndex = widget.initialIndex;
    _tabController = TabController(length: widget.tabs.length, vsync: this, initialIndex: widget.initialIndex);
    _tabController.addListener(() {
      if (_tabController.indexIsChanging) {
        setState(() {
          _currentIndex = _tabController.index;
        });
      }
    });
  }

  @override
  void dispose() {
    _tabController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Column(
        children: [
          TabBar(
            controller: _tabController,
            tabs: widget.tabs,
            dividerHeight: 1,
            indicatorColor: Colors.transparent,
            splashFactory: NoSplash.splashFactory,
            isScrollable: true,
            tabAlignment: TabAlignment.start,
            dividerColor: AppColors.wildSand,
            labelColor: AppColors.licorice,
            overlayColor: WidgetStateProperty.all(Colors.transparent),
            labelStyle: const TextStyle(fontSize: 16),
            labelPadding: const EdgeInsets.only(right: 20, bottom: 5),
            unselectedLabelColor: AppColors.grey,
          ),
          Expanded(
            child: widget.resetOnSwitch
                ? IndexedStack(
              index: _currentIndex,
              children: List.generate(
                widget.views.length,
                    (index) => _currentIndex == index
                    ? KeyedSubtree(
                  key: ValueKey('tab_$index'),
                  child: widget.views[index],
                )
                    : Container(),
              ),
            )
                : TabBarView(
              controller: _tabController,
              children: widget.views,
            ),
          ),
        ],
      ),
    );
  }
}
