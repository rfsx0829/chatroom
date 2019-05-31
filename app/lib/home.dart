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
      length: 2,
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
            var id = jsonDecode(res.data)["id"];
            print(id);
          }).catchError((e) {
            showDialog(
              context: context,
              builder: (BuildContext context) {
                return SimpleDialog(
                  title: Text("Error: $e"),
                );
              }
            );
          }).whenComplete(() {
            dio.get(host+"/gr").then((data) {
              var obj = jsonDecode(data.data);
              List<Room> list = [];
              for (var item in obj) {
                list.add(Room(id: item["id"], name: item["name"], nums: item["nums"]));
              }
              setState(() => rooms = list);
            }).catchError((err) {
              showDialog(
                context: context,
                builder: (BuildContext context) {
                  return SimpleDialog(
                    title: Text("Error: $err"),
                  );
                }
              );
            }).whenComplete(() {});
          });
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
