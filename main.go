package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    // Register routes with middleware
    http.HandleFunc("/health", 
        chainMiddleware(healthHandler, loggingMiddleware, corsMiddleware))
    
    http.HandleFunc("/api/jokes", 
        chainMiddleware(jokesRouter, loggingMiddleware, corsMiddleware))
    
    http.HandleFunc("/api/jokes/random", 
        chainMiddleware(getRandomJokeHandler, loggingMiddleware, corsMiddleware))
    
    http.HandleFunc("/api/jokes/category/", 
        chainMiddleware(getJokesByCategoryHandler, loggingMiddleware, corsMiddleware))
    
    http.HandleFunc("/api/stats", 
        chainMiddleware(statsHandler, loggingMiddleware, corsMiddleware))
    
    // Start server
    port := ":8080"
    fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
    fmt.Println("📝 Available endpoints:")
    fmt.Println("   GET  /health                    - Health check")
    fmt.Println("   GET  /api/jokes                 - Get all jokes")
    fmt.Println("   POST /api/jokes                 - Create new joke")
    fmt.Println("   GET  /api/jokes/random          - Get random joke")
    fmt.Println("   GET  /api/jokes/{id}            - Get joke by ID")
    fmt.Println("   GET  /api/jokes/category/{cat}  - Filter by category")
    fmt.Println("   GET  /api/stats                 - API statistics")
    
    log.Fatal(http.ListenAndServe(port, nil))
}

// jokesRouter routes /api/jokes to appropriate handler based on method and path
func jokesRouter(w http.ResponseWriter, r *http.Request) {
    // Handle POST to /api/jokes
    if r.Method == http.MethodPost {
        createJokeHandler(w, r)
        return
    }
    
    // Handle GET to /api/jokes or /api/jokes/{id}
    if r.Method == http.MethodGet {
        // Check if URL has an ID parameter
        if len(r.URL.Path) > len("/api/jokes/") {
            getJokeByIDHandler(w, r)
        } else {
            getAllJokesHandler(w, r)
        }
        return
    }
    
    // Method not allowed
    errorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
}