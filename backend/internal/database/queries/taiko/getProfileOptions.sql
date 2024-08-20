SELECT MyDonName,
       Title,
       MyDonNameLanguage,
       TitlePlateId,
       DisplayAchievement,
       AchievementDisplayDifficulty,
       DisplayDan,
       DifficultySettingCourse,
       DifficultySettingSort,
       DifficultySettingStar
FROM UserData
WHERE Baid = ?