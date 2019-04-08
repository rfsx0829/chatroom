import 'package:flutter/material.dart';
import 'card.dart';

class InfoCard extends StatefulWidget {
  @override
  InfoCardState createState() => InfoCardState();
}

class InfoCardState extends State<InfoCard> {
  @override
  Widget build(BuildContext context) {
    return CardWidget(
      cardChild: Column(
        children: <Widget>[
          Container(
            child: ClipOval(
              child: Image.asset("avatar.jpg", width: 100.0, height: 100.0,),
            ),
            padding: EdgeInsets.all(20.0),
          ),
          Text("Peter", style: TextStyle(fontSize: 20.0, color: Colors.white),),
          Container(
            child: Text("rfsx0829@163.com", style: TextStyle(fontSize: 16.0, color: Colors.white),),
            padding: EdgeInsets.all(10.0),
          )
        ],
      ),
      lButtonText: "私聊",
      lButtonFunc: () => {},
      rButtonText: "了解",
      rButtonFunc: () => {},
    );
  }
}