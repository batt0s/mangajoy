// database package is for connection with database
package database

import (
	"database/sql"
	"log"
	"strings"

	"github.com/batt0s/mangajoy/settings"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

// global DB variable
var DB *bun.DB

// Init database with given mode
func InitDB(mode string) error {
	var db *bun.DB
	mode = strings.ToLower(mode)
	switch mode {
	case "dev":
		dev := settings.DATABASES["dev"].(map[string]string)
		name := dev["name"]
		sqlite, err := sql.Open(sqliteshim.ShimName, name)
		if err != nil {
			return err
		}
		db = bun.NewDB(sqlite, sqlitedialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	case "test":
		sqldb, err := sql.Open(sqliteshim.ShimName, "test.db") //"flie::memory?cache=shared"
		if err != nil {
			return err
		}
		db = bun.NewDB(sqldb, sqlitedialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	default:
		name := settings.DATABASES["default"].(map[string]string)["name"]
		sqldb, err := sql.Open(sqliteshim.ShimName, name)
		if err != nil {
			return err
		}
		db = bun.NewDB(sqldb, sqlitedialect.New())
	}
	DB = db
	log.Println("Connected to database.")
	return nil
}
