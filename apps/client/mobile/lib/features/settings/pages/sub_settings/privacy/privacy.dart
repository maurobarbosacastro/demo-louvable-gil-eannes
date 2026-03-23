import 'package:flutter/material.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';

class Privacy extends StatelessWidget {
  List<String> restrictions(BuildContext context) =>
      [
        translation(context).privacyRestrictionsDescriptionTopic1,
        translation(context).privacyRestrictionsDescriptionTopic2,
        translation(context).privacyRestrictionsDescriptionTopic3,
        translation(context).privacyRestrictionsDescriptionTopic4,
        translation(context).privacyRestrictionsDescriptionTopic5,
        translation(context).privacyRestrictionsDescriptionTopic6,
        translation(context).privacyRestrictionsDescriptionTopic7,
        translation(context).privacyRestrictionsDescriptionTopic8,
      ];

  const Privacy({super.key});

  @override
  Widget build(BuildContext context) {
    return ScrollableBody(
      child: Column(
        spacing: 30,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          PrivacyClause(
            title: translation(context).privacyIntroductionTitle,
            description: translation(context).privacyIntroductionDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyIntellectualPropertyRightsTitle,
            description: translation(context).privacyIntellectualPropertyRightsDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyRestrictionsTitle,
            description: translation(context).privacyRestrictionsDescriptionMain,
            secondary: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                ...restrictions(context).map((item) => Row(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const Text('• ', style: TextStyle(fontSize: 15)),
                    const SizedBox(width: 5),
                    Expanded(child: Text(item, style: Theme.of(context).textTheme.bodyMedium)),
                  ],
                )),
                const SizedBox(height: 10),
                Text(
                  translation(context).privacyRestrictionsDescriptionSecondary,
                  style: Theme.of(context).textTheme.bodyMedium,
                ),
              ],
            ),
          ),
          PrivacyClause(
            title: translation(context).privacyYourContentTitle,
            description: translation(context).privacyYourContentDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyNoWarrantyTitle,
            description: translation(context).privacyNoWarrantyDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyLimitationOfLiabilityTitle,
            description: translation(context).privacyLimitationOfLiabilityDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyIndemnificationTitle,
            description: translation(context).privacyIndemnificationDescription,
          ),
          PrivacyClause(
            title: translation(context).privacySeverabilityTitle,
            description: translation(context).privacySeverabilityDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyVariationOfTermsTitle,
            description: translation(context).privacyVariationOfTermsDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyAssignmentTitle,
            description: translation(context).privacyAssignmentDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyEntireAgreementTitle,
            description: translation(context).privacyEntireAgreementDescription,
          ),
          PrivacyClause(
            title: translation(context).privacyGoverningLawJurisdictionTitle,
            description: translation(context).privacyGoverningLawJurisdictionDescription,
          ),
        ],
      ),
    );
  }
}

class PrivacyClause extends StatelessWidget {
  final String title;
  final String description;
  final Widget? secondary;
  const PrivacyClause({super.key, required this.title, required this.description, this.secondary});

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
        child: Column(
          spacing: 10,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(title, style: TextTheme.of(context).titleLarge),
            Text(description, style: TextTheme.of(context).bodyMedium),
            if (secondary != null) secondary!
          ],
        )
    );
  }
}
