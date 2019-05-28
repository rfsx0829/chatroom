import 'package:flutter/material.dart';
import 'package:app/common/common.dart';

class UserDrawerHeader extends StatefulWidget {
  const UserDrawerHeader({
    Key key,
    @required this.user,
  }) : super(key: key);

  final User user;

  @override
  DrawerHeaderState createState() => DrawerHeaderState();
}

class DrawerHeaderState extends State<UserDrawerHeader> {
  @override
  Widget build(BuildContext context) {
    return UserAccountsDrawerHeader(
      decoration: BoxDecoration(
        color: Colors.lime,
      ),
      accountName: Text(widget.user.name),
      accountEmail: Text(widget.user.email),
      currentAccountPicture: ClipOval(
        child: Image.network(widget.user.avatar),
      ),
    );
  }
}
