import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_service.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final trackerFutureProvider = FutureProvider((ref) async {
  await ref.read(trackerNotifierProvider.notifier).getTrackers();
  return ref.read(trackerNotifierProvider);
});

final trackerNotifierProvider = StateNotifierProvider<TrackerNotifier, List<TrackerRow>>((ref) {
  return TrackerNotifier(trackerService, ref.read(messageProvider.notifier));
});

class TrackerNotifier extends StateNotifier<List<TrackerRow>> {
  TrackerService trackerService;
  MessageNotifier messageNotifier;

  TrackerNotifier(this.trackerService, this.messageNotifier) : super([]);

  String notificationIntervalAsString(Int64 seconds) {
    Duration duration = Duration(seconds: seconds.toInt());
    if (duration.inDays > 0) {
      var dayText = "day";
      if (duration.inDays > 1) {
        dayText = "days";
      }
      if (duration.inHours > duration.inDays * 24) {
        return "${duration.inDays} $dayText, ${duration.inHours - duration.inDays * 24} hours before";
      } else {
        return "${duration.inDays} $dayText";
      }
    } else if (duration.inHours > 0) {
      return "${duration.inHours} hours before";
    } else if (duration.inMinutes > 0) {
      return "${duration.inMinutes} minutes before";
    }
    return "on time";
  }

  String shortenedBech32Address(String address) {
    if (address.length > 20) {
      return "${address.substring(0, 8)}...${address.substring(address.length - 4)}";
    }
    return address;
  }

  Future<void> getTrackers() async {
    var response = await trackerService.getTrackers(Empty());
    response.trackers.forEach((tracker) {
      state = [
        ...state,
        TrackerRow(
          id: tracker.id,
          address: shortenedBech32Address(tracker.address),
          notificationInterval: notificationIntervalAsString(tracker.notificationInterval.seconds),
          chatId: tracker.discordChannelId,
          isSaved: true,
        )
      ];
    });
    state = [...state, TrackerRow(id: Int64(0), address: "", notificationInterval: "", chatId: Int64(0), isSaved: false, isAddRow: true)];
  }

  void addTracker() {
    state = [...state, TrackerRow(id: Int64(0), address: "", notificationInterval: "", chatId: Int64(0), isSaved: false)];
  }

  Future<void> deleteTracker(TrackerRow tracker) async {
    try {
      await trackerService.deleteTracker(DeleteTrackerRequest(trackerId: tracker.id));
      state = state.where((element) => element.id != tracker.id).toList();
    } catch (e) {
      messageNotifier.sendMsg(error: e.toString());
    }
  }
}
