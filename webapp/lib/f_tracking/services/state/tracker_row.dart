import 'package:fixnum/fixnum.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'tracker_row.freezed.dart';

@freezed
class TrackerRow with _$TrackerRow {
  const factory TrackerRow({
    required Int64 id,
    required String address,
    required String notificationInterval,
    required Int64 chatId,
    required bool isSaved,
    @Default(false) bool isAddRow,
  }) = _TrackerRow;
}

@freezed
class TrackerList with _$TrackerList {
  const factory TrackerList({
    required List<TrackerRow> trackers,
  }) = _TrackerList;
}
