package controllers

import (
	"net/http"

	"github.com/batt0s/mangajoy/models"
	"github.com/gin-gonic/gin"
)

type UserViews struct{}

func (uw UserViews) Register(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "user/register", gin.H{
			"title": "Kullanıcı Kayıt",
		})
	}
	if ctx.Request.Method == "POST" {
		var form models.User
		var err error
		if err = ctx.ShouldBind(&form); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if form.IsValid() {
			form.Create()
			ctx.HTML(http.StatusOK, "user/success", gin.H{
				"title":   "Kayıt Başarılı",
				"alert":   "Başarıyla kayıt oldunuz, yönlendiriliyorsunuz.",
				"success": 1,
			})
		} else {
			ctx.HTML(http.StatusBadRequest, "user/success", gin.H{
				"title": "Kullanıcı kayıt",
				"alert": "Kullanıcı geçerli değil. Yönlendiriliyorsunuz.",
			})
		}
	}
}
