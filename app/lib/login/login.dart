import 'package:flutter/material.dart';
import 'package:dio/dio.dart';

typedef void HandleAuth(String auth);

class ChatLogin extends StatefulWidget {
  final Dio dio;
  final String host;
  final HandleAuth handleAuth;

  ChatLogin(this.dio, this.host, this.handleAuth);

  @override
  State createState() => _ChatLoginState(dio, host, handleAuth);
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
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Form(
        child: Column(
          children: <Widget>[
            TextField(
              decoration: InputDecoration(labelText: 'Username'),
              onChanged: (String str) => setState(() => username = str),
            ),
            TextField(
              decoration: InputDecoration(labelText: 'Password'),
              onChanged: (String str) => setState(() => password = str),
            ),
            sending
                ? CircularProgressIndicator()
                : RaisedButton(
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
                          builder: (BuildContext context) {
                            return SimpleDialog(
                              title: Text("Login Error: $e"),
                            );
                          },
                        );
                      }).whenComplete(() {
                        setState(() => sending = false);
                      });
                    },
                    color: Theme.of(context).primaryColor,
                    highlightColor: Theme.of(context).highlightColor,
                    child: Text(
                      'SUBMIT',
                      style: TextStyle(color: Colors.white),
                    ),
                  )
          ],
        ),
      ),
    );
  }
}
