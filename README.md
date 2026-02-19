# Go Jokes API 

A RESTful API built with Go that serves programming jokes. This project demonstrates fundamental backend development concepts including HTTP server implementation, JSON handling, middleware patterns, and proper error handling.

## Features

- **RESTful API Design** - Clean, intuitive endpoints following REST principles
- **JSON Responses** - All data returned in structured JSON format
- **CRUD Operations** - Create, Read, Update (future), Delete (future) jokes
- **Request Logging** - Middleware-based logging of all HTTP requests
- **CORS Support** - Cross-Origin Resource Sharing enabled for browser access
- **Error Handling** - Comprehensive error responses with appropriate HTTP status codes
- **No External Dependencies** - Built entirely with Go's standard library
- **Category Filtering** - Filter jokes by programming language or topic

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21 or higher** - [Download Go](https://go.dev/dl/)
- **Git** - [Download Git](https://git-scm.com/downloads)
- **A code editor** - VS Code recommended with [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go)
- **curl or Postman** - For testing API endpoints

### Verify Installation

```bash
# Check Go version
go version
# Expected output: go version go1.21.x or higher

# Check Git version
git --version
```

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/983861Peter/go-jokes-api.git
cd go-jokes-api
```

### 2. Initialize Go Module

```bash
# If go.mod doesn't exist
go mod init github.com/yourusername/go-jokes-api

# Download dependencies (none required for this project, but good practice)
go mod tidy
```

### 3. Verify Setup

```bash
# Run the application
go run .

# You should see:
# 🚀 Server starting on http://localhost:8080
# 📝 Available endpoints:
#    GET  /health ...
```

##  Quick Start

### Start the Server

```bash
# Option 1: Run without building
go run .

# Option 2: Build and run executable
go build
./go-jokes-api  # On Windows: go-jokes-api.exe
```

The server will start on `http://localhost:8080`

### Test the API

```bash
# Health check
curl http://localhost:8080/health

# Get all jokes
curl http://localhost:8080/api/jokes

# Get a random joke
curl http://localhost:8080/api/jokes/random
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "setup": "Why do programmers prefer dark mode?",
    "punchline": "Because light attracts bugs!",
    "category": "general"
  }
}
```

## API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. Health Check
Check if the API is running.

```http
GET /health
```

**Response:**
```json
{
  "success": true,
  "message": "API is running",
  "data": {
    "status": "healthy",
    "version": "1.0.0"
  }
}
```

---

#### 2. Get All Jokes
Retrieve all jokes in the database.

```http
GET /api/jokes
```

**Response:**
```json
{
  "success": true,
  "message": "Found 5 jokes",
  "data": [
    {
      "id": 1,
      "setup": "Why do programmers prefer dark mode?",
      "punchline": "Because light attracts bugs!",
      "category": "general"
    },
    ...
  ]
}
```

**Example:**
```bash
curl http://localhost:8080/api/jokes
```

---

#### 3. Get Random Joke
Get a randomly selected joke.

```http
GET /api/jokes/random
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 3,
    "setup": "Why do Java developers wear glasses?",
    "punchline": "Because they don't C#!",
    "category": "languages"
  }
}
```

**Example:**
```bash
curl http://localhost:8080/api/jokes/random
```

---

#### 4. Get Joke by ID
Retrieve a specific joke by its ID.

```http
GET /api/jokes/{id}
```

**Parameters:**
- `id` (integer, required) - The joke ID

**Response (Success):**
```json
{
  "success": true,
  "data": {
    "id": 2,
    "setup": "How many programmers does it take to change a light bulb?",
    "punchline": "None. It's a hardware problem!",
    "category": "general"
  }
}
```

**Response (Not Found):**
```json
{
  "success": false,
  "error": "Joke not found"
}
```

**Example:**
```bash
curl http://localhost:8080/api/jokes/2
```

---

#### 5. Create New Joke
Add a new joke to the collection.

```http
POST /api/jokes
Content-Type: application/json
```

**Request Body:**
```json
{
  "setup": "Why do programmers always mix up Halloween and Christmas?",
  "punchline": "Because Oct 31 equals Dec 25!",
  "category": "general"
}
```

**Fields:**
- `setup` (string, required) - The joke setup/question
- `punchline` (string, required) - The joke punchline/answer
- `category` (string, optional) - Joke category (default: "general")

**Response:**
```json
{
  "success": true,
  "message": "Joke created successfully",
  "data": {
    "id": 6,
    "setup": "Why do programmers always mix up Halloween and Christmas?",
    "punchline": "Because Oct 31 equals Dec 25!",
    "category": "general"
  }
}
```

**Example (curl - single line):**
```bash
curl -X POST http://localhost:8080/api/jokes -H "Content-Type: application/json" -d "{\"setup\":\"Why do programmers always mix up Halloween and Christmas?\",\"punchline\":\"Because Oct 31 equals Dec 25!\",\"category\":\"general\"}"
```

**Example (using a file):**
```bash
# Create joke.json with the JSON body above
curl -X POST http://localhost:8080/api/jokes -H "Content-Type: application/json" -d @joke.json
```

**Example (PowerShell):**
```powershell
$body = @{
    setup = "Why do programmers always mix up Halloween and Christmas?"
    punchline = "Because Oct 31 equals Dec 25!"
    category = "general"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/jokes -Method POST -Body $body -ContentType "application/json"
```

---

#### 6. Filter Jokes by Category
Get all jokes in a specific category.

```http
GET /api/jokes/category/{category}
```

**Parameters:**
- `category` (string, required) - Category name (e.g., "general", "languages")

**Response:**
```json
{
  "success": true,
  "message": "Found 2 jokes in category 'languages'",
  "data": [
    {
      "id": 3,
      "setup": "Why do Java developers wear glasses?",
      "punchline": "Because they don't C#!",
      "category": "languages"
    }
  ]
}
```

**Example:**
```bash
curl http://localhost:8080/api/jokes/category/languages
```

---

#### 7. Get API Statistics
Retrieve statistics about the API.

```http
GET /api/stats
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total_jokes": 5,
    "categories": {
      "general": 4,
      "languages": 1
    },
    "next_id": 6
  }
}
```

**Example:**
```bash
curl http://localhost:8080/api/stats
```

---

### Response Format

All API responses follow a consistent structure:

**Success Response:**
```json
{
  "success": true,
  "message": "Optional message",
  "data": { /* Response data */ }
}
```

**Error Response:**
```json
{
  "success": false,
  "error": "Error message description"
}
```

### HTTP Status Codes

| Code | Meaning | When Used |
|------|---------|-----------|
| 200 | OK | Successful GET request |
| 201 | Created | Successful POST request |
| 400 | Bad Request | Invalid input (malformed JSON, missing fields) |
| 404 | Not Found | Resource doesn't exist |
| 405 | Method Not Allowed | Wrong HTTP method for endpoint |
| 500 | Internal Server Error | Server-side error |

## Project Structure

```
go-jokes-api/
├── main.go           # Server setup and route registration
├── handlers.go       # HTTP request handlers for all endpoints
├── models.go         # Data structures (Joke, Response, etc.)
├── middleware.go     # Logging and CORS middleware
├── go.mod           # Go module definition
├── go.sum           # Dependency checksums (if any)
├── README.md        # This file
└── .gitignore       # Git ignore rules
```

### File Descriptions

- **`main.go`** - Entry point of the application. Sets up routes and starts the HTTP server.
- **`handlers.go`** - Contains all HTTP handler functions for processing requests.
- **`models.go`** - Defines data structures (structs) used throughout the application.
- **`middleware.go`** - Middleware functions for logging requests and enabling CORS.
- **`go.mod`** - Defines the module path and Go version requirements.

## Development

### Running in Development Mode

```bash
# Standard run
go run .

# Run with race detection (finds concurrency bugs)
go run -race .

# Run specific files
go run main.go handlers.go models.go middleware.go
```

### Building the Application

```bash
# Build for current platform
go build -o jokes-api

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o jokes-api-linux
GOOS=windows GOARCH=amd64 go build -o jokes-api.exe
GOOS=darwin GOARCH=amd64 go build -o jokes-api-mac
```

### Code Formatting

Go has built-in formatting tools:

```bash
# Format all Go files
go fmt ./...

# Check for common mistakes
go vet ./...
```

### Hot Reload (Optional)

For automatic reloading during development, install `air`:

```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

##  Testing

### Manual Testing with curl

```bash
# Test health endpoint
curl http://localhost:8080/health

# Test GET all jokes
curl http://localhost:8080/api/jokes

# Test POST (create joke) - save this as test-joke.json
echo '{
  "setup": "Test joke setup",
  "punchline": "Test punchline",
  "category": "test"
}' > test-joke.json

curl -X POST http://localhost:8080/api/jokes \
  -H "Content-Type: application/json" \
  -d @test-joke.json

# Test GET by ID
curl http://localhost:8080/api/jokes/1

# Test category filter
curl http://localhost:8080/api/jokes/category/general

# Test stats
curl http://localhost:8080/api/stats
```

### Using Postman

1. Import this collection URL: `http://localhost:8080`
2. Create requests for each endpoint
3. Save as a collection for reuse

### Testing with HTTPie (Alternative to curl)

```bash
# Install HTTPie
pip install httpie

# Test endpoints with cleaner syntax
http GET localhost:8080/api/jokes
http POST localhost:8080/api/jokes setup="Test" punchline="Answer" category="test"
```

### Unit Testing (Future Enhancement)

Create `handlers_test.go`:

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(healthHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}
```

Run tests:
```bash
go test ./...
```

##  Troubleshooting

### Common Issues

#### 1. "go: command not found"
**Problem:** Go is not installed or not in PATH.

**Solution:**
```bash
# Verify Go installation
which go  # macOS/Linux
where go  # Windows

# If not found, reinstall Go from https://go.dev/dl/
# Then restart your terminal
```

---

#### 2. "address already in use"
**Problem:** Port 8080 is occupied by another process.

**Solution:**
```bash
# macOS/Linux - Kill process on port 8080
lsof -ti:8080 | xargs kill -9

# Windows (PowerShell)
Get-Process -Id (Get-NetTCPConnection -LocalPort 8080).OwningProcess | Stop-Process -Force

# Or change port in main.go
port := ":8081"  // Use different port
```

---

#### 3. "cannot find module"
**Problem:** `go.mod` file is missing.

**Solution:**
```bash
# Initialize Go module
go mod init github.com/yourusername/go-jokes-api

# Verify go.mod exists
ls go.mod
```

---

#### 4. "invalid character '<' looking for beginning of value"
**Problem:** Server returned HTML instead of JSON (usually an error page).

**Solution:**
- Check server logs for errors
- Verify the URL is correct
- Ensure Content-Type header is set for POST requests

---

#### 5. POST request returns 404
**Problem:** Handler not properly routing POST requests.

**Solution:**
- Verify the URL exactly matches: `/api/jokes` (no trailing slash)
- Check that your handler checks for `POST` method
- Use `-v` flag with curl to see request details:
```bash
curl -v -X POST http://localhost:8080/api/jokes ...
```

---

#### 6. Empty JSON response: `{}`
**Problem:** Struct fields not exported (not capitalized).

**Solution:**
```go
// ❌ Wrong
type Joke struct {
    id    int    // lowercase - not exported
    setup string
}

// ✅ Correct
type Joke struct {
    ID    int    `json:"id"`    // Capitalized - exported
    Setup string `json:"setup"`
}
```

---

### Getting Help

If you encounter other issues:

1. **Check server logs** - The terminal running the server shows detailed logs
2. **Test with curl first** - Isolate whether it's a client or server issue
3. **Verify JSON format** - Use [JSONLint](https://jsonlint.com/) to validate JSON
4. **Check Go version** - Ensure you're running Go 1.21+
5. **Read error messages** - Go's error messages are usually descriptive

**Still stuck?** Open an issue on GitHub with:
- Your Go version (`go version`)
- Operating system
- Exact error message
- Steps to reproduce

##  Contributing

Contributions are welcome! Here's how you can help:

### Reporting Bugs

1. Check if the issue already exists
2. Include Go version and OS
3. Provide steps to reproduce
4. Include error messages and logs

### Suggesting Features

1. Open an issue describing the feature
2. Explain the use case
3. Provide examples if possible

### Pull Requests

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run `go fmt ./...` and `go vet ./...`
5. Test your changes
6. Commit with clear messages
7. Push to your fork
8. Open a Pull Request

### Code Style

Follow these Go conventions:
- Use `gofmt` for formatting
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Write clear, descriptive commit messages
- Add comments for exported functions
- Keep functions small and focused

##  License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

##  Learning Resources

### Official Documentation
- [Go Official Website](https://go.dev/)
- [A Tour of Go](https://go.dev/tour/) - Interactive tutorial
- [Go by Example](https://gobyexample.com/) - Annotated examples
- [Effective Go](https://go.dev/doc/effective_go) - Style guide

### Related Topics
- [HTTP Package Documentation](https://pkg.go.dev/net/http)
- [JSON Package Documentation](https://pkg.go.dev/encoding/json)
- [REST API Best Practices](https://restfulapi.net/)

##  Acknowledgments

- Built as part of the Moringa School Capstone Project
- Inspired by the Go community's excellent documentation
- Thanks to all contributors and testers

##  Contact

**Project Maintainer:** [Peter Gitau]
- GitHub: [983861Peter](htps://github.com/983861Peter)
- Email: petdev254@gmail.com 

**Project Link:** [https://github.com/983861Peter/go-jokes-api](https://github.com/983861Peter/go-jokes-api)

---

**⭐ If you found this project helpful, please star it on GitHub!**

---

Made with ❤️ using Go