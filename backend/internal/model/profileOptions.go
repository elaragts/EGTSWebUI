package model

type ProfileOptions struct {
	MyDonName                    string `json:"myDonName"`
	Title                        string `json:"title"`
	Language                     uint   `json:"language"`
	TitlePlateId                 uint   `json:"titlePlateId"`
	DisplayAchievement           bool   `json:"displayAchievement"`
	AchievementDisplayDifficulty uint   `json:"achievementDisplayDifficulty"`
	DisplayDan                   bool   `json:"displayDan"`
	DifficultySettingArray       string `json:"difficultySettingArray"`
}
