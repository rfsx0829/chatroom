import 'package:flutter/material.dart';
import 'pages/about.dart';

class DrawerWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.all(0.0),
        children: <Widget>[
          UserAccountsDrawerHeader(
            decoration: BoxDecoration(
              color: Colors.green,
            ),
            accountName: Text("Peter"),
            accountEmail: Text("rfsx0829@163.com"),
            currentAccountPicture: CircleAvatar(
              backgroundColor: Colors.white,
            ),
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
            title: Text("退出"),
            subtitle: Text("Exit"),
          )
        ],
      )
    );
  }
}
