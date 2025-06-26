# 📡 API Documentation

Complete API reference for the Debt Tracker backend.

## 🔗 Base URL

```
http://localhost:8080/api/v1
```

## 🔐 Authentication

All protected endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## 📋 Response Format

### Success Response
```json
{
  "data": {}, // Response data
  "message": "Success message"
}
```

### Error Response
```json
{
  "error": "Error message",
  "code": "ERROR_CODE"
}
```

## 🛡️ Authentication Endpoints

### Register User
```http
POST /auth/register
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123",
  "name": "John Doe",
  "phone": "+1234567890"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "phone": "+1234567890",
    "created_at": "2025-06-15T10:30:00Z",
    "updated_at": "2025-06-15T10:30:00Z"
  }
}
```

### Login User
```http
POST /auth/login
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "phone": "+1234567890",
    "created_at": "2025-06-15T10:30:00Z",
    "updated_at": "2025-06-15T10:30:00Z"
  }
}
```

### Get User Profile
```http
GET /auth/me
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": 1,
  "email": "user@example.com",
  "name": "John Doe",
  "phone": "+1234567890",
  "created_at": "2025-06-15T10:30:00Z",
  "updated_at": "2025-06-15T10:30:00Z"
}
```

## 👥 Contact Endpoints

### Get All Contacts
```http
GET /contacts
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "name": "John Doe",
    "phone": "+1234567890",
    "email": "john@example.com",
    "is_active": true,
    "created_at": "2025-06-15T10:30:00Z",
    "updated_at": "2025-06-15T10:30:00Z"
  }
]
```

### Create Contact
```http
POST /contacts
```

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Jane Smith",
  "phone": "+1987654321",
  "email": "jane@example.com"
}
```

**Response:**
```json
{
  "id": 2,
  "user_id": 1,
  "name": "Jane Smith",
  "phone": "+1987654321",
  "email": "jane@example.com",
  "is_active": true,
  "created_at": "2025-06-15T11:00:00Z",
  "updated_at": "2025-06-15T11:00:00Z"
}
```

### Get Specific Contact
```http
GET /contacts/{id}
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "name": "John Doe",
  "phone": "+1234567890",
  "email": "john@example.com",
  "is_active": true,
  "created_at": "2025-06-15T10:30:00Z",
  "updated_at": "2025-06-15T10:30:00Z"
}
```

### Update Contact
```http
PUT /contacts/{id}
```

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "John Updated",
  "phone": "+1234567890",
  "email": "johnupdated@example.com"
}
```

**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "name": "John Updated",
  "phone": "+1234567890",
  "email": "johnupdated@example.com",
  "is_active": true,
  "created_at": "2025-06-15T10:30:00Z",
  "updated_at": "2025-06-15T11:30:00Z"
}
```

### Delete Contact
```http
DELETE /contacts/{id}
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Contact deleted successfully"
}
```

## 💰 Debt Endpoints

### Get All Debts
```http
GET /debts
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "contact_id": 1,
    "amount": 50.00,
    "direction": "owe_to",
    "status": "active",
    "description": "Lunch money",
    "created_at": "2025-06-15T12:00:00Z",
    "updated_at": "2025-06-15T12:00:00Z",
    "contact": {
      "id": 1,
      "name": "John Doe",
      "phone": "+1234567890",
      "email": "john@example.com"
    }
  }
]
```

### Create Debt
```http
POST /debts
```

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "contact_id": 1,
  "amount": 75.50,
  "direction": "owe_from",
  "description": "Concert tickets"
}
```

**Debt Directions:**
- `owe_to`: You owe money to the contact
- `owe_from`: The contact owes money to you

**Response:**
```json
{
  "id": 2,
  "user_id": 1,
  "contact_id": 1,
  "amount": 75.50,
  "direction": "owe_from",
  "status": "active",
  "description": "Concert tickets",
  "created_at": "2025-06-15T13:00:00Z",
  "updated_at": "2025-06-15T13:00:00Z",
  "contact": {
    "id": 1,
    "name": "John Doe",
    "phone": "+1234567890",
    "email": "john@example.com"
  }
}
```

### Get Debt Summary
```http
GET /debts/summary
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "total_owed_to_others": 50.00,
  "total_owed_from_others": 75.50,
  "net_balance": 25.50,
  "active_debts_count": 2,
  "contacts_with_debts": 1
}
```

### Get Specific Debt
```http
GET /debts/{id}
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "contact_id": 1,
  "amount": 50.00,
  "direction": "owe_to",
  "status": "active",
  "description": "Lunch money",
  "created_at": "2025-06-15T12:00:00Z",
  "updated_at": "2025-06-15T12:00:00Z",
  "contact": {
    "id": 1,
    "name": "John Doe",
    "phone": "+1234567890",
    "email": "john@example.com"
  }
}
```

### Update Debt
```http
PUT /debts/{id}
```

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "amount": 60.00,
  "description": "Updated lunch money",
  "status": "active"
}
```

**Debt Statuses:**
- `active`: Debt is currently active
- `settled`: Debt has been paid off
- `removed`: Debt has been removed from tracking

**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "contact_id": 1,
  "amount": 60.00,
  "direction": "owe_to",
  "status": "active",
  "description": "Updated lunch money",
  "created_at": "2025-06-15T12:00:00Z",
  "updated_at": "2025-06-15T14:00:00Z",
  "contact": {
    "id": 1,
    "name": "John Doe",
    "phone": "+1234567890",
    "email": "john@example.com"
  }
}
```

### Delete Debt
```http
DELETE /debts/{id}
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Debt deleted successfully"
}
```

## 📊 Transaction Endpoints

### Get All Transactions
```http
GET /transactions
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "debt_id": 1,
    "amount": 25.00,
    "transaction_type": "paid_back",
    "description": "Partial payment",
    "created_at": "2025-06-15T15:00:00Z"
  }
]
```

### Create Transaction
```http
POST /transactions
```

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "debt_id": 1,
  "amount": 30.00,
  "transaction_type": "received_back",
  "description": "Payment received"
}
```

**Transaction Types:**
- `lent`: You lent money to someone
- `borrowed`: You borrowed money from someone
- `paid_back`: You paid back money you owed
- `received_back`: You received money someone owed you

**Response:**
```json
{
  "id": 2,
  "debt_id": 1,
  "amount": 30.00,
  "transaction_type": "received_back",
  "description": "Payment received",
  "created_at": "2025-06-15T16:00:00Z"
}
```

### Get Specific Transaction
```http
GET /transactions/{id}
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": 1,
  "debt_id": 1,
  "amount": 25.00,
  "transaction_type": "paid_back",
  "description": "Partial payment",
  "created_at": "2025-06-15T15:00:00Z"
}
```

### Get Transactions for Specific Debt
```http
GET /debts/{debt_id}/transactions
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "debt_id": 1,
    "amount": 25.00,
    "transaction_type": "paid_back",
    "description": "Partial payment",
    "created_at": "2025-06-15T15:00:00Z"
  },
  {
    "id": 2,
    "debt_id": 1,
    "amount": 30.00,
    "transaction_type": "received_back",
    "description": "Payment received",
    "created_at": "2025-06-15T16:00:00Z"
  }
]
```

### Delete Transaction
```http
DELETE /transactions/{id}
```

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Transaction deleted successfully"
}
```

## 🔧 Utility Endpoints

### Health Check
```http
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "debt-tracker-api"
}
```

## ❌ Error Codes

### HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created successfully |
| 400 | Bad Request - Invalid request data |
| 401 | Unauthorized - Invalid or missing authentication |
| 403 | Forbidden - Access denied |
| 404 | Not Found - Resource not found |
| 409 | Conflict - Resource already exists |
| 500 | Internal Server Error - Server error |

### Application Error Codes

| Code | Description |
|------|-------------|
| `INVALID_CREDENTIALS` | Invalid email or password |
| `USER_EXISTS` | User already exists with this email |
| `CONTACT_NOT_FOUND` | Contact not found or doesn't belong to user |
| `DEBT_NOT_FOUND` | Debt not found or doesn't belong to user |
| `TRANSACTION_NOT_FOUND` | Transaction not found or doesn't belong to user |
| `VALIDATION_ERROR` | Request data validation failed |
| `DATABASE_ERROR` | Database operation failed |

## 🔐 Security Considerations

### Authentication
- JWT tokens expire after 7 days
- Use strong passwords (minimum 6 characters)
- Store tokens securely on client side

### Data Protection
- All requests must use HTTPS in production
- Input validation on all endpoints
- SQL injection protection through parameterized queries
- XSS protection through input sanitization

### Rate Limiting
*Not implemented yet - planned for future release*

## 📝 Request Examples

### Using cURL

**Register a new user:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "name": "Test User",
    "phone": "+1234567890"
  }'
```

**Login and save token:**
```bash
TOKEN=$(curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }' | jq -r '.token')
```

**Create a contact:**
```bash
curl -X POST http://localhost:8080/api/v1/contacts \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "phone": "+1987654321",
    "email": "john@example.com"
  }'
```

**Create a debt:**
```bash
curl -X POST http://localhost:8080/api/v1/debts \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "contact_id": 1,
    "amount": 50.00,
    "direction": "owe_to",
    "description": "Lunch money"
  }'
```

### Using JavaScript/Fetch

```javascript
// Login and get token
const loginResponse = await fetch('http://localhost:8080/api/v1/auth/login', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    email: 'test@example.com',
    password: 'password123'
  })
});

const { token } = await loginResponse.json();

// Get all debts
const debtsResponse = await fetch('http://localhost:8080/api/v1/debts', {
  method: 'GET',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json',
  }
});

const debts = await debtsResponse.json();
console.log(debts);
```

## 🧪 Testing the API

### Recommended Tools
- **Postman**: GUI-based API testing
- **Insomnia**: Lightweight API client
- **cURL**: Command-line testing
- **Thunder Client**: VS Code extension

### Postman Collection
*Coming soon - we'll provide a complete Postman collection for easy testing*

### Test Data Setup
```bash
# Complete test workflow
./scripts/test-api.sh
```

*Script coming soon*

---

**This API documentation covers all current endpoints. As new features are added (like bank integration), this documentation will be updated accordingly.**