import 'package:flutter/material.dart';
import 'package:app/person/person.dart';
import 'dart:convert';

class RoomsTab extends StatefulWidget {
  @override
  RoomsState createState() => RoomsState();
}

class RoomsState extends State<RoomsTab> {
  String newVersion = "";

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Container(
          margin: EdgeInsets.all(20.0),
          child: TextField(
            decoration: InputDecoration(
              icon: Icon(Icons.input),
              hintText: "v1.0.1",
              helperText: "Click To Input New Version",
              labelText: "New Version"
            ),
            onSubmitted: (String str) => {newVersion = str},
            onChanged: (String str) => {newVersion = str},
          )
        ),
        FloatingActionButton.extended(
          icon: Icon(Icons.subject),
          label: Text("Submit !", style: TextStyle(fontSize: 30.0),),
          onPressed: updateVersion,
        )
      ],
    );
  }

  void updateVersion() async {
  }
}
