package settings

var DATABASES = map[string]interface{}{
	"default": map[string]string{
		"NAME": "database.db",
	},
	"dev": map[string]string{
		"NAME": "dev.db",
	},
}

const (
	STATIC_ROOT  string = "static"
	MEDIA_ROOT   string = "static/img"
	UPLOADS_ROOT string = "static/img/uploads"
)
