package controllers

import (
	"net/http"
	"strconv"

	"github.com/batt0s/mangajoy/models"
	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserViews struct{}

func (UserViews) RegisterGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user/register", gin.H{
		"title": "Kullanıcı Kayıt",
	})
}

func (UserViews) RegisterPost(ctx *gin.Context) {
	var form models.User
	var err error
	if err = ctx.ShouldBind(&form); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if form.IsValid() {
		avatar, err := ctx.FormFile("avatar")
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		avatar.Filename = settings.UPLOADS_ROOT + "/avatars/" + (strconv.Itoa(int(form.ID)) + form.Username)
		err = ctx.SaveUploadedFile(avatar, avatar.Filename)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		form.Avatar = avatar.Filename
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

func (UserViews) LoginGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user/login", gin.H{
		"title": "Kullanıcı Giriş",
	})
}

func (UserViews) LoginPost(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var username, password string
	username = ctx.PostForm("username")
	password = ctx.PostForm("password")
	user, err := models.Authenticate(username, password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	session.Set(settings.USER_COOKIE, user.Username)
	if err := session.Save(); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.HTML(http.StatusOK, "user/success", gin.H{
		"title":   "Kullanıcı giriş",
		"alert":   "Başarıyla giriş yaptınız, yönlendiriliyorsunuz.",
		"success": 1,
	})
}

func (UserViews) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(settings.USER_COOKIE)
	if user != nil {
		session.Delete(settings.USER_COOKIE)
		if err := session.Save(); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func (UserViews) Dashboard(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.HTML(http.StatusOK, "user/dashboard", gin.H{
		"title": "Dashboard",
		"user":  user.(*models.User),
	})
}
