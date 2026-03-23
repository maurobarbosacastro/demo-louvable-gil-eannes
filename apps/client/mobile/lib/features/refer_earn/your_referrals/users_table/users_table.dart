import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/referral_info_model/user_referral_model/user_referral_model.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';
import 'package:tagpeak/utils/extensions/currency_extension.dart';

class UsersTable extends StatelessWidget {
  final List<UserReferralModel>? users;
  final String currency;

  const UsersTable({super.key, this.users, required this.currency});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(20),
        border: Border.all(color: AppColors.wildSand),
      ),
      constraints: const BoxConstraints(
        maxHeight: 400,
      ),
      child: Column(
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 14, horizontal: 20),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Row(
                  spacing: 10,
                  children: [
                    SvgPicture.asset(Assets.star),
                    Text(translation(context).referrals, style: Theme.of(context).textTheme.titleSmall),
                  ],
                ),
                Text((users?.length ?? 0).toString(), style: Theme.of(context).textTheme.headlineMedium),
              ],
            ),
          ),
          Expanded(
            child: ListView.separated(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              itemCount: users?.length ?? 0,
              separatorBuilder: (context, index) => const Divider(color: AppColors.wildSand),
              itemBuilder: (context, index) {
                final user = users![index];
                return Padding(
                  padding: const EdgeInsets.symmetric(vertical: 10),
                  child: Row(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      ClipRRect(
                        borderRadius: BorderRadius.circular(100),
                        child: FadeInImage.assetNetwork(
                          fadeInDuration: const Duration(milliseconds: 1),
                          placeholder: Assets.defaultTgPkImage,
                          image: user.profilePicture,
                          fit: BoxFit.cover,
                          width: 33,
                          height: 33,
                          imageErrorBuilder: (context, error, stackTrace) =>
                              Image.asset(Assets.defaultTgPkImage,
                                  fit: BoxFit.cover, width: 33, height: 33),
                        ),
                      ),
                      const SizedBox(width: 10),
                      Expanded(
                        child: Text(
                          user.displayName ?? '${user.firstName} ${user.lastName}',
                          style: Theme.of(context).textTheme.bodyMedium,
                        ),
                      ),
                      Text(
                        user.firstTransactionSuccessful
                            ? user.referredValue.toCurrency(currency)
                            : '-',
                        style: Theme.of(context).textTheme.titleLarge,
                      ),
                    ],
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
