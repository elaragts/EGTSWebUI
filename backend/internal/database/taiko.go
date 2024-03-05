package database

import (
	"database/sql"
	"fmt"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

var db *sql.DB
var leaderboardStmt *sql.Stmt

// Once ensures the database connection is initialized only once
var once sync.Once

func InitTaikoDB(dataSourceName string) {
	once.Do(func() {
		var err error
		db, err = sql.Open("sqlite3", dataSourceName)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}
		leaderboardStmt, err = db.Prepare(`
SELECT ud.MyDonName, sbd.BestScore, sbd.BestCrown, sbd.BestScoreRank
FROM SongBestData sbd
         INNER JOIN UserData ud ON sbd.Baid = ud.Baid
         INNER JOIN Card c ON sbd.Baid = c.Baid
WHERE SongID = ?
  AND Difficulty = ?
ORDER BY sbd.BestScore DESC
LIMIT 10`)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully connected to the database")
	})
}

func GetLeaderboard(songId uint, difficulty uint) ([]model.LeaderboardRecord, error) {
	rows, err := leaderboardStmt.Query(songId, difficulty)
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
