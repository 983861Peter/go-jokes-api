package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
    // Seed random number generator
    rand.Seed(time.Now().UnixNano())
}

// healthHandler returns API health status
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{
        Success: true,
        Message: "API is running",
        Data: map[string]string{
            "status":  "healthy",
            "version": "1.0.0",
        },
    }
    json.NewEncoder(w).Encode(response)
}

// getAllJokesHandler returns all jokes
func getAllJokesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    response := Response{
        Success: true,
        Message: fmt.Sprintf("Found %d jokes", len(jokes)),
        Data:    jokes,
    }
    json.NewEncoder(w).Encode(response)
}

// getRandomJokeHandler returns a random joke
func getRandomJokeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    if len(jokes) == 0 {
        errorResponse(w, "No jokes available", http.StatusNotFound)
        return
    }
    
    randomIndex := rand.Intn(len(jokes))
    randomJoke := jokes[randomIndex]
    
    response := Response{
        Success: true,
        Data:    randomJoke,
    }
    json.NewEncoder(w).Encode(response)
}

// getJokeByIDHandler returns a joke by ID
func getJokeByIDHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Extract ID from URL path: /api/jokes/3
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 4 {
        errorResponse(w, "Invalid URL format", http.StatusBadRequest)
        return
    }
    
    idStr := pathParts[3]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        errorResponse(w, "Invalid ID format", http.StatusBadRequest)
        return
    }
    
    // Find joke with matching ID
    for _, joke := range jokes {
        if joke.ID == id {
            response := Response{
                Success: true,
                Data:    joke,
            }
            json.NewEncoder(w).Encode(response)
            return
        }
    }
    
    errorResponse(w, "Joke not found", http.StatusNotFound)
}

// createJokeHandler adds a new joke
func createJokeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Only allow POST method
    if r.Method != http.MethodPost {
        errorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var newJoke Joke
    
    // Decode JSON from request body
    err := json.NewDecoder(r.Body).Decode(&newJoke)
    if err != nil {
        errorResponse(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }
    
    // Validate required fields
    if newJoke.Setup == "" || newJoke.Punchline == "" {
        errorResponse(w, "Setup and punchline are required", http.StatusBadRequest)
        return
    }
    
    // Set default category if not provided
    if newJoke.Category == "" {
        newJoke.Category = "general"
    }
    
    // Assign new ID and add to collection
    newJoke.ID = nextID
    nextID++
    jokes = append(jokes, newJoke)
    
    response := Response{
        Success: true,
        Message: "Joke created successfully",
        Data:    newJoke,
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// getJokesByCategoryHandler filters jokes by category
func getJokesByCategoryHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Extract category from URL
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 5 {
        errorResponse(w, "Invalid URL format", http.StatusBadRequest)
        return
    }
    
    category := pathParts[4]
    
    // Filter jokes by category
    var filteredJokes []Joke
    for _, joke := range jokes {
        if joke.Category == category {
            filteredJokes = append(filteredJokes, joke)
        }
    }
    
    response := Response{
        Success: true,
        Message: fmt.Sprintf("Found %d jokes in category '%s'", len(filteredJokes), category),
        Data:    filteredJokes,
    }
    json.NewEncoder(w).Encode(response)
}

// statsHandler returns API statistics
func statsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Count jokes by category
    categoryCount := make(map[string]int)
    for _, joke := range jokes {
        categoryCount[joke.Category]++
    }
    
    stats := map[string]interface{}{
        "total_jokes":      len(jokes),
        "categories":       categoryCount,
        "next_id":         nextID,
    }
    
    response := Response{
        Success: true,
        Data:    stats,
    }
    json.NewEncoder(w).Encode(response)
}

// errorResponse sends a JSON error response
func errorResponse(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    errResp := ErrorResponse{
        Success: false,
        Error:   message,
    }
    json.NewEncoder(w).Encode(errResp)
}