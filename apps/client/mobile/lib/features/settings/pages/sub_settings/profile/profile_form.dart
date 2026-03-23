import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/features/auth/screens/login/widgets/password_form_field_widget.dart';
import 'package:tagpeak/features/auth/screens/login/widgets/text_form_field_widget.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/features/settings/pages/sub_settings/profile/text_field_suffix.dart';
import 'package:tagpeak/shared/services/country_service.dart';
import 'package:tagpeak/shared/widgets/date_picker/date_picker.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/shared/widgets/dropdowns/searchable_dropdown.dart';
import 'package:tagpeak/utils/date_helpers.dart';
import 'package:tagpeak/utils/models/form_error_model/form_error_model.dart';

class ProfileFormProps {
  bool readOnly;
  TextEditingController controller;
  FormErrorModel? error;

  ProfileFormProps({Key? key, this.readOnly = false, required this.controller, this.error});
}

enum SuffixAction {
  edit, save, close
}

class ProfileForm extends ConsumerStatefulWidget {
  final Map<String, ProfileFormProps> configs;
  final void Function(String field, SuffixAction suffixAction) onSuffixAction;
  final GlobalKey<FormState> formKey;

  const ProfileForm(
      {super.key, required this.configs, required this.onSuffixAction, required this.formKey});

  @override
  ConsumerState<ProfileForm> createState() => _ProfileFormState();
}

class _ProfileFormState extends ConsumerState<ProfileForm> {
  List<DropdownOption>? dropdownItems;

  @override
  void initState() {
    super.initState();
    _fetchCountries();
  }

  void _fetchCountries() async {
    await ref.read(countryServiceProvider).getCountries().then((pagination) {
      if (!mounted) return;
      setState(() {
        dropdownItems = pagination.data.map((country) {
          return DropdownOption(label: country.name, value: country.abbreviation);
        }).toList();
      });
    });
  }

  DateTime? stringToDateTime() {
    String input = widget.configs['dateOfBirth']!.controller.text;

    if (input.isNotEmpty && (input.contains('-') || input.contains('/'))) {
      final delimiter = input.contains('-') ? '-' : '/';
      List<String> parts = input.split(delimiter);

      return DateTime(
        int.parse(parts[2]),
        int.parse(parts[1]),
        int.parse(parts[0]),
      );
    }

    return null;
  }

  Widget _getSuffixWidget(String field) {
    return TextFieldSuffix(
      inEditMode: !widget.configs[field]!.readOnly,
      handleClose: () => widget.onSuffixAction(field, SuffixAction.close),
      handleEdit: () => widget.onSuffixAction(field, SuffixAction.edit),
      handleSave: () => widget.onSuffixAction(field, SuffixAction.save),
    );
  }

  @override
  Widget build(BuildContext context) {
    final Map<String, ProfileFormProps> configs = widget.configs;

    return Form(
      key: widget.formKey,
      child: Column(
        children: [
          // Name
          TextFormFieldWidget(
            formatters: const [],
            controller: configs['name']!.controller,
            validator: (value) => value!.toString().trim().isEmpty ? translation(context).nameNotEmpty : null,
            placeholder: '',
            label: translation(context).name,
            textInputType: TextInputType.name,
            readOnly: configs['name']!.readOnly,
            suffix: _getSuffixWidget('name')
          ),
          // Display Name
          TextFormFieldWidget(
            controller: configs['displayName']!.controller,
            validator: (value) => value!.toString().trim().isEmpty ? translation(context).displayNameNotEmpty : null,
            placeholder: '',
            label: translation(context).displayName,
            textInputType: TextInputType.name,
            readOnly: configs['displayName']!.readOnly,
            suffix: _getSuffixWidget('displayName')
          ),
          // Password
          PasswordFormFieldWidget(
            controller: configs['password']!.controller,
            validator: (value) => value!.toString().trim().isEmpty ? translation(context).passwordNotEmpty : null,
            placeholder: translation(context).yourPassword,
            obscureText: true,
            label: translation(context).password,
            suffixIcon: false,
            readOnly: configs['password']!.readOnly,
          ),
          // Email Address
          TextFormFieldWidget(
            controller: configs['email']!.controller,
            validator: (value) => value!.toString().trim().isEmpty ? translation(context).emailNotEmpty : null,
            placeholder: '',
            label: translation(context).email,
            textInputType: TextInputType.emailAddress,
            readOnly: configs['email']!.readOnly,
          ),
          const SizedBox(height: 10),
          DatePicker(
            controller: configs['dateOfBirth']!.controller,
            selectedDate: stringToDateTime(),
            label: translation(context).dateBirth,
            placeholder: translation(context).noDate,
            onSelectedDate: (date) => setState(() {
              configs['dateOfBirth']?.controller.text =
                  dateTimeToString(date);
            }),
            readOnly: configs['dateOfBirth']!.readOnly,
            suffix: _getSuffixWidget('dateOfBirth'),
            type: InputType.underline,
          ),
          const SizedBox(height: 25),
          SearchableDropdown(
            controller: configs['country']!.controller,
            label: translation(context).country,
            items: dropdownItems ?? [],
            hint: translation(context).noCountry,
            value: configs['country']?.controller.text,
            onChanged: (value) {},
            readOnly: configs['country']!.readOnly,
            suffix: _getSuffixWidget('country')
          ),
        ],
      ),
    );
  }
}
