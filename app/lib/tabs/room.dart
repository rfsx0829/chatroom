/*
import 'package:flutter/material.dart';
import 'package:app/person/person.dart';
import 'package:app/constant/color.dart';
import 'dart:convert';

class RoomsTab extends StatefulWidget {
  final List<RoomInfo> roomList = [
    RoomInfo(
      rid: 1,
      name: "Room1",
      pass: "password"
    ),
    RoomInfo(
      rid: 2,
      name: "Room2",
      pass: "PASSWORD"
    )
  ];

  @override
  RoomsState createState() => RoomsState();
}

class RoomsState extends State<RoomsTab> {
  @override
  Widget build(BuildContext context) {
    return ListView(
      children: list2widget(widget.roomList, context),
    );
  }

  List<Widget> list2widget(List<RoomInfo> list, BuildContext ctx) {
    List<Widget> newList = [];
    for (var item in list) {
      newList.add(ListTile(
        leading: Icon(Icons.room),
        title: Text(item.name),
        onTap: () => {
          showInfoDialog(item, ctx)
        },
      ));
    }
    return newList;
  }

  void showInfoDialog(RoomInfo info, BuildContext context) {
    showDialog<void>(
      context: context,
      barrierDismissible: true,
      builder: (context) {
        return StatefulBuilder(
          builder: (context, state) {
            return Dialog(
              backgroundColor: ConstantColor.bgColor,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(20.0)
              ),
              child: Container(
                height: 250.0,
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: <Widget>[
                    Text(info.name, style: TextStyle(fontSize: 40.0, color: ConstantColor.fontColor),),
                    Container(
                      child: TextField(
                        decoration: InputDecoration(
                          prefixIcon: Icon(Icons.vpn_key, color: ConstantColor.fontColor,),
                          labelText: "密码",
                          labelStyle: TextStyle(color: ConstantColor.fontColor)
                        ),
                        style: TextStyle(color: ConstantColor.fontColor),
                        keyboardType: TextInputType.emailAddress,
                        onChanged: (str) => {
                          Person.formData.room.pass = str
                        },
                      ),
                      padding: EdgeInsets.all(20.0),
                    ),
                    ButtonBar(
                      children: <Widget>[
                        FlatButton(
                          child: Text("进入", style: TextStyle(color: ConstantColor.fontColor, fontSize: 20.0),),
                          padding: EdgeInsets.all(10.0),
                          onPressed: () => {
                            print(jsonEncode(Person.formData.room))
                          },
                        ),
                        FlatButton(
                          child: Text("返回", style: TextStyle(color: ConstantColor.fontColor, fontSize: 20.0),),
                          padding: EdgeInsets.all(10.0),
                          onPressed: () => {
                            Navigator.of(context).pop()
                          },
                        )
                      ],
                    )
                  ],
                )
              ),
            );
          },
        );
      }
    );
  }
}
*/
