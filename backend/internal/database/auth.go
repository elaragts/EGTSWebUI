package database

import (
	"database/sql"
	"fmt"
	"log"
)

type authPreparedStatements struct {
	GetAuthUserByUsername *sql.Stmt
	GetAuthUserByBaid     *sql.Stmt
	InsertAuthUser        *sql.Stmt
}

var db *sql.DB
var authStmts authPreparedStatements

func initAuthDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	authStmts.GetAuthUserByUsername = prepareQuery(db, "internal/database/queries/auth/getAuthUserByUsername.sql")
	authStmts.GetAuthUserByBaid = prepareQuery(db, "internal/database/queries/auth/getAuthUserByBaid.sql")
	authStmts.InsertAuthUser = prepareQuery(db, "internal/database/queries/auth/insertAuthUser.sql")
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database")
}

const (
	USERNAMEFOUND = 0
	BAIDFOUND     = 1
)

func IsAuthUserUnique(username string, baid uint) (bool, uint, error) {
	usernameRows, err := authStmts.GetAuthUserByUsername.Query(username)
	if err != nil {
		return false, 0, err
	}
	defer usernameRows.Close()
	// Iterate over the rows
	if usernameRows.Next() {
		return false, USERNAMEFOUND, nil
	}
	baidRows, err := authStmts.GetAuthUserByBaid.Query(baid)
	if err != nil {
		return false, 0, err
	}
	defer baidRows.Close()
	// Iterate over the rows
	if baidRows.Next() {
		return false, BAIDFOUND, nil
	}
	return true, 0, nil
}

func InsertAuthUser(baid uint, username string, passwordHash string) error {
	_, err := authStmts.InsertAuthUser.Exec(baid, username, passwordHash)
	return err
}
