import 'package:flutter/material.dart';
import 'package:reactive_forms/reactive_forms.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/shared/widgets/file_uploader/file_uploader.dart';
import 'package:tagpeak/shared/widgets/inputs/tgpk_reactive_text_field.dart';

enum EPaymentMethod {
  bank, paypal
}

class PaymentMethodForm extends StatelessWidget {
  final List<DropdownOption> countries;
  final List<DropdownOption> paymentMethods;
  final FormGroup formGroup;

  const PaymentMethodForm({super.key, required this.countries, required this.paymentMethods, required this.formGroup});

  String? _getDropdownErrors(String field, BuildContext context) {
    final control = formGroup.control(field);
    if (!(control.touched || control.dirty)) {
      return null;
    }

    final errors = control.errors;
    if (errors['required'] == true) {
      return translation(context).fieldRequired;
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    return ReactiveForm(
      formGroup: formGroup,
      child: Column(
        spacing: 20,
        children: [
          // Payment Method
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(translation(context).selectPaymentMethod, style: Theme.of(context).textTheme.titleSmall),
              Dropdown(
                entries: paymentMethods,
                onChanged: (String? value) {
                  formGroup.control('paymentMethod').value = value;
                },
                label: translation(context).selectPaymentMethod,
                isExpanded: true,
                borderStyle: DropdownBorderStyle.underline,
                hasClearIcon: false
              )
            ]
          ),
          // Bank Name
          TgpkReactiveTextField(
            formControl: formGroup.control('bankName') as FormControl<String>,
            label: translation(context).bankName,
            validationMessages: {
              'required': (error) => translation(context).fieldRequired
            },
          ),
          TgpkReactiveTextField(
            formControl: formGroup.control('bankAddress') as FormControl<String>,
            label: translation(context).bankAddress,
            validationMessages: {
              'required': (error) => translation(context).fieldRequired
            },
          ),
          Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(translation(context).bankCountry, style: Theme.of(context).textTheme.titleSmall),
                Dropdown(
                    entries: countries,
                    onChanged: (String? value) {
                      formGroup.control('bankCountry').value = value;
                      formGroup.control('bankCountry').markAsTouched();

                      if (value != null) formGroup.control('bankCountry').removeError('required');
                    },
                    onClickClear: (_) => formGroup.control('bankCountry').value = null,
                    label: translation(context).selectBankCountry,
                    isExpanded: true,
                    borderStyle: DropdownBorderStyle.underline,
                    errorText: _getDropdownErrors('bankCountry', context)
                )
              ]
          ),
          TgpkReactiveTextField(
            formControl: formGroup.control('bankAccountTitle') as FormControl<String>,
            label: translation(context).bankAccountTitle,
            validationMessages: {
              'required': (error) => translation(context).fieldRequired
            },
          ),
          Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(translation(context).country, style: Theme.of(context).textTheme.titleSmall),
                Dropdown(
                    entries: countries,
                    onChanged: (String? value) {
                      formGroup.control('country').value = value;
                      formGroup.control('country').markAsTouched();

                      if (value != null) formGroup.control('country').removeError('required');
                    },
                    onClickClear: (_) => formGroup.control('country').value = null,
                    label: translation(context).selectCountry,
                    isExpanded: true,
                    borderStyle: DropdownBorderStyle.underline,
                    errorText: _getDropdownErrors('country', context)
                )
              ]
          ),
          TgpkReactiveTextField(
            formControl: formGroup.control('fiscalNumber') as FormControl<String>,
            label: translation(context).fiscalNumber,
            validationMessages: {
              'required': (error) => translation(context).fieldRequired
            },
          ),
          TgpkReactiveTextField(
            formControl: formGroup.control('iban') as FormControl<String>,
            label: translation(context).ibanAccountNumber,
            validationMessages: {
              'required': (error) => translation(context).fieldRequired,
              'minLength': (error) => translation(context).minLengthIban,
              'pattern': (error) => translation(context).invalidIban,
            },
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              ConstrainedBox(
                constraints: const BoxConstraints(maxWidth: 150),
                child: Text(
                  translation(context).ibanStatement,
                  softWrap: true,
                  style: Theme.of(context).textTheme.titleSmall,
                ),
              ),
              ConstrainedBox(
                constraints: const BoxConstraints(maxWidth: 170),
                child: FileUploader(
                  onUpload: (ibanStatement) => formGroup.control('ibanStatement').value = ibanStatement,
                  onRemove: () => formGroup.control('ibanStatement').value = null,
                ),
              )
            ],
          ),
        ]
      ),
    );
  }
}
