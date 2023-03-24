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
import 'package:cosmos_notifier/f_tracking/widgets/subwidgets/address_input_widget.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:responsive_framework/responsive_framework.dart';
import 'package:riverpod_messages/riverpod_messages.dart';

class TrackingPage extends StatelessWidget {
  static const iconSize = 24.0;
  static const iconSizeSmall = 12.0;

  const TrackingPage({Key? key, errorCode}) : super(key: key);

  Widget notificationIntervalDialog(BuildContext context, TrackerRow trackerRow, WidgetRef ref) {
    final duration = Duration(seconds: trackerRow.notificationInterval.seconds.toInt());
    var days = duration.inDays;
    var hours = duration.inHours - days * 24;
    return Shortcuts(
      shortcuts: {
        LogicalKeySet(LogicalKeyboardKey.escape): const DismissIntent(),
      },
      child: AlertDialog(
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
              onPressed: () {
                final newDuration = Duration(days: days, hours: hours);
                if (newDuration.compareTo(duration) != 0) {
                  final pbDuration = pb.Duration(seconds: Int64(newDuration.inSeconds));
                  if (pbDuration.seconds != trackerRow.notificationInterval.seconds) {
                    trackerRow = trackerRow.copyWith(notificationInterval: pbDuration);
                    ref.read(trackerNotifierProvider.notifier).updateTracker(trackerRow);
                  }
                }
                Navigator.of(context).pop();
              },
              child: const Text("Save")),
        ],
      ),
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
            return const Center(child: CircularProgressIndicator());
          } else {
            return DataTable(
                columnSpacing: ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? 10 : null,
                columns: const [
                  DataColumn(label: Text("Address")),
                  DataColumn(
                      label: Padding(
                    padding: EdgeInsets.only(left: 10),
                    child: Text("Notification"),
                  )),
                  DataColumn(label: Text("Chat")),
                  DataColumn(label: Text("Action")),
                ],
                rows: trackerRows.map((trackerRow) {
                  return DataRow(cells: [
                    DataCell(trackerRow.isSaved
                        ? Text(trackerRow.shortenedAddress(ResponsiveWrapper.of(context).isSmallerThan(TABLET)))
                        : AddressInputWidget(ref, trackerRow)),
                    DataCell(
                      ElevatedButton(
                        onPressed: () {  },
                        style: ElevatedButton.styleFrom(
                          backgroundColor: Colors.transparent,
                          shadowColor: Colors.transparent,
                          elevation: 0,
                          maximumSize: Size.fromWidth(
                            ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? MediaQuery.of(context).size.width / 5 : 200,
                          ),
                          padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 22),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(0),
                          ),
                        ),
                        child: GestureDetector(
                          onTap: () async =>
                              showDialog(context: context, builder: (context) => notificationIntervalDialog(context, trackerRow, ref)),
                          child: Container(
                            constraints: BoxConstraints(maxWidth: ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? MediaQuery.of(context).size.width / 5 : 200),
                            child: Row(
                              mainAxisAlignment: MainAxisAlignment.start,
                              children: [
                                Flexible(
                                  child: Text(
                                    trackerRow.notificationIntervalPrettyString,
                                    overflow: TextOverflow.ellipsis,
                                  ),
                                ),
                                const SizedBox(width: 5),
                                const Icon(
                                  Icons.edit,
                                  size: iconSizeSmall,
                                ),
                              ],
                            ),
                          ),
                        ),
                      ),
                    ),
                    DataCell(
                      LimitedBox(
                        maxWidth: 200,
                        child: DropdownButton<TrackerChatRoom>(
                          focusColor: Colors.transparent,
                          value: trackerRow.chatRoom,
                          icon: const Icon(Icons.arrow_downward),
                          iconSize: iconSizeSmall,
                          style: Theme.of(context).textTheme.bodyMedium,
                          onChanged: (TrackerChatRoom? newValue) async {
                            if (newValue == null || newValue == trackerRow.chatRoom) {
                              return;
                            }
                            trackerRow = trackerRow.copyWith(chatRoom: newValue);
                            await ref.read(trackerNotifierProvider.notifier).updateTracker(trackerRow);
                          },
                          items: trackerChatRooms.map<DropdownMenuItem<TrackerChatRoom>>((trackerChatRoom) {
                            return DropdownMenuItem<TrackerChatRoom>(
                              value: trackerChatRoom,
                              child: Text(
                                trackerChatRoom.name,
                                overflow: TextOverflow.ellipsis,
                              ),
                            );
                          }).toList(),
                        ),
                      ),
                    ),
                    DataCell(
                      IconButton(
                        padding: const EdgeInsets.all(0),
                        onPressed: () async => {await ref.read(trackerNotifierProvider.notifier).deleteTracker(trackerRow)},
                        icon: const Icon(Icons.delete, size: iconSize),
                      ),
                    ),
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
              padding: EdgeInsets.only(left: 24),
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
          MessageSnackbarListener(
            provider: messageProvider,
            child: Container(
              width: MediaQuery.of(context).size.width,
              padding: const EdgeInsets.only(top: Styles.topPadding, left: Styles.sidePadding, right: Styles.sidePadding),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const HeaderWidget(),
                  const SizedBox(height: 10),
                  Text("Reminders", style: Theme.of(context).textTheme.headlineMedium),
                  const SizedBox(height: 10),
                  const Text("Set up your reminders. You will get a notification if you did not vote on a proposal and it is about to end.", maxLines: 3),
                  const SizedBox(height: 10),
                  Expanded(
                    child: SingleChildScrollView(
                      child: Column(
                        children: [
                          table(context),
                          validationError(context),
                        ],
                      ),
                    ),
                  ),
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
