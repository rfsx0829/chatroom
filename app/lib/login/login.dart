import 'dart:io';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:dio/dio.dart';
import 'package:path_provider/path_provider.dart';

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
  TextEditingController nameController, passController;

  bool sending = false;
  String tempPath = "";

  _ChatLoginState(this.dio, this.host, this.handleAuth);

  @override
  void initState() {
    super.initState();
    nameController = TextEditingController();
    passController = TextEditingController();
    initFilePath();
  }

  @override
  void dispose() {
    nameController.dispose();
    passController.dispose();
    super.dispose();
  }

  void initFilePath() async {
    var directory = await getApplicationDocumentsDirectory();
    setState(() {
      tempPath = directory.path;
    });
    try {
      File f = File(tempPath+"/config.json");
      var data = f.readAsStringSync();
      var obj = jsonDecode(data);
      setState(() {
        username = obj["name"];
        nameController.text = username;
        password = obj["pass"];
        passController.text = password;
      });
    }catch(e) {
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Form(
        child: Column(
          children: <Widget>[
            TextField(
              controller: nameController,
              decoration: InputDecoration(labelText: 'Username'),
              onChanged: (String str) => setState(() => username = str),
            ),
            TextField(
              controller: passController,
              decoration: InputDecoration(labelText: 'Password'),
              keyboardType: TextInputType.emailAddress,
              obscureText: true,
              onChanged: (String str) => setState(() => password = str),
            ),
            sending
                ? CircularProgressIndicator()
                : RaisedButton(
                    onPressed: () {
                      setState(() {
                        sending = true;
                      });

                      var obj = {
                        "name": username,
                        "pass": password,
                      };

                      dio.post(host+"/au", data: obj)
                      .then((res) {
                        File f = File(tempPath + "/config.json");
                        f.writeAsStringSync(jsonEncode(obj));

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
