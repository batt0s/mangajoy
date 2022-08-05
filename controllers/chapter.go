package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/batt0s/mangajoy/models"
	"github.com/batt0s/mangajoy/settings"
	"github.com/gin-gonic/gin"
)

type ChapterViews struct{}

func (ChapterViews) Show(ctx *gin.Context) {
	chapterid, _ := strconv.Atoi(ctx.Param("chapterid"))
	chapter, err := models.GetChapterByID(int64(chapterid))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.HTML(http.StatusOK, "chapter/show", gin.H{
		"title":   chapter.Title,
		"chapter": chapter,
	})
}

func (ChapterViews) New(ctx *gin.Context) {
	mangaid, _ := strconv.Atoi(ctx.Param("mangaid"))
	manga, err := models.GetMangaByID(int64(mangaid))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.HTML(http.StatusOK, "chapter/new", gin.H{
		"title":   manga.Title + " - New chapter",
		"mangaid": manga.ID,
	})
}

func (ChapterViews) Create(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !user.(*models.User).IsStaff {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	var chapter models.Chapter
	if err := ctx.ShouldBind(&chapter); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	chapter.UploaderID = user.(*models.User).ID
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	pages := form.File["pages"]
	for _, page := range pages {
		dst := settings.UPLOADS_ROOT + "/chapters/" + strconv.Itoa(int(chapter.MangaID)) + "/" + strconv.Itoa(int(chapter.ID))
		os.MkdirAll(dst, os.ModePerm)
		page.Filename = dst + "/" + page.Filename
		if err := ctx.SaveUploadedFile(page, page.Filename); err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		chapter.Pages = append(chapter.Pages, page.Filename)
	}
	if err := chapter.Create(); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/manga/"+strconv.Itoa(int(chapter.MangaID)))
}
