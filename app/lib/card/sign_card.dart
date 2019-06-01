import 'package:flutter/material.dart';

import 'package:dio/dio.dart';

import 'package:app/card/card.dart';

typedef void VoidFunc();

class SignCard extends StatefulWidget {
  final Dio dio;
  final String host;
  final VoidFunc callback;

  SignCard(this.host, this.dio, this.callback);

  @override
  SignCardState createState() => SignCardState(host, dio, callback);
}

class SignCardState extends State<SignCard> {
  String name, pass;
  final Dio dio;
  final String host;
  final VoidFunc callback;

  SignCardState(this.host, this.dio, this.callback);

  @override
  Widget build(BuildContext context) {
    var card = CardWidget(
      Column(
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          Container(
            child: TextField(
              decoration: InputDecoration(
                prefixIcon: Icon(Icons.account_circle, color: Colors.white,),
                labelText: "名字",
                labelStyle: TextStyle(color: Colors.white),
              ),
              style: TextStyle(color: Colors.white),
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
      lButtonText: "创建",
      rButtonText: "取消",
      elevation: 10.0,
      lButtonFunc: signUp,
      rButtonFunc: cancel,
    );

    return Center(
      child: card,
    );
  }

  void signUp() async {
    dio.post(host+"/cr", data: {
      "name": name,
      "pass": pass,
    }).catchError((e) => showDialog(
      context: context,
      builder: (BuildContext context) => SimpleDialog(
        title: Text("Create Error $e"),
      )
    )).whenComplete(callback);
    Navigator.of(context).pop();
  }

  void cancel() {
    callback();
    Navigator.of(context).pop();
  }
}
