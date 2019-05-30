import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

class AboutPage extends StatefulWidget {
  final Dio dio;

  AboutPage(this.dio);

  @override
  AboutState createState() => AboutState(dio);
}

class AboutState extends State<AboutPage> {
  String version = "获取版本号失败...";
  bool versionOK = false;
  final Dio dio;

  AboutState(this.dio);

  @override
  Widget build(BuildContext context) {
    if (!versionOK) {
      getVersion();
      versionOK = true;
    }
    return Scaffold(
      body: Center(
        child: Container(
          margin: EdgeInsets.only(top: 200.0),
          child: Column(
            children: <Widget>[
              Text("About Page", style: TextStyle(fontSize: 50.0),),
              Text("Version", style: TextStyle(fontSize: 50.0),),
              Text(version, style: TextStyle(fontSize: 50.0),),
            ],
          )
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      floatingActionButton: FloatButton(),
    );
  }

  void getVersion() async {
    var res = await dio.get("http://localhost:8000/chatroom/latest");
    var data = jsonDecode(res.data.toString());
    setState(() {
      version = data["version"];
    });
  }
}

class FloatButton extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return FloatingActionButton.extended(
      foregroundColor: Colors.white,
      icon: Icon(Icons.flag),
      label: Text("Check For Update !", maxLines: 1, style: TextStyle(fontSize: 20.0),),
      onPressed: () => {
        Scaffold.of(context).showSnackBar(
          SnackBar(
            content: ListTile(
              leading: Icon(Icons.access_time),
              title: Text("还在开发中哦 亲～"),
            ),
            action: SnackBarAction(
              label: "知道了",
              textColor: Colors.white,
              onPressed: () => {
                Navigator.of(context).pop()
              },
            ),
            duration: Duration(seconds: 2),
          )
        )
      },
    );
  }
}
