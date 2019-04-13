import 'package:flutter/material.dart';
import 'package:app/card/card.dart';
import 'package:app/constant/color.dart';
import 'package:app/person/person.dart';

class SignCard extends StatefulWidget {
  @override
  SignCardState createState() => SignCardState();
}

class SignCardState extends State<SignCard> {
  String name, pass;

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
              onChanged: (str) {
                setState(() {
                  name = str;
                });
              }
            ),
            padding: EdgeInsets.fromLTRB(20.0, 20.0, 20.0, 10.0)
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
      lButtonText: "注册",
      rButtonText: "登陆",
      elevation: 10.0,
      lButtonFunc: signUp,
      rButtonFunc: signIn,
    );

    return Center(
      child: card,
    );
  }

  void signUp() async {
    Person.formData.oper = Oper.SignUp;
    Person.formData.user.name = name;
    Person.formData.user.pass = pass;

    Person.sendReq();
  }

  void signIn() async {
  }
}
