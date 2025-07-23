class SouthAfricanConfig {
  // Currency settings
  static const String currency = 'ZAR';
  static const String currencySymbol = 'R';
  static const String currencyCode = 'ZAR';
  
  // Date and number formatting
  static const String dateFormat = 'dd/MM/yyyy';
  static const String timeFormat = 'HH:mm';
  static const String phonePrefix = '+27';
  
  // South African banks
  static const List<String> majorBanks = [
    'First National Bank (FNB)',
    'Standard Bank',
    'ABSA Bank',
    'Nedbank',
    'Capitec Bank',
  ];
  
  // Common transaction categories in SA
  static const List<String> commonCategories = [
    'Municipal Bills',
    'Eskom (Electricity)',
    'School Fees',
    'University Fees',
    'DSTV/Streaming',
    'Groceries',
    'Petrol',
    'Rent',
    'Insurance',
    'Medical Aid',
  ];
}
