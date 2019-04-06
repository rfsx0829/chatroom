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
        CardWidget(
          cardChild: ListTile(
            title: Text("Hello World", style: TextStyle(color: Colors.white, fontSize: 40.0),),
            contentPadding: EdgeInsets.all(20.0),
          ),
          lButtonText: "酷哦",
          lButtonFunc: () {
            print("Hello");
          },
        ),
      ],
    );
  }
}
