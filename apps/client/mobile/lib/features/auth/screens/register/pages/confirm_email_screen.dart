import 'package:carousel_slider/carousel_controller.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/shared/widgets/carousel_slider.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import '../../../../core/application/translation.dart';

class ConfirmEmailScreen extends StatelessWidget {
  final bool isPasswordReset;

  const ConfirmEmailScreen({super.key, this.isPasswordReset = false});

  @override
  Widget build(BuildContext context) {
    final CarouselSliderController carouselSliderController = CarouselSliderController();

    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Visibility(
              visible: !isPasswordReset,
              child: Carousel(
                carouselController: carouselSliderController,
                carouselItems: [
                  ClipRRect(
                    borderRadius: BorderRadius.circular(10.0),
                    child: Image.asset(
                      Assets.dataTransferGraphic,
                      fit: BoxFit.fitWidth,
                      width: double.infinity,
                    ),
                  ),
                  ClipRRect(
                    borderRadius: BorderRadius.circular(10.0),
                    child: Image.asset(
                      Assets.dataTransferGraphic,
                      fit: BoxFit.fitWidth,
                      width: double.infinity,
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 90.0),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 11.0, vertical: 10.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SvgPicture.asset(Assets.checkUnderlined),
                  const SizedBox(height: 25.0),
                  Text(
                    translation(context).checkYourEmail,
                    style: Theme.of(context).textTheme.headlineMedium,
                  ),
                  const SizedBox(height: 10.0),
                  Text(
                    !isPasswordReset ? translation(context).checkYourEmailDescription : translation(context).checkYourEmailToResetPassword,
                    style: Theme.of(context).textTheme.bodyMedium,
                  ),
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}
