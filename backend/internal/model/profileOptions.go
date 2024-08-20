package model

type ProfileOptions struct {
	MyDonName                    string `json:"myDonName"`
	Title                        string `json:"title"`
	Language                     uint   `json:"language"`
	TitlePlateId                 uint   `json:"titlePlateId"`
	DisplayAchievement           bool   `json:"displayAchievement"`
	AchievementDisplayDifficulty uint   `json:"achievementDisplayDifficulty"`
	DisplayDan                   bool   `json:"displayDan"`
	DifficultySettingCourse      uint   `json:"difficultySettingCourse"`
	DifficultySettingStar        uint   `json:"difficultySettingStar"`
	DifficultySettingSort        uint   `json:"difficultySettingSort"`
	CustomTitleOn                bool   `json:"customTitleOn"`
}
