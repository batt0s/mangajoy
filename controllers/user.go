package controllers

import (
	"net/http"

	"github.com/batt0s/mangajoy/models"
	"github.com/go-gin/gin"
)

func RegisterController(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload.",
			"error":   err.Error(),
		})
		return
	}
	if newUser.IsValid() {
		newUser.Save()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User created succesfully.",
			"user":    newUser,
		})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "User couldn't created.",
			"error":   "User is not valid.",
		})
	}
}
