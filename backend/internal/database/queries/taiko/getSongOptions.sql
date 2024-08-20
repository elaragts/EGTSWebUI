SELECT OptionSetting,
       IsSkipOn,
       IsVoiceOn,
       SelectedToneId,
       NotesPosition
FROM UserData
WHERE Baid = ?