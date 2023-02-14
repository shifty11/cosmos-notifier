import 'package:cosmos_notifier/f_home/services/state/auth_state.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final cannySSOProvider = StateProvider<CannySSO>((ref) {
  final companyID = Uri.base.queryParameters.entries.firstWhere((e) => e.key == "companyID", orElse: () => const MapEntry("", "")).value;
  if (companyID.isNotEmpty) {
    final redirect = Uri.base.queryParameters.entries.firstWhere((e) => e.key == "redirect", orElse: () => const MapEntry("", "")).value;
    if (redirect.isNotEmpty) {
      return CannySSO(true, "", redirect, companyID);
    }
  }
  return const CannySSO(false, "", "", "");
});
