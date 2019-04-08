import 'package:flutter/material.dart';
import 'package:app/card/card.dart';

class ChatTab extends StatefulWidget {
  @override
  ChatState createState() => ChatState();
}

class ChatState extends State<ChatTab> {
  @override
  Widget build(BuildContext context) {
    var card = CardWidget(
      cardChild: Column(
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          Container(
            child: TextField(
              decoration: InputDecoration(
                prefixIcon: Icon(Icons.account_circle, color: Colors.white),
                labelText: "用户名",
                labelStyle: TextStyle(color: Colors.white),
              ),
              style: TextStyle(color: Colors.white),
              keyboardType: TextInputType.emailAddress,
            ),
            padding: EdgeInsets.only(top: 20.0, bottom: 10.0, left: 20.0, right: 20.0),
          ),
          Container(
            child: TextField(
              decoration: InputDecoration(
                prefixIcon: Icon(Icons.vpn_key, color: Colors.white,),
                labelText: "密码",
                labelStyle: TextStyle(color: Colors.white),
              ),
              style: TextStyle(color: Colors.white),
              keyboardType: TextInputType.emailAddress,
            ),
            padding: EdgeInsets.all(20.0),
          ),
        ],
      ),
      lButtonText: "注册",
      rButtonText: "登陆",
      elevation: 10.0,
      margin: EdgeInsets.only(left: 20.0, right: 20.0, bottom: 20.0, top: 60.0),
    );

    return Column(
      children: <Widget>[
        card,
      ],
    );
  }
}
