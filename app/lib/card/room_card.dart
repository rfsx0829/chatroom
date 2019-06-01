import 'package:flutter/material.dart';

import 'package:app/card/card.dart';

typedef void Func(String argv);

class RoomCard extends StatefulWidget {
  final Func callback;

  RoomCard(this.callback);

  @override
  RoomCardState createState() => RoomCardState(callback);
}

class RoomCardState extends State<RoomCard> {
  String pass;
  final Func callback;

  RoomCardState(this.callback);

  @override
  Widget build(BuildContext context) {
    var card = CardWidget(
      Column(
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          Container(
            child: TextField(
              decoration: InputDecoration(
                prefixIcon: Icon(Icons.vpn_key, color: Colors.white,),
                labelText: "密码",
                labelStyle: TextStyle(color: Colors.white),
              ),
              style: TextStyle(color: Colors.white),
              keyboardType: TextInputType.emailAddress,
              obscureText: true,
              onChanged: (str) {
                setState(() {
                  pass = str;
                });
              }
            ),
            padding: EdgeInsets.all(20.0),
          ),
        ],
      ),
      lButtonText: "进入",
      rButtonText: "取消",
      elevation: 10.0,
      lButtonFunc: enter,
      rButtonFunc: cancel,
    );

    return Center(
      child: card,
    );
  }

  void enter() {
    callback(pass);
    cancel();
  }

  void cancel() {
    Navigator.of(context).pop();
  }
}
