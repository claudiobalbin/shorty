package routes

import (
	"net/http"
	"shorty/routes/v1/endpoints"
)

func RoutesV1() {
	shortener := endpoints.NewShortenerRoute()
	http.HandleFunc("/shorten", shortener.ShortenURLHandler)
	http.HandleFunc("/", shortener.RedirectHandler)
}
