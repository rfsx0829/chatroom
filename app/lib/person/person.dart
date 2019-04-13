import 'dart:io';
import 'dart:convert';

class Person{
  static int uid;
  static String name;
  static String pass;
  static String email;

  static List<String> messageBox;
  static WebSocket conn;
  static int roomID;
  static String roomToken;
  static bool online = false;

  static FormData formData = FormData();
  static dynamic response;

  static void initConn() async {
    conn = await WebSocket.connect("ws://192.168.2.1:8080/api/main");
    conn.listen(onData);
  }

  static void close() async {
    await conn.close();
  }

  static void sendReq() async {
    if(conn == null) {
      initConn();
    }
    var data = jsonEncode(formData);
    conn.add(data);
  }

  static void onData(str) {
    response = jsonDecode(str);
  }
}

class FormData {
  int oper = Oper.DefaultOper;
  UserInfo user = UserInfo();
  RoomInfo room = RoomInfo();
  Message mes = Message();

  Map<String, dynamic> toJson() => <String, dynamic> {
    "oper": this.oper,
    "user": this.user,
    "room": this.room,
    "mes": this.mes,
  };
}

class Oper{
  static const DefaultOper = 0;
  static const SignUp = 1;
  static const SignIn = 2;
  static const AddEmail = 3;
  static const Create = 4;
  static const Enter = 5;
  static const Leave = 6;
  static const SendMes = 7;
  static const SendBox = 8;
  static const GetRoomList = 9;
  static const GetPersonsInRoom = 10;
  static const Close = 11;
}

class UserInfo {
  UserInfo({
    this.uid,
    this.name,
    this.pass,
    this.email = ""
  });
  
  int uid = 0;
  String name = "";
  String pass = "";
  String email = "";

  Map<String, dynamic> toJson() => <String, dynamic> {
    "uid": this.uid,
    "name": this.name,
    "pass": this.pass,
    "email": this.email,
  };
}

class RoomInfo {
  RoomInfo({
    this.rid,
    this.name,
    this.pass
  });

  int rid = 0;
  String name = "";
  String pass = "";

  Map<String, dynamic> toJson() => <String, dynamic> {
    "rid": this.rid,
    "name": this.name,
    "pass": this.pass,
  };
}

class Message {
  String text = "";
  int from = 0;
  int to = 0;
  DateTime time;

  Map<String, dynamic> toJson() => <String, dynamic> {
    "text": this.text,
    "from": this.from,
    "to": this.to,
    "time": this.timeString(),
  };

  String timeString() {
    var str = (this.time ?? DateTime.now()).toIso8601String();
    var sub = str.substring(0, 19) + "+08:00";
    return sub;
  }
}
