import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../../core/routes/route_names.dart';
import '../widgets/tile_button_widget.dart';

class OptionsListComponent extends StatelessWidget {
  const OptionsListComponent({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        TileButton(
          label: "Edit user",
          onTap: () => context.pushNamed(RouteNames.editUserRouteName),
          icon: Icons.edit,
        ),
        TileButton(
          label: "Terms and Politics",
          onTap: () => context.pushNamed(RouteNames.termsRouteName),
          icon: Icons.contact_page,
        )
      ],
    );
  }
}
