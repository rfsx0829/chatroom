import 'package:flutter/material.dart';

class CardWidget extends StatelessWidget {
  CardWidget(this.cardChild,{
    this.lButtonText,
    this.lButtonFunc,
    this.rButtonText,
    this.rButtonFunc,
    this.cardColor = Colors.lime,
    this.elevation = 20.0,
    this.margin = const EdgeInsets.all(20.0),
  });

  final Color cardColor;
  final Widget cardChild;

  final double elevation;
  final EdgeInsetsGeometry margin;

  final String lButtonText;
  final lButtonFunc;

  final String rButtonText;
  final rButtonFunc;

  @override
  Widget build(BuildContext context) {
    return Card(
      clipBehavior: Clip.antiAlias,
      color: cardColor,
      elevation: elevation,
      margin: margin,
      semanticContainer: true,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20.0)),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          cardChild,
          ButtonTheme.bar(
            child: ButtonBar(
              children: <Widget>[
                FlatButton(
                  child: Text(lButtonText ?? "滚粗", style: TextStyle(color: Colors.white, fontSize: 20.0),),
                  onPressed: lButtonFunc,
                  padding: EdgeInsets.all(10.0),
                ),
                FlatButton(
                  child: Text(rButtonText ?? "了解", style: TextStyle(color: Colors.white, fontSize: 20.0),),
                  onPressed: rButtonFunc,
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
