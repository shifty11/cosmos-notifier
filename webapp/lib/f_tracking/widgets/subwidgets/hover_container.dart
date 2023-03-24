import 'package:flutter/material.dart';

class HoverContainer extends StatefulWidget {
  final Widget child;

  HoverContainer({required this.child});

  @override
  _HoverContainerState createState() => _HoverContainerState();
}

class _HoverContainerState extends State<HoverContainer> {
  bool _isHovered = false;

  @override
  Widget build(BuildContext context) {
    return MouseRegion(
      onEnter: (_) => setState(() => _isHovered = true),
      onExit: (_) => setState(() => _isHovered = false),
      child: Container(
        decoration: BoxDecoration(
          // get hover color from theme
          color: _isHovered ? Theme.of(context).hoverColor : Colors.transparent,
          // border: Border.all(color: Colors.black),
        ),
        child: widget.child,
      ),
    );
  }
}
