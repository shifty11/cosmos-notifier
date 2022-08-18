import 'dart:convert';

import 'package:fixnum/fixnum.dart' as fixnum;
import 'package:freezed_annotation/freezed_annotation.dart';

part 'login_data.freezed.dart';

@freezed
@immutable
class LoginData with _$LoginData {
  const LoginData._();

  const factory LoginData(String dataStr) = _LoginData;

  bool get isValid {
    return data.isNotEmpty && id != fixnum.Int64() && authDate != fixnum.Int64() && hash.isNotEmpty;
  }

  Iterable<List<String>> get _fields {
    return dataStr.replaceAll("&", "\n").split("\n").map((e) => e.split("=")).where((e) => e.length == 2 && e[0] != "chat_id");
  }

  String get _user {
    final user = _fields.where((e) => e[0] == 'user').map((e) => e[1]);
    return user.isNotEmpty ? user.first : "";
  }

  String get data {
    final list = _fields.where((e) => e[0] != "hash").map((e) => "${e[0]}=${e[1]}").toList();
    list.sort(((a, b) => a.compareTo(b)));
    return list.join("\n");
  }

  fixnum.Int64 get id {
    int id = 0;
    if (_user.isNotEmpty) {
      id = jsonDecode(_user)['id'] ?? 0;
    }
    final idField = _fields.where((e) => e[0] == "id");
    if (idField.isNotEmpty) {
      id = int.tryParse(idField.first[1]) ?? 0;
    }
    return fixnum.Int64(id);
  }

  String get username {
    if (_user.isNotEmpty) {
      return jsonDecode(_user)['username'] ?? "";
    }
    final username = _fields.where((e) => e[0] == "username");
    if (username.isNotEmpty) {
      return username.first[1];
    }
    return "";
  }

  fixnum.Int64 get authDate {
    int authDate = 0;
    final authDateField = _fields.where((e) => e[0] == 'auth_date');
    if (authDateField.isNotEmpty) {
      authDate = int.tryParse(authDateField.first[1]) ?? 0;
    }
    return fixnum.Int64(authDate);
  }

  String get hash {
    final hash = _fields.where((e) => e[0] == 'hash');
    return hash.isNotEmpty ? hash.first[1] : "";
  }
}
