import 'package:flutter/material.dart';
import 'package:app/common/common.dart';

class UserDrawerHeader extends StatelessWidget {
  final User user;

  UserDrawerHeader(this.user);

  @override
  Widget build(BuildContext context) {
    return UserAccountsDrawerHeader(
      decoration: BoxDecoration(
        color: Colors.lime,
      ),
      accountName: Text(user.name),
      accountEmail: Text(user.email),
      currentAccountPicture: ClipOval(
        child: Image.network(user.avatar),
      ),
    );
  }
}
