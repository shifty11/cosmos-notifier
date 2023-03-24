import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_service.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final trackerFutureProvider = FutureProvider<GetTrackersResponse>((ref) async {
  return await ref.read(trackerNotifierProvider.notifier).getTrackers();
});

final trackerNotifierProvider = StateNotifierProvider<TrackerNotifier, List<TrackerRow>>((ref) {
  return TrackerNotifier(trackerService, ref.read(messageProvider.notifier), ref);
});

final showAddTrackerButtonProvider = Provider((ref) {
  return !ref.watch(trackerNotifierProvider.select((trackerRows) => trackerRows.any((trackerRow) => !trackerRow.isSaved)));
});

final trackerChatRoomsProvider = Provider<List<TrackerChatRoom>>((ref) {
  return ref.watch(trackerFutureProvider).when(
        data: (data) {
          return data.chatRooms;
        },
        loading: () => [],
        error: (_, __) => [],
      );
});

class TrackerNotifier extends StateNotifier<List<TrackerRow>> {
  TrackerService trackerService;
  MessageNotifier messageNotifier;
  StateNotifierProviderRef<TrackerNotifier, List<TrackerRow>> ref;

  TrackerNotifier(this.trackerService, this.messageNotifier, this.ref) : super([]);

  TrackerRow? _getLastModifiedTrackerRow() {
    final sortedTrackers = state
        .where((trackerRow) => trackerRow.isSaved)
        .toList()
      ..sort((a, b) => b.updatedAt?.compareTo(a.updatedAt ?? DateTime.fromMicrosecondsSinceEpoch(0)) ?? -1);
    if (sortedTrackers.isEmpty) {
      return null;
    }
    return sortedTrackers.first;
  }

  /// Returns the default notification interval for a new tracker.
  /// Uses the latest modified tracker's notification interval or 1 day if there are no trackers.
  pb.Duration _getDefaultNotificationInterval() {
    final lastModifiedTracker = _getLastModifiedTrackerRow();
    if (lastModifiedTracker == null || lastModifiedTracker.updatedAt == null) {
      return pb.Duration(seconds: Int64(60 * 60 * 24)); // 1 day by default
    }
    return lastModifiedTracker.notificationInterval;
  }

  TrackerChatRoom? _getDefaultChatRoom() {
    final lastModifiedTracker = _getLastModifiedTrackerRow();
    if (lastModifiedTracker != null && lastModifiedTracker.chatRoom != null) {
      return lastModifiedTracker.chatRoom;
    }
    if (state.isNotEmpty) {
      return state.first.chatRoom;
    }
    return null;
  }

  Future<GetTrackersResponse> getTrackers() async {
    var response = await trackerService.getTrackers(Empty());
    for (var tracker in response.trackers) {
      state = [
        ...state,
        TrackerRow(
          id: tracker.id,
          address: tracker.address,
          notificationInterval: tracker.notificationInterval,
          chatRoom: tracker.chatRoom,
          updatedAt: tracker.updatedAt.toDateTime(),
        )
      ];
    }
    return response;
  }

  void addTracker() {
    state = [
      ...state,
      TrackerRow(
        id: Int64(0),
        address: "",
        notificationInterval: _getDefaultNotificationInterval(),
        chatRoom: _getDefaultChatRoom(),
        updatedAt: null,
      ),
    ];
  }

  Future<void> updateTracker(TrackerRow tracker) async {
    if (!tracker.isSaved) {
      if (tracker.address.isNotEmpty) {
        if (tracker.address.length < 32) {
          tracker = tracker.copyWith(isAddressValid: false);
        } else {
          try {
            var response = await trackerService.isAddressValid(IsAddressValidRequest(address: tracker.address));
            tracker = tracker.copyWith(isAddressValid: response.isValid);
          } catch (e) {
            messageNotifier.sendMsg(error: e.toString());
            return;
          }
        }
      }
      if (tracker.address.isNotEmpty && tracker.isAddressValid && tracker.chatRoom != null) {
        try {
          var response = await trackerService.addTracker(AddTrackerRequest(
            address: tracker.address,
            notificationInterval: tracker.notificationInterval,
            chatRoom: tracker.chatRoom,
          ));
          state = _updateTrackerRowByResponse(tracker, response, isNewTracker: true);
          messageNotifier.sendMsg(info: "Tracker added");
        } catch (e) {
          messageNotifier.sendMsg(error: e.toString());
          return;
        }
      }
      state = _updateTrackerRow(tracker);
    } else {
      try {
        final response = await trackerService.updateTracker(UpdateTrackerRequest(
          trackerId: tracker.id,
          notificationInterval: tracker.notificationInterval,
          chatRoom: tracker.chatRoom,
        ));
        state = _updateTrackerRowByResponse(tracker, response);
        messageNotifier.sendMsg(info: "Tracker updated");
      } catch (e) {
        messageNotifier.sendMsg(error: e.toString());
      }
    }
  }

  List<TrackerRow> _updateTrackerRow(TrackerRow tracker, {bool? isNewTracker}) {
    return [
      for (final oldTrackerRow in state)
        if (isNewTracker != null && isNewTracker && oldTrackerRow.id == Int64(0))
          tracker
        else if (oldTrackerRow.id == tracker.id)
          tracker
        else
          oldTrackerRow,
    ];
  }

  List<TrackerRow> _updateTrackerRowByResponse(TrackerRow tracker, Tracker response, {bool? isNewTracker}) {
    return _updateTrackerRow(tracker.copyWith(
      id: response.id,
      address: response.address,
      notificationInterval: response.notificationInterval,
      chatRoom: response.chatRoom,
      updatedAt: response.updatedAt.toDateTime(),
    ), isNewTracker: isNewTracker);
  }

  Future<void> deleteTracker(TrackerRow tracker) async {
    if (!tracker.isSaved) {
      state = state.where((element) => element.id != tracker.id).toList();
      return;
    }
    try {
      await trackerService.deleteTracker(DeleteTrackerRequest(trackerId: tracker.id));
      state = state.where((element) => element.id != tracker.id).toList();
      messageNotifier.sendMsg(info: "Tracker deleted");
    } catch (e) {
      messageNotifier.sendMsg(error: e.toString());
    }
  }
}
