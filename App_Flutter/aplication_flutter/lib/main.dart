import 'dart:convert';

import 'package:aplication_flutter/models/gif.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyApp();
}

class _MyApp extends State<MyApp> {
  Future<List<gif>> _listadoGifs;

  Future<List<gif>> _getGif() async {
    final response = await http.get(Uri.parse("http://localhost:3000/"));

    List<gif> gifs = [];

    if (response.statusCode == 200) {
      String body = utf8.decode(response.bodyBytes);

      final jsonData = jsonDecode(body);
    } else {
      throw Exception("fallo de la conexion");
    }
  }

  @override
  void initState() {
    super.initState();
    _listadoGifs = _getGifs();
  }

  @override
  Widget build(BuildContext context) {
    return Container();
  }
}
