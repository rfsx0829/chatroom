import 'dart:convert';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:dio/dio.dart';
import 'package:app/common/common.dart';
import 'package:flutter/material.dart';
import 'chat_login.dart';
import 'chat_message_list.dart';
import 'package:app/drawer/drawer.dart';

class ChatHome extends StatefulWidget {
  @override
  State createState() => new _ChatHomeState();
}

class _ChatHomeState extends State<ChatHome> {
  final Dio dio = new Dio(); 
  final String host = 'http://192.168.137.106:8089';

  User user;
  WebSocketChannel wsApp;
  bool connecting = true, error = false;
  List<Message> messages = [];

  @override
  Widget build(BuildContext context) {
    if (user == null) {
      return Scaffold(
        appBar: new AppBar(
          title: new Text('Log In'),
        ),
        body: new ChatLogin(dio, host, (String auth) {
          var obj = jsonDecode(auth);
          setState(() {
            user = User.parse(obj["user"]);
            wsApp = IOWebSocketChannel.connect("ws://192.168.137.106:8089/ac");
          });

          wsApp.stream.listen((mes) {
            var obj = jsonDecode(mes);
            setState(() {
              messages.add(Message.fromJson(obj));
            });
          });

          setState(() => connecting = false);
        }),
      );
    }

    Widget body;

    // Render different content depending on the state of the application.
    if (connecting)
      body = new Text('Connecting to server...');
    else if (error)
      body = new Text('An error occurred while connecting to the server.');
    else {
      body = new ChatMessageList(wsApp, messages, user);
    }

    return new Scaffold(
      appBar: new AppBar(
        title: new Text('Chat (${messages.length} messages)'),
      ),
      body: body,
      drawer: DrawerWidget(user),
    );
  }

  @override
  void dispose() {
    wsApp.sink.close();
    super.dispose();
  }
}
