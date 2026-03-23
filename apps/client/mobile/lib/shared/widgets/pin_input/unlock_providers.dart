import 'package:flutter_riverpod/flutter_riverpod.dart';

final pinValuesProvider = StateProvider<List<String>>((ref) => []);
final pinShowValuesProvider = StateProvider<List<bool>>((ref) => []);
final attemptsProvider = StateProvider<int>((ref) => 0);
final pinErrorProvider = StateProvider<bool>((ref) => false);



