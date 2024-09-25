package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"shorty/configs"

	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jxskiss/base62"
	"golang.org/x/exp/rand"
)

var (
	urlStore = make(map[string]string) // In-memory URL store
	mu       sync.RWMutex              // Mutex for concurrent access
)

var settings = configs.GetSettings()

const baseURL = "http://localhost"

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

	shortURL := generateShortURL(longURL)

	mu.Lock()
	urlStore[shortURL] = longURL
	mu.Unlock()

	response := map[string]string{
		"short_url": fmt.Sprintf("%s:%s/%s", baseURL, settings["PORT"], shortURL),
	}

	log.Printf("new url: %s", response["short_url"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Redirect Handler
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/"):]

	mu.RLock()
	longURL, ok := urlStore[shortURL]
	mu.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

func generateShortURL(longURL string) string {
	// Generate a random salt to add entropy
	rand.Seed(uint64(time.Now().UnixNano()))
	salt := make([]byte, 8)
	rand.Read(salt)

	// Combine the long URL and salt
	input := append([]byte(longURL), salt...)

	// Calculate the SHA-256 hash
	hash := sha256.Sum256(input)

	// Encode the hash using base62 for a shorter and more readable URL
	encodedHash := base62.EncodeToString(hash[:])

	// Take the first 8 characters of the encoded hash as the short URL
	return encodedHash[:8]
}

func main() {
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Printf("Starting server on :%s", settings["PORT"])
	log.Fatal(http.ListenAndServe(":"+settings["PORT"], nil))
}
