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

## Prompt Journal
This section documents the AI-assisted learning journey from Go beginner to building a working REST API. Each prompt represents a key learning milestone in the project development.

Understanding Go Fundamentals
Curriculum Link: Introduction to Go Programming
Prompt Used:
Explain Go programming language to someone familiar with Python and Java. 
Focus on:
1. Key differences in type systems
2. Compilation vs interpretation
3. Memory management approaches
4. When to choose Go over Python or Java
5. Real-world use cases
AI Response Summary:
The AI explained that Go is statically typed like Java but with cleaner syntax than both languages. Unlike Python's dynamic typing and Java's verbose type declarations, Go uses type inference while maintaining compile-time safety. Go compiles directly to machine code (faster than Java's JVM bytecode), making it ideal for microservices, CLI tools, and cloud infrastructure. Memory is garbage collected but more efficiently than Java. Best use cases: backend APIs, DevOps tools, concurrent systems.
What I Learned:

Go sits between Python's ease and Java's performance
Compilation gives fast startup times crucial for containers
Concurrency is built into the language (goroutines vs threads)

Evaluation: Excellent comparison that positioned Go clearly in the language ecosystem.

Prompt 2: Development Environment Setup
Prompt Used:
How do I set up a Go development environment on Windows?
Include:
1. Installation methods (package manager vs manual)
2. VS Code configuration with Go extension
3. Verifying the installation works
4. Setting up GOPATH vs Go modules
5. Common installation issues and fixes
AI Response Summary:
AI provided step-by-step installation for all platforms. Recommended Homebrew for Mac, apt/manual for Linux, and MSI installer for Windows. Emphasized that Go modules (introduced in Go 1.11) are now standard, so GOPATH is optional. VS Code needs the official Go extension which will prompt to install additional tools. Verification: run go version and create a simple Hello World.
What I Learned:

Modern Go doesn't require GOPATH setup
Go modules make dependency management much easier
VS Code's Go extension auto-installs necessary tools

Evaluation:Clear, actionable steps. Successfully set up environment in 15 minutes.

Prompt 3: First Go Program Structure
Purpose: Understand basic Go syntax
Prompt Used:
Show me the basic structure of a Go program. Explain:
1. What is "package main" and why it's required
2. How import statements work
3. The role of func main()
4. Variable declaration syntax (var vs :=)
5. How Go differs from Python's import system

Include a simple example that prints "Hello, Go!" and declares a variable.
AI Response Summary:
Every Go file starts with a package declaration. package main indicates an executable (entry point). Libraries use other package names. import loads packages (similar to Python but more explicit). func main() is the entry point like Java. Variables can be declared with var name type = value or shorthand :=  (type inferred). Unlike Python's dynamic imports, Go imports are resolved at compile time.
What I Learned:

Package main = executable, other packages = libraries
:= shorthand is commonly used in functions
Imports must be used or code won't compile (stricter than Python)
Capitalized names are exported (public), lowercase is private

Evaluation: Perfect starter explanation with working example.

Prompt 4: Creating HTTP Server
Purpose: Build first web server
Prompt Used:
How does Go's net/http package work? Show me a minimal HTTP server that:
1. Listens on port 8080
2. Has one route returning "Hello, World!"
3. Includes proper error handling
4. Explains the handler function signature

What makes Go's HTTP server different from Python's Flask or Java's Spring?
AI Response Summary:
Go's net/http package is in the standard library (no Flask/Express needed). http.HandleFunc registers routes, http.ListenAndServe starts the server. Handler signature: func(w http.ResponseWriter, r *http.Request). The server is production-ready without frameworks—handles concurrency automatically with goroutines. Unlike Flask's development server, Go's is production-grade.
What I Learned:

No external dependencies needed for HTTP servers
Each request runs in its own goroutine (automatic concurrency)
ResponseWriter is an interface for sending responses
Server blocks on ListenAndServe until shutdown

Evaluation: Mind-blowing that standard library is production-ready. Server running in 10 lines of code.

Prompt 5: JSON Handling in Go
Purpose: Learn to return JSON responses
Prompt Used:
How do I return JSON from a Go HTTP handler? Show me:
1. How to define a struct with JSON tags
2. Converting struct to JSON (marshaling)
3. Setting Content-Type header correctly
4. Handling the order: headers → status → body
5. Best practices for API responses

What's the difference between json.Marshal vs json.Encoder?
AI Response Summary:
Define structs with JSON tags: json:"field_name". Use json.NewEncoder(w).Encode(data) to write JSON directly to response writer (more efficient than Marshal for HTTP). Always set Content-Type: application/json header BEFORE writing body. Marshal returns []byte (for in-memory operations), Encoder writes to io.Writer (better for HTTP). Use json:"omitempty" to skip null fields.
What I Learned:

Struct field names must be capitalized to be exported
JSON tags control the JSON output field names
Encoder is more efficient for HTTP responses than Marshal
Header order matters: set headers before writing body

Evaluation:Critical knowledge. The Marshal vs Encoder distinction saved memory in my API.

Prompt 6: Parsing URL Parameters
Purpose: Handle dynamic routes
Prompt Used:
What's the best way to extract URL parameters in Go's net/http? 
For a route like "/api/jokes/3", I need to extract "3" as the ID.
Show me:
1. Parsing the URL path
2. Converting string to integer safely
3. Handling invalid input (non-numeric IDs)
4. When to use external routers (gorilla/mux) vs manual parsing
AI Response Summary:
Use r.URL.Path to get the path, then strings.Split(path, "/") to break into segments. Access by index but ALWAYS check slice length first to avoid panics. Use strconv.Atoi() to convert string to int—check the error. For simple APIs, manual parsing is fine. For complex routing with many parameters, consider gorilla/mux or chi. Manual parsing keeps zero dependencies.
What I Learned:

strings.Split is the standard approach for simple routing
MUST check slice bounds before accessing elements
strconv.Atoi returns (value, error)—always handle the error
Decision: manual parsing for simplicity, router for complexity

Evaluation:Practical solution. Appreciated the trade-off discussion about routers.

Prompt 7: Designing Data Models
Purpose: Create joke data structure
Prompt Used:
How do I design a Go struct for a REST API? 
For a jokes API, I need:
1. Struct definition with multiple field types
2. JSON tags for API responses
3. Using omitempty for optional fields
4. How to create an in-memory data store (slice of structs)
5. Auto-incrementing ID pattern

Show me best practices for API data models.
AI Response Summary:
Define struct with capitalized field names. Add JSON tags for each field: json:"id". Use omitempty for optional fields: json:"category,omitempty". Store in a package-level slice: var jokes []Joke. For IDs, maintain a nextID counter and increment after each insert. In production, use a database; for learning, in-memory is fine. Consider embedding common fields in a base struct for reuse.
What I Learned:

Struct tags are powerful for controlling JSON output
Package-level variables act as global state (fine for simple apps)
In-memory storage good for prototyping, not production
Go's explicitness makes data models very readable

Evaluation:Clear guidance on structuring data. The nextID pattern worked perfectly.

Prompt 8: Implementing POST Requests
Purpose: Add create functionality
Prompt Used:
How do I handle POST requests in Go to create new resources?
Show me:
1. Reading JSON from request body
2. Decoding into a struct
3. Validating required fields
4. Returning 201 Created status
5. Error handling for invalid JSON

What's the difference between Decode and Unmarshal?
AI Response Summary:
Use json.NewDecoder(r.Body).Decode(&struct) to read request body directly into struct. Check the error—covers both malformed JSON and type mismatches. Validate required fields with simple if statements (or use validator library). Return 201 with w.WriteHeader(http.StatusCreated). Decode reads from io.Reader (request body), Unmarshal works with []byte. Decode is preferred for HTTP.
What I Learned:

Decoder reads directly from request (more efficient)
Error from Decode tells you what's wrong with the JSON
201 Created is the correct status for successful POST
Always validate user input—never trust client data

Evaluation: Essential REST API knowledge. POST endpoint working after this prompt.

Prompt 9: Error Handling Patterns
Purpose: Implement proper error responses
Prompt Used:
What's the proper way to handle errors in Go HTTP handlers?
Show me:
1. Creating consistent error response format
2. Returning appropriate status codes (400, 404, 500)
3. Logging errors server-side vs sending to client
4. When to use panic vs returning error
5. Helper function for error responses

Include examples of different error scenarios.
AI Response Summary:
Create a standard error response struct: {success: false, error: "message"}. Return 400 for client errors (bad input), 404 for not found, 500 for server errors. Log detailed errors server-side with log.Printf(), send user-friendly messages to client. NEVER panic in HTTP handlers—it crashes the server. Create helper: func errorResponse(w, message, statusCode) for consistency.
What I Learned:

Consistent error format improves API usability
Don't expose internal errors to clients (security risk)
Panic is for programmer errors, not user errors
Helper functions reduce code duplication

Evaluation: This improved API quality significantly. Professional error handling established.

Prompt 10: Creating Middleware
Purpose: Add request logging
Prompt Used:
How do I create middleware in Go? 
Show me a logging middleware that:
1. Logs request method and path
2. Measures and logs response time
3. Can be applied to multiple handlers
4. Follows Go's middleware pattern

Explain how middleware wrapping works in Go.
AI Response Summary:
Middleware is a function that takes a handler and returns a wrapped handler: func(http.HandlerFunc) http.HandlerFunc. Inside, capture time.Now() before calling next handler, log details, call next(w, r), then log duration with time.Since(). Apply with: http.HandleFunc("/path", middleware(handler)). Can chain multiple middleware. This pattern uses closures.
What I Learned:

Middleware pattern uses function wrapping (closure)
Order matters: outer middleware runs first/last
time.Now() and time.Since() for measuring duration
Can chain unlimited middleware functions

Evaluation:Elegant pattern. Logging now shows every request with timing.

Prompt 11: CORS Implementation
Purpose: Enable browser access
Prompt Used:
What is CORS and why do I need it? 
For my Go API, show me:
1. What CORS headers are required
2. Handling preflight OPTIONS requests
3. Creating CORS middleware
4. Security considerations (using * vs specific origins)
5. Testing CORS from browser JavaScript

Why does my API work with curl but not browser fetch()?
AI Response Summary:
CORS (Cross-Origin Resource Sharing) is a browser security feature that blocks requests from different origins. Browsers send OPTIONS preflight before POST/PUT. Required headers: Access-Control-Allow-Origin, Access-Control-Allow-Methods, Access-Control-Allow-Headers. Use * for public APIs (development), specific origins for production. curl bypasses CORS (not a browser), fetch respects it. Create middleware to add headers to all responses.
What I Learned:

CORS is browser security, not API security
OPTIONS requests need 200 response with CORS headers
* is convenient but less secure (anyone can call API)
Preflight requests have no body

Evaluation:  Mystery solved! API now accessible from browser. Crucial web development knowledge.

Prompt 12: Code Organization
Purpose: Refactor into multiple files
Prompt Used:
How should I organize a Go REST API project?
I have growing code in main.go. Show me:
1. When to split into multiple files vs packages
2. Naming conventions for files
3. What goes in each file (handlers, models, middleware)
4. Keeping all files in package main vs separate packages
5. Import patterns for multi-file projects

For a small API (5-10 endpoints), what's the best structure?
AI Response Summary:
For small projects, multiple files in package main is fine: main.go (server setup), handlers.go (request handlers), models.go (data structures), middleware.go (middleware functions). No imports needed between files in same package. For larger projects (20+ endpoints), separate into packages: /handlers, /models, /middleware. Keep it simple initially; refactor when complexity justifies it.
What I Learned:

File names should be descriptive and lowercase
Same package = automatic visibility, no imports
Don't over-engineer early; flat structure scales to ~1000 lines
main.go should be minimal (just server setup)

Evaluation:  Pragmatic advice. Refactored into 4 files—much more maintainable.

Prompt 13: Writing API Documentation
Purpose: Create comprehensive API docs
Prompt Used:
Help me document my REST API. 
Show me:
1. Standard format for documenting endpoints
2. How to document request/response examples
3. Including curl commands that actually work
4. Documenting error responses
5. Status code reference table

Provide a template I can follow for each endpoint.
AI Response Summary:
Standard format: HTTP method + path, description, parameters, request body, response format, example curl command. Use code blocks for JSON examples. Include both success and error responses. Create a table of status codes. Template:
### GET /api/resource
Description: ...
Parameters: ...
Response: { example }
Example: curl ...
Consider using API documentation tools like Swagger for larger APIs.
What I Learned:

Consistent format makes API easier to learn
Real curl examples are more valuable than descriptions
Document errors as thoroughly as success cases
Tables improve readability for status codes

Evaluation: Template saved hours. Documentation now looks professional.

Prompt 14: README Best Practices
Purpose: Create project README
Prompt Used:
What makes a great README for a REST API project?
Show me:
1. Essential sections to include
2. Badges and visual elements
3. Quick start that gets someone running in 2 minutes
4. How to balance detail vs brevity
5. Examples of excellent READMEs to reference

The README should work for both users and contributors.
AI Response Summary:
Essential sections: Title/Description, Features, Prerequisites, Installation, Quick Start, API Documentation, Development, Testing, Troubleshooting, Contributing, License. Start with badges (Go version, license). Quick Start should be copy-paste commands. Use emojis sparingly for visual breaks. Include Table of Contents for long READMEs. Good examples: Docker, Kubernetes, go-chi. Balance: detailed enough to get started, not so detailed it overwhelms.
What I Learned:

First impression matters—lead with clear value proposition
Quick Start should deliver working result immediately
Troubleshooting section prevents common support questions
Visual hierarchy (headers, lists, code blocks) improves skimmability

Evaluation:  Comprehensive guidance. README structure now follows industry standards.

Prompt 15: Troubleshooting Common Issues
Purpose: Document common problems
Prompt Used:
What are the most common issues beginners face with Go HTTP servers?
For my troubleshooting section, list:
1. Installation/setup problems
2. Server runtime errors
3. JSON handling issues
4. HTTP-specific errors
5. Development workflow problems

For each: symptom, cause, and solution with commands.
AI Response Summary:
Common issues: "go command not found" (PATH), "address already in use" (port conflict), "cannot find module" (missing go.mod), "empty JSON {}" (unexported fields), "multiple WriteHeader calls" (order violation), "CORS errors" (missing headers). For each: clear symptom, root cause, copy-paste solution. Include platform-specific commands (macOS/Linux/Windows).
What I Learned:

Many errors have standard solutions
Port conflicts are extremely common in development
Struct field capitalization catches everyone initially
Platform differences matter (kill process commands vary)

Evaluation: Invaluable reference. My peer testing found 3 of these issues—already documented.

Prompt 16: Testing Strategies
Purpose: Validate API thoroughly
Prompt Used:
How do I test a Go REST API?
Show me:
1. Manual testing with curl (all HTTP methods)
2. Using Postman effectively
3. Writing basic unit tests with Go's testing package
4. Testing HTTP handlers with httptest
5. What to test (happy path, error cases, edge cases)

Include example test cases for a GET and POST endpoint.
AI Response Summary:
Manual testing: curl for quick checks, Postman for systematic testing with collections. Unit tests: create *_test.go files, use httptest.NewRecorder() to capture responses, test status codes and response bodies. Test coverage: normal inputs, invalid inputs (malformed JSON, missing fields), boundary conditions (empty strings, very long strings), wrong HTTP methods. Use table-driven tests for multiple scenarios.
What I Learned:

httptest package makes testing handlers easy
Test files live alongside code with _test.go suffix
Table-driven tests reduce code duplication
Testing error cases is as important as success cases

Evaluation:Good introduction to Go testing. Wrote basic tests for critical endpoints.

Prompt 17: curl Command Debugging
Purpose: Fix POST request issues
Prompt Used:
My curl POST command isn't working. I'm getting JSON errors and "Bad hostname".
Here's my command: [multi-line curl with backslashes]
Why doesn't this work and what are the alternatives?
Show me:
1. Single-line curl that works everywhere
2. Using a JSON file instead
3. PowerShell-specific syntax
4. How to debug curl commands
AI Response Summary:
Multi-line curl with backslashes (\) only works in bash/zsh. Windows CMD and PowerShell interpret them differently. Solutions: (1) Single-line with escaped quotes, (2) Store JSON in file and use @file.json, (3) PowerShell: use Invoke-RestMethod, (4) Use -v flag to see request details. For cross-platform compatibility, JSON file approach is most reliable.
What I Learned:

Shell differences cause subtle issues
File-based curl is more maintainable
PowerShell has native REST cmdlets (better than curl)
-v flag invaluable for debugging

Evaluation: Solved my immediate problem and taught me platform considerations.

Prompt 18: Code Review & Best Practices
Purpose: Final code quality check
Prompt Used:
Review my Go REST API code for best practices.
Check for:
1. Error handling completeness
2. Code organization and naming
3. Security issues (even for learning projects)
4. Performance considerations
5. Go idioms and conventions

Here's my code: [pasted handlers.go and main.go]
What would make this more production-ready?
AI Response Summary:
Positives: Good error handling, consistent response format, proper status codes. Improvements: (1) Use constants for repeated strings ("general", "languages"), (2) Add request size limits to prevent large POST attacks, (3) Consider context for timeout handling, (4) Add graceful shutdown, (5) Use http.StatusXXX constants instead of numbers. Production additions: rate limiting, authentication, database instead of in-memory storage, proper logging (structured logs), metrics/monitoring, configuration via environment variables.
What I Learned:

Constants improve maintainability and catch typos
Even learning projects should consider basic security
Graceful shutdown prevents data loss
Production-ready requires layers beyond basic functionality

Evaluation: Helpful code review. Implemented constants and improved error messages.

Prompt 19: Deployment Considerations
Purpose: Understand next steps
Prompt Used:
How do I deploy a Go REST API to production?
Explain:
1. Building for different platforms
2. Using environment variables for configuration
3. Popular deployment platforms (Heroku, Railway, fly.io)
4. Docker containerization basics
5. What changes from development to production

I'm not deploying yet, but want to understand the process.
AI Response Summary:
Build with GOOS=linux GOARCH=amd64 go build for Linux servers. Use os.Getenv() for config (ports, database URLs). Platforms: Railway and fly.io are Go-friendly (detect automatically), Heroku needs Procfile. Docker: create Dockerfile with multi-stage build. Production changes: use real database, enable HTTPS, add logging/monitoring, set proper CORS origins, implement authentication, use environment-based config, add health checks for load balancers.
What I Learned:

Cross-compilation is built into Go
Environment variables are standard for config
Modern platforms auto-detect Go apps
Production requires additional infrastructure (logging, monitoring)

Evaluation: Good overview of deployment landscape. Not overwhelming.

Prompt 20: Project Reflection & Documentation
Purpose: Complete project documentation
Prompt Used:
Help me write a compelling project summary for my portfolio.
The project is a Go REST API for learning purposes. Include:
1. What problem it solves
2. Technologies and concepts demonstrated
3. Key learning outcomes
4. Challenges overcome
5. Next steps for improvement

Keep it concise (2-3 paragraphs) but impactful.
AI Response Summary:
Structured project summary emphasizing: (1) Purpose: demonstrating backend development fundamentals while learning Go, (2) Skills: REST API design, HTTP server implementation, JSON encoding, middleware patterns, error handling, API documentation, (3) Learning: transition from Python/Java to Go's concurrency model and type system, using AI as learning accelerator, (4) Challenges: debugging CORS issues, understanding Go's explicit error handling, structuring code effectively, (5) Future: add database persistence, implement authentication, deploy to cloud, add comprehensive testing.
What I Learned:

Portfolio summaries should focus on skills demonstrated
Mentioning challenges shows problem-solving ability
Future improvements show growth mindset
Concrete technologies/concepts more valuable than vague descriptions

Evaluation: Perfect for portfolio. Balances technical detail with accessibility.

## AI Learning Journey Summary
Overall Statistics

Total Prompts Used: 20
Development Time: 5 days
Lines of Code: ~500
Endpoints Created: 7
Documentation Pages: 85+

Key Insights from AI Usage
Most Valuable Prompts:

#5 (JSON Handling) - Fundamental to REST APIs
#11 (CORS) - Solved mysterious browser issues
#8 (POST Requests) - Enabled full CRUD functionality
#17 (curl Debugging) - Immediate practical problem solved

AI Limitations Encountered:

Occasionally suggested complex solutions for simple problems
Some Go version specifics were outdated (verified with official docs)
Didn't always understand skill level initially (asked for clarification)

How AI Improved Learning:

Reduced research time by 70% - Got answers in minutes vs hours of doc reading
Immediate debugging help - Explained errors in context
Comparative learning - Bridged from Python/Java to Go effectively
Best practices early - Learned idiomatic Go from the start
Confidence building - Validated approach at each step

Recommendations for Others:

Start broad, then narrow - "Explain X" → "Show me how to implement Y"
Paste code for debugging - AI excels at spotting issues in context
Ask for comparisons - Leverage existing language knowledge
Request examples - Code examples solidify concepts faster
Verify with official docs - Always cross-reference for accuracy

Prompting Best Practices Discovered
Effective Prompt Structure:
[Clear objective]
Show me: [numbered list of specific items]
[Relevant context about my background/project]
[Specific question or constraint]
Evolution of Prompting:

Day 1: General questions ("How does X work?")
Day 2: Specific implementations ("Show me how to handle POST")
Day 3: Problem-solving ("Why is this code failing?")
Day 4-5: Best practices and polish ("How can I improve this?")


 Learning Resources
Official Documentation

Go Official Website
A Tour of Go - Interactive tutorial
Go by Example - Annotated examples
Effective Go - Style guide

Related Topics

HTTP Package Documentation
JSON Package Documentation
REST API Best Practices
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