import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:tagpeak/shared/models/store_model/public_stores/public_store_model.dart';

final isHiddenUntilNextMonthProvider = StateProvider((ref) => false);

class VisitStoreNotifier {
  final StateController<PublicStoreModel?> visitStoreController;

  VisitStoreNotifier({required this.visitStoreController});

  void saveVisitStore(PublicStoreModel visitStore) {
    visitStoreController.update((state) => visitStore);
  }

  void removeVisitStore() {
    visitStoreController.update((state) => null);
  }
}

final visitStoreStateProvider = StateProvider<PublicStoreModel?>((ref) => null);

final visitStoreNotifierProvider = Provider((ref) => VisitStoreNotifier(
  visitStoreController: ref.watch(visitStoreStateProvider.notifier),
));
