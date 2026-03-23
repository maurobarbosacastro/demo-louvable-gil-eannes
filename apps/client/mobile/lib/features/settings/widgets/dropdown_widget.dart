import 'package:tagpeak/features/settings/widgets/tile_button_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class DropdownWidget extends ConsumerWidget {
  final Widget child;
  final String label;
  final IconData iconData;

  const DropdownWidget({
    super.key,
    required this.child,
    required this.label,
    required this.iconData,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Column(
      children: [
        TileButton(
          label: label,
          onTap: () {
            ref.read(showSecurityProvider.notifier).update((state) => !state);
          },
          icon: iconData,
          endIcon: ref.watch(showSecurityProvider) ? Icons.arrow_upward : Icons.arrow_downward,
        ),
        Visibility(
          visible: ref.watch(showSecurityProvider),
          child: Padding(padding: const EdgeInsets.symmetric(horizontal: 16), child: child),
        )
      ],
    );
  }
}

final showSecurityProvider = StateProvider((ref) => false);
final enableBiometricsProvider = StateProvider((ref) => false);
