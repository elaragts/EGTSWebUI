package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type taikoPreparedStatements struct {
	Leaderboard           *sql.Stmt
	GetBaidFromAccessCode *sql.Stmt
}

var taikodb *sql.DB
var taikoStmts taikoPreparedStatements

// Once ensures the database connection is initialized only once

func initTaikoDB(dataSourceName string) {
	var err error
	taikodb, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	taikoStmts.Leaderboard = prepareQuery(taikodb, "internal/database/queries/taiko/leaderboard.sql")
	taikoStmts.GetBaidFromAccessCode = prepareQuery(taikodb, "internal/database/queries/taiko/getBaidFromAccessCode.sql")
	if err = taikodb.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database")
}

func GetLeaderboard(songId uint, difficulty uint) ([]model.LeaderboardRecord, error) {
	rows, err := taikoStmts.Leaderboard.Query(songId, difficulty)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ret := []model.LeaderboardRecord{}

	// Iterate over the rows
	for rows.Next() {
		var row model.LeaderboardRecord
		var err = rows.Scan(&row.Name, &row.BestScore, &row.BestCrown, &row.BestRank)
		if err != nil {
			return nil, err
		}

		// Append the user to the slice
		ret = append(ret, row)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func GetBaidFromAccessCode(accessCode string) (uint, bool, error) {
	var baid uint
	err := taikoStmts.GetBaidFromAccessCode.QueryRow(accessCode).Scan(&baid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return baid, true, nil
}
