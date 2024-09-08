package model

import "database/sql"

type PublicProfile struct {
	MyDonName                    string        `json:"myDonName"`
	Title                        string        `json:"title"`
	AchievementDisplayDifficulty int           `json:"achievementDisplayDifficulty"`
	CostumeData                  [5]int        `json:"costumeData"`
	ColorBody                    int           `json:"colorBody"`
	ColorFace                    int           `json:"colorFace"`
	PlayCount                    int           `json:"playCount"`
	DanId                        sql.NullInt32 `json:"danId"`
	ClearState                   sql.NullInt32 `json:"clearState"`
	BestScoreRank                [8]int        `json:"bestScoreRank"`
	BestCrown                    [3]int        `json:"bestCrown"`
}

type ProfileSettings struct {
	ProfileOptions ProfileOptions `json:"profileOptions"`
	CostumeOptions CostumeOptions `json:"costumeOptions"`
	SongOptions    SongOptions    `json:"songOptions"`
}

type AccessCode struct {
	AccessCode string `json:"accessCode"`
}

type FavouritedSong struct {
	SongId uint `json:"songId"`
}

