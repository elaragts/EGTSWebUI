SELECT MyDonName,
       Title,
       MyDonNameLanguage,
       TitlePlateId,
       DisplayAchievement,
       AchievementDisplayDifficulty,
       DisplayDan,
       DifficultySettingArray
FROM UserData
WHERE Baid = ?

