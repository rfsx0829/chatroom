import 'dart:convert';

import 'package:app/tabs/rooms.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

import 'package:app/drawer/drawer.dart';
import 'package:app/tabs/messages.dart';
import 'package:app/common/common.dart';

typedef void Func(List<Message> newer);

class ChatApp extends StatefulWidget {
  final WebSocketChannel channel;
  final List<Message> messages;
  final Func updateMessages;
  final User user;
  final Dio dio;
  final String host;
  final ScrollController scrollController;

  ChatApp(this.channel, this.messages, this.updateMessages, this.user, this.dio, this.host, this.scrollController);

  @override
  _ChatAppState createState() => _ChatAppState(channel, messages, updateMessages, user, dio, host, scrollController);
}

class _ChatAppState extends State<ChatApp> with SingleTickerProviderStateMixin {
  TabController controller;
  final WebSocketChannel channel;
  final User user;
  final Dio dio;
  final String host;

  final List<Message> messages;
  final Func updateMessages;
  
  String tempString;
  TextEditingController messageController;
  final ScrollController messageScrollController;

  List<Room> rooms = [];
  Room currentRoom;

  _ChatAppState(this.channel, this.messages, this.updateMessages, this.user, this.dio, this.host, this.messageScrollController);

  @override
  void initState() {
    super.initState();

    controller = TabController(
      length: 2,
      vsync: this,
    );
    messageController = TextEditingController();
  }

  @override
  void dispose() {
    controller.dispose();
    messageController.dispose();
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
        RoomWidget(rooms, () {
          dio.get(host+"/gr").then((data) {
            var obj = jsonDecode(data.data);
            List<Room> list = [];
            for (var item in obj) {
              list.add(Room(id: item["id"], name: item["name"], nums: item["nums"]));
            }
            setState(() => rooms = list);
          });
        }, (Room room, String pass) {
          dio.post(host+"/er", data: {
            "uid": user.id,
            "rid": room.id,
            "pass": pass,
          }).then((res) {
            var obj = jsonDecode(res.data);
            List<Message> messageList = [];
            for (var item in obj) {
              messageList.add(Message(
                content: item["content"],
                type: item["type"],
                user: User.fromJson(item["user"]),
              ));
            }
            updateMessages(messageList);
            setState(() => currentRoom = room);
          }).catchError((e) => showDialog(
            context: context,
            builder: (BuildContext context) => SimpleDialog(
              title: Text("Enter Error $e"),
            )
          ));
        }, host, dio),
        ChatMessageList(messages, user, currentRoom, () {
          if (tempString.isNotEmpty) {
            var m = Message(
              content: tempString,
              user: user,
            );

            channel.sink.add(jsonEncode(m));

            setState(() {tempString = ""; messageController.text = "";});
          }
        }, (String str) => tempString = str, messageController, messageScrollController),
      ]),
      drawer: DrawerWidget(user, dio),
    );
  }
}
