import 'package:flutter/material.dart';

class PersonsTab extends StatefulWidget {
  @override
  PersonsState createState() => PersonsState();
}

class PersonsState extends State<PersonsTab> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Text("Persons"),
    );
  }
}