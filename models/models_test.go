package models_test

import (
	"context"
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
	err = database.DB.ResetModel(context.Background(), (*models.User)(nil))
	if err != nil {
		log.Println("Cannot reset model user")
		os.Exit(1)
	}
	exitVal := m.Run()
	log.Println("Done testing package models.")
	err = os.Remove("test.db")
	if err != nil {
		log.Println("Could not delete test.db")
	}
	os.Exit(exitVal)
}

var (
	username string       = "testuser"
	email    string       = "test@test.com"
	password string       = "testpass"
	testUser *models.User = &models.User{
		Username: username,
		Email:    email,
		Password: password,
		IsAdmin:  false,
		IsStaff:  false,
	}
)

func TestSaveUser(t *testing.T) {
	err := testUser.Save()
	checkErr(t, err)
}

func TestAuthenticate(t *testing.T) {
	user, err := models.Authenticate(testUser.Email, password)
	checkErr(t, err)
	compare(t, username, user.Username)
	compare(t, email, user.Email)
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
