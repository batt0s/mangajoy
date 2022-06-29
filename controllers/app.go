// controllers package is for http controllers
// app.go is for App struct
package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/batt0s/mangajoy/config"
	"github.com/batt0s/mangajoy/database"
	"github.com/go-gin/gin"
)

// App struct has address as string and router as gin.Engine
type App struct {
	Addr   string
	Router *gin.Engine
	Server http.Server
}

// Initialize app
func (app *App) Init(mode string) error {
	// Load config from config.json
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config files.\nError: %s", err.Error())
	}

	// Initialize database
	database.InitDB(mode)

	// Initialize router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	app.Router = router

	// Initialize App
	log.Println("App Mode : ", mode)
	var host, port string
	if mode == "prod" {
		port = os.Getenv("PORT")
		host = "0.0.0.0"
	} else {
		port = config.Conf.GetString(mode + ".port")
		host = config.Conf.GetString(mode + ".host")
	}
	if port == "" {
		port = "8080"
		log.Println("[warning] No port found. Defaulting to 8080.")
	}
	if host == "" {
		host = "0.0.0.0"
		log.Println("[warning] No host found. Defaulting to 0.0.0.0")
	}
	app.Addr = host + ":" + port
	app.Server = http.Server{
		Addr:    app.Addr,
		Handler: app.Router,
	}

	return nil
}

// Run App
func (app *App) Run() {
	log.Printf("[info] App starting at %s", app.Addr)
	app.Server.ListenAndServe()
}
