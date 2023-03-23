import 'package:collection/collection.dart';
import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pbgrpc.dart';
import 'package:cosmos_notifier/common/header_widget.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/bottom_navigation_bar_widget.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/footer_widget.dart';
import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_provider.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:debounce_throttle/debounce_throttle.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_messages/riverpod_messages.dart';

class TrackingPage extends StatelessWidget {
  static const iconSize = 24.0;
  static const iconSizeSmall = 12.0;

  const TrackingPage({Key? key, errorCode}) : super(key: key);

  Widget notificationIntervalDialog(BuildContext context, TrackerRow trackerRow, WidgetRef ref) {
    final duration = Duration(seconds: trackerRow.notificationInterval.seconds.toInt());
    var days = duration.inDays;
    var hours = duration.inHours - days * 24;
    return AlertDialog(
      title: const Text("Notification"),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          const Text("When should the notification be sent before the proposal ends?"),
          const SizedBox(height: 10),
          Row(
            children: [
              const SizedBox(width: 50, child: Text("Days:")),
              const SizedBox(width: 10),
              SizedBox(
                width: 100,
                child: TextFormField(
                  initialValue: days.toString(),
                  keyboardType: TextInputType.number,
                  decoration: const InputDecoration(
                    border: InputBorder.none,
                    enabledBorder: InputBorder.none,
                    hintText: "Days",
                  ),
                  onChanged: (value) {
                    days = int.tryParse(value) ?? 0;
                  },
                ),
              ),
            ],
          ),
          Row(
            children: [
              const SizedBox(width: 50, child: Text("Hours:")),
              const SizedBox(width: 10),
              SizedBox(
                width: 100,
                child: TextFormField(
                  initialValue: hours.toString(),
                  keyboardType: TextInputType.number,
                  decoration: const InputDecoration(
                    border: InputBorder.none,
                    enabledBorder: InputBorder.none,
                    hintText: "Hours",
                  ),
                  onChanged: (value) {
                    hours = int.tryParse(value) ?? 0;
                  },
                ),
              ),
            ],
          ),
        ],
      ),
      actions: [
        TextButton(
            onPressed: () {
              Navigator.of(context).pop();
            },
            child: const Text("Cancel")),
        ElevatedButton(
            // style: ElevatedButton.styleFrom(backgroundColor: Styles.dangerBgColor, foregroundColor: Styles.dangerTextColor),
            onPressed: () {
              final newDuration = Duration(days: days, hours: hours);
              if (newDuration.compareTo(duration) != 0) {
                trackerRow = trackerRow.copyWith(notificationInterval: pb.Duration(seconds: Int64(newDuration.inSeconds)));
                ref.read(trackerNotifierProvider.notifier).updateTracker(trackerRow);
              }
              Navigator.of(context).pop();
            },
            child: const Text("Save")),
      ],
    );
  }

  Widget table(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final trackerFuture = ref.watch(trackerFutureProvider);
      final trackerRows = ref.watch(trackerNotifierProvider);
      final showAddTrackerButton = ref.watch(showAddTrackerButtonProvider);
      final trackerChatRooms = ref.watch(trackerChatRoomsProvider);
      return Builder(
        builder: (BuildContext context) {
          if (trackerFuture.isLoading) {
            return const CircularProgressIndicator();
          } else {
            return DataTable(
                columns: const [
                  DataColumn(label: Text("Track Address")),
                  DataColumn(label: Text("Notification")),
                  DataColumn(label: Text("Chat")),
                  DataColumn(label: Text("Action")),
                ],
                rows: trackerRows.map((trackerRow) {
                  return DataRow(cells: [
                    DataCell(trackerRow.updatedAt == null ? AddressInputWidget(ref, trackerRow) : Text(trackerRow.shortenedAddress)),
                    DataCell(Row(
                      children: [
                        Text(trackerRow.notificationIntervalPrettyString, textAlign: TextAlign.center),
                        IconButton(
                          padding: const EdgeInsets.all(0),
                          onPressed: () async =>
                              {showDialog(context: context, builder: (context) => notificationIntervalDialog(context, trackerRow, ref))},
                          icon: const Icon(Icons.edit, size: iconSizeSmall),
                        ),
                      ],
                    )),
                    DataCell(
                      DropdownButton<TrackerChatRoom>(
                        value: trackerRow.chatRoom,
                        icon: const Icon(Icons.arrow_downward),
                        iconSize: iconSizeSmall,
                        elevation: 16,
                        style: const TextStyle(color: Colors.deepPurple),
                        underline: Container(
                          height: 2,
                          color: Colors.deepPurpleAccent,
                        ),
                        onChanged: (TrackerChatRoom? newValue) async {
                          if (newValue == null) {
                            return;
                          }
                          trackerRow = trackerRow.copyWith(chatRoom: newValue);
                          await ref.read(trackerNotifierProvider.notifier).updateTracker(trackerRow);
                        },
                        items: trackerChatRooms.map<DropdownMenuItem<TrackerChatRoom>>((trackerChatRoom) {
                          return DropdownMenuItem<TrackerChatRoom>(
                            value: trackerChatRoom,
                            child: Text(trackerChatRoom.name),
                          );
                        }).toList(),
                      ),
                    ),
                    DataCell(IconButton(
                      padding: const EdgeInsets.all(0),
                      onPressed: () async => {await ref.read(trackerNotifierProvider.notifier).deleteTracker(trackerRow)},
                      icon: const Icon(Icons.delete, size: iconSize),
                    )),
                  ]);
                }).toList()
                  ..addAll([
                    if (showAddTrackerButton)
                      DataRow(cells: [
                        const DataCell(Text("")),
                        const DataCell(Text("")),
                        const DataCell(Text("")),
                        DataCell(IconButton(
                          padding: const EdgeInsets.all(0),
                          onPressed: () async => {ref.read(trackerNotifierProvider.notifier).addTracker()},
                          icon: const Icon(Icons.add, size: iconSize),
                        ))
                      ])
                  ]));
          }
        },
      );
    });
  }

  Widget validationError(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final trackerRows = ref.watch(trackerNotifierProvider);
      return Builder(
        builder: (BuildContext context) {
          final trackerRowWithValErr =
              trackerRows.firstWhereOrNull((trackerRow) => trackerRow.updatedAt == null && !trackerRow.isAddressValid);
          if (trackerRowWithValErr != null) {
            return const Padding(
              padding: EdgeInsets.only(left: 20),
              child: Text("Invalid address", style: TextStyle(color: Colors.red)),
            );
          } else {
            return const SizedBox.shrink();
          }
        },
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          MessageOverlayListener(
            provider: messageProvider,
            child: Container(
              width: MediaQuery.of(context).size.width,
              padding: const EdgeInsets.only(top: Styles.topPadding, left: Styles.sidePadding, right: Styles.sidePadding),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const HeaderWidget(),
                  const SizedBox(height: 10),
                  table(context),
                  validationError(context),
                  const Spacer(flex: 1),
                  const Flexible(flex: 0, child: FooterWidget()),
                ],
              ),
            ),
          ),
        ],
      ),
      bottomNavigationBar: BottomNavigationBarWidget(jwtManager: jwtManager),
    );
  }
}

class AddressInputWidget extends HookWidget {
  final WidgetRef ref;
  final TrackerRow trackerRow;

  const AddressInputWidget(this.ref, this.trackerRow, {Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final debouncer = Debouncer<String>(const Duration(milliseconds: 200), initialValue: "");
    debouncer.values.listen((value) {
      ref.read(trackerNotifierProvider.notifier).updateTracker(trackerRow.copyWith(address: value));
    });
    final controller = useTextEditingController(text: trackerRow.shortenedAddress);
    return TextField(
      controller: controller,
      decoration: const InputDecoration(
        hintText: 'Enter wallet address',
        border: InputBorder.none,
        enabledBorder: InputBorder.none,
        hintStyle: TextStyle(color: Colors.grey, fontSize: 12),
      ),
      onChanged: (value) => {debouncer.value = value},
    );
  }
}
