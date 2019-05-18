import 'package:flutter/material.dart';
import 'dart:convert';

class AboutPage extends StatefulWidget {
  @override
  AboutState createState() => AboutState();
}

class AboutState extends State<AboutPage> {
  String version = "获取版本号失败...";
  bool versionOK = false;

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
    // TODO:
    dynamic res;
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
