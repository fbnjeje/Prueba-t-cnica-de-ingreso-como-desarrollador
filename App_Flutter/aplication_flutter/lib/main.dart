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

  Future<List<gif>> _getGifs() async {
    final response = await http.get("http://localhost:3000/");

    if (response.statusCode == 200) {
      print(response.body);
    } else {
      throw Exception("fallo de la conexion");
    }
  }

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Container();
  }
}
