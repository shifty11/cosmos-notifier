import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
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
});

final trackerNotifierProvider = StateNotifierProvider<TrackerNotifier, List<TrackerRow>>((ref) {
  return TrackerNotifier(trackerService, ref.read(messageProvider.notifier));
});

final showAddTrackerButtonProvider = Provider((ref) {
  return ref.watch(trackerNotifierProvider.notifier).hasUnsavedChanges;
});

class TrackerNotifier extends StateNotifier<List<TrackerRow>> {
  TrackerService trackerService;
  MessageNotifier messageNotifier;

  TrackerNotifier(this.trackerService, this.messageNotifier) : super([]);

  pb.Duration _getDefaultNotificationInterval() {
    var lastModifiedTracker = state.where((trackerRow) => trackerRow.updatedAt != null).reduce((value, element) {
      if (value.updatedAt != null && element.updatedAt != null && value.updatedAt!.isBefore(element.updatedAt!)) {
        return value;
      }
      return element;
    });
    if (lastModifiedTracker.updatedAt == null) {
      return pb.Duration(seconds: Int64(60 * 60 * 24)); // 1 day by default
    }
    return lastModifiedTracker.notificationInterval;
  }

  bool get hasUnsavedChanges {
    return state.where((trackerRow) => !trackerRow.isSaved).isNotEmpty;
  }

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
          updatedAt: tracker.updatedAt.toDateTime(),
        )
      ];
    }
  }

  void addTracker() {
    if (hasUnsavedChanges) {
      return;
    }
    state = [
      ...state,
      TrackerRow(id: Int64(0), address: "", notificationInterval: _getDefaultNotificationInterval(), chatId: Int64(0), updatedAt: null),
    ];
  }

  Future<void> updateTracker(TrackerRow tracker) async {
    if (!tracker.isSaved) {
      if (tracker.address.isNotEmpty) {
        try {
          var response = await trackerService.isAddressValid(IsAddressValidRequest(address: tracker.address));
          tracker = tracker.copyWith(isAddressValid: response.isValid);
        } catch (e) {
          messageNotifier.sendMsg(error: e.toString());
          return;
        }
      }
      // if (tracker.address.isNotEmpty && !tracker.notificationInterval.seconds.isZero && tracker.chatId.toInt() > 0) {
      if (tracker.address.isNotEmpty && tracker.isAddressValid && !tracker.notificationInterval.seconds.isZero) {
        try {
          var response = await trackerService.addTracker(AddTrackerRequest(
            address: tracker.address,
            notificationInterval: tracker.notificationInterval,
            discordChannelId: tracker.chatId,
          ));
          // TODO: fix state
        } catch (e) {
          messageNotifier.sendMsg(error: e.toString());
          return;
        }
      }
      state = _updateTrackerRow(tracker);
    } else {
      try {
        await trackerService.updateTracker(UpdateTrackerRequest(
          trackerId: tracker.id,
          notificationInterval: tracker.notificationInterval,
          discordChannelId: tracker.chatId,
          // TODO: telegram chat id

        ));
        state = _updateTrackerRow(tracker);
      } catch (e) {
        messageNotifier.sendMsg(error: e.toString());
      }
    }
  }

  List<TrackerRow> _updateTrackerRow(TrackerRow tracker) {
    return [
      for (final oldTrackerRow in state)
        if (oldTrackerRow.id == tracker.id) tracker else oldTrackerRow,
    ];
  }

  Future<void> deleteTracker(TrackerRow tracker) async {
    if (!tracker.isSaved) {
      state = state.where((element) => element.id != tracker.id).toList();
      return;
    }
    try {
      await trackerService.deleteTracker(DeleteTrackerRequest(trackerId: tracker.id));
      state = state.where((element) => element.id != tracker.id).toList();
    } catch (e) {
      messageNotifier.sendMsg(error: e.toString());
    }
  }
}
