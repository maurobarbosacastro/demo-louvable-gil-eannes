import 'package:carousel_slider/carousel_slider.dart';
import 'package:flutter/material.dart';
import 'package:tagpeak/features/stores/stores/store_details/widgets/carousel/carousel_step_item.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class CarouselStepper extends StatefulWidget {
  final List<String> steps;

  final CarouselSliderController carouselController = CarouselSliderController();

  CarouselStepper({super.key, required this.steps});

  @override
  State<CarouselStepper> createState() => _CarouselStepperState();
}

class _CarouselStepperState extends State<CarouselStepper> {
  late CarouselSliderController _carouselController;
  int _currentIndex = 0;

  @override
  void initState() {
    super.initState();
    _carouselController = widget.carouselController;
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 30),
      decoration: const BoxDecoration(color: AppColors.youngBlue),
      child: Column(
        spacing: 20,
        mainAxisSize: MainAxisSize.min,
        children: [
          Container(
            width: double.infinity,
            constraints: const BoxConstraints(
              maxHeight: 70,
            ),
            child: CarouselSlider(
              carouselController: _carouselController,
              items: [
                ...List.generate(widget.steps.length, (index) {
                  return CarouselStepItem(index: index + 1, label: widget.steps[index]);
                })
              ],
              options: CarouselOptions(
                  initialPage: 0,
                  viewportFraction: 1,
                  enableInfiniteScroll: false,
                  onPageChanged: (index, _) => setState(() => _currentIndex = index)
              ),
            ),
          ),
          Row(
            spacing: 10,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              ...List.generate(widget.steps.length, (index) {
                return GestureDetector(
                  onTap: () => _carouselController.animateToPage(
                    index,
                    duration: const Duration(milliseconds: 400),
                    curve: Curves.easeInOut,
                  ),
                  child: Container(
                    height: 10,
                    width: 10,
                    decoration: BoxDecoration(
                      shape: BoxShape.circle,
                      color: _currentIndex == index ? AppColors.grandis : AppColors.wildSand,
                    ),
                  ),
                );
              })
            ]
          )
        ],
      ),
    );
  }
}
