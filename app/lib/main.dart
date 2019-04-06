import 'package:flutter/material.dart';
import 'tabs/chat.dart';
import 'tabs/room.dart';
import 'tabs/persons.dart';
import 'drawer/drawer.dart';
import 'package:app/constant/color.dart';

void main() => runApp(
  MaterialApp(
    home: MyApp(),
    theme: ThemeData(
    ),
    debugShowCheckedModeBanner: false,
  )
);

class MyApp extends StatefulWidget {
  @override
  MyAppState createState() => MyAppState();
}

class MyAppState extends State<MyApp> with SingleTickerProviderStateMixin {
  TabController controller;

  @override
  void initState() {
    super.initState();

    controller = TabController(
      length: 3,
      vsync: this,
    );
  }

  @override
  void dispose() {
    controller.dispose();
    super.dispose();
  }

  TabBar getTabBar() {
    return TabBar(
      tabs: <Widget>[
        Tab(
          icon: Icon(Icons.room),
        ),
        Tab(
          icon: Icon(Icons.room),
        ),
        Tab(
          icon: Icon(Icons.group),
        )
      ],
      controller: controller,
    );
  }

  TabBarView getTabBarView(var tabs) {
    return TabBarView(
      children: tabs,
      controller: controller,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Hello"),
        bottom: getTabBar(),
        backgroundColor: ConstantColor.darkgrey,
      ),
      body: getTabBarView(<StatefulWidget>[
        ChatTab(),
        RoomsTab(),
        PersonsTab(),
      ]),
      drawer: DrawerWidget(),
    );
  }
}
