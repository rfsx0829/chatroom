import 'package:flutter/material.dart';

import 'package:dio/dio.dart';

import 'package:app/card/sign_card.dart';
import 'package:app/common/common.dart';

typedef void VoidFunc();

class RoomWidget extends StatelessWidget {
  final List<Room> rooms;
  final VoidFunc getRoomList;

  final String host;
  final Dio dio;

  RoomWidget(this.rooms, this.getRoomList, this.host, this.dio);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton.extended(
        foregroundColor: Colors.white,
        icon: Icon(Icons.flag),
        label: Text("Create", maxLines: 1, style: TextStyle(fontSize: 20.0),),
        onPressed: () {
          showDialog(
            context: context,
            builder: (BuildContext context) => SignCard(host, dio, getRoomList)
          );
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
