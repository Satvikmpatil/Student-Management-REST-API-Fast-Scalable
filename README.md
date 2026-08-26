# Student Management REST API

A fast, scalable REST API for managing student records built with Go.

## Features

- Create and retrieve student records
- SQLite database storage
- Request validation
- Clean architecture with separation of concerns
- Graceful shutdown handling

## Project Structure

```
├── cmd/students-api/       # Application entrypoint
├── config/                 # Configuration files
├── internal/
│   ├── config/            # Configuration loading
│   ├── http/handlier/     # HTTP handlers
│   ├── storage/           # Database layer
│   ├── types/             # Data models
│   └── Utils/response/    # Response utilities
└── storage/               # SQLite database file
```

## Getting Started

### Prerequisites

- Go 1.26+

### Run the Server

```bash
go run cmd/students-api/main.go -config config/local.yaml
```

Server starts at `http://localhost:8082`

## API Endpoints

### Create Student

```bash
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "age": 20}'
```

**Response:**
```json
{"id": 1}
```

### Get Student by ID

```bash
curl http://localhost:8082/api/students/1
```

**Response:**
```json
{"id": 1, "name": "John Doe", "email": "john@example.com", "age": 20}
```

## Configuration

Edit `config/local.yaml`:

```yaml
env: "dev"
storagePath: "storage/storage.db"
HTTPServer:
  address: "localhost:8082"
```

## Tech Stack

- **Go** - Programming language
- **SQLite** - Database
- **go-playground/validator** - Request validation
- **cleanenv** - Configuration management
