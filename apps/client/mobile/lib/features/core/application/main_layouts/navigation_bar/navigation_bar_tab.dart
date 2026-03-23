import 'package:awesome_bottom_bar/awesome_bottom_bar.dart';

class NavigationBarTab extends TabItem {
  const NavigationBarTab({
    required this.initialLocation,
    required super.icon,
    required super.title
  });

  final String initialLocation;
}
