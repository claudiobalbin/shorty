package routes

import (
	"fmt"
	"net/http"
	"shorty/configs"
	"shorty/routes/v1/endpoints"
)

var settings = configs.GetSettings()

func RoutesV1() *http.ServeMux {
	shortener := endpoints.NewShortenerRoute()
	handler := http.NewServeMux()
	handler.HandleFunc(fmt.Sprintf("%s/shorten", settings["API_V1"]), shortener.ShortenURLHandler)
	handler.HandleFunc(fmt.Sprintf("%s/", settings["API_V1"]), shortener.RedirectHandler)

	return handler
}
