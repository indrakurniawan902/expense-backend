# 📱 Expense API - Mobile Developer Integration Guide

Dokumentasi lengkap untuk mobile developer yang ingin integrate dengan API ini.

---

## Base URL

**Development:**
```
http://localhost:8080
```

**Production (setelah deploy):**
```
https://your-deployed-url
```

---

## API Endpoints

### 1. Health Check
Check apakah API sedang running.

```
GET /health
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Expense API is running",
  "data": {
    "status": "ok"
  }
}
```

---

### 2. Create Expense
Membuat expense baru.

```
POST /expenses
Content-Type: application/json
```

**Request Body:**
```json
{
  "description": "Coffee at Starbucks",
  "amount": 5.5,
  "category": "Food",
  "date": "2025-11-26"
}
```

**Validation Rules:**
- `description`: Required, string
- `amount`: Required, number > 0
- `category`: Required, string
- `date`: Required, format YYYY-MM-DD

**Response (200 OK):**
```json
{
  "statusCode": 200,
  "message": "Expense created successfully",
  "data": {
    "id": 1,
    "description": "Coffee at Starbucks",
    "amount": 5.5,
    "category": "Food",
    "date": "2025-11-26T00:00:00Z",
    "created_at": "2025-11-26T13:47:32.832461+07:00",
    "updated_at": "2025-11-26T13:47:32.832461+07:00"
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "statusCode": 400,
  "message": "Validation error message",
  "data": null
}
```

---

### 3. Get All Expenses
Ambil semua expense.

```
GET /expenses
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "All expenses retrieved successfully",
  "data": [
    {
      "id": 1,
      "description": "Coffee at Starbucks",
      "amount": 5.5,
      "category": "Food",
      "date": "2025-11-26T00:00:00Z",
      "created_at": "2025-11-26T13:47:32.832461+07:00",
      "updated_at": "2025-11-26T13:47:32.832461+07:00"
    },
    {
      "id": 2,
      "description": "Gas",
      "amount": 45.0,
      "category": "Transport",
      "date": "2025-11-25T00:00:00Z",
      "created_at": "2025-11-26T13:47:40.123456+07:00",
      "updated_at": "2025-11-26T13:47:40.123456+07:00"
    }
  ]
}
```

---

### 4. Get Single Expense
Ambil expense berdasarkan ID.

```
GET /expenses/:id
```

**Example:**
```
GET /expenses/1
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Expense retrieved successfully",
  "data": {
    "id": 1,
    "description": "Coffee at Starbucks",
    "amount": 5.5,
    "category": "Food",
    "date": "2025-11-26T00:00:00Z",
    "created_at": "2025-11-26T13:47:32.832461+07:00",
    "updated_at": "2025-11-26T13:47:32.832461+07:00"
  }
}
```

**Error Response (404 Not Found):**
```json
{
  "statusCode": 404,
  "message": "expense not found",
  "data": null
}
```

---

### 5. Update Expense
Update expense yang sudah ada.

```
PUT /expenses/:id
Content-Type: application/json
```

**Example:**
```
PUT /expenses/1
```

**Request Body (semua field optional):**
```json
{
  "description": "Coffee at Starbucks (updated)",
  "amount": 6.5,
  "category": "Food",
  "date": "2025-11-27"
}
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Expense updated successfully",
  "data": {
    "id": 1,
    "description": "Coffee at Starbucks (updated)",
    "amount": 6.5,
    "category": "Food",
    "date": "2025-11-27T00:00:00Z",
    "created_at": "2025-11-26T13:47:32.832461+07:00",
    "updated_at": "2025-11-26T13:47:48.312396+07:00"
  }
}
```

---

### 6. Delete Expense
Hapus expense.

```
DELETE /expenses/:id
```

**Example:**
```
DELETE /expenses/1
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Expense deleted successfully",
  "data": null
}
```

**Error Response (404 Not Found):**
```json
{
  "statusCode": 404,
  "message": "expense not found",
  "data": null
}
}
```

---

## Data Types

### Expense Object
```typescript
{
  id: number,
  description: string,
  amount: number,
  category: string,
  date: string (ISO 8601 datetime),
  created_at: string (ISO 8601 datetime),
  updated_at: string (ISO 8601 datetime)
}
```

---

## Error Handling

### HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK - Request successful |
| 400 | Bad Request - Invalid input |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

### Error Response Format

Semua error response mengikuti format:
```json
{
  "statusCode": <http_status_code>,
  "message": "<error_description>",
  "data": null
}
```

---

## Code Examples

### JavaScript/Fetch API

**Create Expense:**
```javascript
const createExpense = async () => {
  const response = await fetch('https://api-url/expenses', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      description: 'Coffee',
      amount: 5.5,
      category: 'Food',
      date: '2025-11-26'
    })
  });
  
  const result = await response.json();
  if (response.ok) {
    console.log('Expense created:', result.data);
  } else {
    console.error('Error:', result.message);
  }
};
```

**Get All Expenses:**
```javascript
const getExpenses = async () => {
  const response = await fetch('https://api-url/expenses');
  const result = await response.json();
  
  if (response.ok) {
    console.log('Expenses:', result.data);
  }
};
```

---

### Dart/Flutter

**Create Expense:**
```dart
import 'package:http/http.dart' as http;
import 'dart:convert';

Future<void> createExpense() async {
  final response = await http.post(
    Uri.parse('https://api-url/expenses'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({
      'description': 'Coffee',
      'amount': 5.5,
      'category': 'Food',
      'date': '2025-11-26'
    }),
  );

  if (response.statusCode == 200) {
    final json = jsonDecode(response.body);
    print('Expense created: ${json['data']}');
  } else {
    print('Error: ${response.body}');
  }
}
```

**Get All Expenses:**
```dart
Future<void> getExpenses() async {
  final response = await http.get(
    Uri.parse('https://api-url/expenses'),
  );

  if (response.statusCode == 200) {
    final json = jsonDecode(response.body);
    final expenses = json['data'] as List;
    print('Expenses: $expenses');
  }
}
```

---

### Python/Requests

**Create Expense:**
```python
import requests
import json

response = requests.post(
    'https://api-url/expenses',
    json={
        'description': 'Coffee',
        'amount': 5.5,
        'category': 'Food',
        'date': '2025-11-26'
    }
)

if response.status_code == 200:
    data = response.json()
    print(f"Expense created: {data['data']}")
else:
    print(f"Error: {response.json()['message']}")
```

---

### Swift/iOS

```swift
import Foundation

func createExpense() {
    let url = URL(string: "https://api-url/expenses")!
    var request = URLRequest(url: url)
    request.httpMethod = "POST"
    request.addValue("application/json", forHTTPHeaderField: "Content-Type")
    
    let body = [
        "description": "Coffee",
        "amount": 5.5,
        "category": "Food",
        "date": "2025-11-26"
    ] as [String : Any]
    
    request.httpBody = try? JSONSerialization.data(withJSONObject: body)
    
    URLSession.shared.dataTask(with: request) { data, response, error in
        if let data = data {
            let json = try? JSONDecoder().decode(ApiResponse.self, from: data)
            print("Expense created: \(json?.data)")
        }
    }.resume()
}
```

---

## CORS Support

API support CORS, jadi bisa diakses dari:
- Web (JavaScript)
- Mobile (iOS, Android)
- Desktop
- Any client

---

## Rate Limiting & Best Practices

**Current:**
- No rate limiting (tapi akan ditambah di production)

**Best Practices:**
1. Cache data lokal untuk offline support
2. Implement retry logic untuk network errors
3. Validate input sebelum send ke API
4. Handle timeout (recommend 30 detik)
5. Use exponential backoff untuk retry

---

## Troubleshooting

### Connection Error
```
Problem: Can't connect to API
Solution: 
1. Check internet connection
2. Verify API URL is correct
3. Check if server is running
```

### 400 Bad Request
```
Problem: Invalid input
Solution:
1. Check all required fields
2. Validate date format (YYYY-MM-DD)
3. Check amount > 0
4. Verify all fields types
```

### 404 Not Found
```
Problem: Resource not found
Solution:
1. Check expense ID exists
2. Verify you're accessing correct endpoint
```

---

## Support & Questions

Jika ada pertanyaan:
1. Check API documentation (ini file)
2. Test endpoint dengan curl/Postman
3. Check server logs
4. Contact backend developer

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2025-11-26 | Initial release |

---

Last Updated: 2025-11-26
