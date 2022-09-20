import 'dart:convert';

import 'package:fixnum/fixnum.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

/// A [Provider] that returns chat-id if it is present as a query parameter.
final chatIdProvider = StateProvider<Int64?>((ref) {
  final chatIdStr = Uri.base.queryParameters.entries.firstWhere((e) => e.key == "chat-id", orElse: () => const MapEntry("chat-id", "")).value;
  if (chatIdStr.isNotEmpty) {
    try {
      final chatId = Int64.parseInt(chatIdStr);
      return chatId;
    } on FormatException {
      // ignore exceptions since the query param could be anything
    }
  }
  final state = Uri.base.queryParameters.entries.firstWhere((e) => e.key == "state", orElse: () => const MapEntry("state", "")).value;
  if (state.isNotEmpty) {
    Codec<String, String> stringToBase64 = utf8.fuse(base64);
    String decoded = stringToBase64.decode(state);
    if (decoded.startsWith("chat-id=")) {
      try {
        final chatId = Int64.parseInt(decoded.substring("chat-id=".length));
        return chatId;
      } on FormatException {
        // ignore exceptions since the query param could be anything
      }
    }
  }
  return null;
});
