package main

import (
	"fmt"
	"shorty/configs"
	routes "shorty/routes/api"

	"log"
	"net/http"
)

func main() {
	settings := configs.GetSettings()
	routes.RoutesV1()

	fmt.Printf("Starting server on :%s", settings["PORT"])
	log.Fatal(http.ListenAndServe(":"+settings["PORT"], nil))
}
