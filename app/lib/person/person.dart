class FormData {
  int oper = Oper.SendMes;
  int id = 0;
  String mes = "";

  Map<String, dynamic> toJson() => <String, dynamic> {
    "oper": this.oper,
    "id": this.id,
    "mes": this.mes,
  };
}

class Oper{
  static const Create = 0;
  static const Delete = 1;
  static const Enter = 2;
  static const Leave = 3;
  static const SendMes = 4;
}

/*
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
}*/
