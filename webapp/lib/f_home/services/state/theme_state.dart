import 'package:flutter/material.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';

part 'theme_state.freezed.dart';

@freezed
class ThemeState with _$ThemeState {
  const ThemeState._();

  const factory ThemeState.initial({required ThemeData darkStyle, required ThemeData lightStyle}) = Initial;
  const factory ThemeState.custom({required ThemeData style}) = Custom;
}