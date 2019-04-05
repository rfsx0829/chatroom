import 'package:flutter/material.dart';

class DrawerWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: Column(
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
          )
        ],
      )
    );
  }
}