import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart';
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
    required TrackerChatRoom? chatRoom,
    required DateTime? updatedAt,
    @Default(true) bool isAddressValid,
  }) = _TrackerRow;

  bool get isSaved => id != Int64(0);

  String _singularOrPlural(int number, String singular, String plural) {
    if (number == 1) {
      return singular;
    }
    return plural;
  }

  String get notificationIntervalPrettyString {
    Duration duration = Duration(seconds: notificationInterval.seconds.toInt());
    if (duration.inDays > 0) {
      if (duration.inHours > duration.inDays * 24) {
        final hours = duration.inHours - duration.inDays * 24;
        return "${duration.inDays} "
            "${_singularOrPlural(duration.inDays, "day", "days")}, "
            "$hours ${_singularOrPlural(hours, "hour", "hours")} before";
      } else {
        return "${duration.inDays} ${_singularOrPlural(duration.inDays, "day", "days")}";
      }
    } else if (duration.inHours > 0) {
      return "${duration.inHours} ${_singularOrPlural(duration.inHours, "hour", "hours")} before";
    } else if (duration.inMinutes > 0) {
      return "${duration.inMinutes} ${_singularOrPlural(duration.inMinutes, "minute", "minutes")} before";
    }
    return "on time";
  }

  String shortenedAddress(bool veryShort) {
    final maxLength = veryShort ? 12 : 20;
    if (isAddressValid && address.length > maxLength) {
      return "${address.substring(0, maxLength ~/ 2)}...${address.substring(address.length - maxLength ~/ 2)}";
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
