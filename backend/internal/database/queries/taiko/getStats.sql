SELECT
    (SELECT COUNT(Baid) FROM UserData) AS TotalUsers,
    (SELECT COUNT(DISTINCT Baid) FROM SongPlayData WHERE PlayTime >= DATETIME('now', '-1 month')) AS ActiveUsers,
    (SELECT COUNT(*) FROM SongPlayData WHERE Skipped = 0) AS TotalPlayCount;
