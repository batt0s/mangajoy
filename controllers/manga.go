package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/batt0s/mangajoy/models"
	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-gonic/gin"
)

type MangaViews struct{}

func (MangaViews) List(ctx *gin.Context) {
	mangas, err := models.GetLastNManga(20)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.HTML(http.StatusOK, "manga/list", gin.H{
		"title":  "Son mangalar",
		"mangas": mangas,
	})
}

func (MangaViews) Show(ctx *gin.Context) {
	var manga *models.Manga
	id, err := strconv.Atoi(ctx.Param("mangaid"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	manga, err = models.GetMangaByID(int64(id))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.HTML(http.StatusOK, "manga/show", gin.H{
		"title": manga.Title,
		"manga": manga,
	})
}

func (MangaViews) New(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !user.(*models.User).IsStaff {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	artists, err := models.GetAllArtist()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.HTML(http.StatusOK, "manga/new", gin.H{
		"title":   "Manga ekle",
		"artists": artists,
	})
}

func (MangaViews) Create(ctx *gin.Context) {
	var manga models.Manga
	var err error
	if err = ctx.ShouldBind(&manga); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	tags := ctx.PostForm("tags_")
	taglist := strings.Split(tags, ",")
	manga.Tags = taglist
	if !manga.IsValid() {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	coverfile, err := ctx.FormFile("cover")
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	coverfile.Filename = settings.UPLOADS_ROOT + "/covers/" + manga.Title
	err = ctx.SaveUploadedFile(coverfile, coverfile.Filename)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	manga.Cover = coverfile.Filename
	if err = manga.Create(); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	redirectUrl := "/manga/" + strconv.Itoa(int(manga.ID))
	ctx.Redirect(http.StatusCreated, redirectUrl)
}

func (MangaViews) UpdateView(ctx *gin.Context) {
	var manga *models.Manga
	var err error
	mangaid, err := strconv.Atoi(ctx.Param("mangaid"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	manga, err = models.GetMangaByID(int64(mangaid))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.HTML(http.StatusOK, "manga/update", gin.H{
		"title": manga.Title + " - Güncelle",
		"manga": manga,
	})
}

func (MangaViews) Update(ctx *gin.Context) {
	var manga models.Manga
	var err error
	if err = ctx.ShouldBind(&manga); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	tags := ctx.PostForm("tags_")
	taglist := strings.Split(tags, ",")
	manga.Tags = taglist
	if !manga.IsValid() {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	coverfile, err := ctx.FormFile("cover")
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	coverfile.Filename = settings.UPLOADS_ROOT + "/covers/" + manga.Title
	if manga.Cover != coverfile.Filename {
		err = ctx.SaveUploadedFile(coverfile, coverfile.Filename)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		manga.Cover = coverfile.Filename
	}
	if err = manga.Update(); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.Redirect(http.StatusCreated, "/manga/"+strconv.Itoa(int(manga.ID)))
}
