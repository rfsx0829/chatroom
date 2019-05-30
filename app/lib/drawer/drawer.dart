import 'dart:io';

import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

import 'pages/about.dart';
import 'header/header.dart';
import 'package:app/common/common.dart';

class DrawerWidget extends StatelessWidget {
  DrawerWidget(this.user, this.dio);
  final User user;
  final Dio dio;
  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.all(0.0),
        children: <Widget>[
          UserDrawerHeader(user: user,),
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
                builder: (BuildContext context) => AboutPage(dio),
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
