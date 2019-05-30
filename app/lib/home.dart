import 'dart:convert';

import 'package:app/tabs/rooms.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

import 'package:app/drawer/drawer.dart';
import 'package:app/tabs/messages.dart';
import 'package:app/common/common.dart';

class ChatApp extends StatefulWidget {
  final WebSocketChannel channel;
  final List<Message> messages;
  final User user;
  final Dio dio;
  final String host;

  ChatApp(this.channel, this.messages, this.user, this.dio, this.host);

  @override
  ChatAppState createState() => ChatAppState(channel, messages, user, dio, host);
}

class ChatAppState extends State<ChatApp> with SingleTickerProviderStateMixin {
  TabController controller;
  final WebSocketChannel channel;
  final User user;
  final Dio dio;
  final String host;

  final List<Message> messages;
  String tempString;

  List<Room> rooms = [];

  ChatAppState(this.channel, this.messages, this.user, this.dio, this.host);

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
          icon: Icon(Icons.account_balance),
        ),
        Tab(
          icon: Icon(Icons.list),
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
        backgroundColor: Colors.lime,
      ),
      body: getTabBarView(<StatelessWidget>[
        RoomWidget(rooms, (String name, String pass) {
          var req = {
            "name": name,
            "pass": pass,
          };

          dio.post(host+"/cr", data: req).then((res) {
            var id = res.data["id"];
            setState(() => rooms.add(Room(id: id, name: name, nums: 0)));
          }).catchError((e) {}).whenComplete(() {});
        }),
        ChatMessageList(messages, user, () {
          if (tempString.isNotEmpty) {
            var m = Message(
              content: tempString,
              user: user,
            );

            channel.sink.add(jsonEncode(m));

            setState(() => tempString = "");
          }
        }, (String str) => tempString = str),
      ]),
      drawer: DrawerWidget(user, dio),
    );
  }
}
