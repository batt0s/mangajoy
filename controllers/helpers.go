package controllers

import (
	"github.com/batt0s/mangajoy/models"
	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) (*models.User, bool) {
	user, ok := ctx.Get("user")
	if !ok {
		return nil, false
	}
	return user.(*models.User), true
}
