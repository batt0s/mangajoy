package middlewares

import (
	"net/http"

	"github.com/batt0s/mangajoy/models"
	"github.com/gin-gonic/gin"
)

func StaffOnly(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		ctx.Abort()
		return
	}
	if !user.(*models.User).IsStaff {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	ctx.Next()
}

func AdminOnly(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		ctx.Abort()
		return
	}
	if !user.(*models.User).IsAdmin {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	ctx.Next()
}
