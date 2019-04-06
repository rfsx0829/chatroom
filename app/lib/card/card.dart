import 'package:flutter/material.dart';
import 'package:app/constant/color.dart';

class CardWidget extends StatefulWidget {
  const CardWidget({
    Key key,
    @required this.cardChild,
    this.lButtonText,
    this.lButtonFunc,
    this.rButtonText,
    this.rButtonFunc,
    this.cardColor = ConstantColor.darkgrey,
    this.elevation = 20.0,
  }) : super(key: key);

  final Color cardColor;
  final double elevation;
  final Widget cardChild;

  final String lButtonText;
  final lButtonFunc;

  final String rButtonText;
  final rButtonFunc;

  @override
  CardWidgetState createState() => CardWidgetState();
}

class CardWidgetState extends State<CardWidget> {
  @override
  Widget build(BuildContext context) {
    return Card(
      clipBehavior: Clip.antiAlias,
      color: widget.cardColor,
      elevation: widget.elevation,
      margin: EdgeInsets.all(20.0),
      semanticContainer: true,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20.0)),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          widget.cardChild,
          ButtonTheme.bar(
            child: ButtonBar(
              children: <Widget>[
                FlatButton(
                  child: Text(widget.lButtonText ?? "滚粗", style: TextStyle(color: Colors.white, fontSize: 20.0),),
                  onPressed: widget.lButtonFunc,
                  padding: EdgeInsets.all(10.0),
                ),
                FlatButton(
                  child: Text(widget.rButtonText ?? "了解", style: TextStyle(color: Colors.white, fontSize: 20.0),),
                  onPressed: widget.rButtonFunc,
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
