import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/models/user_payment_method_model/user_payment_method_model.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class ViewPaymentMethodDialog extends ConsumerWidget {
  final UserPaymentMethodModel payment;

  const ViewPaymentMethodDialog({super.key, required this.payment});

  String formatIban(String iban) {
    String formattedValue = '';
    for (int i = 0; i < iban.length; i++) {
      if (i > 0 && i % 4 == 0) {
        formattedValue += ' ';
      }
      formattedValue += iban[i];
    }
    return formattedValue;
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Padding(
      padding: const EdgeInsets.only(top: 60, bottom: 0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            spacing: 5,
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Flexible(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      '${payment.paymentMethod} ${translation(context).account}',
                      style: Theme.of(context).textTheme.headlineMedium,
                    ),
                    // i want to wrap this because if not my button exits the screen and gives me overflow error
                    Text(
                      formatIban(payment.iban),
                      style: Theme.of(context).textTheme.bodySmall,
                      softWrap: true,
                      maxLines: 2,
                    ),
                  ],
                ),
              ),
              Button(
                  label: translation(context).delete,
                  mode: Mode.outlinedRed,
                  icon: SvgPicture.asset(Assets.trash,
                      colorFilter: const ColorFilter.mode(AppColors.radicalRed, BlendMode.srcIn)),
                  onPressed: () {})
            ],
          ),
          const SizedBox(height: 10),
          Container(
            height: 1,
            color: AppColors.wildSand,
          ),
          const SizedBox(height: 30),
          Column(
            spacing: 30,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              PaymentMethodDetail(
                  label: translation(context).bankName, value: payment.bankName),
              PaymentMethodDetail(
                  label: translation(context).bankAddress,
                  value: payment.bankAddress),
              PaymentMethodDetail(
                  label: translation(context).country, value: payment.country),
              PaymentMethodDetail(
                  label: translation(context).bankAccountTitle,
                  value: payment.bankAccountTitle),
              PaymentMethodDetail(
                  label: translation(context).fiscalNumber, value: payment.vat),
              PaymentMethodDetail(
                  label: translation(context).ibanAccountNumber,
                  value: formatIban(payment.iban)),
            ],
          ),
          const SizedBox(height: 30),
          SizedBox(
            width: MediaQuery.of(context)
                .size
                .width,
            child: Row(
              mainAxisSize: MainAxisSize.max,
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Text(
                  translation(context).ibanConfirmationStatement,
                  style: Theme.of(context).textTheme.titleSmall,
                ),
                Container(
                  margin: const EdgeInsets.only(right: 0),
                  decoration: BoxDecoration(
                    color: AppColors.youngBlue,
                    borderRadius: BorderRadius.circular(4),
                  ),
                  padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                  constraints: const BoxConstraints(maxWidth: 150),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      SvgPicture.asset(Assets.checkUnderlined,
                          colorFilter: const ColorFilter.mode(Colors.white, BlendMode.srcIn), height: 15),
                      const SizedBox(width: 8),
                      Flexible(
                        fit: FlexFit.loose,
                        child: Text(
                          payment.ibanStatement.fileName,
                          softWrap: true,
                          overflow: TextOverflow.visible,
                          style: Theme.of(context)
                              .textTheme
                              .titleSmall
                              ?.copyWith(color: Colors.white),
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          )
        ],
      ),
    );
  }
}

class PaymentMethodDetail extends StatelessWidget {
  final String label;
  final String value;

  const PaymentMethodDetail(
      {super.key, required this.label, required this.value});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          label,
          style: Theme.of(context).textTheme.titleSmall,
        ),
        Text(
          value,
          style: Theme.of(context)
              .textTheme
              .bodyMedium
              ?.copyWith(color: AppColors.licorice),
        ),
      ],
    );
  }
}
