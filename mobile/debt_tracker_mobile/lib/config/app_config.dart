class AppConfig {
  // API Configuration
  static const String baseUrl = 'http://localhost:8080/api/v1';
  
  // Development vs Production
  static const bool isDevelopment = true;
  
  // Timeouts
  static const int connectTimeout = 30000; // 30 seconds
  static const int receiveTimeout = 30000;
  
  // Storage keys
  static const String tokenKey = 'auth_token';
  static const String userKey = 'user_data';
  
  // App info
  static const String appName = 'Debt Tracker SA';
  static const String appVersion = '1.0.0';
}
