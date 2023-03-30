import 'package:collection/collection.dart';
import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/empty.pb.dart';
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/state/validator_bundle.dart';
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

final hasValidationErrorProvider = Provider<bool>((ref) {
  return ref.watch(trackerNotifierProvider.select((trackerRows) {
    final trackerRow = trackerRows.firstWhereOrNull((trackerRow) => !trackerRow.isSaved && !trackerRow.isAddressValid);
    return trackerRow != null;
  }));
});

final showChatRoomColumnProvider = Provider<bool>((ref) {
  return ref.watch(trackerChatRoomsProvider).length > 1;
});

final trackerSortProvider = StateProvider<TrackerSortState>((ref) {
  return const TrackerSortState(isAscending: true, sortType: TrackerSortType.address);
});

final validatorBundleProvider = StateNotifierProvider<ValidatorBundleNotifier, List<FreezedValidatorBundle>>((ref) {
  return ValidatorBundleNotifier(ref);
});

class ValidatorBundleNotifier extends StateNotifier<List<FreezedValidatorBundle>> {
  StateNotifierProviderRef<ValidatorBundleNotifier, List<FreezedValidatorBundle>> ref;

  ValidatorBundleNotifier(this.ref) : super([]) {
    state = ref.watch(trackerFutureProvider).when(
          data: (data) {
            return data.validatorBundles.map((e) => FreezedValidatorBundle.fromProtobuf(e)).toList()..sort();
          },
          loading: () => [],
          error: (_, __) => [],
        );
  }

  void selectValidatorBundles(List<FreezedValidatorBundle> added, List<FreezedValidatorBundle> deleted) {
    final monikerAdded = added.map((e) => e.moniker);
    final monikerDeleted = deleted.map((e) => e.moniker);
    state = state.map((bundle) {
      if (monikerAdded.contains(bundle.moniker)) {
        return bundle.copyWith(isTracked: true);
      }
      if (monikerDeleted.contains(bundle.moniker)) {
        return bundle.copyWith(isTracked: false);
      }
      return bundle;
    }).toList()
      ..sort();
  }
}

class TrackerNotifier extends StateNotifier<List<TrackerRow>> {
  TrackerService trackerService;
  MessageNotifier messageNotifier;
  StateNotifierProviderRef<TrackerNotifier, List<TrackerRow>> ref;

  TrackerNotifier(this.trackerService, this.messageNotifier, this.ref) : super([]);

  _sortByAddress(bool isAscending) {
    var sorted = state.where((t) => t.isSaved).sorted((a, b) => a.address.compareTo(b.address));
    if (!isAscending) {
      sorted = sorted.reversed.toList();
    }
    state = [...sorted, ...state.where((t) => !t.isSaved)];
  }

  _sortByNotificationInterval(bool isAscending) {
    var sorted = state.where((t) => t.isSaved).sorted((a, b) => a.notificationInterval.seconds.compareTo(b.notificationInterval.seconds));
    if (!isAscending) {
      sorted = sorted.reversed.toList();
    }
    state = [...sorted, ...state.where((t) => !t.isSaved)];
  }

  _sortByChatRoom(bool isAscending) {
    var sorted = state.where((t) => t.isSaved).sorted((a, b) => a.chatRoom!.name.compareTo(b.chatRoom!.name));
    if (!isAscending) {
      sorted = sorted.reversed.toList();
    }
    state = [...sorted, ...state.where((t) => !t.isSaved)];
  }

  sort() {
    final sortState = ref.read(trackerSortProvider);
    switch (sortState.sortType) {
      case TrackerSortType.none:
        break;
      case TrackerSortType.address:
        _sortByAddress(sortState.isAscending);
        break;
      case TrackerSortType.notificationInterval:
        _sortByNotificationInterval(sortState.isAscending);
        break;
      case TrackerSortType.chatRoom:
        _sortByChatRoom(sortState.isAscending);
        break;
    }
  }

  TrackerRow? getLastModifiedTrackerRow() {
    final sortedTrackers = state.where((trackerRow) => trackerRow.isSaved).toList()
      ..sort((a, b) => b.updatedAt?.compareTo(a.updatedAt ?? DateTime.fromMicrosecondsSinceEpoch(0)) ?? -1);
    if (sortedTrackers.isEmpty) {
      return null;
    }
    return sortedTrackers.first;
  }

  /// Returns the default notification interval for a new tracker.
  /// Uses the latest modified tracker's notification interval or 1 day if there are no trackers.
  pb.Duration getDefaultNotificationInterval() {
    final lastModifiedTracker = getLastModifiedTrackerRow();
    if (lastModifiedTracker == null || lastModifiedTracker.updatedAt == null) {
      return pb.Duration(seconds: Int64(60 * 60 * 24)); // 1 day by default
    }
    return lastModifiedTracker.notificationInterval;
  }

  TrackerChatRoom? getDefaultChatRoom() {
    final lastModifiedTracker = getLastModifiedTrackerRow();
    if (lastModifiedTracker != null && lastModifiedTracker.chatRoom != null) {
      return lastModifiedTracker.chatRoom;
    }
    if (state.isNotEmpty) {
      return state.first.chatRoom;
    }
    return ref.watch(trackerChatRoomsProvider).firstOrNull;
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
    sort();
    return response;
  }

  void addTracker() {
    state = [
      ...state,
      TrackerRow(
        id: Int64(0),
        address: "",
        notificationInterval: getDefaultNotificationInterval(),
        chatRoom: getDefaultChatRoom(),
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
          messageNotifier.sendMsg(info: "Reminder added");
          ref.read(trackerSortProvider.notifier).state = ref.read(trackerSortProvider).copyWith(sortType: TrackerSortType.none);
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
        messageNotifier.sendMsg(info: "Reminder updated");
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
    return _updateTrackerRow(
        tracker.copyWith(
          id: response.id,
          address: response.address,
          notificationInterval: response.notificationInterval,
          chatRoom: response.chatRoom,
          updatedAt: response.updatedAt.toDateTime(),
        ),
        isNewTracker: isNewTracker);
  }

  Future<void> deleteTracker(TrackerRow tracker) async {
    if (!tracker.isSaved) {
      state = state.where((element) => element.id != tracker.id).toList();
      return;
    }
    try {
      await trackerService.deleteTracker(DeleteTrackerRequest(trackerId: tracker.id));
      state = state.where((element) => element.id != tracker.id).toList();
      messageNotifier.sendMsg(info: "Reminder deleted");
    } catch (e) {
      messageNotifier.sendMsg(error: e.toString());
    }
  }

  trackValidators(
    List<FreezedValidatorBundle> toBeTracked,
    List<FreezedValidatorBundle> toBeAdded,
    List<FreezedValidatorBundle> toBeDeleted,
    TrackerChatRoom? chatRoom,
    Duration? notificationInterval,
  ) async {
    try {
      final monikers = toBeTracked.map((e) => e.moniker).toList();
      final response = await trackerService.trackValidators(TrackValidatorsRequest(
        monikers: monikers,
        notificationInterval: notificationInterval != null ? pb.Duration(seconds: Int64(notificationInterval.inSeconds)) : null,
        chatRoom: chatRoom,
      ));
      state = [
        ...state.whereNot((trackerRow) => response.deletedTrackerIds.contains(trackerRow.id)),
        ...response.addedTrackers.map((tracker) => TrackerRow(
              id: tracker.id,
              address: tracker.address,
              notificationInterval: tracker.notificationInterval,
              chatRoom: tracker.chatRoom,
              updatedAt: tracker.updatedAt.toDateTime(),
            ))
      ];
      ref.read(validatorBundleProvider.notifier).selectValidatorBundles(toBeAdded, toBeDeleted);
      // TODO: show message
    } catch (e) {
      messageNotifier.sendMsg(error: e.toString());
    }
  }
}
