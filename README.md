# Expense App - RESTful API

A complete RESTful API for managing expenses built with Go and Gin framework.

## Features

- ✅ Create expenses
- ✅ Read/retrieve expenses (single or all)
- ✅ Update existing expenses
- ✅ Delete expenses
- ✅ CORS support
- ✅ Input validation
- ✅ Error handling

## Project Structure

```
expense-backend/
├── main.go                 # Server setup and routes
├── models/
│   └── expense.go         # Data models and request/response types
├── handlers/
│   └── expense.go         # HTTP handlers for CRUD operations
├── repository/
│   └── expense.go         # In-memory data storage and CRUD logic
├── go.mod                 # Go module definition
└── go.sum                 # Dependency checksums
```

## Prerequisites

- Go 1.25.4 or higher
- curl (for testing) or any HTTP client

## Installation

1. Clone or navigate to the project directory:
```bash
cd expense-backend
```

2. Install dependencies:
```bash
go get github.com/gin-gonic/gin@latest
```

3. Build the project:
```bash
go build -o expense-api
```

4. Run the server:
```bash
./expense-api
```

The API will start on `http://localhost:8080`

## API Endpoints

### Health Check
- **GET** `/health` - Check if the API is running
  ```bash
  curl -X GET http://localhost:8080/health
  ```

### Create Expense
- **POST** `/expenses` - Create a new expense
  ```bash
  curl -X POST http://localhost:8080/expenses \
    -H "Content-Type: application/json" \
    -d '{
      "description": "Groceries",
      "amount": 50.5,
      "category": "Food",
      "date": "2025-11-26"
    }'
  ```
  **Request Body:**
  - `description` (string, required) - Description of the expense
  - `amount` (float, required) - Amount spent (must be > 0)
  - `category` (string, required) - Category of expense (e.g., Food, Transport, Entertainment)
  - `date` (string, required) - Date of expense (format: YYYY-MM-DD)

  **Response:** 201 Created
  ```json
  {
    "id": 1,
    "description": "Groceries",
    "amount": 50.5,
    "category": "Food",
    "date": "2025-11-26T00:00:00Z",
    "created_at": "2025-11-26T13:37:17.980089+07:00",
    "updated_at": "2025-11-26T13:37:17.980089+07:00"
  }
  ```

### Get All Expenses
- **GET** `/expenses` - Retrieve all expenses
  ```bash
  curl -X GET http://localhost:8080/expenses
  ```
  **Response:** 200 OK
  ```json
  [
    {
      "id": 1,
      "description": "Groceries",
      "amount": 50.5,
      "category": "Food",
      "date": "2025-11-26T00:00:00Z",
      "created_at": "2025-11-26T13:37:17.980089+07:00",
      "updated_at": "2025-11-26T13:37:17.980089+07:00"
    }
  ]
  ```

### Get Single Expense
- **GET** `/expenses/:id` - Retrieve a specific expense by ID
  ```bash
  curl -X GET http://localhost:8080/expenses/1
  ```
  **Response:** 200 OK
  ```json
  {
    "id": 1,
    "description": "Groceries",
    "amount": 50.5,
    "category": "Food",
    "date": "2025-11-26T00:00:00Z",
    "created_at": "2025-11-26T13:37:17.980089+07:00",
    "updated_at": "2025-11-26T13:37:17.980089+07:00"
  }
  ```

### Update Expense
- **PUT** `/expenses/:id` - Update an existing expense
  ```bash
  curl -X PUT http://localhost:8080/expenses/1 \
    -H "Content-Type: application/json" \
    -d '{
      "description": "Groceries & Fruits",
      "amount": 65.75
    }'
  ```
  **Request Body (all fields optional):**
  - `description` (string) - New description
  - `amount` (float) - New amount (if provided, must be > 0)
  - `category` (string) - New category
  - `date` (string) - New date (format: YYYY-MM-DD)

  **Response:** 200 OK
  ```json
  {
    "id": 1,
    "description": "Groceries & Fruits",
    "amount": 65.75,
    "category": "Food",
    "date": "2025-11-26T00:00:00Z",
    "created_at": "2025-11-26T13:37:17.980089+07:00",
    "updated_at": "2025-11-26T13:37:39.105122+07:00"
  }
  ```

### Delete Expense
- **DELETE** `/expenses/:id` - Delete an expense
  ```bash
  curl -X DELETE http://localhost:8080/expenses/1
  ```
  **Response:** 200 OK
  ```json
  {
    "message": "expense deleted successfully"
  }
  ```

## Error Responses

### Bad Request (400)
```json
{
  "error": "invalid expense ID"
}
```

### Not Found (404)
```json
{
  "error": "expense not found"
}
```

## Testing Examples

1. **Create multiple expenses:**
```bash
curl -X POST http://localhost:8080/expenses \
  -H "Content-Type: application/json" \
  -d '{"description": "Gas", "amount": 45.0, "category": "Transport", "date": "2025-11-25"}'

curl -X POST http://localhost:8080/expenses \
  -H "Content-Type: application/json" \
  -d '{"description": "Movie", "amount": 15.0, "category": "Entertainment", "date": "2025-11-24"}'
```

2. **List all expenses:**
```bash
curl -X GET http://localhost:8080/expenses
```

3. **Update an expense:**
```bash
curl -X PUT http://localhost:8080/expenses/1 \
  -H "Content-Type: application/json" \
  -d '{"amount": 75.0}'
```

4. **Delete an expense:**
```bash
curl -X DELETE http://localhost:8080/expenses/1
```

## Data Model

### Expense
| Field | Type | Description |
|-------|------|-------------|
| id | int | Unique identifier (auto-generated) |
| description | string | Description of the expense |
| amount | float64 | Amount spent |
| category | string | Category of expense |
| date | time.Time | Date of expense |
| created_at | time.Time | Timestamp when created |
| updated_at | time.Time | Timestamp when last updated |

## Future Enhancements

- [ ] Add database support (PostgreSQL, MySQL, SQLite)
- [ ] Add user authentication and authorization
- [ ] Add expense filtering and sorting
- [ ] Add expense categories management
- [ ] Add monthly/yearly expense reports
- [ ] Add data persistence
- [ ] Add unit tests
- [ ] Add API documentation (Swagger/OpenAPI)

## Notes

- The API currently uses in-memory storage, so all data is lost when the server stops
- Dates must be in `YYYY-MM-DD` format
- Amount must be a positive number greater than 0
- All timestamps are returned in ISO 8601 format

## License

MIT
