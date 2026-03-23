import 'package:carousel_slider/carousel_controller.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/carousel_slider.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class OnBoardingDialog extends ConsumerStatefulWidget {
  final void Function() onClickStartShopping;

  const OnBoardingDialog({super.key, required this.onClickStartShopping});

  @override
  ConsumerState<OnBoardingDialog> createState() => _OnBoardingDialogState();
}

class _OnBoardingDialogState extends ConsumerState<OnBoardingDialog> {
  late int _currentPage = 0;
  final CarouselSliderController _carouselSliderController = CarouselSliderController();

  void nextPage() {
    if (_currentPage == 4) {
      widget.onClickStartShopping();
      return;
    }

    _carouselSliderController.nextPage(
      duration: const Duration(milliseconds: 500),
      curve: Curves.easeInOut,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 20),
      child: Carousel(
        carouselController: _carouselSliderController,
        carouselStepperInactiveColor: AppColors.youngBlue,
        carouselStepperAlignment: MainAxisAlignment.start,
        carouselHeight: 648,
        carouselStepperWidth: 30,
        onPageChanged: (index) {
          setState(() {
            _currentPage = index;
          });
        },
        carouselItems: [
          Padding(
            padding: const EdgeInsets.only(top: 63),
            child: OnBoardingSlide(
              imagePath: Assets.onBoardingSlide1,
              title: translation(context).onBoardingSlide1Title,
              description: translation(context).onBoardingSlide1Description,
              button: Button(
                label: translation(context).onBoardingSlide1Button,
                onPressed: nextPage,
                width: 199.4,
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 0),
            child: OnBoardingSlide(
              imagePath: Assets.onBoardingSlide2,
              title: translation(context).onBoardingSlide2Title,
              description: translation(context).onBoardingSlide2Description,
              button: Button(
                label: translation(context).gotIt,
                onPressed: nextPage,
                width: 93.4,
                mode: Mode.outlinedDark
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 63.0),
            child: OnBoardingSlide(
              imagePath: Assets.onBoardingSlide3,
              title: translation(context).onBoardingSlide3Title,
              description: translation(context).onBoardingSlide3Description,
              button: Button(
                label: translation(context).gotIt,
                onPressed: nextPage,
                width: 93.4,
                mode: Mode.outlinedDark
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 63.0),
            child: OnBoardingSlide(
              imagePath: Assets.onBoardingSlide4,
              title: translation(context).onBoardingSlide4Title,
              description: translation(context).onBoardingSlide4Description,
              button: Button(
                label: translation(context).gotIt,
                onPressed: nextPage,
                width: 93.4,
                mode: Mode.outlinedDark
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 63.0, bottom: 20),
            child: OnBoardingSlide(
              imagePath: Assets.onBoardingSlide5,
              title: translation(context).onBoardingSlide5Title,
              description: translation(context).onBoardingSlide5Description,
              button: Button(
                icon: SvgPicture.asset(Assets.checkUnderlined, colorFilter: const ColorFilter.mode(Colors.white, BlendMode.srcIn), height: 15, width: 15),
                label: translation(context).onBoardingSlide5Button,
                onPressed: nextPage,
                width: 163.4,
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class OnBoardingSlide extends StatelessWidget {
  final String imagePath;
  final String title;
  final String description;
  final Button button;
  const OnBoardingSlide({super.key, required this.imagePath, required this.title, required this.description, required this.button});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Image.asset(imagePath),
        const SizedBox(height: 32),
        Text(
          title,
          style: Theme.of(context).textTheme.headlineMedium,
        ),
        const SizedBox(height: 10),
        Text(
          description,
          style: Theme.of(context).textTheme.bodyMedium,
        ),
        const SizedBox(height: 10),
        button
      ],
    );
  }
}
