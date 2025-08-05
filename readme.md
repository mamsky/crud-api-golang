# PT Zahir International - Contact Management API

Proyek ini adalah API RESTful sederhana yang dibangun menggunakan Go (Fiber) untuk mengelola informasi kontak.

## Features

- Buat Kontak
- Validasi Email, Format Telepon Manual
- Get Data Paginasi, Filter
- Pengujian Unit (Hanya Mock, tanpa database)
- Dokumentasi Swagger

## Project Structure

```
📁PT Zahir International
├── config/                # Application configuration
├── controllers/           # API controllers (e.g. contact_controller.go)
├── database/              # DB connection setup
├── docs/                  # Swagger documentation
├── models/                # Data models (e.g. Contact struct)
├── routes/                # API route definitions
├── test/                  # Unit tests and mocks
├── tmp/                   # Temp build files (ignored)
├── utils/                 # Utility functions (e.g. validation, pagination)
├── .env                   # Environment variables
├── .env.example           # Example env file
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
└── main.go                # Entry point
```

## Getting Started

## Clone Repository

```bash
https://github.com/mamsky/crud-api-golang.git
cd crud-api-golang
go mod tidy
```

### Running the Server

```bash
go run main.go
```

### Running Tests

```bash
 go test -v test/contact_controller_test.go
```

## API Documentation

Visit `http://localhost:3000/swagger/index.html` after running the server.

### Example request POST `/contacts`

```json
{
  "name": "Andi",
  "email": "andi@example.com",
  "phone": "08123456789",
  "gender": "male"
}
```
