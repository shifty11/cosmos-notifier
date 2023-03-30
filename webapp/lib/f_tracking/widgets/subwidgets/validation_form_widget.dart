import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart';
import 'package:cosmos_notifier/f_tracking/services/state/validator_bundle.dart';
import 'package:cosmos_notifier/f_tracking/services/tracker_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class ValidatorForm extends ConsumerStatefulWidget {
  final List<FreezedValidatorBundle> toBeTracked;
  final List<FreezedValidatorBundle> toBeAdded;
  final List<FreezedValidatorBundle> toBeDeleted;

  const ValidatorForm(this.toBeTracked, this.toBeAdded, this.toBeDeleted, {Key? key}) : super(key: key);

  @override
  _ValidatorFormState createState() => _ValidatorFormState();
}

class _ValidatorFormState extends ConsumerState<ValidatorForm> {
  final _formKey = GlobalKey<FormState>();

  TrackerChatRoom? _selectedChat;
  int _days = 0;
  int _hours = 0;

  @override
  void initState() {
    super.initState();
    final duration = Duration(seconds: ref.read(trackerNotifierProvider.notifier).getDefaultNotificationInterval().seconds.toInt());
    _days = duration.inDays;
    _hours = duration.inHours - _days * 24;
    _selectedChat = ref.read(trackerNotifierProvider.notifier).getDefaultChatRoom();
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Track validators'),
      content: SizedBox(
        height: 200,
        child: Form(
          key: _formKey,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const SizedBox(height: 16),
              DropdownButtonFormField<TrackerChatRoom>(
                value: _selectedChat,
                onChanged: (value) {
                  setState(() {
                    _selectedChat = value;
                  });
                },
                items: ref
                    .read(trackerChatRoomsProvider)
                    .map((chat) => DropdownMenuItem(
                          value: chat,
                          child: Text(chat.name),
                        ))
                    .toList(),
                decoration: const InputDecoration(
                  labelText: 'Chat',
                ),
                validator: (value) {
                  if (value == null) {
                    return 'Please select a chat.';
                  }
                  return null;
                },
              ),
              const SizedBox(height: 16),
              Row(
                children: [
                  Expanded(
                    child: TextFormField(
                      initialValue: _days.toString(),
                      decoration: const InputDecoration(
                        labelText: 'Days',
                      ),
                      keyboardType: TextInputType.number,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter a number of days.';
                        }
                        if (int.tryParse(value) == null) {
                          return 'Please enter a valid number.';
                        }
                        return null;
                      },
                      onSaved: (value) {
                        _days = int.parse(value!);
                      },
                    ),
                  ),
                  const SizedBox(width: 16),
                  Expanded(
                    child: TextFormField(
                      initialValue: _hours.toString(),
                      decoration: const InputDecoration(
                        labelText: 'Hours',
                      ),
                      keyboardType: TextInputType.number,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter a number of hours.';
                        }
                        if (int.tryParse(value) == null) {
                          return 'Please enter a valid number.';
                        }
                        return null;
                      },
                      onSaved: (value) {
                        _hours = int.parse(value!);
                      },
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.pop(context),
          child: const Text('Cancel'),
        ),
        ElevatedButton(
            onPressed: () {
              if (_formKey.currentState!.validate()) {
                _formKey.currentState!.save();
                final newDuration = Duration(days: _days, hours: _hours);
                ref
                    .read(trackerNotifierProvider.notifier)
                    .trackValidators(widget.toBeTracked, widget.toBeAdded, widget.toBeDeleted, _selectedChat!, newDuration);
                Navigator.pop(context);
              }
            },
            child: const Text('Save'))
      ],
    );
  }
}
