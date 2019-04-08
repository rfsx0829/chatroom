import 'dart:io';

class Person{
  Person(int uid, String name, String pass, String email) {
    this.uid = uid;
    this.name = name;
    this.pass = pass;
    this.email = email;

    this.initConn();
  }

  int uid;
  String name;
  String pass;
  String email;

  List<String> messageBox;
  WebSocket conn;
  FormData formData;
  int roomID;
  String roomToken;

  void initConn() async {
    this.conn = await WebSocket.connect("ws://127.0.0.1:8080/ws");
  }

  void close() async {
    await this.conn.close();
  }

  void signUp() async {
    conn.add()
  }
}

class FormData {
}