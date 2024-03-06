SELECT Baid, username, passwordHash
FROM users
WHERE username = ?
LIMIT 1