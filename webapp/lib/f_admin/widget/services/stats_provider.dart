import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final statsProvider = FutureProvider((ref) async {
  return await adminService.getStats(Empty());
});
