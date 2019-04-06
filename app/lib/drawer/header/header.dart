import 'package:flutter/material.dart';
import 'package:app/constant/color.dart';

class UserDrawerHeader extends StatefulWidget {
  const UserDrawerHeader({
    Key key,
    @required this.userName,
    @required this.userEmail,
  }) : super(key: key);

  final String userName;
  final String userEmail;

  @override
  DrawerHeaderState createState() => DrawerHeaderState();
}

class DrawerHeaderState extends State<UserDrawerHeader> {
  @override
  Widget build(BuildContext context) {
    return UserAccountsDrawerHeader(
      decoration: BoxDecoration(
        color: ConstantColor.darkgrey,
      ),
      accountName: Text(widget.userName),
      accountEmail: Text(widget.userEmail),
      currentAccountPicture: ClipOval(
        child: Image.asset("avatar.jpg"),
      ),
    );
  }
}
