import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:tagpeak/features/core/application/translation.dart';
import 'package:tagpeak/shared/widgets/button.dart';
import 'package:tagpeak/shared/widgets/dropdowns/dropdown.dart';
import 'package:tagpeak/utils/constants/assets.dart';
import 'package:tagpeak/utils/constants/colors.dart';

class StoresFilters extends StatelessWidget {
  final List<DropdownOption> countries;
  final void Function(String?) onSelectCountry;
  final void Function() onClearCountry;

  final List<DropdownOption> categories;
  final void Function(String?) onSelectCategory;
  final void Function() onClearCategory;

  final TextEditingController nameController;
  final void Function() onSearch;

  final String? sortValue;
  final void Function(String?) onChangeSortDropdown;

  StoresFilters(
      {Key? key,
      required this.countries,
      required this.onSelectCountry,
      required this.onClearCountry,
      required this.categories,
      required this.onSelectCategory,
      required this.onClearCategory,
      required this.nameController,
      this.sortValue,
      required this.onChangeSortDropdown,
      required this.onSearch})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      spacing: 10,
      children: [
        // ToDO change this to searchable dropdown
        Row(
          spacing: 5,
          children: [
            Expanded(
              child: Dropdown(
                isExpanded: true,
                entries: countries,
                label: translation(context).chooseCountry,
                onChanged: (String? value) => onSelectCountry(value),
                onClickClear: (_) => onClearCountry(),
              ),
            ),
            Expanded(
              child: Dropdown(
                isExpanded: true,
                entries: categories,
                label: translation(context).chooseCategory,
                onChanged: (String? value) => onSelectCategory(value),
                onClickClear: (_) => onClearCategory(),
              ),
            ),
          ],
        ),

        Row(
          spacing: 5,
          children: [
            Expanded(
              child: TextFormField(
                controller: nameController,
                decoration: InputDecoration(
                  enabledBorder: OutlineInputBorder(
                    borderSide: BorderSide(
                      color: AppColors.waterloo.withAlpha(153),
                    ),
                  ),
                  focusedBorder: OutlineInputBorder(
                    borderSide: BorderSide(
                      color: AppColors.waterloo.withAlpha(153),
                    ),
                  ),
                  floatingLabelBehavior: FloatingLabelBehavior.never,
                  hintText: translation(context).searchByName,
                  hintStyle: Theme.of(context).textTheme.bodyMedium,
                  contentPadding: const EdgeInsets.symmetric(vertical: 5, horizontal: 10),
                  isCollapsed: true,
                ),
              ),
            ),

            Button(
                label: translation(context).search,
                width: 111.5,
                icon: SvgPicture.asset(Assets.search,
                    width: 14, colorFilter: ColorFilter.mode(AppColors.licorice.withAlpha(204), BlendMode.srcIn)),
                onPressed: onSearch, mode: Mode.outlinedDark)
          ],
        ),

        Row(
          spacing: 10,
          children: [
            Text(translation(context).sortBy,
                style: Theme.of(context)
                    .textTheme
                    .bodyMedium
                    ?.copyWith(color: AppColors.licorice)),
            DropdownButtonHideUnderline(
              child: DropdownButton(
                value: sortValue,
                items: [
                  DropdownMenuItem(
                      value: 'created_at desc',
                      child: Text(translation(context).latestStores,
                          style: Theme.of(context)
                              .textTheme
                              .bodyMedium
                              ?.copyWith(color: AppColors.licorice))),
                  DropdownMenuItem(
                      value: 'name',
                      child: Text(translation(context).byTitle,
                          style: Theme.of(context)
                              .textTheme
                              .bodyMedium
                              ?.copyWith(color: AppColors.licorice))),
                  DropdownMenuItem(
                      value: 'most-popular',
                      child: Text(translation(context).mostPopular,
                          style: Theme.of(context)
                              .textTheme
                              .bodyMedium
                              ?.copyWith(color: AppColors.licorice)))
                ],
                onChanged: onChangeSortDropdown
              ),
            ),
          ],
        )
      ],
    );
  }
}
