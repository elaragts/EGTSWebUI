SELECT CurrentBody,
       CurrentFace,
       CurrentHead,
       CurrentKigurumi,
       CurrentPuchi,
       ColorBody,
       ColorFace,
       ColorLimb
FROM UserData
WHERE Baid = ?