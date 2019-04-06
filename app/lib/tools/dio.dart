import 'package:dio/dio.dart';

class Tools {
  static final dio = new Dio();
  static final host = "http://192.168.2.1:8080";

  static get(String path) async {
    var res = await dio.get(host + path);
    return res;
  }

  static post(String path, dynamic data) async {
    var res = await dio.post(host + path, data: data);
    return res;
  }

  static download(String urlPath, dynamic savePath) async {
    var res = await dio.download(host + urlPath, savePath);
    return res;
  }
}
