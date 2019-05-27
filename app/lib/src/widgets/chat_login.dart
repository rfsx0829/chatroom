import 'package:flutter/material.dart';
import 'package:dio/dio.dart';

typedef void HandleAuth(String auth);

class ChatLogin extends StatefulWidget {
  final Dio dio;
  final String host;
  final HandleAuth handleAuth;

  ChatLogin(this.dio, this.host, this.handleAuth);

  @override
  State createState() => new _ChatLoginState(dio, host, handleAuth);
}

class _ChatLoginState extends State<ChatLogin> {
  final Dio dio;
  final String host;
  final HandleAuth handleAuth;
  String username, password;
  bool sending = false;

  _ChatLoginState(this.dio, this.host, this.handleAuth);

  @override
  Widget build(BuildContext context) {
    return new Padding(
      padding: const EdgeInsets.all(16.0),
      child: new Form(
        child: new Column(
          children: <Widget>[
            new TextField(
              decoration: new InputDecoration(labelText: 'Username'),
              onChanged: (String str) => setState(() => username = str),
            ),
            new TextField(
              decoration: new InputDecoration(labelText: 'Password'),
              onChanged: (String str) => setState(() => password = str),
            ),
            sending
                ? new CircularProgressIndicator()
                : new RaisedButton(
                    onPressed: () {
                      setState(() {
                        sending = true;
                      });

                      dio.post(host+"/au", data: {
                        "name": username,
                        "pass": password,
                      })
                      .then((res) {
                        handleAuth(res.data);
                      }).catchError((e) {
                        showDialog(
                          context: context,
                          child: SimpleDialog(
                            title: Text("Login Error: $e"),
                          )
                        );
                      }).whenComplete(() {
                        setState(() => sending = false);
                      });
                    },
                    color: Theme.of(context).primaryColor,
                    highlightColor: Theme.of(context).highlightColor,
                    child: new Text(
                      'SUBMIT',
                      style: new TextStyle(color: Colors.white),
                    ),
                  )
          ],
        ),
      ),
    );
  }
}
