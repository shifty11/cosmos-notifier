import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
import 'package:fixnum/fixnum.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'tracker_row.freezed.dart';

@freezed
class TrackerRow with _$TrackerRow {
  const TrackerRow._();

  const factory TrackerRow({
    required Int64 id,
    required String address,
    required pb.Duration notificationInterval,
    required Int64 chatId,
    required bool isSaved,
    @Default(false) bool isAddRow,
    @Default(true) bool isAddressValid,
  }) = _TrackerRow;

  String get notificationIntervalPrettyString {
    Duration duration = Duration(seconds: notificationInterval.seconds.toInt());
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

  String get shortenedBech32Address {
    if (isAddressValid && address.length > 20) {
      return "${address.substring(0, 8)}...${address.substring(address.length - 4)}";
    }
    return address;
  }
}

@freezed
class TrackerList with _$TrackerList {
  const factory TrackerList({
    required List<TrackerRow> trackers,
  }) = _TrackerList;
}
