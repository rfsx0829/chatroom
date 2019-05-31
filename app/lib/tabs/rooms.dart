import 'package:flutter/material.dart';

import 'package:app/common/common.dart';

typedef void RoomFunc(String name, String pass);

class RoomWidget extends StatelessWidget {
  final List<Room> rooms;
  final RoomFunc createRoom;

  RoomWidget(this.rooms, this.createRoom);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton.extended(
        foregroundColor: Colors.white,
        icon: Icon(Icons.flag),
        label: Text("Create", maxLines: 1, style: TextStyle(fontSize: 20.0),),
        onPressed: () {
          var name = "asd";
          var pass = "passss";
          createRoom(name, pass);
        },
      ),
      body: Column(
        children: <Widget>[
          Flexible(
            child: rooms.isEmpty
              ? Text("No room yet ... Try to create one !")
              : ListView.builder(
                itemCount: rooms.length,
                itemBuilder: (_, int i) {
                  return ListTile(
                    title: Text(
                      rooms[i].name,
                      style: TextStyle(fontWeight: FontWeight.bold),
                    ),
                    subtitle: Text("${rooms[i].nums}人在线"),
                  );
                },
              )
          )
        ],
      ),
    );
  }
}
