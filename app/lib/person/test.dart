import 'dart:io';
import 'dart:convert';

void main() async {
  var socket = await WebSocket.connect("ws://localhost:8080/ws");

  socket.listen((data) => {
    print(data)
  });

  var data = jsonEncode({
    "vers": "",
    "pass": ""
  });

  socket.add(data);

  socket.add("Hello World");

  socket.close();
}
