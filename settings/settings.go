package settings

import (
	"log"
	"os"
)

var BASEPATH string = getBasePath()

var SECRET_KEY string = getSecretFromEnv()

var DATABASES = map[string]interface{}{
	"default": map[string]string{
		"NAME": "database.db",
	},
	"dev": map[string]string{
		"NAME": "dev.db",
	},
}

const (
	TEMPLATES_ROOT string = "templates"
	STATIC_ROOT    string = "static"
	MEDIA_ROOT     string = "static/img"
	UPLOADS_ROOT   string = "static/img/uploads"
)

func getBasePath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while getting BASEPATH\n%s", err.Error())
	}
	return path
}

func getSecretFromEnv() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Println("[warning] No SECRET in env.")
		secret = "ultimatesecretkeyolc0IHuxdn9h8FyIe6GpzQkP3ZbBznJ3"
	}
	return secret
}

const USER_COOKIE string = "user"
