package database_test

import (
	"log"
	"os"
	"testing"

	"github.com/batt0s/mangajoy/database"
)

func TestMain(m *testing.M) {
	log.Println("Starting testing package database.")
	exitVal := m.Run()
	log.Println("Done testing package database.")
	err := os.Remove("test.db")
	if err != nil {
		log.Println("Could not delete test.db")
	}
	os.Exit(exitVal)
}

func TestInitDb(t *testing.T) {
	err := database.InitDB("test")
	checkErr(t, err)
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("[ERROR] -> %s", err.Error())
	}
}
