import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../config/app_config.dart';

class ApiClient {
  static const String _baseUrl = AppConfig.baseUrl;
  
  // Get stored auth token
  Future<String?> _getToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString(AppConfig.tokenKey);
  }
  
  // Common headers
  Future<Map<String, String>> _getHeaders() async {
    final token = await _getToken();
    final headers = {
      'Content-Type': 'application/json',
    };
    
    if (token != null) {
      headers['Authorization'] = 'Bearer $token';
    }
    
    return headers;
  }
  
  // POST request
  Future<http.Response> post(String endpoint, Map<String, dynamic> data) async {
    final url = Uri.parse('$_baseUrl$endpoint');
    final headers = await _getHeaders();
    
    try {
      final response = await http.post(
        url,
        headers: headers,
        body: json.encode(data),
      ).timeout(Duration(milliseconds: AppConfig.connectTimeout));
      
      return response;
    } catch (e) {
      throw Exception('Network error: $e');
    }
  }
  
  // GET request
  Future<http.Response> get(String endpoint) async {
    final url = Uri.parse('$_baseUrl$endpoint');
    final headers = await _getHeaders();
    
    try {
      final response = await http.get(
        url,
        headers: headers,
      ).timeout(Duration(milliseconds: AppConfig.connectTimeout));
      
      return response;
    } catch (e) {
      throw Exception('Network error: $e');
    }
  }
  
  // PUT request
  Future<http.Response> put(String endpoint, Map<String, dynamic> data) async {
    final url = Uri.parse('$_baseUrl$endpoint');
    final headers = await _getHeaders();
    
    try {
      final response = await http.put(
        url,
        headers: headers,
        body: json.encode(data),
      ).timeout(Duration(milliseconds: AppConfig.connectTimeout));
      
      return response;
    } catch (e) {
      throw Exception('Network error: $e');
    }
  }
  
  // DELETE request
  Future<http.Response> delete(String endpoint) async {
    final url = Uri.parse('$_baseUrl$endpoint');
    final headers = await _getHeaders();
    
    try {
      final response = await http.delete(
        url,
        headers: headers,
      ).timeout(Duration(milliseconds: AppConfig.connectTimeout));
      
      return response;
    } catch (e) {
      throw Exception('Network error: $e');
    }
  }
}
