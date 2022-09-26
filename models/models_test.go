package models_test

import (
	"log"
	"os"
	"testing"

	"github.com/batt0s/mangajoy/database"
	"github.com/batt0s/mangajoy/models"
)

func TestMain(m *testing.M) {
	log.Println("Starting testing package models.")
	err := database.InitDB("test")
	if err != nil {
		log.Println("Cannot Init Database")
		os.Exit(1)
	}
	models.Migrate()
	exitVal := m.Run()
	log.Println("Done testing package models.")
	err = os.Remove("test.db")
	if err != nil {
		log.Println("Could not delete test.db")
	}
	os.Exit(exitVal)
}

var (
	// user.go test vars
	username string       = "testuser"
	email    string       = "test@test.com"
	password string       = "testpass"
	avatar   string       = "test.jpg"
	testUser *models.User = &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Avatar:   avatar,
		IsAdmin:  false,
		IsStaff:  false,
	}
	// artist.go test vars
	name       string         = "Test Artist"
	about      string         = "This is the about section of Test Artist."
	image      string         = "test.jpg"
	testArtist *models.Artist = &models.Artist{
		Name:    name,
		About:   about,
		Picture: image,
	}
	// manga.go test vars
	title             string        = "Test Manga"
	alternative_title string        = "ななつのたいざ"
	description       string        = "This is the description section of Test Manga."
	tags              []string      = []string{"shounen", "ecchi"}
	cover             string        = "test.jpg"
	testManga         *models.Manga = &models.Manga{
		Title:            title,
		AlternativeTitle: alternative_title,
		Description:      description,
		Tags:             tags,
		Cover:            cover,
	}
	// chapter.go test vars
	titlec       string          = "Test Manga Chapter 1"
	descriptionc string          = "Test Manga Chapter 1\nKarakterimiz bişiler yapıyor."
	pages        []string        = []string{"1.jpg", "2.jpg"}
	testChapter  *models.Chapter = &models.Chapter{
		Title:       titlec,
		Description: descriptionc,
		Pages:       pages,
	}
	// comment.go test vars
	comment     string          = "güzel chapter"
	testComment *models.Comment = &models.Comment{
		Content: comment,
	}
)

func TestCreateUser(t *testing.T) {
	err := testUser.Create()
	checkErr(t, err)
}

func TestAuthenticate(t *testing.T) {
	user, err := models.Authenticate(testUser.Email, password)
	checkErr(t, err)
	compare(t, username, user.Username)
	compare(t, email, user.Email)
}

func TestUpdateUser(t *testing.T) {
	testUser.Username = "updatedusername"
	testUser.Update()
	user, err := models.Authenticate(testUser.Email, password)
	checkErr(t, err)
	compare(t, user.Username, testUser.Username)
}

func TestSetPassword(t *testing.T) {
	testUser.SetPassword("updatedpass")
	user, err := models.Authenticate(testUser.Email, "updatedpass")
	checkErr(t, err)
	compare(t, user.Username, testUser.Username)
}

func TestCreateArtist(t *testing.T) {
	err := testArtist.Create()
	checkErr(t, err)
}

func TestCreateManga(t *testing.T) {
	testManga.ArtistID = testArtist.ID
	err := testManga.Create()
	checkErr(t, err)
}

func TestCreateChapter(t *testing.T) {
	testChapter.MangaID = testManga.ID
	testChapter.UploaderID = testUser.ID
	err := testChapter.Create()
	checkErr(t, err)
}

func TestCreateComment(t *testing.T) {
	testComment.ChapterID = testChapter.ID
	testComment.UserID = testUser.ID
	err := testComment.Create()
	checkErr(t, err)
}

func TestUpdateComment(t *testing.T) {
	testComment.Content = "This comment is updated."
	err := testComment.Update()
	checkErr(t, err)
	comm, err := models.GetCommentByID(testComment.ID)
	checkErr(t, err)
	compare(t, testComment.Content, comm.Content)
}

func TestUpdateChapter(t *testing.T) {
	testChapter.Title = "Chapter 1 - Black Swordsman"
	err := testChapter.Update()
	checkErr(t, err)
	chp, err := models.GetChapterByID(testChapter.ID)
	checkErr(t, err)
	compare(t, testChapter.Title, chp.Title)
}

func TestUpdateManga(t *testing.T) {
	testManga.AlternativeTitle = testManga.AlternativeTitle + ", İkinci falan"
	err := testManga.Update()
	checkErr(t, err)
	mng, err := models.GetMangaByID(testManga.ID)
	checkErr(t, err)
	compare(t, testManga.AlternativeTitle, mng.AlternativeTitle)
}

func TestUpdateArtist(t *testing.T) {
	testArtist.About = "Updated about of testArtist. zamazingo."
	err := testArtist.Update()
	checkErr(t, err)
	arts, err := models.GetArtistByID(testArtist.ID)
	checkErr(t, err)
	compare(t, testArtist.About, arts.About)
}

func TestDeleteComment(t *testing.T) {
	err := testComment.Delete()
	checkErr(t, err)
}

func TestDeleteChapter(t *testing.T) {
	err := testChapter.Delete()
	checkErr(t, err)
}

func TestDeleteManga(t *testing.T) {
	err := testManga.Delete()
	checkErr(t, err)
}

func TestDeleteArtist(t *testing.T) {
	err := testArtist.Delete()
	checkErr(t, err)
}

func TestDeleteUser(t *testing.T) {
	err := testUser.Delete()
	checkErr(t, err)
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("[ERROR] -> %s", err.Error())
	}
}

func compare[T comparable](t *testing.T, want T, got T) {
	if want != got {
		t.Errorf("Want %v, got %v.", want, got)
	}
}
