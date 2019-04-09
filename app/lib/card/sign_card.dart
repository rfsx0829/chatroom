import 'package:flutter/material.dart';
import 'package:app/card/card.dart';
import 'package:app/constant/color.dart';

class SignCard extends StatefulWidget {
  @override
  SignCardState createState() => SignCardState();
}

class SignCardState extends State<SignCard> {
  @override
  Widget build(BuildContext context) {
    var card = CardWidget(
      cardChild: Column(
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          Container(
            child: TextField(
              decoration: InputDecoration(
                prefixIcon: Icon(Icons.account_circle, color: ConstantColor.fontColor,),
                labelText: "用户名",
                labelStyle: TextStyle(color: ConstantColor.fontColor),
              ),
              style: TextStyle(color: ConstantColor.fontColor),
              keyboardType: TextInputType.emailAddress,
            ),
            padding: EdgeInsets.only(top: 20.0, bottom: 10.0, left: 20.0, right: 20.0),
          ),
          Container(
            child: TextField(
              decoration: InputDecoration(
                prefixIcon: Icon(Icons.vpn_key, color: ConstantColor.fontColor,),
                labelText: "密码",
                labelStyle: TextStyle(color: ConstantColor.fontColor),
              ),
              style: TextStyle(color: ConstantColor.fontColor),
              keyboardType: TextInputType.emailAddress,
            ),
            padding: EdgeInsets.all(20.0),
          ),
        ],
      ),
      lButtonText: "注册",
      rButtonText: "登陆",
      elevation: 10.0,
      lButtonFunc: () => {},
      rButtonFunc: () => {},
    );

    return Center(
      child: card,
    );
  }
}
