CREATE TABLE IF NOT EXISTS users
(
    Baid         INTEGER PRIMARY KEY,
    username     TEXT not null,
    passwordHash TEXT not null
);