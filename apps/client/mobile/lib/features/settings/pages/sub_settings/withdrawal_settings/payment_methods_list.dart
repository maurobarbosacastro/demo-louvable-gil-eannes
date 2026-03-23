import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_svg/svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/dashboard/store_visits/without_results.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/withdrawal_settings/dialogs/delete_payment_method_dialog.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/withdrawal_settings/dialogs/view_payment_method_dialog.dart';
import 'package:tagpeak/shared/models/user_payment_method_model/user_payment_method_model.dart';
import 'package:tagpeak/shared/services/payment_method_service.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_bottom_dialog.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_dialog.dart';
import 'package:tagpeak/shared/widgets/snack_bar/snack_bar.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class PaymentMethodsList extends ConsumerStatefulWidget {
  final List<UserPaymentMethodModel> payments;

  const PaymentMethodsList({super.key, this.payments = const []});

  @override
  ConsumerState<PaymentMethodsList> createState() => _PaymentMethodsListState();
}

class _PaymentMethodsListState extends ConsumerState<PaymentMethodsList> {
  late List<UserPaymentMethodModel> payments;

  @override
  void initState() {
    super.initState();
    payments = List.from(widget.payments);
  }

  @override
  void didUpdateWidget(covariant PaymentMethodsList oldWidget) {
    super.didUpdateWidget(oldWidget);

    if (oldWidget.payments != widget.payments) {
      setState(() {
        payments = List.from(widget.payments);
      });
    }
  }

  void _deletePaymentMethod(String uuid, BuildContext ctx) {
    ref.read(paymentMethodServiceProvider).deletePaymentMethod(uuid).then((_) {
      setState(() {
        payments = payments.where((payment) => payment.uuid != uuid).toList();
      });

      if (!ctx.mounted) return;
      GenericDialog.closeDialog(ctx);

      showSnackBar(context, translation(context).deletePaymentMethodSuccess, SnackBarType.success);
    }).catchError((error) {
      showSnackBar(context, translation(context).somethingWentWrong, SnackBarType.error);
    });
  }

  @override
  Widget build(BuildContext context) {
    if (payments.isEmpty) {
      return WithoutResults(
        title: translation(context).noPaymentAdded,
        description: translation(context).noPaymentAddedDescription,
      );
    }

    return Column(
      children: payments.asMap().entries.map((entry) {
        return PaymentMethodCard(
          index: entry.key,
          payment: entry.value,
          handleDeleteClick: (String uuid) {
            GenericDialog.openDialog(context,
                content: DeletePaymentMethodDialog(
              onDelete: (BuildContext ctx) => _deletePaymentMethod(uuid, ctx),
            ));
          },
          handleViewClick: (payment) {
            GenericBottomDialog.openBottomDialog(
              context,
              content: ViewPaymentMethodDialog(payment: payment),
              hasCloseIcon: true,
              padding: const EdgeInsets.only(right: 20, left: 20, top: 20, bottom: 40),
            );
          },
        );
      }).toList(),
    );
  }
}

class PaymentMethodCard extends StatelessWidget {
  final int index;
  final UserPaymentMethodModel payment;
  final void Function(String) handleDeleteClick;
  final void Function(UserPaymentMethodModel) handleViewClick;

  const PaymentMethodCard({
    super.key,
    this.index = 0,
    required this.payment,
    required this.handleDeleteClick,
    required this.handleViewClick,
  });

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
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => handleViewClick(payment),
      child: Container(
        decoration: const BoxDecoration(
          border: Border(bottom: BorderSide(color: AppColors.wildSand, width: 1)),
        ),
        padding: const EdgeInsets.only(top: 10, bottom: 15),
        margin: EdgeInsets.only(top: index == 0 ? 0 : 20),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  '${payment.paymentMethod} ${translation(context).account}',
                  style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: AppColors.licorice),
                ),
                Text(
                  formatIban(payment.iban),
                  style: Theme.of(context).textTheme.titleSmall,
                ),
              ],
            ),
            Row(
              children: [
                GestureDetector(
                  child: SvgPicture.asset(Assets.trash),
                  onTap: () => handleDeleteClick(payment.uuid),
                ),
                const SizedBox(width: 10),
                GestureDetector(
                  child: SvgPicture.asset(Assets.simpleArrowRight),
                  onTap: () => handleViewClick(payment),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
