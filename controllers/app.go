// controllers package is for http controllers
// app.go is for App struct
package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/batt0s/mangajoy/database"
	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-gonic/gin"
)

// App struct has address as string and router as gin.Engine
type App struct {
	Addr   string
	Router *gin.Engine
	Server http.Server
}

// Initialize app
func (app *App) Init(mode string) error {

	// gin mode
	switch mode {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}

	// Initialize database
	database.InitDB(mode)

	// Initialize router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Load HTML Templates
	router.LoadHTMLGlob(settings.TEMPLATES_ROOT + "/**/*.gohtml")
	// Static
	router.Static("/static", settings.STATIC_ROOT)

	// Routes
	userGroup := router.Group("/user")
	{
		user := UserViews{}
		userGroup.GET("/register", user.Register)
		userGroup.POST("/register", user.Register)
	}

	app.Router = router

	// Initialize App
	log.Println("App Mode : ", mode)
	var host, port string
	host = "0.0.0.0"
	port = os.Getenv("PORT")
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
