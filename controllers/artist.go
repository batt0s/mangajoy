package controllers

import (
	"net/http"
	"strconv"

	"github.com/batt0s/mangajoy/models"
	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-gonic/gin"
)

type ArtistViews struct{}

func (ArtistViews) Show(ctx *gin.Context) {
	artistid, err := strconv.Atoi(ctx.Param("artistid"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	artist, err := models.GetArtistByID(int64(artistid))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.HTML(http.StatusOK, "artist/show", gin.H{
		"title":  artist.Name,
		"artist": artist,
	})
}

func (ArtistViews) New(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "artist/new", gin.H{
		"title": "Yeni artist",
	})
}

func (ArtistViews) Create(ctx *gin.Context) {
	var artist models.Artist
	var err error
	if err = ctx.ShouldBind(&artist); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !artist.IsValid() {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	pic, err := ctx.FormFile("picture")
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	artistid := strconv.Itoa(int(artist.ID))
	pic.Filename = settings.UPLOADS_ROOT + "/artist/" + artist.Name + artistid
	err = ctx.SaveUploadedFile(pic, pic.Filename)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	artist.Picture = pic.Filename
	err = artist.Create()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/artist/"+artistid)
}
