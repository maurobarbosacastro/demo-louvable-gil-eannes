import 'package:flutter/material.dart';
import 'package:carousel_slider/carousel_slider.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class Carousel extends StatefulWidget {
  const Carousel(
      {super.key,
      required this.carouselItems,
      this.carouselHeight = 315,
      this.carouselStepperInactiveColor = Colors.white,
      this.carouselStepperAlignment = MainAxisAlignment.center,
      this.onPageChanged,
      this.carouselCurrentPage = 0,
      this.carouselStepperWidth = 20,
      required this.carouselController});

  final List<Widget> carouselItems;
  final double carouselHeight;
  final Color carouselStepperInactiveColor;
  final MainAxisAlignment carouselStepperAlignment;
  final double carouselStepperWidth;
  final void Function(int)? onPageChanged;
  final int carouselCurrentPage;
  final CarouselSliderController carouselController;

  @override
  State<Carousel> createState() => _CarouselState();
}

class _CarouselState extends State<Carousel> {
  late int _current;
  late CarouselSliderController _carouselController;

  int? _pendingPage;

  @override
  void initState() {
    super.initState();
    _current = widget.carouselCurrentPage;
    _carouselController = widget.carouselController;
  }

  @override
  Widget build(BuildContext context) {
    final List<Widget> carouselItems = widget.carouselItems;

    return Stack(
      children: [
        CarouselSlider(
          carouselController: _carouselController,
          options: CarouselOptions(
            initialPage: _current,
            viewportFraction: 1,
            height: widget.carouselHeight,
            enableInfiniteScroll: false,
            onPageChanged: (index, reason) {
              if (_pendingPage == null || _pendingPage == index) {
                setState(() {
                  _current = index;
                });
                _pendingPage = null;
                widget.onPageChanged?.call(index);
              }
            }
          ),
          items: carouselItems.map((carouselItem) {
            return Builder(
              builder: (BuildContext context) {
                return Stack(
                  children: [
                    SizedBox(
                      width: MediaQuery.of(context).size.width,
                      child: carouselItem,
                    ),
                  ],
                );
              },
            );
          }).toList(),
        ),
        Positioned(
          bottom: 20.0,
          left: 0,
          right: 0,
          child: Center(
            child: Row(
              mainAxisAlignment: widget.carouselStepperAlignment,
              children: carouselItems.asMap().entries.map((entry) {
                int index = entry.key;
                return GestureDetector(
                  child: Container(
                    width: widget.carouselStepperWidth,
                    height: 4,
                    margin: const EdgeInsets.symmetric(horizontal: 2),
                    color: _current == index
                        ? AppColors.youngBlue
                        : widget.carouselStepperInactiveColor.withAlpha(51),
                  ),
                  onTap: () {
                    _pendingPage = index;
                    _carouselController.animateToPage(
                      index,
                      duration: const Duration(milliseconds: 500),
                      curve: Curves.easeInOut,
                    );
                  },
                );
              }).toList(),
            ),
          ),
        ),
      ],
    );
  }
}
