import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import 'application_version.dart';
import 'widgets/update_widgets.dart';

class UpdateApplication extends ConsumerStatefulWidget {
  final Widget child;
  final SnackBar? snackBar;
  final AlertDialog? alertDialog;
  final VoidCallback? onUpdate;
  final VoidCallback? onCancel;
  final bool isRequired;

  const UpdateApplication({super.key, required this.child, this.snackBar, this.alertDialog, this.onUpdate, this.onCancel, this.isRequired = false});

  @override
  ConsumerState<UpdateApplication> createState() => _UpdateApplicationState();
}

class _UpdateApplicationState extends ConsumerState<UpdateApplication> {
  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    checkVersion(ref.watch(updateApplicationProvider));
  }

  checkVersion(ApplicationVersion newVersion) async {
    final dismissProvider = ref.watch(showUpdateProvider);
    final status = await newVersion.getVersionStatus();
    final VoidCallback onTap = widget.onUpdate ?? () => ref.watch(updateApplicationProvider).openStore();
    final VoidCallback onCancel = widget.onCancel ??
        () {
          context.pop();
          ref.read(showUpdateProvider.notifier).update((state) => true);
        };
    if (status != null && status.canUpdate && !dismissProvider) {
      if (mounted) {
        if (widget.isRequired) {
          showDialog(context: context, builder: (context) => UpdateWidgets(context).dialog(onPressed: onTap, version: status.storeVersion, onCancel: onCancel));
        } else {
          ScaffoldMessenger.of(context).showSnackBar(
            UpdateWidgets(context).snackBar(
              onTap,
              widget.onCancel ??
                  () {
                    ScaffoldMessenger.of(context).clearSnackBars();
                    ref.read(showUpdateProvider.notifier).update((state) => true);
                  },
            ),
          );
        }
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return widget.child;
  }
}
