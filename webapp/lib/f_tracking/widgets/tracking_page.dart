import 'package:cosmos_notifier/api/protobuf/dart/google/protobuf/duration.pb.dart' as pb;
import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pbgrpc.dart';
import 'package:cosmos_notifier/common/header_widget.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/bottom_navigation_bar_widget.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/footer_widget.dart';
import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/state/validator_bundle.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_provider.dart';
import 'package:cosmos_notifier/f_tracking/widgets/subwidgets/address_input_widget.dart';
import 'package:cosmos_notifier/f_tracking/widgets/subwidgets/hover_container.dart';
import 'package:cosmos_notifier/f_tracking/widgets/subwidgets/validation_form_widget.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multi_select_flutter/dialog/multi_select_dialog_field.dart';
import 'package:multi_select_flutter/util/multi_select_item.dart';
import 'package:responsive_framework/responsive_framework.dart';
import 'package:riverpod_messages/riverpod_messages.dart';

// ignore: must_be_immutable
class ValidatorMultiSelectDialogField<ValidatorBundle> extends MultiSelectDialogField<FreezedValidatorBundle> {
  ValidatorMultiSelectDialogField({
    required super.items,
    required super.onConfirm,
    required super.initialValue,
    super.searchable = true,
    super.key,
    super.buttonText = const Text("Select validators"),
  });
}

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
      return Builder(
        builder: (BuildContext context) {
          if (ref.watch(trackerFutureProvider).isLoading) {
            return const Center(child: CircularProgressIndicator());
          } else {
            final trackerRows = ref.watch(trackerNotifierProvider);
            final showChatRoomColumn = ref.watch(showChatRoomColumnProvider);
            return DataTable(
                columnSpacing: ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? 10 : null,
                sortAscending: ref.watch(trackerSortProvider).isAscending,
                sortColumnIndex: ref.watch(trackerSortProvider).sortType.index,
                columns: [
                  DataColumn(
                      label: Text("Address (${trackerRows.length})"),
                      tooltip: "Wallet address that is being tracked",
                      onSort: (columnIndex, ascending) {
                        ref.read(trackerSortProvider.notifier).state =
                            TrackerSortState(isAscending: ascending, sortType: TrackerSortType.address);
                        ref.read(trackerNotifierProvider.notifier).sort();
                      }),
                  DataColumn(
                      label: const Padding(
                        padding: EdgeInsets.only(left: 10),
                        child: Text("Notification"),
                      ),
                      tooltip: "Time when the notification will be sent before the proposal ends",
                      onSort: (columnIndex, ascending) {
                        ref.read(trackerSortProvider.notifier).state =
                            TrackerSortState(isAscending: ascending, sortType: TrackerSortType.notificationInterval);
                        ref.read(trackerNotifierProvider.notifier).sort();
                      }),
                  if (showChatRoomColumn)
                    DataColumn(
                        label: const Text("Chat"),
                        tooltip: "Reminder will be sent to this chat",
                        onSort: (columnIndex, ascending) {
                          ref.read(trackerSortProvider.notifier).state =
                              TrackerSortState(isAscending: ascending, sortType: TrackerSortType.chatRoom);
                          ref.read(trackerNotifierProvider.notifier).sort();
                        }),
                  const DataColumn(label: Text("Action")),
                ],
                rows: trackerRows.map((trackerRow) {
                  var addressSize = ResponsiveWrapper.of(context).isSmallerThan(MOBILE) ? AddressSize.veryShort : AddressSize.short;
                  if (ResponsiveWrapper.of(context).isLargerThan(TABLET)) {
                    addressSize = AddressSize.long;
                  }
                  return DataRow(cells: [
                    DataCell(trackerRow.isSaved ? Text(trackerRow.shortenedAddress(addressSize)) : AddressInputWidget(ref, trackerRow)),
                    DataCell(
                      GestureDetector(
                        onTap: () async =>
                            showDialog(context: context, builder: (context) => notificationIntervalDialog(context, trackerRow, ref)),
                        child: HoverContainer(
                          child: Container(
                            padding: const EdgeInsets.symmetric(horizontal: 10),
                            constraints: BoxConstraints(
                              maxWidth: ResponsiveWrapper.of(context).isSmallerThan(TABLET) ? MediaQuery.of(context).size.width / 5 : 200,
                              minHeight: 56,
                            ),
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
                    if (showChatRoomColumn)
                      DataCell(
                        DropdownButton<TrackerChatRoom>(
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
                          items: ref.watch(trackerChatRoomsProvider).map<DropdownMenuItem<TrackerChatRoom>>((trackerChatRoom) {
                            return DropdownMenuItem<TrackerChatRoom>(
                              value: trackerChatRoom,
                              child: LimitedBox(
                                maxWidth: ResponsiveWrapper.of(context).isSmallerThan(MOBILE) ? 80 : 200,
                                child: Text(
                                  trackerChatRoom.name,
                                  overflow: TextOverflow.ellipsis,
                                ),
                              ),
                            );
                          }).toList(),
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
                    if (ref.watch(showAddTrackerButtonProvider))
                      DataRow(cells: [
                        const DataCell(Text("")),
                        const DataCell(Text("")),
                        if (showChatRoomColumn) const DataCell(Text("")),
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
      return Builder(
        builder: (BuildContext context) {
          if (ref.watch(hasValidationErrorProvider)) {
            return const Text("Invalid address", style: TextStyle(color: Colors.red));
          } else {
            return const SizedBox.shrink();
          }
        },
      );
    });
  }

  showValidationChangePopup(BuildContext context, List<FreezedValidatorBundle> toBeTracked, List<FreezedValidatorBundle> toBeAdded,
      List<FreezedValidatorBundle> toBeDeleted) {
    showDialog(
      context: context,
      builder: (_) {
        return ValidatorForm(toBeTracked, toBeAdded, toBeDeleted);
      },
    );
  }

  Widget validatorSelection() {
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      final validators = ref.watch(validatorBundleProvider);
      final items = validators
          .map((bundle) => MultiSelectItem(bundle, "${bundle.moniker} (${bundle.validators.length})")..selected = bundle.isTracked)
          .toList();
      return ValidatorMultiSelectDialogField(
        items: items,
        initialValue: List<FreezedValidatorBundle>.from(items.where((item) => item.selected).map((item) => item.value)),
        onConfirm: (List<FreezedValidatorBundle> result) {
          final toBeAdded = result.where((val) => !val.isTracked).toList();
          final toBeTracked = items.where((item) => item.selected).map((item) => item.value).toList();
          final toBeDeleted = items.where((item) => !item.selected).map((item) => item.value).where((val) => val.isTracked).toList();
          if (toBeAdded.isNotEmpty) {
            showValidationChangePopup(context, toBeTracked, toBeAdded, toBeDeleted);
          } else if (toBeDeleted.isNotEmpty) {
            ref
                .read(trackerNotifierProvider.notifier)
                .trackValidators(toBeTracked, toBeAdded, toBeDeleted, TrackerChatRoom(), Duration.zero);
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
                  const SizedBox(height: 20),
                  Text("Reminders", style: Theme.of(context).textTheme.headlineMedium),
                  const SizedBox(height: 5),
                  const Text("Track your validator and get notified when it's time to vote.\nYou can also add custom addresses.",
                      maxLines: 3),
                  const SizedBox(height: 10),
                  Expanded(
                    child: SingleChildScrollView(
                      child: Column(
                        children: [
                          validatorSelection(),
                          const SizedBox(height: 10),
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
