import 'package:flutter/material.dart';

import 'package:app/card/card.dart';

class InfoCard extends StatefulWidget {
  const InfoCard({
    Key key,
    @required this.uid,
    @required this.name,
    this.email,
    this.imageUrl = "http://39.98.162.91:9572/files/picture/8c184d0c40897470b10db9e589afc361.jpg",
  }) : super(key: key);

  final int uid;
  final String name;
  final String email;
  final String imageUrl;

  @override
  InfoCardState createState() => InfoCardState();
}

class InfoCardState extends State<InfoCard> {
  @override
  Widget build(BuildContext context) {
    var card = CardWidget(
      Column(
        children: <Widget>[
          Container(
            child: ClipOval(
              child: Image.network(widget.imageUrl, width: 100.0, height: 100.0,),
            ),
            padding: EdgeInsets.all(20.0),
          ),
          Text(widget.name, style: TextStyle(fontSize: 20.0, color: Colors.white),),
          Container(
            child: Text(widget.email ?? "暂未添加邮箱", style: TextStyle(fontSize: 16.0, color: Colors.white),),
            padding: EdgeInsets.all(10.0),
          )
        ],
      ),
      lButtonText: "私聊",
      lButtonFunc: lfunc,
      rButtonText: "了解",
      rButtonFunc: rfunc,

      elevation: 10.0,
    );

    return card;
  }

  void lfunc() async {
  }

  void rfunc() async {
  }
}