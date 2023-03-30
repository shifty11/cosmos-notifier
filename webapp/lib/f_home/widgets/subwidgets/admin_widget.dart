import 'package:cosmos_notifier/api/protobuf/dart/dev_service.pbenum.dart';
import 'package:cosmos_notifier/f_home/services/auth_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class AdminWidget extends StatelessWidget {
  const AdminWidget({Key? key}) : super(key: key);
  static const iconSize = 30.0;
  static const spaceBetween = 15.0;

  showAdminPopup(BuildContext context) {
    showDialog(
      context: context,
      builder: (_) {
        return AlertDialog(
          title: const Text('Admin Actions'),
          content: const LoginForm(),
          actions: [
            TextButton(
              onPressed: () => Navigator.pop(context),
              child: const Text('Cancel'),
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return OutlinedButton(onPressed: () => showAdminPopup(context), child: const Text("Admin"));
  }
}

class FormData {
  int userId;
  DevLoginRequest_UserType userType;
  DevLoginRequest_Role roleType;

  FormData({required this.userId, required this.userType, required this.roleType});
}


class LoginForm extends StatefulWidget {
  const LoginForm({Key? key}) : super(key: key);

  @override
  _LoginFormState createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  final _formKey = GlobalKey<FormState>();
  final _formData = FormData(userId: 0, userType: DevLoginRequest_UserType.DISCORD, roleType: DevLoginRequest_Role.ADMIN);

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          TextFormField(
            decoration: const InputDecoration(
              labelText: 'User ID',
            ),
            keyboardType: TextInputType.number,
            validator: (value) {
              if (value == null || value.isEmpty) {
                return null;
              }
              if (int.tryParse(value) == null) {
                return 'Please enter a valid number.';
              }
              return null;
            },
            onSaved: (value) {
              _formData.userId = int.tryParse(value ?? '0') ?? 0;
            },
          ),
          const SizedBox(height: 16),
          DropdownButtonFormField<DevLoginRequest_UserType>(
            value: _formData.userType,
            onChanged: (value) {
              setState(() {
                _formData.userType = value!;
              });
            },
            items: DevLoginRequest_UserType.values
                .map((userType) => DropdownMenuItem(
              value: userType,
              child: Text(userType.toString().split('.').last.toLowerCase()),
            ))
                .toList(),
            decoration: const InputDecoration(
              labelText: 'User Type',
            ),
          ),
          const SizedBox(height: 16),
          DropdownButtonFormField<DevLoginRequest_Role>(
            value: _formData.roleType,
            onChanged: (value) {
              setState(() {
                _formData.roleType = value!;
              });
            },
            items: DevLoginRequest_Role.values
                .map((roleType) => DropdownMenuItem(
              value: roleType,
              child: Text(roleType.toString().split('.').last.toLowerCase()),
            ))
                .toList(),
            decoration: const InputDecoration(
              labelText: 'Role Type',
            ),
          ),
          const SizedBox(height: 16),
          Consumer(
            builder: (BuildContext context, WidgetRef ref, Widget? child) {
              return ElevatedButton(
                onPressed: () {
                  if (_formKey.currentState!.validate()) {
                    _formKey.currentState!.save();
                    ref.read(authStateProvider.notifier).devLogin(
                      userId: _formData.userId,
                      userType: _formData.userType,
                      roleType: _formData.roleType,
                    );
                  }
                },
                child: const Text('Login'),
              );
            }
          ),
        ],
      ),
    );
  }
}
