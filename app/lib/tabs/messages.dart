import 'package:flutter/material.dart';

import 'package:app/common/common.dart';

typedef void VoidFunc();
typedef void StrFunc(String str);

class ChatMessageList extends StatelessWidget  {
  final List<Message> messages;
  final User user;
  final VoidFunc sendMes;
  final StrFunc onChanged;

  ChatMessageList(this.messages, this.user, this.sendMes, this.onChanged);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Flexible(
          child: messages.isEmpty
              ? Text('Nobody has said anything yet... Break the silence!')
              : ListView.builder(
                  itemCount: messages.length,
                  itemBuilder: (_, int i) {
                    return ListTile(
                      leading: Image.network(
                          '${messages[i].user.avatar}'),
                      title: Text(
                        messages[i].user.name,
                        style: TextStyle(fontWeight: FontWeight.bold),
                      ),
                      subtitle: Text(messages[i].content),
                    );
                  }),
        ),
        Divider(height: 1.0),
        Container(
          decoration: BoxDecoration(color: Theme.of(context).cardColor),
          child: Padding(
            padding: const EdgeInsets.all(12.0),
            child: TextField(
              decoration: InputDecoration(
                suffixIcon: IconButton(
                  icon: Icon(Icons.send),
                  onPressed: sendMes,
                )
              ),
              onChanged: onChanged,
            ),
          ),
        )
      ],
    );
  }
}
