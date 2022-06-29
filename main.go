package main

import (
	"log"
	"os"
)

func main() {
	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		appMode = "dev"
		log.Println("[warning] No APP_MODE in env. Defaulting to dev.")
	}

}
