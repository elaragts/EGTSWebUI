package database

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"sync"
)

//go:embed queries/taiko/*.sql queries/auth/*.sql
var sqlFiles embed.FS

func prepareQuery(db *sql.DB, filename string) *sql.Stmt {
	content, err := fs.ReadFile(sqlFiles, filename)
	if err != nil {
		log.Fatalf("Could not read SQL file: %s, error: %v", filename, err)
	}

	stmt, err := db.Prepare(string(content))
	if err != nil {
		log.Fatalf("Could not prepare query from file: %s, error: %v", filename, err)
	}

	return stmt
}

var once sync.Once

func InitDBs(taikoDBPath string, authDBPath string) {
	once.Do(func() {
		initTaikoDB(taikoDBPath)
		initAuthDB(authDBPath)
	})
}
