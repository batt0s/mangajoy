package main

import (
	"fmt"
	"log"
	"os"

	"github.com/batt0s/mangajoy/controllers"
	"github.com/batt0s/mangajoy/database"
	"github.com/batt0s/mangajoy/models"
)

func main() {
	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		appMode = "dev"
		log.Println("[warning] No APP_MODE in env. Defaulting to dev.")
	}
	if len(os.Args) < 2 {
		printhelp()
		os.Exit(0)
	} else {
		if os.Args[1] == "runserver" {
			run(appMode)
		}
		if os.Args[1] == "migrate" {
			migrate(appMode)
		}
	}

}

func run(mode string) {
	app := new(controllers.App)
	err := app.Init(mode)
	if err != nil {
		log.Printf("[error] Error while init app. %s", err.Error())
		os.Exit(1)
	}
	app.Run()
}

func migrate(mode string) {
	log.Println("Migrating models...")
	err := database.InitDB(mode)
	if err != nil {
		log.Printf("Error while init database.\n%s", err.Error())
		os.Exit(1)
	}
	models.Migrate()
	log.Println("Done.")
}

func printhelp() {
	fmt.Println(`
	MangaJoy Server
	
To run: mangajoy runserver
To migrate models: mangajoy migrate 
	`)
}
