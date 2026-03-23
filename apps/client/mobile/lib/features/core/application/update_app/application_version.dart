import 'dart:convert';
import 'dart:io';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:package_info_plus/package_info_plus.dart';
import 'package:url_launcher/url_launcher.dart';

import '../../../../utils/logging.dart';


class VersionData {

  /// The current version of the app.
  final String localVersion;

  /// The most recent version of the app in the store.
  final String storeVersion;

  /// A link to the app store page where the app can be updated.
  final String appStoreLink;

  /// The release notes for the store version of the app.
  final String? releaseNotes;

  bool get canUpdate {
    final local = localVersion.split('.').map(int.parse).toList();
    final store = storeVersion.split('.').map(int.parse).toList();

    // version is greater than the local version.
    for (var i = 0; i < store.length; i++) {
      // The store version field is newer than the local version.
      if (store[i] > local[i]) return true;
      // The local version field is newer than the store version.
      if (local[i] > store[i]) return false;
    }

    // The local and store versions are the same.
    return false;
  }
  // bool get isRequired {
  //   final local = localVersion.split('.').map(int.parse).toList();
  //   final store = storeVersion.split('.').map(int.parse).toList();
  //   // version is greater than the local version.
  //   if(store.length >= 3){
  //     if(store[0] > local[0] || store[1] > local[1]){
  //       return true;
  //     }
  //     if(store[2] > local[2]){
  //       return false;
  //     }
  //   }
  //   for (var i = 0; i < store.length; i++) {
  //     // The store version field is newer than the local version.
  //     if (store[i] > local[i]) return true;
  //     // The local version field is newer than the store version.
  //     if (local[i] > store[i]) return false;
  //   }
  //
  //   // The local and store versions are the same.
  //   return false;
  // }

  VersionData._({
    required this.localVersion,
    required this.storeVersion,
    required this.appStoreLink,
    this.releaseNotes,
  });
}


class ApplicationVersion {
  ApplicationVersion();

  Future<VersionData?> getVersionStatus() async {
    PackageInfo packageInfo = await PackageInfo.fromPlatform();
    if (Platform.isIOS) {
      return _getiOSStoreVersion(packageInfo);
    } else if (Platform.isAndroid) {
      return _getAndroidStoreVersion(packageInfo);
    } else {
      Log.info('The target platform "${Platform.operatingSystem}" is not yet supported by this package.');
      return null;
    }
  }

  String _getCleanVersion(String version) => RegExp(r'\d+\.\d+(\.\d+)?').stringMatch(version) ?? '0.0.0';

  Future<VersionData?> _getiOSStoreVersion(PackageInfo packageInfo) async {
    final id = packageInfo.packageName;
    final parameters = {"bundleId": id};
    var uri = Uri.https("itunes.apple.com", "/lookup", parameters);
    final response = await http.get(uri);
    if (response.statusCode != 200) {
      Log.info('Failed to query iOS App Store');
      return null;
    }
    final jsonObj = json.decode(response.body);
    final List results = jsonObj['results'];
    if (results.isEmpty) {
      Log.info('Can\'t find an app in the App Store with the id: $id');
      return null;
    }
    return VersionData._(
      localVersion: _getCleanVersion(packageInfo.version),
      storeVersion: _getCleanVersion(jsonObj['results'][0]['version']),
      appStoreLink: jsonObj['results'][0]['trackViewUrl'],
      releaseNotes: jsonObj['results'][0]['releaseNotes'],
    );
  }


  Future<VersionData?> _getAndroidStoreVersion(PackageInfo packageInfo) async {
    final id =  packageInfo.packageName;
    final uri = Uri.https("play.google.com", "/store/apps/details", {"id": id.toString()});
    final response = await http.get(uri);
    if (response.statusCode != 200) throw Exception("Invalid response code: ${response.statusCode}");

    // Supports 1.2.3 (most of the apps) and 1.2.prod.3 (e.g. Google Cloud)
    //final regexp = RegExp(r'\[\[\["(\d+\.\d+(\.[a-z]+)?\.\d+)"\]\]');
    final regexp = RegExp(r'\[\[\[\"(\d+\.\d+(\.[a-z]+)?(\.([^"]|\\")*)?)\"\]\]');
    final storeVersion = regexp.firstMatch(response.body)?.group(1);

    //Description
    //final regexpDescription = RegExp(r'\[\[(null,)\"((\.[a-z]+)?(([^"]|\\")*)?)\"\]\]');

    //Release
    final regexpRelease = RegExp(r'\[(null,)\[(null,)\"((\.[a-z]+)?(([^"]|\\")*)?)\"\]\]');

    final expRemoveSc = RegExp(r"\\u003c[A-Za-z]{1,10}\\u003e", multiLine: true, caseSensitive: true);

    final releaseNotes = regexpRelease.firstMatch(response.body)?.group(3);
    //final descriptionNotes = regexpDescription.firstMatch(response.body)?.group(2);

    return VersionData._(
      localVersion: _getCleanVersion(packageInfo.version),
      storeVersion: _getCleanVersion(storeVersion ?? ""),
      appStoreLink: uri.toString(),
      releaseNotes: releaseNotes?.replaceAll(expRemoveSc, ''),
    );
  }

  openStore() async {
    if (Platform.isIOS) _openAppStore();
    if (Platform.isAndroid) _openPlayStore();
  }

  _launchStore(Uri uri) => launchUrl(uri, mode: LaunchMode.externalApplication);

  _openPlayStore() async {
    final id = PackageInfo.fromPlatform().then((value) => value.packageName);
    final link = ("market://details?id=$id");
    _launchStore(Uri.parse(link));
  }

  _openAppStore() async {
    final packageInfo = await PackageInfo.fromPlatform();
    final parameters = {"bundleId": packageInfo.packageName};
    var uri = Uri.https("itunes.apple.com", "/lookup", parameters);
    final response = await http.get(uri);
    if (response.statusCode != 200) {
      Log.info("Failed to query iOS App Store");
      return null;
    }
    final jsonObj = json.decode(response.body);
    final List results = jsonObj['results'];
    if (results.isEmpty) {
      Log.info('Can\'t find an app in the App Store with the id: ${packageInfo.packageName}');
      return null;
    }
    _launchStore(Uri.parse(jsonObj['results'][0]['trackViewUrl']));
  }

}


final updateApplicationProvider = Provider((ref) => ApplicationVersion());
final showUpdateProvider = StateProvider((ref) => false);
