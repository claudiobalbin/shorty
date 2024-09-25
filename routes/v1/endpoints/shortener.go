package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shorty/configs"
	cache "shorty/repositories"
	services "shorty/services"
)

var settings = configs.GetSettings()

type ShortenerRoute struct {
	ShortenerService services.ShortenerService
	CacheService     cache.CacheService
}

func NewShortenerRoute() *ShortenerRoute {
	service := &ShortenerRoute{
		ShortenerService: *services.NewShortenerService(),
		CacheService:     *cache.NewCacheService(),
	}

	return service
}

// Redirect Handler
func (s *ShortenerRoute) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/"):]

	longURL, ok := s.CacheService.GetUrl(key)
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

// URL Shortener Handler
func (s *ShortenerRoute) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
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

	key := s.ShortenerService.GenerateShortURL(longURL)

	ok = s.CacheService.SetUrl(key, longURL)
	if !ok {
		http.Error(w, "Error creating short url, please try again later", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"short_url": fmt.Sprintf("%s:%s/%s", settings["BASE_URL"], settings["PORT"], key),
	}

	log.Printf("new url: %s", response["short_url"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
