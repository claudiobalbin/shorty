package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetSettings() map[string]string {
	ignoreEnviroment, err := strconv.ParseBool(os.Getenv("IGNORE_ENVIRONMENT"))
	if !ignoreEnviroment {
		err = godotenv.Load()
	}

	if err != nil {
		log.Panicf("Error loading .env file: %v\n", err)
	}

	settings := make(map[string]string)

	settings["PORT"] = "8080"
	if os.Getenv("PORT") != "" {
		settings["PORT"] = os.Getenv("PORT")
	}

	return settings
}
