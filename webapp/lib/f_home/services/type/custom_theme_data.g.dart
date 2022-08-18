// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'custom_theme_data.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_CustomThemeData _$$_CustomThemeDataFromJson(Map<String, dynamic> json) =>
    _$_CustomThemeData(
      bgColor: colorFromJson(json['bg_color'] as String),
      textColor: colorFromJson(json['text_color'] as String),
      hintColor: colorFromJson(json['hint_color'] as String),
      linkColor: colorFromJson(json['link_color'] as String),
      buttonColor: colorFromJson(json['button_color'] as String),
      buttonTextColor: colorFromJson(json['button_text_color'] as String),
    );

Map<String, dynamic> _$$_CustomThemeDataToJson(_$_CustomThemeData instance) =>
    <String, dynamic>{
      'bg_color': colorToJson(instance.bgColor),
      'text_color': colorToJson(instance.textColor),
      'hint_color': colorToJson(instance.hintColor),
      'link_color': colorToJson(instance.linkColor),
      'button_color': colorToJson(instance.buttonColor),
      'button_text_color': colorToJson(instance.buttonTextColor),
    };
