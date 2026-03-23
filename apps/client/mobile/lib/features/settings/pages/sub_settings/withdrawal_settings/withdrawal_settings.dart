import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/core/application/main_layouts/scrollable_body.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/withdrawal_settings/create_payment_card.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/withdrawal_settings/create_payment_method/create_payment_method.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/withdrawal_settings/payment_methods_list.dart';
import 'package:tagpeak/shared/models/user_payment_method_model/user_payment_method_model.dart';
import 'package:tagpeak/shared/services/payment_method_service.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_bottom_dialog.dart';
import 'package:tagpeak/shared/widgets/snack_bar/snack_bar.dart';

class WithdrawalSettings extends ConsumerStatefulWidget {
  const WithdrawalSettings({super.key});

  @override
  ConsumerState<WithdrawalSettings> createState() => _WithdrawalSettingsState();
}

class _WithdrawalSettingsState extends ConsumerState<WithdrawalSettings> {
  List<UserPaymentMethodModel> payments = [];

  @override
  void initState() {
    super.initState();

    _fetchPaymentMethods();
  }

  void _fetchPaymentMethods() {
    ref.read(paymentMethodServiceProvider).getPaymentMethods().then((response) {
      setState(() => payments = response.data);
    });
  }

  void handleAddPaymentMethodClick(BuildContext context) {
    GenericBottomDialog.openBottomDialog(
        context,
        content: CreatePaymentMethod(
          onDialogClose: () {
            _fetchPaymentMethods();
            showSnackBar(context, translation(context).createPaymentMethodSuccess, SnackBarType.success);
          },
        ),
        hasCloseIcon: true,
        padding: const EdgeInsets.only(right: 20, left: 20, top: 20, bottom: 40),
    );
  }

  @override
  Widget build(BuildContext context) {
    return ScrollableBody(
        child: Column(
          spacing: 30,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            CreatePaymentCard(onAddPaymentMethodPressed: () => handleAddPaymentMethodClick(context)),
            PaymentMethodsList(payments: payments)
          ],
        )
    );
  }
}
