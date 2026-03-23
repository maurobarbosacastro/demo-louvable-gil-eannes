import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../routes/route_names.dart';
import 'hive_cache.dart';
import 'http_interceptor.dart';

final GlobalKey<ScaffoldMessengerState> scaffoldKey = GlobalKey<ScaffoldMessengerState>();

final scaffoldKeyProvider = Provider((ref) => scaffoldKey);

final hiveProvider = Provider((ref) => HiveDatabase());

final lifecycleRouteProvider = StateProvider((ref) => RouteNames.splashRouteLocation);

final httpInterceptorProvider = Provider((ref) => ApiInterceptor(ref.watch(hiveProvider), ref));
