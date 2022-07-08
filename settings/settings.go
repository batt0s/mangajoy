package settings

import (
	"log"
	"os"
)

var BASEPATH string = getBasePath()

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
