package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/model"
	"log"
)

type authPreparedStatements struct {
	GetAuthUserByUsername   *sql.Stmt
	GetBaidFromBaid         *sql.Stmt
	InsertAuthUser          *sql.Stmt
	GetUsernameFromBaid     *sql.Stmt
	GetPasswordHashFromBaid *sql.Stmt
	GetCustomTitleOn        *sql.Stmt
	UpdateCustomTitleOn     *sql.Stmt
	ChangeUsername          *sql.Stmt
	ChangePassword          *sql.Stmt
}

var db *sql.DB
var authStmts authPreparedStatements

func initAuthDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	authStmts.GetAuthUserByUsername = prepareQuery(db, "queries/auth/getAuthUserByUsername.sql")
	authStmts.GetBaidFromBaid = prepareQuery(db, "queries/auth/getBaidFromBaid.sql")
	authStmts.InsertAuthUser = prepareQuery(db, "queries/auth/insertAuthUser.sql")
	authStmts.GetUsernameFromBaid = prepareQuery(db, "queries/auth/getUsernameFromBaid.sql")
	authStmts.GetPasswordHashFromBaid = prepareQuery(db, "queries/auth/getPasswordHashFromBaid.sql")
	authStmts.GetCustomTitleOn = prepareQuery(db, "queries/auth/getCustomTitleOn.sql")
	authStmts.UpdateCustomTitleOn = prepareQuery(db, "queries/auth/updateCustomTitleOn.sql")
	authStmts.ChangeUsername = prepareQuery(db, "queries/auth/changeUsername.sql")
	authStmts.ChangePassword = prepareQuery(db, "queries/auth/changePassword.sql")

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
	baidRows, err := authStmts.GetBaidFromBaid.Query(baid)
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

func InsertAuthUser(user model.AuthUser) error {
	_, err := authStmts.InsertAuthUser.Exec(user.Baid, user.Username, user.PasswordHash)
	return err
}

func GetAuthUserByUsername(username string) (model.AuthUser, bool, error) {
	var user model.AuthUser
	err := authStmts.GetAuthUserByUsername.QueryRow(username).Scan(&user.Baid, &user.Username, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

func GetUsernameByBaid(baid uint) (string, bool, error) {
	var username string
	err := authStmts.GetUsernameFromBaid.QueryRow(baid).Scan(&username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false, nil
		}
		return "", false, err
	}
	return username, true, nil
}

func GetPasswordHashByBaid(baid uint) (string, bool, error) {
	var passwordHash string
	err := authStmts.GetPasswordHashFromBaid.QueryRow(baid).Scan(&passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false, nil
		}
		return "", false, err
	}
	return passwordHash, true, nil
}

func GetCustomTitleOn(baid uint) (bool, error) {
	var customTitleOn bool
	err := authStmts.GetCustomTitleOn.QueryRow(baid).Scan(&customTitleOn)
	if err != nil {
		return true, err // doesn't matter if return true or false here
	}
	return customTitleOn, nil
}

func UpdateCustomTitleOn(baid uint, customTitleOn bool) error {
	_, err := authStmts.UpdateCustomTitleOn.Exec(customTitleOn, baid)
	return err
}

func ChangeUsername(baid uint, username string) error {
	_, err := authStmts.ChangeUsername.Exec(username, baid)
	return err
}

func ChangePassword(baid uint, passwordHash string) error {
	_, err := authStmts.ChangePassword.Exec(passwordHash, baid)
	return err
}
