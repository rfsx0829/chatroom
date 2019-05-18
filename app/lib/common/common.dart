class User {
  int id;
  String name;
  String avatar;

  User({
    this.id,
    this.name,
    this.avatar,
  });

  factory User.fromJson(Map data) {
    return new User(
      id: data["id"],
      name: data["name"],
      avatar: data["avatar"],
    );
  }

  Map<String, dynamic> toJson() => {
    "id": id,
    "name": name,
    "avatar": avatar,
  };

  static User parse(Map map) => new User.fromJson(map);
}

class Message {
  int id;
  String str;

  User user;

  Message({
    this.id,
    this.str,
    this.user,
  });

  factory Message.fromJson(Map data) {
    return new Message(
      id: data["id"],
      str: data["str"],
      user: data["user"] == null
        ? null
        : data["user"] is User
          ? data["user"]
          : User.fromJson(data["user"]),
    );
  }

  Map<String, dynamic> toJson() => {
    "id": id,
    "str": str,
    "user": user,
  };

  static Message parse(Map map) => new Message.fromJson(map);
}
