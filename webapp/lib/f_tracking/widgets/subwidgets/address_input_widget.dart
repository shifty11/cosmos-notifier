import 'package:cosmos_notifier/f_tracking/services/state/tracker_row.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_provider.dart';
import 'package:debounce_throttle/debounce_throttle.dart';
import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

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
      style: Theme.of(context).textTheme.bodyMedium,
      decoration: InputDecoration(
        hintText: 'Enter wallet address',
        border: InputBorder.none,
        enabledBorder: InputBorder.none,
        hintStyle: TextStyle(color: Colors.grey, fontSize: Theme.of(context).textTheme.bodyMedium!.fontSize),
      ),
      onChanged: (value) => {debouncer.value = value},
    );
  }
}
