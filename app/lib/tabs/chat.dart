import 'package:flutter/material.dart';
import 'package:app/card/card.dart';

class ChatTab extends StatefulWidget {
  @override
  ChatState createState() => ChatState();
}

class ChatState extends State<ChatTab> {
  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisSize: MainAxisSize.min,
      children: <Widget>[
        CardFulDefault(),
        CardFulDefault(),
        CardFulDefault(),
      ],
    );
  }
}
