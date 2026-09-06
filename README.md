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
│       ├── api/
│       │   └── calculator/
│       └── calculator/
│           └── operations/
├── frontend/
│   └── src/
│       ├── components/
│       │   └── Calculator/
│       ├── services/
│       └── types/
├── .gitignore
└── README.md
```

## How It Works

The frontend handles user input and expression evaluation, including operator precedence.

The frontend communicates with the Go backend through a REST API. The backend validates the requested operation and operands, executes the corresponding calculator operation, and returns either the result or a structured error response.

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

Open a second terminal from the project root::

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


## Docker

The application can be run locally using Docker Compose.

From the project root:

```bash
docker compose up --build
```

The application will be available at:

```text
http://localhost:5173
```

Docker Compose runs two services:

- **Frontend:** React application served by Nginx.
- **Backend:** Go REST API.

Nginx serves the frontend and proxies `/api` requests to the Go backend through the Docker network.

To stop the containers:

```bash
docker compose down
```


## API

### Calculate
Calculates the result of the requested operation using the provided operands.

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

Unsupported HTTP methods return `405 Method Not Allowed`.

## Design Decisions

### Backend

The backend uses a simple layered structure that separates HTTP concerns from calculator business logic.

- The API layer handles HTTP requests, JSON serialization, status codes, and error responses.
- The calculator layer contains the business rules and validation.
- Each calculator operation implements a common interface, making operations easier to test and extend.

The architecture intentionally avoids additional layers such as repositories or database infrastructure because the application does not require persistence.

### Frontend

The frontend uses React with TypeScript.

UI components are kept separate from backend communication, while the expression evaluator is isolated from the UI. This allows expression evaluation and operator precedence to be tested independently.

### API Design

The calculator uses a single REST endpoint because all supported operations share the same request and response structure.

The backend is responsible for validating operations and operands so that business rules are not dependent on frontend behavior.

### Error Handling

The backend keeps calculator errors independent from HTTP concerns. The API layer maps those errors to structured JSON responses and appropriate HTTP status codes.

This keeps the calculator business logic independent from the transport layer.

### Scope

The implementation intentionally avoids a database, authentication, or other infrastructure that is not required for the problem domain.


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

Frontend Lint:

```bash
cd frontend
npm run lint
```

Frontend Production Build:

```bash
cd frontend
npm run build
```

The tests cover the calculator operations, expression evaluation, operator precedence, and error propagation.