import 'dart:io';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:reactive_forms/reactive_forms.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/withdrawal_settings/create_payment_method/payment_method_form.dart';
import 'package:tagpeak/shared/models/create_payment_method_model/create_payment_method_model.dart';
import 'package:tagpeak/shared/services/country_service.dart';
import 'package:tagpeak/shared/services/payment_method_service.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/dialogs/generic_bottom_dialog.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/shared/widgets/snack_bar/snack_bar.dart';

class CreatePaymentMethod extends ConsumerStatefulWidget {
  final void Function() onDialogClose;

  const CreatePaymentMethod({super.key, required this.onDialogClose});

  @override
  ConsumerState<CreatePaymentMethod> createState() => _CreatePaymentMethodState();
}

class _CreatePaymentMethodState extends ConsumerState<CreatePaymentMethod> {
  late List<DropdownOption> countries = [];
  late List<DropdownOption> paymentMethods = [];
  final FormGroup formGroup = FormGroup({
    'paymentMethod': FormControl<String>(validators: [Validators.required]),
    'bankName': FormControl<String>(validators: [Validators.required]),
    'bankAddress': FormControl<String>(validators: [Validators.required]),
    'bankCountry': FormControl<String>(validators: [Validators.required]),
    'bankAccountTitle': FormControl<String>(validators: [Validators.required]),
    'fiscalNumber': FormControl<String>(validators: [Validators.required]),
    'iban': FormControl<String>(validators: [
      Validators.required,
      Validators.minLength(25),
      Validators.pattern(r'^[A-Z0-9 ]+$'),
    ]),
    'ibanStatement': FormControl<PlatformFile>(validators: [Validators.required]),
    'country': FormControl<String>(validators: [Validators.required])
  });

  void _fetchCountries() async {
    await ref.read(countryServiceProvider).getCountries().then((pagination) {
      if (!mounted) return;
      setState(() {
        countries = pagination.data.map((country) {
          return DropdownOption(label: country.name, value: country.uuid);
        }).toList();
      });
    });
  }

  void _fetchPaymentMethods() async {
    await ref.read(paymentMethodServiceProvider).getPaymentMethodsAvailable().then((paymentMethods) {
      if (!mounted) return;
      setState(() {
        this.paymentMethods = paymentMethods.map((paymentMethod) {
          return DropdownOption(label: paymentMethod.name, value: paymentMethod.uuid, disabled: paymentMethod.code == 'paypal', selected: paymentMethod.code == 'bank');
        }).toList();

        final bankPaymentMethods = paymentMethods.where((element) => element.code == 'bank');

        if (bankPaymentMethods.isNotEmpty) {
          final paymentMethodOption = bankPaymentMethods.first;
          formGroup.control('paymentMethod').value = paymentMethodOption.uuid;
        }
      });
    });
  }

  void finalize() async {
    if (!formGroup.valid) {
      formGroup.markAllAsTouched();
      setState(() {});
      return;
    }

    CreatePaymentMethodModel paymentMethod = CreatePaymentMethodModel(
      paymentMethod: formGroup.control('paymentMethod').value,
      bankName: formGroup.control('bankName').value,
      bankAddress: formGroup.control('bankAddress').value,
      bankCountry: formGroup.control('bankCountry').value,
      country: formGroup.control('country').value,
      vat: formGroup.control('fiscalNumber').value,
      bankAccountTitle: formGroup.control('bankAccountTitle').value,
      iban: formGroup
          .control('iban')
          .value
          .toString()
          .replaceAll(RegExp(r'\s'), ''),
      ibanStatement: '',
      state: '',
    );

    try {
      final platformFile = formGroup.control('ibanStatement').value as PlatformFile?;
      final filePath = platformFile?.path;

      if (filePath == null) {
        return;
      }

      final File file = File(filePath);

      final String fileId = await ref.read(paymentMethodServiceProvider).loadIbanStatement(file);

      paymentMethod = paymentMethod.copyWith(ibanStatement: fileId.replaceAll('"', '').trim());

      await ref.read(paymentMethodServiceProvider).createPaymentMethod(paymentMethod).then((_) {
        if (!context.mounted) return;
        GenericBottomDialog.closeBottomDialog(context);

        widget.onDialogClose();
      }).catchError((error) {
        showSnackBar(context, translation(context).somethingWentWrong, SnackBarType.error);
      });
    } catch (e) {
      showSnackBar(context, translation(context).somethingWentWrong, SnackBarType.error);
    }
  }

  @override
  void initState() {
    _fetchCountries();
    _fetchPaymentMethods();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: double.infinity,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 40),
          Text(
            translation(context).createNewPaymentMethod,
            style: Theme.of(context).textTheme.headlineMedium,
          ),
          Text(
            translation(context).fillOutPaymentDetails,
            style: Theme.of(context).textTheme.bodyMedium,
          ),
          const SizedBox(height: 20),
          Expanded(
            child: SingleChildScrollView(
              padding: const EdgeInsets.only(
                  bottom: 10,
                  top: 10
              ),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  PaymentMethodForm(countries: countries, paymentMethods: paymentMethods, formGroup: formGroup),
                ],
              ),
            ),
          ),
          const SizedBox(height: 20),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Button(
                label: translation(context).cancel,
                mode: Mode.outlinedDark,
                onPressed: () => context.pop(),
              ),
              Button(
                label: translation(context).add,
                onPressed: () => finalize(),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
