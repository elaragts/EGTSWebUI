package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type taikoPreparedStatements struct {
	Leaderboard           *sql.Stmt
	GetBaidFromAccessCode *sql.Stmt
	GetPublicProfile      *sql.Stmt
	GetProfileOptions     *sql.Stmt
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
	taikoStmts.Leaderboard = prepareQuery(taikodb, "queries/taiko/leaderboard.sql")
	taikoStmts.GetBaidFromAccessCode = prepareQuery(taikodb, "queries/taiko/getBaidFromAccessCode.sql")
	taikoStmts.GetPublicProfile = prepareQuery(taikodb, "queries/taiko/getPublicProfile.sql")
	taikoStmts.GetProfileOptions = prepareQuery(taikodb, "queries/taiko/getProfileOptions.sql")
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

func GetPublicProfile(baid uint) (model.PublicProfile, error) {
	var profile model.PublicProfile
	var costumeDataStr string

	err := taikoStmts.GetPublicProfile.QueryRow(sql.Named("baid", baid)).Scan(
		&profile.MyDonName,
		&profile.Title,
		&profile.AchievementDisplayDifficulty,
		&costumeDataStr, // Costume data as a JSON string
		&profile.ColorBody,
		&profile.ColorFace,
		&profile.PlayCount,
		&profile.DanId,
		&profile.ClearState,
		&profile.BestScoreRank[0], // Directly scanning into the array elements
		&profile.BestScoreRank[1],
		&profile.BestScoreRank[2],
		&profile.BestScoreRank[3],
		&profile.BestScoreRank[4],
		&profile.BestScoreRank[5],
		&profile.BestScoreRank[6],
		&profile.BestScoreRank[7],
		&profile.BestCrown[0],
		&profile.BestCrown[1],
		&profile.BestCrown[2],
	)

	if err != nil {
		return model.PublicProfile{}, err
	}

	err = json.Unmarshal([]byte(costumeDataStr), &profile.CostumeData)
	if err != nil {
		return model.PublicProfile{}, err
	}
	return profile, nil
}

func GetProfileOptions(baid uint) (model.ProfileOptions, error) {
	var profileOptions model.ProfileOptions

	err := taikoStmts.GetProfileOptions.QueryRow(baid).Scan(
		&profileOptions.MyDonName,
		&profileOptions.Title,
		&profileOptions.Language,
		&profileOptions.TitlePlateId,
		&profileOptions.DisplayAchievement,
		&profileOptions.AchievementDisplayDifficulty,
		&profileOptions.DisplayDan,
		&profileOptions.DifficultySettingArray,
	)

	if err != nil {
		return model.ProfileOptions{}, err
	}

	return profileOptions, nil
}
