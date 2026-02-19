package main

// Joke represents a programming joke
type Joke struct {
    ID       int    `json:"id"`
    Setup    string `json:"setup"`
    Punchline string `json:"punchline"`
    Category string `json:"category"`
}

// Response represents a standard API response
type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
    Success bool   `json:"success"`
    Error   string `json:"error"`
}

// In-memory storage for jokes (in real app, would use a database)
var jokes = []Joke{
    {
        ID:       1,
        Setup:    "Why do programmers prefer dark mode?",
        Punchline: "Because light attracts bugs!",
        Category: "general",
    },
    {
        ID:       2,
        Setup:    "How many programmers does it take to change a light bulb?",
        Punchline: "None. It's a hardware problem!",
        Category: "general",
    },
    {
        ID:       3,
        Setup:    "Why do Java developers wear glasses?",
        Punchline: "Because they don't C#!",
        Category: "languages",
    },
    {
        ID:       4,
        Setup:    "What's a programmer's favorite hangout place?",
        Punchline: "Foo Bar!",
        Category: "general",
    },
    {
        ID:       5,
        Setup:    "Why did the programmer quit his job?",
        Punchline: "Because he didn't get arrays!",
        Category: "general",
    },
}

// nextID tracks the next available ID for new jokes
var nextID = 6