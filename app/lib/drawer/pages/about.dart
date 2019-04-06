import 'package:flutter/material.dart';

class AboutPage extends StatefulWidget {
  @override
  AboutState createState() => AboutState();
}

class AboutState extends State<AboutPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text("About Page", style: TextStyle(fontSize: 30.0),),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      floatingActionButton: FloatButton(),
    );
  }
}

class FloatButton extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return FloatingActionButton.extended(
      foregroundColor: Colors.white,
      icon: Icon(Icons.flag),
      label: Text("Check For Update !", maxLines: 1, style: TextStyle(fontSize: 20.0),),
      onPressed: () => {
        Scaffold.of(context).showSnackBar(
          SnackBar(
            content: Text("还在开发中哦 亲～"),
            action: SnackBarAction(
              label: "知道了",
              textColor: Colors.white,
              onPressed: () => {},
            ),
            duration: Duration(seconds: 2),
          )
        )
      },
    );
  }
}
