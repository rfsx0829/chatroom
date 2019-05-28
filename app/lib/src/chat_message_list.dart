import 'dart:convert';

import 'package:app/common/common.dart';
import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class ChatMessageList extends StatelessWidget {
  final WebSocketChannel channel;

  final List<Message> messages;
  final User user;

  ChatMessageList(this.channel, this.messages, this.user);

  @override
  Widget build(BuildContext context) {
    return new Column(
      children: <Widget>[
        new Flexible(
          child: messages.isEmpty
              ? new Text('Nobody has said anything yet... Break the silence!')
              : new ListView.builder(
                  itemCount: messages.length,
                  itemBuilder: (_, int i) {
                    return new ListTile(
                      leading: new Image.network(
                          '${messages[i].user.avatar}'),
                      title: new Text(
                        messages[i].user.name,
                        style: new TextStyle(fontWeight: FontWeight.bold),
                      ),
                      subtitle: new Text(messages[i].content),
                    );
                  }),
        ),
        new Divider(height: 1.0),
        new Container(
          decoration: new BoxDecoration(color: Theme.of(context).cardColor),
          child: new Padding(
            padding: const EdgeInsets.only(left: 8.0, right: 8.0),
            child: new TextField(
              decoration: new InputDecoration(labelText: 'Send a message...'),
              onSubmitted: (String msg) {
                if (msg.isNotEmpty) {
                  var m = Message(
                    content: msg,
                    user: user,
                  );
                  
                  channel.sink.add(jsonEncode(m));
                }
              },
            ),
          ),
        )
      ],
    );
  }
}