import 'package:flutter/material.dart';
import 'src/widgets/chat_home.dart';

void main() {
  runApp(new ChatApp());
}

class ChatApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      theme: ThemeData(primarySwatch: Colors.teal),
      home: ChatHome(),
    );
  }
}
