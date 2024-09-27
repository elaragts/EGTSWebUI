package model

type LeaderboardRecord struct {
	Name      string
	BestScore uint
	BestCrown uint
	BestRank  uint
}

type Stats struct {
	TotalUsers     uint `json:"totalUsers"`
	ActiveUsers    uint `json:"activeUsers"`
	TotalPlayCount uint `json:"totalPlayCount"`
}
