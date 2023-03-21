import 'package:cosmos_notifier/common/header_widget.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/bottom_navigation_bar_widget.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/footer_widget.dart';
import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_provider.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_messages/riverpod_messages.dart';

class TrackingPage extends StatelessWidget {
  static const iconSize = 24.0;

  const TrackingPage({Key? key, errorCode}) : super(key: key);

  Widget table(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      return FutureBuilder(
        future: ref.watch(trackerFutureProvider.future),
        builder: (BuildContext context, AsyncSnapshot<List<TrackerRow>> snapshot) {
          if (snapshot.hasData) {
            return DataTable(
                columns: const [
                  DataColumn(label: Text("Track Address")),
                  DataColumn(label: Text("Notification")),
                  DataColumn(label: Text("Chat")),
                  DataColumn(label: Text("Action")),
                ],
                rows: snapshot.data!.map((trackerRow) {
                  return DataRow(cells: [
                    DataCell(Text(trackerRow.address)),
                    DataCell(Text(trackerRow.notificationInterval)),
                    DataCell(Text(trackerRow.chatId.toString())),
                    trackerRow.isAddRow
                        ? DataCell(IconButton(
                            padding: const EdgeInsets.all(0),
                            onPressed: () async => {
                              ref.read(trackerNotifierProvider.notifier).addTracker()
                            },
                            icon: const Icon(Icons.add, size: iconSize),
                          ))
                        : DataCell(IconButton(
                            padding: const EdgeInsets.all(0),
                            onPressed: () async => {
                              await ref.read(trackerNotifierProvider.notifier).deleteTracker(trackerRow)
                            },
                            icon: const Icon(Icons.delete, size: iconSize),
                          )),
                  ]);
                }).toList());
            // rows: snapshot.data!.trackers.map((tracker) {
            //   return DataRow(cells: [
            //     DataCell(Text(shortenedBech32Address(tracker.address))),
            //     DataCell(Text(notificationIntervalAsString(tracker.notificationInterval.seconds))),
            //     DataCell(Text(tracker.discordChannelId.toString())),
            //     DataCell(IconButton(
            //       padding: const EdgeInsets.all(0),
            //       onPressed: () async => {},
            //       icon: const Icon(Icons.delete, size: iconSize),
            //     )),
            //   ]);
            // }).toList()
            //   ..add(
            //     DataRow(cells: [
            //       const DataCell(Text("")),
            //       const DataCell(Text("")),
            //       const DataCell(Text("")),
            //       DataCell(IconButton(
            //         padding: const EdgeInsets.all(0),
            //         onPressed: () async => {},
            //         icon: const Icon(Icons.add, size: iconSize),
            //       )),
            //     ]),
            //   ));
          } else {
            return const CircularProgressIndicator();
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
