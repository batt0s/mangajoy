package controllers

import (
	"net/http"

	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomePageHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get(settings.USER_COOKIE)
	ctx.HTML(http.StatusOK, "homepage", gin.H{"user": username})
}
