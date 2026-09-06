# Calculator App

A full-stack calculator application built with React, TypeScript, and Go.

## Features

- Addition
- Subtraction
- Multiplication
- Division
- Power
- Square root
- Operator precedence
- Division by zero validation
- Responsive UI
- REST API
- Unit tests for frontend and backend

## Project Structure

```text
calculator-app/
├── backend/
│   ├── cmd/
│   │   └── api/
│   └── internal/
│       └── calculator/
├── frontend/
│   └── src/
│       ├── components/
│       ├── services/
│       └── types/
├── .gitignore
└── README.md
```

The frontend is responsible for the UI, user input, expression evaluation, and communication with the backend.

The backend exposes the REST API and contains the calculator business logic and validation.

## Getting Started

### Prerequisites

- Go 1.27.1
- Node.js 22.13+
- npm

### Backend

From the project root:

```bash
cd backend
go run ./cmd/api
```

The API will run on:

```text
http://localhost:8080
```

### Frontend

Open a second terminal:

```bash
cd frontend
npm install
npm run dev
```

Vite will display the URL where the application is running, usually:

```text
http://localhost:5173
```

The frontend development server proxies `/api` requests to the Go backend.

## API

### Calculate

```http
POST /api/v1/calculator
```

Request:

```json
{
  "operation": "add",
  "operands": [10, 5, 3]
}
```

Response:

```json
{
  "result": 18
}
```

### Supported Operations

| Operation | Description | Operands |
|---|---|---:|
| `add` | Adds all operands | 2+ |
| `subtract` | Subtracts from left to right | 2+ |
| `multiply` | Multiplies all operands | 2+ |
| `divide` | Divides from left to right | 2+ |
| `power` | Raises the first operand to the second | 2 |
| `sqrt` | Calculates the square root | 1 |

### Errors

Errors use the following format:

```json
{
  "error": {
    "code": "DIVISION_BY_ZERO",
    "message": "division by zero"
  }
}
```

Invalid operations or operands return `400 Bad Request`.

## Design Decisions

I kept the architecture simple because the application does not require a database, authentication, or other additional infrastructure.

The backend separates the HTTP layer from the calculator business logic. Each operation implements a common interface, which makes the operations easier to test and extend.

On the frontend, components handle the UI, the API service handles backend communication, and the expression evaluator handles operator precedence.

The calculator uses a single API endpoint because all operations follow the same request/response structure.

## Testing

Backend:

```bash
cd backend
go test ./...
```

Frontend:

```bash
cd frontend
npm run test
```

The tests cover the calculator operations, expression evaluation, operator precedence, and error propagation.