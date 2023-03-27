import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart';
import 'package:fixnum/fixnum.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'tracker_row.freezed.dart';

enum AddressSize {
  veryShort,
  short,
  long,
}

@freezed
class TrackerRow with _$TrackerRow {
  const TrackerRow._();

  const factory TrackerRow({
    required Int64 id,
    required String address,
    required pb.Duration notificationInterval,
    required TrackerChatRoom? chatRoom,
    required DateTime? updatedAt,
    @Default(true) bool isAddressValid,
  }) = _TrackerRow;

  bool get isSaved => id != Int64(0);

  String get notificationIntervalPrettyString {
    Duration duration = Duration(seconds: notificationInterval.seconds.toInt());
    if (duration.inDays > 0) {
      if (duration.inHours > duration.inDays * 24) {
        final hours = duration.inHours - duration.inDays * 24;
        return "${duration.inDays}d ${hours}h";
      } else {
        return "${duration.inDays}d";
      }
    } else if (duration.inHours > 0) {
      return "${duration.inHours}h";
    } else if (duration.inMinutes > 0) {
      return "${duration.inMinutes}m";
    }
    return "on time";
  }

  String shortenedAddress(AddressSize addressSize) {
    final maxLength = addressSize == AddressSize.veryShort ? 8 : addressSize == AddressSize.short ? 12 : 16;
    if (isAddressValid && address.length > maxLength) {
      return "${address.substring(0, maxLength ~/ 2)}...${address.substring(address.length - maxLength ~/ 2)}";
    }
    return address;
  }
}

// needs same order as in UI (none is not used in UI so it's at the end)
enum TrackerSortType {
  address,
  notificationInterval,
  chatRoom,
  none,
}

@freezed
class TrackerSortState with _$TrackerSortState {
  const factory TrackerSortState({
    required bool isAscending,
    required TrackerSortType sortType,
  }) = _TrackerSortState;

  const TrackerSortState._();

  factory TrackerSortState.initial() => const TrackerSortState(isAscending: false, sortType: TrackerSortType.none);
}