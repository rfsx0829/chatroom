import 'package:flutter/material.dart';
import 'package:app/card/info_card.dart';

class PersonsTab extends StatefulWidget {
  @override
  PersonsState createState() => PersonsState();
}

class PersonsState extends State<PersonsTab> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: InfoCard(),
    );
  }
}