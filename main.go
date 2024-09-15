package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	urlStore = make(map[string]string) // In-memory URL store
	mu       sync.RWMutex              // Mutex for concurrent access
)

const baseURL = "http://localhost:8080/"

// URL Shortener Handler
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	longURL, ok := request["long_url"]
	if !ok {
		http.Error(w, "Missing long_url field", http.StatusBadRequest)
		return
	}

	shortID := generateShortID(longURL)

	mu.Lock()
	urlStore[shortID] = longURL
	mu.Unlock()

	response := map[string]string{
		"short_url": baseURL + shortID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Redirect Handler
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[len("/"):]

	mu.RLock()
	longURL, ok := urlStore[shortID]
	mu.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

// Generate a short ID based on the long URL (for demo purposes)
func generateShortID(longURL string) string {
	// Simple hash-based ID generation for demo purposes
	return fmt.Sprintf("%x", len(longURL)) // Use length of URL as ID
}

func main() {
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
