import 'package:flutter/material.dart';
import 'pages/about.dart';
import 'header/header.dart';
import 'dart:io';

class DrawerWidget extends StatefulWidget {
  @override
  DrawerState createState() => DrawerState();
}

class DrawerState extends State<DrawerWidget> {
  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.all(0.0),
        children: <Widget>[
          UserDrawerHeader(
            userName: "Tom",
            userEmail: "someone@163.com",
          ),
          ListTile(
            title: Text("主页"),
            subtitle: Text("Main Page"),
            onTap: () {
              Navigator.of(context).pop();
            },
          ),
          ListTile(
            title: Text("关于"),
            subtitle: Text("About"),
            onTap: () {
              Navigator.of(context).pop();
              Navigator.of(context).push(MaterialPageRoute(
                builder: (BuildContext context) => AboutPage(),
              ));
            },
          ),
          ListTile(
            title: Text("长按退出"),
            subtitle: Text("Exit"),
            onLongPress: () => exit(0),
          )
        ],
      )
    );
  }
}
