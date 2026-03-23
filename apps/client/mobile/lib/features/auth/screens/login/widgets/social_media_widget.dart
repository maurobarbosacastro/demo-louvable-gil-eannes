import 'package:tagpeak/features/auth/screens/login/provider/login_provider.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_fonts/google_fonts.dart';

class SocialMediaWidget extends StatelessWidget {
  const SocialMediaWidget({super.key});

  static final List socialMedia = ["Facebook", "Google"];

  String getAsset(String media) {
    switch (media.toLowerCase()) {
      case "facebook":
        return Assets.facebook;
      case "instagram":
        return Assets.instagram;
      case "apple":
        return Assets.apple;
      case "google":
        return Assets.google;
      case "twitter":
        return Assets.twitter;
      default:
        return "";
    }
  }

  Widget horizontalLine() => Container(height: 1.0, color: const Color(0xFFF0EDFF));

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 85,
      width: MediaQuery.of(context).size.width,
      child: Column(
        children: [
          Consumer(builder: (context, ref, child) {
            return Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: List.generate(
                socialMedia.length,
                    (index) => TextButton(
                  style: TextButton.styleFrom(foregroundColor: const Color(0xFF606060), shape: const CircleBorder()),
                  onPressed: () => ref.read(loginNotifier).signInWithSocial(socialMedia[index].toString().toLowerCase()),
                  child: Image.asset(getAsset(socialMedia[index]), width: 42, height: 42),
                ),
              ),
            );
          }),
        ],
      ),
    );
  }
}

