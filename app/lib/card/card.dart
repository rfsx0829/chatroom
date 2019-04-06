import 'package:flutter/material.dart';

class CardFulDefault extends StatefulWidget {
  const CardFulDefault() : super();

  @override
  State<StatefulWidget> createState() => _CardFulDefault();
}

class _CardFulDefault extends State {
  @override
  Widget build(BuildContext context) {
    return Card(
      clipBehavior: Clip.antiAlias,
      color: Colors.green,
      elevation: 20.0,
      margin: EdgeInsets.all(20.0),
      semanticContainer: true,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20.0)),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          const ListTile(
            leading: Icon(Icons.access_time, color: Colors.white,),
            title: Text("Hello World", style: TextStyle(color: Colors.white, fontSize: 40.0),),
            subtitle: Text("Sub Title", style: TextStyle(color: Colors.white, fontSize: 20.0),),
            contentPadding: EdgeInsets.all(20.0),
          ),
          ButtonTheme.bar(
            child: ButtonBar(
              children: <Widget>[
                FlatButton(
                  child: Text("滚粗", style: TextStyle(color: Colors.white, fontSize: 28.0),),
                  onPressed: () => {
                    Navigator.of(context).pop()
                  },
                  padding: EdgeInsets.all(10.0),
                ),
                FlatButton(
                  child: Text("了解", style: TextStyle(color: Colors.white, fontSize: 28.0),),
                  onPressed: () => {
                    Navigator.of(context).pop()
                  },
                  padding: EdgeInsets.all(10.0),
                )
              ],
            ),
          )
        ],
      ),
    );
  }
}
