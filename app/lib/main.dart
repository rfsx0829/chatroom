import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

import 'home.dart';
import 'login/login.dart';
import 'package:app/common/common.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(primarySwatch: Colors.lime),
      home: ChatHome(),
    );
  }
}

class ChatHome extends StatefulWidget {
  @override
  State createState() => new _ChatHomeState();
}

class _ChatHomeState extends State<ChatHome> {
  final Dio dio = new Dio();
  final String host = 'http://39.98.162.91:8089';

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
        body: new ChatLogin(dio, host, (String auth) async {
          var obj = jsonDecode(auth);
          setState(() {
            obj["id"] = int.parse(obj["id"]);
            user = User.parse(obj);
            wsApp = IOWebSocketChannel.connect("ws://39.98.162.91:8089/ac");
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

    if (connecting)
      body = Text('Connecting to server...');
    else if (error)
      body = Text('An error occurred while connecting to the server.');
    else {
      body = ChatApp(wsApp, messages, user, dio);
    }

    return body;
  }

  @override
  void dispose() {
    if(wsApp != null) {
      wsApp.sink.close();
    }
    super.dispose();
  }
}

