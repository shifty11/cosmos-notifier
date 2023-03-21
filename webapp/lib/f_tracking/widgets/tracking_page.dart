import 'package:cosmos_notifier/common/header_widget.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/bottom_navigation_bar_widget.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/footer_widget.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_provider.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:debounce_throttle/debounce_throttle.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_messages/riverpod_messages.dart';

class TrackingPage extends StatelessWidget {
  static const iconSize = 24.0;

  const TrackingPage({Key? key, errorCode}) : super(key: key);

  Widget table(BuildContext context) {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final trackerFuture = ref.watch(trackerFutureProvider);
      final trackerRows = ref.watch(trackerNotifierProvider);
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
                  final debouncer = Debouncer<String>(const Duration(milliseconds: 200), initialValue: "");
                  debouncer.values.listen((value) {
                    ref.read(trackerNotifierProvider.notifier).updateTracker(trackerRow.copyWith(address: value));
                  });
                  return DataRow(cells: [
                    DataCell(TextField(
                      controller: TextEditingController(text: trackerRow.shortenedBech32Address),
                      onChanged: (value) => {debouncer.value = value},
                    )),
                    DataCell(Text(trackerRow.notificationIntervalPrettyString)),
                    DataCell(Text(trackerRow.chatId.toString())),
                    trackerRow.isAddRow
                        ? DataCell(IconButton(
                            padding: const EdgeInsets.all(0),
                            onPressed: () async => {ref.read(trackerNotifierProvider.notifier).addTracker()},
                            icon: const Icon(Icons.add, size: iconSize),
                          ))
                        : DataCell(IconButton(
                            padding: const EdgeInsets.all(0),
                            onPressed: () async => {await ref.read(trackerNotifierProvider.notifier).deleteTracker(trackerRow)},
                            icon: const Icon(Icons.delete, size: iconSize),
                          )),
                  ]);
                }).toList());
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
