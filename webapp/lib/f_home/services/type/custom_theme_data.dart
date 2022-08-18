import 'package:flutter/material.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:string_to_hex/string_to_hex.dart';

part 'custom_theme_data.freezed.dart';
part 'custom_theme_data.g.dart';

@freezed
@immutable
class CustomThemeData with _$CustomThemeData {
  factory CustomThemeData({
    @JsonKey(name: 'bg_color', fromJson: colorFromJson, toJson: colorToJson) required Color bgColor,
    @JsonKey(name: 'text_color', fromJson: colorFromJson, toJson: colorToJson) required Color textColor,
    @JsonKey(name: 'hint_color', fromJson: colorFromJson, toJson: colorToJson) required Color hintColor,
    @JsonKey(name: 'link_color', fromJson: colorFromJson, toJson: colorToJson) required Color linkColor,
    @JsonKey(name: 'button_color', fromJson: colorFromJson, toJson: colorToJson) required Color buttonColor,
    @JsonKey(name: 'button_text_color', fromJson: colorFromJson, toJson: colorToJson) required Color buttonTextColor,
  }) = _CustomThemeData;

  factory CustomThemeData.fromJson(Map<String, dynamic> json) => _$CustomThemeDataFromJson(json);
}

Color colorFromJson(String color) {
  return Color(int.parse('FF${color.replaceAll("#", "")}', radix: 16));
}

String colorToJson(Color color) => StringToHex.toHexString(color.value);
