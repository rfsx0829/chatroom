class User {
  int id;
  String name;
  String email;
  String avatar;

  User({
    this.id,
    this.name,
    this.email,
    this.avatar,
  });

  factory User.fromJson(Map data) {
    return User(
      id: data["id"],
      name: data["name"],
      email: data["email"],
      avatar: data["avatar"],
    );
  }

  Map<String, dynamic> toJson() => {
    "id": id,
    "name": name,
    "email": email,
    "avatar": avatar,
  };

  static User parse(Map map) => User.fromJson(map);
}

class Room {
  int id;
  int nums;
  String name;

  Room({
    this.id,
    this.nums,
    this.name,
  });

  factory Room.fromJson(Map data) {
    return Room(
      id: data["id"],
      nums: data["nums"],
      name: data["name"],
    );
  }

  Map<String, dynamic> toJson() => {
    "id": id,
    "nums": nums,
    "name": name,
  };

  static Room parse(Map map) => Room.fromJson(map);
}

class Message {
  int type;
  String content;
  User user;

  Message({
    this.type,
    this.content,
    this.user,
  });

  factory Message.fromJson(Map data) {
    return Message(
      type: data["type"],
      content: data["content"],
      user: data["user"] == null
        ? null
        : data["user"] is User
          ? data["user"]
          : User.fromJson(data["user"]),
    );
  }

  Map<String, dynamic> toJson() => {
    "type": type,
    "content": content,
    "user": user,
  };

  static Message parse(Map map) => Message.fromJson(map);
}
