package middlewares

import (
	"net/http"

	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(settings.USER_COOKIE)
	if user == nil {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		ctx.Abort()
		return
	}
	ctx.Next()
}
