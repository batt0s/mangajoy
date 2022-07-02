package controllers

import (
	"net/http"

	"github.com/batt0s/mangajoy/models"
	"github.com/go-gin/gin"
)

type UserViews struct{}

func (uw UserViews) Register(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "user/register", nil)
	}
	if ctx.Request.Method == "POST" {
		var newUser models.User
		var err error
		if err = ctx.Bind(&newUser); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.Redirect(http.StatusOK, "/")
	}
}
