import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_service.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;

final trackerFutureProvider = FutureProvider((ref) async {
  await ref.read(trackerNotifierProvider.notifier).getTrackers();
});

final trackerNotifierProvider = StateNotifierProvider<TrackerNotifier, List<TrackerRow>>((ref) {
  return TrackerNotifier(trackerService, ref.read(messageProvider.notifier));
});

class TrackerNotifier extends StateNotifier<List<TrackerRow>> {
  TrackerService trackerService;
  MessageNotifier messageNotifier;

  TrackerNotifier(this.trackerService, this.messageNotifier) : super([]);

  Future<void> getTrackers() async {
    var response = await trackerService.getTrackers(Empty());
    for (var tracker in response.trackers) {
      state = [
        ...state,
        TrackerRow(
          id: tracker.id,
          address: tracker.address,
          notificationInterval: tracker.notificationInterval,
          chatId: tracker.discordChannelId,
          isSaved: true,
        )
      ];
    }
    state = [...state, TrackerRow(id: Int64(0), address: "", notificationInterval: pb.Duration(), chatId: Int64(0), isSaved: false, isAddRow: true)];
  }

  void addTracker() {
    if (state.last.isAddRow) {
      return;
    }
    state = [...state, TrackerRow(id: Int64(0), address: "", notificationInterval: pb.Duration(), chatId: Int64(0), isSaved: false)];
  }

  Future<void> updateTracker(TrackerRow tracker) async {
    if (tracker.isAddRow) {
      if (tracker.address.isNotEmpty) {
        try {
          var response = await trackerService.isAddressValid(IsAddressValidRequest(address: tracker.address));
          tracker = tracker.copyWith(isAddressValid: response.isValid);
          state = [
            for (final oldTrackerRow in state)
              if (oldTrackerRow.id == tracker.id)
                tracker
              else
                oldTrackerRow,
          ];
        } catch (e) {
          messageNotifier.sendMsg(error: e.toString());
          return;
        }
      }
      // if (tracker.address.isNotEmpty && !tracker.notificationInterval.seconds.isZero && tracker.chatId.toInt() > 0) {
      if (tracker.isAddressValid && !tracker.notificationInterval.seconds.isZero) {
        try {
          var response = await trackerService.addTracker(AddTrackerRequest(
            address: tracker.address,
            notificationInterval: tracker.notificationInterval,
            discordChannelId: tracker.chatId,
          ));
          // TODO: fix state
        } catch (e) {
          messageNotifier.sendMsg(error: e.toString());
        }
      }
    }
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
