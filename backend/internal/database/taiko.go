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
	GetCostumeOptions     *sql.Stmt
	GetSongOptions        *sql.Stmt
	UpdateUser            *sql.Stmt
	GetAccessCodes        *sql.Stmt
	AddAccessCode         *sql.Stmt
	DeleteAccessCode      *sql.Stmt
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
	taikoStmts.GetCostumeOptions = prepareQuery(taikodb, "queries/taiko/getCostumeOptions.sql")
	taikoStmts.GetSongOptions = prepareQuery(taikodb, "queries/taiko/getSongOptions.sql")
	taikoStmts.UpdateUser = prepareQuery(taikodb, "queries/taiko/updateUser.sql")
	taikoStmts.GetAccessCodes = prepareQuery(taikodb, "queries/taiko/getAccessCodes.sql")
	taikoStmts.AddAccessCode = prepareQuery(taikodb, "queries/taiko/addAccessCode.sql")
	taikoStmts.DeleteAccessCode = prepareQuery(taikodb, "queries/taiko/deleteAccessCode.sql")

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
	var ret []model.LeaderboardRecord

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
		&profileOptions.DifficultySettingCourse,
		&profileOptions.DifficultySettingSort,
		&profileOptions.DifficultySettingStar,
	)

	if err != nil {
		return model.ProfileOptions{}, err
	}

	return profileOptions, nil
}

func GetCostumeOptions(baid uint) (model.CostumeOptions, error) {
	var costumeOptions model.CostumeOptions

	err := taikoStmts.GetCostumeOptions.QueryRow(baid).Scan(
		&costumeOptions.CurrentBody,
		&costumeOptions.CurrentFace,
		&costumeOptions.CurrentHead,
		&costumeOptions.CurrentKigurumi,
		&costumeOptions.CurrentPuchi,
		&costumeOptions.ColorBody,
		&costumeOptions.ColorFace,
		&costumeOptions.ColorLimb,
	)

	if err != nil {
		return model.CostumeOptions{}, err
	}

	return costumeOptions, nil
}

func GetSongOptions(baid uint) (model.SongOptions, error) {
	var songOptions model.SongOptions

	var optionSetting uint8

	err := taikoStmts.GetSongOptions.QueryRow(baid).Scan(
		&optionSetting,
		&songOptions.IsSkipOn,
		&songOptions.IsVoiceOn,
		&songOptions.SelectedToneId,
		&songOptions.NotesPosition,
	)

	if err != nil {
		return model.SongOptions{}, err
	}

	// OptionSetting is an 8-bit uint where when counting bits LEFT TO RIGHT:
	// bits 1-4 are speedId,
	// bit 5 is isVanish bool, bit 6 is isInverse bool
	// and bits 7-8 are randomId where 01 = messy and 10 = whimsical and 00 = none
	songOptions.SpeedId = uint((optionSetting & 0b11110000) >> 4) // first 4 bits on the left
	songOptions.IsVanishOn = (optionSetting & 0b00001000) != 0    // 5th bit from the left
	songOptions.IsInverseOn = (optionSetting & 0b00000100) != 0   // 6th bit from the left
	songOptions.RandomId = uint(optionSetting & 0b00000011)       // last 2 bits

	return songOptions, nil
}

func UpdateUser(baid uint, profileSettings model.ProfileSettings) error {

	var optionSetting uint8

	// note that bits are being counted left to right
	optionSetting = uint8(profileSettings.SongOptions.SpeedId << 4) // storing speed id in the first 4 bits
	if profileSettings.SongOptions.IsVanishOn {
		optionSetting |= 0b00001000
		// storing vanish boolean in the 5th bit
	}
	if profileSettings.SongOptions.IsInverseOn {
		optionSetting |= 0b00000100
		// storing inverse boolean in the 6th bit
	}
	optionSetting |= uint8(profileSettings.SongOptions.RandomId) // storing random id in the last 2 bits

	_, err := taikoStmts.UpdateUser.Exec(
		profileSettings.ProfileOptions.MyDonName,
		profileSettings.ProfileOptions.Title,
		profileSettings.ProfileOptions.Language,
		profileSettings.ProfileOptions.TitlePlateId,
		profileSettings.ProfileOptions.DisplayAchievement,
		profileSettings.ProfileOptions.AchievementDisplayDifficulty,
		profileSettings.ProfileOptions.DisplayDan,
		profileSettings.ProfileOptions.DifficultySettingCourse,
		profileSettings.ProfileOptions.DifficultySettingStar,
		profileSettings.ProfileOptions.DifficultySettingSort,
		profileSettings.CostumeOptions.CurrentBody,
		profileSettings.CostumeOptions.CurrentFace,
		profileSettings.CostumeOptions.CurrentHead,
		profileSettings.CostumeOptions.CurrentKigurumi,
		profileSettings.CostumeOptions.CurrentPuchi,
		profileSettings.CostumeOptions.ColorBody,
		profileSettings.CostumeOptions.ColorFace,
		profileSettings.CostumeOptions.ColorLimb,
		profileSettings.SongOptions.IsSkipOn,
		profileSettings.SongOptions.IsVoiceOn,
		profileSettings.SongOptions.SelectedToneId,
		profileSettings.SongOptions.NotesPosition,
		optionSetting,
		baid,
	)

	return err
}

func GetAccessCodes(baid uint) ([]string, error) {

	var accessCodes []string

	rows, err := taikoStmts.GetAccessCodes.Query(baid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accessCode string
		if err := rows.Scan(&accessCode); err != nil {
			return nil, err
		}
		accessCodes = append(accessCodes, accessCode)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accessCodes, nil
}

func AddAccessCode(baid uint, accessCode string) error {

	_, err := taikoStmts.AddAccessCode.Exec(
		accessCode,
		baid,
	)

	return err
}

func DeleteAccessCode(baid uint, accessCode string) error {

	// sending with baid just in case
	_, err := taikoStmts.DeleteAccessCode.Exec(
		accessCode,
		baid,
	)

	return err
}
