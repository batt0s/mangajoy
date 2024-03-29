package middlewares

import (
	"net/http"

	"github.com/batt0s/mangajoy/models"
	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get(settings.USER_COOKIE)
	if username == nil {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		ctx.Abort()
		return
	}
	user, err := models.GetUserWithUsername(username.(string))
	if err != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}
